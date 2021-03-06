package eet

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/beevik/etree"
	"github.com/chutommy/eetgateway/pkg/wsse"
)

// ErrInvalidXMLDigest is returned if the referenced computed digest differs from the digest in the XML.
var ErrInvalidXMLDigest = errors.New("computed digest differs from the digest in the XML")

// ErrInvalidSOAPMessage is returned if a SOAP message has an unexpected structure.
var ErrInvalidSOAPMessage = errors.New("SOAP message with an unexpected structure")

// ErrInvalidUUID is returned if the response UUID is different, or it has invalid format.
var ErrInvalidUUID = errors.New("invalid UUID value or format")

// ErrInvalidBKP is returned if the response BKP code is different.
var ErrInvalidBKP = errors.New("incorrect response BKP")

// NewRequestEnvelope returns a populated and signed SOAP request envelope.
func NewRequestEnvelope(t *TrzbaType, cert *x509.Certificate, pk *rsa.PrivateKey) ([]byte, error) {
	if err := t.setSecurityCodes(pk); err != nil {
		return nil, fmt.Errorf("setting security codes: %w", err)
	}

	trzba, err := t.etree()
	if err != nil {
		return nil, fmt.Errorf("marshal trzba to etree element: %w", err)
	}

	binCert, err := certToB64(cert)
	if err != nil {
		return nil, fmt.Errorf("convert certificate to base64: %w", err)
	}

	// build request message
	body := buildBodyElem()
	body.AddChild(trzba)
	env := getSoapEnvelope()
	envElem, err := findElement(env.Root(), ".")
	if err != nil {
		return nil, err
	}

	envElem.AddChild(body)

	// set binary security token
	tokenElem, err := findElement(envElem, "./Header/Security/BinarySecurityToken")
	if err != nil {
		return nil, err
	}

	tokenElem.SetText(string(binCert))

	// set signature
	signature, err := findElement(env.Root(), "./Header/Security/Signature")
	if err != nil {
		return nil, err
	}

	if err = setDigestVal(body, signature); err != nil {
		return nil, fmt.Errorf("set digest value: %w", err)
	}

	if err = setSignatureVal(pk, signature); err != nil {
		return nil, fmt.Errorf("set signature value: %w", err)
	}

	signedEnv, err := env.WriteToBytes()
	if err != nil {
		return nil, fmt.Errorf("parse envelope document to bytes: %w", err)
	}

	return signedEnv, nil
}

func certToB64(cert *x509.Certificate) ([]byte, error) {
	b64 := make([]byte, base64.StdEncoding.EncodedLen(len(cert.Raw)))
	base64.StdEncoding.Encode(b64, cert.Raw)
	return b64, nil
}

func setDigestVal(body *etree.Element, sign *etree.Element) error {
	digest, err := wsse.CalcDigest(body)
	if err != nil {
		return fmt.Errorf("calculate digest of the body element: %w", err)
	}

	digestVal := base64.StdEncoding.EncodeToString(digest)
	digestValElem, err := findElement(sign, "./SignedInfo/Reference/DigestValue")
	if err != nil {
		return err
	}

	digestValElem.SetText(digestVal)

	return nil
}

func setSignatureVal(pk *rsa.PrivateKey, signature *etree.Element) error {
	signedInfo, err := findElement(signature, "./SignedInfo")
	if err != nil {
		return err
	}

	rawSig, err := wsse.CalcSignature(pk, signedInfo.Copy())
	if err != nil {
		return fmt.Errorf("calculate signature value: %w", err)
	}

	sigVal := base64.StdEncoding.EncodeToString(rawSig)
	sigValElem, err := findElement(signature, "./SignatureValue")
	if err != nil {
		return err
	}

	sigValElem.SetText(sigVal)

	return nil
}

// OdpovedBody represents a SOAP Body of the response envelope.
type OdpovedBody struct {
	Odpoved OdpovedType `xml:"Odpoved"`
}

// ParseResponseEnvelope returns a parsed SOAP response envelope.
func ParseResponseEnvelope(env []byte) (*OdpovedType, error) {
	doc := etree.NewDocument()
	err := doc.ReadFromBytes(env)
	if err != nil {
		return nil, fmt.Errorf("parse envelope to etree: %w", err)
	}

	bodyElem, err := findElement(doc.Root(), "./Body")
	if err != nil {
		return nil, err
	}

	doc.SetRoot(bodyElem.Copy())
	odpovedBytes, err := doc.WriteToBytes()
	if err != nil {
		return nil, fmt.Errorf("serialize etree document to bytes: %w", err)
	}

	var odpoved OdpovedBody
	err = xml.Unmarshal(odpovedBytes, &odpoved)
	if err != nil {
		return nil, fmt.Errorf("decode odpoved bytes: %w", err)
	}

	return &odpoved.Odpoved, nil
}

// VerifyResponse checks whether the response envelope is valid and is signed by a trusted certificate.
func VerifyResponse(trzba *TrzbaType, respEnv []byte, odpoved *OdpovedType, verifyCert func(*x509.Certificate) error) error {
	envelope := etree.NewDocument()
	err := envelope.ReadFromBytes(respEnv)
	if err != nil {
		return fmt.Errorf("parse envelope to etree: %w", err)
	}

	// verify only if successful
	if !trzba.Hlavicka.Overeni && odpoved.Chyba.Kod == 0 {
		if trzba.Hlavicka.Uuidzpravy != odpoved.Hlavicka.Uuidzpravy {
			return fmt.Errorf("different uuid: %w", ErrInvalidUUID)
		}

		if trzba.KontrolniKody.Bkp.BkpType != odpoved.Hlavicka.Bkp {
			return fmt.Errorf("different bkp: %w", ErrInvalidBKP)
		}

		if err := checkDigSig(envelope); err != nil {
			return fmt.Errorf("check digital signature: %w", err)
		}

		if err := verifyCertificate(envelope, verifyCert); err != nil {
			return fmt.Errorf("check certificate: %w", err)
		}
	}

	return nil
}

func verifyCertificate(envelope *etree.Document, verifyCert func(*x509.Certificate) error) error {
	tokenElem, err := findElement(envelope.Root(), "./Header/Security/BinarySecurityToken")
	if err != nil {
		return err
	}

	tokenB64 := tokenElem.Text()
	rawCert, err := base64.StdEncoding.DecodeString(tokenB64)
	if err != nil {
		return fmt.Errorf("decode binary security token from base64 encoding: %w", err)
	}

	cert, err := x509.ParseCertificate(rawCert)
	if err != nil {
		return fmt.Errorf("parse raw certificate: %w", err)
	}

	if err = verifyCert(cert); err != nil {
		return fmt.Errorf("verify security token: %w", err)
	}

	return nil
}

func checkDigSig(envelope *etree.Document) error {
	if err := validateDigestValue(envelope); err != nil {
		return fmt.Errorf("invalid digest value: %w", err)
	}

	if err := verifySignature(envelope); err != nil {
		return fmt.Errorf("verify signature: %w", err)
	}

	return nil
}

func verifySignature(envelope *etree.Document) error {
	cert, err := getCertFromToken(envelope)
	if err != nil {
		return fmt.Errorf("retrieve certificate from the binary security token: %w", err)
	}

	sigVal, digest, err := getSignatureDate(envelope)
	if err != nil {
		return fmt.Errorf("retrieve signature value and digest: %w", err)
	}

	err = rsa.VerifyPKCS1v15(cert.PublicKey.(*rsa.PublicKey), crypto.SHA256, digest, sigVal)
	if err != nil {
		return fmt.Errorf("verify PKCS1v15 signature: %w", err)
	}

	return nil
}

func getSignatureDate(envelope *etree.Document) (sigVal []byte, digest []byte, err error) {
	sigElem, err := findElement(envelope.Root(), "./Header/Security/Signature")
	if err != nil {
		return nil, nil, err
	}

	// signature value
	{
		sigValElem, err := findElement(sigElem, "./SignatureValue")
		if err != nil {
			return nil, nil, err
		}

		sigValB64 := sigValElem.Text()
		sigVal, err = base64.StdEncoding.DecodeString(sigValB64)
		if err != nil {
			return nil, nil, fmt.Errorf("decode base64 signature value: %w", err)
		}
	}

	// digest value
	{
		signedInfo, err := findElement(sigElem, "./SignedInfo")
		if err != nil {
			return nil, nil, err
		}

		// namespace for the XML digital signature
		signedInfo.CreateAttr("xmlns", "http://www.w3.org/2000/09/xmldsig#")

		digest, err = wsse.CalcDigest(signedInfo)
		if err != nil {
			return nil, nil, fmt.Errorf("calculate digest value of signed info: %w", err)
		}
	}

	return sigVal, digest, nil
}

func getCertFromToken(envelope *etree.Document) (*x509.Certificate, error) {
	tokenElem, err := findElement(envelope.Root(), "./Header/Security/BinarySecurityToken")
	if err != nil {
		return nil, err
	}

	token := tokenElem.Text()
	rawCert, _ := base64.StdEncoding.DecodeString(token)
	cert, err := x509.ParseCertificate(rawCert)
	if err != nil {
		return nil, fmt.Errorf("parse x509 certificate: %w", err)
	}

	return cert, nil
}

func validateDigestValue(envelope *etree.Document) error {
	digest, err := bodyDigest(envelope)
	if err != nil {
		return fmt.Errorf("get digest value of the body element: %w", err)
	}

	digestValElem, err := findElement(envelope.Root(), "./Header/Security/Signature/SignedInfo/Reference/DigestValue")
	if err != nil {
		return err
	}

	expDigestVal := base64.StdEncoding.EncodeToString(digest)
	actDigestVal := digestValElem.Text()

	if expDigestVal != actDigestVal {
		return ErrInvalidXMLDigest
	}

	return nil
}

func bodyDigest(envelope *etree.Document) ([]byte, error) {
	bodyElem, err := findElement(envelope.Root(), "./Body")
	if err != nil {
		return nil, err
	}

	// set namespaces
	bodyElem.CreateAttr("xmlns:eet", "http://fs.mfcr.cz/eet/schema/v3")
	bodyElem.CreateAttr("xmlns:soapenv", "http://schemas.xmlsoap.org/soap/envelope/")
	bodyElem.CreateAttr("xmlns:wsu", "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd")

	digest, _ := wsse.CalcDigest(bodyElem)

	return digest, nil
}

func findElement(root *etree.Element, path string) (*etree.Element, error) {
	e := root.FindElement(path)
	if e == nil {
		return nil, fmt.Errorf("element in %s of %s element not found: %w", path, root.FullTag(), ErrInvalidSOAPMessage)
	}

	return e, nil
}

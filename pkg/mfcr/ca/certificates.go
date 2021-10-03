package ca

const (
	// OrganizationName is the legal name that the organization is registered with authority at the national level.
	OrganizationName = "Česká republika - Generální finanční ředitelství"
)

var (
	// ICACertificate is the certificate of subordinate CA for issuing qualified
	// certificates in PEM format.
	// CN = I.CA QUALIFIED 2 CA/RSA 02/2016; SN:  100001006 (5F5E4EE HEX)
	// https://www.ica.cz/HCA-qualificate
	ICACertificate = []byte(`
-----BEGIN CERTIFICATE-----
MIIHpTCCBY2gAwIBAgIEBfXk7jANBgkqhkiG9w0BAQsFADBwMQswCQYDVQQGEwJD
WjEtMCsGA1UECgwkUHJ2bsOtIGNlcnRpZmlrYcSNbsOtIGF1dG9yaXRhLCBhLnMu
MRkwFwYDVQQDDBBJLkNBIFJvb3QgQ0EvUlNBMRcwFQYDVQQFEw5OVFJDWi0yNjQz
OTM5NTAeFw0xNjAyMTExMjE3MTFaFw0yNjAyMDgxMjE3MTFaMH8xCzAJBgNVBAYT
AkNaMSgwJgYDVQQDDB9JLkNBIFF1YWxpZmllZCAyIENBL1JTQSAwMi8yMDE2MS0w
KwYDVQQKDCRQcnZuw60gY2VydGlmaWthxI1uw60gYXV0b3JpdGEsIGEucy4xFzAV
BgNVBAUTDk5UUkNaLTI2NDM5Mzk1MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIIC
CgKCAgEAyMALHP/YeXEtOEHHJXTrhOrZWZ5SeX3f8phlvUiCIxoJt2yZ4CI2Y26t
SuD+Ard7c539buJlzoZnuFs6xswvSVJwpwoKF3pflZ5DZjyqUhPpDZdEXQyne1U9
uo1T9wD1WWKQ/yONzKcawxfH2tr0ourILIjVtB6I99u5uA7flA/mynGucR1C4PC9
WbY4MrRV+YkSAzWb88K1wyhVZ0Tq50+jINrL8xCGzRNLSPbMw9lBsWNPfcom2ajP
bmIfyaf3uMBGNdNxUjQoiBjC0mYWkrEd95K6S0dkOA8KgelI/3Kyut/kxc1RsLXg
Io0DNSQ9F38q2I8KWpmxm2sOAHBR191fNEwhnfomCi1jjx6nHpIhHR1Vs5KcjL6z
8Qr42otM55qtEBhOnM1juPZs5+GYjpHG08e9cATWBC3GLd59hN6uSdZjNSb6LVg0
hB194Jb29WpaNj0wzEx98zR1W4NQy+EXSaBfj8bb7UZrxtSoJzF2YMNAPb/oYlRV
NuP4tmnUsW3m6r09j7cltBXCo/YfXDRX0rWNlJ7p+gDRHU1+nPlih6LWgyI/yrhJ
qGg4dg63YyywvuuoDI0zfjlhBSkqQymNwNelg1mDcEFUVxk8LKzXPXJlFNEt33+q
T+CMXlR+IkUC0jOI1SZV3uwcAwgbQWazNljKFpoJjGXc4fwh2A8CAwEAAaOCAjYw
ggIyMIHXBgNVHSAEgc8wgcwwgckGBFUdIAAwgcAwgb0GCCsGAQUFBwICMIGwGoGt
VGVudG8ga3ZhbGlmaWtvdmFueSBzeXN0ZW1vdnkgY2VydGlmaWthdCBieWwgdnlk
YW4gcG9kbGUgemFrb25hIDIyNy8yMDAwIFNiLiB2IHBsYXRuZW0gem5lbmkvVGhp
cyBxdWFsaWZpZWQgc3lzdGVtIGNlcnRpZmljYXRlIHdhcyBpc3N1ZWQgYWNjb3Jk
aW5nIHRvIEFjdCBOby4gMjI3LzIwMDAgQ29sbC4wEgYDVR0TAQH/BAgwBgEB/wIB
ADAOBgNVHQ8BAf8EBAMCAQYwHQYDVR0OBBYEFHSCCJHj2WRocYXW6zHkct+LJrFt
MB8GA1UdIwQYMBaAFHa5A0j71RihoTeg7cxogkxSNDYNMIGMBgNVHR8EgYQwgYEw
KaAnoCWGI2h0dHA6Ly9xY3JsZHAxLmljYS5jei9yY2ExNV9yc2EuY3JsMCmgJ6Al
hiNodHRwOi8vcWNybGRwMi5pY2EuY3ovcmNhMTVfcnNhLmNybDApoCegJYYjaHR0
cDovL3FjcmxkcDMuaWNhLmN6L3JjYTE1X3JzYS5jcmwwYwYIKwYBBQUHAQEEVzBV
MCkGCCsGAQUFBzAChh1odHRwOi8vci5pY2EuY3ovcmNhMTVfcnNhLmNlcjAoBggr
BgEFBQcwAYYcaHR0cDovL29jc3AuaWNhLmN6L3JjYTE1X3JzYTANBgkqhkiG9w0B
AQsFAAOCAgEAELc0LBZqU0XQuG/F43zqtPRspgixVwl4TQBW+9uQXPz0Og3C2Qf7
FHZwlB93EXz9D4jxQwffA0fugp/eRu6eZ6v55tR7M5Vvl3rlBPFVlDs1+8CWLABL
tX61hcXslU1Sdtqi6lGab9pDoBMdvLOky/CLMdQvA01XMEjCUIslT+U6UlCUhGG3
Oh/KBqIORdFcWaseoInsJrBpiAA8+wohMKZGomKSXYlUtuwywZ/GNrkHhJd5nN7a
uEDnM39uAYINSeQ7pHYFtyb4Xik8jOsk5LaQcgC/yOOcVVcZhmPJFamwA+xBhJY+
ynoB7cJyLx2IxiO/7PHSBNsobUaFobfAVNJgoY8X+FYmlcGv5526v8dHH6FEdyq/
0mxeXlFpqLrscfJj4zWNcs8+zmrphCrRgeWrrZkciJ+f6tceW+hdDYtpoHDhpJHn
UJRqc2R67x88t55DL9vjcbGNB8CTOthlOUv1UWzmIVO0FOEomUKy7d6cf4g2qbF6
Fbq9I3WzkYyxlizNmEAFVDhT2YdK19lWK8dlabxjIH9KF1yuhIG71NJWM6EVz905
8ebJcfPdpTRhNkZd+X84+YeFDsxYtOd8Q+L3CmX2Xzj9GrssN9ewTVeW7acSLa5g
cdzAiTUF92rQUfuVwr0zGuvZLnsoLIIsaWrx+pgHcBnL49PVJQV5w4c=
-----END CERTIFICATE-----`)

	// CAEET1Playground is a root certificate of the Certificate Authority EET for development purposes.
	// The certificate is valid from September 29, 2016 to September 29, 2022.
	// It can be downloaded at https://www.etrzby.cz/cs/certifikacni-autorita-EET.
	CAEET1Playground = []byte(`
-----BEGIN CERTIFICATE-----
MIIE2TCCA8GgAwIBAgIFAIPD9YUwDQYJKoZIhvcNAQELBQAwdzESMBAGCgmSJomT
8ixkARkWAkNaMUMwQQYDVQQKDDrEjGVza8OhIFJlcHVibGlrYSDigJMgR2VuZXLD
oWxuw60gZmluYW7EjW7DrSDFmWVkaXRlbHN0dsOtMRwwGgYDVQQDExNFRVQgQ0Eg
MSBQbGF5Z3JvdW5kMB4XDTE2MDkyOTE5NTQ0MFoXDTIyMDkyOTE5NTQ0MFowdzES
MBAGCgmSJomT8ixkARkWAkNaMUMwQQYDVQQKDDrEjGVza8OhIFJlcHVibGlrYSDi
gJMgR2VuZXLDoWxuw60gZmluYW7EjW7DrSDFmWVkaXRlbHN0dsOtMRwwGgYDVQQD
ExNFRVQgQ0EgMSBQbGF5Z3JvdW5kMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB
CgKCAQEA1IyU2zS/gm+66J9H5mL5W5071y88EF0f4X440TXuCjNvwdjvQhaQy2mw
m5+hG3RnuavQnJQOoIi532XLJNTawzP23MUtChjoa0B4ngAbnSRXXsjSscJde+eP
U8WKkxmxfd5BeuW4sHFh4CukI1UmwDs3cLy4BQ3ec0tYmn+HzQ+xzOgTO8EmDdr5
oTsoxV0TSoiIQVaeS+p5Qohx9ZUsB6H2oyg68GCSk/otZUo8wz71LW2bWNTxvDvx
R6YPpnKtQ+j9FNU9JeX76a3vEZ578DbcGrS/0iCZ4sTzrU47zBmdhn4mIJ3yAB6U
0y4dSKVd6TMqldVy5h6ep6hTQoUFdwIDAQABo4IBajCCAWYwEgYDVR0TAQH/BAgw
BgEB/wIBADAgBgNVHQ4BAf8EFgQUfDB2rMzWh9HsyR/icAgs41/eDAcwDgYDVR0P
AQH/BAQDAgEGMB8GA1UdIwQYMBaAFHwwdqzM1ofR7Mkf4nAILONf3gwHMGMGA1Ud
IARcMFowWAYKYIZIAWUDAgEwATBKMEgGCCsGAQUFBwICMDwMOlRlbnRvIGNlcnRp
Zmlrw6F0IGJ5bCB2eWTDoW4gcG91emUgcHJvIHRlc3RvdmFjw60gw7rEjWVseS4w
gZcGA1UdHwSBjzCBjDCBiaCBhqCBg4YpaHR0cDovL2NybC5jYTEtcGcuZWV0LmN6
L2VldGNhMXBnL2FsbC5jcmyGKmh0dHA6Ly9jcmwyLmNhMS1wZy5lZXQuY3ovZWV0
Y2ExcGcvYWxsLmNybIYqaHR0cDovL2NybDMuY2ExLXBnLmVldC5jei9lZXRjYTFw
Zy9hbGwuY3JsMA0GCSqGSIb3DQEBCwUAA4IBAQBHieHV+n7agbBRYYzHbWruqi1i
F7dX1g8cotPPg590FfQEAhK+Nwef8/sPNeo8gT99idzyRSq60c2f1nVlca+5W7YV
jUV2KrVqbE+1Ku4GT9K/ZFW6yyIOSeHBkCCjjoNJYLVBFgJMeepSoHFYsNk0pzzZ
g7Amemh0kxxd4YcxcxZHe0o2tNzdcUJ6pQxgwOYI67uepsBSor30qXTneAouMqLY
QHHc8v6JsMXFzrvg2tDAtQzNC3Ibsquw+Sur6ItgYMmkmOk9WfK33q7lUfXx34X5
F9OTF6UdKfXkvt+NmW7ayYwd+F4+3pfFr3wvBNdrG6tm/SUZBQ41Tt4OTKWg
-----END CERTIFICATE-----`)

	// CAEET1Playground2025 is a root certificate of the Certificate Authority EET for development purposes.
	// The certificate is valid from September 29, 2016 to September 29, 2025.
	// It can be downloaded at https://www.etrzby.cz/cs/certifikacni-autorita-EET.
	CAEET1Playground2025 = []byte(`
-----BEGIN CERTIFICATE-----
MIIE2TCCA8GgAwIBAgIFAKXDnhgwDQYJKoZIhvcNAQELBQAwdzESMBAGCgmSJomT
8ixkARkWAkNaMUMwQQYDVQQKDDrEjGVza8OhIFJlcHVibGlrYSDigJMgR2VuZXLD
oWxuw60gZmluYW7EjW7DrSDFmWVkaXRlbHN0dsOtMRwwGgYDVQQDExNFRVQgQ0Eg
MSBQbGF5Z3JvdW5kMB4XDTE2MDkyOTE5NTQ0MFoXDTI1MDkyOTEwMDAwMFowdzES
MBAGCgmSJomT8ixkARkWAkNaMUMwQQYDVQQKDDrEjGVza8OhIFJlcHVibGlrYSDi
gJMgR2VuZXLDoWxuw60gZmluYW7EjW7DrSDFmWVkaXRlbHN0dsOtMRwwGgYDVQQD
ExNFRVQgQ0EgMSBQbGF5Z3JvdW5kMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB
CgKCAQEA1IyU2zS/gm+66J9H5mL5W5071y88EF0f4X440TXuCjNvwdjvQhaQy2mw
m5+hG3RnuavQnJQOoIi532XLJNTawzP23MUtChjoa0B4ngAbnSRXXsjSscJde+eP
U8WKkxmxfd5BeuW4sHFh4CukI1UmwDs3cLy4BQ3ec0tYmn+HzQ+xzOgTO8EmDdr5
oTsoxV0TSoiIQVaeS+p5Qohx9ZUsB6H2oyg68GCSk/otZUo8wz71LW2bWNTxvDvx
R6YPpnKtQ+j9FNU9JeX76a3vEZ578DbcGrS/0iCZ4sTzrU47zBmdhn4mIJ3yAB6U
0y4dSKVd6TMqldVy5h6ep6hTQoUFdwIDAQABo4IBajCCAWYwEgYDVR0TAQH/BAgw
BgEB/wIBADAgBgNVHQ4BAf8EFgQUfDB2rMzWh9HsyR/icAgs41/eDAcwDgYDVR0P
AQH/BAQDAgEGMB8GA1UdIwQYMBaAFHwwdqzM1ofR7Mkf4nAILONf3gwHMGMGA1Ud
IARcMFowWAYKYIZIAWUDAgEwATBKMEgGCCsGAQUFBwICMDwMOlRlbnRvIGNlcnRp
Zmlrw6F0IGJ5bCB2eWTDoW4gcG91emUgcHJvIHRlc3RvdmFjw60gw7rEjWVseS4w
gZcGA1UdHwSBjzCBjDCBiaCBhqCBg4YpaHR0cDovL2NybC5jYTEtcGcuZWV0LmN6
L2VldGNhMXBnL2FsbC5jcmyGKmh0dHA6Ly9jcmwyLmNhMS1wZy5lZXQuY3ovZWV0
Y2ExcGcvYWxsLmNybIYqaHR0cDovL2NybDMuY2ExLXBnLmVldC5jei9lZXRjYTFw
Zy9hbGwuY3JsMA0GCSqGSIb3DQEBCwUAA4IBAQBX9p/YNy5/WzcIuDDCfHTzn/ig
qJSnl9Um5lCZrzI9u7k6ZbHec2sOiAF/5/O58kHfCUL1b49U57uJzuPEzDqY8nXE
TynDWLYKyzdPfQZWkOxROhOv0GfQaJN5v98eNazQKurlWHaCHDBIt65i1K1wQ7HP
lwSuScJSo7DV28N+iusQmJg6D5mCTHNuOKVr2RktlzP72XjMi8lYZvK9j7wwXPk2
dZLmkDd+9QirvDeO+gJX3JZvYLAVV4jqrF7vdxi5J57oIFP6cToNk3JJColhclgn
QZkGDCaEbqi3U+Rcx3B54gPdhZSDuiW4efE6LI8ogUmCvhsrn//L2ULJfPLf
-----END CERTIFICATE-----`)

	// CAEET1Production is a root certificate of the Certificate Authority EET for production purposes.
	// The certificate is valid from August 31, 2016 to August 31, 2022.
	// It can be downloaded at https://www.etrzby.cz/cs/certifikacni-autorita-EET.
	CAEET1Production = []byte(`
-----BEGIN CERTIFICATE-----
MIIEZjCCA06gAwIBAgIFAL46sBMwDQYJKoZIhvcNAQELBQAwbDESMBAGCgmSJomT
8ixkARkWAkNaMUMwQQYDVQQKDDrEjGVza8OhIFJlcHVibGlrYSDigJMgR2VuZXLD
oWxuw60gZmluYW7EjW7DrSDFmWVkaXRlbHN0dsOtMREwDwYDVQQDEwhFRVQgQ0Eg
MTAeFw0xNjA4MzEyMTU2NTBaFw0yMjA4MzEyMTU2NTBaMGwxEjAQBgoJkiaJk/Is
ZAEZFgJDWjFDMEEGA1UECgw6xIxlc2vDoSBSZXB1Ymxpa2Eg4oCTIEdlbmVyw6Fs
bsOtIGZpbmFuxI1uw60gxZllZGl0ZWxzdHbDrTERMA8GA1UEAxMIRUVUIENBIDEw
ggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDo+EO/DA+QL6EjYyiAE1Qn
CbrVYegyFagUybUkCN+e2oJk/wA8okBjROKKAushJTXfd5CxkuXkYbtKXoycjSXj
RoC/4ExSLvck6qMCnu5gg6RwB+oIN56XMAuv4aYd1IbwYjuahO//c8LlEK+Um1fe
ESe4c5+lVzR+9IXBBw8rMJMoYZaZMSRTL18sHujaWCoI5zCa6ZT8xOtUM1Bcd/vJ
VRek3H+AL70TnvpYWxW6OEisoUaDLMph90LAjPzQRMzQEvTM/1KOPqZOKDkHGbCt
rAnCgMnE8Jr2kXH+vefrxLx6qff0ZBvyXDVMVNUqrVmet751dFZuPu51uU5RUcoL
AgMBAAGjggENMIIBCTASBgNVHRMBAf8ECDAGAQH/AgEAMCAGA1UdDgEB/wQWBBTj
Ui97DA7NpK9Aijee4f5xWN6UMjAOBgNVHQ8BAf8EBAMCAQYwHwYDVR0jBBgwFoAU
41IvewwOzaSvQIo3nuH+cVjelDIwGgYDVR0gBBMwETAPBg0qgUuM9ugpAYFIAQEB
MIGDBgNVHR8EfDB6MHigdqB0hiRodHRwOi8vY3JsLmNhMS5lZXQuY3ovZWV0Y2Ex
L2FsbC5jcmyGJWh0dHA6Ly9jcmwyLmNhMS5lZXQuY3ovZWV0Y2ExL2FsbC5jcmyG
JWh0dHA6Ly9jcmwzLmNhMS5lZXQuY3ovZWV0Y2ExL2FsbC5jcmwwDQYJKoZIhvcN
AQELBQADggEBAFBCuyRWUk2dAvZIW6WIAglEsQza/7MOY4oGOhYts7GvMrYiI9LC
qgxLl6nCUXjR/BSIEdl2oNtJhmm2HEare9gFF7d4cG58po7GaaPmgSES0PRVECS2
TpKYhrbEzlAgKU3nssK281YWLbP27T2byXJ5YK6ba1uRXt/SzDPn/Pf3im+BORMi
g1KszhTDknvVRf9Q+R8a0LoSg73DiTsob3dgke70eK6+EXEOIdYKWsJKwWK9wqhH
Hgtr96Y+W0iqmTcJFaPUh6Ije6T3EjdqJhQYiiIINWbq0M+OBUY07VrBoTERClnl
UMIkISGhjf2xjWpz5Jau3aIurVCCQwOz/jw=
-----END CERTIFICATE-----`)

	// CAEET1Production2025 is a root certificate of the Certificate Authority EET for production purposes.
	// The certificate is valid from August 31, 2016 to September 29, 2025.
	// It can be downloaded at https://www.etrzby.cz/cs/certifikacni-autorita-EET.
	CAEET1Production2025 = []byte(`
-----BEGIN CERTIFICATE-----
MIIEZjCCA06gAwIBAgIFAJA6c9AwDQYJKoZIhvcNAQELBQAwbDESMBAGCgmSJomT
8ixkARkWAkNaMUMwQQYDVQQKDDrEjGVza8OhIFJlcHVibGlrYSDigJMgR2VuZXLD
oWxuw60gZmluYW7EjW7DrSDFmWVkaXRlbHN0dsOtMREwDwYDVQQDEwhFRVQgQ0Eg
MTAeFw0xNjA4MzEyMTU2NTBaFw0yNTA5MjkxMDAwMDBaMGwxEjAQBgoJkiaJk/Is
ZAEZFgJDWjFDMEEGA1UECgw6xIxlc2vDoSBSZXB1Ymxpa2Eg4oCTIEdlbmVyw6Fs
bsOtIGZpbmFuxI1uw60gxZllZGl0ZWxzdHbDrTERMA8GA1UEAxMIRUVUIENBIDEw
ggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDo+EO/DA+QL6EjYyiAE1Qn
CbrVYegyFagUybUkCN+e2oJk/wA8okBjROKKAushJTXfd5CxkuXkYbtKXoycjSXj
RoC/4ExSLvck6qMCnu5gg6RwB+oIN56XMAuv4aYd1IbwYjuahO//c8LlEK+Um1fe
ESe4c5+lVzR+9IXBBw8rMJMoYZaZMSRTL18sHujaWCoI5zCa6ZT8xOtUM1Bcd/vJ
VRek3H+AL70TnvpYWxW6OEisoUaDLMph90LAjPzQRMzQEvTM/1KOPqZOKDkHGbCt
rAnCgMnE8Jr2kXH+vefrxLx6qff0ZBvyXDVMVNUqrVmet751dFZuPu51uU5RUcoL
AgMBAAGjggENMIIBCTASBgNVHRMBAf8ECDAGAQH/AgEAMCAGA1UdDgEB/wQWBBTj
Ui97DA7NpK9Aijee4f5xWN6UMjAOBgNVHQ8BAf8EBAMCAQYwHwYDVR0jBBgwFoAU
41IvewwOzaSvQIo3nuH+cVjelDIwGgYDVR0gBBMwETAPBg0qgUuM9ugpAYFIAQEB
MIGDBgNVHR8EfDB6MHigdqB0hiRodHRwOi8vY3JsLmNhMS5lZXQuY3ovZWV0Y2Ex
L2FsbC5jcmyGJWh0dHA6Ly9jcmwyLmNhMS5lZXQuY3ovZWV0Y2ExL2FsbC5jcmyG
JWh0dHA6Ly9jcmwzLmNhMS5lZXQuY3ovZWV0Y2ExL2FsbC5jcmwwDQYJKoZIhvcN
AQELBQADggEBANQ1FCrSN41Z0p1QZpiXlz78mEZK2YArXGVSIpme3KVfpZ1D+wwR
5jxc2BT+nA20LONuTDgnMJ0YXSvpcGoL3j48xBhjBIX+0fdxvvgGmFuT6MHYP2DB
4eFuhutkXPPTJnwhY1yH8dLCBR5mal8F1aa43RtCOLDyXB9V878WgknJsLt+67OH
63vf7Q12LfR7aQlVuw+Arp1Owx98OR+V76odpWGScb96qFZHpLBNGH3eMWDfQNJy
dJPXwwGwSf4CLfhTYpSOWLp1+eppP5BTOA/DfBPPaTVwyfjx7lul2fyRAN706TFS
7S2mM1UWd4um1ARWkEOrlMGNmzpTHWC/esc=
-----END CERTIFICATE-----`)
)

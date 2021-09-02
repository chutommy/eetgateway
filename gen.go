package main

//go:generate go run bin/soap-defs/main.go -dir data/eet-specs/soap-definition
//go:generate wsdlgen -o pkg/eet/definition-gen.go -pkg eet data/eet-specs/soap-definition/EETXMLSchema.xsd data/eet-specs/soap-definition/EETServiceSOAP.wsdl

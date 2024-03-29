FROM golang:latest AS build

WORKDIR /build
COPY data data

# https://github.com/droyo/go-xml
RUN go install aqwari.net/xml/cmd/wsdlgen@latest
RUN wsdlgen -o ./eet-gen.go -pkg eet data/eet-specs/soap-definition/EETXMLSchema.xsd data/eet-specs/soap-definition/EETServiceSOAP.wsdl

# insert XML chardata attributes for: OdpovedChybaType, OdpovedVarovaniType
RUN sed -i '52 i Zprava string `xml:",chardata"`' ./eet-gen.go
RUN sed -i '77 i Zprava string `xml:",chardata"`' ./eet-gen.go

# generate custom XML tag for: OdpovedType, TrzbaType, TrzbaKontrolniKodyType
# https://github.com/fatih/gomodifytags
RUN go install github.com/fatih/gomodifytags@latest
RUN gomodifytags -file ./eet-gen.go -struct OdpovedType -remove-tags xml -w --quiet
RUN gomodifytags -file ./eet-gen.go -struct OdpovedType -add-tags xml -add-options xml=omitempty -transform pascalcase -w --quiet
RUN gomodifytags -file ./eet-gen.go -struct TrzbaType -remove-tags xml -w --quiet
RUN gomodifytags -file ./eet-gen.go -struct TrzbaType -add-tags xml -transform pascalcase -w --quiet
RUN gomodifytags -file ./eet-gen.go -struct TrzbaKontrolniKodyType -remove-tags xml -w --quiet
RUN gomodifytags -file ./eet-gen.go -struct TrzbaKontrolniKodyType -add-tags xml -transform snakecase -w --quiet

RUN chmod 666 eet-gen.go

CMD mv /build/eet-gen.go /gen/
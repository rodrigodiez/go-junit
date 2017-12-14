package types

import (
  "encoding/xml"
)

type Failure struct {
  XMLName xml.Name  `xml:"failure"`
  Message string  `xml:"message,attr"`
  Type string `xml:"type,attr"`
  Text  string  `xml:",chardata"`
}

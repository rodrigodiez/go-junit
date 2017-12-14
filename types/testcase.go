package types

import (
  "encoding/xml"
)


type Testcase struct {
  XMLName xml.Name  `xml:"testcase"`
  Id string  `xml:"id,attr"`
  Name string `xml:"name,attr"`
  Time  float32 `xml:"time,attr"`
  Failures []*Failure `xml:,`
}

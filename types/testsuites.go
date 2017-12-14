package types

import (
	"encoding/xml"
)

type Testsuites struct {
	XMLName    xml.Name     `xml:"testsuites"`
	Id         string       `xml:"id,attr"`
	Name       string       `xml:"name,attr"`
	Tests      int          `xml:"tests,attr"`
	Failures   int          `xml:"failures,attr"`
	Time       float32      `xml:"time,attr"`
	Testsuites []*Testsuite `xml:,`
}

package types

import (
	"encoding/xml"
)

type Testsuite struct {
	XMLName   xml.Name    `xml:"testsuite"`
	Id        string      `xml:"id,attr"`
	Name      string      `xml:"name,attr"`
	Tests     int         `xml:"tests,attr"`
	Failures  int         `xml:"failures,attr"`
	Time      float32     `xml:"time,attr"`
	Testcases []*Testcase `xml:,`
}

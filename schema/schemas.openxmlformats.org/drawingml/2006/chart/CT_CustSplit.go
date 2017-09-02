// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_CustSplit struct {
	SecondPiePt []*CT_UnsignedInt
}

func NewCT_CustSplit() *CT_CustSplit {
	ret := &CT_CustSplit{}
	return ret
}

func (m *CT_CustSplit) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.SecondPiePt != nil {
		sesecondPiePt := xml.StartElement{Name: xml.Name{Local: "secondPiePt"}}
		e.EncodeElement(m.SecondPiePt, sesecondPiePt)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CustSplit) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CustSplit:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "secondPiePt":
				tmp := NewCT_UnsignedInt()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SecondPiePt = append(m.SecondPiePt, tmp)
			default:
				log.Printf("skipping unsupported element on CT_CustSplit %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustSplit
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CustSplit and its children
func (m *CT_CustSplit) Validate() error {
	return m.ValidateWithPath("CT_CustSplit")
}

// ValidateWithPath validates the CT_CustSplit and its children, prefixing error messages with path
func (m *CT_CustSplit) ValidateWithPath(path string) error {
	for i, v := range m.SecondPiePt {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SecondPiePt[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}

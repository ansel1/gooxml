// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"log"
)

type CT_Rad struct {
	RadPr *CT_RadPr
	Deg   *CT_OMathArg
	E     *CT_OMathArg
}

func NewCT_Rad() *CT_Rad {
	ret := &CT_Rad{}
	ret.Deg = NewCT_OMathArg()
	ret.E = NewCT_OMathArg()
	return ret
}

func (m *CT_Rad) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.RadPr != nil {
		seradPr := xml.StartElement{Name: xml.Name{Local: "m:radPr"}}
		e.EncodeElement(m.RadPr, seradPr)
	}
	sedeg := xml.StartElement{Name: xml.Name{Local: "m:deg"}}
	e.EncodeElement(m.Deg, sedeg)
	see := xml.StartElement{Name: xml.Name{Local: "m:e"}}
	e.EncodeElement(m.E, see)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Rad) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Deg = NewCT_OMathArg()
	m.E = NewCT_OMathArg()
lCT_Rad:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "radPr":
				m.RadPr = NewCT_RadPr()
				if err := d.DecodeElement(m.RadPr, &el); err != nil {
					return err
				}
			case "deg":
				if err := d.DecodeElement(m.Deg, &el); err != nil {
					return err
				}
			case "e":
				if err := d.DecodeElement(m.E, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Rad %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Rad
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Rad and its children
func (m *CT_Rad) Validate() error {
	return m.ValidateWithPath("CT_Rad")
}

// ValidateWithPath validates the CT_Rad and its children, prefixing error messages with path
func (m *CT_Rad) ValidateWithPath(path string) error {
	if m.RadPr != nil {
		if err := m.RadPr.ValidateWithPath(path + "/RadPr"); err != nil {
			return err
		}
	}
	if err := m.Deg.ValidateWithPath(path + "/Deg"); err != nil {
		return err
	}
	if err := m.E.ValidateWithPath(path + "/E"); err != nil {
		return err
	}
	return nil
}
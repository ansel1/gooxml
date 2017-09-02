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

type CT_Box struct {
	BoxPr *CT_BoxPr
	E     *CT_OMathArg
}

func NewCT_Box() *CT_Box {
	ret := &CT_Box{}
	ret.E = NewCT_OMathArg()
	return ret
}

func (m *CT_Box) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.BoxPr != nil {
		seboxPr := xml.StartElement{Name: xml.Name{Local: "m:boxPr"}}
		e.EncodeElement(m.BoxPr, seboxPr)
	}
	see := xml.StartElement{Name: xml.Name{Local: "m:e"}}
	e.EncodeElement(m.E, see)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Box) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.E = NewCT_OMathArg()
lCT_Box:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "boxPr":
				m.BoxPr = NewCT_BoxPr()
				if err := d.DecodeElement(m.BoxPr, &el); err != nil {
					return err
				}
			case "e":
				if err := d.DecodeElement(m.E, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Box %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Box
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Box and its children
func (m *CT_Box) Validate() error {
	return m.ValidateWithPath("CT_Box")
}

// ValidateWithPath validates the CT_Box and its children, prefixing error messages with path
func (m *CT_Box) ValidateWithPath(path string) error {
	if m.BoxPr != nil {
		if err := m.BoxPr.ValidateWithPath(path + "/BoxPr"); err != nil {
			return err
		}
	}
	if err := m.E.ValidateWithPath(path + "/E"); err != nil {
		return err
	}
	return nil
}

// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"log"
)

type EG_RPrMath struct {
	// Inserted Math Control Character
	Ins *CT_MathCtrlIns
	// Deleted Math Control Character
	Del *CT_MathCtrlDel
	// Run Properties
	RPr *CT_RPr
}

func NewEG_RPrMath() *EG_RPrMath {
	ret := &EG_RPrMath{}
	return ret
}

func (m *EG_RPrMath) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:EG_RPrMath"
	if m.Ins != nil {
		seins := xml.StartElement{Name: xml.Name{Local: "w:ins"}}
		e.EncodeElement(m.Ins, seins)
	}
	if m.Del != nil {
		sedel := xml.StartElement{Name: xml.Name{Local: "w:del"}}
		e.EncodeElement(m.Del, sedel)
	}
	if m.RPr != nil {
		serPr := xml.StartElement{Name: xml.Name{Local: "w:rPr"}}
		e.EncodeElement(m.RPr, serPr)
	}
	return nil
}

func (m *EG_RPrMath) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_RPrMath:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "ins":
				m.Ins = NewCT_MathCtrlIns()
				if err := d.DecodeElement(m.Ins, &el); err != nil {
					return err
				}
			case "del":
				m.Del = NewCT_MathCtrlDel()
				if err := d.DecodeElement(m.Del, &el); err != nil {
					return err
				}
			case "rPr":
				m.RPr = NewCT_RPr()
				if err := d.DecodeElement(m.RPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on EG_RPrMath %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_RPrMath
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_RPrMath and its children
func (m *EG_RPrMath) Validate() error {
	return m.ValidateWithPath("EG_RPrMath")
}

// ValidateWithPath validates the EG_RPrMath and its children, prefixing error messages with path
func (m *EG_RPrMath) ValidateWithPath(path string) error {
	if m.Ins != nil {
		if err := m.Ins.ValidateWithPath(path + "/Ins"); err != nil {
			return err
		}
	}
	if m.Del != nil {
		if err := m.Del.ValidateWithPath(path + "/Del"); err != nil {
			return err
		}
	}
	if m.RPr != nil {
		if err := m.RPr.ValidateWithPath(path + "/RPr"); err != nil {
			return err
		}
	}
	return nil
}
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

type CT_MathPrChoice struct {
	WrapIndent *CT_TwipsMeasure
	WrapRight  *CT_OnOff
}

func NewCT_MathPrChoice() *CT_MathPrChoice {
	ret := &CT_MathPrChoice{}
	return ret
}

func (m *CT_MathPrChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.WrapIndent != nil {
		sewrapIndent := xml.StartElement{Name: xml.Name{Local: "m:wrapIndent"}}
		e.EncodeElement(m.WrapIndent, sewrapIndent)
	}
	if m.WrapRight != nil {
		sewrapRight := xml.StartElement{Name: xml.Name{Local: "m:wrapRight"}}
		e.EncodeElement(m.WrapRight, sewrapRight)
	}
	return nil
}

func (m *CT_MathPrChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_MathPrChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "wrapIndent":
				m.WrapIndent = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.WrapIndent, &el); err != nil {
					return err
				}
			case "wrapRight":
				m.WrapRight = NewCT_OnOff()
				if err := d.DecodeElement(m.WrapRight, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_MathPrChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MathPrChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_MathPrChoice and its children
func (m *CT_MathPrChoice) Validate() error {
	return m.ValidateWithPath("CT_MathPrChoice")
}

// ValidateWithPath validates the CT_MathPrChoice and its children, prefixing error messages with path
func (m *CT_MathPrChoice) ValidateWithPath(path string) error {
	if m.WrapIndent != nil {
		if err := m.WrapIndent.ValidateWithPath(path + "/WrapIndent"); err != nil {
			return err
		}
	}
	if m.WrapRight != nil {
		if err := m.WrapRight.ValidateWithPath(path + "/WrapRight"); err != nil {
			return err
		}
	}
	return nil
}

// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_TLAnimVariant struct {
	// Boolean Variant
	BoolVal *CT_TLAnimVariantBooleanVal
	// Integer
	IntVal *CT_TLAnimVariantIntegerVal
	// Float Value
	FltVal *CT_TLAnimVariantFloatVal
	// String Value
	StrVal *CT_TLAnimVariantStringVal
	// Color Value
	ClrVal *drawingml.CT_Color
}

func NewCT_TLAnimVariant() *CT_TLAnimVariant {
	ret := &CT_TLAnimVariant{}
	return ret
}

func (m *CT_TLAnimVariant) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.BoolVal != nil {
		seboolVal := xml.StartElement{Name: xml.Name{Local: "p:boolVal"}}
		e.EncodeElement(m.BoolVal, seboolVal)
	}
	if m.IntVal != nil {
		seintVal := xml.StartElement{Name: xml.Name{Local: "p:intVal"}}
		e.EncodeElement(m.IntVal, seintVal)
	}
	if m.FltVal != nil {
		sefltVal := xml.StartElement{Name: xml.Name{Local: "p:fltVal"}}
		e.EncodeElement(m.FltVal, sefltVal)
	}
	if m.StrVal != nil {
		sestrVal := xml.StartElement{Name: xml.Name{Local: "p:strVal"}}
		e.EncodeElement(m.StrVal, sestrVal)
	}
	if m.ClrVal != nil {
		seclrVal := xml.StartElement{Name: xml.Name{Local: "p:clrVal"}}
		e.EncodeElement(m.ClrVal, seclrVal)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLAnimVariant) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TLAnimVariant:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "boolVal":
				m.BoolVal = NewCT_TLAnimVariantBooleanVal()
				if err := d.DecodeElement(m.BoolVal, &el); err != nil {
					return err
				}
			case "intVal":
				m.IntVal = NewCT_TLAnimVariantIntegerVal()
				if err := d.DecodeElement(m.IntVal, &el); err != nil {
					return err
				}
			case "fltVal":
				m.FltVal = NewCT_TLAnimVariantFloatVal()
				if err := d.DecodeElement(m.FltVal, &el); err != nil {
					return err
				}
			case "strVal":
				m.StrVal = NewCT_TLAnimVariantStringVal()
				if err := d.DecodeElement(m.StrVal, &el); err != nil {
					return err
				}
			case "clrVal":
				m.ClrVal = drawingml.NewCT_Color()
				if err := d.DecodeElement(m.ClrVal, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_TLAnimVariant %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLAnimVariant
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLAnimVariant and its children
func (m *CT_TLAnimVariant) Validate() error {
	return m.ValidateWithPath("CT_TLAnimVariant")
}

// ValidateWithPath validates the CT_TLAnimVariant and its children, prefixing error messages with path
func (m *CT_TLAnimVariant) ValidateWithPath(path string) error {
	if m.BoolVal != nil {
		if err := m.BoolVal.ValidateWithPath(path + "/BoolVal"); err != nil {
			return err
		}
	}
	if m.IntVal != nil {
		if err := m.IntVal.ValidateWithPath(path + "/IntVal"); err != nil {
			return err
		}
	}
	if m.FltVal != nil {
		if err := m.FltVal.ValidateWithPath(path + "/FltVal"); err != nil {
			return err
		}
	}
	if m.StrVal != nil {
		if err := m.StrVal.ValidateWithPath(path + "/StrVal"); err != nil {
			return err
		}
	}
	if m.ClrVal != nil {
		if err := m.ClrVal.ValidateWithPath(path + "/ClrVal"); err != nil {
			return err
		}
	}
	return nil
}
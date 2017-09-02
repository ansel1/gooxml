// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_ColorChangeEffect struct {
	UseAAttr *bool
	ClrFrom  *CT_Color
	ClrTo    *CT_Color
}

func NewCT_ColorChangeEffect() *CT_ColorChangeEffect {
	ret := &CT_ColorChangeEffect{}
	ret.ClrFrom = NewCT_Color()
	ret.ClrTo = NewCT_Color()
	return ret
}

func (m *CT_ColorChangeEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.UseAAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "useA"},
			Value: fmt.Sprintf("%v", *m.UseAAttr)})
	}
	e.EncodeToken(start)
	seclrFrom := xml.StartElement{Name: xml.Name{Local: "a:clrFrom"}}
	e.EncodeElement(m.ClrFrom, seclrFrom)
	seclrTo := xml.StartElement{Name: xml.Name{Local: "a:clrTo"}}
	e.EncodeElement(m.ClrTo, seclrTo)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ColorChangeEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ClrFrom = NewCT_Color()
	m.ClrTo = NewCT_Color()
	for _, attr := range start.Attr {
		if attr.Name.Local == "useA" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UseAAttr = &parsed
		}
	}
lCT_ColorChangeEffect:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "clrFrom":
				if err := d.DecodeElement(m.ClrFrom, &el); err != nil {
					return err
				}
			case "clrTo":
				if err := d.DecodeElement(m.ClrTo, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_ColorChangeEffect %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ColorChangeEffect
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ColorChangeEffect and its children
func (m *CT_ColorChangeEffect) Validate() error {
	return m.ValidateWithPath("CT_ColorChangeEffect")
}

// ValidateWithPath validates the CT_ColorChangeEffect and its children, prefixing error messages with path
func (m *CT_ColorChangeEffect) ValidateWithPath(path string) error {
	if err := m.ClrFrom.ValidateWithPath(path + "/ClrFrom"); err != nil {
		return err
	}
	if err := m.ClrTo.ValidateWithPath(path + "/ClrTo"); err != nil {
		return err
	}
	return nil
}
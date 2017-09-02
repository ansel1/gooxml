// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type CT_ColorSchemeAndMapping struct {
	ClrScheme *CT_ColorScheme
	ClrMap    *CT_ColorMapping
}

func NewCT_ColorSchemeAndMapping() *CT_ColorSchemeAndMapping {
	ret := &CT_ColorSchemeAndMapping{}
	ret.ClrScheme = NewCT_ColorScheme()
	return ret
}

func (m *CT_ColorSchemeAndMapping) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	seclrScheme := xml.StartElement{Name: xml.Name{Local: "a:clrScheme"}}
	e.EncodeElement(m.ClrScheme, seclrScheme)
	if m.ClrMap != nil {
		seclrMap := xml.StartElement{Name: xml.Name{Local: "a:clrMap"}}
		e.EncodeElement(m.ClrMap, seclrMap)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ColorSchemeAndMapping) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ClrScheme = NewCT_ColorScheme()
lCT_ColorSchemeAndMapping:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "clrScheme":
				if err := d.DecodeElement(m.ClrScheme, &el); err != nil {
					return err
				}
			case "clrMap":
				m.ClrMap = NewCT_ColorMapping()
				if err := d.DecodeElement(m.ClrMap, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_ColorSchemeAndMapping %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ColorSchemeAndMapping
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ColorSchemeAndMapping and its children
func (m *CT_ColorSchemeAndMapping) Validate() error {
	return m.ValidateWithPath("CT_ColorSchemeAndMapping")
}

// ValidateWithPath validates the CT_ColorSchemeAndMapping and its children, prefixing error messages with path
func (m *CT_ColorSchemeAndMapping) ValidateWithPath(path string) error {
	if err := m.ClrScheme.ValidateWithPath(path + "/ClrScheme"); err != nil {
		return err
	}
	if m.ClrMap != nil {
		if err := m.ClrMap.ValidateWithPath(path + "/ClrMap"); err != nil {
			return err
		}
	}
	return nil
}

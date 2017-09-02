// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_Object struct {
	// Original Image Width
	DxaOrigAttr *sharedTypes.ST_TwipsMeasure
	// Original Image Height
	DyaOrigAttr *sharedTypes.ST_TwipsMeasure
	Drawing     *CT_Drawing
	Choice      *CT_ObjectChoice
}

func NewCT_Object() *CT_Object {
	ret := &CT_Object{}
	return ret
}

func (m *CT_Object) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.DxaOrigAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:dxaOrig"},
			Value: fmt.Sprintf("%v", *m.DxaOrigAttr)})
	}
	if m.DyaOrigAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:dyaOrig"},
			Value: fmt.Sprintf("%v", *m.DyaOrigAttr)})
	}
	e.EncodeToken(start)
	if m.Drawing != nil {
		sedrawing := xml.StartElement{Name: xml.Name{Local: "w:drawing"}}
		e.EncodeElement(m.Drawing, sedrawing)
	}
	if m.Choice != nil {
		m.Choice.MarshalXML(e, start)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Object) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "dxaOrig" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.DxaOrigAttr = &parsed
		}
		if attr.Name.Local == "dyaOrig" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.DyaOrigAttr = &parsed
		}
	}
lCT_Object:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "drawing":
				m.Drawing = NewCT_Drawing()
				if err := d.DecodeElement(m.Drawing, &el); err != nil {
					return err
				}
			case "control":
				m.Choice = NewCT_ObjectChoice()
				if err := d.DecodeElement(&m.Choice.Control, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "objectLink":
				m.Choice = NewCT_ObjectChoice()
				if err := d.DecodeElement(&m.Choice.ObjectLink, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "objectEmbed":
				m.Choice = NewCT_ObjectChoice()
				if err := d.DecodeElement(&m.Choice.ObjectEmbed, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "movie":
				m.Choice = NewCT_ObjectChoice()
				if err := d.DecodeElement(&m.Choice.Movie, &el); err != nil {
					return err
				}
				_ = m.Choice
			default:
				log.Printf("skipping unsupported element on CT_Object %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Object
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Object and its children
func (m *CT_Object) Validate() error {
	return m.ValidateWithPath("CT_Object")
}

// ValidateWithPath validates the CT_Object and its children, prefixing error messages with path
func (m *CT_Object) ValidateWithPath(path string) error {
	if m.DxaOrigAttr != nil {
		if err := m.DxaOrigAttr.ValidateWithPath(path + "/DxaOrigAttr"); err != nil {
			return err
		}
	}
	if m.DyaOrigAttr != nil {
		if err := m.DyaOrigAttr.ValidateWithPath(path + "/DyaOrigAttr"); err != nil {
			return err
		}
	}
	if m.Drawing != nil {
		if err := m.Drawing.ValidateWithPath(path + "/Drawing"); err != nil {
			return err
		}
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	return nil
}
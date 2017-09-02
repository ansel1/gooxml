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
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_Columns struct {
	// Equal Column Widths
	EqualWidthAttr *sharedTypes.ST_OnOff
	// Spacing Between Equal Width Columns
	SpaceAttr *sharedTypes.ST_TwipsMeasure
	// Number of Equal Width Columns
	NumAttr *int64
	// Draw Line Between Columns
	SepAttr *sharedTypes.ST_OnOff
	// Single Column Definition
	Col []*CT_Column
}

func NewCT_Columns() *CT_Columns {
	ret := &CT_Columns{}
	return ret
}

func (m *CT_Columns) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.EqualWidthAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:equalWidth"},
			Value: fmt.Sprintf("%v", *m.EqualWidthAttr)})
	}
	if m.SpaceAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:space"},
			Value: fmt.Sprintf("%v", *m.SpaceAttr)})
	}
	if m.NumAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:num"},
			Value: fmt.Sprintf("%v", *m.NumAttr)})
	}
	if m.SepAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:sep"},
			Value: fmt.Sprintf("%v", *m.SepAttr)})
	}
	e.EncodeToken(start)
	secol := xml.StartElement{Name: xml.Name{Local: "w:col"}}
	e.EncodeElement(m.Col, secol)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Columns) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "equalWidth" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.EqualWidthAttr = &parsed
		}
		if attr.Name.Local == "space" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.SpaceAttr = &parsed
		}
		if attr.Name.Local == "num" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.NumAttr = &parsed
		}
		if attr.Name.Local == "sep" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.SepAttr = &parsed
		}
	}
lCT_Columns:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "col":
				tmp := NewCT_Column()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Col = append(m.Col, tmp)
			default:
				log.Printf("skipping unsupported element on CT_Columns %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Columns
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Columns and its children
func (m *CT_Columns) Validate() error {
	return m.ValidateWithPath("CT_Columns")
}

// ValidateWithPath validates the CT_Columns and its children, prefixing error messages with path
func (m *CT_Columns) ValidateWithPath(path string) error {
	if m.EqualWidthAttr != nil {
		if err := m.EqualWidthAttr.ValidateWithPath(path + "/EqualWidthAttr"); err != nil {
			return err
		}
	}
	if m.SpaceAttr != nil {
		if err := m.SpaceAttr.ValidateWithPath(path + "/SpaceAttr"); err != nil {
			return err
		}
	}
	if m.SepAttr != nil {
		if err := m.SepAttr.ValidateWithPath(path + "/SepAttr"); err != nil {
			return err
		}
	}
	for i, v := range m.Col {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Col[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
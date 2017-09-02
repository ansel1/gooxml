// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_BandFmt struct {
	Idx  *CT_UnsignedInt
	SpPr *drawingml.CT_ShapeProperties
}

func NewCT_BandFmt() *CT_BandFmt {
	ret := &CT_BandFmt{}
	ret.Idx = NewCT_UnsignedInt()
	return ret
}

func (m *CT_BandFmt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	seidx := xml.StartElement{Name: xml.Name{Local: "idx"}}
	e.EncodeElement(m.Idx, seidx)
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BandFmt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Idx = NewCT_UnsignedInt()
lCT_BandFmt:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "idx":
				if err := d.DecodeElement(m.Idx, &el); err != nil {
					return err
				}
			case "spPr":
				m.SpPr = drawingml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_BandFmt %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BandFmt
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BandFmt and its children
func (m *CT_BandFmt) Validate() error {
	return m.ValidateWithPath("CT_BandFmt")
}

// ValidateWithPath validates the CT_BandFmt and its children, prefixing error messages with path
func (m *CT_BandFmt) ValidateWithPath(path string) error {
	if err := m.Idx.ValidateWithPath(path + "/Idx"); err != nil {
		return err
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
			return err
		}
	}
	return nil
}

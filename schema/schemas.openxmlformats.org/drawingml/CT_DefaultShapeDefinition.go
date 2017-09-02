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

type CT_DefaultShapeDefinition struct {
	SpPr     *CT_ShapeProperties
	BodyPr   *CT_TextBodyProperties
	LstStyle *CT_TextListStyle
	Style    *CT_ShapeStyle
	ExtLst   *CT_OfficeArtExtensionList
}

func NewCT_DefaultShapeDefinition() *CT_DefaultShapeDefinition {
	ret := &CT_DefaultShapeDefinition{}
	ret.SpPr = NewCT_ShapeProperties()
	ret.BodyPr = NewCT_TextBodyProperties()
	ret.LstStyle = NewCT_TextListStyle()
	return ret
}

func (m *CT_DefaultShapeDefinition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sespPr := xml.StartElement{Name: xml.Name{Local: "a:spPr"}}
	e.EncodeElement(m.SpPr, sespPr)
	sebodyPr := xml.StartElement{Name: xml.Name{Local: "a:bodyPr"}}
	e.EncodeElement(m.BodyPr, sebodyPr)
	selstStyle := xml.StartElement{Name: xml.Name{Local: "a:lstStyle"}}
	e.EncodeElement(m.LstStyle, selstStyle)
	if m.Style != nil {
		sestyle := xml.StartElement{Name: xml.Name{Local: "a:style"}}
		e.EncodeElement(m.Style, sestyle)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DefaultShapeDefinition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.SpPr = NewCT_ShapeProperties()
	m.BodyPr = NewCT_TextBodyProperties()
	m.LstStyle = NewCT_TextListStyle()
lCT_DefaultShapeDefinition:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "spPr":
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case "bodyPr":
				if err := d.DecodeElement(m.BodyPr, &el); err != nil {
					return err
				}
			case "lstStyle":
				if err := d.DecodeElement(m.LstStyle, &el); err != nil {
					return err
				}
			case "style":
				m.Style = NewCT_ShapeStyle()
				if err := d.DecodeElement(m.Style, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_DefaultShapeDefinition %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DefaultShapeDefinition
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DefaultShapeDefinition and its children
func (m *CT_DefaultShapeDefinition) Validate() error {
	return m.ValidateWithPath("CT_DefaultShapeDefinition")
}

// ValidateWithPath validates the CT_DefaultShapeDefinition and its children, prefixing error messages with path
func (m *CT_DefaultShapeDefinition) ValidateWithPath(path string) error {
	if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
		return err
	}
	if err := m.BodyPr.ValidateWithPath(path + "/BodyPr"); err != nil {
		return err
	}
	if err := m.LstStyle.ValidateWithPath(path + "/LstStyle"); err != nil {
		return err
	}
	if m.Style != nil {
		if err := m.Style.ValidateWithPath(path + "/Style"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
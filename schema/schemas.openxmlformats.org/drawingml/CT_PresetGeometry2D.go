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
)

type CT_PresetGeometry2D struct {
	PrstAttr ST_ShapeType
	AvLst    *CT_GeomGuideList
}

func NewCT_PresetGeometry2D() *CT_PresetGeometry2D {
	ret := &CT_PresetGeometry2D{}
	ret.PrstAttr = ST_ShapeType(1)
	return ret
}

func (m *CT_PresetGeometry2D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.PrstAttr.MarshalXMLAttr(xml.Name{Local: "prst"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	if m.AvLst != nil {
		seavLst := xml.StartElement{Name: xml.Name{Local: "a:avLst"}}
		e.EncodeElement(m.AvLst, seavLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PresetGeometry2D) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.PrstAttr = ST_ShapeType(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "prst" {
			m.PrstAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_PresetGeometry2D:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "avLst":
				m.AvLst = NewCT_GeomGuideList()
				if err := d.DecodeElement(m.AvLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_PresetGeometry2D %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PresetGeometry2D
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PresetGeometry2D and its children
func (m *CT_PresetGeometry2D) Validate() error {
	return m.ValidateWithPath("CT_PresetGeometry2D")
}

// ValidateWithPath validates the CT_PresetGeometry2D and its children, prefixing error messages with path
func (m *CT_PresetGeometry2D) ValidateWithPath(path string) error {
	if m.PrstAttr == ST_ShapeTypeUnset {
		return fmt.Errorf("%s/PrstAttr is a mandatory field", path)
	}
	if err := m.PrstAttr.ValidateWithPath(path + "/PrstAttr"); err != nil {
		return err
	}
	if m.AvLst != nil {
		if err := m.AvLst.ValidateWithPath(path + "/AvLst"); err != nil {
			return err
		}
	}
	return nil
}

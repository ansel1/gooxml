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

type CT_PresetTextShape struct {
	PrstAttr ST_TextShapeType
	AvLst    *CT_GeomGuideList
}

func NewCT_PresetTextShape() *CT_PresetTextShape {
	ret := &CT_PresetTextShape{}
	ret.PrstAttr = ST_TextShapeType(1)
	return ret
}

func (m *CT_PresetTextShape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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

func (m *CT_PresetTextShape) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.PrstAttr = ST_TextShapeType(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "prst" {
			m.PrstAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_PresetTextShape:
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
				log.Printf("skipping unsupported element on CT_PresetTextShape %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PresetTextShape
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PresetTextShape and its children
func (m *CT_PresetTextShape) Validate() error {
	return m.ValidateWithPath("CT_PresetTextShape")
}

// ValidateWithPath validates the CT_PresetTextShape and its children, prefixing error messages with path
func (m *CT_PresetTextShape) ValidateWithPath(path string) error {
	if m.PrstAttr == ST_TextShapeTypeUnset {
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
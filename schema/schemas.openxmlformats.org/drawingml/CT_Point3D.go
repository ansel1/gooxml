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
)

type CT_Point3D struct {
	XAttr ST_Coordinate
	YAttr ST_Coordinate
	ZAttr ST_Coordinate
}

func NewCT_Point3D() *CT_Point3D {
	ret := &CT_Point3D{}
	return ret
}

func (m *CT_Point3D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "x"},
		Value: fmt.Sprintf("%v", m.XAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "y"},
		Value: fmt.Sprintf("%v", m.YAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "z"},
		Value: fmt.Sprintf("%v", m.ZAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Point3D) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "x" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.XAttr = parsed
		}
		if attr.Name.Local == "y" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.YAttr = parsed
		}
		if attr.Name.Local == "z" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.ZAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Point3D: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Point3D and its children
func (m *CT_Point3D) Validate() error {
	return m.ValidateWithPath("CT_Point3D")
}

// ValidateWithPath validates the CT_Point3D and its children, prefixing error messages with path
func (m *CT_Point3D) ValidateWithPath(path string) error {
	if err := m.XAttr.ValidateWithPath(path + "/XAttr"); err != nil {
		return err
	}
	if err := m.YAttr.ValidateWithPath(path + "/YAttr"); err != nil {
		return err
	}
	if err := m.ZAttr.ValidateWithPath(path + "/ZAttr"); err != nil {
		return err
	}
	return nil
}

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

type CT_PolarAdjustHandle struct {
	GdRefRAttr   *string
	MinRAttr     *ST_AdjCoordinate
	MaxRAttr     *ST_AdjCoordinate
	GdRefAngAttr *string
	MinAngAttr   *ST_AdjAngle
	MaxAngAttr   *ST_AdjAngle
	Pos          *CT_AdjPoint2D
}

func NewCT_PolarAdjustHandle() *CT_PolarAdjustHandle {
	ret := &CT_PolarAdjustHandle{}
	ret.Pos = NewCT_AdjPoint2D()
	return ret
}

func (m *CT_PolarAdjustHandle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.GdRefRAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "gdRefR"},
			Value: fmt.Sprintf("%v", *m.GdRefRAttr)})
	}
	if m.MinRAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minR"},
			Value: fmt.Sprintf("%v", *m.MinRAttr)})
	}
	if m.MaxRAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "maxR"},
			Value: fmt.Sprintf("%v", *m.MaxRAttr)})
	}
	if m.GdRefAngAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "gdRefAng"},
			Value: fmt.Sprintf("%v", *m.GdRefAngAttr)})
	}
	if m.MinAngAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minAng"},
			Value: fmt.Sprintf("%v", *m.MinAngAttr)})
	}
	if m.MaxAngAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "maxAng"},
			Value: fmt.Sprintf("%v", *m.MaxAngAttr)})
	}
	e.EncodeToken(start)
	sepos := xml.StartElement{Name: xml.Name{Local: "a:pos"}}
	e.EncodeElement(m.Pos, sepos)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PolarAdjustHandle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Pos = NewCT_AdjPoint2D()
	for _, attr := range start.Attr {
		if attr.Name.Local == "gdRefR" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GdRefRAttr = &parsed
		}
		if attr.Name.Local == "minR" {
			parsed, err := ParseUnionST_AdjCoordinate(attr.Value)
			if err != nil {
				return err
			}
			m.MinRAttr = &parsed
		}
		if attr.Name.Local == "maxR" {
			parsed, err := ParseUnionST_AdjCoordinate(attr.Value)
			if err != nil {
				return err
			}
			m.MaxRAttr = &parsed
		}
		if attr.Name.Local == "gdRefAng" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GdRefAngAttr = &parsed
		}
		if attr.Name.Local == "minAng" {
			parsed, err := ParseUnionST_AdjAngle(attr.Value)
			if err != nil {
				return err
			}
			m.MinAngAttr = &parsed
		}
		if attr.Name.Local == "maxAng" {
			parsed, err := ParseUnionST_AdjAngle(attr.Value)
			if err != nil {
				return err
			}
			m.MaxAngAttr = &parsed
		}
	}
lCT_PolarAdjustHandle:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "pos":
				if err := d.DecodeElement(m.Pos, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_PolarAdjustHandle %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PolarAdjustHandle
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PolarAdjustHandle and its children
func (m *CT_PolarAdjustHandle) Validate() error {
	return m.ValidateWithPath("CT_PolarAdjustHandle")
}

// ValidateWithPath validates the CT_PolarAdjustHandle and its children, prefixing error messages with path
func (m *CT_PolarAdjustHandle) ValidateWithPath(path string) error {
	if m.MinRAttr != nil {
		if err := m.MinRAttr.ValidateWithPath(path + "/MinRAttr"); err != nil {
			return err
		}
	}
	if m.MaxRAttr != nil {
		if err := m.MaxRAttr.ValidateWithPath(path + "/MaxRAttr"); err != nil {
			return err
		}
	}
	if m.MinAngAttr != nil {
		if err := m.MinAngAttr.ValidateWithPath(path + "/MinAngAttr"); err != nil {
			return err
		}
	}
	if m.MaxAngAttr != nil {
		if err := m.MaxAngAttr.ValidateWithPath(path + "/MaxAngAttr"); err != nil {
			return err
		}
	}
	if err := m.Pos.ValidateWithPath(path + "/Pos"); err != nil {
		return err
	}
	return nil
}
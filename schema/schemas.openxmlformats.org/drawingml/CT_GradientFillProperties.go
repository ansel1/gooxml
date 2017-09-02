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

type CT_GradientFillProperties struct {
	FlipAttr         ST_TileFlipMode
	RotWithShapeAttr *bool
	GsLst            *CT_GradientStopList
	Lin              *CT_LinearShadeProperties
	Path             *CT_PathShadeProperties
	TileRect         *CT_RelativeRect
}

func NewCT_GradientFillProperties() *CT_GradientFillProperties {
	ret := &CT_GradientFillProperties{}
	return ret
}

func (m *CT_GradientFillProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.FlipAttr != ST_TileFlipModeUnset {
		attr, err := m.FlipAttr.MarshalXMLAttr(xml.Name{Local: "flip"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.RotWithShapeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rotWithShape"},
			Value: fmt.Sprintf("%v", *m.RotWithShapeAttr)})
	}
	e.EncodeToken(start)
	if m.GsLst != nil {
		segsLst := xml.StartElement{Name: xml.Name{Local: "a:gsLst"}}
		e.EncodeElement(m.GsLst, segsLst)
	}
	if m.Lin != nil {
		selin := xml.StartElement{Name: xml.Name{Local: "a:lin"}}
		e.EncodeElement(m.Lin, selin)
	}
	if m.Path != nil {
		sepath := xml.StartElement{Name: xml.Name{Local: "a:path"}}
		e.EncodeElement(m.Path, sepath)
	}
	if m.TileRect != nil {
		setileRect := xml.StartElement{Name: xml.Name{Local: "a:tileRect"}}
		e.EncodeElement(m.TileRect, setileRect)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GradientFillProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "flip" {
			m.FlipAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "rotWithShape" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RotWithShapeAttr = &parsed
		}
	}
lCT_GradientFillProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "gsLst":
				m.GsLst = NewCT_GradientStopList()
				if err := d.DecodeElement(m.GsLst, &el); err != nil {
					return err
				}
			case "lin":
				m.Lin = NewCT_LinearShadeProperties()
				if err := d.DecodeElement(m.Lin, &el); err != nil {
					return err
				}
			case "path":
				m.Path = NewCT_PathShadeProperties()
				if err := d.DecodeElement(m.Path, &el); err != nil {
					return err
				}
			case "tileRect":
				m.TileRect = NewCT_RelativeRect()
				if err := d.DecodeElement(m.TileRect, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_GradientFillProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GradientFillProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GradientFillProperties and its children
func (m *CT_GradientFillProperties) Validate() error {
	return m.ValidateWithPath("CT_GradientFillProperties")
}

// ValidateWithPath validates the CT_GradientFillProperties and its children, prefixing error messages with path
func (m *CT_GradientFillProperties) ValidateWithPath(path string) error {
	if err := m.FlipAttr.ValidateWithPath(path + "/FlipAttr"); err != nil {
		return err
	}
	if m.GsLst != nil {
		if err := m.GsLst.ValidateWithPath(path + "/GsLst"); err != nil {
			return err
		}
	}
	if m.Lin != nil {
		if err := m.Lin.ValidateWithPath(path + "/Lin"); err != nil {
			return err
		}
	}
	if m.Path != nil {
		if err := m.Path.ValidateWithPath(path + "/Path"); err != nil {
			return err
		}
	}
	if m.TileRect != nil {
		if err := m.TileRect.ValidateWithPath(path + "/TileRect"); err != nil {
			return err
		}
	}
	return nil
}
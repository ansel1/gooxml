// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingDrawing

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_Inline struct {
	DistTAttr         *uint32
	DistBAttr         *uint32
	DistLAttr         *uint32
	DistRAttr         *uint32
	Extent            *drawingml.CT_PositiveSize2D
	EffectExtent      *CT_EffectExtent
	DocPr             *drawingml.CT_NonVisualDrawingProps
	CNvGraphicFramePr *drawingml.CT_NonVisualGraphicFrameProperties
	Graphic           *drawingml.Graphic
}

func NewCT_Inline() *CT_Inline {
	ret := &CT_Inline{}
	ret.Extent = drawingml.NewCT_PositiveSize2D()
	ret.DocPr = drawingml.NewCT_NonVisualDrawingProps()
	ret.Graphic = drawingml.NewGraphic()
	return ret
}

func (m *CT_Inline) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.DistTAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distT"},
			Value: fmt.Sprintf("%v", *m.DistTAttr)})
	}
	if m.DistBAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distB"},
			Value: fmt.Sprintf("%v", *m.DistBAttr)})
	}
	if m.DistLAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distL"},
			Value: fmt.Sprintf("%v", *m.DistLAttr)})
	}
	if m.DistRAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distR"},
			Value: fmt.Sprintf("%v", *m.DistRAttr)})
	}
	e.EncodeToken(start)
	seextent := xml.StartElement{Name: xml.Name{Local: "wp:extent"}}
	e.EncodeElement(m.Extent, seextent)
	if m.EffectExtent != nil {
		seeffectExtent := xml.StartElement{Name: xml.Name{Local: "wp:effectExtent"}}
		e.EncodeElement(m.EffectExtent, seeffectExtent)
	}
	sedocPr := xml.StartElement{Name: xml.Name{Local: "wp:docPr"}}
	e.EncodeElement(m.DocPr, sedocPr)
	if m.CNvGraphicFramePr != nil {
		secNvGraphicFramePr := xml.StartElement{Name: xml.Name{Local: "wp:cNvGraphicFramePr"}}
		e.EncodeElement(m.CNvGraphicFramePr, secNvGraphicFramePr)
	}
	segraphic := xml.StartElement{Name: xml.Name{Local: "a:graphic"}}
	e.EncodeElement(m.Graphic, segraphic)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Inline) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Extent = drawingml.NewCT_PositiveSize2D()
	m.DocPr = drawingml.NewCT_NonVisualDrawingProps()
	m.Graphic = drawingml.NewGraphic()
	for _, attr := range start.Attr {
		if attr.Name.Local == "distT" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistTAttr = &pt
		}
		if attr.Name.Local == "distB" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistBAttr = &pt
		}
		if attr.Name.Local == "distL" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistLAttr = &pt
		}
		if attr.Name.Local == "distR" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistRAttr = &pt
		}
	}
lCT_Inline:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "extent":
				if err := d.DecodeElement(m.Extent, &el); err != nil {
					return err
				}
			case "effectExtent":
				m.EffectExtent = NewCT_EffectExtent()
				if err := d.DecodeElement(m.EffectExtent, &el); err != nil {
					return err
				}
			case "docPr":
				if err := d.DecodeElement(m.DocPr, &el); err != nil {
					return err
				}
			case "cNvGraphicFramePr":
				m.CNvGraphicFramePr = drawingml.NewCT_NonVisualGraphicFrameProperties()
				if err := d.DecodeElement(m.CNvGraphicFramePr, &el); err != nil {
					return err
				}
			case "graphic":
				if err := d.DecodeElement(m.Graphic, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Inline %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Inline
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Inline and its children
func (m *CT_Inline) Validate() error {
	return m.ValidateWithPath("CT_Inline")
}

// ValidateWithPath validates the CT_Inline and its children, prefixing error messages with path
func (m *CT_Inline) ValidateWithPath(path string) error {
	if err := m.Extent.ValidateWithPath(path + "/Extent"); err != nil {
		return err
	}
	if m.EffectExtent != nil {
		if err := m.EffectExtent.ValidateWithPath(path + "/EffectExtent"); err != nil {
			return err
		}
	}
	if err := m.DocPr.ValidateWithPath(path + "/DocPr"); err != nil {
		return err
	}
	if m.CNvGraphicFramePr != nil {
		if err := m.CNvGraphicFramePr.ValidateWithPath(path + "/CNvGraphicFramePr"); err != nil {
			return err
		}
	}
	if err := m.Graphic.ValidateWithPath(path + "/Graphic"); err != nil {
		return err
	}
	return nil
}
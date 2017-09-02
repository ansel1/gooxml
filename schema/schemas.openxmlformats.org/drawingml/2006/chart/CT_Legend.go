// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_Legend struct {
	LegendPos   *CT_LegendPos
	LegendEntry []*CT_LegendEntry
	Layout      *CT_Layout
	Overlay     *CT_Boolean
	SpPr        *drawingml.CT_ShapeProperties
	TxPr        *drawingml.CT_TextBody
	ExtLst      *CT_ExtensionList
}

func NewCT_Legend() *CT_Legend {
	ret := &CT_Legend{}
	return ret
}

func (m *CT_Legend) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.LegendPos != nil {
		selegendPos := xml.StartElement{Name: xml.Name{Local: "legendPos"}}
		e.EncodeElement(m.LegendPos, selegendPos)
	}
	if m.LegendEntry != nil {
		selegendEntry := xml.StartElement{Name: xml.Name{Local: "legendEntry"}}
		e.EncodeElement(m.LegendEntry, selegendEntry)
	}
	if m.Layout != nil {
		selayout := xml.StartElement{Name: xml.Name{Local: "layout"}}
		e.EncodeElement(m.Layout, selayout)
	}
	if m.Overlay != nil {
		seoverlay := xml.StartElement{Name: xml.Name{Local: "overlay"}}
		e.EncodeElement(m.Overlay, seoverlay)
	}
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	if m.TxPr != nil {
		setxPr := xml.StartElement{Name: xml.Name{Local: "txPr"}}
		e.EncodeElement(m.TxPr, setxPr)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Legend) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Legend:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "legendPos":
				m.LegendPos = NewCT_LegendPos()
				if err := d.DecodeElement(m.LegendPos, &el); err != nil {
					return err
				}
			case "legendEntry":
				tmp := NewCT_LegendEntry()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.LegendEntry = append(m.LegendEntry, tmp)
			case "layout":
				m.Layout = NewCT_Layout()
				if err := d.DecodeElement(m.Layout, &el); err != nil {
					return err
				}
			case "overlay":
				m.Overlay = NewCT_Boolean()
				if err := d.DecodeElement(m.Overlay, &el); err != nil {
					return err
				}
			case "spPr":
				m.SpPr = drawingml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case "txPr":
				m.TxPr = drawingml.NewCT_TextBody()
				if err := d.DecodeElement(m.TxPr, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Legend %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Legend
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Legend and its children
func (m *CT_Legend) Validate() error {
	return m.ValidateWithPath("CT_Legend")
}

// ValidateWithPath validates the CT_Legend and its children, prefixing error messages with path
func (m *CT_Legend) ValidateWithPath(path string) error {
	if m.LegendPos != nil {
		if err := m.LegendPos.ValidateWithPath(path + "/LegendPos"); err != nil {
			return err
		}
	}
	for i, v := range m.LegendEntry {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/LegendEntry[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.Layout != nil {
		if err := m.Layout.ValidateWithPath(path + "/Layout"); err != nil {
			return err
		}
	}
	if m.Overlay != nil {
		if err := m.Overlay.ValidateWithPath(path + "/Overlay"); err != nil {
			return err
		}
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
			return err
		}
	}
	if m.TxPr != nil {
		if err := m.TxPr.ValidateWithPath(path + "/TxPr"); err != nil {
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
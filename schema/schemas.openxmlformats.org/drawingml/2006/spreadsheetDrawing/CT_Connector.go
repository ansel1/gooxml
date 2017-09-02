// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetDrawing

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_Connector struct {
	MacroAttr      *string
	FPublishedAttr *bool
	NvCxnSpPr      *CT_ConnectorNonVisual
	SpPr           *drawingml.CT_ShapeProperties
	Style          *drawingml.CT_ShapeStyle
}

func NewCT_Connector() *CT_Connector {
	ret := &CT_Connector{}
	ret.NvCxnSpPr = NewCT_ConnectorNonVisual()
	ret.SpPr = drawingml.NewCT_ShapeProperties()
	return ret
}

func (m *CT_Connector) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.MacroAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "macro"},
			Value: fmt.Sprintf("%v", *m.MacroAttr)})
	}
	if m.FPublishedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fPublished"},
			Value: fmt.Sprintf("%v", *m.FPublishedAttr)})
	}
	e.EncodeToken(start)
	senvCxnSpPr := xml.StartElement{Name: xml.Name{Local: "xdr:nvCxnSpPr"}}
	e.EncodeElement(m.NvCxnSpPr, senvCxnSpPr)
	sespPr := xml.StartElement{Name: xml.Name{Local: "xdr:spPr"}}
	e.EncodeElement(m.SpPr, sespPr)
	if m.Style != nil {
		sestyle := xml.StartElement{Name: xml.Name{Local: "xdr:style"}}
		e.EncodeElement(m.Style, sestyle)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Connector) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.NvCxnSpPr = NewCT_ConnectorNonVisual()
	m.SpPr = drawingml.NewCT_ShapeProperties()
	for _, attr := range start.Attr {
		if attr.Name.Local == "macro" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MacroAttr = &parsed
		}
		if attr.Name.Local == "fPublished" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FPublishedAttr = &parsed
		}
	}
lCT_Connector:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "nvCxnSpPr":
				if err := d.DecodeElement(m.NvCxnSpPr, &el); err != nil {
					return err
				}
			case "spPr":
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case "style":
				m.Style = drawingml.NewCT_ShapeStyle()
				if err := d.DecodeElement(m.Style, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Connector %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Connector
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Connector and its children
func (m *CT_Connector) Validate() error {
	return m.ValidateWithPath("CT_Connector")
}

// ValidateWithPath validates the CT_Connector and its children, prefixing error messages with path
func (m *CT_Connector) ValidateWithPath(path string) error {
	if err := m.NvCxnSpPr.ValidateWithPath(path + "/NvCxnSpPr"); err != nil {
		return err
	}
	if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
		return err
	}
	if m.Style != nil {
		if err := m.Style.ValidateWithPath(path + "/Style"); err != nil {
			return err
		}
	}
	return nil
}
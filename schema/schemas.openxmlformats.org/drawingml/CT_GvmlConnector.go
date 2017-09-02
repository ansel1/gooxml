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

type CT_GvmlConnector struct {
	NvCxnSpPr *CT_GvmlConnectorNonVisual
	SpPr      *CT_ShapeProperties
	Style     *CT_ShapeStyle
	ExtLst    *CT_OfficeArtExtensionList
}

func NewCT_GvmlConnector() *CT_GvmlConnector {
	ret := &CT_GvmlConnector{}
	ret.NvCxnSpPr = NewCT_GvmlConnectorNonVisual()
	ret.SpPr = NewCT_ShapeProperties()
	return ret
}

func (m *CT_GvmlConnector) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	senvCxnSpPr := xml.StartElement{Name: xml.Name{Local: "a:nvCxnSpPr"}}
	e.EncodeElement(m.NvCxnSpPr, senvCxnSpPr)
	sespPr := xml.StartElement{Name: xml.Name{Local: "a:spPr"}}
	e.EncodeElement(m.SpPr, sespPr)
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

func (m *CT_GvmlConnector) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.NvCxnSpPr = NewCT_GvmlConnectorNonVisual()
	m.SpPr = NewCT_ShapeProperties()
lCT_GvmlConnector:
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
				log.Printf("skipping unsupported element on CT_GvmlConnector %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GvmlConnector
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GvmlConnector and its children
func (m *CT_GvmlConnector) Validate() error {
	return m.ValidateWithPath("CT_GvmlConnector")
}

// ValidateWithPath validates the CT_GvmlConnector and its children, prefixing error messages with path
func (m *CT_GvmlConnector) ValidateWithPath(path string) error {
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
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
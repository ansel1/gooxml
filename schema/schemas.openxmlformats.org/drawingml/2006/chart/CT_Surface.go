// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_Surface struct {
	Thickness      *CT_Thickness
	SpPr           *drawingml.CT_ShapeProperties
	PictureOptions *CT_PictureOptions
	ExtLst         *CT_ExtensionList
}

func NewCT_Surface() *CT_Surface {
	ret := &CT_Surface{}
	return ret
}

func (m *CT_Surface) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Thickness != nil {
		sethickness := xml.StartElement{Name: xml.Name{Local: "thickness"}}
		e.EncodeElement(m.Thickness, sethickness)
	}
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	if m.PictureOptions != nil {
		sepictureOptions := xml.StartElement{Name: xml.Name{Local: "pictureOptions"}}
		e.EncodeElement(m.PictureOptions, sepictureOptions)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Surface) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Surface:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "thickness":
				m.Thickness = NewCT_Thickness()
				if err := d.DecodeElement(m.Thickness, &el); err != nil {
					return err
				}
			case "spPr":
				m.SpPr = drawingml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case "pictureOptions":
				m.PictureOptions = NewCT_PictureOptions()
				if err := d.DecodeElement(m.PictureOptions, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Surface %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Surface
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Surface and its children
func (m *CT_Surface) Validate() error {
	return m.ValidateWithPath("CT_Surface")
}

// ValidateWithPath validates the CT_Surface and its children, prefixing error messages with path
func (m *CT_Surface) ValidateWithPath(path string) error {
	if m.Thickness != nil {
		if err := m.Thickness.ValidateWithPath(path + "/Thickness"); err != nil {
			return err
		}
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
			return err
		}
	}
	if m.PictureOptions != nil {
		if err := m.PictureOptions.ValidateWithPath(path + "/PictureOptions"); err != nil {
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

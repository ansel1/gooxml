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

type CT_Backdrop struct {
	Anchor *CT_Point3D
	Norm   *CT_Vector3D
	Up     *CT_Vector3D
	ExtLst *CT_OfficeArtExtensionList
}

func NewCT_Backdrop() *CT_Backdrop {
	ret := &CT_Backdrop{}
	ret.Anchor = NewCT_Point3D()
	ret.Norm = NewCT_Vector3D()
	ret.Up = NewCT_Vector3D()
	return ret
}

func (m *CT_Backdrop) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	seanchor := xml.StartElement{Name: xml.Name{Local: "a:anchor"}}
	e.EncodeElement(m.Anchor, seanchor)
	senorm := xml.StartElement{Name: xml.Name{Local: "a:norm"}}
	e.EncodeElement(m.Norm, senorm)
	seup := xml.StartElement{Name: xml.Name{Local: "a:up"}}
	e.EncodeElement(m.Up, seup)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Backdrop) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Anchor = NewCT_Point3D()
	m.Norm = NewCT_Vector3D()
	m.Up = NewCT_Vector3D()
lCT_Backdrop:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "anchor":
				if err := d.DecodeElement(m.Anchor, &el); err != nil {
					return err
				}
			case "norm":
				if err := d.DecodeElement(m.Norm, &el); err != nil {
					return err
				}
			case "up":
				if err := d.DecodeElement(m.Up, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Backdrop %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Backdrop
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Backdrop and its children
func (m *CT_Backdrop) Validate() error {
	return m.ValidateWithPath("CT_Backdrop")
}

// ValidateWithPath validates the CT_Backdrop and its children, prefixing error messages with path
func (m *CT_Backdrop) ValidateWithPath(path string) error {
	if err := m.Anchor.ValidateWithPath(path + "/Anchor"); err != nil {
		return err
	}
	if err := m.Norm.ValidateWithPath(path + "/Norm"); err != nil {
		return err
	}
	if err := m.Up.ValidateWithPath(path + "/Up"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
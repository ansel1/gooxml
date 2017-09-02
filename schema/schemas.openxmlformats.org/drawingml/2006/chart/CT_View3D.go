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
)

type CT_View3D struct {
	RotX         *CT_RotX
	HPercent     *CT_HPercent
	RotY         *CT_RotY
	DepthPercent *CT_DepthPercent
	RAngAx       *CT_Boolean
	Perspective  *CT_Perspective
	ExtLst       *CT_ExtensionList
}

func NewCT_View3D() *CT_View3D {
	ret := &CT_View3D{}
	return ret
}

func (m *CT_View3D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.RotX != nil {
		serotX := xml.StartElement{Name: xml.Name{Local: "rotX"}}
		e.EncodeElement(m.RotX, serotX)
	}
	if m.HPercent != nil {
		sehPercent := xml.StartElement{Name: xml.Name{Local: "hPercent"}}
		e.EncodeElement(m.HPercent, sehPercent)
	}
	if m.RotY != nil {
		serotY := xml.StartElement{Name: xml.Name{Local: "rotY"}}
		e.EncodeElement(m.RotY, serotY)
	}
	if m.DepthPercent != nil {
		sedepthPercent := xml.StartElement{Name: xml.Name{Local: "depthPercent"}}
		e.EncodeElement(m.DepthPercent, sedepthPercent)
	}
	if m.RAngAx != nil {
		serAngAx := xml.StartElement{Name: xml.Name{Local: "rAngAx"}}
		e.EncodeElement(m.RAngAx, serAngAx)
	}
	if m.Perspective != nil {
		seperspective := xml.StartElement{Name: xml.Name{Local: "perspective"}}
		e.EncodeElement(m.Perspective, seperspective)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_View3D) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_View3D:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "rotX":
				m.RotX = NewCT_RotX()
				if err := d.DecodeElement(m.RotX, &el); err != nil {
					return err
				}
			case "hPercent":
				m.HPercent = NewCT_HPercent()
				if err := d.DecodeElement(m.HPercent, &el); err != nil {
					return err
				}
			case "rotY":
				m.RotY = NewCT_RotY()
				if err := d.DecodeElement(m.RotY, &el); err != nil {
					return err
				}
			case "depthPercent":
				m.DepthPercent = NewCT_DepthPercent()
				if err := d.DecodeElement(m.DepthPercent, &el); err != nil {
					return err
				}
			case "rAngAx":
				m.RAngAx = NewCT_Boolean()
				if err := d.DecodeElement(m.RAngAx, &el); err != nil {
					return err
				}
			case "perspective":
				m.Perspective = NewCT_Perspective()
				if err := d.DecodeElement(m.Perspective, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_View3D %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_View3D
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_View3D and its children
func (m *CT_View3D) Validate() error {
	return m.ValidateWithPath("CT_View3D")
}

// ValidateWithPath validates the CT_View3D and its children, prefixing error messages with path
func (m *CT_View3D) ValidateWithPath(path string) error {
	if m.RotX != nil {
		if err := m.RotX.ValidateWithPath(path + "/RotX"); err != nil {
			return err
		}
	}
	if m.HPercent != nil {
		if err := m.HPercent.ValidateWithPath(path + "/HPercent"); err != nil {
			return err
		}
	}
	if m.RotY != nil {
		if err := m.RotY.ValidateWithPath(path + "/RotY"); err != nil {
			return err
		}
	}
	if m.DepthPercent != nil {
		if err := m.DepthPercent.ValidateWithPath(path + "/DepthPercent"); err != nil {
			return err
		}
	}
	if m.RAngAx != nil {
		if err := m.RAngAx.ValidateWithPath(path + "/RAngAx"); err != nil {
			return err
		}
	}
	if m.Perspective != nil {
		if err := m.Perspective.ValidateWithPath(path + "/Perspective"); err != nil {
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

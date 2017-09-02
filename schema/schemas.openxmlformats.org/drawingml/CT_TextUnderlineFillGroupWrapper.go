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

type CT_TextUnderlineFillGroupWrapper struct {
	NoFill    *CT_NoFillProperties
	SolidFill *CT_SolidColorFillProperties
	GradFill  *CT_GradientFillProperties
	BlipFill  *CT_BlipFillProperties
	PattFill  *CT_PatternFillProperties
	GrpFill   *CT_GroupFillProperties
}

func NewCT_TextUnderlineFillGroupWrapper() *CT_TextUnderlineFillGroupWrapper {
	ret := &CT_TextUnderlineFillGroupWrapper{}
	return ret
}

func (m *CT_TextUnderlineFillGroupWrapper) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.NoFill != nil {
		senoFill := xml.StartElement{Name: xml.Name{Local: "a:noFill"}}
		e.EncodeElement(m.NoFill, senoFill)
	}
	if m.SolidFill != nil {
		sesolidFill := xml.StartElement{Name: xml.Name{Local: "a:solidFill"}}
		e.EncodeElement(m.SolidFill, sesolidFill)
	}
	if m.GradFill != nil {
		segradFill := xml.StartElement{Name: xml.Name{Local: "a:gradFill"}}
		e.EncodeElement(m.GradFill, segradFill)
	}
	if m.BlipFill != nil {
		seblipFill := xml.StartElement{Name: xml.Name{Local: "a:blipFill"}}
		e.EncodeElement(m.BlipFill, seblipFill)
	}
	if m.PattFill != nil {
		sepattFill := xml.StartElement{Name: xml.Name{Local: "a:pattFill"}}
		e.EncodeElement(m.PattFill, sepattFill)
	}
	if m.GrpFill != nil {
		segrpFill := xml.StartElement{Name: xml.Name{Local: "a:grpFill"}}
		e.EncodeElement(m.GrpFill, segrpFill)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextUnderlineFillGroupWrapper) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TextUnderlineFillGroupWrapper:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "noFill":
				m.NoFill = NewCT_NoFillProperties()
				if err := d.DecodeElement(m.NoFill, &el); err != nil {
					return err
				}
			case "solidFill":
				m.SolidFill = NewCT_SolidColorFillProperties()
				if err := d.DecodeElement(m.SolidFill, &el); err != nil {
					return err
				}
			case "gradFill":
				m.GradFill = NewCT_GradientFillProperties()
				if err := d.DecodeElement(m.GradFill, &el); err != nil {
					return err
				}
			case "blipFill":
				m.BlipFill = NewCT_BlipFillProperties()
				if err := d.DecodeElement(m.BlipFill, &el); err != nil {
					return err
				}
			case "pattFill":
				m.PattFill = NewCT_PatternFillProperties()
				if err := d.DecodeElement(m.PattFill, &el); err != nil {
					return err
				}
			case "grpFill":
				m.GrpFill = NewCT_GroupFillProperties()
				if err := d.DecodeElement(m.GrpFill, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_TextUnderlineFillGroupWrapper %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TextUnderlineFillGroupWrapper
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TextUnderlineFillGroupWrapper and its children
func (m *CT_TextUnderlineFillGroupWrapper) Validate() error {
	return m.ValidateWithPath("CT_TextUnderlineFillGroupWrapper")
}

// ValidateWithPath validates the CT_TextUnderlineFillGroupWrapper and its children, prefixing error messages with path
func (m *CT_TextUnderlineFillGroupWrapper) ValidateWithPath(path string) error {
	if m.NoFill != nil {
		if err := m.NoFill.ValidateWithPath(path + "/NoFill"); err != nil {
			return err
		}
	}
	if m.SolidFill != nil {
		if err := m.SolidFill.ValidateWithPath(path + "/SolidFill"); err != nil {
			return err
		}
	}
	if m.GradFill != nil {
		if err := m.GradFill.ValidateWithPath(path + "/GradFill"); err != nil {
			return err
		}
	}
	if m.BlipFill != nil {
		if err := m.BlipFill.ValidateWithPath(path + "/BlipFill"); err != nil {
			return err
		}
	}
	if m.PattFill != nil {
		if err := m.PattFill.ValidateWithPath(path + "/PattFill"); err != nil {
			return err
		}
	}
	if m.GrpFill != nil {
		if err := m.GrpFill.ValidateWithPath(path + "/GrpFill"); err != nil {
			return err
		}
	}
	return nil
}

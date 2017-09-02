// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chartDrawing

import (
	"encoding/xml"
	"log"
)

type EG_ObjectChoicesChoice struct {
	Sp           *CT_Shape
	GrpSp        *CT_GroupShape
	GraphicFrame *CT_GraphicFrame
	CxnSp        *CT_Connector
	Pic          *CT_Picture
}

func NewEG_ObjectChoicesChoice() *EG_ObjectChoicesChoice {
	ret := &EG_ObjectChoicesChoice{}
	return ret
}

func (m *EG_ObjectChoicesChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Sp != nil {
		sesp := xml.StartElement{Name: xml.Name{Local: "sp"}}
		e.EncodeElement(m.Sp, sesp)
	}
	if m.GrpSp != nil {
		segrpSp := xml.StartElement{Name: xml.Name{Local: "grpSp"}}
		e.EncodeElement(m.GrpSp, segrpSp)
	}
	if m.GraphicFrame != nil {
		segraphicFrame := xml.StartElement{Name: xml.Name{Local: "graphicFrame"}}
		e.EncodeElement(m.GraphicFrame, segraphicFrame)
	}
	if m.CxnSp != nil {
		secxnSp := xml.StartElement{Name: xml.Name{Local: "cxnSp"}}
		e.EncodeElement(m.CxnSp, secxnSp)
	}
	if m.Pic != nil {
		sepic := xml.StartElement{Name: xml.Name{Local: "pic"}}
		e.EncodeElement(m.Pic, sepic)
	}
	return nil
}

func (m *EG_ObjectChoicesChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_ObjectChoicesChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sp":
				m.Sp = NewCT_Shape()
				if err := d.DecodeElement(m.Sp, &el); err != nil {
					return err
				}
			case "grpSp":
				m.GrpSp = NewCT_GroupShape()
				if err := d.DecodeElement(m.GrpSp, &el); err != nil {
					return err
				}
			case "graphicFrame":
				m.GraphicFrame = NewCT_GraphicFrame()
				if err := d.DecodeElement(m.GraphicFrame, &el); err != nil {
					return err
				}
			case "cxnSp":
				m.CxnSp = NewCT_Connector()
				if err := d.DecodeElement(m.CxnSp, &el); err != nil {
					return err
				}
			case "pic":
				m.Pic = NewCT_Picture()
				if err := d.DecodeElement(m.Pic, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on EG_ObjectChoicesChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_ObjectChoicesChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ObjectChoicesChoice and its children
func (m *EG_ObjectChoicesChoice) Validate() error {
	return m.ValidateWithPath("EG_ObjectChoicesChoice")
}

// ValidateWithPath validates the EG_ObjectChoicesChoice and its children, prefixing error messages with path
func (m *EG_ObjectChoicesChoice) ValidateWithPath(path string) error {
	if m.Sp != nil {
		if err := m.Sp.ValidateWithPath(path + "/Sp"); err != nil {
			return err
		}
	}
	if m.GrpSp != nil {
		if err := m.GrpSp.ValidateWithPath(path + "/GrpSp"); err != nil {
			return err
		}
	}
	if m.GraphicFrame != nil {
		if err := m.GraphicFrame.ValidateWithPath(path + "/GraphicFrame"); err != nil {
			return err
		}
	}
	if m.CxnSp != nil {
		if err := m.CxnSp.ValidateWithPath(path + "/CxnSp"); err != nil {
			return err
		}
	}
	if m.Pic != nil {
		if err := m.Pic.ValidateWithPath(path + "/Pic"); err != nil {
			return err
		}
	}
	return nil
}

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

type CT_AnimationElementChoice struct {
	Dgm   *CT_AnimationDgmElement
	Chart *CT_AnimationChartElement
}

func NewCT_AnimationElementChoice() *CT_AnimationElementChoice {
	ret := &CT_AnimationElementChoice{}
	return ret
}

func (m *CT_AnimationElementChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:CT_AnimationElementChoice"
	e.EncodeToken(start)
	if m.Dgm != nil {
		sedgm := xml.StartElement{Name: xml.Name{Local: "a:dgm"}}
		e.EncodeElement(m.Dgm, sedgm)
	}
	if m.Chart != nil {
		sechart := xml.StartElement{Name: xml.Name{Local: "a:chart"}}
		e.EncodeElement(m.Chart, sechart)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AnimationElementChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_AnimationElementChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "dgm":
				m.Dgm = NewCT_AnimationDgmElement()
				if err := d.DecodeElement(m.Dgm, &el); err != nil {
					return err
				}
			case "chart":
				m.Chart = NewCT_AnimationChartElement()
				if err := d.DecodeElement(m.Chart, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_AnimationElementChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AnimationElementChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AnimationElementChoice and its children
func (m *CT_AnimationElementChoice) Validate() error {
	return m.ValidateWithPath("CT_AnimationElementChoice")
}

// ValidateWithPath validates the CT_AnimationElementChoice and its children, prefixing error messages with path
func (m *CT_AnimationElementChoice) ValidateWithPath(path string) error {
	if m.Dgm != nil {
		if err := m.Dgm.ValidateWithPath(path + "/Dgm"); err != nil {
			return err
		}
	}
	if m.Chart != nil {
		if err := m.Chart.ValidateWithPath(path + "/Chart"); err != nil {
			return err
		}
	}
	return nil
}

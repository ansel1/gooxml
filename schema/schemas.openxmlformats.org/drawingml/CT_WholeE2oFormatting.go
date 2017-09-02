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

type CT_WholeE2oFormatting struct {
	Ln        *CT_LineProperties
	EffectLst *CT_EffectList
	EffectDag *CT_EffectContainer
}

func NewCT_WholeE2oFormatting() *CT_WholeE2oFormatting {
	ret := &CT_WholeE2oFormatting{}
	return ret
}

func (m *CT_WholeE2oFormatting) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:CT_WholeE2oFormatting"
	e.EncodeToken(start)
	if m.Ln != nil {
		seln := xml.StartElement{Name: xml.Name{Local: "a:ln"}}
		e.EncodeElement(m.Ln, seln)
	}
	if m.EffectLst != nil {
		seeffectLst := xml.StartElement{Name: xml.Name{Local: "a:effectLst"}}
		e.EncodeElement(m.EffectLst, seeffectLst)
	}
	if m.EffectDag != nil {
		seeffectDag := xml.StartElement{Name: xml.Name{Local: "a:effectDag"}}
		e.EncodeElement(m.EffectDag, seeffectDag)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_WholeE2oFormatting) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_WholeE2oFormatting:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "ln":
				m.Ln = NewCT_LineProperties()
				if err := d.DecodeElement(m.Ln, &el); err != nil {
					return err
				}
			case "effectLst":
				m.EffectLst = NewCT_EffectList()
				if err := d.DecodeElement(m.EffectLst, &el); err != nil {
					return err
				}
			case "effectDag":
				m.EffectDag = NewCT_EffectContainer()
				if err := d.DecodeElement(m.EffectDag, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_WholeE2oFormatting %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_WholeE2oFormatting
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_WholeE2oFormatting and its children
func (m *CT_WholeE2oFormatting) Validate() error {
	return m.ValidateWithPath("CT_WholeE2oFormatting")
}

// ValidateWithPath validates the CT_WholeE2oFormatting and its children, prefixing error messages with path
func (m *CT_WholeE2oFormatting) ValidateWithPath(path string) error {
	if m.Ln != nil {
		if err := m.Ln.ValidateWithPath(path + "/Ln"); err != nil {
			return err
		}
	}
	if m.EffectLst != nil {
		if err := m.EffectLst.ValidateWithPath(path + "/EffectLst"); err != nil {
			return err
		}
	}
	if m.EffectDag != nil {
		if err := m.EffectDag.ValidateWithPath(path + "/EffectDag"); err != nil {
			return err
		}
	}
	return nil
}
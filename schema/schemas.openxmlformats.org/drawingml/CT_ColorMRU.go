// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_ColorMRU struct {
	EG_ColorChoice []*EG_ColorChoice
}

func NewCT_ColorMRU() *CT_ColorMRU {
	ret := &CT_ColorMRU{}
	return ret
}

func (m *CT_ColorMRU) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:CT_ColorMRU"
	e.EncodeToken(start)
	if m.EG_ColorChoice != nil {
		for _, c := range m.EG_ColorChoice {
			c.MarshalXML(e, start)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ColorMRU) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ColorMRU:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "scrgbClr":
				tmpcolorchoice := NewEG_ColorChoice()
				tmpcolorchoice.ScrgbClr = NewCT_ScRgbColor()
				if err := d.DecodeElement(tmpcolorchoice.ScrgbClr, &el); err != nil {
					return err
				}
				m.EG_ColorChoice = append(m.EG_ColorChoice, tmpcolorchoice)
			case "srgbClr":
				tmpcolorchoice := NewEG_ColorChoice()
				tmpcolorchoice.SrgbClr = NewCT_SRgbColor()
				if err := d.DecodeElement(tmpcolorchoice.SrgbClr, &el); err != nil {
					return err
				}
				m.EG_ColorChoice = append(m.EG_ColorChoice, tmpcolorchoice)
			case "hslClr":
				tmpcolorchoice := NewEG_ColorChoice()
				tmpcolorchoice.HslClr = NewCT_HslColor()
				if err := d.DecodeElement(tmpcolorchoice.HslClr, &el); err != nil {
					return err
				}
				m.EG_ColorChoice = append(m.EG_ColorChoice, tmpcolorchoice)
			case "sysClr":
				tmpcolorchoice := NewEG_ColorChoice()
				tmpcolorchoice.SysClr = NewCT_SystemColor()
				if err := d.DecodeElement(tmpcolorchoice.SysClr, &el); err != nil {
					return err
				}
				m.EG_ColorChoice = append(m.EG_ColorChoice, tmpcolorchoice)
			case "schemeClr":
				tmpcolorchoice := NewEG_ColorChoice()
				tmpcolorchoice.SchemeClr = NewCT_SchemeColor()
				if err := d.DecodeElement(tmpcolorchoice.SchemeClr, &el); err != nil {
					return err
				}
				m.EG_ColorChoice = append(m.EG_ColorChoice, tmpcolorchoice)
			case "prstClr":
				tmpcolorchoice := NewEG_ColorChoice()
				tmpcolorchoice.PrstClr = NewCT_PresetColor()
				if err := d.DecodeElement(tmpcolorchoice.PrstClr, &el); err != nil {
					return err
				}
				m.EG_ColorChoice = append(m.EG_ColorChoice, tmpcolorchoice)
			default:
				log.Printf("skipping unsupported element on CT_ColorMRU %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ColorMRU
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ColorMRU and its children
func (m *CT_ColorMRU) Validate() error {
	return m.ValidateWithPath("CT_ColorMRU")
}

// ValidateWithPath validates the CT_ColorMRU and its children, prefixing error messages with path
func (m *CT_ColorMRU) ValidateWithPath(path string) error {
	for i, v := range m.EG_ColorChoice {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_ColorChoice[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}

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

type CT_TextParagraph struct {
	PPr        *CT_TextParagraphProperties
	EG_TextRun []*EG_TextRun
	EndParaRPr *CT_TextCharacterProperties
}

func NewCT_TextParagraph() *CT_TextParagraph {
	ret := &CT_TextParagraph{}
	return ret
}

func (m *CT_TextParagraph) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.PPr != nil {
		sepPr := xml.StartElement{Name: xml.Name{Local: "a:pPr"}}
		e.EncodeElement(m.PPr, sepPr)
	}
	if m.EG_TextRun != nil {
		for _, c := range m.EG_TextRun {
			c.MarshalXML(e, start)
		}
	}
	if m.EndParaRPr != nil {
		seendParaRPr := xml.StartElement{Name: xml.Name{Local: "a:endParaRPr"}}
		e.EncodeElement(m.EndParaRPr, seendParaRPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextParagraph) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TextParagraph:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "pPr":
				m.PPr = NewCT_TextParagraphProperties()
				if err := d.DecodeElement(m.PPr, &el); err != nil {
					return err
				}
			case "r":
				tmptextrun := NewEG_TextRun()
				tmptextrun.R = NewCT_RegularTextRun()
				if err := d.DecodeElement(tmptextrun.R, &el); err != nil {
					return err
				}
				m.EG_TextRun = append(m.EG_TextRun, tmptextrun)
			case "br":
				tmptextrun := NewEG_TextRun()
				tmptextrun.Br = NewCT_TextLineBreak()
				if err := d.DecodeElement(tmptextrun.Br, &el); err != nil {
					return err
				}
				m.EG_TextRun = append(m.EG_TextRun, tmptextrun)
			case "fld":
				tmptextrun := NewEG_TextRun()
				tmptextrun.Fld = NewCT_TextField()
				if err := d.DecodeElement(tmptextrun.Fld, &el); err != nil {
					return err
				}
				m.EG_TextRun = append(m.EG_TextRun, tmptextrun)
			case "endParaRPr":
				m.EndParaRPr = NewCT_TextCharacterProperties()
				if err := d.DecodeElement(m.EndParaRPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_TextParagraph %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TextParagraph
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TextParagraph and its children
func (m *CT_TextParagraph) Validate() error {
	return m.ValidateWithPath("CT_TextParagraph")
}

// ValidateWithPath validates the CT_TextParagraph and its children, prefixing error messages with path
func (m *CT_TextParagraph) ValidateWithPath(path string) error {
	if m.PPr != nil {
		if err := m.PPr.ValidateWithPath(path + "/PPr"); err != nil {
			return err
		}
	}
	for i, v := range m.EG_TextRun {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_TextRun[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.EndParaRPr != nil {
		if err := m.EndParaRPr.ValidateWithPath(path + "/EndParaRPr"); err != nil {
			return err
		}
	}
	return nil
}

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

type CT_TextBody struct {
	BodyPr   *CT_TextBodyProperties
	LstStyle *CT_TextListStyle
	P        []*CT_TextParagraph
}

func NewCT_TextBody() *CT_TextBody {
	ret := &CT_TextBody{}
	ret.BodyPr = NewCT_TextBodyProperties()
	return ret
}

func (m *CT_TextBody) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sebodyPr := xml.StartElement{Name: xml.Name{Local: "a:bodyPr"}}
	e.EncodeElement(m.BodyPr, sebodyPr)
	if m.LstStyle != nil {
		selstStyle := xml.StartElement{Name: xml.Name{Local: "a:lstStyle"}}
		e.EncodeElement(m.LstStyle, selstStyle)
	}
	sep := xml.StartElement{Name: xml.Name{Local: "a:p"}}
	e.EncodeElement(m.P, sep)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.BodyPr = NewCT_TextBodyProperties()
lCT_TextBody:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "bodyPr":
				if err := d.DecodeElement(m.BodyPr, &el); err != nil {
					return err
				}
			case "lstStyle":
				m.LstStyle = NewCT_TextListStyle()
				if err := d.DecodeElement(m.LstStyle, &el); err != nil {
					return err
				}
			case "p":
				tmp := NewCT_TextParagraph()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.P = append(m.P, tmp)
			default:
				log.Printf("skipping unsupported element on CT_TextBody %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TextBody
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TextBody and its children
func (m *CT_TextBody) Validate() error {
	return m.ValidateWithPath("CT_TextBody")
}

// ValidateWithPath validates the CT_TextBody and its children, prefixing error messages with path
func (m *CT_TextBody) ValidateWithPath(path string) error {
	if err := m.BodyPr.ValidateWithPath(path + "/BodyPr"); err != nil {
		return err
	}
	if m.LstStyle != nil {
		if err := m.LstStyle.ValidateWithPath(path + "/LstStyle"); err != nil {
			return err
		}
	}
	for i, v := range m.P {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/P[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}

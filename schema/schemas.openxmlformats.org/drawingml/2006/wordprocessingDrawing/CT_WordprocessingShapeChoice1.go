// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingDrawing

import (
	"encoding/xml"
	"log"
)

type CT_WordprocessingShapeChoice1 struct {
	Txbx       *CT_TextboxInfo
	LinkedTxbx *CT_LinkedTextboxInformation
}

func NewCT_WordprocessingShapeChoice1() *CT_WordprocessingShapeChoice1 {
	ret := &CT_WordprocessingShapeChoice1{}
	return ret
}

func (m *CT_WordprocessingShapeChoice1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Txbx != nil {
		setxbx := xml.StartElement{Name: xml.Name{Local: "wp:txbx"}}
		e.EncodeElement(m.Txbx, setxbx)
	}
	if m.LinkedTxbx != nil {
		selinkedTxbx := xml.StartElement{Name: xml.Name{Local: "wp:linkedTxbx"}}
		e.EncodeElement(m.LinkedTxbx, selinkedTxbx)
	}
	return nil
}

func (m *CT_WordprocessingShapeChoice1) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_WordprocessingShapeChoice1:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "txbx":
				m.Txbx = NewCT_TextboxInfo()
				if err := d.DecodeElement(m.Txbx, &el); err != nil {
					return err
				}
			case "linkedTxbx":
				m.LinkedTxbx = NewCT_LinkedTextboxInformation()
				if err := d.DecodeElement(m.LinkedTxbx, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_WordprocessingShapeChoice1 %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_WordprocessingShapeChoice1
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_WordprocessingShapeChoice1 and its children
func (m *CT_WordprocessingShapeChoice1) Validate() error {
	return m.ValidateWithPath("CT_WordprocessingShapeChoice1")
}

// ValidateWithPath validates the CT_WordprocessingShapeChoice1 and its children, prefixing error messages with path
func (m *CT_WordprocessingShapeChoice1) ValidateWithPath(path string) error {
	if m.Txbx != nil {
		if err := m.Txbx.ValidateWithPath(path + "/Txbx"); err != nil {
			return err
		}
	}
	if m.LinkedTxbx != nil {
		if err := m.LinkedTxbx.ValidateWithPath(path + "/LinkedTxbx"); err != nil {
			return err
		}
	}
	return nil
}
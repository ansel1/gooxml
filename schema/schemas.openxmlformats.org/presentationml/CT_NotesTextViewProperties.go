// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"log"
)

type CT_NotesTextViewProperties struct {
	// Base properties for Notes View
	CViewPr *CT_CommonViewProperties
	ExtLst  *CT_ExtensionList
}

func NewCT_NotesTextViewProperties() *CT_NotesTextViewProperties {
	ret := &CT_NotesTextViewProperties{}
	ret.CViewPr = NewCT_CommonViewProperties()
	return ret
}

func (m *CT_NotesTextViewProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secViewPr := xml.StartElement{Name: xml.Name{Local: "p:cViewPr"}}
	e.EncodeElement(m.CViewPr, secViewPr)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_NotesTextViewProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CViewPr = NewCT_CommonViewProperties()
lCT_NotesTextViewProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cViewPr":
				if err := d.DecodeElement(m.CViewPr, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_NotesTextViewProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_NotesTextViewProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_NotesTextViewProperties and its children
func (m *CT_NotesTextViewProperties) Validate() error {
	return m.ValidateWithPath("CT_NotesTextViewProperties")
}

// ValidateWithPath validates the CT_NotesTextViewProperties and its children, prefixing error messages with path
func (m *CT_NotesTextViewProperties) ValidateWithPath(path string) error {
	if err := m.CViewPr.ValidateWithPath(path + "/CViewPr"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}

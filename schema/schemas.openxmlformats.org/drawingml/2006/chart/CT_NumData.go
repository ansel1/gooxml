// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"
	"log"

	"baliance.com/gooxml"
)

type CT_NumData struct {
	FormatCode *string
	PtCount    *CT_UnsignedInt
	Pt         []*CT_NumVal
	ExtLst     *CT_ExtensionList
}

func NewCT_NumData() *CT_NumData {
	ret := &CT_NumData{}
	return ret
}

func (m *CT_NumData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.FormatCode != nil {
		seformatCode := xml.StartElement{Name: xml.Name{Local: "formatCode"}}
		gooxml.AddPreserveSpaceAttr(&seformatCode, *m.FormatCode)
		e.EncodeElement(m.FormatCode, seformatCode)
	}
	if m.PtCount != nil {
		septCount := xml.StartElement{Name: xml.Name{Local: "ptCount"}}
		e.EncodeElement(m.PtCount, septCount)
	}
	if m.Pt != nil {
		sept := xml.StartElement{Name: xml.Name{Local: "pt"}}
		e.EncodeElement(m.Pt, sept)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_NumData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_NumData:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "formatCode":
				m.FormatCode = new(string)
				if err := d.DecodeElement(m.FormatCode, &el); err != nil {
					return err
				}
			case "ptCount":
				m.PtCount = NewCT_UnsignedInt()
				if err := d.DecodeElement(m.PtCount, &el); err != nil {
					return err
				}
			case "pt":
				tmp := NewCT_NumVal()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Pt = append(m.Pt, tmp)
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_NumData %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_NumData
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_NumData and its children
func (m *CT_NumData) Validate() error {
	return m.ValidateWithPath("CT_NumData")
}

// ValidateWithPath validates the CT_NumData and its children, prefixing error messages with path
func (m *CT_NumData) ValidateWithPath(path string) error {
	if m.PtCount != nil {
		if err := m.PtCount.ValidateWithPath(path + "/PtCount"); err != nil {
			return err
		}
	}
	for i, v := range m.Pt {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Pt[%d]", path, i)); err != nil {
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

// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_FtnDocProps struct {
	// Footnote Placement
	Pos *CT_FtnPos
	// Footnote Numbering Format
	NumFmt *CT_NumFmt
	// Footnote and Endnote Numbering Starting Value
	NumStart *CT_DecimalNumber
	// Footnote and Endnote Numbering Restart Location
	NumRestart *CT_NumRestart
	Footnote   []*CT_FtnEdnSepRef
}

func NewCT_FtnDocProps() *CT_FtnDocProps {
	ret := &CT_FtnDocProps{}
	return ret
}

func (m *CT_FtnDocProps) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Pos != nil {
		sepos := xml.StartElement{Name: xml.Name{Local: "w:pos"}}
		e.EncodeElement(m.Pos, sepos)
	}
	if m.NumFmt != nil {
		senumFmt := xml.StartElement{Name: xml.Name{Local: "w:numFmt"}}
		e.EncodeElement(m.NumFmt, senumFmt)
	}
	if m.NumStart != nil {
		senumStart := xml.StartElement{Name: xml.Name{Local: "w:numStart"}}
		e.EncodeElement(m.NumStart, senumStart)
	}
	if m.NumRestart != nil {
		senumRestart := xml.StartElement{Name: xml.Name{Local: "w:numRestart"}}
		e.EncodeElement(m.NumRestart, senumRestart)
	}
	if m.Footnote != nil {
		sefootnote := xml.StartElement{Name: xml.Name{Local: "w:footnote"}}
		e.EncodeElement(m.Footnote, sefootnote)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FtnDocProps) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_FtnDocProps:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "pos":
				m.Pos = NewCT_FtnPos()
				if err := d.DecodeElement(m.Pos, &el); err != nil {
					return err
				}
			case "numFmt":
				m.NumFmt = NewCT_NumFmt()
				if err := d.DecodeElement(m.NumFmt, &el); err != nil {
					return err
				}
			case "numStart":
				m.NumStart = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.NumStart, &el); err != nil {
					return err
				}
			case "numRestart":
				m.NumRestart = NewCT_NumRestart()
				if err := d.DecodeElement(m.NumRestart, &el); err != nil {
					return err
				}
			case "footnote":
				tmp := NewCT_FtnEdnSepRef()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Footnote = append(m.Footnote, tmp)
			default:
				log.Printf("skipping unsupported element on CT_FtnDocProps %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FtnDocProps
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_FtnDocProps and its children
func (m *CT_FtnDocProps) Validate() error {
	return m.ValidateWithPath("CT_FtnDocProps")
}

// ValidateWithPath validates the CT_FtnDocProps and its children, prefixing error messages with path
func (m *CT_FtnDocProps) ValidateWithPath(path string) error {
	if m.Pos != nil {
		if err := m.Pos.ValidateWithPath(path + "/Pos"); err != nil {
			return err
		}
	}
	if m.NumFmt != nil {
		if err := m.NumFmt.ValidateWithPath(path + "/NumFmt"); err != nil {
			return err
		}
	}
	if m.NumStart != nil {
		if err := m.NumStart.ValidateWithPath(path + "/NumStart"); err != nil {
			return err
		}
	}
	if m.NumRestart != nil {
		if err := m.NumRestart.ValidateWithPath(path + "/NumRestart"); err != nil {
			return err
		}
	}
	for i, v := range m.Footnote {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Footnote[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}

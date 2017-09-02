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

type CT_OleObjectChoice struct {
	Embed *CT_OleObjectEmbed
	Link  *CT_OleObjectLink
}

func NewCT_OleObjectChoice() *CT_OleObjectChoice {
	ret := &CT_OleObjectChoice{}
	return ret
}

func (m *CT_OleObjectChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Embed != nil {
		seembed := xml.StartElement{Name: xml.Name{Local: "p:embed"}}
		e.EncodeElement(m.Embed, seembed)
	}
	if m.Link != nil {
		selink := xml.StartElement{Name: xml.Name{Local: "p:link"}}
		e.EncodeElement(m.Link, selink)
	}
	return nil
}

func (m *CT_OleObjectChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_OleObjectChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "embed":
				m.Embed = NewCT_OleObjectEmbed()
				if err := d.DecodeElement(m.Embed, &el); err != nil {
					return err
				}
			case "link":
				m.Link = NewCT_OleObjectLink()
				if err := d.DecodeElement(m.Link, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_OleObjectChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OleObjectChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OleObjectChoice and its children
func (m *CT_OleObjectChoice) Validate() error {
	return m.ValidateWithPath("CT_OleObjectChoice")
}

// ValidateWithPath validates the CT_OleObjectChoice and its children, prefixing error messages with path
func (m *CT_OleObjectChoice) ValidateWithPath(path string) error {
	if m.Embed != nil {
		if err := m.Embed.ValidateWithPath(path + "/Embed"); err != nil {
			return err
		}
	}
	if m.Link != nil {
		if err := m.Link.ValidateWithPath(path + "/Link"); err != nil {
			return err
		}
	}
	return nil
}

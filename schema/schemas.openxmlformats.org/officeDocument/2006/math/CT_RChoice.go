// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_RChoice struct {
	T []*CT_Text
}

func NewCT_RChoice() *CT_RChoice {
	ret := &CT_RChoice{}
	return ret
}

func (m *CT_RChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.T != nil {
		set := xml.StartElement{Name: xml.Name{Local: "m:t"}}
		e.EncodeElement(m.T, set)
	}
	return nil
}

func (m *CT_RChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_RChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "t":
				tmp := NewCT_Text()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.T = append(m.T, tmp)
			default:
				log.Printf("skipping unsupported element on CT_RChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_RChoice and its children
func (m *CT_RChoice) Validate() error {
	return m.ValidateWithPath("CT_RChoice")
}

// ValidateWithPath validates the CT_RChoice and its children, prefixing error messages with path
func (m *CT_RChoice) ValidateWithPath(path string) error {
	for i, v := range m.T {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/T[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
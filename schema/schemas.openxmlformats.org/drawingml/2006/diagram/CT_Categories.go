// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_Categories struct {
	Cat []*CT_Category
}

func NewCT_Categories() *CT_Categories {
	ret := &CT_Categories{}
	return ret
}

func (m *CT_Categories) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Cat != nil {
		secat := xml.StartElement{Name: xml.Name{Local: "cat"}}
		e.EncodeElement(m.Cat, secat)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Categories) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Categories:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cat":
				tmp := NewCT_Category()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Cat = append(m.Cat, tmp)
			default:
				log.Printf("skipping unsupported element on CT_Categories %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Categories
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Categories and its children
func (m *CT_Categories) Validate() error {
	return m.ValidateWithPath("CT_Categories")
}

// ValidateWithPath validates the CT_Categories and its children, prefixing error messages with path
func (m *CT_Categories) ValidateWithPath(path string) error {
	for i, v := range m.Cat {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Cat[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
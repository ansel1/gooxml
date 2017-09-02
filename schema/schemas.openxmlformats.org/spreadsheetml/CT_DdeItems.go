// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_DdeItems struct {
	// DDE Item definition
	DdeItem []*CT_DdeItem
}

func NewCT_DdeItems() *CT_DdeItems {
	ret := &CT_DdeItems{}
	return ret
}

func (m *CT_DdeItems) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.DdeItem != nil {
		seddeItem := xml.StartElement{Name: xml.Name{Local: "x:ddeItem"}}
		e.EncodeElement(m.DdeItem, seddeItem)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DdeItems) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DdeItems:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "ddeItem":
				tmp := NewCT_DdeItem()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.DdeItem = append(m.DdeItem, tmp)
			default:
				log.Printf("skipping unsupported element on CT_DdeItems %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DdeItems
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DdeItems and its children
func (m *CT_DdeItems) Validate() error {
	return m.ValidateWithPath("CT_DdeItems")
}

// ValidateWithPath validates the CT_DdeItems and its children, prefixing error messages with path
func (m *CT_DdeItems) ValidateWithPath(path string) error {
	for i, v := range m.DdeItem {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/DdeItem[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
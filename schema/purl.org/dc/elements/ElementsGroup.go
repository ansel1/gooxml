// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package elements

import (
	"encoding/xml"
	"fmt"
	"log"
)

type ElementsGroup struct {
	Choice []*ElementsGroupChoice
}

func NewElementsGroup() *ElementsGroup {
	ret := &ElementsGroup{}
	return ret
}

func (m *ElementsGroup) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Choice != nil {
		for _, c := range m.Choice {
			c.MarshalXML(e, start)
		}
	}
	return nil
}

func (m *ElementsGroup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lElementsGroup:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "any":
				tmp := NewElementsGroupChoice()
				if err := d.DecodeElement(&tmp.Any, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			default:
				log.Printf("skipping unsupported element on ElementsGroup %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lElementsGroup
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the ElementsGroup and its children
func (m *ElementsGroup) Validate() error {
	return m.ValidateWithPath("ElementsGroup")
}

// ValidateWithPath validates the ElementsGroup and its children, prefixing error messages with path
func (m *ElementsGroup) ValidateWithPath(path string) error {
	for i, v := range m.Choice {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Choice[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
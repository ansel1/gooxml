// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package terms

import (
	"encoding/xml"
	"fmt"
	"log"
)

type ElementOrRefinementContainer struct {
	Choice []*ElementsAndRefinementsGroupChoice
}

func NewElementOrRefinementContainer() *ElementOrRefinementContainer {
	ret := &ElementOrRefinementContainer{}
	return ret
}

func (m *ElementOrRefinementContainer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "elementOrRefinementContainer"
	e.EncodeToken(start)
	if m.Choice != nil {
		for _, c := range m.Choice {
			c.MarshalXML(e, start)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *ElementOrRefinementContainer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lElementOrRefinementContainer:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "any":
				tmp := NewElementsAndRefinementsGroupChoice()
				if err := d.DecodeElement(&tmp.Any, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			default:
				log.Printf("skipping unsupported element on ElementOrRefinementContainer %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lElementOrRefinementContainer
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the ElementOrRefinementContainer and its children
func (m *ElementOrRefinementContainer) Validate() error {
	return m.ValidateWithPath("ElementOrRefinementContainer")
}

// ValidateWithPath validates the ElementOrRefinementContainer and its children, prefixing error messages with path
func (m *ElementOrRefinementContainer) ValidateWithPath(path string) error {
	for i, v := range m.Choice {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Choice[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package relationships

import (
	"encoding/xml"
	"fmt"
)

type Relationship struct {
	CT_Relationship
}

func NewRelationship() *Relationship {
	ret := &Relationship{}
	ret.CT_Relationship = *NewCT_Relationship()
	return ret
}

func (m *Relationship) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_Relationship.MarshalXML(e, start)
}

func (m *Relationship) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Relationship = *NewCT_Relationship()
	for _, attr := range start.Attr {
		if attr.Name.Local == "TargetMode" {
			m.TargetModeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "Target" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TargetAttr = parsed
		}
		if attr.Name.Local == "Type" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TypeAttr = parsed
		}
		if attr.Name.Local == "Id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing Relationship: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Relationship and its children
func (m *Relationship) Validate() error {
	return m.ValidateWithPath("Relationship")
}

// ValidateWithPath validates the Relationship and its children, prefixing error messages with path
func (m *Relationship) ValidateWithPath(path string) error {
	if err := m.CT_Relationship.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}

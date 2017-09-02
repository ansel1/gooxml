// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package content_types

import (
	"encoding/xml"
	"fmt"
)

type Override struct {
	CT_Override
}

func NewOverride() *Override {
	ret := &Override{}
	ret.CT_Override = *NewCT_Override()
	return ret
}

func (m *Override) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_Override.MarshalXML(e, start)
}

func (m *Override) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Override = *NewCT_Override()
	for _, attr := range start.Attr {
		if attr.Name.Local == "ContentType" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ContentTypeAttr = parsed
		}
		if attr.Name.Local == "PartName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PartNameAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing Override: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Override and its children
func (m *Override) Validate() error {
	return m.ValidateWithPath("Override")
}

// ValidateWithPath validates the Override and its children, prefixing error messages with path
func (m *Override) ValidateWithPath(path string) error {
	if err := m.CT_Override.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}

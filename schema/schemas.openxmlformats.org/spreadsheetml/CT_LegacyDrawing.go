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
)

type CT_LegacyDrawing struct {
	IdAttr string
}

func NewCT_LegacyDrawing() *CT_LegacyDrawing {
	ret := &CT_LegacyDrawing{}
	return ret
}

func (m *CT_LegacyDrawing) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LegacyDrawing) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
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
			return fmt.Errorf("parsing CT_LegacyDrawing: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_LegacyDrawing and its children
func (m *CT_LegacyDrawing) Validate() error {
	return m.ValidateWithPath("CT_LegacyDrawing")
}

// ValidateWithPath validates the CT_LegacyDrawing and its children, prefixing error messages with path
func (m *CT_LegacyDrawing) ValidateWithPath(path string) error {
	return nil
}
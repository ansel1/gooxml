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
)

type CT_Charset struct {
	// Value
	ValAttr *string
	// IANA Name of Character Set
	CharacterSetAttr *string
}

func NewCT_Charset() *CT_Charset {
	ret := &CT_Charset{}
	return ret
}

func (m *CT_Charset) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
			Value: fmt.Sprintf("%v", *m.ValAttr)})
	}
	if m.CharacterSetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:characterSet"},
			Value: fmt.Sprintf("%v", *m.CharacterSetAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Charset) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
		}
		if attr.Name.Local == "characterSet" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CharacterSetAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Charset: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Charset and its children
func (m *CT_Charset) Validate() error {
	return m.ValidateWithPath("CT_Charset")
}

// ValidateWithPath validates the CT_Charset and its children, prefixing error messages with path
func (m *CT_Charset) ValidateWithPath(path string) error {
	return nil
}
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

type CT_Text struct {
	SpaceAttr *string
	Content   string
}

func NewCT_Text() *CT_Text {
	ret := &CT_Text{}
	return ret
}

func (m *CT_Text) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.SpaceAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xml:space"},
			Value: fmt.Sprintf("%v", *m.SpaceAttr)})
	}
	e.EncodeElement(m.Content, start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Text) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "space" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SpaceAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Text: %s", err)
		}
		if cd, ok := tok.(xml.CharData); ok {
			m.Content = string(cd)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Text and its children
func (m *CT_Text) Validate() error {
	return m.ValidateWithPath("CT_Text")
}

// ValidateWithPath validates the CT_Text and its children, prefixing error messages with path
func (m *CT_Text) ValidateWithPath(path string) error {
	return nil
}

// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package docPropsVTypes

import (
	"encoding/xml"
	"fmt"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_Vstream struct {
	VersionAttr string
	Content     string
}

func NewCT_Vstream() *CT_Vstream {
	ret := &CT_Vstream{}
	ret.VersionAttr = "{00000000-0000-0000-0000-000000000000}"
	return ret
}

func (m *CT_Vstream) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "version"},
		Value: fmt.Sprintf("%v", m.VersionAttr)})
	e.EncodeElement(m.Content, start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Vstream) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.VersionAttr = "{00000000-0000-0000-0000-000000000000}"
	for _, attr := range start.Attr {
		if attr.Name.Local == "version" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.VersionAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Vstream: %s", err)
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

// Validate validates the CT_Vstream and its children
func (m *CT_Vstream) Validate() error {
	return m.ValidateWithPath("CT_Vstream")
}

// ValidateWithPath validates the CT_Vstream and its children, prefixing error messages with path
func (m *CT_Vstream) ValidateWithPath(path string) error {
	if !sharedTypes.ST_GuidPatternRe.MatchString(m.VersionAttr) {
		return fmt.Errorf(`%s/m.VersionAttr must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, m.VersionAttr)
	}
	return nil
}

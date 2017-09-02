// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
)

type CT_LineJoinMiterProperties struct {
	LimAttr *ST_PositivePercentage
}

func NewCT_LineJoinMiterProperties() *CT_LineJoinMiterProperties {
	ret := &CT_LineJoinMiterProperties{}
	return ret
}

func (m *CT_LineJoinMiterProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.LimAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lim"},
			Value: fmt.Sprintf("%v", *m.LimAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LineJoinMiterProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "lim" {
			parsed, err := ParseUnionST_PositivePercentage(attr.Value)
			if err != nil {
				return err
			}
			m.LimAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_LineJoinMiterProperties: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_LineJoinMiterProperties and its children
func (m *CT_LineJoinMiterProperties) Validate() error {
	return m.ValidateWithPath("CT_LineJoinMiterProperties")
}

// ValidateWithPath validates the CT_LineJoinMiterProperties and its children, prefixing error messages with path
func (m *CT_LineJoinMiterProperties) ValidateWithPath(path string) error {
	if m.LimAttr != nil {
		if err := m.LimAttr.ValidateWithPath(path + "/LimAttr"); err != nil {
			return err
		}
	}
	return nil
}

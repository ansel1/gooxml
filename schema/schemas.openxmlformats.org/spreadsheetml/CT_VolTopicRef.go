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
	"strconv"
)

type CT_VolTopicRef struct {
	// Reference
	RAttr string
	// Sheet Id
	SAttr uint32
}

func NewCT_VolTopicRef() *CT_VolTopicRef {
	ret := &CT_VolTopicRef{}
	return ret
}

func (m *CT_VolTopicRef) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"},
		Value: fmt.Sprintf("%v", m.RAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "s"},
		Value: fmt.Sprintf("%v", m.SAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_VolTopicRef) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "r" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RAttr = parsed
		}
		if attr.Name.Local == "s" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.SAttr = uint32(parsed)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_VolTopicRef: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_VolTopicRef and its children
func (m *CT_VolTopicRef) Validate() error {
	return m.ValidateWithPath("CT_VolTopicRef")
}

// ValidateWithPath validates the CT_VolTopicRef and its children, prefixing error messages with path
func (m *CT_VolTopicRef) ValidateWithPath(path string) error {
	return nil
}

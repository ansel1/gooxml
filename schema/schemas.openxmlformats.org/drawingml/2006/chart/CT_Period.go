// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_Period struct {
	ValAttr *uint32
}

func NewCT_Period() *CT_Period {
	ret := &CT_Period{}
	return ret
}

func (m *CT_Period) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
			Value: fmt.Sprintf("%v", *m.ValAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Period) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ValAttr = &pt
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Period: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Period and its children
func (m *CT_Period) Validate() error {
	return m.ValidateWithPath("CT_Period")
}

// ValidateWithPath validates the CT_Period and its children, prefixing error messages with path
func (m *CT_Period) ValidateWithPath(path string) error {
	if m.ValAttr != nil {
		if *m.ValAttr < 2 {
			return fmt.Errorf("%s/m.ValAttr must be >= 2 (have %v)", path, *m.ValAttr)
		}
	}
	return nil
}

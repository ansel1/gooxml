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

type CT_LogBase struct {
	ValAttr float64
}

func NewCT_LogBase() *CT_LogBase {
	ret := &CT_LogBase{}
	ret.ValAttr = 2
	return ret
}

func (m *CT_LogBase) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LogBase) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = 2
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.ValAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_LogBase: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_LogBase and its children
func (m *CT_LogBase) Validate() error {
	return m.ValidateWithPath("CT_LogBase")
}

// ValidateWithPath validates the CT_LogBase and its children, prefixing error messages with path
func (m *CT_LogBase) ValidateWithPath(path string) error {
	if m.ValAttr < 2 {
		return fmt.Errorf("%s/m.ValAttr must be >= 2 (have %v)", path, m.ValAttr)
	}
	if m.ValAttr > 1000 {
		return fmt.Errorf("%s/m.ValAttr must be <= 1000 (have %v)", path, m.ValAttr)
	}
	return nil
}

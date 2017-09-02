// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"fmt"
)

type CT_BreakBinSub struct {
	ValAttr ST_BreakBinSub
}

func NewCT_BreakBinSub() *CT_BreakBinSub {
	ret := &CT_BreakBinSub{}
	return ret
}

func (m *CT_BreakBinSub) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != ST_BreakBinSubUnset {
		attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "m:val"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BreakBinSub) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			m.ValAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_BreakBinSub: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_BreakBinSub and its children
func (m *CT_BreakBinSub) Validate() error {
	return m.ValidateWithPath("CT_BreakBinSub")
}

// ValidateWithPath validates the CT_BreakBinSub and its children, prefixing error messages with path
func (m *CT_BreakBinSub) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
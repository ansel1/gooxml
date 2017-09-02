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

type CT_SheetCalcPr struct {
	// Full Calculation On Load
	FullCalcOnLoadAttr *bool
}

func NewCT_SheetCalcPr() *CT_SheetCalcPr {
	ret := &CT_SheetCalcPr{}
	return ret
}

func (m *CT_SheetCalcPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.FullCalcOnLoadAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fullCalcOnLoad"},
			Value: fmt.Sprintf("%v", *m.FullCalcOnLoadAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SheetCalcPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "fullCalcOnLoad" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FullCalcOnLoadAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SheetCalcPr: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_SheetCalcPr and its children
func (m *CT_SheetCalcPr) Validate() error {
	return m.ValidateWithPath("CT_SheetCalcPr")
}

// ValidateWithPath validates the CT_SheetCalcPr and its children, prefixing error messages with path
func (m *CT_SheetCalcPr) ValidateWithPath(path string) error {
	return nil
}

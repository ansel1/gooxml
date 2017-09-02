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

type CT_TableStyleInfo struct {
	// Style Name
	NameAttr *string
	// Show First Column
	ShowFirstColumnAttr *bool
	// Show Last Column
	ShowLastColumnAttr *bool
	// Show Row Stripes
	ShowRowStripesAttr *bool
	// Show Column Stripes
	ShowColumnStripesAttr *bool
}

func NewCT_TableStyleInfo() *CT_TableStyleInfo {
	ret := &CT_TableStyleInfo{}
	return ret
}

func (m *CT_TableStyleInfo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	if m.ShowFirstColumnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showFirstColumn"},
			Value: fmt.Sprintf("%v", *m.ShowFirstColumnAttr)})
	}
	if m.ShowLastColumnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showLastColumn"},
			Value: fmt.Sprintf("%v", *m.ShowLastColumnAttr)})
	}
	if m.ShowRowStripesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showRowStripes"},
			Value: fmt.Sprintf("%v", *m.ShowRowStripesAttr)})
	}
	if m.ShowColumnStripesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showColumnStripes"},
			Value: fmt.Sprintf("%v", *m.ShowColumnStripesAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TableStyleInfo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
		}
		if attr.Name.Local == "showFirstColumn" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowFirstColumnAttr = &parsed
		}
		if attr.Name.Local == "showLastColumn" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowLastColumnAttr = &parsed
		}
		if attr.Name.Local == "showRowStripes" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowRowStripesAttr = &parsed
		}
		if attr.Name.Local == "showColumnStripes" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowColumnStripesAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TableStyleInfo: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TableStyleInfo and its children
func (m *CT_TableStyleInfo) Validate() error {
	return m.ValidateWithPath("CT_TableStyleInfo")
}

// ValidateWithPath validates the CT_TableStyleInfo and its children, prefixing error messages with path
func (m *CT_TableStyleInfo) ValidateWithPath(path string) error {
	return nil
}

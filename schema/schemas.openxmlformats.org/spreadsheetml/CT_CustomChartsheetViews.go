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
	"log"
)

type CT_CustomChartsheetViews struct {
	// Custom Chart Sheet View
	CustomSheetView []*CT_CustomChartsheetView
}

func NewCT_CustomChartsheetViews() *CT_CustomChartsheetViews {
	ret := &CT_CustomChartsheetViews{}
	return ret
}

func (m *CT_CustomChartsheetViews) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.CustomSheetView != nil {
		secustomSheetView := xml.StartElement{Name: xml.Name{Local: "x:customSheetView"}}
		e.EncodeElement(m.CustomSheetView, secustomSheetView)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CustomChartsheetViews) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CustomChartsheetViews:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "customSheetView":
				tmp := NewCT_CustomChartsheetView()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CustomSheetView = append(m.CustomSheetView, tmp)
			default:
				log.Printf("skipping unsupported element on CT_CustomChartsheetViews %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustomChartsheetViews
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CustomChartsheetViews and its children
func (m *CT_CustomChartsheetViews) Validate() error {
	return m.ValidateWithPath("CT_CustomChartsheetViews")
}

// ValidateWithPath validates the CT_CustomChartsheetViews and its children, prefixing error messages with path
func (m *CT_CustomChartsheetViews) ValidateWithPath(path string) error {
	for i, v := range m.CustomSheetView {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CustomSheetView[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
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

type CT_Col struct {
	// Minimum Column
	MinAttr uint32
	// Maximum Column
	MaxAttr uint32
	// Column Width
	WidthAttr *float64
	// Style
	StyleAttr *uint32
	// Hidden Columns
	HiddenAttr *bool
	// Best Fit Column Width
	BestFitAttr *bool
	// Custom Width
	CustomWidthAttr *bool
	// Show Phonetic Information
	PhoneticAttr *bool
	// Outline Level
	OutlineLevelAttr *uint8
	// Collapsed
	CollapsedAttr *bool
}

func NewCT_Col() *CT_Col {
	ret := &CT_Col{}
	return ret
}

func (m *CT_Col) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "min"},
		Value: fmt.Sprintf("%v", m.MinAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "max"},
		Value: fmt.Sprintf("%v", m.MaxAttr)})
	if m.WidthAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "width"},
			Value: fmt.Sprintf("%v", *m.WidthAttr)})
	}
	if m.StyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "style"},
			Value: fmt.Sprintf("%v", *m.StyleAttr)})
	}
	if m.HiddenAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hidden"},
			Value: fmt.Sprintf("%v", *m.HiddenAttr)})
	}
	if m.BestFitAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "bestFit"},
			Value: fmt.Sprintf("%v", *m.BestFitAttr)})
	}
	if m.CustomWidthAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "customWidth"},
			Value: fmt.Sprintf("%v", *m.CustomWidthAttr)})
	}
	if m.PhoneticAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "phonetic"},
			Value: fmt.Sprintf("%v", *m.PhoneticAttr)})
	}
	if m.OutlineLevelAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "outlineLevel"},
			Value: fmt.Sprintf("%v", *m.OutlineLevelAttr)})
	}
	if m.CollapsedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "collapsed"},
			Value: fmt.Sprintf("%v", *m.CollapsedAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Col) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "min" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.MinAttr = uint32(parsed)
		}
		if attr.Name.Local == "max" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.MaxAttr = uint32(parsed)
		}
		if attr.Name.Local == "width" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.WidthAttr = &parsed
		}
		if attr.Name.Local == "style" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.StyleAttr = &pt
		}
		if attr.Name.Local == "hidden" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HiddenAttr = &parsed
		}
		if attr.Name.Local == "bestFit" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BestFitAttr = &parsed
		}
		if attr.Name.Local == "customWidth" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CustomWidthAttr = &parsed
		}
		if attr.Name.Local == "phonetic" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PhoneticAttr = &parsed
		}
		if attr.Name.Local == "outlineLevel" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			pt := uint8(parsed)
			m.OutlineLevelAttr = &pt
		}
		if attr.Name.Local == "collapsed" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CollapsedAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Col: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Col and its children
func (m *CT_Col) Validate() error {
	return m.ValidateWithPath("CT_Col")
}

// ValidateWithPath validates the CT_Col and its children, prefixing error messages with path
func (m *CT_Col) ValidateWithPath(path string) error {
	return nil
}

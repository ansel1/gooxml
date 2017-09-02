// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_WritingStyle struct {
	// Writing Style Language
	LangAttr string
	// Grammatical Engine ID
	VendorIDAttr string
	// Grammatical Check Engine Version
	DllVersionAttr string
	// Natural Language Grammar Check
	NlCheckAttr *sharedTypes.ST_OnOff
	// Check Stylistic Rules With Grammar
	CheckStyleAttr sharedTypes.ST_OnOff
	// Application Name
	AppNameAttr string
}

func NewCT_WritingStyle() *CT_WritingStyle {
	ret := &CT_WritingStyle{}
	return ret
}

func (m *CT_WritingStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:lang"},
		Value: fmt.Sprintf("%v", m.LangAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:vendorID"},
		Value: fmt.Sprintf("%v", m.VendorIDAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:dllVersion"},
		Value: fmt.Sprintf("%v", m.DllVersionAttr)})
	if m.NlCheckAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:nlCheck"},
			Value: fmt.Sprintf("%v", *m.NlCheckAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:checkStyle"},
		Value: fmt.Sprintf("%v", m.CheckStyleAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:appName"},
		Value: fmt.Sprintf("%v", m.AppNameAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_WritingStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "lang" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LangAttr = parsed
		}
		if attr.Name.Local == "vendorID" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.VendorIDAttr = parsed
		}
		if attr.Name.Local == "dllVersion" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DllVersionAttr = parsed
		}
		if attr.Name.Local == "nlCheck" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.NlCheckAttr = &parsed
		}
		if attr.Name.Local == "checkStyle" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.CheckStyleAttr = parsed
		}
		if attr.Name.Local == "appName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AppNameAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_WritingStyle: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_WritingStyle and its children
func (m *CT_WritingStyle) Validate() error {
	return m.ValidateWithPath("CT_WritingStyle")
}

// ValidateWithPath validates the CT_WritingStyle and its children, prefixing error messages with path
func (m *CT_WritingStyle) ValidateWithPath(path string) error {
	if m.NlCheckAttr != nil {
		if err := m.NlCheckAttr.ValidateWithPath(path + "/NlCheckAttr"); err != nil {
			return err
		}
	}
	if err := m.CheckStyleAttr.ValidateWithPath(path + "/CheckStyleAttr"); err != nil {
		return err
	}
	return nil
}

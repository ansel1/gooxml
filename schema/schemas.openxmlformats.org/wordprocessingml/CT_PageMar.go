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

type CT_PageMar struct {
	// Top Margin Spacing
	TopAttr ST_SignedTwipsMeasure
	// Right Margin Spacing
	RightAttr sharedTypes.ST_TwipsMeasure
	// Page Bottom Spacing
	BottomAttr ST_SignedTwipsMeasure
	// Left Margin Spacing
	LeftAttr sharedTypes.ST_TwipsMeasure
	// Spacing to Top of Header
	HeaderAttr sharedTypes.ST_TwipsMeasure
	// Spacing to Bottom of Footer
	FooterAttr sharedTypes.ST_TwipsMeasure
	// Page Gutter Spacing
	GutterAttr sharedTypes.ST_TwipsMeasure
}

func NewCT_PageMar() *CT_PageMar {
	ret := &CT_PageMar{}
	return ret
}

func (m *CT_PageMar) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:top"},
		Value: fmt.Sprintf("%v", m.TopAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:right"},
		Value: fmt.Sprintf("%v", m.RightAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:bottom"},
		Value: fmt.Sprintf("%v", m.BottomAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:left"},
		Value: fmt.Sprintf("%v", m.LeftAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:header"},
		Value: fmt.Sprintf("%v", m.HeaderAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:footer"},
		Value: fmt.Sprintf("%v", m.FooterAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:gutter"},
		Value: fmt.Sprintf("%v", m.GutterAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PageMar) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "top" {
			parsed, err := ParseUnionST_SignedTwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.TopAttr = parsed
		}
		if attr.Name.Local == "right" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.RightAttr = parsed
		}
		if attr.Name.Local == "bottom" {
			parsed, err := ParseUnionST_SignedTwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.BottomAttr = parsed
		}
		if attr.Name.Local == "left" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.LeftAttr = parsed
		}
		if attr.Name.Local == "header" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.HeaderAttr = parsed
		}
		if attr.Name.Local == "footer" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.FooterAttr = parsed
		}
		if attr.Name.Local == "gutter" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.GutterAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PageMar: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_PageMar and its children
func (m *CT_PageMar) Validate() error {
	return m.ValidateWithPath("CT_PageMar")
}

// ValidateWithPath validates the CT_PageMar and its children, prefixing error messages with path
func (m *CT_PageMar) ValidateWithPath(path string) error {
	if err := m.TopAttr.ValidateWithPath(path + "/TopAttr"); err != nil {
		return err
	}
	if err := m.RightAttr.ValidateWithPath(path + "/RightAttr"); err != nil {
		return err
	}
	if err := m.BottomAttr.ValidateWithPath(path + "/BottomAttr"); err != nil {
		return err
	}
	if err := m.LeftAttr.ValidateWithPath(path + "/LeftAttr"); err != nil {
		return err
	}
	if err := m.HeaderAttr.ValidateWithPath(path + "/HeaderAttr"); err != nil {
		return err
	}
	if err := m.FooterAttr.ValidateWithPath(path + "/FooterAttr"); err != nil {
		return err
	}
	if err := m.GutterAttr.ValidateWithPath(path + "/GutterAttr"); err != nil {
		return err
	}
	return nil
}

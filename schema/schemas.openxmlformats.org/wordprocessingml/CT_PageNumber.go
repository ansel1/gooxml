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
	"strconv"
)

type CT_PageNumber struct {
	// Page Number Format
	FmtAttr ST_NumberFormat
	// Starting Page Number
	StartAttr *int64
	// Chapter Heading Style
	ChapStyleAttr *int64
	// Chapter Separator Character
	ChapSepAttr ST_ChapterSep
}

func NewCT_PageNumber() *CT_PageNumber {
	ret := &CT_PageNumber{}
	return ret
}

func (m *CT_PageNumber) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.FmtAttr != ST_NumberFormatUnset {
		attr, err := m.FmtAttr.MarshalXMLAttr(xml.Name{Local: "w:fmt"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.StartAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:start"},
			Value: fmt.Sprintf("%v", *m.StartAttr)})
	}
	if m.ChapStyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:chapStyle"},
			Value: fmt.Sprintf("%v", *m.ChapStyleAttr)})
	}
	if m.ChapSepAttr != ST_ChapterSepUnset {
		attr, err := m.ChapSepAttr.MarshalXMLAttr(xml.Name{Local: "w:chapSep"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PageNumber) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "fmt" {
			m.FmtAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "start" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.StartAttr = &parsed
		}
		if attr.Name.Local == "chapStyle" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.ChapStyleAttr = &parsed
		}
		if attr.Name.Local == "chapSep" {
			m.ChapSepAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PageNumber: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_PageNumber and its children
func (m *CT_PageNumber) Validate() error {
	return m.ValidateWithPath("CT_PageNumber")
}

// ValidateWithPath validates the CT_PageNumber and its children, prefixing error messages with path
func (m *CT_PageNumber) ValidateWithPath(path string) error {
	if err := m.FmtAttr.ValidateWithPath(path + "/FmtAttr"); err != nil {
		return err
	}
	if err := m.ChapSepAttr.ValidateWithPath(path + "/ChapSepAttr"); err != nil {
		return err
	}
	return nil
}
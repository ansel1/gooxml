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

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_FileVersion struct {
	// Application Name
	AppNameAttr *string
	// Last Edited Version
	LastEditedAttr *string
	// Lowest Edited Version
	LowestEditedAttr *string
	// Build Version
	RupBuildAttr *string
	// Code Name
	CodeNameAttr *string
}

func NewCT_FileVersion() *CT_FileVersion {
	ret := &CT_FileVersion{}
	return ret
}

func (m *CT_FileVersion) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.AppNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "appName"},
			Value: fmt.Sprintf("%v", *m.AppNameAttr)})
	}
	if m.LastEditedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lastEdited"},
			Value: fmt.Sprintf("%v", *m.LastEditedAttr)})
	}
	if m.LowestEditedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lowestEdited"},
			Value: fmt.Sprintf("%v", *m.LowestEditedAttr)})
	}
	if m.RupBuildAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rupBuild"},
			Value: fmt.Sprintf("%v", *m.RupBuildAttr)})
	}
	if m.CodeNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "codeName"},
			Value: fmt.Sprintf("%v", *m.CodeNameAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FileVersion) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "appName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AppNameAttr = &parsed
		}
		if attr.Name.Local == "lastEdited" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LastEditedAttr = &parsed
		}
		if attr.Name.Local == "lowestEdited" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LowestEditedAttr = &parsed
		}
		if attr.Name.Local == "rupBuild" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RupBuildAttr = &parsed
		}
		if attr.Name.Local == "codeName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CodeNameAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_FileVersion: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_FileVersion and its children
func (m *CT_FileVersion) Validate() error {
	return m.ValidateWithPath("CT_FileVersion")
}

// ValidateWithPath validates the CT_FileVersion and its children, prefixing error messages with path
func (m *CT_FileVersion) ValidateWithPath(path string) error {
	if m.CodeNameAttr != nil {
		if !sharedTypes.ST_GuidPatternRe.MatchString(*m.CodeNameAttr) {
			return fmt.Errorf(`%s/m.CodeNameAttr must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, *m.CodeNameAttr)
		}
	}
	return nil
}

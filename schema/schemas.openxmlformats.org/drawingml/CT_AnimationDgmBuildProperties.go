// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_AnimationDgmBuildProperties struct {
	BldAttr *ST_AnimationDgmBuildType
	RevAttr *bool
}

func NewCT_AnimationDgmBuildProperties() *CT_AnimationDgmBuildProperties {
	ret := &CT_AnimationDgmBuildProperties{}
	return ret
}

func (m *CT_AnimationDgmBuildProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.BldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "bld"},
			Value: fmt.Sprintf("%v", *m.BldAttr)})
	}
	if m.RevAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rev"},
			Value: fmt.Sprintf("%v", *m.RevAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AnimationDgmBuildProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "bld" {
			parsed, err := ParseUnionST_AnimationDgmBuildType(attr.Value)
			if err != nil {
				return err
			}
			m.BldAttr = &parsed
		}
		if attr.Name.Local == "rev" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RevAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_AnimationDgmBuildProperties: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_AnimationDgmBuildProperties and its children
func (m *CT_AnimationDgmBuildProperties) Validate() error {
	return m.ValidateWithPath("CT_AnimationDgmBuildProperties")
}

// ValidateWithPath validates the CT_AnimationDgmBuildProperties and its children, prefixing error messages with path
func (m *CT_AnimationDgmBuildProperties) ValidateWithPath(path string) error {
	if m.BldAttr != nil {
		if err := m.BldAttr.ValidateWithPath(path + "/BldAttr"); err != nil {
			return err
		}
	}
	return nil
}

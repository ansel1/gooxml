// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_OptionalBlackTransition struct {
	// Transition Through Black
	ThruBlkAttr *bool
}

func NewCT_OptionalBlackTransition() *CT_OptionalBlackTransition {
	ret := &CT_OptionalBlackTransition{}
	return ret
}

func (m *CT_OptionalBlackTransition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ThruBlkAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "thruBlk"},
			Value: fmt.Sprintf("%v", *m.ThruBlkAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OptionalBlackTransition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "thruBlk" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ThruBlkAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_OptionalBlackTransition: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_OptionalBlackTransition and its children
func (m *CT_OptionalBlackTransition) Validate() error {
	return m.ValidateWithPath("CT_OptionalBlackTransition")
}

// ValidateWithPath validates the CT_OptionalBlackTransition and its children, prefixing error messages with path
func (m *CT_OptionalBlackTransition) ValidateWithPath(path string) error {
	return nil
}
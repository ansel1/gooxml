// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package lockedCanvas

import (
	"encoding/xml"
	"fmt"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type LockedCanvas struct {
	drawingml.CT_GvmlGroupShape
}

func NewLockedCanvas() *LockedCanvas {
	ret := &LockedCanvas{}
	ret.CT_GvmlGroupShape = *drawingml.NewCT_GvmlGroupShape()
	return ret
}

func (m *LockedCanvas) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/lockedCanvas"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "lockedCanvas"
	return m.CT_GvmlGroupShape.MarshalXML(e, start)
}

func (m *LockedCanvas) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_GvmlGroupShape = *drawingml.NewCT_GvmlGroupShape()
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing LockedCanvas: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the LockedCanvas and its children
func (m *LockedCanvas) Validate() error {
	return m.ValidateWithPath("LockedCanvas")
}

// ValidateWithPath validates the LockedCanvas and its children, prefixing error messages with path
func (m *LockedCanvas) ValidateWithPath(path string) error {
	if err := m.CT_GvmlGroupShape.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
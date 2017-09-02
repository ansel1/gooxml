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

type CT_Pane struct {
	// Horizontal Split Position
	XSplitAttr *float64
	// Vertical Split Position
	YSplitAttr *float64
	// Top Left Visible Cell
	TopLeftCellAttr *string
	// Active Pane
	ActivePaneAttr ST_Pane
	// Split State
	StateAttr ST_PaneState
}

func NewCT_Pane() *CT_Pane {
	ret := &CT_Pane{}
	return ret
}

func (m *CT_Pane) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.XSplitAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xSplit"},
			Value: fmt.Sprintf("%v", *m.XSplitAttr)})
	}
	if m.YSplitAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ySplit"},
			Value: fmt.Sprintf("%v", *m.YSplitAttr)})
	}
	if m.TopLeftCellAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "topLeftCell"},
			Value: fmt.Sprintf("%v", *m.TopLeftCellAttr)})
	}
	if m.ActivePaneAttr != ST_PaneUnset {
		attr, err := m.ActivePaneAttr.MarshalXMLAttr(xml.Name{Local: "activePane"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.StateAttr != ST_PaneStateUnset {
		attr, err := m.StateAttr.MarshalXMLAttr(xml.Name{Local: "state"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Pane) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "xSplit" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.XSplitAttr = &parsed
		}
		if attr.Name.Local == "ySplit" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.YSplitAttr = &parsed
		}
		if attr.Name.Local == "topLeftCell" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TopLeftCellAttr = &parsed
		}
		if attr.Name.Local == "activePane" {
			m.ActivePaneAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "state" {
			m.StateAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Pane: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Pane and its children
func (m *CT_Pane) Validate() error {
	return m.ValidateWithPath("CT_Pane")
}

// ValidateWithPath validates the CT_Pane and its children, prefixing error messages with path
func (m *CT_Pane) ValidateWithPath(path string) error {
	if err := m.ActivePaneAttr.ValidateWithPath(path + "/ActivePaneAttr"); err != nil {
		return err
	}
	if err := m.StateAttr.ValidateWithPath(path + "/StateAttr"); err != nil {
		return err
	}
	return nil
}
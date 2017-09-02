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
)

type CT_PCDKPI struct {
	// KPI Unique Name
	UniqueNameAttr string
	// KPI Display Name
	CaptionAttr *string
	// KPI Display Folder
	DisplayFolderAttr *string
	// KPI Measure Group Name
	MeasureGroupAttr *string
	// Parent KPI
	ParentAttr *string
	// KPI Value Unique Name
	ValueAttr string
	// KPI Goal Unique Name
	GoalAttr *string
	// KPI Status Unique Name
	StatusAttr *string
	// KPI Trend Unique Name
	TrendAttr *string
	// KPI Weight Unique Name
	WeightAttr *string
	// Time Member KPI Unique Name
	TimeAttr *string
}

func NewCT_PCDKPI() *CT_PCDKPI {
	ret := &CT_PCDKPI{}
	return ret
}

func (m *CT_PCDKPI) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uniqueName"},
		Value: fmt.Sprintf("%v", m.UniqueNameAttr)})
	if m.CaptionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "caption"},
			Value: fmt.Sprintf("%v", *m.CaptionAttr)})
	}
	if m.DisplayFolderAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "displayFolder"},
			Value: fmt.Sprintf("%v", *m.DisplayFolderAttr)})
	}
	if m.MeasureGroupAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "measureGroup"},
			Value: fmt.Sprintf("%v", *m.MeasureGroupAttr)})
	}
	if m.ParentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "parent"},
			Value: fmt.Sprintf("%v", *m.ParentAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "value"},
		Value: fmt.Sprintf("%v", m.ValueAttr)})
	if m.GoalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "goal"},
			Value: fmt.Sprintf("%v", *m.GoalAttr)})
	}
	if m.StatusAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "status"},
			Value: fmt.Sprintf("%v", *m.StatusAttr)})
	}
	if m.TrendAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "trend"},
			Value: fmt.Sprintf("%v", *m.TrendAttr)})
	}
	if m.WeightAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "weight"},
			Value: fmt.Sprintf("%v", *m.WeightAttr)})
	}
	if m.TimeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "time"},
			Value: fmt.Sprintf("%v", *m.TimeAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PCDKPI) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "uniqueName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UniqueNameAttr = parsed
		}
		if attr.Name.Local == "caption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CaptionAttr = &parsed
		}
		if attr.Name.Local == "displayFolder" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DisplayFolderAttr = &parsed
		}
		if attr.Name.Local == "measureGroup" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MeasureGroupAttr = &parsed
		}
		if attr.Name.Local == "parent" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ParentAttr = &parsed
		}
		if attr.Name.Local == "value" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ValueAttr = parsed
		}
		if attr.Name.Local == "goal" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GoalAttr = &parsed
		}
		if attr.Name.Local == "status" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StatusAttr = &parsed
		}
		if attr.Name.Local == "trend" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TrendAttr = &parsed
		}
		if attr.Name.Local == "weight" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.WeightAttr = &parsed
		}
		if attr.Name.Local == "time" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TimeAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PCDKPI: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_PCDKPI and its children
func (m *CT_PCDKPI) Validate() error {
	return m.ValidateWithPath("CT_PCDKPI")
}

// ValidateWithPath validates the CT_PCDKPI and its children, prefixing error messages with path
func (m *CT_PCDKPI) ValidateWithPath(path string) error {
	return nil
}
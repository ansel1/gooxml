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

type CT_FileRecoveryPr struct {
	// Auto Recover
	AutoRecoverAttr *bool
	// Crash Save
	CrashSaveAttr *bool
	// Data Extract Load
	DataExtractLoadAttr *bool
	// Repair Load
	RepairLoadAttr *bool
}

func NewCT_FileRecoveryPr() *CT_FileRecoveryPr {
	ret := &CT_FileRecoveryPr{}
	return ret
}

func (m *CT_FileRecoveryPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.AutoRecoverAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoRecover"},
			Value: fmt.Sprintf("%v", *m.AutoRecoverAttr)})
	}
	if m.CrashSaveAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "crashSave"},
			Value: fmt.Sprintf("%v", *m.CrashSaveAttr)})
	}
	if m.DataExtractLoadAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dataExtractLoad"},
			Value: fmt.Sprintf("%v", *m.DataExtractLoadAttr)})
	}
	if m.RepairLoadAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "repairLoad"},
			Value: fmt.Sprintf("%v", *m.RepairLoadAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FileRecoveryPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "autoRecover" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoRecoverAttr = &parsed
		}
		if attr.Name.Local == "crashSave" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CrashSaveAttr = &parsed
		}
		if attr.Name.Local == "dataExtractLoad" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DataExtractLoadAttr = &parsed
		}
		if attr.Name.Local == "repairLoad" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RepairLoadAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_FileRecoveryPr: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_FileRecoveryPr and its children
func (m *CT_FileRecoveryPr) Validate() error {
	return m.ValidateWithPath("CT_FileRecoveryPr")
}

// ValidateWithPath validates the CT_FileRecoveryPr and its children, prefixing error messages with path
func (m *CT_FileRecoveryPr) ValidateWithPath(path string) error {
	return nil
}
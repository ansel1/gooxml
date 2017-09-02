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

type CT_WorkbookProtection struct {
	// Legacy Workbook Password
	WorkbookPasswordAttr *string
	// Workbook Password Character Set
	WorkbookPasswordCharacterSetAttr *string
	// Legacy Revisions Password
	RevisionsPasswordAttr *string
	// Revisions Password Character Set
	RevisionsPasswordCharacterSetAttr *string
	// Lock Structure
	LockStructureAttr *bool
	// Lock Windows
	LockWindowsAttr *bool
	// Lock Revisions
	LockRevisionAttr *bool
	// Cryptographic Algorithm Name
	RevisionsAlgorithmNameAttr *string
	// Password Hash Value
	RevisionsHashValueAttr *string
	// Salt Value for Password Verifier
	RevisionsSaltValueAttr *string
	// Iterations to Run Hashing Algorithm
	RevisionsSpinCountAttr *uint32
	// Cryptographic Algorithm Name
	WorkbookAlgorithmNameAttr *string
	// Password Hash Value
	WorkbookHashValueAttr *string
	// Salt Value for Password Verifier
	WorkbookSaltValueAttr *string
	// Iterations to Run Hashing Algorithm
	WorkbookSpinCountAttr *uint32
}

func NewCT_WorkbookProtection() *CT_WorkbookProtection {
	ret := &CT_WorkbookProtection{}
	return ret
}

func (m *CT_WorkbookProtection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.WorkbookPasswordAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "workbookPassword"},
			Value: fmt.Sprintf("%v", *m.WorkbookPasswordAttr)})
	}
	if m.WorkbookPasswordCharacterSetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "workbookPasswordCharacterSet"},
			Value: fmt.Sprintf("%v", *m.WorkbookPasswordCharacterSetAttr)})
	}
	if m.RevisionsPasswordAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "revisionsPassword"},
			Value: fmt.Sprintf("%v", *m.RevisionsPasswordAttr)})
	}
	if m.RevisionsPasswordCharacterSetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "revisionsPasswordCharacterSet"},
			Value: fmt.Sprintf("%v", *m.RevisionsPasswordCharacterSetAttr)})
	}
	if m.LockStructureAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lockStructure"},
			Value: fmt.Sprintf("%v", *m.LockStructureAttr)})
	}
	if m.LockWindowsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lockWindows"},
			Value: fmt.Sprintf("%v", *m.LockWindowsAttr)})
	}
	if m.LockRevisionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lockRevision"},
			Value: fmt.Sprintf("%v", *m.LockRevisionAttr)})
	}
	if m.RevisionsAlgorithmNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "revisionsAlgorithmName"},
			Value: fmt.Sprintf("%v", *m.RevisionsAlgorithmNameAttr)})
	}
	if m.RevisionsHashValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "revisionsHashValue"},
			Value: fmt.Sprintf("%v", *m.RevisionsHashValueAttr)})
	}
	if m.RevisionsSaltValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "revisionsSaltValue"},
			Value: fmt.Sprintf("%v", *m.RevisionsSaltValueAttr)})
	}
	if m.RevisionsSpinCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "revisionsSpinCount"},
			Value: fmt.Sprintf("%v", *m.RevisionsSpinCountAttr)})
	}
	if m.WorkbookAlgorithmNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "workbookAlgorithmName"},
			Value: fmt.Sprintf("%v", *m.WorkbookAlgorithmNameAttr)})
	}
	if m.WorkbookHashValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "workbookHashValue"},
			Value: fmt.Sprintf("%v", *m.WorkbookHashValueAttr)})
	}
	if m.WorkbookSaltValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "workbookSaltValue"},
			Value: fmt.Sprintf("%v", *m.WorkbookSaltValueAttr)})
	}
	if m.WorkbookSpinCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "workbookSpinCount"},
			Value: fmt.Sprintf("%v", *m.WorkbookSpinCountAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_WorkbookProtection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "workbookPassword" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.WorkbookPasswordAttr = &parsed
		}
		if attr.Name.Local == "workbookPasswordCharacterSet" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.WorkbookPasswordCharacterSetAttr = &parsed
		}
		if attr.Name.Local == "revisionsPassword" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RevisionsPasswordAttr = &parsed
		}
		if attr.Name.Local == "revisionsPasswordCharacterSet" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RevisionsPasswordCharacterSetAttr = &parsed
		}
		if attr.Name.Local == "lockStructure" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LockStructureAttr = &parsed
		}
		if attr.Name.Local == "lockWindows" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LockWindowsAttr = &parsed
		}
		if attr.Name.Local == "lockRevision" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LockRevisionAttr = &parsed
		}
		if attr.Name.Local == "revisionsAlgorithmName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RevisionsAlgorithmNameAttr = &parsed
		}
		if attr.Name.Local == "revisionsHashValue" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RevisionsHashValueAttr = &parsed
		}
		if attr.Name.Local == "revisionsSaltValue" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RevisionsSaltValueAttr = &parsed
		}
		if attr.Name.Local == "revisionsSpinCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RevisionsSpinCountAttr = &pt
		}
		if attr.Name.Local == "workbookAlgorithmName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.WorkbookAlgorithmNameAttr = &parsed
		}
		if attr.Name.Local == "workbookHashValue" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.WorkbookHashValueAttr = &parsed
		}
		if attr.Name.Local == "workbookSaltValue" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.WorkbookSaltValueAttr = &parsed
		}
		if attr.Name.Local == "workbookSpinCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.WorkbookSpinCountAttr = &pt
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_WorkbookProtection: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_WorkbookProtection and its children
func (m *CT_WorkbookProtection) Validate() error {
	return m.ValidateWithPath("CT_WorkbookProtection")
}

// ValidateWithPath validates the CT_WorkbookProtection and its children, prefixing error messages with path
func (m *CT_WorkbookProtection) ValidateWithPath(path string) error {
	return nil
}
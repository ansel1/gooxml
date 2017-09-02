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

type CT_MetadataType struct {
	// Metadata Type Name
	NameAttr string
	// Minimum Supported Version
	MinSupportedVersionAttr uint32
	// Metadata Ghost Row
	GhostRowAttr *bool
	// Metadata Ghost Column
	GhostColAttr *bool
	// Metadata Edit
	EditAttr *bool
	// Metadata Cell Value Delete
	DeleteAttr *bool
	// Metadata Copy
	CopyAttr *bool
	// Metadata Paste All
	PasteAllAttr *bool
	// Metadata Paste Formulas
	PasteFormulasAttr *bool
	// Metadata Paste Special Values
	PasteValuesAttr *bool
	// Metadata Paste Formats
	PasteFormatsAttr *bool
	// Metadata Paste Comments
	PasteCommentsAttr *bool
	// Metadata Paste Data Validation
	PasteDataValidationAttr *bool
	// Metadata Paste Borders
	PasteBordersAttr *bool
	// Metadata Paste Column Widths
	PasteColWidthsAttr *bool
	// Metadata Paste Number Formats
	PasteNumberFormatsAttr *bool
	// Metadata Merge
	MergeAttr *bool
	// Meatadata Split First
	SplitFirstAttr *bool
	// Metadata Split All
	SplitAllAttr *bool
	// Metadata Insert Delete
	RowColShiftAttr *bool
	// Metadata Clear All
	ClearAllAttr *bool
	// Metadata Clear Formats
	ClearFormatsAttr *bool
	// Metadata Clear Contents
	ClearContentsAttr *bool
	// Metadata Clear Comments
	ClearCommentsAttr *bool
	// Metadata Formula Assignment
	AssignAttr *bool
	// Metadata Coercion
	CoerceAttr *bool
	// Adjust Metadata
	AdjustAttr *bool
	// Cell Metadata
	CellMetaAttr *bool
}

func NewCT_MetadataType() *CT_MetadataType {
	ret := &CT_MetadataType{}
	return ret
}

func (m *CT_MetadataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minSupportedVersion"},
		Value: fmt.Sprintf("%v", m.MinSupportedVersionAttr)})
	if m.GhostRowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ghostRow"},
			Value: fmt.Sprintf("%v", *m.GhostRowAttr)})
	}
	if m.GhostColAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ghostCol"},
			Value: fmt.Sprintf("%v", *m.GhostColAttr)})
	}
	if m.EditAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "edit"},
			Value: fmt.Sprintf("%v", *m.EditAttr)})
	}
	if m.DeleteAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "delete"},
			Value: fmt.Sprintf("%v", *m.DeleteAttr)})
	}
	if m.CopyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "copy"},
			Value: fmt.Sprintf("%v", *m.CopyAttr)})
	}
	if m.PasteAllAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pasteAll"},
			Value: fmt.Sprintf("%v", *m.PasteAllAttr)})
	}
	if m.PasteFormulasAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pasteFormulas"},
			Value: fmt.Sprintf("%v", *m.PasteFormulasAttr)})
	}
	if m.PasteValuesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pasteValues"},
			Value: fmt.Sprintf("%v", *m.PasteValuesAttr)})
	}
	if m.PasteFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pasteFormats"},
			Value: fmt.Sprintf("%v", *m.PasteFormatsAttr)})
	}
	if m.PasteCommentsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pasteComments"},
			Value: fmt.Sprintf("%v", *m.PasteCommentsAttr)})
	}
	if m.PasteDataValidationAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pasteDataValidation"},
			Value: fmt.Sprintf("%v", *m.PasteDataValidationAttr)})
	}
	if m.PasteBordersAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pasteBorders"},
			Value: fmt.Sprintf("%v", *m.PasteBordersAttr)})
	}
	if m.PasteColWidthsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pasteColWidths"},
			Value: fmt.Sprintf("%v", *m.PasteColWidthsAttr)})
	}
	if m.PasteNumberFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pasteNumberFormats"},
			Value: fmt.Sprintf("%v", *m.PasteNumberFormatsAttr)})
	}
	if m.MergeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "merge"},
			Value: fmt.Sprintf("%v", *m.MergeAttr)})
	}
	if m.SplitFirstAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "splitFirst"},
			Value: fmt.Sprintf("%v", *m.SplitFirstAttr)})
	}
	if m.SplitAllAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "splitAll"},
			Value: fmt.Sprintf("%v", *m.SplitAllAttr)})
	}
	if m.RowColShiftAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rowColShift"},
			Value: fmt.Sprintf("%v", *m.RowColShiftAttr)})
	}
	if m.ClearAllAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "clearAll"},
			Value: fmt.Sprintf("%v", *m.ClearAllAttr)})
	}
	if m.ClearFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "clearFormats"},
			Value: fmt.Sprintf("%v", *m.ClearFormatsAttr)})
	}
	if m.ClearContentsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "clearContents"},
			Value: fmt.Sprintf("%v", *m.ClearContentsAttr)})
	}
	if m.ClearCommentsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "clearComments"},
			Value: fmt.Sprintf("%v", *m.ClearCommentsAttr)})
	}
	if m.AssignAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "assign"},
			Value: fmt.Sprintf("%v", *m.AssignAttr)})
	}
	if m.CoerceAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "coerce"},
			Value: fmt.Sprintf("%v", *m.CoerceAttr)})
	}
	if m.AdjustAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "adjust"},
			Value: fmt.Sprintf("%v", *m.AdjustAttr)})
	}
	if m.CellMetaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cellMeta"},
			Value: fmt.Sprintf("%v", *m.CellMetaAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MetadataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
		}
		if attr.Name.Local == "minSupportedVersion" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.MinSupportedVersionAttr = uint32(parsed)
		}
		if attr.Name.Local == "ghostRow" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.GhostRowAttr = &parsed
		}
		if attr.Name.Local == "ghostCol" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.GhostColAttr = &parsed
		}
		if attr.Name.Local == "edit" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EditAttr = &parsed
		}
		if attr.Name.Local == "delete" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DeleteAttr = &parsed
		}
		if attr.Name.Local == "copy" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CopyAttr = &parsed
		}
		if attr.Name.Local == "pasteAll" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PasteAllAttr = &parsed
		}
		if attr.Name.Local == "pasteFormulas" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PasteFormulasAttr = &parsed
		}
		if attr.Name.Local == "pasteValues" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PasteValuesAttr = &parsed
		}
		if attr.Name.Local == "pasteFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PasteFormatsAttr = &parsed
		}
		if attr.Name.Local == "pasteComments" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PasteCommentsAttr = &parsed
		}
		if attr.Name.Local == "pasteDataValidation" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PasteDataValidationAttr = &parsed
		}
		if attr.Name.Local == "pasteBorders" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PasteBordersAttr = &parsed
		}
		if attr.Name.Local == "pasteColWidths" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PasteColWidthsAttr = &parsed
		}
		if attr.Name.Local == "pasteNumberFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PasteNumberFormatsAttr = &parsed
		}
		if attr.Name.Local == "merge" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MergeAttr = &parsed
		}
		if attr.Name.Local == "splitFirst" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SplitFirstAttr = &parsed
		}
		if attr.Name.Local == "splitAll" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SplitAllAttr = &parsed
		}
		if attr.Name.Local == "rowColShift" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RowColShiftAttr = &parsed
		}
		if attr.Name.Local == "clearAll" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ClearAllAttr = &parsed
		}
		if attr.Name.Local == "clearFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ClearFormatsAttr = &parsed
		}
		if attr.Name.Local == "clearContents" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ClearContentsAttr = &parsed
		}
		if attr.Name.Local == "clearComments" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ClearCommentsAttr = &parsed
		}
		if attr.Name.Local == "assign" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AssignAttr = &parsed
		}
		if attr.Name.Local == "coerce" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CoerceAttr = &parsed
		}
		if attr.Name.Local == "adjust" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AdjustAttr = &parsed
		}
		if attr.Name.Local == "cellMeta" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CellMetaAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_MetadataType: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_MetadataType and its children
func (m *CT_MetadataType) Validate() error {
	return m.ValidateWithPath("CT_MetadataType")
}

// ValidateWithPath validates the CT_MetadataType and its children, prefixing error messages with path
func (m *CT_MetadataType) ValidateWithPath(path string) error {
	return nil
}

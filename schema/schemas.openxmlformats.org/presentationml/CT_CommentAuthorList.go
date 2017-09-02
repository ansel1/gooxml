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
	"log"
)

type CT_CommentAuthorList struct {
	// Comment Author
	CmAuthor []*CT_CommentAuthor
}

func NewCT_CommentAuthorList() *CT_CommentAuthorList {
	ret := &CT_CommentAuthorList{}
	return ret
}

func (m *CT_CommentAuthorList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.CmAuthor != nil {
		secmAuthor := xml.StartElement{Name: xml.Name{Local: "p:cmAuthor"}}
		e.EncodeElement(m.CmAuthor, secmAuthor)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CommentAuthorList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CommentAuthorList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cmAuthor":
				tmp := NewCT_CommentAuthor()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CmAuthor = append(m.CmAuthor, tmp)
			default:
				log.Printf("skipping unsupported element on CT_CommentAuthorList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CommentAuthorList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CommentAuthorList and its children
func (m *CT_CommentAuthorList) Validate() error {
	return m.ValidateWithPath("CT_CommentAuthorList")
}

// ValidateWithPath validates the CT_CommentAuthorList and its children, prefixing error messages with path
func (m *CT_CommentAuthorList) ValidateWithPath(path string) error {
	for i, v := range m.CmAuthor {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CmAuthor[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
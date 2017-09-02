// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"fmt"
)

// ST_HoleSize is a union type
type ST_HoleSize struct {
	ST_HoleSizePercent *string
	ST_HoleSizeUByte   *uint8
}

func (m *ST_HoleSize) Validate() error {
	return m.ValidateWithPath("")
}

func (m *ST_HoleSize) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_HoleSizePercent != nil {
		mems = append(mems, "ST_HoleSizePercent")
	}
	if m.ST_HoleSizeUByte != nil {
		mems = append(mems, "ST_HoleSizeUByte")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_HoleSize) String() string {
	if m.ST_HoleSizePercent != nil {
		return fmt.Sprintf("%v", *m.ST_HoleSizePercent)
	}
	if m.ST_HoleSizeUByte != nil {
		return fmt.Sprintf("%v", *m.ST_HoleSizeUByte)
	}
	return ""
}
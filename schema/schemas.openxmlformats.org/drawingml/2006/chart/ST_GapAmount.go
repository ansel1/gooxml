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

// ST_GapAmount is a union type
type ST_GapAmount struct {
	ST_GapAmountPercent *string
	ST_GapAmountUShort  *uint16
}

func (m *ST_GapAmount) Validate() error {
	return m.ValidateWithPath("")
}

func (m *ST_GapAmount) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_GapAmountPercent != nil {
		mems = append(mems, "ST_GapAmountPercent")
	}
	if m.ST_GapAmountUShort != nil {
		mems = append(mems, "ST_GapAmountUShort")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_GapAmount) String() string {
	if m.ST_GapAmountPercent != nil {
		return fmt.Sprintf("%v", *m.ST_GapAmountPercent)
	}
	if m.ST_GapAmountUShort != nil {
		return fmt.Sprintf("%v", *m.ST_GapAmountUShort)
	}
	return ""
}
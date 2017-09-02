// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/wordprocessingml"
)

func TestCT_CellMergeTrackChangeConstructor(t *testing.T) {
	v := wordprocessingml.NewCT_CellMergeTrackChange()
	if v == nil {
		t.Errorf("wordprocessingml.NewCT_CellMergeTrackChange must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wordprocessingml.CT_CellMergeTrackChange should validate: %s", err)
	}
}

func TestCT_CellMergeTrackChangeMarshalUnmarshal(t *testing.T) {
	v := wordprocessingml.NewCT_CellMergeTrackChange()
	buf, _ := xml.Marshal(v)
	v2 := wordprocessingml.NewCT_CellMergeTrackChange()
	xml.Unmarshal(buf, v2)
}

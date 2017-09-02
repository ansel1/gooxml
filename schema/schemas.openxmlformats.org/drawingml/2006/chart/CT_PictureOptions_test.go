// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/chart"
)

func TestCT_PictureOptionsConstructor(t *testing.T) {
	v := chart.NewCT_PictureOptions()
	if v == nil {
		t.Errorf("chart.NewCT_PictureOptions must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_PictureOptions should validate: %s", err)
	}
}

func TestCT_PictureOptionsMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_PictureOptions()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_PictureOptions()
	xml.Unmarshal(buf, v2)
}
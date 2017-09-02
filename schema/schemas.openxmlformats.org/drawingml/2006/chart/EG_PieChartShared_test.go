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

func TestEG_PieChartSharedConstructor(t *testing.T) {
	v := chart.NewEG_PieChartShared()
	if v == nil {
		t.Errorf("chart.NewEG_PieChartShared must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.EG_PieChartShared should validate: %s", err)
	}
}

func TestEG_PieChartSharedMarshalUnmarshal(t *testing.T) {
	v := chart.NewEG_PieChartShared()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewEG_PieChartShared()
	xml.Unmarshal(buf, v2)
}

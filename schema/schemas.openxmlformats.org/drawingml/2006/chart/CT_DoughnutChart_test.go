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

func TestCT_DoughnutChartConstructor(t *testing.T) {
	v := chart.NewCT_DoughnutChart()
	if v == nil {
		t.Errorf("chart.NewCT_DoughnutChart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_DoughnutChart should validate: %s", err)
	}
}

func TestCT_DoughnutChartMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_DoughnutChart()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_DoughnutChart()
	xml.Unmarshal(buf, v2)
}

// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"
)

func TestPivotCacheDefinitionConstructor(t *testing.T) {
	v := spreadsheetml.NewPivotCacheDefinition()
	if v == nil {
		t.Errorf("spreadsheetml.NewPivotCacheDefinition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetml.PivotCacheDefinition should validate: %s", err)
	}
}

func TestPivotCacheDefinitionMarshalUnmarshal(t *testing.T) {
	v := spreadsheetml.NewPivotCacheDefinition()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetml.NewPivotCacheDefinition()
	xml.Unmarshal(buf, v2)
}

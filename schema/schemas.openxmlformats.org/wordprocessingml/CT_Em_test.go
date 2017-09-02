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

func TestCT_EmConstructor(t *testing.T) {
	v := wordprocessingml.NewCT_Em()
	if v == nil {
		t.Errorf("wordprocessingml.NewCT_Em must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wordprocessingml.CT_Em should validate: %s", err)
	}
}

func TestCT_EmMarshalUnmarshal(t *testing.T) {
	v := wordprocessingml.NewCT_Em()
	buf, _ := xml.Marshal(v)
	v2 := wordprocessingml.NewCT_Em()
	xml.Unmarshal(buf, v2)
}

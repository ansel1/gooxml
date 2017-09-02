// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package content_types_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/package/2006/content_types"
)

func TestDefaultConstructor(t *testing.T) {
	v := content_types.NewDefault()
	if v == nil {
		t.Errorf("content_types.NewDefault must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed content_types.Default should validate: %s", err)
	}
}

func TestDefaultMarshalUnmarshal(t *testing.T) {
	v := content_types.NewDefault()
	buf, _ := xml.Marshal(v)
	v2 := content_types.NewDefault()
	xml.Unmarshal(buf, v2)
}

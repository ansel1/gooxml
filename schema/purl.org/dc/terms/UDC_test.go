// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package terms_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/purl.org/dc/terms"
)

func TestUDCConstructor(t *testing.T) {
	v := terms.NewUDC()
	if v == nil {
		t.Errorf("terms.NewUDC must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.UDC should validate: %s", err)
	}
}

func TestUDCMarshalUnmarshal(t *testing.T) {
	v := terms.NewUDC()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewUDC()
	xml.Unmarshal(buf, v2)
}

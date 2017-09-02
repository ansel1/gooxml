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

func TestRFC3066Constructor(t *testing.T) {
	v := terms.NewRFC3066()
	if v == nil {
		t.Errorf("terms.NewRFC3066 must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.RFC3066 should validate: %s", err)
	}
}

func TestRFC3066MarshalUnmarshal(t *testing.T) {
	v := terms.NewRFC3066()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewRFC3066()
	xml.Unmarshal(buf, v2)
}

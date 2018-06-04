// Copyright 2015 go-smpp authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package pdutlv

import (
	"encoding/json"
	"testing"
)

func TestMapSet(t *testing.T) {
	m := make(Map)
	test := []struct {
		k  Tag
		v  interface{}
		ok bool
	}{
		{DestAddrSubunit, nil, true},
		{DestAddrSubunit, "hello", true},
		{DestAddrSubunit, []byte("hello"), true},
		{DestBearerType, nil, true},
		{DestBearerType, uint8(1), true},
		{DestBearerType, int(1), true},
		{DestBearerType, t, false},
		{DestBearerType, String("hello"), true},
		{DestBearerType, CString("hello\x00"), true},
		{DestBearerType, CString("hello"), true},
		{DestBearerType, NewTLV(DestBearerType, []byte{0x03}), true},
	}
	for _, el := range test {
		if err := m.Set(el.k, el.v); el.ok && err != nil {
			t.Fatal(err)
		} else if !el.ok && err == nil {
			t.Fatalf("unexpected set of %q=%#v", el.k, el.v)
		}
	}
}

func TestTLVMarshalJSON(t *testing.T) {
	m := make(Map)
	tlvTypeA := Tag(1)
	tlvTypeB := Tag(2)
	dataA := []byte("tlvBodyA")
	dataB := []byte("tlvBodyB")
	tlvBodyA := NewTLV(tlvTypeA, dataA)
	tlvBodyB := NewTLV(tlvTypeB, dataB)

	m[tlvTypeA] = tlvBodyA
	m[tlvTypeB] = tlvBodyB
	bytes, err := json.Marshal(m)
	if err != nil {
		t.Fatal("error marshaling:", err)
	}

	other := make(Map)
	err = json.Unmarshal(bytes, &other)
	if err != nil {
		t.Fatal("error unmarshaling:", err)
	}

	for k, v := range m {
		if val, ok := other[k]; ok {
			valStr := string(val.Bytes())
			vStr := string(v.Bytes())
			if valStr != vStr {
				t.Fatalf("in key %v, expected to contain: %v, got %v instead", k, vStr, valStr)
			} else if val.Len() != v.Len() {
				t.Fatalf("in key %v, expected to contain len: %v, got %v instead", k, v.Len, val.Len)
			}
		} else {
			t.Fatalf("did not find key: %v", k)
		}
	}
}

func TestTLVUnmarshalJSON(t *testing.T) {
	jsonBytes := []byte(`{"Map" :{
		  "1": {
		    "tag": 1,
		    "len": 8,
		    "data": "dGx2Qm9keUE=",
		    "text": "tlvBodyA"
		  },
		  "2": {
		    "tag": 2,
		    "len": 8,
		    "data": "dGx2Qm9keUI=",
		    "text": "tlvBodyB"
		  }
		}}`)

	s := struct {
		Map Map
	}{
		Map: nil,
	}

	err := json.Unmarshal(jsonBytes, &s)
	if err != nil {
		t.Fatal("error unmarshaling:", err)
	}
	if _, ok := s.Map[Tag(1)]; !ok {
		t.Fatalf("did not find key: %v", 1)
	}
	if _, ok := s.Map[Tag(2)]; !ok {
		t.Fatalf("did not find key: %v", 2)
	}
}

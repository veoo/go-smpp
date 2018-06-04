// Copyright 2015 go-smpp authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package pdutlv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

// Map is a collection of PDU TLV field data indexed by tag.
type Map map[Tag]Body

// Set updates the PDU map with the given tag and value, and
// returns error if the value cannot be converted to type Data.
//
// This is a shortcut for m[t] = NewTLV(t, v) converting v properly.
func (m Map) Set(t Tag, v interface{}) error {
	switch v.(type) {
	case nil:
		m[t] = NewTLV(t, nil) // use default value
	case uint8:
		m[t] = NewTLV(t, []byte{v.(uint8)})
	case int:
		m[t] = NewTLV(t, []byte{uint8(v.(int))})
	case string:
		m[t] = NewTLV(t, []byte(v.(string)))
	case String:
		m[t] = NewTLV(t, []byte(v.(String)))
	case CString:
		value := []byte(v.(CString))
		if len(value) == 0 || value[len(value)-1] != 0x00 {
			value = append(value, 0x00)
		}
		m[t] = NewTLV(t, value)
	case []byte:
		m[t] = NewTLV(t, []byte(v.([]byte)))
	case Body:
		m[t] = v.(Body)
	case MessageStateType:
		m[t] = NewTLV(MessageStateOption, []byte{uint8(v.(MessageStateType))})
	default:
		return fmt.Errorf("unsupported Tag-Length-Value field data: %#v", v)
	}
	return nil
}

func (m Map) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	length := len(m)
	count := 0
	for k, v := range m {
		jsonValue, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		buffer.WriteString(fmt.Sprintf("\"%d\":%s", k, string(jsonValue)))
		count++
		if count < length {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

func (m *Map) UnmarshalJSON(b []byte) error {
	if *m == nil {
		*m = Map{}
	}
	var tmp map[string]*Field
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}
	mtmp := map[Tag]Body{}
	for k, v := range tmp {
		numericKey, err := strconv.Atoi(k)
		if err != nil {
			return err
		}
		mtmp[Tag(numericKey)] = v
	}
	*m = mtmp
	return nil
}

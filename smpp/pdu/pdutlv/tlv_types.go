// Copyright 2015 go-smpp authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package pdutlv

import (
	"encoding/binary"
	"encoding/hex"
	"io"
)

// Fields is a map of tagged TLV fields
type Fields map[Tag]interface{}

// String is a text string that is not null-terminated.
type String string

// CString is a text string that is automatically null-terminated (e.g., final 00 byte at the end).
type CString string

// Tag is the tag of a Tag-Length-Value (TLV) field.
type Tag uint16

// Hex returns hexadecimal representation of tag
func (t Tag) Hex() string {
	bin := make([]byte, 2, 2)
	binary.BigEndian.PutUint16(bin, uint16(t))
	return hex.EncodeToString(bin)
}

// Common Tag-Length-Value (TLV) tags.
const (
	DestAddrSubunit          Tag = 0x0005
	DestNetworkType          Tag = 0x0006
	DestBearerType           Tag = 0x0007
	DestTelematicsID         Tag = 0x0008
	SourceAddrSubunit        Tag = 0x000D
	SourceNetworkType        Tag = 0x000E
	SourceBearerType         Tag = 0x000F
	SourceTelematicsID       Tag = 0x0010
	QosTimeToLive            Tag = 0x0017
	PayloadType              Tag = 0x0019
	AdditionalStatusInfoText Tag = 0x001D
	ReceiptedMessageID       Tag = 0x001E
	MsMsgWaitFacilities      Tag = 0x0030
	PrivacyIndicator         Tag = 0x0201
	SourceSubaddress         Tag = 0x0202
	DestSubaddress           Tag = 0x0203
	UserMessageReference     Tag = 0x0204
	UserResponseCode         Tag = 0x0205
	SourcePort               Tag = 0x020A
	DestinationPort          Tag = 0x020B
	SarMsgRefNum             Tag = 0x020C
	LanguageIndicator        Tag = 0x020D
	SarTotalSegments         Tag = 0x020E
	SarSegmentSeqnum         Tag = 0x020F
	CallbackNumPresInd       Tag = 0x0302
	CallbackNumAtag          Tag = 0x0303
	NumberOfMessages         Tag = 0x0304
	CallbackNum              Tag = 0x0381
	DpfResult                Tag = 0x0420
	SetDpf                   Tag = 0x0421
	MsAvailabilityStatus     Tag = 0x0422
	NetworkErrorCode         Tag = 0x0423
	MessagePayload           Tag = 0x0424
	DeliveryFailureReason    Tag = 0x0425
	MoreMessagesToSend       Tag = 0x0426
	MessageStateOption       Tag = 0x0427
	UssdServiceOp            Tag = 0x0501
	DisplayTime              Tag = 0x1201
	SmsSignal                Tag = 0x1203
	MsValidity               Tag = 0x1204
	AlertOnMessageDelivery   Tag = 0x130C
	ItsReplyType             Tag = 0x1380
	ItsSessionInfo           Tag = 0x1383
)

// Field is a PDU Tag-Length-Value (TLV) field
type Field struct {
	Tag  Tag
	Data []byte
}

// Len implements the Data interface.
func (t *Field) Len() int {
	return len(t.Bytes()) + 4
}

// Raw implements the Data interface.
func (t *Field) Raw() interface{} {
	return t.Bytes()
}

// String implements the Data interface.
func (t *Field) String() string {
	if l := len(t.Data); l > 0 && t.Data[l-1] == 0x00 {
		return string(t.Data[:l-1])
	}
	return string(t.Data)
}

// Bytes implements the Data interface.
func (t *Field) Bytes() []byte {
	return t.Data
}

// SerializeTo implements the Data interface.
func (t *Field) SerializeTo(w io.Writer) error {
	b := make([]byte, len(t.Data)+4)
	binary.BigEndian.PutUint16(b[0:2], uint16(t.Tag))
	binary.BigEndian.PutUint16(b[2:4], uint16(len(t.Data)))
	copy(b[4:], t.Data)

	_, err := w.Write(b)
	return err
}

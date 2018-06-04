// Copyright 2015 go-smpp authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package pdutlv

import (
	"io"
	"strconv"
)

// Body is an interface for manipulating binary PDU Tag-Length-Value field data.
type Body interface {
	Len() int
	Raw() interface{}
	String() string
	Bytes() []byte
	SerializeTo(w io.Writer) error
}

// NewTLV parses the given binary data and returns a Data object,
// or nil if the field Name is unknown.
func NewTLV(tag Tag, value []byte) Body {
	return &Field{Tag: tag, Data: value}
}

var tlvTypeMap = map[Tag]string{
	DestAddrSubunit:          "dest_addr_subunit",
	DestNetworkType:          "dest_network_type",
	DestBearerType:           "dest_bearer_type",
	DestTelematicsID:         "dest_telematics_id",
	SourceAddrSubunit:        "source_addr_subunit",
	SourceNetworkType:        "source_network_type",
	SourceBearerType:         "source_bearer_type",
	SourceTelematicsID:       "source_telematics_id",
	QosTimeToLive:            "qos_time_to_live",
	PayloadType:              "payload_type",
	AdditionalStatusInfoText: "additional_status_info_text",
	ReceiptedMessageID:       "receipted_message_id",
	MsMsgWaitFacilities:      "ms_msg_wait_facilities",
	PrivacyIndicator:         "privacy_indicator",
	SourceSubaddress:         "source_subaddress",
	DestSubaddress:           "dest_subaddress",
	UserMessageReference:     "user_message_reference",
	UserResponseCode:         "user_response_code",
	SourcePort:               "source_port",
	DestinationPort:          "destination_port",
	SarMsgRefNum:             "sar_msg_ref_num",
	LanguageIndicator:        "language_indicator",
	SarTotalSegments:         "sar_total_segments",
	SarSegmentSeqnum:         "sar_segment_seqnum",
	CallbackNumPresInd:       "callback_num_pres_ind",
	CallbackNumAtag:          "callback_num_atag",
	NumberOfMessages:         "number_of_messages",
	CallbackNum:              "callback_num",
	DpfResult:                "dpf_result",
	SetDpf:                   "set_dpf",
	MsAvailabilityStatus:     "ms_availability_status",
	NetworkErrorCode:         "network_error_code",
	MessagePayload:           "message_payload",
	DeliveryFailureReason:    "delivery_failure_reason",
	MoreMessagesToSend:       "more_messages_to_send",
	MessageStateOption:       "message_state_option",
	UssdServiceOp:            "ussd_service_op",
	DisplayTime:              "display_time",
	SmsSignal:                "sms_signal",
	MsValidity:               "ms_validity",
	AlertOnMessageDelivery:   "alert_on_message_delivery",
	ItsReplyType:             "its_reply_type",
	ItsSessionInfo:           "its_session_info",
}

func (t Tag) String() string {
	s := tlvTypeMap[t]
	if s == "" {
		s = strconv.Itoa(int(t))
	}
	return s
}

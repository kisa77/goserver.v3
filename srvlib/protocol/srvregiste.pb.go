// Code generated by protoc-gen-go.
// source: protocol/srvregiste.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type SSSrvRegiste struct {
	Id               *int32  `protobuf:"varint,1,req" json:"Id,omitempty"`
	Type             *int32  `protobuf:"varint,2,req" json:"Type,omitempty"`
	AreaId           *int32  `protobuf:"varint,3,req" json:"AreaId,omitempty"`
	Name             *string `protobuf:"bytes,4,req" json:"Name,omitempty"`
	Data             *string `protobuf:"bytes,5,opt" json:"Data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SSSrvRegiste) Reset()         { *m = SSSrvRegiste{} }
func (m *SSSrvRegiste) String() string { return proto.CompactTextString(m) }
func (*SSSrvRegiste) ProtoMessage()    {}

func (m *SSSrvRegiste) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *SSSrvRegiste) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *SSSrvRegiste) GetAreaId() int32 {
	if m != nil && m.AreaId != nil {
		return *m.AreaId
	}
	return 0
}

func (m *SSSrvRegiste) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *SSSrvRegiste) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

func init() {
}

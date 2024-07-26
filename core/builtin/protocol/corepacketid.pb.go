// Code generated by protoc-gen-go.
// source: protocol/corepacketid.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type CoreBuiltinPacketID int32

const (
	CoreBuiltinPacketID_PACKET_SS_TX_START  CoreBuiltinPacketID = -1000
	CoreBuiltinPacketID_PACKET_SS_TX_CMD    CoreBuiltinPacketID = -1001
	CoreBuiltinPacketID_PACKET_SS_TX_RESULT CoreBuiltinPacketID = -1002
	CoreBuiltinPacketID_PACKET_SS_SLICES    CoreBuiltinPacketID = -1003
	CoreBuiltinPacketID_PACKET_SS_AUTH      CoreBuiltinPacketID = -1004
	CoreBuiltinPacketID_PACKET_SS_KEEPALIVE CoreBuiltinPacketID = -1005
	CoreBuiltinPacketID_PACKET_SS_AUTH_ACK  CoreBuiltinPacketID = -1006
)

var CoreBuiltinPacketID_name = map[int32]string{
	-1000: "PACKET_SS_TX_START",
	-1001: "PACKET_SS_TX_CMD",
	-1002: "PACKET_SS_TX_RESULT",
	-1003: "PACKET_SS_SLICES",
	-1004: "PACKET_SS_AUTH",
	-1005: "PACKET_SS_KEEPALIVE",
	-1006: "PACKET_SS_AUTH_ACK",
}
var CoreBuiltinPacketID_value = map[string]int32{
	"PACKET_SS_TX_START":  -1000,
	"PACKET_SS_TX_CMD":    -1001,
	"PACKET_SS_TX_RESULT": -1002,
	"PACKET_SS_SLICES":    -1003,
	"PACKET_SS_AUTH":      -1004,
	"PACKET_SS_KEEPALIVE": -1005,
	"PACKET_SS_AUTH_ACK":  -1006,
}

func (x CoreBuiltinPacketID) Enum() *CoreBuiltinPacketID {
	p := new(CoreBuiltinPacketID)
	*p = x
	return p
}
func (x CoreBuiltinPacketID) String() string {
	return proto.EnumName(CoreBuiltinPacketID_name, int32(x))
}
func (x CoreBuiltinPacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *CoreBuiltinPacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CoreBuiltinPacketID_value, data, "CoreBuiltinPacketID")
	if err != nil {
		return err
	}
	*x = CoreBuiltinPacketID(value)
	return nil
}

func init() {
	proto.RegisterEnum("protocol.CoreBuiltinPacketID", CoreBuiltinPacketID_name, CoreBuiltinPacketID_value)
}

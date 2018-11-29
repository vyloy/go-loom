// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/karma/karma.proto

/*
Package karma is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/go-loom/builtin/types/karma/karma.proto

It has these top-level messages:
	KarmaInitRequest
	KarmaSources
	KarmaSourcesValidator
	KarmaNewOracleValidator
	KarmaUserTarget
	KarmaUserAmount
	KarmaSourceReward
	KarmaSource
	KarmaUpkeepParams
	KarmaAddressSource
	KarmaState
	KarmaStateUser
	KarmaStateKeyUser
	KarmaTotal
*/
package karma

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type KarmaSourceTarget int32

const (
	KarmaSourceTarget_CALL   KarmaSourceTarget = 0
	KarmaSourceTarget_DEPLOY KarmaSourceTarget = 1
	KarmaSourceTarget_ALL    KarmaSourceTarget = 2
)

var KarmaSourceTarget_name = map[int32]string{
	0: "CALL",
	1: "DEPLOY",
	2: "ALL",
}
var KarmaSourceTarget_value = map[string]int32{
	"CALL":   0,
	"DEPLOY": 1,
	"ALL":    2,
}

func (x KarmaSourceTarget) String() string {
	return proto.EnumName(KarmaSourceTarget_name, int32(x))
}
func (KarmaSourceTarget) EnumDescriptor() ([]byte, []int) { return fileDescriptorKarma, []int{0} }

type KarmaInitRequest struct {
	Oracle  *types.Address        `protobuf:"bytes,1,opt,name=oracle" json:"oracle,omitempty"`
	Sources []*KarmaSourceReward  `protobuf:"bytes,2,rep,name=sources" json:"sources,omitempty"`
	Users   []*KarmaAddressSource `protobuf:"bytes,3,rep,name=users" json:"users,omitempty"`
	Upkeep  *KarmaUpkeepParams    `protobuf:"bytes,4,opt,name=upkeep" json:"upkeep,omitempty"`
}

func (m *KarmaInitRequest) Reset()                    { *m = KarmaInitRequest{} }
func (m *KarmaInitRequest) String() string            { return proto.CompactTextString(m) }
func (*KarmaInitRequest) ProtoMessage()               {}
func (*KarmaInitRequest) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{0} }

func (m *KarmaInitRequest) GetOracle() *types.Address {
	if m != nil {
		return m.Oracle
	}
	return nil
}

func (m *KarmaInitRequest) GetSources() []*KarmaSourceReward {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *KarmaInitRequest) GetUsers() []*KarmaAddressSource {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *KarmaInitRequest) GetUpkeep() *KarmaUpkeepParams {
	if m != nil {
		return m.Upkeep
	}
	return nil
}

type KarmaSources struct {
	Sources []*KarmaSourceReward `protobuf:"bytes,1,rep,name=sources" json:"sources,omitempty"`
}

func (m *KarmaSources) Reset()                    { *m = KarmaSources{} }
func (m *KarmaSources) String() string            { return proto.CompactTextString(m) }
func (*KarmaSources) ProtoMessage()               {}
func (*KarmaSources) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{1} }

func (m *KarmaSources) GetSources() []*KarmaSourceReward {
	if m != nil {
		return m.Sources
	}
	return nil
}

type KarmaSourcesValidator struct {
	Sources []*KarmaSourceReward `protobuf:"bytes,1,rep,name=sources" json:"sources,omitempty"`
	Oracle  *types.Address       `protobuf:"bytes,2,opt,name=oracle" json:"oracle,omitempty"`
}

func (m *KarmaSourcesValidator) Reset()                    { *m = KarmaSourcesValidator{} }
func (m *KarmaSourcesValidator) String() string            { return proto.CompactTextString(m) }
func (*KarmaSourcesValidator) ProtoMessage()               {}
func (*KarmaSourcesValidator) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{2} }

func (m *KarmaSourcesValidator) GetSources() []*KarmaSourceReward {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *KarmaSourcesValidator) GetOracle() *types.Address {
	if m != nil {
		return m.Oracle
	}
	return nil
}

type KarmaNewOracleValidator struct {
	NewOracle *types.Address `protobuf:"bytes,1,opt,name=new_oracle,json=newOracle" json:"new_oracle,omitempty"`
	OldOracle *types.Address `protobuf:"bytes,2,opt,name=old_oracle,json=oldOracle" json:"old_oracle,omitempty"`
}

func (m *KarmaNewOracleValidator) Reset()                    { *m = KarmaNewOracleValidator{} }
func (m *KarmaNewOracleValidator) String() string            { return proto.CompactTextString(m) }
func (*KarmaNewOracleValidator) ProtoMessage()               {}
func (*KarmaNewOracleValidator) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{3} }

func (m *KarmaNewOracleValidator) GetNewOracle() *types.Address {
	if m != nil {
		return m.NewOracle
	}
	return nil
}

func (m *KarmaNewOracleValidator) GetOldOracle() *types.Address {
	if m != nil {
		return m.OldOracle
	}
	return nil
}

type KarmaUserTarget struct {
	User   *types.Address    `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Target KarmaSourceTarget `protobuf:"varint,2,opt,name=target,proto3,enum=karma.KarmaSourceTarget" json:"target,omitempty"`
}

func (m *KarmaUserTarget) Reset()                    { *m = KarmaUserTarget{} }
func (m *KarmaUserTarget) String() string            { return proto.CompactTextString(m) }
func (*KarmaUserTarget) ProtoMessage()               {}
func (*KarmaUserTarget) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{4} }

func (m *KarmaUserTarget) GetUser() *types.Address {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *KarmaUserTarget) GetTarget() KarmaSourceTarget {
	if m != nil {
		return m.Target
	}
	return KarmaSourceTarget_CALL
}

type KarmaUserAmount struct {
	User   *types.Address `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Amount *types.BigUInt `protobuf:"bytes,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *KarmaUserAmount) Reset()                    { *m = KarmaUserAmount{} }
func (m *KarmaUserAmount) String() string            { return proto.CompactTextString(m) }
func (*KarmaUserAmount) ProtoMessage()               {}
func (*KarmaUserAmount) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{5} }

func (m *KarmaUserAmount) GetUser() *types.Address {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *KarmaUserAmount) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type KarmaSourceReward struct {
	Name   string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Reward int64             `protobuf:"varint,2,opt,name=reward,proto3" json:"reward,omitempty"`
	Target KarmaSourceTarget `protobuf:"varint,3,opt,name=target,proto3,enum=karma.KarmaSourceTarget" json:"target,omitempty"`
}

func (m *KarmaSourceReward) Reset()                    { *m = KarmaSourceReward{} }
func (m *KarmaSourceReward) String() string            { return proto.CompactTextString(m) }
func (*KarmaSourceReward) ProtoMessage()               {}
func (*KarmaSourceReward) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{6} }

func (m *KarmaSourceReward) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KarmaSourceReward) GetReward() int64 {
	if m != nil {
		return m.Reward
	}
	return 0
}

func (m *KarmaSourceReward) GetTarget() KarmaSourceTarget {
	if m != nil {
		return m.Target
	}
	return KarmaSourceTarget_CALL
}

type KarmaSource struct {
	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Count int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *KarmaSource) Reset()                    { *m = KarmaSource{} }
func (m *KarmaSource) String() string            { return proto.CompactTextString(m) }
func (*KarmaSource) ProtoMessage()               {}
func (*KarmaSource) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{7} }

func (m *KarmaSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KarmaSource) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type KarmaUpkeepParams struct {
	Cost   int64  `protobuf:"varint,1,opt,name=cost,proto3" json:"cost,omitempty"`
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Period int64  `protobuf:"varint,3,opt,name=period,proto3" json:"period,omitempty"`
}

func (m *KarmaUpkeepParams) Reset()                    { *m = KarmaUpkeepParams{} }
func (m *KarmaUpkeepParams) String() string            { return proto.CompactTextString(m) }
func (*KarmaUpkeepParams) ProtoMessage()               {}
func (*KarmaUpkeepParams) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{8} }

func (m *KarmaUpkeepParams) GetCost() int64 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func (m *KarmaUpkeepParams) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *KarmaUpkeepParams) GetPeriod() int64 {
	if m != nil {
		return m.Period
	}
	return 0
}

type KarmaAddressSource struct {
	User    *types.Address `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Sources []*KarmaSource `protobuf:"bytes,2,rep,name=sources" json:"sources,omitempty"`
}

func (m *KarmaAddressSource) Reset()                    { *m = KarmaAddressSource{} }
func (m *KarmaAddressSource) String() string            { return proto.CompactTextString(m) }
func (*KarmaAddressSource) ProtoMessage()               {}
func (*KarmaAddressSource) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{9} }

func (m *KarmaAddressSource) GetUser() *types.Address {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *KarmaAddressSource) GetSources() []*KarmaSource {
	if m != nil {
		return m.Sources
	}
	return nil
}

type KarmaState struct {
	SourceStates     []*KarmaSource `protobuf:"bytes,1,rep,name=source_states,json=sourceStates" json:"source_states,omitempty"`
	DeployKarmaTotal int64          `protobuf:"varint,2,opt,name=deploy_karma_total,json=deployKarmaTotal,proto3" json:"deploy_karma_total,omitempty"`
	CallKarmaTotal   int64          `protobuf:"varint,3,opt,name=call_karma_total,json=callKarmaTotal,proto3" json:"call_karma_total,omitempty"`
	LastUpdateTime   int64          `protobuf:"varint,4,opt,name=last_update_time,json=lastUpdateTime,proto3" json:"last_update_time,omitempty"`
}

func (m *KarmaState) Reset()                    { *m = KarmaState{} }
func (m *KarmaState) String() string            { return proto.CompactTextString(m) }
func (*KarmaState) ProtoMessage()               {}
func (*KarmaState) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{10} }

func (m *KarmaState) GetSourceStates() []*KarmaSource {
	if m != nil {
		return m.SourceStates
	}
	return nil
}

func (m *KarmaState) GetDeployKarmaTotal() int64 {
	if m != nil {
		return m.DeployKarmaTotal
	}
	return 0
}

func (m *KarmaState) GetCallKarmaTotal() int64 {
	if m != nil {
		return m.CallKarmaTotal
	}
	return 0
}

func (m *KarmaState) GetLastUpdateTime() int64 {
	if m != nil {
		return m.LastUpdateTime
	}
	return 0
}

type KarmaStateUser struct {
	SourceStates []*KarmaSource `protobuf:"bytes,1,rep,name=source_states,json=sourceStates" json:"source_states,omitempty"`
	User         *types.Address `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Oracle       *types.Address `protobuf:"bytes,3,opt,name=oracle" json:"oracle,omitempty"`
}

func (m *KarmaStateUser) Reset()                    { *m = KarmaStateUser{} }
func (m *KarmaStateUser) String() string            { return proto.CompactTextString(m) }
func (*KarmaStateUser) ProtoMessage()               {}
func (*KarmaStateUser) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{11} }

func (m *KarmaStateUser) GetSourceStates() []*KarmaSource {
	if m != nil {
		return m.SourceStates
	}
	return nil
}

func (m *KarmaStateUser) GetUser() *types.Address {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *KarmaStateUser) GetOracle() *types.Address {
	if m != nil {
		return m.Oracle
	}
	return nil
}

type KarmaStateKeyUser struct {
	StateKeys []string       `protobuf:"bytes,1,rep,name=state_keys,json=stateKeys" json:"state_keys,omitempty"`
	User      *types.Address `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Oracle    *types.Address `protobuf:"bytes,3,opt,name=oracle" json:"oracle,omitempty"`
}

func (m *KarmaStateKeyUser) Reset()                    { *m = KarmaStateKeyUser{} }
func (m *KarmaStateKeyUser) String() string            { return proto.CompactTextString(m) }
func (*KarmaStateKeyUser) ProtoMessage()               {}
func (*KarmaStateKeyUser) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{12} }

func (m *KarmaStateKeyUser) GetStateKeys() []string {
	if m != nil {
		return m.StateKeys
	}
	return nil
}

func (m *KarmaStateKeyUser) GetUser() *types.Address {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *KarmaStateKeyUser) GetOracle() *types.Address {
	if m != nil {
		return m.Oracle
	}
	return nil
}

type KarmaTotal struct {
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *KarmaTotal) Reset()                    { *m = KarmaTotal{} }
func (m *KarmaTotal) String() string            { return proto.CompactTextString(m) }
func (*KarmaTotal) ProtoMessage()               {}
func (*KarmaTotal) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{13} }

func (m *KarmaTotal) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*KarmaInitRequest)(nil), "karma.KarmaInitRequest")
	proto.RegisterType((*KarmaSources)(nil), "karma.KarmaSources")
	proto.RegisterType((*KarmaSourcesValidator)(nil), "karma.KarmaSourcesValidator")
	proto.RegisterType((*KarmaNewOracleValidator)(nil), "karma.KarmaNewOracleValidator")
	proto.RegisterType((*KarmaUserTarget)(nil), "karma.KarmaUserTarget")
	proto.RegisterType((*KarmaUserAmount)(nil), "karma.KarmaUserAmount")
	proto.RegisterType((*KarmaSourceReward)(nil), "karma.KarmaSourceReward")
	proto.RegisterType((*KarmaSource)(nil), "karma.KarmaSource")
	proto.RegisterType((*KarmaUpkeepParams)(nil), "karma.KarmaUpkeepParams")
	proto.RegisterType((*KarmaAddressSource)(nil), "karma.KarmaAddressSource")
	proto.RegisterType((*KarmaState)(nil), "karma.KarmaState")
	proto.RegisterType((*KarmaStateUser)(nil), "karma.KarmaStateUser")
	proto.RegisterType((*KarmaStateKeyUser)(nil), "karma.KarmaStateKeyUser")
	proto.RegisterType((*KarmaTotal)(nil), "karma.KarmaTotal")
	proto.RegisterEnum("karma.KarmaSourceTarget", KarmaSourceTarget_name, KarmaSourceTarget_value)
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/karma/karma.proto", fileDescriptorKarma)
}

var fileDescriptorKarma = []byte{
	// 647 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x71, 0x9a, 0x36, 0xd3, 0x12, 0xc2, 0x8a, 0x9f, 0x80, 0x40, 0x8a, 0x7c, 0xa1, 0x42,
	0x25, 0xae, 0xc2, 0xa1, 0x37, 0xa4, 0x16, 0x38, 0x54, 0xad, 0x68, 0x59, 0x5a, 0x10, 0x27, 0xb3,
	0xb5, 0x57, 0xc1, 0x8a, 0xed, 0x75, 0x77, 0xd7, 0x8a, 0xf2, 0x08, 0x3c, 0x19, 0x07, 0x5e, 0x0a,
	0xed, 0xec, 0x26, 0x75, 0x68, 0xa2, 0x56, 0xe5, 0x62, 0x79, 0x66, 0xbe, 0x6f, 0xe6, 0x9b, 0xd9,
	0x59, 0x1b, 0xde, 0x8d, 0x52, 0xfd, 0xb3, 0xba, 0x18, 0xc4, 0x22, 0x0f, 0x33, 0x21, 0xf2, 0x82,
	0xeb, 0x89, 0x90, 0xe3, 0x70, 0x24, 0xde, 0x18, 0x33, 0xbc, 0xa8, 0xd2, 0x4c, 0xa7, 0x45, 0xa8,
	0xa7, 0x25, 0x57, 0xe1, 0x98, 0xc9, 0x9c, 0xd9, 0xe7, 0xa0, 0x94, 0x42, 0x0b, 0xb2, 0x86, 0xc6,
	0xf3, 0xdd, 0x1b, 0xd2, 0x58, 0x3a, 0x3e, 0x2d, 0x31, 0xf8, 0xe3, 0x41, 0xf7, 0xc8, 0x70, 0x0f,
	0x8b, 0x54, 0x53, 0x7e, 0x59, 0x71, 0xa5, 0x49, 0x1f, 0x5a, 0x42, 0xb2, 0x38, 0xe3, 0x3d, 0xaf,
	0xef, 0x6d, 0x6f, 0x0e, 0x37, 0x06, 0xfb, 0x49, 0x22, 0xb9, 0x52, 0xd4, 0xf9, 0xc9, 0x10, 0xd6,
	0x95, 0xa8, 0x64, 0xcc, 0x55, 0xaf, 0xd1, 0xf7, 0xb7, 0x37, 0x87, 0xbd, 0x81, 0x95, 0x83, 0xb9,
	0xbe, 0x60, 0x88, 0xf2, 0x09, 0x93, 0x09, 0x9d, 0x01, 0x49, 0x08, 0x6b, 0x95, 0xe2, 0x52, 0xf5,
	0x7c, 0x64, 0x3c, 0xab, 0x33, 0x5c, 0x7e, 0x47, 0xb4, 0x38, 0xb2, 0x0b, 0xad, 0xaa, 0x1c, 0x73,
	0x5e, 0xf6, 0x9a, 0x28, 0x63, 0xa1, 0xc6, 0x39, 0x46, 0x4e, 0x99, 0x64, 0xb9, 0xa2, 0x0e, 0x17,
	0x1c, 0xc0, 0x56, 0x4d, 0x80, 0xaa, 0xcb, 0xf4, 0x6e, 0x29, 0x33, 0xc8, 0xe1, 0x71, 0x3d, 0xc7,
	0x57, 0x96, 0xa5, 0x09, 0xd3, 0x42, 0xde, 0x25, 0x59, 0x6d, 0x92, 0x8d, 0xe5, 0x93, 0x0c, 0xc6,
	0xf0, 0x14, 0xf9, 0x9f, 0xf8, 0xe4, 0x04, 0x3d, 0x57, 0x05, 0x5f, 0x01, 0x14, 0x7c, 0x12, 0xad,
	0x38, 0x8a, 0x76, 0x31, 0x63, 0x18, 0xa0, 0xc8, 0x92, 0x68, 0x45, 0xa5, 0xb6, 0xc8, 0x12, 0x0b,
	0x0c, 0x18, 0x3c, 0xb0, 0xc3, 0x53, 0x5c, 0x9e, 0x31, 0x39, 0xe2, 0x9a, 0xbc, 0x80, 0xa6, 0x99,
	0xf6, 0xb5, 0xf4, 0xe8, 0x35, 0x47, 0xa0, 0x11, 0x87, 0x59, 0x3b, 0xcb, 0x5a, 0xb6, 0x79, 0xa8,
	0xc3, 0x05, 0x9f, 0x6b, 0x25, 0xf6, 0x73, 0x51, 0x15, 0x37, 0x95, 0xe8, 0x43, 0x8b, 0x21, 0x6e,
	0x2e, 0xfc, 0x20, 0x1d, 0x9d, 0x1f, 0x16, 0x9a, 0x3a, 0x7f, 0x70, 0x09, 0x0f, 0xaf, 0x8d, 0x98,
	0x10, 0x68, 0x16, 0x2c, 0xb7, 0x63, 0x69, 0x53, 0x7c, 0x27, 0x4f, 0xa0, 0x25, 0x31, 0x8a, 0xa9,
	0x7c, 0xea, 0xac, 0x5a, 0x17, 0xfe, 0x2d, 0xbb, 0xd8, 0x83, 0xcd, 0x5a, 0x70, 0x69, 0xb1, 0x47,
	0xb0, 0x16, 0xcf, 0x65, 0xfb, 0xd4, 0x1a, 0xc1, 0x37, 0xa7, 0xb5, 0xbe, 0x9e, 0x86, 0x1e, 0x0b,
	0xa5, 0x91, 0xee, 0x53, 0x7c, 0x37, 0x5a, 0xed, 0x92, 0x20, 0xbf, 0x4d, 0x9d, 0x65, 0xfc, 0x25,
	0x97, 0xa9, 0x48, 0x50, 0xab, 0x4f, 0x9d, 0x15, 0xfc, 0x00, 0x72, 0xfd, 0xa6, 0xdc, 0x30, 0xda,
	0x9d, 0x7f, 0x6f, 0x29, 0x59, 0xb2, 0xb1, 0xf3, 0xc5, 0xff, 0xed, 0x01, 0xd8, 0x80, 0x66, 0x9a,
	0x93, 0x3d, 0xb8, 0x6f, 0x23, 0x91, 0x32, 0xf6, 0x6c, 0xe9, 0x97, 0xa5, 0xd8, 0xb2, 0x40, 0xe4,
	0x29, 0xb2, 0x03, 0x24, 0xe1, 0x65, 0x26, 0xa6, 0x11, 0x22, 0x23, 0x2d, 0x34, 0xcb, 0xdc, 0x94,
	0xba, 0x36, 0x82, 0xe4, 0x33, 0xe3, 0x27, 0xdb, 0xd0, 0x8d, 0x59, 0x96, 0x2d, 0x60, 0x6d, 0xe7,
	0x1d, 0xe3, 0x5f, 0x44, 0x66, 0x4c, 0xe9, 0xa8, 0x2a, 0x13, 0xa6, 0x79, 0xa4, 0xd3, 0x9c, 0xe3,
	0x87, 0xc1, 0xa7, 0x1d, 0xe3, 0x3f, 0x47, 0xf7, 0x59, 0x9a, 0xf3, 0xe0, 0x97, 0x07, 0x9d, 0xab,
	0x4e, 0xcc, 0x26, 0xde, 0xbd, 0x9b, 0xd9, 0x84, 0x1b, 0xab, 0x96, 0xd7, 0xdd, 0x3a, 0x7f, 0xc5,
	0xfd, 0x96, 0xb3, 0xe5, 0x35, 0xe9, 0x8e, 0xf8, 0x14, 0xd5, 0xbc, 0x04, 0x40, 0x19, 0xd1, 0x98,
	0x4f, 0xad, 0x94, 0x36, 0x6d, 0x2b, 0x87, 0xf8, 0xff, 0x9a, 0x81, 0x3b, 0x48, 0x3b, 0xb7, 0xf9,
	0xa2, 0x7a, 0xb5, 0x45, 0x7d, 0x3d, 0x5c, 0xb8, 0x54, 0xee, 0x63, 0xb0, 0x01, 0xcd, 0xf7, 0xfb,
	0xc7, 0xc7, 0xdd, 0x7b, 0x04, 0xa0, 0xf5, 0xe1, 0xe3, 0xe9, 0xf1, 0xc9, 0xf7, 0xae, 0x47, 0xd6,
	0xc1, 0x37, 0xce, 0xc6, 0x45, 0x0b, 0xff, 0x19, 0x6f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe2,
	0x31, 0x74, 0x16, 0xae, 0x06, 0x00, 0x00,
}

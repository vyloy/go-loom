// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/plasma_cash/plasma_cash.proto

/*
Package plasma_cash is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/go-loom/builtin/types/plasma_cash/plasma_cash.proto

It has these top-level messages:
	PlasmaCashCoin
	PlasmaCashAccount
	PlasmaBlock
	PlasmaTx
	PlasmaBookKeeping
	GetCurrentBlockRequest
	GetCurrentBlockResponse
	GetBlockRequest
	GetBlockResponse
	SubmitBlockToMainnetRequest
	SubmitBlockToMainnetResponse
	PlasmaTxRequest
	PlasmaTxResponse
	GetPlasmaTxRequest
	GetPlasmaTxResponse
	GetUserSlotsRequest
	GetUserSlotsResponse
	DepositRequest
	DepositResponse
	PlasmaCashExitCoinRequest
	PlasmaCashWithdrawCoinRequest
	Pending
	PlasmaCashParams
	PlasmaCashInitRequest
	PlasmaCashBalanceOfRequest
	PlasmaCashBalanceOfResponse
	PlasmaDepositEvent
	PlasmaCashStartedExitEvent
	PlasmaCashFinalizedExitEvent
	PlasmaCashWithdrewEvent
	PlasmaCashTransferConfirmed
*/
package plasma_cash

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
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

type PlasmaCashCoinState int32

const (
	PlasmaCashCoinState_DEPOSITED  PlasmaCashCoinState = 0
	PlasmaCashCoinState_EXITING    PlasmaCashCoinState = 1
	PlasmaCashCoinState_CHALLENGED PlasmaCashCoinState = 2
	PlasmaCashCoinState_EXITED     PlasmaCashCoinState = 3
)

var PlasmaCashCoinState_name = map[int32]string{
	0: "DEPOSITED",
	1: "EXITING",
	2: "CHALLENGED",
	3: "EXITED",
}
var PlasmaCashCoinState_value = map[string]int32{
	"DEPOSITED":  0,
	"EXITING":    1,
	"CHALLENGED": 2,
	"EXITED":     3,
}

func (x PlasmaCashCoinState) String() string {
	return proto.EnumName(PlasmaCashCoinState_name, int32(x))
}
func (PlasmaCashCoinState) EnumDescriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{0} }

// Plasma Cash coin holds a single ERC721 token
type PlasmaCashCoin struct {
	// Unique ID
	Slot  uint64              `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	State PlasmaCashCoinState `protobuf:"varint,2,opt,name=state,proto3,enum=PlasmaCashCoinState" json:"state,omitempty"`
	// ERC721 token ID
	Token *types.BigUInt `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	// ERC721 token contract address
	Contract *types.Address `protobuf:"bytes,4,opt,name=contract" json:"contract,omitempty"`
}

func (m *PlasmaCashCoin) Reset()                    { *m = PlasmaCashCoin{} }
func (m *PlasmaCashCoin) String() string            { return proto.CompactTextString(m) }
func (*PlasmaCashCoin) ProtoMessage()               {}
func (*PlasmaCashCoin) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{0} }

func (m *PlasmaCashCoin) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *PlasmaCashCoin) GetState() PlasmaCashCoinState {
	if m != nil {
		return m.State
	}
	return PlasmaCashCoinState_DEPOSITED
}

func (m *PlasmaCashCoin) GetToken() *types.BigUInt {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *PlasmaCashCoin) GetContract() *types.Address {
	if m != nil {
		return m.Contract
	}
	return nil
}

type PlasmaCashAccount struct {
	Owner *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	// Plasma coins in this account, identified by their slot number.
	Slots []uint64 `protobuf:"varint,3,rep,packed,name=slots" json:"slots,omitempty"`
}

func (m *PlasmaCashAccount) Reset()                    { *m = PlasmaCashAccount{} }
func (m *PlasmaCashAccount) String() string            { return proto.CompactTextString(m) }
func (*PlasmaCashAccount) ProtoMessage()               {}
func (*PlasmaCashAccount) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{1} }

func (m *PlasmaCashAccount) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *PlasmaCashAccount) GetSlots() []uint64 {
	if m != nil {
		return m.Slots
	}
	return nil
}

type PlasmaBlock struct {
	Uid          *types.BigUInt `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	Transactions []*PlasmaTx    `protobuf:"bytes,2,rep,name=transactions" json:"transactions,omitempty"`
	Signature    []byte         `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	MerkleHash   []byte         `protobuf:"bytes,4,opt,name=merkle_hash,json=merkleHash,proto3" json:"merkle_hash,omitempty"`
	Hash         []byte         `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
	Proof        []byte         `protobuf:"bytes,6,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *PlasmaBlock) Reset()                    { *m = PlasmaBlock{} }
func (m *PlasmaBlock) String() string            { return proto.CompactTextString(m) }
func (*PlasmaBlock) ProtoMessage()               {}
func (*PlasmaBlock) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{2} }

func (m *PlasmaBlock) GetUid() *types.BigUInt {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *PlasmaBlock) GetTransactions() []*PlasmaTx {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *PlasmaBlock) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *PlasmaBlock) GetMerkleHash() []byte {
	if m != nil {
		return m.MerkleHash
	}
	return nil
}

func (m *PlasmaBlock) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *PlasmaBlock) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

type PlasmaTx struct {
	Slot          uint64         `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	PreviousBlock *types.BigUInt `protobuf:"bytes,2,opt,name=previous_block,json=previousBlock" json:"previous_block,omitempty"`
	Denomination  *types.BigUInt `protobuf:"bytes,3,opt,name=denomination" json:"denomination,omitempty"`
	NewOwner      *types.Address `protobuf:"bytes,4,opt,name=new_owner,json=newOwner" json:"new_owner,omitempty"`
	Signature     []byte         `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	Hash          []byte         `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
	MerkleHash    []byte         `protobuf:"bytes,7,opt,name=merkle_hash,json=merkleHash,proto3" json:"merkle_hash,omitempty"`
	Sender        *types.Address `protobuf:"bytes,8,opt,name=sender" json:"sender,omitempty"`
	Proof         []byte         `protobuf:"bytes,9,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *PlasmaTx) Reset()                    { *m = PlasmaTx{} }
func (m *PlasmaTx) String() string            { return proto.CompactTextString(m) }
func (*PlasmaTx) ProtoMessage()               {}
func (*PlasmaTx) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{3} }

func (m *PlasmaTx) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *PlasmaTx) GetPreviousBlock() *types.BigUInt {
	if m != nil {
		return m.PreviousBlock
	}
	return nil
}

func (m *PlasmaTx) GetDenomination() *types.BigUInt {
	if m != nil {
		return m.Denomination
	}
	return nil
}

func (m *PlasmaTx) GetNewOwner() *types.Address {
	if m != nil {
		return m.NewOwner
	}
	return nil
}

func (m *PlasmaTx) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *PlasmaTx) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *PlasmaTx) GetMerkleHash() []byte {
	if m != nil {
		return m.MerkleHash
	}
	return nil
}

func (m *PlasmaTx) GetSender() *types.Address {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *PlasmaTx) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

type PlasmaBookKeeping struct {
	CurrentHeight *types.BigUInt `protobuf:"bytes,1,opt,name=current_height,json=currentHeight" json:"current_height,omitempty"`
}

func (m *PlasmaBookKeeping) Reset()                    { *m = PlasmaBookKeeping{} }
func (m *PlasmaBookKeeping) String() string            { return proto.CompactTextString(m) }
func (*PlasmaBookKeeping) ProtoMessage()               {}
func (*PlasmaBookKeeping) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{4} }

func (m *PlasmaBookKeeping) GetCurrentHeight() *types.BigUInt {
	if m != nil {
		return m.CurrentHeight
	}
	return nil
}

type GetCurrentBlockRequest struct {
}

func (m *GetCurrentBlockRequest) Reset()                    { *m = GetCurrentBlockRequest{} }
func (m *GetCurrentBlockRequest) String() string            { return proto.CompactTextString(m) }
func (*GetCurrentBlockRequest) ProtoMessage()               {}
func (*GetCurrentBlockRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{5} }

type GetCurrentBlockResponse struct {
	BlockHeight *types.BigUInt `protobuf:"bytes,1,opt,name=block_height,json=blockHeight" json:"block_height,omitempty"`
}

func (m *GetCurrentBlockResponse) Reset()         { *m = GetCurrentBlockResponse{} }
func (m *GetCurrentBlockResponse) String() string { return proto.CompactTextString(m) }
func (*GetCurrentBlockResponse) ProtoMessage()    {}
func (*GetCurrentBlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{6}
}

func (m *GetCurrentBlockResponse) GetBlockHeight() *types.BigUInt {
	if m != nil {
		return m.BlockHeight
	}
	return nil
}

type GetBlockRequest struct {
	BlockHeight *types.BigUInt `protobuf:"bytes,1,opt,name=block_height,json=blockHeight" json:"block_height,omitempty"`
}

func (m *GetBlockRequest) Reset()                    { *m = GetBlockRequest{} }
func (m *GetBlockRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBlockRequest) ProtoMessage()               {}
func (*GetBlockRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{7} }

func (m *GetBlockRequest) GetBlockHeight() *types.BigUInt {
	if m != nil {
		return m.BlockHeight
	}
	return nil
}

type GetBlockResponse struct {
	Block *PlasmaBlock `protobuf:"bytes,1,opt,name=block" json:"block,omitempty"`
}

func (m *GetBlockResponse) Reset()                    { *m = GetBlockResponse{} }
func (m *GetBlockResponse) String() string            { return proto.CompactTextString(m) }
func (*GetBlockResponse) ProtoMessage()               {}
func (*GetBlockResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{8} }

func (m *GetBlockResponse) GetBlock() *PlasmaBlock {
	if m != nil {
		return m.Block
	}
	return nil
}

// This only originates from the validator
type SubmitBlockToMainnetRequest struct {
}

func (m *SubmitBlockToMainnetRequest) Reset()         { *m = SubmitBlockToMainnetRequest{} }
func (m *SubmitBlockToMainnetRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitBlockToMainnetRequest) ProtoMessage()    {}
func (*SubmitBlockToMainnetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{9}
}

type SubmitBlockToMainnetResponse struct {
	MerkleHash []byte `protobuf:"bytes,1,opt,name=merkle_hash,json=merkleHash,proto3" json:"merkle_hash,omitempty"`
}

func (m *SubmitBlockToMainnetResponse) Reset()         { *m = SubmitBlockToMainnetResponse{} }
func (m *SubmitBlockToMainnetResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitBlockToMainnetResponse) ProtoMessage()    {}
func (*SubmitBlockToMainnetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{10}
}

func (m *SubmitBlockToMainnetResponse) GetMerkleHash() []byte {
	if m != nil {
		return m.MerkleHash
	}
	return nil
}

type PlasmaTxRequest struct {
	Plasmatx *PlasmaTx `protobuf:"bytes,1,opt,name=plasmatx" json:"plasmatx,omitempty"`
}

func (m *PlasmaTxRequest) Reset()                    { *m = PlasmaTxRequest{} }
func (m *PlasmaTxRequest) String() string            { return proto.CompactTextString(m) }
func (*PlasmaTxRequest) ProtoMessage()               {}
func (*PlasmaTxRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{11} }

func (m *PlasmaTxRequest) GetPlasmatx() *PlasmaTx {
	if m != nil {
		return m.Plasmatx
	}
	return nil
}

type PlasmaTxResponse struct {
}

func (m *PlasmaTxResponse) Reset()                    { *m = PlasmaTxResponse{} }
func (m *PlasmaTxResponse) String() string            { return proto.CompactTextString(m) }
func (*PlasmaTxResponse) ProtoMessage()               {}
func (*PlasmaTxResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{12} }

type GetPlasmaTxRequest struct {
	Slot        uint64         `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	BlockHeight *types.BigUInt `protobuf:"bytes,2,opt,name=block_height,json=blockHeight" json:"block_height,omitempty"`
}

func (m *GetPlasmaTxRequest) Reset()                    { *m = GetPlasmaTxRequest{} }
func (m *GetPlasmaTxRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPlasmaTxRequest) ProtoMessage()               {}
func (*GetPlasmaTxRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{13} }

func (m *GetPlasmaTxRequest) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *GetPlasmaTxRequest) GetBlockHeight() *types.BigUInt {
	if m != nil {
		return m.BlockHeight
	}
	return nil
}

type GetPlasmaTxResponse struct {
	Plasmatx *PlasmaTx `protobuf:"bytes,1,opt,name=plasmatx" json:"plasmatx,omitempty"`
}

func (m *GetPlasmaTxResponse) Reset()                    { *m = GetPlasmaTxResponse{} }
func (m *GetPlasmaTxResponse) String() string            { return proto.CompactTextString(m) }
func (*GetPlasmaTxResponse) ProtoMessage()               {}
func (*GetPlasmaTxResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{14} }

func (m *GetPlasmaTxResponse) GetPlasmatx() *PlasmaTx {
	if m != nil {
		return m.Plasmatx
	}
	return nil
}

type GetUserSlotsRequest struct {
	From *types.Address `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
}

func (m *GetUserSlotsRequest) Reset()                    { *m = GetUserSlotsRequest{} }
func (m *GetUserSlotsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserSlotsRequest) ProtoMessage()               {}
func (*GetUserSlotsRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{15} }

func (m *GetUserSlotsRequest) GetFrom() *types.Address {
	if m != nil {
		return m.From
	}
	return nil
}

type GetUserSlotsResponse struct {
	Slots []uint64 `protobuf:"varint,1,rep,packed,name=slots" json:"slots,omitempty"`
}

func (m *GetUserSlotsResponse) Reset()                    { *m = GetUserSlotsResponse{} }
func (m *GetUserSlotsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUserSlotsResponse) ProtoMessage()               {}
func (*GetUserSlotsResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{16} }

func (m *GetUserSlotsResponse) GetSlots() []uint64 {
	if m != nil {
		return m.Slots
	}
	return nil
}

// This only originates from the validator
type DepositRequest struct {
	Slot         uint64         `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	DepositBlock *types.BigUInt `protobuf:"bytes,2,opt,name=deposit_block,json=depositBlock" json:"deposit_block,omitempty"`
	// For ERC20 this is the number of coins deposited, for ERC721 this is a token ID.
	Denomination *types.BigUInt `protobuf:"bytes,3,opt,name=denomination" json:"denomination,omitempty"`
	// Entity that made the deposit
	From *types.Address `protobuf:"bytes,4,opt,name=from" json:"from,omitempty"`
	// Contract from which the coins originated (i.e. the currency of the coins)
	Contract *types.Address `protobuf:"bytes,5,opt,name=contract" json:"contract,omitempty"`
}

func (m *DepositRequest) Reset()                    { *m = DepositRequest{} }
func (m *DepositRequest) String() string            { return proto.CompactTextString(m) }
func (*DepositRequest) ProtoMessage()               {}
func (*DepositRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{17} }

func (m *DepositRequest) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *DepositRequest) GetDepositBlock() *types.BigUInt {
	if m != nil {
		return m.DepositBlock
	}
	return nil
}

func (m *DepositRequest) GetDenomination() *types.BigUInt {
	if m != nil {
		return m.Denomination
	}
	return nil
}

func (m *DepositRequest) GetFrom() *types.Address {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *DepositRequest) GetContract() *types.Address {
	if m != nil {
		return m.Contract
	}
	return nil
}

type DepositResponse struct {
}

func (m *DepositResponse) Reset()                    { *m = DepositResponse{} }
func (m *DepositResponse) String() string            { return proto.CompactTextString(m) }
func (*DepositResponse) ProtoMessage()               {}
func (*DepositResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{18} }

type PlasmaCashExitCoinRequest struct {
	Owner *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Slot  uint64         `protobuf:"varint,2,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (m *PlasmaCashExitCoinRequest) Reset()         { *m = PlasmaCashExitCoinRequest{} }
func (m *PlasmaCashExitCoinRequest) String() string { return proto.CompactTextString(m) }
func (*PlasmaCashExitCoinRequest) ProtoMessage()    {}
func (*PlasmaCashExitCoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{19}
}

func (m *PlasmaCashExitCoinRequest) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *PlasmaCashExitCoinRequest) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

type PlasmaCashWithdrawCoinRequest struct {
	Owner *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Slot  uint64         `protobuf:"varint,2,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (m *PlasmaCashWithdrawCoinRequest) Reset()         { *m = PlasmaCashWithdrawCoinRequest{} }
func (m *PlasmaCashWithdrawCoinRequest) String() string { return proto.CompactTextString(m) }
func (*PlasmaCashWithdrawCoinRequest) ProtoMessage()    {}
func (*PlasmaCashWithdrawCoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{20}
}

func (m *PlasmaCashWithdrawCoinRequest) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *PlasmaCashWithdrawCoinRequest) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

type Pending struct {
	Transactions []*PlasmaTx `protobuf:"bytes,1,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *Pending) Reset()                    { *m = Pending{} }
func (m *Pending) String() string            { return proto.CompactTextString(m) }
func (*Pending) ProtoMessage()               {}
func (*Pending) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{21} }

func (m *Pending) GetTransactions() []*PlasmaTx {
	if m != nil {
		return m.Transactions
	}
	return nil
}

// Initialization of state from Genesis.json
type PlasmaCashParams struct {
	BlockInterval uint64 `protobuf:"varint,1,opt,name=block_interval,json=blockInterval,proto3" json:"block_interval,omitempty"`
}

func (m *PlasmaCashParams) Reset()                    { *m = PlasmaCashParams{} }
func (m *PlasmaCashParams) String() string            { return proto.CompactTextString(m) }
func (*PlasmaCashParams) ProtoMessage()               {}
func (*PlasmaCashParams) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{22} }

func (m *PlasmaCashParams) GetBlockInterval() uint64 {
	if m != nil {
		return m.BlockInterval
	}
	return 0
}

type PlasmaCashInitRequest struct {
	Params *PlasmaCashParams `protobuf:"bytes,1,opt,name=params" json:"params,omitempty"`
}

func (m *PlasmaCashInitRequest) Reset()                    { *m = PlasmaCashInitRequest{} }
func (m *PlasmaCashInitRequest) String() string            { return proto.CompactTextString(m) }
func (*PlasmaCashInitRequest) ProtoMessage()               {}
func (*PlasmaCashInitRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{23} }

func (m *PlasmaCashInitRequest) GetParams() *PlasmaCashParams {
	if m != nil {
		return m.Params
	}
	return nil
}

type PlasmaCashBalanceOfRequest struct {
	Owner    *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Contract *types.Address `protobuf:"bytes,2,opt,name=contract" json:"contract,omitempty"`
}

func (m *PlasmaCashBalanceOfRequest) Reset()         { *m = PlasmaCashBalanceOfRequest{} }
func (m *PlasmaCashBalanceOfRequest) String() string { return proto.CompactTextString(m) }
func (*PlasmaCashBalanceOfRequest) ProtoMessage()    {}
func (*PlasmaCashBalanceOfRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{24}
}

func (m *PlasmaCashBalanceOfRequest) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *PlasmaCashBalanceOfRequest) GetContract() *types.Address {
	if m != nil {
		return m.Contract
	}
	return nil
}

type PlasmaCashBalanceOfResponse struct {
	Coins []*PlasmaCashCoin `protobuf:"bytes,1,rep,name=coins" json:"coins,omitempty"`
}

func (m *PlasmaCashBalanceOfResponse) Reset()         { *m = PlasmaCashBalanceOfResponse{} }
func (m *PlasmaCashBalanceOfResponse) String() string { return proto.CompactTextString(m) }
func (*PlasmaCashBalanceOfResponse) ProtoMessage()    {}
func (*PlasmaCashBalanceOfResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{25}
}

func (m *PlasmaCashBalanceOfResponse) GetCoins() []*PlasmaCashCoin {
	if m != nil {
		return m.Coins
	}
	return nil
}

type PlasmaDepositEvent struct {
	Slot         uint64         `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	DepositBlock *types.BigUInt `protobuf:"bytes,2,opt,name=deposit_block,json=depositBlock" json:"deposit_block,omitempty"`
	// For ERC20 this is the number of coins deposited, for ERC721 this is a token ID.
	Denomination *types.BigUInt `protobuf:"bytes,3,opt,name=denomination" json:"denomination,omitempty"`
	// Entity that made the deposit
	From *types.Address `protobuf:"bytes,4,opt,name=from" json:"from,omitempty"`
	// Contract from which the coins originated (i.e. the currency of the coins)
	Contract *types.Address `protobuf:"bytes,5,opt,name=contract" json:"contract,omitempty"`
}

func (m *PlasmaDepositEvent) Reset()                    { *m = PlasmaDepositEvent{} }
func (m *PlasmaDepositEvent) String() string            { return proto.CompactTextString(m) }
func (*PlasmaDepositEvent) ProtoMessage()               {}
func (*PlasmaDepositEvent) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{26} }

func (m *PlasmaDepositEvent) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *PlasmaDepositEvent) GetDepositBlock() *types.BigUInt {
	if m != nil {
		return m.DepositBlock
	}
	return nil
}

func (m *PlasmaDepositEvent) GetDenomination() *types.BigUInt {
	if m != nil {
		return m.Denomination
	}
	return nil
}

func (m *PlasmaDepositEvent) GetFrom() *types.Address {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *PlasmaDepositEvent) GetContract() *types.Address {
	if m != nil {
		return m.Contract
	}
	return nil
}

type PlasmaCashStartedExitEvent struct {
	Owner *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Slot  uint64         `protobuf:"varint,2,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (m *PlasmaCashStartedExitEvent) Reset()         { *m = PlasmaCashStartedExitEvent{} }
func (m *PlasmaCashStartedExitEvent) String() string { return proto.CompactTextString(m) }
func (*PlasmaCashStartedExitEvent) ProtoMessage()    {}
func (*PlasmaCashStartedExitEvent) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{27}
}

func (m *PlasmaCashStartedExitEvent) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *PlasmaCashStartedExitEvent) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

type PlasmaCashFinalizedExitEvent struct {
	Owner *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Slot  uint64         `protobuf:"varint,2,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (m *PlasmaCashFinalizedExitEvent) Reset()         { *m = PlasmaCashFinalizedExitEvent{} }
func (m *PlasmaCashFinalizedExitEvent) String() string { return proto.CompactTextString(m) }
func (*PlasmaCashFinalizedExitEvent) ProtoMessage()    {}
func (*PlasmaCashFinalizedExitEvent) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{28}
}

func (m *PlasmaCashFinalizedExitEvent) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *PlasmaCashFinalizedExitEvent) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

type PlasmaCashWithdrewEvent struct {
	Owner        *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Mode         uint32         `protobuf:"varint,2,opt,name=mode,proto3" json:"mode,omitempty"`
	Contract     *types.Address `protobuf:"bytes,3,opt,name=contract" json:"contract,omitempty"`
	Uid          *types.BigUInt `protobuf:"bytes,4,opt,name=uid" json:"uid,omitempty"`
	Denomination *types.BigUInt `protobuf:"bytes,5,opt,name=denomination" json:"denomination,omitempty"`
	Slot         uint64         `protobuf:"varint,6,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (m *PlasmaCashWithdrewEvent) Reset()         { *m = PlasmaCashWithdrewEvent{} }
func (m *PlasmaCashWithdrewEvent) String() string { return proto.CompactTextString(m) }
func (*PlasmaCashWithdrewEvent) ProtoMessage()    {}
func (*PlasmaCashWithdrewEvent) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{29}
}

func (m *PlasmaCashWithdrewEvent) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *PlasmaCashWithdrewEvent) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *PlasmaCashWithdrewEvent) GetContract() *types.Address {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *PlasmaCashWithdrewEvent) GetUid() *types.BigUInt {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *PlasmaCashWithdrewEvent) GetDenomination() *types.BigUInt {
	if m != nil {
		return m.Denomination
	}
	return nil
}

func (m *PlasmaCashWithdrewEvent) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

type PlasmaCashTransferConfirmed struct {
	From *types.Address `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To   *types.Address `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	Slot uint64         `protobuf:"varint,3,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (m *PlasmaCashTransferConfirmed) Reset()         { *m = PlasmaCashTransferConfirmed{} }
func (m *PlasmaCashTransferConfirmed) String() string { return proto.CompactTextString(m) }
func (*PlasmaCashTransferConfirmed) ProtoMessage()    {}
func (*PlasmaCashTransferConfirmed) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{30}
}

func (m *PlasmaCashTransferConfirmed) GetFrom() *types.Address {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *PlasmaCashTransferConfirmed) GetTo() *types.Address {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *PlasmaCashTransferConfirmed) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func init() {
	proto.RegisterType((*PlasmaCashCoin)(nil), "PlasmaCashCoin")
	proto.RegisterType((*PlasmaCashAccount)(nil), "PlasmaCashAccount")
	proto.RegisterType((*PlasmaBlock)(nil), "PlasmaBlock")
	proto.RegisterType((*PlasmaTx)(nil), "PlasmaTx")
	proto.RegisterType((*PlasmaBookKeeping)(nil), "PlasmaBookKeeping")
	proto.RegisterType((*GetCurrentBlockRequest)(nil), "GetCurrentBlockRequest")
	proto.RegisterType((*GetCurrentBlockResponse)(nil), "GetCurrentBlockResponse")
	proto.RegisterType((*GetBlockRequest)(nil), "GetBlockRequest")
	proto.RegisterType((*GetBlockResponse)(nil), "GetBlockResponse")
	proto.RegisterType((*SubmitBlockToMainnetRequest)(nil), "SubmitBlockToMainnetRequest")
	proto.RegisterType((*SubmitBlockToMainnetResponse)(nil), "SubmitBlockToMainnetResponse")
	proto.RegisterType((*PlasmaTxRequest)(nil), "PlasmaTxRequest")
	proto.RegisterType((*PlasmaTxResponse)(nil), "PlasmaTxResponse")
	proto.RegisterType((*GetPlasmaTxRequest)(nil), "GetPlasmaTxRequest")
	proto.RegisterType((*GetPlasmaTxResponse)(nil), "GetPlasmaTxResponse")
	proto.RegisterType((*GetUserSlotsRequest)(nil), "GetUserSlotsRequest")
	proto.RegisterType((*GetUserSlotsResponse)(nil), "GetUserSlotsResponse")
	proto.RegisterType((*DepositRequest)(nil), "DepositRequest")
	proto.RegisterType((*DepositResponse)(nil), "DepositResponse")
	proto.RegisterType((*PlasmaCashExitCoinRequest)(nil), "PlasmaCashExitCoinRequest")
	proto.RegisterType((*PlasmaCashWithdrawCoinRequest)(nil), "PlasmaCashWithdrawCoinRequest")
	proto.RegisterType((*Pending)(nil), "Pending")
	proto.RegisterType((*PlasmaCashParams)(nil), "PlasmaCashParams")
	proto.RegisterType((*PlasmaCashInitRequest)(nil), "PlasmaCashInitRequest")
	proto.RegisterType((*PlasmaCashBalanceOfRequest)(nil), "PlasmaCashBalanceOfRequest")
	proto.RegisterType((*PlasmaCashBalanceOfResponse)(nil), "PlasmaCashBalanceOfResponse")
	proto.RegisterType((*PlasmaDepositEvent)(nil), "PlasmaDepositEvent")
	proto.RegisterType((*PlasmaCashStartedExitEvent)(nil), "PlasmaCashStartedExitEvent")
	proto.RegisterType((*PlasmaCashFinalizedExitEvent)(nil), "PlasmaCashFinalizedExitEvent")
	proto.RegisterType((*PlasmaCashWithdrewEvent)(nil), "PlasmaCashWithdrewEvent")
	proto.RegisterType((*PlasmaCashTransferConfirmed)(nil), "PlasmaCashTransferConfirmed")
	proto.RegisterEnum("PlasmaCashCoinState", PlasmaCashCoinState_name, PlasmaCashCoinState_value)
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/plasma_cash/plasma_cash.proto", fileDescriptorPlasmaCash)
}

var fileDescriptorPlasmaCash = []byte{
	// 1097 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xdd, 0x4e, 0x1b, 0x47,
	0x14, 0xee, 0xfa, 0x0f, 0x73, 0xfc, 0x83, 0x99, 0xd0, 0xc6, 0x25, 0xa4, 0x58, 0xab, 0x46, 0xa2,
	0x51, 0xb0, 0x2b, 0x90, 0xaa, 0x54, 0x8a, 0x54, 0x61, 0xec, 0x82, 0xdb, 0x04, 0xac, 0x35, 0x11,
	0xbd, 0x43, 0xeb, 0xdd, 0xb1, 0x3d, 0xc2, 0x9e, 0x71, 0x67, 0xc6, 0x98, 0xf6, 0x41, 0xfa, 0x3e,
	0xbd, 0xe8, 0x45, 0xef, 0xda, 0x17, 0xe0, 0x82, 0x27, 0xa9, 0x76, 0x66, 0xd7, 0xbb, 0x5e, 0x1b,
	0x42, 0xc3, 0x55, 0x6e, 0x56, 0x3b, 0xe7, 0xcc, 0xf9, 0xf9, 0xbe, 0x73, 0xe6, 0xcc, 0xc0, 0x4f,
	0x7d, 0x22, 0x07, 0x93, 0x6e, 0xd5, 0x61, 0xa3, 0xda, 0x90, 0xb1, 0x11, 0xc5, 0x72, 0xca, 0xf8,
	0x65, 0xad, 0xcf, 0x76, 0xbd, 0x65, 0xad, 0x3b, 0x21, 0x43, 0x49, 0x68, 0x4d, 0xfe, 0x36, 0xc6,
	0xa2, 0x36, 0x1e, 0xda, 0x62, 0x64, 0x5f, 0x38, 0xb6, 0x18, 0x44, 0xff, 0xab, 0x63, 0xce, 0x24,
	0xdb, 0xdc, 0x8d, 0xf8, 0xea, 0xb3, 0x3e, 0xab, 0x29, 0x71, 0x77, 0xd2, 0x53, 0x2b, 0xb5, 0x50,
	0x7f, 0xfe, 0xf6, 0x6f, 0x3f, 0x10, 0x5a, 0x87, 0x54, 0x5f, 0x6d, 0x61, 0xfe, 0x61, 0x40, 0xb1,
	0xad, 0xc2, 0x1e, 0xda, 0x62, 0x70, 0xc8, 0x08, 0x45, 0x08, 0x52, 0x62, 0xc8, 0x64, 0xd9, 0xa8,
	0x18, 0x3b, 0x29, 0x4b, 0xfd, 0xa3, 0x97, 0x90, 0x16, 0xd2, 0x96, 0xb8, 0x9c, 0xa8, 0x18, 0x3b,
	0xc5, 0xbd, 0x8d, 0xea, 0xbc, 0x4d, 0xc7, 0xd3, 0x59, 0x7a, 0x0b, 0xfa, 0x0a, 0xd2, 0x92, 0x5d,
	0x62, 0x5a, 0x4e, 0x56, 0x8c, 0x9d, 0xdc, 0x5e, 0xb6, 0x5a, 0x27, 0xfd, 0xf7, 0x2d, 0x2a, 0x2d,
	0x2d, 0x46, 0x5f, 0x43, 0xd6, 0x61, 0x54, 0x72, 0xdb, 0x91, 0xe5, 0x94, 0xbf, 0xe5, 0xc0, 0x75,
	0x39, 0x16, 0xc2, 0x9a, 0x69, 0xcc, 0x16, 0xac, 0x87, 0x31, 0x0e, 0x1c, 0x87, 0x4d, 0xa8, 0xf4,
	0x5c, 0xb3, 0x29, 0xc5, 0x5c, 0xe5, 0x16, 0xb5, 0xd3, 0x62, 0xb4, 0x01, 0x69, 0x2f, 0x5d, 0x51,
	0x4e, 0x56, 0x92, 0x3b, 0x29, 0x4b, 0x2f, 0xcc, 0x3f, 0x0d, 0xc8, 0x69, 0x5f, 0xf5, 0x21, 0x73,
	0x2e, 0xd1, 0x26, 0x24, 0x27, 0xc4, 0x9d, 0xf9, 0x08, 0xd2, 0xf3, 0x84, 0x68, 0x17, 0xf2, 0x92,
	0xdb, 0x54, 0xd8, 0x8e, 0x24, 0x8c, 0x8a, 0x72, 0xa2, 0x92, 0xdc, 0xc9, 0xed, 0xad, 0xfa, 0x78,
	0xcf, 0xae, 0xad, 0x39, 0x35, 0xda, 0x82, 0x55, 0x41, 0xfa, 0xd4, 0x96, 0x13, 0x8e, 0x15, 0xde,
	0xbc, 0x15, 0x0a, 0xd0, 0x36, 0xe4, 0x46, 0x98, 0x5f, 0x0e, 0xf1, 0xc5, 0xc0, 0x16, 0x03, 0x05,
	0x36, 0x6f, 0x81, 0x16, 0x1d, 0xdb, 0x62, 0xe0, 0x51, 0xad, 0x34, 0x69, 0xa5, 0x51, 0xff, 0x1e,
	0x86, 0x31, 0x67, 0xac, 0x57, 0xce, 0x28, 0xa1, 0x5e, 0x98, 0xff, 0x24, 0x20, 0x1b, 0xe4, 0xb0,
	0xb4, 0x42, 0x75, 0x28, 0x8e, 0x39, 0xbe, 0x22, 0x6c, 0x22, 0x2e, 0xba, 0x1e, 0x4c, 0x55, 0xaa,
	0x08, 0xbe, 0xfa, 0xfa, 0xed, 0xcd, 0x76, 0xa1, 0xed, 0xef, 0x51, 0x4c, 0x58, 0x85, 0x71, 0x74,
	0x89, 0x5e, 0x41, 0xde, 0xc5, 0x94, 0x8d, 0x08, 0xb5, 0x3d, 0x78, 0x0b, 0x05, 0x9c, 0xd3, 0xa2,
	0x7d, 0x58, 0xa5, 0x78, 0x7a, 0xa1, 0x0b, 0x12, 0x2b, 0x64, 0x3d, 0x7f, 0x7b, 0xb3, 0x9d, 0x3d,
	0xc1, 0xd3, 0x53, 0x4f, 0x6b, 0x65, 0xa9, 0xff, 0x37, 0x4f, 0x58, 0x3a, 0x4e, 0x58, 0xc0, 0x47,
	0x26, 0xc2, 0x47, 0x8c, 0xc4, 0x95, 0x05, 0x12, 0x2b, 0x90, 0x11, 0x98, 0xba, 0x98, 0x97, 0xb3,
	0xb1, 0xae, 0xf0, 0xe5, 0x21, 0xa5, 0xab, 0x51, 0x4a, 0xcf, 0x83, 0x0e, 0xab, 0x33, 0x76, 0xf9,
	0x33, 0xc6, 0x63, 0x42, 0xfb, 0x1e, 0x8d, 0xce, 0x84, 0x73, 0x4c, 0xe5, 0xc5, 0x00, 0x93, 0xfe,
	0x40, 0xc6, 0xdb, 0x44, 0xd3, 0x78, 0xa8, 0xf7, 0x1c, 0xab, 0x2d, 0x56, 0xc1, 0x89, 0x2e, 0xcd,
	0x32, 0x7c, 0x71, 0x84, 0xa5, 0xbf, 0x45, 0x13, 0x8d, 0x7f, 0x9d, 0x60, 0x21, 0xcd, 0x73, 0x78,
	0xba, 0xa0, 0x11, 0x63, 0x46, 0x05, 0x46, 0x6f, 0x20, 0xaf, 0xca, 0x76, 0x57, 0xd8, 0xb5, 0xdb,
	0x9b, 0xed, 0x9c, 0x32, 0xf1, 0x83, 0xe6, 0xba, 0xe1, 0xc2, 0x3c, 0x85, 0xb5, 0x23, 0x3c, 0x17,
	0xeb, 0x91, 0x0e, 0xbf, 0x83, 0x52, 0xe8, 0xd0, 0x4f, 0xd1, 0x84, 0xb4, 0xee, 0x2c, 0xed, 0x2a,
	0x5f, 0x8d, 0x1c, 0x2a, 0x4b, 0xab, 0xcc, 0xe7, 0xf0, 0xac, 0x33, 0xe9, 0x8e, 0x88, 0x36, 0x3d,
	0x63, 0xef, 0x6c, 0x42, 0x29, 0x96, 0x01, 0x01, 0x3f, 0xc0, 0xd6, 0x72, 0xb5, 0x1f, 0x22, 0x56,
	0x6c, 0x23, 0x5e, 0x6c, 0xf3, 0x35, 0xac, 0xcd, 0x8e, 0xa2, 0x0f, 0xf4, 0x05, 0x64, 0xf5, 0xe0,
	0x94, 0xd7, 0x7e, 0x66, 0x91, 0xe3, 0x3a, 0x53, 0x99, 0x08, 0x4a, 0xa1, 0xa5, 0x0e, 0x67, 0xf6,
	0x00, 0x1d, 0x61, 0x19, 0x77, 0xb8, 0xec, 0x78, 0xc5, 0xd9, 0x4c, 0xfc, 0x2f, 0x36, 0xdf, 0xc0,
	0x93, 0xb9, 0x38, 0x3e, 0xda, 0x07, 0x66, 0xbe, 0xaf, 0xac, 0xdf, 0x0b, 0xcc, 0x3b, 0xde, 0x3c,
	0x0b, 0xd2, 0xdc, 0x82, 0x54, 0x8f, 0xb3, 0xd1, 0xc2, 0x2c, 0x54, 0x52, 0xf3, 0x15, 0x6c, 0xcc,
	0x1b, 0xf9, 0x31, 0x67, 0x23, 0xd2, 0x88, 0x8e, 0xc8, 0xbf, 0x0c, 0x28, 0x36, 0xf0, 0x98, 0x09,
	0x22, 0xef, 0x63, 0x61, 0x17, 0x0a, 0xae, 0xde, 0xb5, 0x7c, 0xc6, 0x78, 0x13, 0x42, 0xa9, 0x3f,
	0x66, 0x9e, 0x04, 0x78, 0x52, 0xcb, 0xf0, 0xcc, 0xdd, 0x1a, 0xe9, 0x3b, 0x6f, 0x8d, 0x75, 0x58,
	0x9b, 0xc1, 0xf0, 0x6b, 0x7c, 0x0a, 0x5f, 0x86, 0x17, 0x49, 0xf3, 0x9a, 0x48, 0xef, 0xc2, 0x0a,
	0x40, 0x7e, 0xe8, 0x42, 0x09, 0x48, 0x48, 0x84, 0x24, 0x98, 0x1d, 0x78, 0x1e, 0x3a, 0x3c, 0x27,
	0x72, 0xe0, 0x72, 0x7b, 0xfa, 0x58, 0xa7, 0xaf, 0x61, 0xa5, 0x8d, 0xa9, 0xeb, 0x8d, 0xa0, 0xf8,
	0x15, 0x64, 0xdc, 0x7b, 0x05, 0x99, 0xdf, 0x07, 0x7d, 0xed, 0xa5, 0xd3, 0xb6, 0xb9, 0x3d, 0x12,
	0xe8, 0x05, 0x14, 0x75, 0xb7, 0x12, 0x2a, 0x31, 0xbf, 0xb2, 0x87, 0x7e, 0x15, 0x0b, 0x4a, 0xda,
	0xf2, 0x85, 0x66, 0x1d, 0x3e, 0x0f, 0x4d, 0x5b, 0x34, 0xac, 0xfd, 0x37, 0x90, 0x19, 0x2b, 0x4f,
	0x3e, 0x84, 0xf5, 0x6a, 0x3c, 0x84, 0xe5, 0x6f, 0x30, 0xbb, 0xb0, 0x19, 0xea, 0xea, 0xf6, 0xd0,
	0xa6, 0x0e, 0x3e, 0xed, 0x3d, 0x94, 0x8a, 0x68, 0x55, 0x13, 0x77, 0x56, 0xb5, 0x01, 0xcf, 0x96,
	0xc6, 0x98, 0x1d, 0xa3, 0xb4, 0xc3, 0xc8, 0x8c, 0xa9, 0xb5, 0xd8, 0xe3, 0xc4, 0xd2, 0x5a, 0xf3,
	0x6f, 0x03, 0x90, 0xd6, 0xf8, 0x2d, 0xd2, 0xbc, 0xc2, 0xf4, 0x13, 0xed, 0xf3, 0x76, 0x94, 0xf5,
	0x8e, 0xb4, 0xb9, 0xc4, 0xae, 0xd7, 0xdb, 0x1a, 0xd2, 0xc7, 0x34, 0xa0, 0x05, 0x5b, 0xa1, 0xc7,
	0x1f, 0x09, 0xb5, 0x87, 0xe4, 0xf7, 0xc7, 0xfa, 0xfc, 0xd7, 0x80, 0xa7, 0xf1, 0xa3, 0x82, 0xa7,
	0x0f, 0xf6, 0x37, 0x62, 0xae, 0x7e, 0x70, 0x16, 0x2c, 0xf5, 0x3f, 0xc7, 0x4d, 0xf2, 0x2e, 0x6e,
	0x82, 0xe7, 0x5d, 0x6a, 0xd9, 0xf3, 0x2e, 0x5e, 0xa9, 0xf4, 0xbd, 0x95, 0x0a, 0x30, 0x65, 0x22,
	0x98, 0x48, 0xb4, 0x17, 0xcf, 0xbc, 0x83, 0xd8, 0xc3, 0xfc, 0x90, 0xd1, 0x1e, 0xe1, 0x23, 0xec,
	0xde, 0x3f, 0x94, 0x51, 0x19, 0x12, 0x92, 0x2d, 0x34, 0x7a, 0x42, 0xb2, 0x59, 0xa8, 0x64, 0x18,
	0xea, 0xe5, 0x3b, 0x78, 0xb2, 0xe4, 0x99, 0x8d, 0x0a, 0xb0, 0xda, 0x68, 0xb6, 0x4f, 0x3b, 0xad,
	0xb3, 0x66, 0xa3, 0xf4, 0x19, 0xca, 0xc1, 0x4a, 0xf3, 0x97, 0xd6, 0x59, 0xeb, 0xe4, 0xa8, 0x64,
	0xa0, 0x22, 0xc0, 0xe1, 0xf1, 0xc1, 0xdb, 0xb7, 0xcd, 0x93, 0xa3, 0x66, 0xa3, 0x94, 0x40, 0x00,
	0x19, 0x4f, 0xd9, 0x6c, 0x94, 0x92, 0xdd, 0x8c, 0x7a, 0xf1, 0xef, 0xff, 0x17, 0x00, 0x00, 0xff,
	0xff, 0x6c, 0xa4, 0x24, 0xe6, 0xa0, 0x0c, 0x00, 0x00,
}

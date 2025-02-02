// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: penumbra/core/transparent_proofs/v1alpha1/transparent_proofs.proto

package transparent_proofsv1alpha1

import (
	v1alpha1 "github.com/penumbra-zone/penumbra/proto/go/gen/penumbra/core/crypto/v1alpha1"
	v1alpha11 "github.com/penumbra-zone/penumbra/proto/go/gen/penumbra/core/dex/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A Penumbra transparent Spend Proof.
type SpendProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Auxiliary inputs
	StateCommitmentProof *v1alpha1.StateCommitmentProof `protobuf:"bytes,1,opt,name=state_commitment_proof,json=stateCommitmentProof,proto3" json:"state_commitment_proof,omitempty"`
	// @exclude
	// From the note being spent
	Note                *v1alpha1.Note `protobuf:"bytes,2,opt,name=note,proto3" json:"note,omitempty"`
	VBlinding           []byte         `protobuf:"bytes,6,opt,name=v_blinding,json=vBlinding,proto3" json:"v_blinding,omitempty"`
	SpendAuthRandomizer []byte         `protobuf:"bytes,9,opt,name=spend_auth_randomizer,json=spendAuthRandomizer,proto3" json:"spend_auth_randomizer,omitempty"`
	Ak                  []byte         `protobuf:"bytes,10,opt,name=ak,proto3" json:"ak,omitempty"`
	Nk                  []byte         `protobuf:"bytes,11,opt,name=nk,proto3" json:"nk,omitempty"`
}

func (x *SpendProof) Reset() {
	*x = SpendProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpendProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpendProof) ProtoMessage() {}

func (x *SpendProof) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpendProof.ProtoReflect.Descriptor instead.
func (*SpendProof) Descriptor() ([]byte, []int) {
	return file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_rawDescGZIP(), []int{0}
}

func (x *SpendProof) GetStateCommitmentProof() *v1alpha1.StateCommitmentProof {
	if x != nil {
		return x.StateCommitmentProof
	}
	return nil
}

func (x *SpendProof) GetNote() *v1alpha1.Note {
	if x != nil {
		return x.Note
	}
	return nil
}

func (x *SpendProof) GetVBlinding() []byte {
	if x != nil {
		return x.VBlinding
	}
	return nil
}

func (x *SpendProof) GetSpendAuthRandomizer() []byte {
	if x != nil {
		return x.SpendAuthRandomizer
	}
	return nil
}

func (x *SpendProof) GetAk() []byte {
	if x != nil {
		return x.Ak
	}
	return nil
}

func (x *SpendProof) GetNk() []byte {
	if x != nil {
		return x.Nk
	}
	return nil
}

// A Penumbra transparent output proof.
type OutputProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Auxiliary inputs
	Note      *v1alpha1.Note `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	VBlinding []byte         `protobuf:"bytes,5,opt,name=v_blinding,json=vBlinding,proto3" json:"v_blinding,omitempty"`
}

func (x *OutputProof) Reset() {
	*x = OutputProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutputProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputProof) ProtoMessage() {}

func (x *OutputProof) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutputProof.ProtoReflect.Descriptor instead.
func (*OutputProof) Descriptor() ([]byte, []int) {
	return file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_rawDescGZIP(), []int{1}
}

func (x *OutputProof) GetNote() *v1alpha1.Note {
	if x != nil {
		return x.Note
	}
	return nil
}

func (x *OutputProof) GetVBlinding() []byte {
	if x != nil {
		return x.VBlinding
	}
	return nil
}

// A Penumbra transparent SwapClaimProof.
type SwapClaimProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The swap being claimed
	SwapPlaintext *v1alpha11.SwapPlaintext `protobuf:"bytes,1,opt,name=swap_plaintext,json=swapPlaintext,proto3" json:"swap_plaintext,omitempty"`
	// Inclusion proof for the swap commitment
	SwapCommitmentProof *v1alpha1.StateCommitmentProof `protobuf:"bytes,4,opt,name=swap_commitment_proof,json=swapCommitmentProof,proto3" json:"swap_commitment_proof,omitempty"`
	// The nullifier key used to derive the swap nullifier
	Nk []byte `protobuf:"bytes,6,opt,name=nk,proto3" json:"nk,omitempty"`
	// @exclude
	// Describes output amounts
	Lambda_1I uint64 `protobuf:"varint,20,opt,name=lambda_1_i,json=lambda1I,proto3" json:"lambda_1_i,omitempty"`
	Lambda_2I uint64 `protobuf:"varint,21,opt,name=lambda_2_i,json=lambda2I,proto3" json:"lambda_2_i,omitempty"`
}

func (x *SwapClaimProof) Reset() {
	*x = SwapClaimProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwapClaimProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwapClaimProof) ProtoMessage() {}

func (x *SwapClaimProof) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwapClaimProof.ProtoReflect.Descriptor instead.
func (*SwapClaimProof) Descriptor() ([]byte, []int) {
	return file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_rawDescGZIP(), []int{2}
}

func (x *SwapClaimProof) GetSwapPlaintext() *v1alpha11.SwapPlaintext {
	if x != nil {
		return x.SwapPlaintext
	}
	return nil
}

func (x *SwapClaimProof) GetSwapCommitmentProof() *v1alpha1.StateCommitmentProof {
	if x != nil {
		return x.SwapCommitmentProof
	}
	return nil
}

func (x *SwapClaimProof) GetNk() []byte {
	if x != nil {
		return x.Nk
	}
	return nil
}

func (x *SwapClaimProof) GetLambda_1I() uint64 {
	if x != nil {
		return x.Lambda_1I
	}
	return 0
}

func (x *SwapClaimProof) GetLambda_2I() uint64 {
	if x != nil {
		return x.Lambda_2I
	}
	return 0
}

// A Penumbra transparent SwapProof.
type SwapProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SwapPlaintext *v1alpha11.SwapPlaintext `protobuf:"bytes,1,opt,name=swap_plaintext,json=swapPlaintext,proto3" json:"swap_plaintext,omitempty"`
	// The blinding factor used for the Swap action's fee commitment.
	FeeBlinding []byte `protobuf:"bytes,2,opt,name=fee_blinding,json=feeBlinding,proto3" json:"fee_blinding,omitempty"`
}

func (x *SwapProof) Reset() {
	*x = SwapProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwapProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwapProof) ProtoMessage() {}

func (x *SwapProof) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwapProof.ProtoReflect.Descriptor instead.
func (*SwapProof) Descriptor() ([]byte, []int) {
	return file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_rawDescGZIP(), []int{3}
}

func (x *SwapProof) GetSwapPlaintext() *v1alpha11.SwapPlaintext {
	if x != nil {
		return x.SwapPlaintext
	}
	return nil
}

func (x *SwapProof) GetFeeBlinding() []byte {
	if x != nil {
		return x.FeeBlinding
	}
	return nil
}

type UndelegateClaimProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnbondingAmount *v1alpha1.Amount `protobuf:"bytes,1,opt,name=unbonding_amount,json=unbondingAmount,proto3" json:"unbonding_amount,omitempty"`
	BalanceBlinding []byte           `protobuf:"bytes,2,opt,name=balance_blinding,json=balanceBlinding,proto3" json:"balance_blinding,omitempty"`
}

func (x *UndelegateClaimProof) Reset() {
	*x = UndelegateClaimProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UndelegateClaimProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UndelegateClaimProof) ProtoMessage() {}

func (x *UndelegateClaimProof) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UndelegateClaimProof.ProtoReflect.Descriptor instead.
func (*UndelegateClaimProof) Descriptor() ([]byte, []int) {
	return file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_rawDescGZIP(), []int{4}
}

func (x *UndelegateClaimProof) GetUnbondingAmount() *v1alpha1.Amount {
	if x != nil {
		return x.UnbondingAmount
	}
	return nil
}

func (x *UndelegateClaimProof) GetBalanceBlinding() []byte {
	if x != nil {
		return x.BalanceBlinding
	}
	return nil
}

var File_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto protoreflect.FileDescriptor

var file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_rawDesc = []byte{
	0x0a, 0x42, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x6f, 0x66, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a,
	0x2a, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x70, 0x65, 0x6e,
	0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x64, 0x65, 0x78, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa3, 0x02, 0x0a, 0x0a, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x12, 0x69, 0x0a, 0x16, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x33, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x14, 0x73, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x37, 0x0a, 0x04, 0x6e,
	0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x65, 0x6e, 0x75,
	0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x04,
	0x6e, 0x6f, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x5f, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x76, 0x42, 0x6c, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x13, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x41, 0x75, 0x74, 0x68, 0x52, 0x61, 0x6e,
	0x64, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x6b, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x02, 0x61, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x6b, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x02, 0x6e, 0x6b, 0x22, 0x65, 0x0a, 0x0b, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x37, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x76, 0x5f, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x76, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x97,
	0x02, 0x0a, 0x0e, 0x53, 0x77, 0x61, 0x70, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x50, 0x72, 0x6f, 0x6f,
	0x66, 0x12, 0x50, 0x0a, 0x0e, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x65, 0x6e, 0x75,
	0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x77, 0x61, 0x70, 0x50, 0x6c, 0x61, 0x69, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x52, 0x0d, 0x73, 0x77, 0x61, 0x70, 0x50, 0x6c, 0x61, 0x69, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x67, 0x0a, 0x15, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x33, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x13, 0x73, 0x77, 0x61, 0x70, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x0e, 0x0a, 0x02,
	0x6e, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x6e, 0x6b, 0x12, 0x1c, 0x0a, 0x0a,
	0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x5f, 0x31, 0x5f, 0x69, 0x18, 0x14, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x31, 0x49, 0x12, 0x1c, 0x0a, 0x0a, 0x6c, 0x61,
	0x6d, 0x62, 0x64, 0x61, 0x5f, 0x32, 0x5f, 0x69, 0x18, 0x15, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x32, 0x49, 0x22, 0x80, 0x01, 0x0a, 0x09, 0x53, 0x77, 0x61,
	0x70, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x50, 0x0a, 0x0e, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x70,
	0x6c, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x64,
	0x65, 0x78, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x77, 0x61, 0x70,
	0x50, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0d, 0x73, 0x77, 0x61, 0x70, 0x50,
	0x6c, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x65, 0x65, 0x5f,
	0x62, 0x6c, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b,
	0x66, 0x65, 0x65, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x93, 0x01, 0x0a, 0x14,
	0x55, 0x6e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x50,
	0x72, 0x6f, 0x6f, 0x66, 0x12, 0x50, 0x0a, 0x10, 0x75, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0f, 0x75, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x42, 0xff, 0x02, 0x0a, 0x2d, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62,
	0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x42, 0x16, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x6f, 0x66, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x73, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62,
	0x72, 0x61, 0x2d, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x65,
	0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x43, 0x54, 0xaa, 0x02, 0x28, 0x50, 0x65, 0x6e, 0x75, 0x6d,
	0x62, 0x72, 0x61, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x73, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xca, 0x02, 0x28, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x5c, 0x43,
	0x6f, 0x72, 0x65, 0x5c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x6f, 0x66, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02,
	0x34, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x5c, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x73,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x2b, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61,
	0x3a, 0x3a, 0x43, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_rawDescOnce sync.Once
	file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_rawDescData = file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_rawDesc
)

func file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_rawDescGZIP() []byte {
	file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_rawDescOnce.Do(func() {
		file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_rawDescData = protoimpl.X.CompressGZIP(file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_rawDescData)
	})
	return file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_rawDescData
}

var file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_goTypes = []interface{}{
	(*SpendProof)(nil),                    // 0: penumbra.core.transparent_proofs.v1alpha1.SpendProof
	(*OutputProof)(nil),                   // 1: penumbra.core.transparent_proofs.v1alpha1.OutputProof
	(*SwapClaimProof)(nil),                // 2: penumbra.core.transparent_proofs.v1alpha1.SwapClaimProof
	(*SwapProof)(nil),                     // 3: penumbra.core.transparent_proofs.v1alpha1.SwapProof
	(*UndelegateClaimProof)(nil),          // 4: penumbra.core.transparent_proofs.v1alpha1.UndelegateClaimProof
	(*v1alpha1.StateCommitmentProof)(nil), // 5: penumbra.core.crypto.v1alpha1.StateCommitmentProof
	(*v1alpha1.Note)(nil),                 // 6: penumbra.core.crypto.v1alpha1.Note
	(*v1alpha11.SwapPlaintext)(nil),       // 7: penumbra.core.dex.v1alpha1.SwapPlaintext
	(*v1alpha1.Amount)(nil),               // 8: penumbra.core.crypto.v1alpha1.Amount
}
var file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_depIdxs = []int32{
	5, // 0: penumbra.core.transparent_proofs.v1alpha1.SpendProof.state_commitment_proof:type_name -> penumbra.core.crypto.v1alpha1.StateCommitmentProof
	6, // 1: penumbra.core.transparent_proofs.v1alpha1.SpendProof.note:type_name -> penumbra.core.crypto.v1alpha1.Note
	6, // 2: penumbra.core.transparent_proofs.v1alpha1.OutputProof.note:type_name -> penumbra.core.crypto.v1alpha1.Note
	7, // 3: penumbra.core.transparent_proofs.v1alpha1.SwapClaimProof.swap_plaintext:type_name -> penumbra.core.dex.v1alpha1.SwapPlaintext
	5, // 4: penumbra.core.transparent_proofs.v1alpha1.SwapClaimProof.swap_commitment_proof:type_name -> penumbra.core.crypto.v1alpha1.StateCommitmentProof
	7, // 5: penumbra.core.transparent_proofs.v1alpha1.SwapProof.swap_plaintext:type_name -> penumbra.core.dex.v1alpha1.SwapPlaintext
	8, // 6: penumbra.core.transparent_proofs.v1alpha1.UndelegateClaimProof.unbonding_amount:type_name -> penumbra.core.crypto.v1alpha1.Amount
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_init() }
func file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_init() {
	if File_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpendProof); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutputProof); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwapClaimProof); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwapProof); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UndelegateClaimProof); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_goTypes,
		DependencyIndexes: file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_depIdxs,
		MessageInfos:      file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_msgTypes,
	}.Build()
	File_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto = out.File
	file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_rawDesc = nil
	file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_goTypes = nil
	file_penumbra_core_transparent_proofs_v1alpha1_transparent_proofs_proto_depIdxs = nil
}

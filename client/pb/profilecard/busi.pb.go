// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/profilecard/busi.proto

package profilecard

import (
	proto "github.com/RomiChan/protobuf/proto"
)

type BusiColor struct {
	R proto.Option[int32] `protobuf:"varint,1,opt"`
	G proto.Option[int32] `protobuf:"varint,2,opt"`
	B proto.Option[int32] `protobuf:"varint,3,opt"`
}

type BusiComm struct {
	Ver            proto.Option[int32]  `protobuf:"varint,1,opt"`
	Seq            proto.Option[int32]  `protobuf:"varint,2,opt"`
	Fromuin        proto.Option[int64]  `protobuf:"varint,3,opt"`
	Touin          proto.Option[int64]  `protobuf:"varint,4,opt"`
	Service        proto.Option[int32]  `protobuf:"varint,5,opt"`
	SessionType    proto.Option[int32]  `protobuf:"varint,6,opt"`
	SessionKey     []byte               `protobuf:"bytes,7,opt"`
	ClientIp       proto.Option[int32]  `protobuf:"varint,8,opt"`
	Display        *BusiUi              `protobuf:"bytes,9,opt"`
	Result         proto.Option[int32]  `protobuf:"varint,10,opt"`
	ErrMsg         proto.Option[string] `protobuf:"bytes,11,opt"`
	Platform       proto.Option[int32]  `protobuf:"varint,12,opt"`
	Qqver          proto.Option[string] `protobuf:"bytes,13,opt"`
	Build          proto.Option[int32]  `protobuf:"varint,14,opt"`
	MsgLoginSig    *BusiLoginSig        `protobuf:"bytes,15,opt"`
	Version        proto.Option[int32]  `protobuf:"varint,17,opt"`
	MsgUinInfo     *BusiUinInfo         `protobuf:"bytes,18,opt"`
	MsgRichDisplay *BusiRichUi          `protobuf:"bytes,19,opt"`
}

type BusiCommonReq struct {
	ServiceCmd proto.Option[string] `protobuf:"bytes,1,opt"`
	VcReq      *BusiVisitorCountReq `protobuf:"bytes,2,opt"`
	HrReq      *BusiHideRecordsReq  `protobuf:"bytes,3,opt"`
}

type BusiDetailRecord struct {
	Fuin     proto.Option[int32] `protobuf:"varint,1,opt"`
	Source   proto.Option[int32] `protobuf:"varint,2,opt"`
	Vtime    proto.Option[int32] `protobuf:"varint,3,opt"`
	Mod      proto.Option[int32] `protobuf:"varint,4,opt"`
	HideFlag proto.Option[int32] `protobuf:"varint,5,opt"`
}

type BusiHideRecordsReq struct {
	Huin    proto.Option[int32] `protobuf:"varint,1,opt"`
	Fuin    proto.Option[int32] `protobuf:"varint,2,opt"`
	Records []*BusiDetailRecord `protobuf:"bytes,3,rep"`
}

type BusiLabel struct {
	Name        []byte              `protobuf:"bytes,1,opt"`
	EnumType    proto.Option[int32] `protobuf:"varint,2,opt"`
	TextColor   *BusiColor          `protobuf:"bytes,3,opt"`
	EdgingColor *BusiColor          `protobuf:"bytes,4,opt"`
	LabelAttr   proto.Option[int32] `protobuf:"varint,5,opt"`
	LabelType   proto.Option[int32] `protobuf:"varint,6,opt"`
}

type BusiLoginSig struct {
	Type  proto.Option[int32] `protobuf:"varint,1,opt"`
	Sig   []byte              `protobuf:"bytes,2,opt"`
	Appid proto.Option[int32] `protobuf:"varint,3,opt"`
}

type BusiRichUi struct {
	Name       proto.Option[string] `protobuf:"bytes,1,opt"`
	ServiceUrl proto.Option[string] `protobuf:"bytes,2,opt"` //repeated UiInfo uiList = 3;
}

type BusiUi struct {
	Url     proto.Option[string] `protobuf:"bytes,1,opt"`
	Title   proto.Option[string] `protobuf:"bytes,2,opt"`
	Content proto.Option[string] `protobuf:"bytes,3,opt"`
	JumpUrl proto.Option[string] `protobuf:"bytes,4,opt"`
}

type BusiUinInfo struct {
	Int64Longitude proto.Option[int64] `protobuf:"varint,1,opt"`
	Int64Latitude  proto.Option[int64] `protobuf:"varint,2,opt"`
}

type BusiVisitorCountReq struct {
	Requireuin proto.Option[int32] `protobuf:"varint,1,opt"`
	Operuin    proto.Option[int32] `protobuf:"varint,2,opt"`
	Mod        proto.Option[int32] `protobuf:"varint,3,opt"`
	ReportFlag proto.Option[int32] `protobuf:"varint,4,opt"`
}

type BusiVisitorCountRsp struct {
	Requireuin proto.Option[int32] `protobuf:"varint,1,opt"`
	TotalLike  proto.Option[int32] `protobuf:"varint,2,opt"`
	TotalView  proto.Option[int32] `protobuf:"varint,3,opt"`
	HotValue   proto.Option[int32] `protobuf:"varint,4,opt"`
	RedValue   proto.Option[int32] `protobuf:"varint,5,opt"`
	HotDiff    proto.Option[int32] `protobuf:"varint,6,opt"`
}

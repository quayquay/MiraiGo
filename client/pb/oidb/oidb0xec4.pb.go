// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/oidb/oidb0xec4.proto

package oidb

import (
	proto "github.com/RomiChan/protobuf/proto"
)

type Comment struct {
	Id       proto.Option[string] `protobuf:"bytes,1,opt"`
	Comment  proto.Option[string] `protobuf:"bytes,2,opt"`
	Time     proto.Option[uint64] `protobuf:"varint,3,opt"`
	FromUin  proto.Option[uint64] `protobuf:"varint,4,opt"`
	ToUin    proto.Option[uint64] `protobuf:"varint,5,opt"`
	ReplyId  proto.Option[string] `protobuf:"bytes,6,opt"`
	FromNick proto.Option[string] `protobuf:"bytes,7,opt"`
}

type Praise struct {
	FromUin  proto.Option[uint64] `protobuf:"varint,1,opt"`
	ToUin    proto.Option[uint64] `protobuf:"varint,2,opt"`
	Time     proto.Option[uint64] `protobuf:"varint,3,opt"`
	FromNick proto.Option[string] `protobuf:"bytes,4,opt"`
}

type Quest struct {
	Id          proto.Option[string] `protobuf:"bytes,1,opt"`
	Quest       proto.Option[string] `protobuf:"bytes,2,opt"`
	QuestUin    proto.Option[uint64] `protobuf:"varint,3,opt"`
	Time        proto.Option[uint64] `protobuf:"varint,4,opt"`
	Ans         proto.Option[string] `protobuf:"bytes,5,opt"`
	AnsTime     proto.Option[uint64] `protobuf:"varint,6,opt"`
	Comment     []*Comment           `protobuf:"bytes,7,rep"`
	Praise      []*Praise            `protobuf:"bytes,8,rep"`
	PraiseNum   proto.Option[uint64] `protobuf:"varint,9,opt"`
	LikeKey     proto.Option[string] `protobuf:"bytes,10,opt"`
	SystemId    proto.Option[uint64] `protobuf:"varint,11,opt"`
	CommentNum  proto.Option[uint64] `protobuf:"varint,12,opt"`
	ShowType    proto.Option[uint64] `protobuf:"varint,13,opt"`
	ShowTimes   proto.Option[uint64] `protobuf:"varint,14,opt"`
	BeenPraised proto.Option[uint64] `protobuf:"varint,15,opt"`
	QuestRead   proto.Option[bool]   `protobuf:"varint,16,opt"`
	AnsShowType proto.Option[uint64] `protobuf:"varint,17,opt"`
}

type DEC4ReqBody struct {
	Uin        proto.Option[uint64] `protobuf:"varint,1,opt"`
	QuestNum   proto.Option[uint64] `protobuf:"varint,2,opt"`
	CommentNum proto.Option[uint64] `protobuf:"varint,3,opt"`
	Cookie     []byte               `protobuf:"bytes,4,opt"`
	FetchType  proto.Option[uint32] `protobuf:"varint,5,opt"`
}

type DEC4RspBody struct {
	Quest            []*Quest             `protobuf:"bytes,1,rep"`
	IsFetchOver      proto.Option[bool]   `protobuf:"varint,2,opt"`
	TotalQuestNum    proto.Option[uint32] `protobuf:"varint,3,opt"`
	Cookie           []byte               `protobuf:"bytes,4,opt"`
	Ret              proto.Option[uint32] `protobuf:"varint,5,opt"`
	AnsweredQuestNum proto.Option[uint32] `protobuf:"varint,6,opt"`
}

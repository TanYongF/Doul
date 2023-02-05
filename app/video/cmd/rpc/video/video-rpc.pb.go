// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: video-rpc.proto

package video

import (
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

type VideoPO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author        *UserPO `protobuf:"bytes,1,opt,name=Author,proto3" json:"Author,omitempty"`                // 视频作者信息
	CommentCount  int64   `protobuf:"varint,2,opt,name=CommentCount,proto3" json:"CommentCount,omitempty"`   // 视频的评论总数
	CoverURL      string  `protobuf:"bytes,3,opt,name=CoverURL,proto3" json:"CoverURL,omitempty"`            // 视频封面地址
	FavoriteCount int64   `protobuf:"varint,4,opt,name=FavoriteCount,proto3" json:"FavoriteCount,omitempty"` // 视频的点赞总数
	VideoId       int64   `protobuf:"varint,5,opt,name=VideoId,proto3" json:"VideoId,omitempty"`             // 视频唯一标识
	IsFavorite    bool    `protobuf:"varint,6,opt,name=IsFavorite,proto3" json:"IsFavorite,omitempty"`       // true-已点赞，false-未点赞
	PlayURL       string  `protobuf:"bytes,7,opt,name=PlayURL,proto3" json:"PlayURL,omitempty"`              // 视频播放地址
	Title         string  `protobuf:"bytes,8,opt,name=Title,proto3" json:"Title,omitempty"`                  // 视频标题
}

func (x *VideoPO) Reset() {
	*x = VideoPO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoPO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoPO) ProtoMessage() {}

func (x *VideoPO) ProtoReflect() protoreflect.Message {
	mi := &file_video_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoPO.ProtoReflect.Descriptor instead.
func (*VideoPO) Descriptor() ([]byte, []int) {
	return file_video_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *VideoPO) GetAuthor() *UserPO {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *VideoPO) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *VideoPO) GetCoverURL() string {
	if x != nil {
		return x.CoverURL
	}
	return ""
}

func (x *VideoPO) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *VideoPO) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *VideoPO) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

func (x *VideoPO) GetPlayURL() string {
	if x != nil {
		return x.PlayURL
	}
	return ""
}

func (x *VideoPO) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type UserPO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowCount   int64  `protobuf:"varint,1,opt,name=FollowCount,proto3" json:"FollowCount,omitempty"`     // 关注总数
	FollowerCount int64  `protobuf:"varint,2,opt,name=FollowerCount,proto3" json:"FollowerCount,omitempty"` // 粉丝总数
	UserId        int64  `protobuf:"varint,3,opt,name=UserId,proto3" json:"UserId,omitempty"`               // 用户id
	IsFollow      bool   `protobuf:"varint,4,opt,name=IsFollow,proto3" json:"IsFollow,omitempty"`           // true-已关注，false-未关注
	Name          string `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`                    // 用户名称
}

func (x *UserPO) Reset() {
	*x = UserPO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPO) ProtoMessage() {}

func (x *UserPO) ProtoReflect() protoreflect.Message {
	mi := &file_video_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPO.ProtoReflect.Descriptor instead.
func (*UserPO) Descriptor() ([]byte, []int) {
	return file_video_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *UserPO) GetFollowCount() int64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *UserPO) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *UserPO) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserPO) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

func (x *UserPO) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FeedReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatestTime string `protobuf:"bytes,1,opt,name=Latest_time,json=LatestTime,proto3" json:"Latest_time,omitempty"`
}

func (x *FeedReq) Reset() {
	*x = FeedReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedReq) ProtoMessage() {}

func (x *FeedReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedReq.ProtoReflect.Descriptor instead.
func (*FeedReq) Descriptor() ([]byte, []int) {
	return file_video_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *FeedReq) GetLatestTime() string {
	if x != nil {
		return x.LatestTime
	}
	return ""
}

type FeedResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextTime  string     `protobuf:"bytes,1,opt,name=NextTime,proto3" json:"NextTime,omitempty"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	VideoList []*VideoPO `protobuf:"bytes,2,rep,name=VideoList,proto3" json:"VideoList,omitempty"` // 视频列表
}

func (x *FeedResp) Reset() {
	*x = FeedResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedResp) ProtoMessage() {}

func (x *FeedResp) ProtoReflect() protoreflect.Message {
	mi := &file_video_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedResp.ProtoReflect.Descriptor instead.
func (*FeedResp) Descriptor() ([]byte, []int) {
	return file_video_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *FeedResp) GetNextTime() string {
	if x != nil {
		return x.NextTime
	}
	return ""
}

func (x *FeedResp) GetVideoList() []*VideoPO {
	if x != nil {
		return x.VideoList
	}
	return nil
}

type PublishListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryId int64 `protobuf:"varint,1,opt,name=QueryId,proto3" json:"QueryId,omitempty"`
}

func (x *PublishListReq) Reset() {
	*x = PublishListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishListReq) ProtoMessage() {}

func (x *PublishListReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishListReq.ProtoReflect.Descriptor instead.
func (*PublishListReq) Descriptor() ([]byte, []int) {
	return file_video_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *PublishListReq) GetQueryId() int64 {
	if x != nil {
		return x.QueryId
	}
	return 0
}

type PublishListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoList []*VideoPO `protobuf:"bytes,1,rep,name=VideoList,proto3" json:"VideoList,omitempty"` // 视频列表
}

func (x *PublishListResp) Reset() {
	*x = PublishListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishListResp) ProtoMessage() {}

func (x *PublishListResp) ProtoReflect() protoreflect.Message {
	mi := &file_video_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishListResp.ProtoReflect.Descriptor instead.
func (*PublishListResp) Descriptor() ([]byte, []int) {
	return file_video_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *PublishListResp) GetVideoList() []*VideoPO {
	if x != nil {
		return x.VideoList
	}
	return nil
}

type FavoriteListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryId int64 `protobuf:"varint,1,opt,name=QueryId,proto3" json:"QueryId,omitempty"`
}

func (x *FavoriteListReq) Reset() {
	*x = FavoriteListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteListReq) ProtoMessage() {}

func (x *FavoriteListReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteListReq.ProtoReflect.Descriptor instead.
func (*FavoriteListReq) Descriptor() ([]byte, []int) {
	return file_video_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *FavoriteListReq) GetQueryId() int64 {
	if x != nil {
		return x.QueryId
	}
	return 0
}

type FavoriteListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoList []*VideoPO `protobuf:"bytes,1,rep,name=VideoList,proto3" json:"VideoList,omitempty"` // 视频列表
}

func (x *FavoriteListResp) Reset() {
	*x = FavoriteListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_rpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteListResp) ProtoMessage() {}

func (x *FavoriteListResp) ProtoReflect() protoreflect.Message {
	mi := &file_video_rpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteListResp.ProtoReflect.Descriptor instead.
func (*FavoriteListResp) Descriptor() ([]byte, []int) {
	return file_video_rpc_proto_rawDescGZIP(), []int{7}
}

func (x *FavoriteListResp) GetVideoList() []*VideoPO {
	if x != nil {
		return x.VideoList
	}
	return nil
}

var File_video_rpc_proto protoreflect.FileDescriptor

var file_video_rpc_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2d, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x80, 0x02, 0x0a, 0x07, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x50, 0x4f, 0x12, 0x25, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x4f, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x12, 0x24, 0x0a, 0x0d, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x49,
	0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x49, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50,
	0x6c, 0x61, 0x79, 0x55, 0x52, 0x4c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x6c,
	0x61, 0x79, 0x55, 0x52, 0x4c, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x98, 0x01, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x49, 0x73, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x07, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x54, 0x0a, 0x08, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a,
	0x0a, 0x08, 0x4e, 0x65, 0x78, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4e, 0x65, 0x78, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x50, 0x4f, 0x52, 0x09, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x2a, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x0f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2c, 0x0a, 0x09, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x50, 0x4f, 0x52, 0x09, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x2b, 0x0a, 0x0f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x49, 0x64, 0x22, 0x40, 0x0a, 0x10, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2c, 0x0a, 0x09, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x50, 0x4f, 0x52, 0x09, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x32, 0xaf, 0x01, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x27,
	0x0a, 0x04, 0x66, 0x65, 0x65, 0x64, 0x12, 0x0e, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x46,
	0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x46,
	0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3f, 0x0a, 0x0c, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x66, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_video_rpc_proto_rawDescOnce sync.Once
	file_video_rpc_proto_rawDescData = file_video_rpc_proto_rawDesc
)

func file_video_rpc_proto_rawDescGZIP() []byte {
	file_video_rpc_proto_rawDescOnce.Do(func() {
		file_video_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_rpc_proto_rawDescData)
	})
	return file_video_rpc_proto_rawDescData
}

var file_video_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_video_rpc_proto_goTypes = []interface{}{
	(*VideoPO)(nil),          // 0: video.VideoPO
	(*UserPO)(nil),           // 1: video.UserPO
	(*FeedReq)(nil),          // 2: video.FeedReq
	(*FeedResp)(nil),         // 3: video.FeedResp
	(*PublishListReq)(nil),   // 4: video.publishListReq
	(*PublishListResp)(nil),  // 5: video.publishListResp
	(*FavoriteListReq)(nil),  // 6: video.favoriteListReq
	(*FavoriteListResp)(nil), // 7: video.favoriteListResp
}
var file_video_rpc_proto_depIdxs = []int32{
	1, // 0: video.VideoPO.Author:type_name -> video.UserPO
	0, // 1: video.FeedResp.VideoList:type_name -> video.VideoPO
	0, // 2: video.publishListResp.VideoList:type_name -> video.VideoPO
	0, // 3: video.favoriteListResp.VideoList:type_name -> video.VideoPO
	2, // 4: video.Video.feed:input_type -> video.FeedReq
	4, // 5: video.Video.publishList:input_type -> video.publishListReq
	6, // 6: video.Video.favoriteList:input_type -> video.favoriteListReq
	3, // 7: video.Video.feed:output_type -> video.FeedResp
	5, // 8: video.Video.publishList:output_type -> video.publishListResp
	7, // 9: video.Video.favoriteList:output_type -> video.favoriteListResp
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_video_rpc_proto_init() }
func file_video_rpc_proto_init() {
	if File_video_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_video_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoPO); i {
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
		file_video_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPO); i {
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
		file_video_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedReq); i {
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
		file_video_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedResp); i {
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
		file_video_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishListReq); i {
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
		file_video_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishListResp); i {
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
		file_video_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteListReq); i {
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
		file_video_rpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteListResp); i {
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
			RawDescriptor: file_video_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_rpc_proto_goTypes,
		DependencyIndexes: file_video_rpc_proto_depIdxs,
		MessageInfos:      file_video_rpc_proto_msgTypes,
	}.Build()
	File_video_rpc_proto = out.File
	file_video_rpc_proto_rawDesc = nil
	file_video_rpc_proto_goTypes = nil
	file_video_rpc_proto_depIdxs = nil
}

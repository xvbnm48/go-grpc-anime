// Code generated by protoc-gen-go. DO NOT EDIT.
// source: animepb/anime.proto

package animepb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Anime struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Anime) Reset()         { *m = Anime{} }
func (m *Anime) String() string { return proto.CompactTextString(m) }
func (*Anime) ProtoMessage()    {}
func (*Anime) Descriptor() ([]byte, []int) {
	return fileDescriptor_b15d889cbf9faebf, []int{0}
}

func (m *Anime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Anime.Unmarshal(m, b)
}
func (m *Anime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Anime.Marshal(b, m, deterministic)
}
func (m *Anime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Anime.Merge(m, src)
}
func (m *Anime) XXX_Size() int {
	return xxx_messageInfo_Anime.Size(m)
}
func (m *Anime) XXX_DiscardUnknown() {
	xxx_messageInfo_Anime.DiscardUnknown(m)
}

var xxx_messageInfo_Anime proto.InternalMessageInfo

func (m *Anime) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Anime) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Anime) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CreateAnimeRequest struct {
	Anime                *Anime   `protobuf:"bytes,1,opt,name=anime,proto3" json:"anime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAnimeRequest) Reset()         { *m = CreateAnimeRequest{} }
func (m *CreateAnimeRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAnimeRequest) ProtoMessage()    {}
func (*CreateAnimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b15d889cbf9faebf, []int{1}
}

func (m *CreateAnimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAnimeRequest.Unmarshal(m, b)
}
func (m *CreateAnimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAnimeRequest.Marshal(b, m, deterministic)
}
func (m *CreateAnimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAnimeRequest.Merge(m, src)
}
func (m *CreateAnimeRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAnimeRequest.Size(m)
}
func (m *CreateAnimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAnimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAnimeRequest proto.InternalMessageInfo

func (m *CreateAnimeRequest) GetAnime() *Anime {
	if m != nil {
		return m.Anime
	}
	return nil
}

type CreateAnimeResponse struct {
	Anime                *Anime   `protobuf:"bytes,1,opt,name=anime,proto3" json:"anime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAnimeResponse) Reset()         { *m = CreateAnimeResponse{} }
func (m *CreateAnimeResponse) String() string { return proto.CompactTextString(m) }
func (*CreateAnimeResponse) ProtoMessage()    {}
func (*CreateAnimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b15d889cbf9faebf, []int{2}
}

func (m *CreateAnimeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAnimeResponse.Unmarshal(m, b)
}
func (m *CreateAnimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAnimeResponse.Marshal(b, m, deterministic)
}
func (m *CreateAnimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAnimeResponse.Merge(m, src)
}
func (m *CreateAnimeResponse) XXX_Size() int {
	return xxx_messageInfo_CreateAnimeResponse.Size(m)
}
func (m *CreateAnimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAnimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAnimeResponse proto.InternalMessageInfo

func (m *CreateAnimeResponse) GetAnime() *Anime {
	if m != nil {
		return m.Anime
	}
	return nil
}

type ReadAnimeRequest struct {
	AnimeId              string   `protobuf:"bytes,1,opt,name=anime_id,json=animeId,proto3" json:"anime_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAnimeRequest) Reset()         { *m = ReadAnimeRequest{} }
func (m *ReadAnimeRequest) String() string { return proto.CompactTextString(m) }
func (*ReadAnimeRequest) ProtoMessage()    {}
func (*ReadAnimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b15d889cbf9faebf, []int{3}
}

func (m *ReadAnimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAnimeRequest.Unmarshal(m, b)
}
func (m *ReadAnimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAnimeRequest.Marshal(b, m, deterministic)
}
func (m *ReadAnimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAnimeRequest.Merge(m, src)
}
func (m *ReadAnimeRequest) XXX_Size() int {
	return xxx_messageInfo_ReadAnimeRequest.Size(m)
}
func (m *ReadAnimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAnimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAnimeRequest proto.InternalMessageInfo

func (m *ReadAnimeRequest) GetAnimeId() string {
	if m != nil {
		return m.AnimeId
	}
	return ""
}

type ReadAnimeResponse struct {
	Anime                *Anime   `protobuf:"bytes,1,opt,name=anime,proto3" json:"anime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAnimeResponse) Reset()         { *m = ReadAnimeResponse{} }
func (m *ReadAnimeResponse) String() string { return proto.CompactTextString(m) }
func (*ReadAnimeResponse) ProtoMessage()    {}
func (*ReadAnimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b15d889cbf9faebf, []int{4}
}

func (m *ReadAnimeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAnimeResponse.Unmarshal(m, b)
}
func (m *ReadAnimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAnimeResponse.Marshal(b, m, deterministic)
}
func (m *ReadAnimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAnimeResponse.Merge(m, src)
}
func (m *ReadAnimeResponse) XXX_Size() int {
	return xxx_messageInfo_ReadAnimeResponse.Size(m)
}
func (m *ReadAnimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAnimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAnimeResponse proto.InternalMessageInfo

func (m *ReadAnimeResponse) GetAnime() *Anime {
	if m != nil {
		return m.Anime
	}
	return nil
}

type ListAnimeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAnimeRequest) Reset()         { *m = ListAnimeRequest{} }
func (m *ListAnimeRequest) String() string { return proto.CompactTextString(m) }
func (*ListAnimeRequest) ProtoMessage()    {}
func (*ListAnimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b15d889cbf9faebf, []int{5}
}

func (m *ListAnimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAnimeRequest.Unmarshal(m, b)
}
func (m *ListAnimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAnimeRequest.Marshal(b, m, deterministic)
}
func (m *ListAnimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAnimeRequest.Merge(m, src)
}
func (m *ListAnimeRequest) XXX_Size() int {
	return xxx_messageInfo_ListAnimeRequest.Size(m)
}
func (m *ListAnimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAnimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAnimeRequest proto.InternalMessageInfo

type ListAnimeResponse struct {
	Anime                *Anime   `protobuf:"bytes,1,opt,name=anime,proto3" json:"anime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAnimeResponse) Reset()         { *m = ListAnimeResponse{} }
func (m *ListAnimeResponse) String() string { return proto.CompactTextString(m) }
func (*ListAnimeResponse) ProtoMessage()    {}
func (*ListAnimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b15d889cbf9faebf, []int{6}
}

func (m *ListAnimeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAnimeResponse.Unmarshal(m, b)
}
func (m *ListAnimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAnimeResponse.Marshal(b, m, deterministic)
}
func (m *ListAnimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAnimeResponse.Merge(m, src)
}
func (m *ListAnimeResponse) XXX_Size() int {
	return xxx_messageInfo_ListAnimeResponse.Size(m)
}
func (m *ListAnimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAnimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAnimeResponse proto.InternalMessageInfo

func (m *ListAnimeResponse) GetAnime() *Anime {
	if m != nil {
		return m.Anime
	}
	return nil
}

type UpdateAnimeRequest struct {
	Anime                *Anime   `protobuf:"bytes,1,opt,name=anime,proto3" json:"anime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAnimeRequest) Reset()         { *m = UpdateAnimeRequest{} }
func (m *UpdateAnimeRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateAnimeRequest) ProtoMessage()    {}
func (*UpdateAnimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b15d889cbf9faebf, []int{7}
}

func (m *UpdateAnimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAnimeRequest.Unmarshal(m, b)
}
func (m *UpdateAnimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAnimeRequest.Marshal(b, m, deterministic)
}
func (m *UpdateAnimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAnimeRequest.Merge(m, src)
}
func (m *UpdateAnimeRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateAnimeRequest.Size(m)
}
func (m *UpdateAnimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAnimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAnimeRequest proto.InternalMessageInfo

func (m *UpdateAnimeRequest) GetAnime() *Anime {
	if m != nil {
		return m.Anime
	}
	return nil
}

type UpdateAnimeResponse struct {
	Anime                *Anime   `protobuf:"bytes,1,opt,name=anime,proto3" json:"anime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAnimeResponse) Reset()         { *m = UpdateAnimeResponse{} }
func (m *UpdateAnimeResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateAnimeResponse) ProtoMessage()    {}
func (*UpdateAnimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b15d889cbf9faebf, []int{8}
}

func (m *UpdateAnimeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAnimeResponse.Unmarshal(m, b)
}
func (m *UpdateAnimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAnimeResponse.Marshal(b, m, deterministic)
}
func (m *UpdateAnimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAnimeResponse.Merge(m, src)
}
func (m *UpdateAnimeResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateAnimeResponse.Size(m)
}
func (m *UpdateAnimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAnimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAnimeResponse proto.InternalMessageInfo

func (m *UpdateAnimeResponse) GetAnime() *Anime {
	if m != nil {
		return m.Anime
	}
	return nil
}

type DeleteAnimeRequest struct {
	AnimeId              string   `protobuf:"bytes,1,opt,name=anime_id,json=animeId,proto3" json:"anime_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAnimeRequest) Reset()         { *m = DeleteAnimeRequest{} }
func (m *DeleteAnimeRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAnimeRequest) ProtoMessage()    {}
func (*DeleteAnimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b15d889cbf9faebf, []int{9}
}

func (m *DeleteAnimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAnimeRequest.Unmarshal(m, b)
}
func (m *DeleteAnimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAnimeRequest.Marshal(b, m, deterministic)
}
func (m *DeleteAnimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAnimeRequest.Merge(m, src)
}
func (m *DeleteAnimeRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAnimeRequest.Size(m)
}
func (m *DeleteAnimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAnimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAnimeRequest proto.InternalMessageInfo

func (m *DeleteAnimeRequest) GetAnimeId() string {
	if m != nil {
		return m.AnimeId
	}
	return ""
}

type DeleteAnimeResponse struct {
	AnimeId              string   `protobuf:"bytes,1,opt,name=anime_id,json=animeId,proto3" json:"anime_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAnimeResponse) Reset()         { *m = DeleteAnimeResponse{} }
func (m *DeleteAnimeResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteAnimeResponse) ProtoMessage()    {}
func (*DeleteAnimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b15d889cbf9faebf, []int{10}
}

func (m *DeleteAnimeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAnimeResponse.Unmarshal(m, b)
}
func (m *DeleteAnimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAnimeResponse.Marshal(b, m, deterministic)
}
func (m *DeleteAnimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAnimeResponse.Merge(m, src)
}
func (m *DeleteAnimeResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteAnimeResponse.Size(m)
}
func (m *DeleteAnimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAnimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAnimeResponse proto.InternalMessageInfo

func (m *DeleteAnimeResponse) GetAnimeId() string {
	if m != nil {
		return m.AnimeId
	}
	return ""
}

func init() {
	proto.RegisterType((*Anime)(nil), "anime.Anime")
	proto.RegisterType((*CreateAnimeRequest)(nil), "anime.CreateAnimeRequest")
	proto.RegisterType((*CreateAnimeResponse)(nil), "anime.CreateAnimeResponse")
	proto.RegisterType((*ReadAnimeRequest)(nil), "anime.ReadAnimeRequest")
	proto.RegisterType((*ReadAnimeResponse)(nil), "anime.ReadAnimeResponse")
	proto.RegisterType((*ListAnimeRequest)(nil), "anime.ListAnimeRequest")
	proto.RegisterType((*ListAnimeResponse)(nil), "anime.ListAnimeResponse")
	proto.RegisterType((*UpdateAnimeRequest)(nil), "anime.UpdateAnimeRequest")
	proto.RegisterType((*UpdateAnimeResponse)(nil), "anime.UpdateAnimeResponse")
	proto.RegisterType((*DeleteAnimeRequest)(nil), "anime.DeleteAnimeRequest")
	proto.RegisterType((*DeleteAnimeResponse)(nil), "anime.DeleteAnimeResponse")
}

func init() {
	proto.RegisterFile("animepb/anime.proto", fileDescriptor_b15d889cbf9faebf)
}

var fileDescriptor_b15d889cbf9faebf = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0x56, 0x03, 0xa5, 0xf4, 0x52, 0x50, 0xb9, 0x0c, 0xa4, 0x99, 0xaa, 0x4c, 0x2c, 0xb4, 0xa8,
	0x0c, 0x94, 0xb1, 0x10, 0x09, 0x21, 0xc1, 0x12, 0x89, 0x85, 0x05, 0x39, 0xf8, 0x06, 0x4b, 0x34,
	0x09, 0xb1, 0x79, 0x23, 0x1e, 0x14, 0xc5, 0xb1, 0x22, 0xa7, 0x09, 0x88, 0x74, 0xaa, 0x7d, 0x77,
	0xdf, 0x4f, 0xbf, 0x8b, 0xc1, 0x63, 0xa9, 0xd8, 0x52, 0x9e, 0x2c, 0xf5, 0xef, 0x22, 0x2f, 0x32,
	0x95, 0xe1, 0x50, 0x5f, 0xc2, 0x67, 0x18, 0x6e, 0xca, 0x03, 0x9e, 0x82, 0x23, 0xb8, 0x3f, 0x98,
	0x0f, 0x2e, 0xc6, 0xb1, 0x23, 0x38, 0x22, 0x1c, 0xa6, 0x6c, 0x4b, 0xbe, 0xa3, 0x2b, 0xfa, 0x8c,
	0x73, 0x70, 0x39, 0xc9, 0xf7, 0x42, 0xe4, 0x4a, 0x64, 0xa9, 0x7f, 0xa0, 0x5b, 0x76, 0x29, 0x5c,
	0x03, 0xde, 0x17, 0xc4, 0x14, 0x69, 0xd2, 0x98, 0x3e, 0xbf, 0x48, 0x2a, 0x0c, 0xa1, 0x52, 0xd3,
	0xf4, 0xee, 0x6a, 0xb2, 0xa8, 0x8c, 0x54, 0x33, 0xc6, 0xc8, 0x2d, 0x78, 0x0d, 0xa4, 0xcc, 0xb3,
	0x54, 0xd2, 0xbf, 0xa0, 0x97, 0x30, 0x8d, 0x89, 0xf1, 0x86, 0xe4, 0x0c, 0x8e, 0x75, 0xf3, 0xad,
	0xfe, 0x53, 0x23, 0x7d, 0x7f, 0xe4, 0xe1, 0x0d, 0x9c, 0x59, 0xe3, 0x3d, 0x74, 0x10, 0xa6, 0x4f,
	0x42, 0x2a, 0x5b, 0xa7, 0x24, 0xb3, 0x6a, 0x3d, 0xc8, 0xd6, 0x80, 0x2f, 0x39, 0xdf, 0x33, 0xa9,
	0x06, 0xb2, 0x87, 0xe8, 0x12, 0x30, 0xa2, 0x0f, 0xda, 0x11, 0xfd, 0x23, 0xab, 0x2b, 0xf0, 0x1a,
	0x00, 0xa3, 0xf5, 0x3b, 0x62, 0xf5, 0xed, 0xc0, 0x49, 0xc4, 0x14, 0x4b, 0x98, 0xac, 0x40, 0x18,
	0x81, 0x6b, 0x6d, 0x16, 0x67, 0xc6, 0x58, 0xfb, 0x3b, 0x09, 0x82, 0xae, 0x96, 0x91, 0xdc, 0xc0,
	0xe4, 0x81, 0xaa, 0x9c, 0x4b, 0x7a, 0x3c, 0x37, 0xb3, 0xbb, 0x9b, 0x0f, 0xfc, 0x76, 0xc3, 0x50,
	0x44, 0xe0, 0x5a, 0xc1, 0xd5, 0x46, 0xda, 0x6b, 0xa8, 0x8d, 0x74, 0xe5, 0x1c, 0x81, 0x6b, 0x45,
	0x52, 0xb3, 0xb4, 0x73, 0xad, 0x59, 0x3a, 0x12, 0xbc, 0x1b, 0xbf, 0x8e, 0xcc, 0xab, 0x4c, 0x8e,
	0xf4, 0x83, 0xbc, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xac, 0xc0, 0xfd, 0xf3, 0xa7, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DatabaseAnimeClient is the client API for DatabaseAnime service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatabaseAnimeClient interface {
	// CreateAnime: Creates a new anime.
	CreateAnime(ctx context.Context, in *CreateAnimeRequest, opts ...grpc.CallOption) (*CreateAnimeResponse, error)
	// ReadAnime: Gets an anime by ID.
	GetAnimeData(ctx context.Context, in *ReadAnimeRequest, opts ...grpc.CallOption) (*ReadAnimeResponse, error)
	// UpdateAnime: Updates an anime.
	UpdateAnime(ctx context.Context, in *UpdateAnimeRequest, opts ...grpc.CallOption) (*UpdateAnimeResponse, error)
	// DeleteAnime: Deletes an anime.
	DeleteAnime(ctx context.Context, in *DeleteAnimeRequest, opts ...grpc.CallOption) (*DeleteAnimeResponse, error)
}

type databaseAnimeClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseAnimeClient(cc grpc.ClientConnInterface) DatabaseAnimeClient {
	return &databaseAnimeClient{cc}
}

func (c *databaseAnimeClient) CreateAnime(ctx context.Context, in *CreateAnimeRequest, opts ...grpc.CallOption) (*CreateAnimeResponse, error) {
	out := new(CreateAnimeResponse)
	err := c.cc.Invoke(ctx, "/anime.DatabaseAnime/CreateAnime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseAnimeClient) GetAnimeData(ctx context.Context, in *ReadAnimeRequest, opts ...grpc.CallOption) (*ReadAnimeResponse, error) {
	out := new(ReadAnimeResponse)
	err := c.cc.Invoke(ctx, "/anime.DatabaseAnime/GetAnimeData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseAnimeClient) UpdateAnime(ctx context.Context, in *UpdateAnimeRequest, opts ...grpc.CallOption) (*UpdateAnimeResponse, error) {
	out := new(UpdateAnimeResponse)
	err := c.cc.Invoke(ctx, "/anime.DatabaseAnime/UpdateAnime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseAnimeClient) DeleteAnime(ctx context.Context, in *DeleteAnimeRequest, opts ...grpc.CallOption) (*DeleteAnimeResponse, error) {
	out := new(DeleteAnimeResponse)
	err := c.cc.Invoke(ctx, "/anime.DatabaseAnime/DeleteAnime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseAnimeServer is the server API for DatabaseAnime service.
type DatabaseAnimeServer interface {
	// CreateAnime: Creates a new anime.
	CreateAnime(context.Context, *CreateAnimeRequest) (*CreateAnimeResponse, error)
	// ReadAnime: Gets an anime by ID.
	GetAnimeData(context.Context, *ReadAnimeRequest) (*ReadAnimeResponse, error)
	// UpdateAnime: Updates an anime.
	UpdateAnime(context.Context, *UpdateAnimeRequest) (*UpdateAnimeResponse, error)
	// DeleteAnime: Deletes an anime.
	DeleteAnime(context.Context, *DeleteAnimeRequest) (*DeleteAnimeResponse, error)
}

// UnimplementedDatabaseAnimeServer can be embedded to have forward compatible implementations.
type UnimplementedDatabaseAnimeServer struct {
}

func (*UnimplementedDatabaseAnimeServer) CreateAnime(ctx context.Context, req *CreateAnimeRequest) (*CreateAnimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAnime not implemented")
}
func (*UnimplementedDatabaseAnimeServer) GetAnimeData(ctx context.Context, req *ReadAnimeRequest) (*ReadAnimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnimeData not implemented")
}
func (*UnimplementedDatabaseAnimeServer) UpdateAnime(ctx context.Context, req *UpdateAnimeRequest) (*UpdateAnimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAnime not implemented")
}
func (*UnimplementedDatabaseAnimeServer) DeleteAnime(ctx context.Context, req *DeleteAnimeRequest) (*DeleteAnimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAnime not implemented")
}

func RegisterDatabaseAnimeServer(s *grpc.Server, srv DatabaseAnimeServer) {
	s.RegisterService(&_DatabaseAnime_serviceDesc, srv)
}

func _DatabaseAnime_CreateAnime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAnimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseAnimeServer).CreateAnime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anime.DatabaseAnime/CreateAnime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseAnimeServer).CreateAnime(ctx, req.(*CreateAnimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseAnime_GetAnimeData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAnimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseAnimeServer).GetAnimeData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anime.DatabaseAnime/GetAnimeData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseAnimeServer).GetAnimeData(ctx, req.(*ReadAnimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseAnime_UpdateAnime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAnimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseAnimeServer).UpdateAnime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anime.DatabaseAnime/UpdateAnime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseAnimeServer).UpdateAnime(ctx, req.(*UpdateAnimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseAnime_DeleteAnime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAnimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseAnimeServer).DeleteAnime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anime.DatabaseAnime/DeleteAnime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseAnimeServer).DeleteAnime(ctx, req.(*DeleteAnimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DatabaseAnime_serviceDesc = grpc.ServiceDesc{
	ServiceName: "anime.DatabaseAnime",
	HandlerType: (*DatabaseAnimeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAnime",
			Handler:    _DatabaseAnime_CreateAnime_Handler,
		},
		{
			MethodName: "GetAnimeData",
			Handler:    _DatabaseAnime_GetAnimeData_Handler,
		},
		{
			MethodName: "UpdateAnime",
			Handler:    _DatabaseAnime_UpdateAnime_Handler,
		},
		{
			MethodName: "DeleteAnime",
			Handler:    _DatabaseAnime_DeleteAnime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "animepb/anime.proto",
}

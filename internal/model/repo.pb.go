// Code generated by protoc-gen-go. DO NOT EDIT.
// source: repo.proto

package model // import "github.com/guilhermesteves/github-repos-sort-api/internal/model"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Repo struct {
	Id                   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GithubRepo           *GithubRepo `protobuf:"bytes,2,opt,name=githubRepo,proto3" json:"githubRepo,omitempty"`
	Likes                int64       `protobuf:"varint,3,opt,name=likes,proto3" json:"likes,omitempty"`
	LikedBy              []string    `protobuf:"bytes,4,rep,name=likedBy,proto3" json:"likedBy,omitempty"`
	Liked                bool        `protobuf:"varint,5,opt,name=liked,proto3" json:"liked,omitempty"`
	CreatedAt            string      `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            string      `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Repo) Reset()         { *m = Repo{} }
func (m *Repo) String() string { return proto.CompactTextString(m) }
func (*Repo) ProtoMessage()    {}
func (*Repo) Descriptor() ([]byte, []int) {
	return fileDescriptor_repo_949795d72d739c3b, []int{0}
}
func (m *Repo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repo.Unmarshal(m, b)
}
func (m *Repo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repo.Marshal(b, m, deterministic)
}
func (dst *Repo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repo.Merge(dst, src)
}
func (m *Repo) XXX_Size() int {
	return xxx_messageInfo_Repo.Size(m)
}
func (m *Repo) XXX_DiscardUnknown() {
	xxx_messageInfo_Repo.DiscardUnknown(m)
}

var xxx_messageInfo_Repo proto.InternalMessageInfo

func (m *Repo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Repo) GetGithubRepo() *GithubRepo {
	if m != nil {
		return m.GithubRepo
	}
	return nil
}

func (m *Repo) GetLikes() int64 {
	if m != nil {
		return m.Likes
	}
	return 0
}

func (m *Repo) GetLikedBy() []string {
	if m != nil {
		return m.LikedBy
	}
	return nil
}

func (m *Repo) GetLiked() bool {
	if m != nil {
		return m.Liked
	}
	return false
}

func (m *Repo) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Repo) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*Repo)(nil), "Repo")
}

func init() { proto.RegisterFile("repo.proto", fileDescriptor_repo_949795d72d739c3b) }

var fileDescriptor_repo_949795d72d739c3b = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0xc9, 0x4c, 0x7f, 0x9c, 0x5b, 0x10, 0x0c, 0x2e, 0x82, 0xb8, 0x08, 0xae, 0x02, 0x32,
	0x33, 0xa0, 0x0f, 0x20, 0xed, 0xc6, 0x7d, 0x96, 0x6e, 0x64, 0xda, 0x7b, 0x69, 0x83, 0x33, 0x4d,
	0x48, 0x32, 0x82, 0x0f, 0xe8, 0x7b, 0x49, 0x1a, 0xdb, 0xe9, 0x2e, 0xe7, 0xfb, 0xce, 0xe2, 0xe4,
	0x02, 0x78, 0x72, 0xb6, 0x71, 0xde, 0x46, 0xfb, 0x70, 0xb7, 0x37, 0xf1, 0x30, 0x6e, 0x3f, 0x27,
	0xf4, 0xf4, 0xcb, 0x60, 0xa6, 0xc9, 0x59, 0x7e, 0x0b, 0x85, 0x41, 0xc1, 0x24, 0x53, 0x95, 0x2e,
	0x0c, 0xf2, 0x67, 0x80, 0xdc, 0x4e, 0x56, 0x14, 0x92, 0xa9, 0xd5, 0xcb, 0xaa, 0x79, 0xbf, 0x20,
	0x7d, 0xa5, 0xf9, 0x3d, 0xcc, 0x7b, 0xf3, 0x45, 0x41, 0x94, 0x92, 0xa9, 0x52, 0xe7, 0xc0, 0x05,
	0x2c, 0xd3, 0x03, 0x37, 0x3f, 0x62, 0x26, 0x4b, 0x55, 0xe9, 0x73, 0x3c, 0xf7, 0x51, 0xcc, 0x25,
	0x53, 0x37, 0xb9, 0x8f, 0xfc, 0x11, 0xaa, 0x9d, 0xa7, 0x2e, 0x12, 0xae, 0xa3, 0x58, 0x9c, 0x96,
	0x4c, 0x20, 0xd9, 0xd1, 0xe1, 0xbf, 0x5d, 0x66, 0x7b, 0x01, 0x9b, 0xf5, 0xc7, 0x5b, 0xde, 0xd3,
	0xec, 0xec, 0xd0, 0xee, 0x47, 0xd3, 0x1f, 0xc8, 0x0f, 0x14, 0x22, 0x7d, 0x53, 0x68, 0xb3, 0xaa,
	0xd3, 0xbf, 0x43, 0x1d, 0xac, 0x8f, 0x75, 0xe7, 0x4c, 0x6b, 0x8e, 0x91, 0xfc, 0xb1, 0xeb, 0xdb,
	0xc1, 0x22, 0xf5, 0xdb, 0xc5, 0xe9, 0x22, 0xaf, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x09, 0xa2,
	0xff, 0xa3, 0x32, 0x01, 0x00, 0x00,
}

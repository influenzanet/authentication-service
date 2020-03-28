// Code generated by protoc-gen-go. DO NOT EDIT.
// source: study-service-api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Study struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Study) Reset()         { *m = Study{} }
func (m *Study) String() string { return proto.CompactTextString(m) }
func (*Study) ProtoMessage()    {}
func (*Study) Descriptor() ([]byte, []int) {
	return fileDescriptor_81f0d8f98f9be15c, []int{0}
}

func (m *Study) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Study.Unmarshal(m, b)
}
func (m *Study) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Study.Marshal(b, m, deterministic)
}
func (m *Study) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Study.Merge(m, src)
}
func (m *Study) XXX_Size() int {
	return xxx_messageInfo_Study.Size(m)
}
func (m *Study) XXX_DiscardUnknown() {
	xxx_messageInfo_Study.DiscardUnknown(m)
}

var xxx_messageInfo_Study proto.InternalMessageInfo

type CreateSurveyReq struct {
	Token                *TokenInfos `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	SurveyDef            *SurveyItem `protobuf:"bytes,2,opt,name=surveyDef,proto3" json:"surveyDef,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateSurveyReq) Reset()         { *m = CreateSurveyReq{} }
func (m *CreateSurveyReq) String() string { return proto.CompactTextString(m) }
func (*CreateSurveyReq) ProtoMessage()    {}
func (*CreateSurveyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_81f0d8f98f9be15c, []int{1}
}

func (m *CreateSurveyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSurveyReq.Unmarshal(m, b)
}
func (m *CreateSurveyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSurveyReq.Marshal(b, m, deterministic)
}
func (m *CreateSurveyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSurveyReq.Merge(m, src)
}
func (m *CreateSurveyReq) XXX_Size() int {
	return xxx_messageInfo_CreateSurveyReq.Size(m)
}
func (m *CreateSurveyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSurveyReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSurveyReq proto.InternalMessageInfo

func (m *CreateSurveyReq) GetToken() *TokenInfos {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *CreateSurveyReq) GetSurveyDef() *SurveyItem {
	if m != nil {
		return m.SurveyDef
	}
	return nil
}

type SubmitResponseReq struct {
	Token                *TokenInfos         `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	StudyId              string              `protobuf:"bytes,2,opt,name=study_id,json=studyId,proto3" json:"study_id,omitempty"`
	ProfileId            string              `protobuf:"bytes,3,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
	Responses            *SurveyItemResponse `protobuf:"bytes,4,opt,name=responses,proto3" json:"responses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SubmitResponseReq) Reset()         { *m = SubmitResponseReq{} }
func (m *SubmitResponseReq) String() string { return proto.CompactTextString(m) }
func (*SubmitResponseReq) ProtoMessage()    {}
func (*SubmitResponseReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_81f0d8f98f9be15c, []int{2}
}

func (m *SubmitResponseReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitResponseReq.Unmarshal(m, b)
}
func (m *SubmitResponseReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitResponseReq.Marshal(b, m, deterministic)
}
func (m *SubmitResponseReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitResponseReq.Merge(m, src)
}
func (m *SubmitResponseReq) XXX_Size() int {
	return xxx_messageInfo_SubmitResponseReq.Size(m)
}
func (m *SubmitResponseReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitResponseReq.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitResponseReq proto.InternalMessageInfo

func (m *SubmitResponseReq) GetToken() *TokenInfos {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *SubmitResponseReq) GetStudyId() string {
	if m != nil {
		return m.StudyId
	}
	return ""
}

func (m *SubmitResponseReq) GetProfileId() string {
	if m != nil {
		return m.ProfileId
	}
	return ""
}

func (m *SubmitResponseReq) GetResponses() *SurveyItemResponse {
	if m != nil {
		return m.Responses
	}
	return nil
}

func init() {
	proto.RegisterType((*Study)(nil), "inf.study_service_api.Study")
	proto.RegisterType((*CreateSurveyReq)(nil), "inf.study_service_api.CreateSurveyReq")
	proto.RegisterType((*SubmitResponseReq)(nil), "inf.study_service_api.SubmitResponseReq")
}

func init() {
	proto.RegisterFile("study-service-api.proto", fileDescriptor_81f0d8f98f9be15c)
}

var fileDescriptor_81f0d8f98f9be15c = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xdf, 0x4e, 0xf2, 0x30,
	0x18, 0xc6, 0xb3, 0x8f, 0x0f, 0xf8, 0x56, 0xc8, 0x47, 0x6c, 0x02, 0xc2, 0x8c, 0x89, 0x21, 0x51,
	0x39, 0x59, 0x49, 0xd0, 0x1b, 0xf0, 0x0f, 0x31, 0x3b, 0x74, 0x33, 0x1e, 0x78, 0x42, 0x36, 0xf7,
	0x8e, 0x34, 0x8e, 0xb5, 0xb6, 0x1d, 0x09, 0xd7, 0xe6, 0xad, 0x78, 0x31, 0x66, 0x6d, 0x11, 0x10,
	0x3d, 0xf0, 0xf0, 0xed, 0xf3, 0xf4, 0x7d, 0x9e, 0xfc, 0x5e, 0x74, 0x28, 0x55, 0x99, 0xae, 0x7c,
	0x09, 0x62, 0x49, 0x9f, 0xc1, 0x8f, 0x39, 0x25, 0x5c, 0x30, 0xc5, 0x70, 0x97, 0x16, 0x19, 0xd1,
	0xe2, 0xcc, 0x8a, 0xb3, 0x98, 0x53, 0x0f, 0xcf, 0x73, 0x96, 0xc4, 0xb9, 0xaf, 0x56, 0x1c, 0xa4,
	0xb1, 0x7a, 0x6d, 0x59, 0x8a, 0x25, 0xac, 0xec, 0xd4, 0x35, 0x93, 0x2f, 0x40, 0x72, 0x56, 0x48,
	0xb0, 0xcf, 0x47, 0x73, 0xc6, 0xe6, 0x39, 0x8c, 0xf5, 0x94, 0x94, 0xd9, 0x18, 0x16, 0x5c, 0xd9,
	0x3f, 0xc3, 0x26, 0xaa, 0x47, 0x55, 0xd4, 0xb0, 0x40, 0x9d, 0x1b, 0x01, 0xb1, 0x82, 0x48, 0x2f,
	0x09, 0xe1, 0x15, 0x9f, 0xa2, 0xba, 0x62, 0x2f, 0x50, 0xf4, 0x9d, 0x13, 0x67, 0xd4, 0x9a, 0x74,
	0x48, 0x55, 0xec, 0xa1, 0x7a, 0x09, 0x8a, 0x8c, 0xc9, 0xd0, 0xa8, 0xf8, 0x12, 0xb9, 0x26, 0xf8,
	0x16, 0xb2, 0xfe, 0x1f, 0x6d, 0xed, 0x69, 0xab, 0x2d, 0x67, 0x16, 0x06, 0x0a, 0x16, 0xe1, 0xc6,
	0x38, 0x7c, 0x73, 0xd0, 0x41, 0x54, 0x26, 0x0b, 0xaa, 0x42, 0x5b, 0xf7, 0x17, 0x91, 0x03, 0xf4,
	0xcf, 0x00, 0xa2, 0xa9, 0x4e, 0x74, 0xc3, 0xa6, 0x9e, 0x83, 0x14, 0x1f, 0x23, 0xc4, 0x05, 0xcb,
	0x68, 0x0e, 0x95, 0x58, 0xd3, 0xa2, 0x6b, 0x5f, 0x82, 0x14, 0x4f, 0x91, 0xbb, 0xc6, 0x23, 0xfb,
	0x7f, 0x75, 0xc8, 0xf9, 0x56, 0xd9, 0xd9, 0x27, 0xbb, 0xad, 0xd6, 0xeb, 0x7e, 0x9b, 0x9f, 0x93,
	0x77, 0x07, 0x75, 0x34, 0xb7, 0xc8, 0x5c, 0xe8, 0x8a, 0x53, 0xec, 0xa3, 0x46, 0xa4, 0x62, 0x55,
	0x4a, 0xdc, 0x23, 0x06, 0x39, 0x59, 0x23, 0x27, 0xd3, 0x0a, 0xb9, 0xd7, 0xd2, 0x49, 0xd6, 0x74,
	0x8f, 0xda, 0xdb, 0xc0, 0xf1, 0x19, 0xf9, 0xf6, 0xee, 0xe4, 0xcb, 0x55, 0xbc, 0xc1, 0x3e, 0xdb,
	0x47, 0x10, 0x92, 0xb2, 0x02, 0xdf, 0xa1, 0xff, 0xbb, 0x48, 0xf1, 0xe8, 0x87, 0xa5, 0x7b, 0xe4,
	0x77, 0xba, 0x5d, 0xd7, 0x9f, 0x6a, 0x31, 0xa7, 0x49, 0x43, 0xf7, 0xbf, 0xf8, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x6f, 0x9c, 0xb9, 0x90, 0xab, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StudyServiceApiClient is the client API for StudyServiceApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StudyServiceApiClient interface {
	Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Status, error)
	CreateSurvey(ctx context.Context, in *CreateSurveyReq, opts ...grpc.CallOption) (*SurveyVersion, error)
	SubmitResponse(ctx context.Context, in *SubmitResponseReq, opts ...grpc.CallOption) (*Status, error)
}

type studyServiceApiClient struct {
	cc grpc.ClientConnInterface
}

func NewStudyServiceApiClient(cc grpc.ClientConnInterface) StudyServiceApiClient {
	return &studyServiceApiClient{cc}
}

func (c *studyServiceApiClient) Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/inf.study_service_api.StudyServiceApi/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studyServiceApiClient) CreateSurvey(ctx context.Context, in *CreateSurveyReq, opts ...grpc.CallOption) (*SurveyVersion, error) {
	out := new(SurveyVersion)
	err := c.cc.Invoke(ctx, "/inf.study_service_api.StudyServiceApi/CreateSurvey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studyServiceApiClient) SubmitResponse(ctx context.Context, in *SubmitResponseReq, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/inf.study_service_api.StudyServiceApi/SubmitResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudyServiceApiServer is the server API for StudyServiceApi service.
type StudyServiceApiServer interface {
	Status(context.Context, *empty.Empty) (*Status, error)
	CreateSurvey(context.Context, *CreateSurveyReq) (*SurveyVersion, error)
	SubmitResponse(context.Context, *SubmitResponseReq) (*Status, error)
}

// UnimplementedStudyServiceApiServer can be embedded to have forward compatible implementations.
type UnimplementedStudyServiceApiServer struct {
}

func (*UnimplementedStudyServiceApiServer) Status(ctx context.Context, req *empty.Empty) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (*UnimplementedStudyServiceApiServer) CreateSurvey(ctx context.Context, req *CreateSurveyReq) (*SurveyVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSurvey not implemented")
}
func (*UnimplementedStudyServiceApiServer) SubmitResponse(ctx context.Context, req *SubmitResponseReq) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitResponse not implemented")
}

func RegisterStudyServiceApiServer(s *grpc.Server, srv StudyServiceApiServer) {
	s.RegisterService(&_StudyServiceApi_serviceDesc, srv)
}

func _StudyServiceApi_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudyServiceApiServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inf.study_service_api.StudyServiceApi/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudyServiceApiServer).Status(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudyServiceApi_CreateSurvey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSurveyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudyServiceApiServer).CreateSurvey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inf.study_service_api.StudyServiceApi/CreateSurvey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudyServiceApiServer).CreateSurvey(ctx, req.(*CreateSurveyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudyServiceApi_SubmitResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitResponseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudyServiceApiServer).SubmitResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inf.study_service_api.StudyServiceApi/SubmitResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudyServiceApiServer).SubmitResponse(ctx, req.(*SubmitResponseReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _StudyServiceApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "inf.study_service_api.StudyServiceApi",
	HandlerType: (*StudyServiceApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _StudyServiceApi_Status_Handler,
		},
		{
			MethodName: "CreateSurvey",
			Handler:    _StudyServiceApi_CreateSurvey_Handler,
		},
		{
			MethodName: "SubmitResponse",
			Handler:    _StudyServiceApi_SubmitResponse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "study-service-api.proto",
}

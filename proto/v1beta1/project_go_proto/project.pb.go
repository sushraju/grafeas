// Code generated by protoc-gen-go. DO NOT EDIT.
// source: project.proto

package project_go_proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

// Request to create a new project.
type CreateProjectRequest struct {
	// The project to create.
	Project              *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateProjectRequest) Reset()         { *m = CreateProjectRequest{} }
func (m *CreateProjectRequest) String() string { return proto.CompactTextString(m) }
func (*CreateProjectRequest) ProtoMessage()    {}
func (*CreateProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8340e6318dfdfac2, []int{0}
}

func (m *CreateProjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProjectRequest.Unmarshal(m, b)
}
func (m *CreateProjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProjectRequest.Marshal(b, m, deterministic)
}
func (m *CreateProjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProjectRequest.Merge(m, src)
}
func (m *CreateProjectRequest) XXX_Size() int {
	return xxx_messageInfo_CreateProjectRequest.Size(m)
}
func (m *CreateProjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProjectRequest proto.InternalMessageInfo

func (m *CreateProjectRequest) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

// Request to get a project.
type GetProjectRequest struct {
	// The name of the project in the form of `projects/{PROJECT_ID}`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProjectRequest) Reset()         { *m = GetProjectRequest{} }
func (m *GetProjectRequest) String() string { return proto.CompactTextString(m) }
func (*GetProjectRequest) ProtoMessage()    {}
func (*GetProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8340e6318dfdfac2, []int{1}
}

func (m *GetProjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProjectRequest.Unmarshal(m, b)
}
func (m *GetProjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProjectRequest.Marshal(b, m, deterministic)
}
func (m *GetProjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProjectRequest.Merge(m, src)
}
func (m *GetProjectRequest) XXX_Size() int {
	return xxx_messageInfo_GetProjectRequest.Size(m)
}
func (m *GetProjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProjectRequest proto.InternalMessageInfo

func (m *GetProjectRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request to list projects.
type ListProjectsRequest struct {
	// The filter expression.
	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// Number of projects to return in the list.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Token to provide to skip to a particular spot in the list.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListProjectsRequest) Reset()         { *m = ListProjectsRequest{} }
func (m *ListProjectsRequest) String() string { return proto.CompactTextString(m) }
func (*ListProjectsRequest) ProtoMessage()    {}
func (*ListProjectsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8340e6318dfdfac2, []int{2}
}

func (m *ListProjectsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProjectsRequest.Unmarshal(m, b)
}
func (m *ListProjectsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProjectsRequest.Marshal(b, m, deterministic)
}
func (m *ListProjectsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProjectsRequest.Merge(m, src)
}
func (m *ListProjectsRequest) XXX_Size() int {
	return xxx_messageInfo_ListProjectsRequest.Size(m)
}
func (m *ListProjectsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProjectsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListProjectsRequest proto.InternalMessageInfo

func (m *ListProjectsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListProjectsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListProjectsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Request to delete a project.
type DeleteProjectRequest struct {
	// The name of the project in the form of `projects/{PROJECT_ID}`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProjectRequest) Reset()         { *m = DeleteProjectRequest{} }
func (m *DeleteProjectRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteProjectRequest) ProtoMessage()    {}
func (*DeleteProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8340e6318dfdfac2, []int{3}
}

func (m *DeleteProjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProjectRequest.Unmarshal(m, b)
}
func (m *DeleteProjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProjectRequest.Marshal(b, m, deterministic)
}
func (m *DeleteProjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProjectRequest.Merge(m, src)
}
func (m *DeleteProjectRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteProjectRequest.Size(m)
}
func (m *DeleteProjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProjectRequest proto.InternalMessageInfo

func (m *DeleteProjectRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Response for listing projects.
type ListProjectsResponse struct {
	// The projects requested.
	Projects []*Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	// The next pagination token in the list response. It should be used as
	// `page_token` for the following request. An empty value means no more
	// results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListProjectsResponse) Reset()         { *m = ListProjectsResponse{} }
func (m *ListProjectsResponse) String() string { return proto.CompactTextString(m) }
func (*ListProjectsResponse) ProtoMessage()    {}
func (*ListProjectsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8340e6318dfdfac2, []int{4}
}

func (m *ListProjectsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProjectsResponse.Unmarshal(m, b)
}
func (m *ListProjectsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProjectsResponse.Marshal(b, m, deterministic)
}
func (m *ListProjectsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProjectsResponse.Merge(m, src)
}
func (m *ListProjectsResponse) XXX_Size() int {
	return xxx_messageInfo_ListProjectsResponse.Size(m)
}
func (m *ListProjectsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProjectsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListProjectsResponse proto.InternalMessageInfo

func (m *ListProjectsResponse) GetProjects() []*Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

func (m *ListProjectsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Describes a Grafeas project.
type Project struct {
	// The name of the project in the form of `projects/{PROJECT_ID}`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_8340e6318dfdfac2, []int{5}
}

func (m *Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Project.Unmarshal(m, b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Project.Marshal(b, m, deterministic)
}
func (m *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(m, src)
}
func (m *Project) XXX_Size() int {
	return xxx_messageInfo_Project.Size(m)
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateProjectRequest)(nil), "grafeas.v1beta1.project.CreateProjectRequest")
	proto.RegisterType((*GetProjectRequest)(nil), "grafeas.v1beta1.project.GetProjectRequest")
	proto.RegisterType((*ListProjectsRequest)(nil), "grafeas.v1beta1.project.ListProjectsRequest")
	proto.RegisterType((*DeleteProjectRequest)(nil), "grafeas.v1beta1.project.DeleteProjectRequest")
	proto.RegisterType((*ListProjectsResponse)(nil), "grafeas.v1beta1.project.ListProjectsResponse")
	proto.RegisterType((*Project)(nil), "grafeas.v1beta1.project.Project")
}

func init() { proto.RegisterFile("project.proto", fileDescriptor_8340e6318dfdfac2) }

var fileDescriptor_8340e6318dfdfac2 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x93, 0xd2, 0xa6, 0x03, 0x11, 0xea, 0x34, 0x2a, 0xc1, 0x6d, 0xa5, 0xc8, 0x07, 0xa8,
	0x22, 0xba, 0x56, 0xcb, 0x89, 0x0a, 0x0e, 0xfc, 0xa9, 0x17, 0x0e, 0x91, 0xe1, 0xc4, 0xc5, 0xda,
	0x54, 0x13, 0xb3, 0x90, 0x78, 0x8d, 0x77, 0x83, 0x20, 0xfc, 0x1c, 0x10, 0xe2, 0x05, 0x90, 0x78,
	0x00, 0x5e, 0x89, 0x57, 0xe0, 0x41, 0x90, 0xd7, 0xbb, 0xa1, 0x69, 0xec, 0x92, 0x93, 0xbd, 0x3b,
	0xdf, 0xcc, 0xf7, 0xcd, 0x7c, 0xb3, 0xd0, 0xce, 0x72, 0xf9, 0x9a, 0xce, 0x34, 0xcb, 0x72, 0xa9,
	0x25, 0xde, 0x48, 0x72, 0x3e, 0x22, 0xae, 0xd8, 0xbb, 0xa3, 0x21, 0x69, 0x7e, 0xc4, 0x6c, 0xd8,
	0xdf, 0x4b, 0xa4, 0x4c, 0xc6, 0x14, 0xf2, 0x4c, 0x84, 0x3c, 0x4d, 0xa5, 0xe6, 0x5a, 0xc8, 0x54,
	0x95, 0x69, 0xfe, 0xae, 0x8d, 0x9a, 0xd3, 0x70, 0x3a, 0x0a, 0x69, 0x92, 0xe9, 0x0f, 0x65, 0x30,
	0x88, 0xa0, 0xf3, 0x38, 0x27, 0xae, 0x69, 0x50, 0xd6, 0x8a, 0xe8, 0xed, 0x94, 0x94, 0xc6, 0x13,
	0xd8, 0xb0, 0xd5, 0xbb, 0x5e, 0xcf, 0x3b, 0xb8, 0x7a, 0xdc, 0x63, 0x35, 0xec, 0xcc, 0x65, 0xba,
	0x84, 0xe0, 0x36, 0x6c, 0x9d, 0x92, 0xbe, 0x50, 0x10, 0x61, 0x2d, 0xe5, 0x13, 0x32, 0xd5, 0x36,
	0x23, 0xf3, 0x1f, 0x08, 0xd8, 0x7e, 0x26, 0x94, 0x43, 0x2a, 0x07, 0xdd, 0x81, 0xf5, 0x91, 0x18,
	0x6b, 0xca, 0x2d, 0xd8, 0x9e, 0x70, 0x17, 0x36, 0x33, 0x9e, 0x50, 0xac, 0xc4, 0x8c, 0xba, 0x8d,
	0x9e, 0x77, 0x70, 0x25, 0x6a, 0x15, 0x17, 0xcf, 0xc5, 0x8c, 0x70, 0x1f, 0xc0, 0x04, 0xb5, 0x7c,
	0x43, 0x69, 0xb7, 0x69, 0x12, 0x0d, 0xfc, 0x45, 0x71, 0x11, 0xf4, 0xa1, 0xf3, 0x84, 0xc6, 0xb4,
	0xd4, 0x67, 0x95, 0xac, 0x4f, 0xd0, 0x59, 0x94, 0xa5, 0x32, 0x99, 0x2a, 0xc2, 0xfb, 0xd0, 0xb2,
	0x2d, 0xaa, 0xae, 0xd7, 0x6b, 0xae, 0x34, 0x94, 0x79, 0x06, 0xde, 0x82, 0xeb, 0x29, 0xbd, 0xd7,
	0xf1, 0x39, 0x95, 0x0d, 0x43, 0xda, 0x2e, 0xae, 0x07, 0x73, 0xa5, 0xfb, 0xb0, 0x61, 0x93, 0xab,
	0xc4, 0x1d, 0xff, 0x5c, 0x83, 0x96, 0x53, 0x86, 0xdf, 0x3c, 0x68, 0x2f, 0xd8, 0x87, 0x87, 0xb5,
	0x8a, 0xaa, 0x6c, 0xf6, 0xff, 0xdb, 0x40, 0x10, 0x7c, 0xfd, 0xfd, 0xe7, 0x47, 0x63, 0x2f, 0xd8,
	0x0a, 0x2d, 0x22, 0x74, 0x1d, 0x9d, 0x38, 0xc3, 0xf1, 0x0b, 0xc0, 0x3f, 0xc3, 0xb1, 0x5f, 0x5b,
	0x73, 0x69, 0x2b, 0x56, 0xe7, 0x47, 0x7f, 0xce, 0xff, 0xb1, 0x98, 0xc3, 0x03, 0xa7, 0x22, 0xec,
	0x7f, 0xc6, 0xef, 0x1e, 0x5c, 0x3b, 0xef, 0x18, 0xde, 0xa9, 0x2d, 0x5b, 0xb1, 0x6f, 0xfe, 0xe1,
	0x8a, 0xe8, 0x72, 0x0d, 0x82, 0x9b, 0x46, 0xd1, 0x36, 0x2e, 0x4f, 0x04, 0x67, 0xd0, 0x5e, 0xd8,
	0xb2, 0x4b, 0xec, 0xa8, 0xda, 0x46, 0x7f, 0x87, 0x95, 0x6f, 0x95, 0xb9, 0xb7, 0xca, 0x9e, 0x16,
	0x6f, 0xd5, 0x0d, 0xa1, 0x7f, 0xc9, 0x10, 0x1e, 0xc5, 0xe0, 0x0b, 0x59, 0x47, 0x37, 0xf0, 0x5e,
	0xde, 0x4b, 0x84, 0x7e, 0x35, 0x1d, 0xb2, 0x33, 0x39, 0x09, 0x2d, 0x6a, 0xfe, 0x35, 0x74, 0x17,
	0xdb, 0x89, 0x13, 0x19, 0x9b, 0xc0, 0xaf, 0x46, 0xf3, 0x34, 0x7a, 0x38, 0x5c, 0x37, 0x87, 0xbb,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xec, 0x21, 0x92, 0xed, 0x96, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProjectsClient is the client API for Projects service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProjectsClient interface {
	// Creates a new project.
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*Project, error)
	// Gets the specified project.
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error)
	// Lists projects.
	ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error)
	// Deletes the specified project.
	DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type projectsClient struct {
	cc *grpc.ClientConn
}

func NewProjectsClient(cc *grpc.ClientConn) ProjectsClient {
	return &projectsClient{cc}
}

func (c *projectsClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, "/grafeas.v1beta1.project.Projects/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, "/grafeas.v1beta1.project.Projects/GetProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error) {
	out := new(ListProjectsResponse)
	err := c.cc.Invoke(ctx, "/grafeas.v1beta1.project.Projects/ListProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grafeas.v1beta1.project.Projects/DeleteProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectsServer is the server API for Projects service.
type ProjectsServer interface {
	// Creates a new project.
	CreateProject(context.Context, *CreateProjectRequest) (*Project, error)
	// Gets the specified project.
	GetProject(context.Context, *GetProjectRequest) (*Project, error)
	// Lists projects.
	ListProjects(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error)
	// Deletes the specified project.
	DeleteProject(context.Context, *DeleteProjectRequest) (*empty.Empty, error)
}

func RegisterProjectsServer(s *grpc.Server, srv ProjectsServer) {
	s.RegisterService(&_Projects_serviceDesc, srv)
}

func _Projects_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grafeas.v1beta1.project.Projects/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grafeas.v1beta1.project.Projects/GetProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).GetProject(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_ListProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).ListProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grafeas.v1beta1.project.Projects/ListProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).ListProjects(ctx, req.(*ListProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grafeas.v1beta1.project.Projects/DeleteProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).DeleteProject(ctx, req.(*DeleteProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Projects_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grafeas.v1beta1.project.Projects",
	HandlerType: (*ProjectsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _Projects_CreateProject_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _Projects_GetProject_Handler,
		},
		{
			MethodName: "ListProjects",
			Handler:    _Projects_ListProjects_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _Projects_DeleteProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "project.proto",
}
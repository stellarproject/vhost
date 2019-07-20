// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/stellarproject/terraos/api/node/v1/agent.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	v1 "github.com/stellarproject/terraos/api/types/v1"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type InfoResponse struct {
	Node                 *v1.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Image                string   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoResponse) Reset()      { *m = InfoResponse{} }
func (*InfoResponse) ProtoMessage() {}
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a976e96e0a4f5eb9, []int{0}
}
func (m *InfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoResponse.Merge(m, src)
}
func (m *InfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *InfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InfoResponse proto.InternalMessageInfo

type ProvisionRequest struct {
	Node                 *v1.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Nameservers          []string `protobuf:"bytes,2,rep,name=nameservers,proto3" json:"nameservers,omitempty"`
	SshKeys              []string `protobuf:"bytes,3,rep,name=ssh_keys,json=sshKeys,proto3" json:"ssh_keys,omitempty"`
	Gateway              string   `protobuf:"bytes,4,opt,name=gateway,proto3" json:"gateway,omitempty"`
	ClusterFs            string   `protobuf:"bytes,5,opt,name=cluster_fs,json=clusterFs,proto3" json:"cluster_fs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProvisionRequest) Reset()      { *m = ProvisionRequest{} }
func (*ProvisionRequest) ProtoMessage() {}
func (*ProvisionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a976e96e0a4f5eb9, []int{1}
}
func (m *ProvisionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProvisionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProvisionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProvisionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProvisionRequest.Merge(m, src)
}
func (m *ProvisionRequest) XXX_Size() int {
	return m.Size()
}
func (m *ProvisionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProvisionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProvisionRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*InfoResponse)(nil), "io.stellarproject.node.v1.InfoResponse")
	proto.RegisterType((*ProvisionRequest)(nil), "io.stellarproject.node.v1.ProvisionRequest")
}

func init() {
	proto.RegisterFile("github.com/stellarproject/terraos/api/node/v1/agent.proto", fileDescriptor_a976e96e0a4f5eb9)
}

var fileDescriptor_a976e96e0a4f5eb9 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xce, 0xb4, 0x89, 0x35, 0x13, 0x0f, 0x32, 0x14, 0xd9, 0x46, 0x5c, 0x96, 0x5e, 0x0c, 0x08,
	0x33, 0x26, 0x8a, 0x20, 0x9e, 0x14, 0x54, 0x54, 0x90, 0xb0, 0x27, 0xe9, 0x25, 0x4c, 0xd2, 0x97,
	0xc9, 0xea, 0xee, 0xbe, 0x75, 0xde, 0xec, 0xca, 0xde, 0x3c, 0xfa, 0xab, 0x3c, 0xf7, 0xe8, 0xd1,
	0xa3, 0xdd, 0x5f, 0x22, 0x3b, 0x69, 0x4b, 0x55, 0x56, 0xb4, 0xb7, 0xf9, 0xde, 0xf7, 0xbe, 0x37,
	0x1f, 0xef, 0x7b, 0xfc, 0xb1, 0x49, 0xdc, 0xa6, 0x5c, 0xca, 0x15, 0x66, 0x8a, 0x1c, 0xa4, 0xa9,
	0xb6, 0x85, 0xc5, 0xf7, 0xb0, 0x72, 0xca, 0x81, 0xb5, 0x1a, 0x49, 0xe9, 0x22, 0x51, 0x39, 0x1e,
	0x83, 0xaa, 0xa6, 0x4a, 0x1b, 0xc8, 0x9d, 0x2c, 0x2c, 0x3a, 0x14, 0x07, 0x09, 0xca, 0x5f, 0x25,
	0xb2, 0x6d, 0x93, 0xd5, 0x74, 0xbc, 0x6f, 0xd0, 0xa0, 0xef, 0x52, 0xed, 0x6b, 0x2b, 0x18, 0xdf,
	0x36, 0x88, 0x26, 0x05, 0xe5, 0xd1, 0xb2, 0x5c, 0x2b, 0xc8, 0x0a, 0x57, 0x9f, 0x91, 0xff, 0x68,
	0xc4, 0xd5, 0x05, 0x50, 0xeb, 0xc4, 0x7f, 0xe5, 0xa5, 0x87, 0x47, 0xfc, 0xc6, 0xab, 0x7c, 0x8d,
	0x31, 0x50, 0x81, 0x39, 0x81, 0x78, 0xc8, 0xfb, 0x2d, 0x1b, 0xb0, 0x88, 0x4d, 0x46, 0xb3, 0x48,
	0xfe, 0xe9, 0xd3, 0x4f, 0x91, 0xd5, 0x54, 0xbe, 0xc5, 0x63, 0x88, 0x7d, 0xb7, 0xd8, 0xe7, 0x83,
	0x24, 0xd3, 0x06, 0x82, 0x9d, 0x88, 0x4d, 0x86, 0xf1, 0x16, 0x1c, 0x7e, 0x65, 0xfc, 0xe6, 0xdc,
	0x62, 0x95, 0x50, 0x82, 0x79, 0x0c, 0x1f, 0x4b, 0x20, 0x77, 0xc5, 0x0f, 0x22, 0x3e, 0xca, 0x75,
	0x06, 0x04, 0xb6, 0x02, 0x4b, 0xc1, 0x4e, 0xb4, 0x3b, 0x19, 0xc6, 0x97, 0x4b, 0xe2, 0x80, 0x5f,
	0x27, 0xda, 0x2c, 0x3e, 0x40, 0x4d, 0xc1, 0xae, 0xa7, 0xf7, 0x88, 0x36, 0x6f, 0xa0, 0x26, 0x11,
	0xf0, 0x3d, 0xa3, 0x1d, 0x7c, 0xd2, 0x75, 0xd0, 0xf7, 0xfe, 0xce, 0xa1, 0xb8, 0xc3, 0xf9, 0x2a,
	0x2d, 0xc9, 0x81, 0x5d, 0xac, 0x29, 0x18, 0x78, 0x72, 0x78, 0x56, 0x79, 0x41, 0xb3, 0x2f, 0x8c,
	0x0f, 0x9e, 0xb6, 0xa9, 0x89, 0x47, 0xbc, 0x3f, 0x4f, 0x72, 0x23, 0x6e, 0xc9, 0x6d, 0x0e, 0xf2,
	0x3c, 0x07, 0xf9, 0xbc, 0xcd, 0x61, 0xdc, 0x51, 0x17, 0x2f, 0x79, 0xbf, 0x5d, 0x6f, 0xa7, 0xee,
	0xae, 0xec, 0x3c, 0x04, 0x79, 0x39, 0x97, 0xd9, 0x82, 0x8f, 0x2e, 0x56, 0x09, 0x56, 0xcc, 0xf9,
	0xf0, 0x02, 0x8a, 0x7b, 0x7f, 0x19, 0xf2, 0xfb, 0xfe, 0xbb, 0x9c, 0x3e, 0x7b, 0x7d, 0x72, 0x1a,
	0xf6, 0xbe, 0x9f, 0x86, 0xbd, 0xcf, 0x4d, 0xc8, 0x4e, 0x9a, 0x90, 0x7d, 0x6b, 0x42, 0xf6, 0xa3,
	0x09, 0xd9, 0xd1, 0xfd, 0xff, 0x3a, 0xf3, 0x27, 0xd5, 0xf4, 0x5d, 0x6f, 0x79, 0xcd, 0x4f, 0x7f,
	0xf0, 0x33, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x6b, 0xe7, 0x25, 0x23, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentClient interface {
	Ping(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error)
	Info(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*InfoResponse, error)
}

type agentClient struct {
	cc *grpc.ClientConn
}

func NewAgentClient(cc *grpc.ClientConn) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) Ping(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/io.stellarproject.node.v1.Agent/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) Info(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/io.stellarproject.node.v1.Agent/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
type AgentServer interface {
	Ping(context.Context, *types.Empty) (*types.Empty, error)
	Info(context.Context, *types.Empty) (*InfoResponse, error)
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.stellarproject.node.v1.Agent/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Ping(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.stellarproject.node.v1.Agent/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Info(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "io.stellarproject.node.v1.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Agent_Ping_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Agent_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/stellarproject/terraos/api/node/v1/agent.proto",
}

// ProvisionerClient is the client API for Provisioner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProvisionerClient interface {
	Provision(ctx context.Context, in *ProvisionRequest, opts ...grpc.CallOption) (*types.Empty, error)
}

type provisionerClient struct {
	cc *grpc.ClientConn
}

func NewProvisionerClient(cc *grpc.ClientConn) ProvisionerClient {
	return &provisionerClient{cc}
}

func (c *provisionerClient) Provision(ctx context.Context, in *ProvisionRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/io.stellarproject.node.v1.Provisioner/Provision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProvisionerServer is the server API for Provisioner service.
type ProvisionerServer interface {
	Provision(context.Context, *ProvisionRequest) (*types.Empty, error)
}

func RegisterProvisionerServer(s *grpc.Server, srv ProvisionerServer) {
	s.RegisterService(&_Provisioner_serviceDesc, srv)
}

func _Provisioner_Provision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProvisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionerServer).Provision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.stellarproject.node.v1.Provisioner/Provision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionerServer).Provision(ctx, req.(*ProvisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Provisioner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "io.stellarproject.node.v1.Provisioner",
	HandlerType: (*ProvisionerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Provision",
			Handler:    _Provisioner_Provision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/stellarproject/terraos/api/node/v1/agent.proto",
}

func (m *InfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InfoResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Node != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAgent(dAtA, i, uint64(m.Node.Size()))
		n1, err := m.Node.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Image) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.Image)))
		i += copy(dAtA[i:], m.Image)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ProvisionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProvisionRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Node != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAgent(dAtA, i, uint64(m.Node.Size()))
		n2, err := m.Node.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Nameservers) > 0 {
		for _, s := range m.Nameservers {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.SshKeys) > 0 {
		for _, s := range m.SshKeys {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Gateway) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.Gateway)))
		i += copy(dAtA[i:], m.Gateway)
	}
	if len(m.ClusterFs) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.ClusterFs)))
		i += copy(dAtA[i:], m.ClusterFs)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintAgent(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *InfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Node != nil {
		l = m.Node.Size()
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.Image)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ProvisionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Node != nil {
		l = m.Node.Size()
		n += 1 + l + sovAgent(uint64(l))
	}
	if len(m.Nameservers) > 0 {
		for _, s := range m.Nameservers {
			l = len(s)
			n += 1 + l + sovAgent(uint64(l))
		}
	}
	if len(m.SshKeys) > 0 {
		for _, s := range m.SshKeys {
			l = len(s)
			n += 1 + l + sovAgent(uint64(l))
		}
	}
	l = len(m.Gateway)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.ClusterFs)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAgent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAgent(x uint64) (n int) {
	return sovAgent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *InfoResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&InfoResponse{`,
		`Node:` + strings.Replace(fmt.Sprintf("%v", this.Node), "Node", "v1.Node", 1) + `,`,
		`Image:` + fmt.Sprintf("%v", this.Image) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProvisionRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProvisionRequest{`,
		`Node:` + strings.Replace(fmt.Sprintf("%v", this.Node), "Node", "v1.Node", 1) + `,`,
		`Nameservers:` + fmt.Sprintf("%v", this.Nameservers) + `,`,
		`SshKeys:` + fmt.Sprintf("%v", this.SshKeys) + `,`,
		`Gateway:` + fmt.Sprintf("%v", this.Gateway) + `,`,
		`ClusterFs:` + fmt.Sprintf("%v", this.ClusterFs) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAgent(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *InfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAgent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Node == nil {
				m.Node = &v1.Node{}
			}
			if err := m.Node.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Image = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAgent
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAgent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProvisionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAgent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProvisionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProvisionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Node == nil {
				m.Node = &v1.Node{}
			}
			if err := m.Node.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nameservers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nameservers = append(m.Nameservers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SshKeys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SshKeys = append(m.SshKeys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gateway", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gateway = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterFs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterFs = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAgent
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAgent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAgent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAgent
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAgent
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAgent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAgent
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAgent(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAgent
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAgent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAgent   = fmt.Errorf("proto: integer overflow")
)
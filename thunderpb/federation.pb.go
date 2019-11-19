// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: federation.proto

package thunderpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ExecuteRequest_Kind int32

const (
	ExecuteRequest_QUERY    ExecuteRequest_Kind = 0
	ExecuteRequest_MUTATION ExecuteRequest_Kind = 1
)

var ExecuteRequest_Kind_name = map[int32]string{
	0: "QUERY",
	1: "MUTATION",
}
var ExecuteRequest_Kind_value = map[string]int32{
	"QUERY":    0,
	"MUTATION": 1,
}

func (x ExecuteRequest_Kind) String() string {
	return proto.EnumName(ExecuteRequest_Kind_name, int32(x))
}
func (ExecuteRequest_Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorFederation, []int{3, 0}
}

type Selection struct {
	Name         string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Alias        string        `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	SelectionSet *SelectionSet `protobuf:"bytes,3,opt,name=selectionSet" json:"selectionSet,omitempty"`
	Arguments    []byte        `protobuf:"bytes,4,opt,name=arguments,proto3" json:"arguments,omitempty"`
}

func (m *Selection) Reset()                    { *m = Selection{} }
func (m *Selection) String() string            { return proto.CompactTextString(m) }
func (*Selection) ProtoMessage()               {}
func (*Selection) Descriptor() ([]byte, []int) { return fileDescriptorFederation, []int{0} }

func (m *Selection) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Selection) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *Selection) GetSelectionSet() *SelectionSet {
	if m != nil {
		return m.SelectionSet
	}
	return nil
}

func (m *Selection) GetArguments() []byte {
	if m != nil {
		return m.Arguments
	}
	return nil
}

type Fragment struct {
	On           string        `protobuf:"bytes,1,opt,name=on,proto3" json:"on,omitempty"`
	SelectionSet *SelectionSet `protobuf:"bytes,2,opt,name=selectionSet" json:"selectionSet,omitempty"`
}

func (m *Fragment) Reset()                    { *m = Fragment{} }
func (m *Fragment) String() string            { return proto.CompactTextString(m) }
func (*Fragment) ProtoMessage()               {}
func (*Fragment) Descriptor() ([]byte, []int) { return fileDescriptorFederation, []int{1} }

func (m *Fragment) GetOn() string {
	if m != nil {
		return m.On
	}
	return ""
}

func (m *Fragment) GetSelectionSet() *SelectionSet {
	if m != nil {
		return m.SelectionSet
	}
	return nil
}

type SelectionSet struct {
	Selections []*Selection `protobuf:"bytes,1,rep,name=selections" json:"selections,omitempty"`
	Fragments  []*Fragment  `protobuf:"bytes,2,rep,name=fragments" json:"fragments,omitempty"`
}

func (m *SelectionSet) Reset()                    { *m = SelectionSet{} }
func (m *SelectionSet) String() string            { return proto.CompactTextString(m) }
func (*SelectionSet) ProtoMessage()               {}
func (*SelectionSet) Descriptor() ([]byte, []int) { return fileDescriptorFederation, []int{2} }

func (m *SelectionSet) GetSelections() []*Selection {
	if m != nil {
		return m.Selections
	}
	return nil
}

func (m *SelectionSet) GetFragments() []*Fragment {
	if m != nil {
		return m.Fragments
	}
	return nil
}

type ExecuteRequest struct {
	Kind         ExecuteRequest_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=thunderpb.ExecuteRequest_Kind" json:"kind,omitempty"`
	Name         string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SelectionSet *SelectionSet       `protobuf:"bytes,3,opt,name=selectionSet" json:"selectionSet,omitempty"`
}

func (m *ExecuteRequest) Reset()                    { *m = ExecuteRequest{} }
func (m *ExecuteRequest) String() string            { return proto.CompactTextString(m) }
func (*ExecuteRequest) ProtoMessage()               {}
func (*ExecuteRequest) Descriptor() ([]byte, []int) { return fileDescriptorFederation, []int{3} }

func (m *ExecuteRequest) GetKind() ExecuteRequest_Kind {
	if m != nil {
		return m.Kind
	}
	return ExecuteRequest_QUERY
}

func (m *ExecuteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ExecuteRequest) GetSelectionSet() *SelectionSet {
	if m != nil {
		return m.SelectionSet
	}
	return nil
}

type ExecuteResponse struct {
	Result []byte `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *ExecuteResponse) Reset()                    { *m = ExecuteResponse{} }
func (m *ExecuteResponse) String() string            { return proto.CompactTextString(m) }
func (*ExecuteResponse) ProtoMessage()               {}
func (*ExecuteResponse) Descriptor() ([]byte, []int) { return fileDescriptorFederation, []int{4} }

func (m *ExecuteResponse) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*Selection)(nil), "thunderpb.Selection")
	proto.RegisterType((*Fragment)(nil), "thunderpb.Fragment")
	proto.RegisterType((*SelectionSet)(nil), "thunderpb.SelectionSet")
	proto.RegisterType((*ExecuteRequest)(nil), "thunderpb.ExecuteRequest")
	proto.RegisterType((*ExecuteResponse)(nil), "thunderpb.ExecuteResponse")
	proto.RegisterEnum("thunderpb.ExecuteRequest_Kind", ExecuteRequest_Kind_name, ExecuteRequest_Kind_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Executor service

type ExecutorClient interface {
	Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error)
}

type executorClient struct {
	cc *grpc.ClientConn
}

func NewExecutorClient(cc *grpc.ClientConn) ExecutorClient {
	return &executorClient{cc}
}

func (c *executorClient) Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error) {
	out := new(ExecuteResponse)
	err := grpc.Invoke(ctx, "/thunderpb.Executor/Execute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Executor service

type ExecutorServer interface {
	Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error)
}

func RegisterExecutorServer(s *grpc.Server, srv ExecutorServer) {
	s.RegisterService(&_Executor_serviceDesc, srv)
}

func _Executor_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thunderpb.Executor/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).Execute(ctx, req.(*ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Executor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "thunderpb.Executor",
	HandlerType: (*ExecutorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _Executor_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "federation.proto",
}

func (m *Selection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Selection) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFederation(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Alias) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFederation(dAtA, i, uint64(len(m.Alias)))
		i += copy(dAtA[i:], m.Alias)
	}
	if m.SelectionSet != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintFederation(dAtA, i, uint64(m.SelectionSet.Size()))
		n1, err := m.SelectionSet.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Arguments) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintFederation(dAtA, i, uint64(len(m.Arguments)))
		i += copy(dAtA[i:], m.Arguments)
	}
	return i, nil
}

func (m *Fragment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fragment) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.On) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFederation(dAtA, i, uint64(len(m.On)))
		i += copy(dAtA[i:], m.On)
	}
	if m.SelectionSet != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFederation(dAtA, i, uint64(m.SelectionSet.Size()))
		n2, err := m.SelectionSet.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *SelectionSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SelectionSet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Selections) > 0 {
		for _, msg := range m.Selections {
			dAtA[i] = 0xa
			i++
			i = encodeVarintFederation(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Fragments) > 0 {
		for _, msg := range m.Fragments {
			dAtA[i] = 0x12
			i++
			i = encodeVarintFederation(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ExecuteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExecuteRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Kind != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFederation(dAtA, i, uint64(m.Kind))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFederation(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.SelectionSet != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintFederation(dAtA, i, uint64(m.SelectionSet.Size()))
		n3, err := m.SelectionSet.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *ExecuteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExecuteResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Result) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFederation(dAtA, i, uint64(len(m.Result)))
		i += copy(dAtA[i:], m.Result)
	}
	return i, nil
}

func encodeVarintFederation(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Selection) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovFederation(uint64(l))
	}
	l = len(m.Alias)
	if l > 0 {
		n += 1 + l + sovFederation(uint64(l))
	}
	if m.SelectionSet != nil {
		l = m.SelectionSet.Size()
		n += 1 + l + sovFederation(uint64(l))
	}
	l = len(m.Arguments)
	if l > 0 {
		n += 1 + l + sovFederation(uint64(l))
	}
	return n
}

func (m *Fragment) Size() (n int) {
	var l int
	_ = l
	l = len(m.On)
	if l > 0 {
		n += 1 + l + sovFederation(uint64(l))
	}
	if m.SelectionSet != nil {
		l = m.SelectionSet.Size()
		n += 1 + l + sovFederation(uint64(l))
	}
	return n
}

func (m *SelectionSet) Size() (n int) {
	var l int
	_ = l
	if len(m.Selections) > 0 {
		for _, e := range m.Selections {
			l = e.Size()
			n += 1 + l + sovFederation(uint64(l))
		}
	}
	if len(m.Fragments) > 0 {
		for _, e := range m.Fragments {
			l = e.Size()
			n += 1 + l + sovFederation(uint64(l))
		}
	}
	return n
}

func (m *ExecuteRequest) Size() (n int) {
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovFederation(uint64(m.Kind))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovFederation(uint64(l))
	}
	if m.SelectionSet != nil {
		l = m.SelectionSet.Size()
		n += 1 + l + sovFederation(uint64(l))
	}
	return n
}

func (m *ExecuteResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovFederation(uint64(l))
	}
	return n
}

func sovFederation(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFederation(x uint64) (n int) {
	return sovFederation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Selection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFederation
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Selection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Selection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFederation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFederation
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alias", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFederation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFederation
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Alias = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelectionSet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFederation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFederation
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SelectionSet == nil {
				m.SelectionSet = &SelectionSet{}
			}
			if err := m.SelectionSet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Arguments", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFederation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFederation
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Arguments = append(m.Arguments[:0], dAtA[iNdEx:postIndex]...)
			if m.Arguments == nil {
				m.Arguments = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFederation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFederation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Fragment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFederation
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Fragment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fragment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field On", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFederation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFederation
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.On = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelectionSet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFederation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFederation
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SelectionSet == nil {
				m.SelectionSet = &SelectionSet{}
			}
			if err := m.SelectionSet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFederation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFederation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SelectionSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFederation
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SelectionSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SelectionSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Selections", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFederation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFederation
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Selections = append(m.Selections, &Selection{})
			if err := m.Selections[len(m.Selections)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fragments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFederation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFederation
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fragments = append(m.Fragments, &Fragment{})
			if err := m.Fragments[len(m.Fragments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFederation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFederation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExecuteRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFederation
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExecuteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExecuteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			m.Kind = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFederation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Kind |= (ExecuteRequest_Kind(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFederation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFederation
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelectionSet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFederation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFederation
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SelectionSet == nil {
				m.SelectionSet = &SelectionSet{}
			}
			if err := m.SelectionSet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFederation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFederation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExecuteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFederation
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExecuteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExecuteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFederation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFederation
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = append(m.Result[:0], dAtA[iNdEx:postIndex]...)
			if m.Result == nil {
				m.Result = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFederation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFederation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFederation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFederation
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
					return 0, ErrIntOverflowFederation
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
					return 0, ErrIntOverflowFederation
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthFederation
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFederation
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
				next, err := skipFederation(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthFederation = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFederation   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("federation.proto", fileDescriptorFederation) }

var fileDescriptorFederation = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x9d, 0xb3, 0x6c, 0x34, 0x77, 0x51, 0xa9, 0xcc, 0x04, 0xa1, 0x42, 0x21, 0xca, 0x03, 0x0a,
	0x0f, 0x24, 0x22, 0xec, 0x8d, 0x27, 0x26, 0x0d, 0x09, 0x21, 0x86, 0x70, 0x37, 0x21, 0x78, 0x73,
	0x1a, 0x37, 0x8d, 0x68, 0xec, 0xce, 0x76, 0x04, 0x9f, 0xc1, 0xd7, 0xf0, 0x0d, 0x3c, 0xf2, 0x09,
	0xa8, 0x5f, 0x82, 0xe2, 0x64, 0x59, 0x2a, 0xca, 0x03, 0xda, 0x9b, 0x8f, 0xef, 0x39, 0xe7, 0x1e,
	0xfb, 0x5e, 0x98, 0x2c, 0x58, 0xce, 0x24, 0xd5, 0xa5, 0xe0, 0xf1, 0x5a, 0x0a, 0x2d, 0xb0, 0xa3,
	0x97, 0x35, 0xcf, 0x99, 0x5c, 0x67, 0xd3, 0x67, 0x45, 0xa9, 0x97, 0x75, 0x16, 0xcf, 0x45, 0x95,
	0x14, 0xa2, 0x10, 0x89, 0x61, 0x64, 0xf5, 0xc2, 0x20, 0x03, 0xcc, 0xa9, 0x55, 0x86, 0xdf, 0x11,
	0x38, 0x33, 0xb6, 0x62, 0xf3, 0xc6, 0x0d, 0x63, 0xb0, 0x39, 0xad, 0x98, 0x87, 0x02, 0x14, 0x39,
	0xc4, 0x9c, 0xf1, 0x31, 0x1c, 0xd0, 0x55, 0x49, 0x95, 0x67, 0x99, 0xcb, 0x16, 0xe0, 0x97, 0xe0,
	0xaa, 0x6b, 0xd9, 0x8c, 0x69, 0x6f, 0x3f, 0x40, 0xd1, 0x51, 0xfa, 0x20, 0xee, 0x83, 0xc4, 0xb3,
	0x41, 0x99, 0x6c, 0x91, 0xf1, 0x23, 0x70, 0xa8, 0x2c, 0xea, 0x8a, 0x71, 0xad, 0x3c, 0x3b, 0x40,
	0x91, 0x4b, 0x6e, 0x2e, 0xc2, 0x8f, 0x30, 0x7a, 0x2d, 0x69, 0xd1, 0x00, 0x3c, 0x06, 0x4b, 0xf0,
	0x2e, 0x8e, 0x25, 0xf8, 0x5f, 0x6d, 0xad, 0xff, 0x68, 0x1b, 0x7e, 0x05, 0x77, 0x58, 0xc5, 0x27,
	0x00, 0x7d, 0x5d, 0x79, 0x28, 0xd8, 0x8f, 0x8e, 0xd2, 0xe3, 0x5d, 0x56, 0x64, 0xc0, 0xc3, 0xcf,
	0xc1, 0x59, 0x74, 0xf1, 0x9a, 0x3f, 0x69, 0x44, 0xf7, 0x06, 0xa2, 0xeb, 0xe8, 0xe4, 0x86, 0x15,
	0xfe, 0x40, 0x30, 0x3e, 0xfb, 0xc6, 0xe6, 0xb5, 0x66, 0x84, 0x5d, 0xd5, 0x4c, 0x69, 0x9c, 0x82,
	0xfd, 0xa5, 0xe4, 0xb9, 0x79, 0xda, 0x38, 0xf5, 0x07, 0x06, 0xdb, 0xc4, 0xf8, 0x6d, 0xc9, 0x73,
	0x62, 0xb8, 0xfd, 0x74, 0xac, 0xc1, 0x74, 0x6e, 0x33, 0x87, 0xf0, 0x31, 0xd8, 0x8d, 0x3d, 0x76,
	0xe0, 0xe0, 0xc3, 0xe5, 0x19, 0xf9, 0x34, 0xd9, 0xc3, 0x2e, 0x8c, 0xde, 0x5d, 0x5e, 0xbc, 0xba,
	0x78, 0xf3, 0xfe, 0x7c, 0x82, 0xc2, 0xa7, 0x70, 0xb7, 0x8f, 0xa3, 0xd6, 0x82, 0x2b, 0x86, 0xef,
	0xc3, 0xa1, 0x64, 0xaa, 0x5e, 0x69, 0x13, 0xdd, 0x25, 0x1d, 0x4a, 0xcf, 0x61, 0xd4, 0x52, 0x85,
	0xc4, 0xa7, 0x70, 0xa7, 0x93, 0xe1, 0x87, 0xff, 0x7c, 0xd9, 0x74, 0xba, 0xab, 0xd4, 0x76, 0x09,
	0xf7, 0x4e, 0x4f, 0x7e, 0x6e, 0x7c, 0xf4, 0x6b, 0xe3, 0xa3, 0xdf, 0x1b, 0x1f, 0x7d, 0x7e, 0x32,
	0xd8, 0x6a, 0x45, 0x2b, 0x45, 0x25, 0x5d, 0x5e, 0x25, 0x9d, 0x3e, 0xe9, 0x7d, 0xb2, 0x43, 0xb3,
	0xd5, 0x2f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x09, 0x67, 0x13, 0xbc, 0x23, 0x03, 0x00, 0x00,
}

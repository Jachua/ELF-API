// Code generated by protoc-gen-go. DO NOT EDIT.
// source: play.proto

package play

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Step struct {
	X                    int32    `protobuf:"zigzag32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"zigzag32,2,opt,name=y,proto3" json:"y,omitempty"`
	Player               *Player  `protobuf:"bytes,3,opt,name=player,proto3" json:"player,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Step) Reset()         { *m = Step{} }
func (m *Step) String() string { return proto.CompactTextString(m) }
func (*Step) ProtoMessage()    {}
func (*Step) Descriptor() ([]byte, []int) {
	return fileDescriptor_play_97261441d1e200c2, []int{0}
}
func (m *Step) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Step.Unmarshal(m, b)
}
func (m *Step) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Step.Marshal(b, m, deterministic)
}
func (dst *Step) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Step.Merge(dst, src)
}
func (m *Step) XXX_Size() int {
	return xxx_messageInfo_Step.Size(m)
}
func (m *Step) XXX_DiscardUnknown() {
	xxx_messageInfo_Step.DiscardUnknown(m)
}

var xxx_messageInfo_Step proto.InternalMessageInfo

func (m *Step) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Step) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Step) GetPlayer() *Player {
	if m != nil {
		return m.Player
	}
	return nil
}

type Player struct {
	Color                uint32   `protobuf:"varint,1,opt,name=color,proto3" json:"color,omitempty"`
	ID                   string   `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_play_97261441d1e200c2, []int{1}
}
func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (dst *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(dst, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetColor() uint32 {
	if m != nil {
		return m.Color
	}
	return 0
}

func (m *Player) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type State struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	ID                   string   `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID               string   `protobuf:"bytes,3,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_play_97261441d1e200c2, []int{2}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (dst *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(dst, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *State) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *State) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type Resumed struct {
	Move                 []string `protobuf:"bytes,1,rep,name=move,proto3" json:"move,omitempty"`
	ID                   string   `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resumed) Reset()         { *m = Resumed{} }
func (m *Resumed) String() string { return proto.CompactTextString(m) }
func (*Resumed) ProtoMessage()    {}
func (*Resumed) Descriptor() ([]byte, []int) {
	return fileDescriptor_play_97261441d1e200c2, []int{3}
}
func (m *Resumed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resumed.Unmarshal(m, b)
}
func (m *Resumed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resumed.Marshal(b, m, deterministic)
}
func (dst *Resumed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resumed.Merge(dst, src)
}
func (m *Resumed) XXX_Size() int {
	return xxx_messageInfo_Resumed.Size(m)
}
func (m *Resumed) XXX_DiscardUnknown() {
	xxx_messageInfo_Resumed.DiscardUnknown(m)
}

var xxx_messageInfo_Resumed proto.InternalMessageInfo

func (m *Resumed) GetMove() []string {
	if m != nil {
		return m.Move
	}
	return nil
}

func (m *Resumed) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*Step)(nil), "play.Step")
	proto.RegisterType((*Player)(nil), "play.Player")
	proto.RegisterType((*State)(nil), "play.State")
	proto.RegisterType((*Resumed)(nil), "play.Resumed")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TurnClient is the client API for Turn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TurnClient interface {
	SetMove(ctx context.Context, in *Step, opts ...grpc.CallOption) (*State, error)
	GetMove(ctx context.Context, in *Player, opts ...grpc.CallOption) (*Step, error)
	UpdateNext(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error)
	IsNextPlayer(ctx context.Context, in *Player, opts ...grpc.CallOption) (*State, error)
	SetPlayer(ctx context.Context, in *Player, opts ...grpc.CallOption) (*State, error)
	GetAIPlayer(ctx context.Context, in *State, opts ...grpc.CallOption) (*Player, error)
	HasChosen(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error)
	HasMoved(ctx context.Context, in *Player, opts ...grpc.CallOption) (*State, error)
	SetResumed(ctx context.Context, in *Resumed, opts ...grpc.CallOption) (*State, error)
	GetResumed(ctx context.Context, in *State, opts ...grpc.CallOption) (*Resumed, error)
	NewRoom(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error)
	GetID(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error)
	SetExit(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error)
	CheckExit(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error)
	SetUserID(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error)
	GetUserID(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error)
}

type turnClient struct {
	cc *grpc.ClientConn
}

func NewTurnClient(cc *grpc.ClientConn) TurnClient {
	return &turnClient{cc}
}

func (c *turnClient) SetMove(ctx context.Context, in *Step, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/play.Turn/SetMove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turnClient) GetMove(ctx context.Context, in *Player, opts ...grpc.CallOption) (*Step, error) {
	out := new(Step)
	err := c.cc.Invoke(ctx, "/play.Turn/GetMove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turnClient) UpdateNext(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/play.Turn/UpdateNext", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turnClient) IsNextPlayer(ctx context.Context, in *Player, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/play.Turn/IsNextPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turnClient) SetPlayer(ctx context.Context, in *Player, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/play.Turn/SetPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turnClient) GetAIPlayer(ctx context.Context, in *State, opts ...grpc.CallOption) (*Player, error) {
	out := new(Player)
	err := c.cc.Invoke(ctx, "/play.Turn/GetAIPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turnClient) HasChosen(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/play.Turn/HasChosen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turnClient) HasMoved(ctx context.Context, in *Player, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/play.Turn/HasMoved", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turnClient) SetResumed(ctx context.Context, in *Resumed, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/play.Turn/SetResumed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turnClient) GetResumed(ctx context.Context, in *State, opts ...grpc.CallOption) (*Resumed, error) {
	out := new(Resumed)
	err := c.cc.Invoke(ctx, "/play.Turn/GetResumed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turnClient) NewRoom(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/play.Turn/NewRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turnClient) GetID(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/play.Turn/GetID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turnClient) SetExit(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/play.Turn/SetExit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turnClient) CheckExit(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/play.Turn/CheckExit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turnClient) SetUserID(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/play.Turn/SetUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turnClient) GetUserID(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/play.Turn/GetUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TurnServer is the server API for Turn service.
type TurnServer interface {
	SetMove(context.Context, *Step) (*State, error)
	GetMove(context.Context, *Player) (*Step, error)
	UpdateNext(context.Context, *State) (*State, error)
	IsNextPlayer(context.Context, *Player) (*State, error)
	SetPlayer(context.Context, *Player) (*State, error)
	GetAIPlayer(context.Context, *State) (*Player, error)
	HasChosen(context.Context, *State) (*State, error)
	HasMoved(context.Context, *Player) (*State, error)
	SetResumed(context.Context, *Resumed) (*State, error)
	GetResumed(context.Context, *State) (*Resumed, error)
	NewRoom(context.Context, *State) (*State, error)
	GetID(context.Context, *State) (*State, error)
	SetExit(context.Context, *State) (*State, error)
	CheckExit(context.Context, *State) (*State, error)
	SetUserID(context.Context, *State) (*State, error)
	GetUserID(context.Context, *State) (*State, error)
}

func RegisterTurnServer(s *grpc.Server, srv TurnServer) {
	s.RegisterService(&_Turn_serviceDesc, srv)
}

func _Turn_SetMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Step)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurnServer).SetMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/play.Turn/SetMove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurnServer).SetMove(ctx, req.(*Step))
	}
	return interceptor(ctx, in, info, handler)
}

func _Turn_GetMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Player)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurnServer).GetMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/play.Turn/GetMove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurnServer).GetMove(ctx, req.(*Player))
	}
	return interceptor(ctx, in, info, handler)
}

func _Turn_UpdateNext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(State)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurnServer).UpdateNext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/play.Turn/UpdateNext",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurnServer).UpdateNext(ctx, req.(*State))
	}
	return interceptor(ctx, in, info, handler)
}

func _Turn_IsNextPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Player)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurnServer).IsNextPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/play.Turn/IsNextPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurnServer).IsNextPlayer(ctx, req.(*Player))
	}
	return interceptor(ctx, in, info, handler)
}

func _Turn_SetPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Player)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurnServer).SetPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/play.Turn/SetPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurnServer).SetPlayer(ctx, req.(*Player))
	}
	return interceptor(ctx, in, info, handler)
}

func _Turn_GetAIPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(State)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurnServer).GetAIPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/play.Turn/GetAIPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurnServer).GetAIPlayer(ctx, req.(*State))
	}
	return interceptor(ctx, in, info, handler)
}

func _Turn_HasChosen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(State)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurnServer).HasChosen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/play.Turn/HasChosen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurnServer).HasChosen(ctx, req.(*State))
	}
	return interceptor(ctx, in, info, handler)
}

func _Turn_HasMoved_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Player)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurnServer).HasMoved(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/play.Turn/HasMoved",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurnServer).HasMoved(ctx, req.(*Player))
	}
	return interceptor(ctx, in, info, handler)
}

func _Turn_SetResumed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Resumed)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurnServer).SetResumed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/play.Turn/SetResumed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurnServer).SetResumed(ctx, req.(*Resumed))
	}
	return interceptor(ctx, in, info, handler)
}

func _Turn_GetResumed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(State)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurnServer).GetResumed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/play.Turn/GetResumed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurnServer).GetResumed(ctx, req.(*State))
	}
	return interceptor(ctx, in, info, handler)
}

func _Turn_NewRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(State)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurnServer).NewRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/play.Turn/NewRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurnServer).NewRoom(ctx, req.(*State))
	}
	return interceptor(ctx, in, info, handler)
}

func _Turn_GetID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(State)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurnServer).GetID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/play.Turn/GetID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurnServer).GetID(ctx, req.(*State))
	}
	return interceptor(ctx, in, info, handler)
}

func _Turn_SetExit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(State)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurnServer).SetExit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/play.Turn/SetExit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurnServer).SetExit(ctx, req.(*State))
	}
	return interceptor(ctx, in, info, handler)
}

func _Turn_CheckExit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(State)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurnServer).CheckExit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/play.Turn/CheckExit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurnServer).CheckExit(ctx, req.(*State))
	}
	return interceptor(ctx, in, info, handler)
}

func _Turn_SetUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(State)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurnServer).SetUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/play.Turn/SetUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurnServer).SetUserID(ctx, req.(*State))
	}
	return interceptor(ctx, in, info, handler)
}

func _Turn_GetUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(State)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurnServer).GetUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/play.Turn/GetUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurnServer).GetUserID(ctx, req.(*State))
	}
	return interceptor(ctx, in, info, handler)
}

var _Turn_serviceDesc = grpc.ServiceDesc{
	ServiceName: "play.Turn",
	HandlerType: (*TurnServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetMove",
			Handler:    _Turn_SetMove_Handler,
		},
		{
			MethodName: "GetMove",
			Handler:    _Turn_GetMove_Handler,
		},
		{
			MethodName: "UpdateNext",
			Handler:    _Turn_UpdateNext_Handler,
		},
		{
			MethodName: "IsNextPlayer",
			Handler:    _Turn_IsNextPlayer_Handler,
		},
		{
			MethodName: "SetPlayer",
			Handler:    _Turn_SetPlayer_Handler,
		},
		{
			MethodName: "GetAIPlayer",
			Handler:    _Turn_GetAIPlayer_Handler,
		},
		{
			MethodName: "HasChosen",
			Handler:    _Turn_HasChosen_Handler,
		},
		{
			MethodName: "HasMoved",
			Handler:    _Turn_HasMoved_Handler,
		},
		{
			MethodName: "SetResumed",
			Handler:    _Turn_SetResumed_Handler,
		},
		{
			MethodName: "GetResumed",
			Handler:    _Turn_GetResumed_Handler,
		},
		{
			MethodName: "NewRoom",
			Handler:    _Turn_NewRoom_Handler,
		},
		{
			MethodName: "GetID",
			Handler:    _Turn_GetID_Handler,
		},
		{
			MethodName: "SetExit",
			Handler:    _Turn_SetExit_Handler,
		},
		{
			MethodName: "CheckExit",
			Handler:    _Turn_CheckExit_Handler,
		},
		{
			MethodName: "SetUserID",
			Handler:    _Turn_SetUserID_Handler,
		},
		{
			MethodName: "GetUserID",
			Handler:    _Turn_GetUserID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "play.proto",
}

func init() { proto.RegisterFile("play.proto", fileDescriptor_play_97261441d1e200c2) }

var fileDescriptor_play_97261441d1e200c2 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x6e, 0xaa, 0x40,
	0x10, 0xc6, 0x45, 0x01, 0x65, 0xd4, 0x93, 0x9c, 0x4d, 0xd3, 0x10, 0xaf, 0x0c, 0xd5, 0x48, 0xda,
	0xd4, 0x0b, 0xfb, 0x04, 0x8d, 0x34, 0xc8, 0x45, 0x4d, 0xb3, 0xd4, 0x07, 0xa0, 0x3a, 0x89, 0x4d,
	0xd5, 0x25, 0xec, 0xd2, 0xe2, 0x9b, 0xf4, 0x71, 0x9b, 0x5d, 0x30, 0x15, 0xd2, 0xa8, 0x77, 0xf3,
	0xed, 0xfe, 0xe6, 0xdb, 0xf9, 0x03, 0x00, 0xf1, 0x26, 0xda, 0x8f, 0xe3, 0x84, 0x09, 0x46, 0x74,
	0x19, 0x3b, 0x33, 0xd0, 0x43, 0x81, 0x31, 0xe9, 0x80, 0x96, 0xd9, 0x5a, 0x5f, 0x73, 0xff, 0x53,
	0x2d, 0x93, 0x6a, 0x6f, 0xd7, 0x73, 0xb5, 0x27, 0x03, 0x30, 0x25, 0x8b, 0x89, 0xdd, 0xe8, 0x6b,
	0x6e, 0x7b, 0xd2, 0x19, 0x2b, 0x9b, 0x17, 0x75, 0x46, 0x8b, 0x3b, 0x67, 0x0c, 0x66, 0x7e, 0x42,
	0xae, 0xc0, 0x58, 0xb2, 0x0d, 0x4b, 0x94, 0x5f, 0x97, 0xe6, 0x82, 0xfc, 0x83, 0x7a, 0xe0, 0x29,
	0x53, 0x8b, 0xd6, 0x03, 0xcf, 0xf1, 0xc1, 0x08, 0x45, 0x24, 0x90, 0x5c, 0x83, 0xc9, 0x45, 0x24,
	0x52, 0xae, 0xf8, 0x16, 0x2d, 0x54, 0x35, 0x41, 0x72, 0x29, 0xc7, 0x24, 0xf0, 0x54, 0x19, 0x16,
	0x2d, 0x94, 0x73, 0x0f, 0x4d, 0x8a, 0x3c, 0xdd, 0xe2, 0x8a, 0x10, 0xd0, 0xb7, 0xec, 0x13, 0x6d,
	0xad, 0xdf, 0x70, 0x2d, 0xaa, 0xe2, 0xaa, 0xcd, 0xe4, 0xdb, 0x00, 0xfd, 0x35, 0x4d, 0x76, 0x64,
	0x00, 0xcd, 0x10, 0xc5, 0xb3, 0x64, 0x20, 0xef, 0x48, 0x4e, 0xa2, 0xd7, 0x3e, 0xc4, 0x91, 0x40,
	0xa7, 0x46, 0x86, 0xd0, 0xf4, 0x0b, 0xaa, 0xd4, 0x77, 0xef, 0x28, 0xc7, 0xa9, 0x11, 0x17, 0x60,
	0x11, 0xaf, 0x22, 0x81, 0x73, 0xcc, 0x04, 0x39, 0xf6, 0xa8, 0x1a, 0xde, 0x41, 0x27, 0xe0, 0x92,
	0x2a, 0xa6, 0x55, 0x76, 0xad, 0xc0, 0x2e, 0x58, 0x21, 0x5e, 0x44, 0xde, 0x42, 0xdb, 0x47, 0xf1,
	0x18, 0x14, 0x6c, 0xa9, 0x82, 0x52, 0xa2, 0x53, 0x23, 0x23, 0xb0, 0x66, 0x11, 0x9f, 0xae, 0x19,
	0xc7, 0xdd, 0xc9, 0x5a, 0x47, 0xd0, 0x9a, 0x45, 0x5c, 0x36, 0xbf, 0x3a, 0xf7, 0x3a, 0x84, 0x28,
	0x0e, 0x6b, 0xe8, 0xe6, 0x97, 0x85, 0xfc, 0x83, 0xf5, 0x7f, 0xd9, 0xd2, 0xf3, 0xe5, 0xc4, 0x7c,
	0xfa, 0x73, 0xfc, 0xa2, 0x8c, 0x6d, 0x4f, 0xd6, 0x79, 0x03, 0x86, 0x8f, 0x22, 0xf0, 0x4e, 0x42,
	0x43, 0xb5, 0xef, 0xa7, 0xec, 0x5d, 0x9c, 0xe9, 0xd9, 0x9a, 0xae, 0x71, 0xf9, 0x71, 0x09, 0x18,
	0xa2, 0x58, 0xa8, 0x8f, 0xf0, 0x1c, 0xe8, 0x5f, 0x02, 0xbe, 0x99, 0xea, 0xcf, 0x7c, 0xf8, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x5c, 0x2c, 0x63, 0x21, 0xa7, 0x03, 0x00, 0x00,
}

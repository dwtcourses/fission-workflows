// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/scheduler/scheduler.proto

/*
Package scheduler is a generated protocol buffer package.

It is generated from these files:
	pkg/scheduler/scheduler.proto

It has these top-level messages:
	Schedule
	AbortAction
	RunTaskAction
	PrepareTaskAction
*/
package scheduler

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import fission_workflows_types1 "github.com/fission/fission-workflows/pkg/types"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

type Schedule struct {
	InvocationId string                     `protobuf:"bytes,1,opt,name=invocationId" json:"invocationId,omitempty"`
	CreatedAt    *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=createdAt" json:"createdAt,omitempty"`
	Abort        *AbortAction               `protobuf:"bytes,4,opt,name=abort" json:"abort,omitempty"`
	RunTasks     []*RunTaskAction           `protobuf:"bytes,5,rep,name=runTasks" json:"runTasks,omitempty"`
	PrepareTasks []*PrepareTaskAction       `protobuf:"bytes,6,rep,name=prepareTasks" json:"prepareTasks,omitempty"`
}

func (m *Schedule) Reset()                    { *m = Schedule{} }
func (m *Schedule) String() string            { return proto.CompactTextString(m) }
func (*Schedule) ProtoMessage()               {}
func (*Schedule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Schedule) GetInvocationId() string {
	if m != nil {
		return m.InvocationId
	}
	return ""
}

func (m *Schedule) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Schedule) GetAbort() *AbortAction {
	if m != nil {
		return m.Abort
	}
	return nil
}

func (m *Schedule) GetRunTasks() []*RunTaskAction {
	if m != nil {
		return m.RunTasks
	}
	return nil
}

func (m *Schedule) GetPrepareTasks() []*PrepareTaskAction {
	if m != nil {
		return m.PrepareTasks
	}
	return nil
}

type AbortAction struct {
	Reason string `protobuf:"bytes,1,opt,name=reason" json:"reason,omitempty"`
}

func (m *AbortAction) Reset()                    { *m = AbortAction{} }
func (m *AbortAction) String() string            { return proto.CompactTextString(m) }
func (*AbortAction) ProtoMessage()               {}
func (*AbortAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AbortAction) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type RunTaskAction struct {
	// Id of the task in the workflow
	TaskID string `protobuf:"bytes,1,opt,name=taskID" json:"taskID,omitempty"`
}

func (m *RunTaskAction) Reset()                    { *m = RunTaskAction{} }
func (m *RunTaskAction) String() string            { return proto.CompactTextString(m) }
func (*RunTaskAction) ProtoMessage()               {}
func (*RunTaskAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RunTaskAction) GetTaskID() string {
	if m != nil {
		return m.TaskID
	}
	return ""
}

type PrepareTaskAction struct {
	TaskID     string                     `protobuf:"bytes,1,opt,name=taskID" json:"taskID,omitempty"`
	ExpectedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=expectedAt" json:"expectedAt,omitempty"`
}

func (m *PrepareTaskAction) Reset()                    { *m = PrepareTaskAction{} }
func (m *PrepareTaskAction) String() string            { return proto.CompactTextString(m) }
func (*PrepareTaskAction) ProtoMessage()               {}
func (*PrepareTaskAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PrepareTaskAction) GetTaskID() string {
	if m != nil {
		return m.TaskID
	}
	return ""
}

func (m *PrepareTaskAction) GetExpectedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpectedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Schedule)(nil), "fission.workflows.scheduler.Schedule")
	proto.RegisterType((*AbortAction)(nil), "fission.workflows.scheduler.AbortAction")
	proto.RegisterType((*RunTaskAction)(nil), "fission.workflows.scheduler.RunTaskAction")
	proto.RegisterType((*PrepareTaskAction)(nil), "fission.workflows.scheduler.PrepareTaskAction")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Scheduler service

type SchedulerClient interface {
	Evaluate(ctx context.Context, in *fission_workflows_types1.WorkflowInvocation, opts ...grpc.CallOption) (*Schedule, error)
}

type schedulerClient struct {
	cc *grpc.ClientConn
}

func NewSchedulerClient(cc *grpc.ClientConn) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) Evaluate(ctx context.Context, in *fission_workflows_types1.WorkflowInvocation, opts ...grpc.CallOption) (*Schedule, error) {
	out := new(Schedule)
	err := grpc.Invoke(ctx, "/fission.workflows.scheduler.Scheduler/evaluate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Scheduler service

type SchedulerServer interface {
	Evaluate(context.Context, *fission_workflows_types1.WorkflowInvocation) (*Schedule, error)
}

func RegisterSchedulerServer(s *grpc.Server, srv SchedulerServer) {
	s.RegisterService(&_Scheduler_serviceDesc, srv)
}

func _Scheduler_Evaluate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types1.WorkflowInvocation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).Evaluate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.scheduler.Scheduler/Evaluate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).Evaluate(ctx, req.(*fission_workflows_types1.WorkflowInvocation))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fission.workflows.scheduler.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "evaluate",
			Handler:    _Scheduler_Evaluate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/scheduler/scheduler.proto",
}

func init() { proto.RegisterFile("pkg/scheduler/scheduler.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4b, 0xeb, 0x40,
	0x14, 0xc5, 0x5f, 0xfb, 0x5e, 0x4b, 0x73, 0xd3, 0xb7, 0x70, 0x16, 0x12, 0x22, 0x62, 0x09, 0x14,
	0x83, 0xe2, 0x04, 0xea, 0x46, 0xba, 0x10, 0x2a, 0x22, 0x74, 0x27, 0xb1, 0x20, 0xb8, 0x72, 0x92,
	0x4e, 0xd3, 0x90, 0x3f, 0x13, 0x66, 0x26, 0xad, 0x7e, 0x3f, 0x3f, 0x98, 0x24, 0x93, 0xa4, 0x2d,
	0x6a, 0x70, 0x93, 0xe4, 0x0e, 0xbf, 0x73, 0xce, 0xcd, 0x9d, 0x0b, 0xa7, 0x59, 0x14, 0x38, 0xc2,
	0x5f, 0xd3, 0x65, 0x1e, 0x53, 0xbe, 0xfb, 0xc2, 0x19, 0x67, 0x92, 0xa1, 0x93, 0x55, 0x28, 0x44,
	0xc8, 0x52, 0xbc, 0x65, 0x3c, 0x5a, 0xc5, 0x6c, 0x2b, 0x70, 0x83, 0x98, 0xd3, 0x20, 0x94, 0xeb,
	0xdc, 0xc3, 0x3e, 0x4b, 0x9c, 0x8a, 0xab, 0xdf, 0x57, 0x0d, 0xef, 0x14, 0x01, 0xf2, 0x3d, 0xa3,
	0x42, 0x3d, 0x95, 0xb1, 0x79, 0x16, 0x30, 0x16, 0xc4, 0xd4, 0x29, 0x2b, 0x2f, 0x5f, 0x39, 0x32,
	0x4c, 0xa8, 0x90, 0x24, 0xc9, 0x14, 0x60, 0x7d, 0x74, 0x61, 0xf0, 0x54, 0x45, 0x21, 0x0b, 0x86,
	0x61, 0xba, 0x61, 0x3e, 0x91, 0x21, 0x4b, 0xe7, 0x4b, 0xa3, 0x33, 0xea, 0xd8, 0x9a, 0x7b, 0x70,
	0x86, 0x6e, 0x40, 0xf3, 0x39, 0x25, 0x92, 0x2e, 0x67, 0xd2, 0xe8, 0x8e, 0x3a, 0xb6, 0x3e, 0x31,
	0xb1, 0x4a, 0xc1, 0x75, 0x0a, 0x5e, 0xd4, 0x29, 0xee, 0x0e, 0x46, 0xb7, 0xd0, 0x23, 0x1e, 0xe3,
	0xd2, 0xf8, 0x57, 0xaa, 0x6c, 0xdc, 0xf2, 0xd3, 0x78, 0x56, 0x90, 0x33, 0xbf, 0x08, 0x75, 0x95,
	0x0c, 0x3d, 0xc0, 0x80, 0xe7, 0xe9, 0x82, 0x88, 0x48, 0x18, 0xbd, 0xd1, 0x5f, 0x5b, 0x9f, 0x5c,
	0xb4, 0x5a, 0xb8, 0x0a, 0xae, 0x4c, 0x1a, 0x2d, 0x72, 0x61, 0x98, 0x71, 0x9a, 0x11, 0x4e, 0x95,
	0x57, 0xbf, 0xf4, 0xc2, 0xad, 0x5e, 0x8f, 0x3b, 0x41, 0xe5, 0x77, 0xe0, 0x61, 0x8d, 0x41, 0xdf,
	0xeb, 0x18, 0x1d, 0x43, 0x9f, 0x53, 0x22, 0x58, 0x5a, 0x8d, 0xb0, 0xaa, 0xac, 0x73, 0xf8, 0x7f,
	0xd0, 0x55, 0x01, 0x4a, 0x22, 0xa2, 0xf9, 0x7d, 0x0d, 0xaa, 0xca, 0x0a, 0xe0, 0xe8, 0x4b, 0xe4,
	0x4f, 0x30, 0x9a, 0x02, 0xd0, 0xb7, 0x8c, 0xfa, 0xbf, 0xbd, 0x93, 0x3d, 0x7a, 0x92, 0x80, 0x56,
	0x5f, 0x3f, 0x47, 0xaf, 0x30, 0xa0, 0x1b, 0x12, 0xe7, 0x44, 0x52, 0x74, 0xf9, 0xcd, 0x3c, 0xd4,
	0x66, 0x3d, 0x57, 0xf5, 0xbc, 0x59, 0x0b, 0x73, 0xdc, 0x3a, 0xbc, 0x3a, 0xc0, 0xfa, 0x73, 0xa7,
	0xbf, 0x68, 0xcd, 0xb9, 0xd7, 0x2f, 0x7b, 0xbb, 0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x32,
	0xe7, 0xf1, 0x1d, 0x03, 0x00, 0x00,
}

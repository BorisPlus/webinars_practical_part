// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: elections.proto

package main

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Vote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Passport    string               `protobuf:"bytes,1,opt,name=passport,proto3" json:"passport,omitempty"`
	CandidateId uint32               `protobuf:"varint,2,opt,name=candidate_id,json=candidateId,proto3" json:"candidate_id,omitempty"`
	Note        string               `protobuf:"bytes,3,opt,name=note,proto3" json:"note,omitempty"`
	Time        *timestamp.Timestamp `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Vote) Reset() {
	*x = Vote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elections_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vote) ProtoMessage() {}

func (x *Vote) ProtoReflect() protoreflect.Message {
	mi := &file_elections_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vote.ProtoReflect.Descriptor instead.
func (*Vote) Descriptor() ([]byte, []int) {
	return file_elections_proto_rawDescGZIP(), []int{0}
}

func (x *Vote) GetPassport() string {
	if x != nil {
		return x.Passport
	}
	return ""
}

func (x *Vote) GetCandidateId() uint32 {
	if x != nil {
		return x.CandidateId
	}
	return 0
}

func (x *Vote) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *Vote) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records map[uint32]uint32    `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Time    *timestamp.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Stats) Reset() {
	*x = Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elections_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats) ProtoMessage() {}

func (x *Stats) ProtoReflect() protoreflect.Message {
	mi := &file_elections_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats.ProtoReflect.Descriptor instead.
func (*Stats) Descriptor() ([]byte, []int) {
	return file_elections_proto_rawDescGZIP(), []int{1}
}

func (x *Stats) GetRecords() map[uint32]uint32 {
	if x != nil {
		return x.Records
	}
	return nil
}

func (x *Stats) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type StatsVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Body:
	//	*StatsVote_Stats
	//	*StatsVote_Vote
	Body isStatsVote_Body `protobuf_oneof:"body"`
}

func (x *StatsVote) Reset() {
	*x = StatsVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elections_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsVote) ProtoMessage() {}

func (x *StatsVote) ProtoReflect() protoreflect.Message {
	mi := &file_elections_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsVote.ProtoReflect.Descriptor instead.
func (*StatsVote) Descriptor() ([]byte, []int) {
	return file_elections_proto_rawDescGZIP(), []int{2}
}

func (m *StatsVote) GetBody() isStatsVote_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *StatsVote) GetStats() *Stats {
	if x, ok := x.GetBody().(*StatsVote_Stats); ok {
		return x.Stats
	}
	return nil
}

func (x *StatsVote) GetVote() *Vote {
	if x, ok := x.GetBody().(*StatsVote_Vote); ok {
		return x.Vote
	}
	return nil
}

type isStatsVote_Body interface {
	isStatsVote_Body()
}

type StatsVote_Stats struct {
	Stats *Stats `protobuf:"bytes,1,opt,name=stats,proto3,oneof"`
}

type StatsVote_Vote struct {
	Vote *Vote `protobuf:"bytes,2,opt,name=vote,proto3,oneof"`
}

func (*StatsVote_Stats) isStatsVote_Body() {}

func (*StatsVote_Vote) isStatsVote_Body() {}

var File_elections_proto protoreflect.FileDescriptor

var file_elections_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x77, 0x69, 0x74,
	0x68, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61,
	0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x22, 0xb7, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x42, 0x0a, 0x07, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x1a,
	0x3a, 0x0a, 0x0c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7a, 0x0a, 0x09, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x30, 0x0a,
	0x04, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x48, 0x00, 0x52, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x42,
	0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x32, 0x9e, 0x01, 0x0a, 0x09, 0x45, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x42, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x56,
	0x6f, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f,
	0x77, 0x69, 0x74, 0x68, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x08, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x1a, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x56, 0x6f, 0x74,
	0x65, 0x1a, 0x1f, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x77, 0x69,
	0x74, 0x68, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x56, 0x6f,
	0x74, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x6d, 0x61,
	0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_elections_proto_rawDescOnce sync.Once
	file_elections_proto_rawDescData = file_elections_proto_rawDesc
)

func file_elections_proto_rawDescGZIP() []byte {
	file_elections_proto_rawDescOnce.Do(func() {
		file_elections_proto_rawDescData = protoimpl.X.CompressGZIP(file_elections_proto_rawDescData)
	})
	return file_elections_proto_rawDescData
}

var file_elections_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_elections_proto_goTypes = []interface{}{
	(*Vote)(nil),                // 0: elections_with_admin.Vote
	(*Stats)(nil),               // 1: elections_with_admin.Stats
	(*StatsVote)(nil),           // 2: elections_with_admin.StatsVote
	nil,                         // 3: elections_with_admin.Stats.RecordsEntry
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*empty.Empty)(nil),         // 5: google.protobuf.Empty
}
var file_elections_proto_depIdxs = []int32{
	4, // 0: elections_with_admin.Vote.time:type_name -> google.protobuf.Timestamp
	3, // 1: elections_with_admin.Stats.records:type_name -> elections_with_admin.Stats.RecordsEntry
	4, // 2: elections_with_admin.Stats.time:type_name -> google.protobuf.Timestamp
	1, // 3: elections_with_admin.StatsVote.stats:type_name -> elections_with_admin.Stats
	0, // 4: elections_with_admin.StatsVote.vote:type_name -> elections_with_admin.Vote
	0, // 5: elections_with_admin.Elections.SubmitVote:input_type -> elections_with_admin.Vote
	0, // 6: elections_with_admin.Elections.Internal:input_type -> elections_with_admin.Vote
	5, // 7: elections_with_admin.Elections.SubmitVote:output_type -> google.protobuf.Empty
	2, // 8: elections_with_admin.Elections.Internal:output_type -> elections_with_admin.StatsVote
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_elections_proto_init() }
func file_elections_proto_init() {
	if File_elections_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_elections_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vote); i {
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
		file_elections_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stats); i {
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
		file_elections_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsVote); i {
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
	file_elections_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*StatsVote_Stats)(nil),
		(*StatsVote_Vote)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_elections_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_elections_proto_goTypes,
		DependencyIndexes: file_elections_proto_depIdxs,
		MessageInfos:      file_elections_proto_msgTypes,
	}.Build()
	File_elections_proto = out.File
	file_elections_proto_rawDesc = nil
	file_elections_proto_goTypes = nil
	file_elections_proto_depIdxs = nil
}

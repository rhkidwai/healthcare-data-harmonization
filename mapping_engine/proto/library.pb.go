// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: proto/library.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The LibraryConfig message defines a set of standalone projectors and Google
// cloud functions that can be loaded along-side and shared between mappings.
type LibraryConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Basic projectors that can be shared across various mappings.
	Projector []*ProjectorDefinition `protobuf:"bytes,1,rep,name=projector,proto3" json:"projector,omitempty"`
	// The remote servers that the configuration can call.
	Servers []*ServerDefinition `protobuf:"bytes,2,rep,name=servers,proto3" json:"servers,omitempty"`
	// The CloudFunction message defines a cloud function that can be called
	// as a projector.
	CloudFunction []*CloudFunction `protobuf:"bytes,3,rep,name=cloud_function,json=cloudFunction,proto3" json:"cloud_function,omitempty"`
	// HttpQuery defines a list of fetch queries, which are essentially projectors
	// with external lookup. All fetch query configs are expected to contain a
	// name field to be used as the projector name.
	HttpQuery []*HttpFetchQuery `protobuf:"bytes,4,rep,name=http_query,json=httpQuery,proto3" json:"http_query,omitempty"`
	// DEPRECATED: Builtin projector libraries.
	//
	// Deprecated: Do not use.
	DeprecateImportedLibraries []string `protobuf:"bytes,5,rep,name=deprecate_imported_libraries,json=deprecateImportedLibraries,proto3" json:"deprecate_imported_libraries,omitempty"`
	// User-defined projector libraries that will be included in the mappings.
	UserLibraries []*UserLibrary `protobuf:"bytes,6,rep,name=user_libraries,json=userLibraries,proto3" json:"user_libraries,omitempty"`
}

func (x *LibraryConfig) Reset() {
	*x = LibraryConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_library_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LibraryConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LibraryConfig) ProtoMessage() {}

func (x *LibraryConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_library_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LibraryConfig.ProtoReflect.Descriptor instead.
func (*LibraryConfig) Descriptor() ([]byte, []int) {
	return file_proto_library_proto_rawDescGZIP(), []int{0}
}

func (x *LibraryConfig) GetProjector() []*ProjectorDefinition {
	if x != nil {
		return x.Projector
	}
	return nil
}

func (x *LibraryConfig) GetServers() []*ServerDefinition {
	if x != nil {
		return x.Servers
	}
	return nil
}

func (x *LibraryConfig) GetCloudFunction() []*CloudFunction {
	if x != nil {
		return x.CloudFunction
	}
	return nil
}

func (x *LibraryConfig) GetHttpQuery() []*HttpFetchQuery {
	if x != nil {
		return x.HttpQuery
	}
	return nil
}

// Deprecated: Do not use.
func (x *LibraryConfig) GetDeprecateImportedLibraries() []string {
	if x != nil {
		return x.DeprecateImportedLibraries
	}
	return nil
}

func (x *LibraryConfig) GetUserLibraries() []*UserLibrary {
	if x != nil {
		return x.UserLibraries
	}
	return nil
}

// Specifies a user-defined projector library.
type UserLibrary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of mapping.
	Type MappingType `protobuf:"varint,1,opt,name=type,proto3,enum=cloud.healthcare.cdw.etl.mapping.proto.MappingType" json:"type,omitempty"`
	// The path defining the custom mapping library file.
	Path *Location `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *UserLibrary) Reset() {
	*x = UserLibrary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_library_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLibrary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLibrary) ProtoMessage() {}

func (x *UserLibrary) ProtoReflect() protoreflect.Message {
	mi := &file_proto_library_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLibrary.ProtoReflect.Descriptor instead.
func (*UserLibrary) Descriptor() ([]byte, []int) {
	return file_proto_library_proto_rawDescGZIP(), []int{1}
}

func (x *UserLibrary) GetType() MappingType {
	if x != nil {
		return x.Type
	}
	return MappingType_INVALID
}

func (x *UserLibrary) GetPath() *Location {
	if x != nil {
		return x.Path
	}
	return nil
}

var File_proto_library_proto protoreflect.FileDescriptor

var file_proto_library_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x63, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x64, 0x77, 0x2e, 0x65, 0x74, 0x6c, 0x2e,
	0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x95, 0x04, 0x0a, 0x0d, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x59, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x63, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x64, 0x77, 0x2e, 0x65, 0x74, 0x6c, 0x2e,
	0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x52, 0x0a, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x61, 0x72, 0x65,
	0x2e, 0x63, 0x64, 0x77, 0x2e, 0x65, 0x74, 0x6c, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73,
	0x12, 0x5c, 0x0a, 0x0e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x64, 0x77, 0x2e,
	0x65, 0x74, 0x6c, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x55,
	0x0a, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x63, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x64, 0x77, 0x2e, 0x65, 0x74, 0x6c, 0x2e, 0x6d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x74, 0x74, 0x70,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x09, 0x68, 0x74, 0x74, 0x70,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x44, 0x0a, 0x1c, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61,
	0x74, 0x65, 0x5f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52,
	0x1a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x5a, 0x0a, 0x0e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x63, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x64, 0x77, 0x2e, 0x65, 0x74, 0x6c, 0x2e, 0x6d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x22, 0x9c, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x47, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x63, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x64, 0x77, 0x2e, 0x65, 0x74, 0x6c,
	0x2e, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x44, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x61, 0x72,
	0x65, 0x2e, 0x63, 0x64, 0x77, 0x2e, 0x65, 0x74, 0x6c, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x42, 0x8c, 0x01, 0x0a, 0x35, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x63, 0x61, 0x72, 0x65, 0x2e, 0x65, 0x74, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x50, 0x01, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x61, 0x72, 0x65, 0x2d, 0x64, 0x61,
	0x74, 0x61, 0x2d, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_library_proto_rawDescOnce sync.Once
	file_proto_library_proto_rawDescData = file_proto_library_proto_rawDesc
)

func file_proto_library_proto_rawDescGZIP() []byte {
	file_proto_library_proto_rawDescOnce.Do(func() {
		file_proto_library_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_library_proto_rawDescData)
	})
	return file_proto_library_proto_rawDescData
}

var file_proto_library_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_library_proto_goTypes = []interface{}{
	(*LibraryConfig)(nil),       // 0: cloud.healthcare.cdw.etl.mapping.proto.LibraryConfig
	(*UserLibrary)(nil),         // 1: cloud.healthcare.cdw.etl.mapping.proto.UserLibrary
	(*ProjectorDefinition)(nil), // 2: cloud.healthcare.cdw.etl.mapping.proto.ProjectorDefinition
	(*ServerDefinition)(nil),    // 3: cloud.healthcare.cdw.etl.mapping.proto.ServerDefinition
	(*CloudFunction)(nil),       // 4: cloud.healthcare.cdw.etl.mapping.proto.CloudFunction
	(*HttpFetchQuery)(nil),      // 5: cloud.healthcare.cdw.etl.mapping.proto.HttpFetchQuery
	(MappingType)(0),            // 6: cloud.healthcare.cdw.etl.mapping.proto.MappingType
	(*Location)(nil),            // 7: cloud.healthcare.cdw.etl.mapping.proto.Location
}
var file_proto_library_proto_depIdxs = []int32{
	2, // 0: cloud.healthcare.cdw.etl.mapping.proto.LibraryConfig.projector:type_name -> cloud.healthcare.cdw.etl.mapping.proto.ProjectorDefinition
	3, // 1: cloud.healthcare.cdw.etl.mapping.proto.LibraryConfig.servers:type_name -> cloud.healthcare.cdw.etl.mapping.proto.ServerDefinition
	4, // 2: cloud.healthcare.cdw.etl.mapping.proto.LibraryConfig.cloud_function:type_name -> cloud.healthcare.cdw.etl.mapping.proto.CloudFunction
	5, // 3: cloud.healthcare.cdw.etl.mapping.proto.LibraryConfig.http_query:type_name -> cloud.healthcare.cdw.etl.mapping.proto.HttpFetchQuery
	1, // 4: cloud.healthcare.cdw.etl.mapping.proto.LibraryConfig.user_libraries:type_name -> cloud.healthcare.cdw.etl.mapping.proto.UserLibrary
	6, // 5: cloud.healthcare.cdw.etl.mapping.proto.UserLibrary.type:type_name -> cloud.healthcare.cdw.etl.mapping.proto.MappingType
	7, // 6: cloud.healthcare.cdw.etl.mapping.proto.UserLibrary.path:type_name -> cloud.healthcare.cdw.etl.mapping.proto.Location
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_library_proto_init() }
func file_proto_library_proto_init() {
	if File_proto_library_proto != nil {
		return
	}
	file_proto_harmonization_proto_init()
	file_proto_http_proto_init()
	file_proto_mapping_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_library_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LibraryConfig); i {
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
		file_proto_library_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLibrary); i {
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
			RawDescriptor: file_proto_library_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_library_proto_goTypes,
		DependencyIndexes: file_proto_library_proto_depIdxs,
		MessageInfos:      file_proto_library_proto_msgTypes,
	}.Build()
	File_proto_library_proto = out.File
	file_proto_library_proto_rawDesc = nil
	file_proto_library_proto_goTypes = nil
	file_proto_library_proto_depIdxs = nil
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: ChemistryBackbone.proto

package ChemistryBackbone

import (
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

type EquationSection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Equations []*Equation `protobuf:"bytes,3,rep,name=equations,proto3" json:"equations,omitempty"`
}

func (x *EquationSection) Reset() {
	*x = EquationSection{}
	mi := &file_ChemistryBackbone_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EquationSection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EquationSection) ProtoMessage() {}

func (x *EquationSection) ProtoReflect() protoreflect.Message {
	mi := &file_ChemistryBackbone_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EquationSection.ProtoReflect.Descriptor instead.
func (*EquationSection) Descriptor() ([]byte, []int) {
	return file_ChemistryBackbone_proto_rawDescGZIP(), []int{0}
}

func (x *EquationSection) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EquationSection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EquationSection) GetEquations() []*Equation {
	if x != nil {
		return x.Equations
	}
	return nil
}

type Equation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Filters     []string `protobuf:"bytes,4,rep,name=filters,proto3" json:"filters,omitempty"`
	FieldLabels []string `protobuf:"bytes,5,rep,name=fieldLabels,proto3" json:"fieldLabels,omitempty"`
}

func (x *Equation) Reset() {
	*x = Equation{}
	mi := &file_ChemistryBackbone_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Equation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Equation) ProtoMessage() {}

func (x *Equation) ProtoReflect() protoreflect.Message {
	mi := &file_ChemistryBackbone_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Equation.ProtoReflect.Descriptor instead.
func (*Equation) Descriptor() ([]byte, []int) {
	return file_ChemistryBackbone_proto_rawDescGZIP(), []int{1}
}

func (x *Equation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Equation) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Equation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Equation) GetFilters() []string {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *Equation) GetFieldLabels() []string {
	if x != nil {
		return x.FieldLabels
	}
	return nil
}

type EquationSectionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sections []*EquationSection `protobuf:"bytes,1,rep,name=sections,proto3" json:"sections,omitempty"`
}

func (x *EquationSectionList) Reset() {
	*x = EquationSectionList{}
	mi := &file_ChemistryBackbone_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EquationSectionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EquationSectionList) ProtoMessage() {}

func (x *EquationSectionList) ProtoReflect() protoreflect.Message {
	mi := &file_ChemistryBackbone_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EquationSectionList.ProtoReflect.Descriptor instead.
func (*EquationSectionList) Descriptor() ([]byte, []int) {
	return file_ChemistryBackbone_proto_rawDescGZIP(), []int{2}
}

func (x *EquationSectionList) GetSections() []*EquationSection {
	if x != nil {
		return x.Sections
	}
	return nil
}

var File_ChemistryBackbone_proto protoreflect.FileDescriptor

var file_ChemistryBackbone_proto_rawDesc = []byte{
	0x0a, 0x17, 0x43, 0x68, 0x65, 0x6d, 0x69, 0x73, 0x74, 0x72, 0x79, 0x42, 0x61, 0x63, 0x6b, 0x62,
	0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x43, 0x68, 0x65, 0x6d, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x42, 0x61, 0x63, 0x6b, 0x62, 0x6f, 0x6e, 0x65, 0x22, 0x70, 0x0a, 0x0f,
	0x45, 0x71, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x65, 0x71, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x43, 0x68, 0x65, 0x6d, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x42, 0x61, 0x63, 0x6b, 0x62, 0x6f, 0x6e, 0x65, 0x2e, 0x45, 0x71, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x65, 0x71, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x8e,
	0x01, 0x0a, 0x08, 0x45, 0x71, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22,
	0x55, 0x0a, 0x13, 0x45, 0x71, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x43, 0x68, 0x65, 0x6d, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x42, 0x61, 0x63, 0x6b, 0x62, 0x6f, 0x6e, 0x65, 0x2e, 0x45, 0x71, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x43, 0x68, 0x65, 0x6d,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x42, 0x61, 0x63, 0x6b, 0x62, 0x6f, 0x6e, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChemistryBackbone_proto_rawDescOnce sync.Once
	file_ChemistryBackbone_proto_rawDescData = file_ChemistryBackbone_proto_rawDesc
)

func file_ChemistryBackbone_proto_rawDescGZIP() []byte {
	file_ChemistryBackbone_proto_rawDescOnce.Do(func() {
		file_ChemistryBackbone_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChemistryBackbone_proto_rawDescData)
	})
	return file_ChemistryBackbone_proto_rawDescData
}

var file_ChemistryBackbone_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ChemistryBackbone_proto_goTypes = []any{
	(*EquationSection)(nil),     // 0: ChemistryBackbone.EquationSection
	(*Equation)(nil),            // 1: ChemistryBackbone.Equation
	(*EquationSectionList)(nil), // 2: ChemistryBackbone.EquationSectionList
}
var file_ChemistryBackbone_proto_depIdxs = []int32{
	1, // 0: ChemistryBackbone.EquationSection.equations:type_name -> ChemistryBackbone.Equation
	0, // 1: ChemistryBackbone.EquationSectionList.sections:type_name -> ChemistryBackbone.EquationSection
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ChemistryBackbone_proto_init() }
func file_ChemistryBackbone_proto_init() {
	if File_ChemistryBackbone_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ChemistryBackbone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChemistryBackbone_proto_goTypes,
		DependencyIndexes: file_ChemistryBackbone_proto_depIdxs,
		MessageInfos:      file_ChemistryBackbone_proto_msgTypes,
	}.Build()
	File_ChemistryBackbone_proto = out.File
	file_ChemistryBackbone_proto_rawDesc = nil
	file_ChemistryBackbone_proto_goTypes = nil
	file_ChemistryBackbone_proto_depIdxs = nil
}

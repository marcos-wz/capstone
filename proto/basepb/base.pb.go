// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: base.proto

package basepb

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

type Unit int32

const (
	Unit_UNIT_UNDEFINED Unit = 0
	Unit_UNIT_KG        Unit = 1
	Unit_UNIT_LB        Unit = 2
)

// Enum value maps for Unit.
var (
	Unit_name = map[int32]string{
		0: "UNIT_UNDEFINED",
		1: "UNIT_KG",
		2: "UNIT_LB",
	}
	Unit_value = map[string]int32{
		"UNIT_UNDEFINED": 0,
		"UNIT_KG":        1,
		"UNIT_LB":        2,
	}
)

func (x Unit) Enum() *Unit {
	p := new(Unit)
	*p = x
	return p
}

func (x Unit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Unit) Descriptor() protoreflect.EnumDescriptor {
	return file_base_proto_enumTypes[0].Descriptor()
}

func (Unit) Type() protoreflect.EnumType {
	return &file_base_proto_enumTypes[0]
}

func (x Unit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Unit.Descriptor instead.
func (Unit) EnumDescriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

type Country int32

const (
	Country_COUNTRY_UNDEFINED Country = 0
	Country_COUNTRY_MEXICO    Country = 1
	Country_COUNTRY_BRAZIL    Country = 2
	Country_COUNTRY_CANADA    Country = 3
	Country_COUNTRY_USA       Country = 4
)

// Enum value maps for Country.
var (
	Country_name = map[int32]string{
		0: "COUNTRY_UNDEFINED",
		1: "COUNTRY_MEXICO",
		2: "COUNTRY_BRAZIL",
		3: "COUNTRY_CANADA",
		4: "COUNTRY_USA",
	}
	Country_value = map[string]int32{
		"COUNTRY_UNDEFINED": 0,
		"COUNTRY_MEXICO":    1,
		"COUNTRY_BRAZIL":    2,
		"COUNTRY_CANADA":    3,
		"COUNTRY_USA":       4,
	}
)

func (x Country) Enum() *Country {
	p := new(Country)
	*p = x
	return p
}

func (x Country) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Country) Descriptor() protoreflect.EnumDescriptor {
	return file_base_proto_enumTypes[1].Descriptor()
}

func (Country) Type() protoreflect.EnumType {
	return &file_base_proto_enumTypes[1]
}

func (x Country) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Country.Descriptor instead.
func (Country) EnumDescriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{1}
}

type Currency int32

const (
	Currency_CURRENCY_UNDEFINED Currency = 0
	Currency_CURRENCY_MXN       Currency = 1
	Currency_CURRENCY_BRL       Currency = 2
	Currency_CURRENCY_CAD       Currency = 3
	Currency_CURRENCY_USD       Currency = 4
)

// Enum value maps for Currency.
var (
	Currency_name = map[int32]string{
		0: "CURRENCY_UNDEFINED",
		1: "CURRENCY_MXN",
		2: "CURRENCY_BRL",
		3: "CURRENCY_CAD",
		4: "CURRENCY_USD",
	}
	Currency_value = map[string]int32{
		"CURRENCY_UNDEFINED": 0,
		"CURRENCY_MXN":       1,
		"CURRENCY_BRL":       2,
		"CURRENCY_CAD":       3,
		"CURRENCY_USD":       4,
	}
)

func (x Currency) Enum() *Currency {
	p := new(Currency)
	*p = x
	return p
}

func (x Currency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Currency) Descriptor() protoreflect.EnumDescriptor {
	return file_base_proto_enumTypes[2].Descriptor()
}

func (Currency) Type() protoreflect.EnumType {
	return &file_base_proto_enumTypes[2]
}

func (x Currency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Currency.Descriptor instead.
func (Currency) EnumDescriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{2}
}

type Fruit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description  string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Color        string   `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
	Unit         Unit     `protobuf:"varint,5,opt,name=unit,proto3,enum=capstone.Unit" json:"unit,omitempty"`
	Price        float32  `protobuf:"fixed32,6,opt,name=price,proto3" json:"price,omitempty"`
	Currency     Currency `protobuf:"varint,7,opt,name=currency,proto3,enum=capstone.Currency" json:"currency,omitempty"`
	Stock        uint32   `protobuf:"varint,8,opt,name=stock,proto3" json:"stock,omitempty"`
	CaducateDays uint32   `protobuf:"varint,9,opt,name=caducate_days,json=caducateDays,proto3" json:"caducate_days,omitempty"`
	Country      Country  `protobuf:"varint,10,opt,name=country,proto3,enum=capstone.Country" json:"country,omitempty"`
	// Create time time in microseconds (µs).
	CreateTime uint64 `protobuf:"varint,11,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Updated time time in microseconds (µs).
	UpdateTime uint64 `protobuf:"varint,12,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Fruit) Reset() {
	*x = Fruit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fruit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fruit) ProtoMessage() {}

func (x *Fruit) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fruit.ProtoReflect.Descriptor instead.
func (*Fruit) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

func (x *Fruit) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Fruit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Fruit) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Fruit) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Fruit) GetUnit() Unit {
	if x != nil {
		return x.Unit
	}
	return Unit_UNIT_UNDEFINED
}

func (x *Fruit) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Fruit) GetCurrency() Currency {
	if x != nil {
		return x.Currency
	}
	return Currency_CURRENCY_UNDEFINED
}

func (x *Fruit) GetStock() uint32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *Fruit) GetCaducateDays() uint32 {
	if x != nil {
		return x.CaducateDays
	}
	return 0
}

func (x *Fruit) GetCountry() Country {
	if x != nil {
		return x.Country
	}
	return Country_COUNTRY_UNDEFINED
}

func (x *Fruit) GetCreateTime() uint64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Fruit) GetUpdateTime() uint64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

var File_base_proto protoreflect.FileDescriptor

var file_base_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x61,
	0x70, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x22, 0xf7, 0x02, 0x0a, 0x05, 0x46, 0x72, 0x75, 0x69, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x04,
	0x75, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x63, 0x61, 0x70,
	0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x74,
	0x6f, 0x6e, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x61, 0x64, 0x75, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x61, 0x64, 0x75, 0x63, 0x61, 0x74, 0x65, 0x44, 0x61, 0x79,
	0x73, 0x12, 0x2b, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x2a, 0x34, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x49, 0x54,
	0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4b, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x49,
	0x54, 0x5f, 0x4c, 0x42, 0x10, 0x02, 0x2a, 0x6d, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x55, 0x4e, 0x44,
	0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x55, 0x4e,
	0x54, 0x52, 0x59, 0x5f, 0x4d, 0x45, 0x58, 0x49, 0x43, 0x4f, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e,
	0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x42, 0x52, 0x41, 0x5a, 0x49, 0x4c, 0x10, 0x02,
	0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x43, 0x41, 0x4e, 0x41,
	0x44, 0x41, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x5f,
	0x55, 0x53, 0x41, 0x10, 0x04, 0x2a, 0x6a, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x55, 0x4e,
	0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x55, 0x52,
	0x52, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x4d, 0x58, 0x4e, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x43,
	0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x42, 0x52, 0x4c, 0x10, 0x02, 0x12, 0x10, 0x0a,
	0x0c, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x43, 0x41, 0x44, 0x10, 0x03, 0x12,
	0x10, 0x0a, 0x0c, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x55, 0x53, 0x44, 0x10,
	0x04, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x61, 0x72, 0x63, 0x6f, 0x73, 0x2d, 0x77, 0x7a, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x74, 0x6f,
	0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_proto_rawDescOnce sync.Once
	file_base_proto_rawDescData = file_base_proto_rawDesc
)

func file_base_proto_rawDescGZIP() []byte {
	file_base_proto_rawDescOnce.Do(func() {
		file_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_proto_rawDescData)
	})
	return file_base_proto_rawDescData
}

var file_base_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_base_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_base_proto_goTypes = []interface{}{
	(Unit)(0),     // 0: capstone.Unit
	(Country)(0),  // 1: capstone.Country
	(Currency)(0), // 2: capstone.Currency
	(*Fruit)(nil), // 3: capstone.Fruit
}
var file_base_proto_depIdxs = []int32{
	0, // 0: capstone.Fruit.unit:type_name -> capstone.Unit
	2, // 1: capstone.Fruit.currency:type_name -> capstone.Currency
	1, // 2: capstone.Fruit.country:type_name -> capstone.Country
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_base_proto_init() }
func file_base_proto_init() {
	if File_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fruit); i {
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
			RawDescriptor: file_base_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_proto_goTypes,
		DependencyIndexes: file_base_proto_depIdxs,
		EnumInfos:         file_base_proto_enumTypes,
		MessageInfos:      file_base_proto_msgTypes,
	}.Build()
	File_base_proto = out.File
	file_base_proto_rawDesc = nil
	file_base_proto_goTypes = nil
	file_base_proto_depIdxs = nil
}

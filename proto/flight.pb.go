package proto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type QueryFlightDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlightId int32 `protobuf:"varint,1,opt,name=flightId,proto3" json:"flightId,omitempty"`
}

func (x *QueryFlightDetail) Reset() {
	*x = QueryFlightDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryFlightDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryFlightDetail) ProtoMessage() {}

func (x *QueryFlightDetail) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryFlightDetail.ProtoReflect.Descriptor instead.
func (*QueryFlightDetail) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{0}
}

func (x *QueryFlightDetail) GetFlightId() int32 {
	if x != nil {
		return x.FlightId
	}
	return 0
}

type FlightDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code          string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	TotalSlot     int32  `protobuf:"varint,3,opt,name=totalSlot,proto3" json:"totalSlot,omitempty"`
	DepartureTime string `protobuf:"bytes,4,opt,name=departureTime,proto3" json:"departureTime,omitempty"`
	ArrivalTime   string `protobuf:"bytes,5,opt,name=arrivalTime,proto3" json:"arrivalTime,omitempty"`
}

func (x *FlightDetailResponse) Reset() {
	*x = FlightDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlightDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlightDetailResponse) ProtoMessage() {}

func (x *FlightDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlightDetailResponse.ProtoReflect.Descriptor instead.
func (*FlightDetailResponse) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{1}
}

func (x *FlightDetailResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FlightDetailResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *FlightDetailResponse) GetTotalSlot() int32 {
	if x != nil {
		return x.TotalSlot
	}
	return 0
}

func (x *FlightDetailResponse) GetDepartureTime() string {
	if x != nil {
		return x.DepartureTime
	}
	return ""
}

func (x *FlightDetailResponse) GetArrivalTime() string {
	if x != nil {
		return x.ArrivalTime
	}
	return ""
}

var File_flight_proto protoreflect.FileDescriptor

var file_flight_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x22, 0x2f, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x49, 0x64, 0x22, 0xa0, 0x01, 0x0a, 0x14, 0x46, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x6c, 0x6f,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x6c,
	0x6f, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x75, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x72, 0x72, 0x69,
	0x76, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x58, 0x0a, 0x0d, 0x46, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x19, 0x2e, 0x66, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x1c, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x79, 0x70, 0x68, 0x61, 0x6d, 0x6a, 0x2f, 0x66, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_flight_proto_rawDescOnce sync.Once
	file_flight_proto_rawDescData = file_flight_proto_rawDesc
)

func file_flight_proto_rawDescGZIP() []byte {
	file_flight_proto_rawDescOnce.Do(func() {
		file_flight_proto_rawDescData = protoimpl.X.CompressGZIP(file_flight_proto_rawDescData)
	})
	return file_flight_proto_rawDescData
}

var file_flight_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_flight_proto_goTypes = []interface{}{
	(*QueryFlightDetail)(nil),    // 0: flight.QueryFlightDetail
	(*FlightDetailResponse)(nil), // 1: flight.FlightDetailResponse
}
var file_flight_proto_depIdxs = []int32{
	0, // 0: flight.FlightService.SearchFlight:input_type -> flight.QueryFlightDetail
	1, // 1: flight.FlightService.SearchFlight:output_type -> flight.FlightDetailResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_flight_proto_init() }
func file_flight_proto_init() {
	if File_flight_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flight_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryFlightDetail); i {
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
		file_flight_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlightDetailResponse); i {
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
			RawDescriptor: file_flight_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flight_proto_goTypes,
		DependencyIndexes: file_flight_proto_depIdxs,
		MessageInfos:      file_flight_proto_msgTypes,
	}.Build()
	File_flight_proto = out.File
	file_flight_proto_rawDesc = nil
	file_flight_proto_goTypes = nil
	file_flight_proto_depIdxs = nil
}

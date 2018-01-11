// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tsdb_plugin.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	tsdb_plugin.proto

It has these top-level messages:
	TsdbQuery
	Query
	TimeRange
	Response
	QueryResult
	DatasourceInfo
	TimeSeries
	Point
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type TsdbQuery struct {
	TimeRange  *TimeRange      `protobuf:"bytes,1,opt,name=timeRange" json:"timeRange,omitempty"`
	Datasource *DatasourceInfo `protobuf:"bytes,2,opt,name=datasource" json:"datasource,omitempty"`
	Queries    []*Query        `protobuf:"bytes,3,rep,name=queries" json:"queries,omitempty"`
}

func (m *TsdbQuery) Reset()                    { *m = TsdbQuery{} }
func (m *TsdbQuery) String() string            { return proto1.CompactTextString(m) }
func (*TsdbQuery) ProtoMessage()               {}
func (*TsdbQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TsdbQuery) GetTimeRange() *TimeRange {
	if m != nil {
		return m.TimeRange
	}
	return nil
}

func (m *TsdbQuery) GetDatasource() *DatasourceInfo {
	if m != nil {
		return m.Datasource
	}
	return nil
}

func (m *TsdbQuery) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

type Query struct {
	RefId         string `protobuf:"bytes,1,opt,name=refId" json:"refId,omitempty"`
	MaxDataPoints int64  `protobuf:"varint,2,opt,name=maxDataPoints" json:"maxDataPoints,omitempty"`
	IntervalMs    int64  `protobuf:"varint,3,opt,name=intervalMs" json:"intervalMs,omitempty"`
	ModelJson     string `protobuf:"bytes,4,opt,name=modelJson" json:"modelJson,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto1.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Query) GetRefId() string {
	if m != nil {
		return m.RefId
	}
	return ""
}

func (m *Query) GetMaxDataPoints() int64 {
	if m != nil {
		return m.MaxDataPoints
	}
	return 0
}

func (m *Query) GetIntervalMs() int64 {
	if m != nil {
		return m.IntervalMs
	}
	return 0
}

func (m *Query) GetModelJson() string {
	if m != nil {
		return m.ModelJson
	}
	return ""
}

type TimeRange struct {
	FromRaw     string `protobuf:"bytes,1,opt,name=fromRaw" json:"fromRaw,omitempty"`
	ToRaw       string `protobuf:"bytes,2,opt,name=toRaw" json:"toRaw,omitempty"`
	FromEpochMs int64  `protobuf:"varint,3,opt,name=fromEpochMs" json:"fromEpochMs,omitempty"`
	ToEpochMs   int64  `protobuf:"varint,4,opt,name=toEpochMs" json:"toEpochMs,omitempty"`
}

func (m *TimeRange) Reset()                    { *m = TimeRange{} }
func (m *TimeRange) String() string            { return proto1.CompactTextString(m) }
func (*TimeRange) ProtoMessage()               {}
func (*TimeRange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TimeRange) GetFromRaw() string {
	if m != nil {
		return m.FromRaw
	}
	return ""
}

func (m *TimeRange) GetToRaw() string {
	if m != nil {
		return m.ToRaw
	}
	return ""
}

func (m *TimeRange) GetFromEpochMs() int64 {
	if m != nil {
		return m.FromEpochMs
	}
	return 0
}

func (m *TimeRange) GetToEpochMs() int64 {
	if m != nil {
		return m.ToEpochMs
	}
	return 0
}

type Response struct {
	Results []*QueryResult `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto1.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetResults() []*QueryResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type QueryResult struct {
	Error    string        `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	RefId    string        `protobuf:"bytes,2,opt,name=refId" json:"refId,omitempty"`
	MetaJson string        `protobuf:"bytes,3,opt,name=metaJson" json:"metaJson,omitempty"`
	Series   []*TimeSeries `protobuf:"bytes,4,rep,name=series" json:"series,omitempty"`
}

func (m *QueryResult) Reset()                    { *m = QueryResult{} }
func (m *QueryResult) String() string            { return proto1.CompactTextString(m) }
func (*QueryResult) ProtoMessage()               {}
func (*QueryResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *QueryResult) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *QueryResult) GetRefId() string {
	if m != nil {
		return m.RefId
	}
	return ""
}

func (m *QueryResult) GetMetaJson() string {
	if m != nil {
		return m.MetaJson
	}
	return ""
}

func (m *QueryResult) GetSeries() []*TimeSeries {
	if m != nil {
		return m.Series
	}
	return nil
}

type DatasourceInfo struct {
	Id             int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	OrgId          int64  `protobuf:"varint,2,opt,name=orgId" json:"orgId,omitempty"`
	Name           string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Type           string `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
	Url            string `protobuf:"bytes,5,opt,name=url" json:"url,omitempty"`
	JsonData       string `protobuf:"bytes,6,opt,name=jsonData" json:"jsonData,omitempty"`
	SecureJsonData string `protobuf:"bytes,7,opt,name=secureJsonData" json:"secureJsonData,omitempty"`
}

func (m *DatasourceInfo) Reset()                    { *m = DatasourceInfo{} }
func (m *DatasourceInfo) String() string            { return proto1.CompactTextString(m) }
func (*DatasourceInfo) ProtoMessage()               {}
func (*DatasourceInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DatasourceInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DatasourceInfo) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *DatasourceInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DatasourceInfo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DatasourceInfo) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *DatasourceInfo) GetJsonData() string {
	if m != nil {
		return m.JsonData
	}
	return ""
}

func (m *DatasourceInfo) GetSecureJsonData() string {
	if m != nil {
		return m.SecureJsonData
	}
	return ""
}

type TimeSeries struct {
	Name   string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Tags   map[string]string `protobuf:"bytes,2,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Points []*Point          `protobuf:"bytes,3,rep,name=points" json:"points,omitempty"`
}

func (m *TimeSeries) Reset()                    { *m = TimeSeries{} }
func (m *TimeSeries) String() string            { return proto1.CompactTextString(m) }
func (*TimeSeries) ProtoMessage()               {}
func (*TimeSeries) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TimeSeries) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TimeSeries) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *TimeSeries) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

type Point struct {
	Timestamp int64   `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Value     float64 `protobuf:"fixed64,2,opt,name=value" json:"value,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto1.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Point) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Point) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto1.RegisterType((*TsdbQuery)(nil), "plugins.TsdbQuery")
	proto1.RegisterType((*Query)(nil), "plugins.Query")
	proto1.RegisterType((*TimeRange)(nil), "plugins.TimeRange")
	proto1.RegisterType((*Response)(nil), "plugins.Response")
	proto1.RegisterType((*QueryResult)(nil), "plugins.QueryResult")
	proto1.RegisterType((*DatasourceInfo)(nil), "plugins.DatasourceInfo")
	proto1.RegisterType((*TimeSeries)(nil), "plugins.TimeSeries")
	proto1.RegisterType((*Point)(nil), "plugins.Point")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TsdbPlugin service

type TsdbPluginClient interface {
	Query(ctx context.Context, in *TsdbQuery, opts ...grpc.CallOption) (*Response, error)
}

type tsdbPluginClient struct {
	cc *grpc.ClientConn
}

func NewTsdbPluginClient(cc *grpc.ClientConn) TsdbPluginClient {
	return &tsdbPluginClient{cc}
}

func (c *tsdbPluginClient) Query(ctx context.Context, in *TsdbQuery, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/plugins.TsdbPlugin/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TsdbPlugin service

type TsdbPluginServer interface {
	Query(context.Context, *TsdbQuery) (*Response, error)
}

func RegisterTsdbPluginServer(s *grpc.Server, srv TsdbPluginServer) {
	s.RegisterService(&_TsdbPlugin_serviceDesc, srv)
}

func _TsdbPlugin_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TsdbQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TsdbPluginServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugins.TsdbPlugin/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TsdbPluginServer).Query(ctx, req.(*TsdbQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _TsdbPlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "plugins.TsdbPlugin",
	HandlerType: (*TsdbPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _TsdbPlugin_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tsdb_plugin.proto",
}

func init() { proto1.RegisterFile("tsdb_plugin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0x51, 0x8b, 0xd3, 0x4c,
	0x14, 0x25, 0x49, 0xdb, 0x6c, 0x6e, 0xf9, 0xca, 0xb7, 0xe3, 0x82, 0xa1, 0xa8, 0x94, 0x20, 0x4b,
	0x41, 0x08, 0x5a, 0x1f, 0x56, 0x56, 0x9f, 0xc4, 0x7d, 0xd8, 0x05, 0x61, 0x1d, 0xfb, 0xe4, 0x8b,
	0x4c, 0x9b, 0x69, 0x8c, 0x26, 0x99, 0x38, 0x33, 0x59, 0x2d, 0x3e, 0xe9, 0x2f, 0xf1, 0x47, 0xf8,
	0x03, 0x65, 0x6e, 0x32, 0x49, 0x5a, 0x7c, 0xea, 0xdc, 0x73, 0xce, 0xcc, 0x3d, 0x39, 0x73, 0x3b,
	0x70, 0xaa, 0x55, 0xb2, 0xf9, 0x58, 0xe5, 0x75, 0x9a, 0x95, 0x71, 0x25, 0x85, 0x16, 0xc4, 0x6f,
	0x2a, 0x15, 0xfd, 0x76, 0x20, 0x58, 0xab, 0x64, 0xf3, 0xae, 0xe6, 0x72, 0x4f, 0x9e, 0x42, 0xa0,
	0xb3, 0x82, 0x53, 0x56, 0xa6, 0x3c, 0x74, 0x16, 0xce, 0x72, 0xba, 0x22, 0x71, 0x2b, 0x8d, 0xd7,
	0x96, 0xa1, 0xbd, 0x88, 0x5c, 0x00, 0x24, 0x4c, 0x33, 0x25, 0x6a, 0xb9, 0xe5, 0xa1, 0x8b, 0x5b,
	0xee, 0x77, 0x5b, 0xde, 0x74, 0xd4, 0x75, 0xb9, 0x13, 0x74, 0x20, 0x25, 0x4b, 0xf0, 0xbf, 0xd6,
	0x5c, 0x66, 0x5c, 0x85, 0xde, 0xc2, 0x5b, 0x4e, 0x57, 0xb3, 0x6e, 0x17, 0x7a, 0xa1, 0x96, 0x8e,
	0x7e, 0x3a, 0x30, 0x6e, 0xec, 0x9d, 0xc1, 0x58, 0xf2, 0xdd, 0x75, 0x82, 0xd6, 0x02, 0xda, 0x14,
	0xe4, 0x31, 0xfc, 0x57, 0xb0, 0xef, 0xa6, 0xd5, 0xad, 0xc8, 0x4a, 0xad, 0xd0, 0x85, 0x47, 0x0f,
	0x41, 0xf2, 0x08, 0x20, 0x2b, 0x35, 0x97, 0x77, 0x2c, 0x7f, 0x6b, 0x5a, 0x1a, 0xc9, 0x00, 0x21,
	0x0f, 0x20, 0x28, 0x44, 0xc2, 0xf3, 0x1b, 0x25, 0xca, 0x70, 0x84, 0xe7, 0xf7, 0x40, 0xf4, 0x03,
	0x82, 0xee, 0xf3, 0x49, 0x08, 0xfe, 0x4e, 0x8a, 0x82, 0xb2, 0x6f, 0xad, 0x11, 0x5b, 0x1a, 0x83,
	0x5a, 0x18, 0xdc, 0x6d, 0x0c, 0x62, 0x41, 0x16, 0x30, 0x35, 0x82, 0xab, 0x4a, 0x6c, 0x3f, 0x75,
	0xbd, 0x87, 0x90, 0x69, 0xae, 0x85, 0xe5, 0x47, 0xc8, 0xf7, 0x40, 0x74, 0x09, 0x27, 0x94, 0xab,
	0x4a, 0x94, 0x8a, 0x93, 0x18, 0x7c, 0xc9, 0x55, 0x9d, 0x6b, 0x15, 0x3a, 0x18, 0xdb, 0xd9, 0x51,
	0x6c, 0x48, 0x52, 0x2b, 0x8a, 0x7e, 0x39, 0x30, 0x1d, 0x10, 0xc6, 0x21, 0x97, 0x52, 0x48, 0x1b,
	0x21, 0x16, 0x7d, 0xb0, 0xee, 0x30, 0xd8, 0x39, 0x9c, 0x14, 0x5c, 0x33, 0x4c, 0xc4, 0x43, 0xa2,
	0xab, 0xc9, 0x13, 0x98, 0xa8, 0xe6, 0xf6, 0x46, 0x68, 0xe3, 0xde, 0xc1, 0x98, 0xbc, 0x47, 0x8a,
	0xb6, 0x92, 0xe8, 0x8f, 0x03, 0xb3, 0xc3, 0x51, 0x20, 0x33, 0x70, 0xb3, 0xe6, 0x1e, 0x3d, 0xea,
	0x66, 0x89, 0x71, 0x20, 0x64, 0xda, 0x3a, 0xf0, 0x68, 0x53, 0x10, 0x02, 0xa3, 0x92, 0x15, 0xbc,
	0xed, 0x8e, 0x6b, 0x83, 0xe9, 0x7d, 0xc5, 0xdb, 0x3b, 0xc2, 0x35, 0xf9, 0x1f, 0xbc, 0x5a, 0xe6,
	0xe1, 0x18, 0x21, 0xb3, 0x34, 0xde, 0x3f, 0x2b, 0x51, 0x9a, 0xae, 0xe1, 0xa4, 0xf1, 0x6e, 0x6b,
	0x72, 0x0e, 0x33, 0xc5, 0xb7, 0xb5, 0xe4, 0x37, 0x56, 0xe1, 0xa3, 0xe2, 0x08, 0x35, 0xb6, 0xa1,
	0xff, 0x9a, 0xce, 0x8c, 0x33, 0x30, 0xf3, 0x0c, 0x46, 0x9a, 0xa5, 0x66, 0xe4, 0x4c, 0x08, 0x0f,
	0xff, 0x11, 0x42, 0xbc, 0x66, 0xa9, 0xba, 0x2a, 0xb5, 0xdc, 0x53, 0x94, 0x92, 0x73, 0x98, 0x54,
	0xcd, 0x9c, 0x1e, 0xcf, 0x3d, 0x4e, 0x2a, 0x6d, 0xd9, 0xf9, 0x05, 0x04, 0xdd, 0x56, 0xf3, 0x81,
	0x5f, 0xf8, 0xbe, 0x6d, 0x6d, 0x96, 0x26, 0xb0, 0x3b, 0x96, 0xd7, 0xdc, 0x5e, 0x19, 0x16, 0x97,
	0xee, 0x0b, 0x27, 0x7a, 0x09, 0x63, 0x3c, 0x09, 0xa7, 0x2a, 0x2b, 0xb8, 0xd2, 0xac, 0xa8, 0xda,
	0xa8, 0x7b, 0xe0, 0xf0, 0x00, 0xa7, 0x3d, 0x60, 0xf5, 0x0a, 0xc0, 0x3c, 0x07, 0xb7, 0x68, 0x89,
	0xc4, 0xf6, 0x9f, 0x37, 0x78, 0x05, 0xec, 0x63, 0x31, 0x3f, 0xed, 0x30, 0x3b, 0x9d, 0xaf, 0xfd,
	0x0f, 0x63, 0x7c, 0x5f, 0x36, 0x13, 0xfc, 0x79, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x69,
	0x3b, 0x74, 0x7b, 0x04, 0x00, 0x00,
}

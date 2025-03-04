// Code generated by protoc-gen-go. DO NOT EDIT.
// source: as/external/api/profiles.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RatePolicy int32

const (
	// Drop
	RatePolicy_DROP RatePolicy = 0
	// Mark
	RatePolicy_MARK RatePolicy = 1
)

var RatePolicy_name = map[int32]string{
	0: "DROP",
	1: "MARK",
}

var RatePolicy_value = map[string]int32{
	"DROP": 0,
	"MARK": 1,
}

func (x RatePolicy) String() string {
	return proto.EnumName(RatePolicy_name, int32(x))
}

func (RatePolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a1c507fa3dbc9903, []int{0}
}

type ServiceProfile struct {
	// Service-profile ID (UUID string).
	// This will be automatically set on create.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Service-profile name.
	Name string `protobuf:"bytes,21,opt,name=name,proto3" json:"name,omitempty"`
	// Organization ID to which the service-profile is assigned.
	OrganizationId int64 `protobuf:"varint,22,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	// Network-server ID on which the service-profile is provisioned.
	NetworkServerId int64 `protobuf:"varint,23,opt,name=network_server_id,json=networkServerID,proto3" json:"network_server_id,omitempty"`
	// Token bucket filling rate, including ACKs (packet/h).
	UlRate uint32 `protobuf:"varint,2,opt,name=ul_rate,json=ulRate,proto3" json:"ul_rate,omitempty"`
	// Token bucket burst size.
	UlBucketSize uint32 `protobuf:"varint,3,opt,name=ul_bucket_size,json=ulBucketSize,proto3" json:"ul_bucket_size,omitempty"`
	// Drop or mark when exceeding ULRate.
	UlRatePolicy RatePolicy `protobuf:"varint,4,opt,name=ul_rate_policy,json=ulRatePolicy,proto3,enum=api.RatePolicy" json:"ul_rate_policy,omitempty"`
	// Token bucket filling rate, including ACKs (packet/h).
	DlRate uint32 `protobuf:"varint,5,opt,name=dl_rate,json=dlRate,proto3" json:"dl_rate,omitempty"`
	// Token bucket burst size.
	DlBucketSize uint32 `protobuf:"varint,6,opt,name=dl_bucket_size,json=dlBucketSize,proto3" json:"dl_bucket_size,omitempty"`
	// Drop or mark when exceeding DLRate.
	DlRatePolicy RatePolicy `protobuf:"varint,7,opt,name=dl_rate_policy,json=dlRatePolicy,proto3,enum=api.RatePolicy" json:"dl_rate_policy,omitempty"`
	// GW metadata (RSSI, SNR, GW geoloc., etc.) are added to the packet sent to AS.
	AddGwMetadata bool `protobuf:"varint,8,opt,name=add_gw_metadata,json=addGWMetaData,proto3" json:"add_gw_metadata,omitempty"`
	// Frequency to initiate an End-Device status request (request/day).
	DevStatusReqFreq uint32 `protobuf:"varint,9,opt,name=dev_status_req_freq,json=devStatusReqFreq,proto3" json:"dev_status_req_freq,omitempty"`
	// Report End-Device battery level to AS.
	ReportDevStatusBattery bool `protobuf:"varint,10,opt,name=report_dev_status_battery,json=reportDevStatusBattery,proto3" json:"report_dev_status_battery,omitempty"`
	// Report End-Device margin to AS.
	ReportDevStatusMargin bool `protobuf:"varint,11,opt,name=report_dev_status_margin,json=reportDevStatusMargin,proto3" json:"report_dev_status_margin,omitempty"`
	// Minimum allowed data rate. Used for ADR.
	DrMin uint32 `protobuf:"varint,12,opt,name=dr_min,json=drMin,proto3" json:"dr_min,omitempty"`
	// Maximum allowed data rate. Used for ADR.
	DrMax uint32 `protobuf:"varint,13,opt,name=dr_max,json=drMax,proto3" json:"dr_max,omitempty"`
	// Channel mask. sNS does not have to obey (i.e., informative).
	ChannelMask []byte `protobuf:"bytes,14,opt,name=channel_mask,json=channelMask,proto3" json:"channel_mask,omitempty"`
	// Passive Roaming allowed.
	PrAllowed bool `protobuf:"varint,15,opt,name=pr_allowed,json=prAllowed,proto3" json:"pr_allowed,omitempty"`
	// Handover Roaming allowed.
	HrAllowed bool `protobuf:"varint,16,opt,name=hr_allowed,json=hrAllowed,proto3" json:"hr_allowed,omitempty"`
	// Roaming Activation allowed.
	RaAllowed bool `protobuf:"varint,17,opt,name=ra_allowed,json=raAllowed,proto3" json:"ra_allowed,omitempty"`
	// Enable network geolocation service.
	NwkGeoLoc bool `protobuf:"varint,18,opt,name=nwk_geo_loc,json=nwkGeoLoc,proto3" json:"nwk_geo_loc,omitempty"`
	// Target Packet Error Rate.
	TargetPer uint32 `protobuf:"varint,19,opt,name=target_per,json=targetPER,proto3" json:"target_per,omitempty"`
	// Minimum number of receiving GWs (informative).
	MinGwDiversity       uint32   `protobuf:"varint,20,opt,name=min_gw_diversity,json=minGWDiversity,proto3" json:"min_gw_diversity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceProfile) Reset()         { *m = ServiceProfile{} }
func (m *ServiceProfile) String() string { return proto.CompactTextString(m) }
func (*ServiceProfile) ProtoMessage()    {}
func (*ServiceProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c507fa3dbc9903, []int{0}
}

func (m *ServiceProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceProfile.Unmarshal(m, b)
}
func (m *ServiceProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceProfile.Marshal(b, m, deterministic)
}
func (m *ServiceProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceProfile.Merge(m, src)
}
func (m *ServiceProfile) XXX_Size() int {
	return xxx_messageInfo_ServiceProfile.Size(m)
}
func (m *ServiceProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceProfile.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceProfile proto.InternalMessageInfo

func (m *ServiceProfile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServiceProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceProfile) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

func (m *ServiceProfile) GetNetworkServerId() int64 {
	if m != nil {
		return m.NetworkServerId
	}
	return 0
}

func (m *ServiceProfile) GetUlRate() uint32 {
	if m != nil {
		return m.UlRate
	}
	return 0
}

func (m *ServiceProfile) GetUlBucketSize() uint32 {
	if m != nil {
		return m.UlBucketSize
	}
	return 0
}

func (m *ServiceProfile) GetUlRatePolicy() RatePolicy {
	if m != nil {
		return m.UlRatePolicy
	}
	return RatePolicy_DROP
}

func (m *ServiceProfile) GetDlRate() uint32 {
	if m != nil {
		return m.DlRate
	}
	return 0
}

func (m *ServiceProfile) GetDlBucketSize() uint32 {
	if m != nil {
		return m.DlBucketSize
	}
	return 0
}

func (m *ServiceProfile) GetDlRatePolicy() RatePolicy {
	if m != nil {
		return m.DlRatePolicy
	}
	return RatePolicy_DROP
}

func (m *ServiceProfile) GetAddGwMetadata() bool {
	if m != nil {
		return m.AddGwMetadata
	}
	return false
}

func (m *ServiceProfile) GetDevStatusReqFreq() uint32 {
	if m != nil {
		return m.DevStatusReqFreq
	}
	return 0
}

func (m *ServiceProfile) GetReportDevStatusBattery() bool {
	if m != nil {
		return m.ReportDevStatusBattery
	}
	return false
}

func (m *ServiceProfile) GetReportDevStatusMargin() bool {
	if m != nil {
		return m.ReportDevStatusMargin
	}
	return false
}

func (m *ServiceProfile) GetDrMin() uint32 {
	if m != nil {
		return m.DrMin
	}
	return 0
}

func (m *ServiceProfile) GetDrMax() uint32 {
	if m != nil {
		return m.DrMax
	}
	return 0
}

func (m *ServiceProfile) GetChannelMask() []byte {
	if m != nil {
		return m.ChannelMask
	}
	return nil
}

func (m *ServiceProfile) GetPrAllowed() bool {
	if m != nil {
		return m.PrAllowed
	}
	return false
}

func (m *ServiceProfile) GetHrAllowed() bool {
	if m != nil {
		return m.HrAllowed
	}
	return false
}

func (m *ServiceProfile) GetRaAllowed() bool {
	if m != nil {
		return m.RaAllowed
	}
	return false
}

func (m *ServiceProfile) GetNwkGeoLoc() bool {
	if m != nil {
		return m.NwkGeoLoc
	}
	return false
}

func (m *ServiceProfile) GetTargetPer() uint32 {
	if m != nil {
		return m.TargetPer
	}
	return 0
}

func (m *ServiceProfile) GetMinGwDiversity() uint32 {
	if m != nil {
		return m.MinGwDiversity
	}
	return 0
}

type DeviceProfile struct {
	// Device-profile ID (UUID string).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Device-profile name.
	Name string `protobuf:"bytes,21,opt,name=name,proto3" json:"name,omitempty"`
	// Organization ID to which the service-profile is assigned.
	OrganizationId int64 `protobuf:"varint,22,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	// Network-server ID on which the service-profile is provisioned.
	NetworkServerId int64 `protobuf:"varint,23,opt,name=network_server_id,json=networkServerID,proto3" json:"network_server_id,omitempty"`
	// End-Device supports Class B.
	SupportsClassB bool `protobuf:"varint,2,opt,name=supports_class_b,json=supportsClassB,proto3" json:"supports_class_b,omitempty"`
	// Maximum delay for the End-Device to answer a MAC request or a confirmed DL frame (mandatory if class B mode supported).
	ClassBTimeout uint32 `protobuf:"varint,3,opt,name=class_b_timeout,json=classBTimeout,proto3" json:"class_b_timeout,omitempty"`
	// Mandatory if class B mode supported.
	PingSlotPeriod uint32 `protobuf:"varint,4,opt,name=ping_slot_period,json=pingSlotPeriod,proto3" json:"ping_slot_period,omitempty"`
	// Mandatory if class B mode supported.
	PingSlotDr uint32 `protobuf:"varint,5,opt,name=ping_slot_dr,json=pingSlotDR,proto3" json:"ping_slot_dr,omitempty"`
	// Mandatory if class B mode supported.
	PingSlotFreq uint32 `protobuf:"varint,6,opt,name=ping_slot_freq,json=pingSlotFreq,proto3" json:"ping_slot_freq,omitempty"`
	// End-Device supports Class C.
	SupportsClassC bool `protobuf:"varint,7,opt,name=supports_class_c,json=supportsClassC,proto3" json:"supports_class_c,omitempty"`
	// Maximum delay for the End-Device to answer a MAC request or a confirmed DL frame (mandatory if class C mode supported).
	ClassCTimeout uint32 `protobuf:"varint,8,opt,name=class_c_timeout,json=classCTimeout,proto3" json:"class_c_timeout,omitempty"`
	// Version of the LoRaWAN supported by the End-Device.
	MacVersion string `protobuf:"bytes,9,opt,name=mac_version,json=macVersion,proto3" json:"mac_version,omitempty"`
	// Revision of the Regional Parameters document supported by the End-Device.
	RegParamsRevision string `protobuf:"bytes,10,opt,name=reg_params_revision,json=regParamsRevision,proto3" json:"reg_params_revision,omitempty"`
	// Class A RX1 delay (mandatory for ABP).
	RxDelay_1 uint32 `protobuf:"varint,11,opt,name=rx_delay_1,json=rxDelay1,proto3" json:"rx_delay_1,omitempty"`
	// RX1 data rate offset (mandatory for ABP).
	RxDrOffset_1 uint32 `protobuf:"varint,12,opt,name=rx_dr_offset_1,json=rxDROffset1,proto3" json:"rx_dr_offset_1,omitempty"`
	// RX2 data rate (mandatory for ABP).
	RxDatarate_2 uint32 `protobuf:"varint,13,opt,name=rx_datarate_2,json=rxDataRate2,proto3" json:"rx_datarate_2,omitempty"`
	// RX2 channel frequency (mandatory for ABP).
	RxFreq_2 uint32 `protobuf:"varint,14,opt,name=rx_freq_2,json=rxFreq2,proto3" json:"rx_freq_2,omitempty"`
	// List of factory-preset frequencies (mandatory for ABP).
	FactoryPresetFreqs []uint32 `protobuf:"varint,15,rep,packed,name=factory_preset_freqs,json=factoryPresetFreqs,proto3" json:"factory_preset_freqs,omitempty"`
	// Maximum EIRP supported by the End-Device.
	MaxEirp uint32 `protobuf:"varint,16,opt,name=max_eirp,json=maxEIRP,proto3" json:"max_eirp,omitempty"`
	// Maximum duty cycle supported by the End-Device.
	MaxDutyCycle uint32 `protobuf:"varint,17,opt,name=max_duty_cycle,json=maxDutyCycle,proto3" json:"max_duty_cycle,omitempty"`
	// End-Device supports Join (OTAA) or not (ABP).
	SupportsJoin bool `protobuf:"varint,18,opt,name=supports_join,json=supportsJoin,proto3" json:"supports_join,omitempty"`
	// RF region name.
	RfRegion string `protobuf:"bytes,19,opt,name=rf_region,json=rfRegion,proto3" json:"rf_region,omitempty"`
	// End-Device uses 32bit FCnt (mandatory for LoRaWAN 1.0 End-Device).
	Supports_32BitFCnt bool `protobuf:"varint,20,opt,name=supports_32bit_f_cnt,json=supports32BitFCnt,proto3" json:"supports_32bit_f_cnt,omitempty"`
	// Payload codec.
	// Leave blank to disable the codec feature.
	PayloadCodec string `protobuf:"bytes,24,opt,name=payload_codec,json=payloadCodec,proto3" json:"payload_codec,omitempty"`
	// Payload encoder script.
	// Depending the codec, it is possible to provide a script which implements
	// the encoder function.
	PayloadEncoderScript string `protobuf:"bytes,25,opt,name=payload_encoder_script,json=payloadEncoderScript,proto3" json:"payload_encoder_script,omitempty"`
	// Payload decoder script.
	// Depending the codec, it is possible to provide a script which implements
	// the decoder function.
	PayloadDecoderScript string `protobuf:"bytes,26,opt,name=payload_decoder_script,json=payloadDecoderScript,proto3" json:"payload_decoder_script,omitempty"`
	// Geolocation buffer TTL (in seconds).
	// When > 0, uplink RX meta-data will be stored in a buffer so that
	// the meta-data of multiple uplinks can be used for geolocation.
	GeolocBufferTtl uint32 `protobuf:"varint,27,opt,name=geoloc_buffer_ttl,json=geolocBufferTTL,proto3" json:"geoloc_buffer_ttl,omitempty"`
	// Geolocation minimum buffer size.
	// When > 0, geolocation will only be performed when the buffer has
	// at least the given size.
	GeolocMinBufferSize  uint32   `protobuf:"varint,28,opt,name=geoloc_min_buffer_size,json=geolocMinBufferSize,proto3" json:"geoloc_min_buffer_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceProfile) Reset()         { *m = DeviceProfile{} }
func (m *DeviceProfile) String() string { return proto.CompactTextString(m) }
func (*DeviceProfile) ProtoMessage()    {}
func (*DeviceProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c507fa3dbc9903, []int{1}
}

func (m *DeviceProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceProfile.Unmarshal(m, b)
}
func (m *DeviceProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceProfile.Marshal(b, m, deterministic)
}
func (m *DeviceProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceProfile.Merge(m, src)
}
func (m *DeviceProfile) XXX_Size() int {
	return xxx_messageInfo_DeviceProfile.Size(m)
}
func (m *DeviceProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceProfile.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceProfile proto.InternalMessageInfo

func (m *DeviceProfile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeviceProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceProfile) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

func (m *DeviceProfile) GetNetworkServerId() int64 {
	if m != nil {
		return m.NetworkServerId
	}
	return 0
}

func (m *DeviceProfile) GetSupportsClassB() bool {
	if m != nil {
		return m.SupportsClassB
	}
	return false
}

func (m *DeviceProfile) GetClassBTimeout() uint32 {
	if m != nil {
		return m.ClassBTimeout
	}
	return 0
}

func (m *DeviceProfile) GetPingSlotPeriod() uint32 {
	if m != nil {
		return m.PingSlotPeriod
	}
	return 0
}

func (m *DeviceProfile) GetPingSlotDr() uint32 {
	if m != nil {
		return m.PingSlotDr
	}
	return 0
}

func (m *DeviceProfile) GetPingSlotFreq() uint32 {
	if m != nil {
		return m.PingSlotFreq
	}
	return 0
}

func (m *DeviceProfile) GetSupportsClassC() bool {
	if m != nil {
		return m.SupportsClassC
	}
	return false
}

func (m *DeviceProfile) GetClassCTimeout() uint32 {
	if m != nil {
		return m.ClassCTimeout
	}
	return 0
}

func (m *DeviceProfile) GetMacVersion() string {
	if m != nil {
		return m.MacVersion
	}
	return ""
}

func (m *DeviceProfile) GetRegParamsRevision() string {
	if m != nil {
		return m.RegParamsRevision
	}
	return ""
}

func (m *DeviceProfile) GetRxDelay_1() uint32 {
	if m != nil {
		return m.RxDelay_1
	}
	return 0
}

func (m *DeviceProfile) GetRxDrOffset_1() uint32 {
	if m != nil {
		return m.RxDrOffset_1
	}
	return 0
}

func (m *DeviceProfile) GetRxDatarate_2() uint32 {
	if m != nil {
		return m.RxDatarate_2
	}
	return 0
}

func (m *DeviceProfile) GetRxFreq_2() uint32 {
	if m != nil {
		return m.RxFreq_2
	}
	return 0
}

func (m *DeviceProfile) GetFactoryPresetFreqs() []uint32 {
	if m != nil {
		return m.FactoryPresetFreqs
	}
	return nil
}

func (m *DeviceProfile) GetMaxEirp() uint32 {
	if m != nil {
		return m.MaxEirp
	}
	return 0
}

func (m *DeviceProfile) GetMaxDutyCycle() uint32 {
	if m != nil {
		return m.MaxDutyCycle
	}
	return 0
}

func (m *DeviceProfile) GetSupportsJoin() bool {
	if m != nil {
		return m.SupportsJoin
	}
	return false
}

func (m *DeviceProfile) GetRfRegion() string {
	if m != nil {
		return m.RfRegion
	}
	return ""
}

func (m *DeviceProfile) GetSupports_32BitFCnt() bool {
	if m != nil {
		return m.Supports_32BitFCnt
	}
	return false
}

func (m *DeviceProfile) GetPayloadCodec() string {
	if m != nil {
		return m.PayloadCodec
	}
	return ""
}

func (m *DeviceProfile) GetPayloadEncoderScript() string {
	if m != nil {
		return m.PayloadEncoderScript
	}
	return ""
}

func (m *DeviceProfile) GetPayloadDecoderScript() string {
	if m != nil {
		return m.PayloadDecoderScript
	}
	return ""
}

func (m *DeviceProfile) GetGeolocBufferTtl() uint32 {
	if m != nil {
		return m.GeolocBufferTtl
	}
	return 0
}

func (m *DeviceProfile) GetGeolocMinBufferSize() uint32 {
	if m != nil {
		return m.GeolocMinBufferSize
	}
	return 0
}

func init() {
	proto.RegisterEnum("api.RatePolicy", RatePolicy_name, RatePolicy_value)
	proto.RegisterType((*ServiceProfile)(nil), "api.ServiceProfile")
	proto.RegisterType((*DeviceProfile)(nil), "api.DeviceProfile")
}

func init() { proto.RegisterFile("as/external/api/profiles.proto", fileDescriptor_a1c507fa3dbc9903) }

var fileDescriptor_a1c507fa3dbc9903 = []byte{
	// 1072 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xdd, 0x6e, 0xdb, 0x36,
	0x1b, 0xfe, 0xdc, 0xb4, 0x89, 0xcd, 0x58, 0x76, 0xc2, 0xa4, 0xa9, 0xfa, 0xf3, 0x75, 0x5e, 0x3b,
	0x6c, 0x46, 0x81, 0xc6, 0x8b, 0xb3, 0x62, 0xd8, 0x61, 0x1d, 0xb5, 0x45, 0xb7, 0x1a, 0x35, 0xe8,
	0x62, 0x05, 0x76, 0x42, 0xbc, 0x16, 0x69, 0x85, 0xb3, 0x24, 0x2a, 0x14, 0xe5, 0xd8, 0xbd, 0xc4,
	0xdd, 0xc0, 0x6e, 0x67, 0xe0, 0x2b, 0xf9, 0x27, 0x29, 0x76, 0xbe, 0x33, 0xe9, 0xf9, 0xe1, 0x23,
	0x92, 0x7a, 0x24, 0x92, 0xa7, 0x90, 0xf7, 0xe4, 0xc2, 0x4a, 0x93, 0x42, 0xdc, 0x83, 0x4c, 0xf5,
	0x32, 0xa3, 0xa7, 0x2a, 0x96, 0xf9, 0x69, 0x66, 0xb4, 0xd5, 0x74, 0x07, 0x32, 0xf5, 0xec, 0xef,
	0x5d, 0xd2, 0x1a, 0x4b, 0x33, 0x57, 0xa1, 0x1c, 0x95, 0x34, 0x6d, 0x91, 0x3b, 0x4a, 0xf8, 0xb5,
	0x4e, 0xad, 0xdb, 0x60, 0x77, 0x94, 0xa0, 0x94, 0xdc, 0x4d, 0x21, 0x91, 0xfe, 0x7d, 0x44, 0xf0,
	0x9a, 0xfe, 0x40, 0xda, 0xda, 0x44, 0x90, 0xaa, 0x2f, 0x60, 0x95, 0x4e, 0xb9, 0x12, 0xfe, 0x49,
	0xa7, 0xd6, 0xdd, 0x61, 0xad, 0x6d, 0xf8, 0x7d, 0x40, 0x5f, 0x90, 0xc3, 0x54, 0xda, 0x6b, 0x6d,
	0x66, 0x3c, 0x97, 0x66, 0x2e, 0x8d, 0x93, 0x3e, 0x40, 0x69, 0xbb, 0x22, 0xc6, 0x88, 0xbf, 0x0f,
	0xe8, 0x03, 0xb2, 0x57, 0xc4, 0xdc, 0x80, 0x95, 0xfe, 0x9d, 0x4e, 0xad, 0xeb, 0xb1, 0xdd, 0x22,
	0x66, 0x60, 0x25, 0xfd, 0x8e, 0xb4, 0x8a, 0x98, 0x4f, 0x8a, 0x70, 0x26, 0x2d, 0xcf, 0xd5, 0x17,
	0xe9, 0xef, 0x20, 0xdf, 0x2c, 0xe2, 0x01, 0x82, 0x63, 0xf5, 0x45, 0xd2, 0x57, 0xa8, 0x72, 0x76,
	0x9e, 0xe9, 0x58, 0x85, 0x4b, 0xff, 0x6e, 0xa7, 0xd6, 0x6d, 0xf5, 0xdb, 0xa7, 0x90, 0xa9, 0x53,
	0x37, 0xd0, 0x08, 0x61, 0x67, 0xdb, 0xdc, 0xb9, 0x54, 0x51, 0xa5, 0xde, 0x2b, 0x53, 0xc5, 0x3a,
	0x55, 0xdc, 0x4c, 0xdd, 0x2d, 0x53, 0xc5, 0xad, 0x54, 0x71, 0x33, 0x75, 0xef, 0x5f, 0x52, 0xc5,
	0x76, 0xea, 0xf7, 0xa4, 0x0d, 0x42, 0xf0, 0xe8, 0x9a, 0x27, 0xd2, 0x82, 0x00, 0x0b, 0x7e, 0xbd,
	0x53, 0xeb, 0xd6, 0x99, 0x07, 0x42, 0xbc, 0xfb, 0x3c, 0x94, 0x16, 0x02, 0xb0, 0x40, 0x5f, 0x92,
	0x23, 0x21, 0xe7, 0x3c, 0xb7, 0x60, 0x8b, 0x9c, 0x1b, 0x79, 0xc5, 0xa7, 0x46, 0x5e, 0xf9, 0x0d,
	0x7c, 0x92, 0x03, 0x21, 0xe7, 0x63, 0x64, 0x98, 0xbc, 0x7a, 0x6b, 0xe4, 0x15, 0xfd, 0x85, 0x3c,
	0x34, 0x32, 0xd3, 0xc6, 0xf2, 0x2d, 0xd7, 0x04, 0xac, 0x95, 0x66, 0xe9, 0x13, 0x0c, 0x38, 0x29,
	0x05, 0xc1, 0xca, 0x3a, 0x28, 0x59, 0xfa, 0x33, 0xf1, 0xbf, 0xb6, 0x26, 0x60, 0x22, 0x95, 0xfa,
	0xfb, 0xe8, 0xbc, 0x7f, 0xcb, 0x39, 0x44, 0x92, 0xde, 0x27, 0xbb, 0xc2, 0xf0, 0x44, 0xa5, 0x7e,
	0x13, 0x9f, 0xea, 0x9e, 0x30, 0xc3, 0x0d, 0x0c, 0x0b, 0xdf, 0x5b, 0xc3, 0xb0, 0xa0, 0xdf, 0x92,
	0x66, 0x78, 0x09, 0x69, 0x2a, 0x63, 0x9e, 0x40, 0x3e, 0xf3, 0x5b, 0x9d, 0x5a, 0xb7, 0xc9, 0xf6,
	0x2b, 0x6c, 0x08, 0xf9, 0x8c, 0xfe, 0x9f, 0x90, 0xcc, 0x70, 0x88, 0x63, 0x7d, 0x2d, 0x85, 0xdf,
	0xc6, 0xec, 0x46, 0x66, 0x5e, 0x97, 0x80, 0xa3, 0x2f, 0x37, 0xf4, 0x41, 0x49, 0x5f, 0x6e, 0xd3,
	0x06, 0xd6, 0xf4, 0x61, 0x49, 0x1b, 0x58, 0xd1, 0x4f, 0xc9, 0x7e, 0x7a, 0x3d, 0xe3, 0x91, 0xd4,
	0x3c, 0xd6, 0xa1, 0x4f, 0x4b, 0x3e, 0xbd, 0x9e, 0xbd, 0x93, 0xfa, 0x83, 0x0e, 0x9d, 0xdd, 0x82,
	0x89, 0xa4, 0xe5, 0x99, 0x34, 0xfe, 0x11, 0x3e, 0x7a, 0xa3, 0x44, 0x46, 0x6f, 0x18, 0xed, 0x92,
	0x83, 0x44, 0xa5, 0x6e, 0xdf, 0x84, 0x9a, 0x4b, 0x93, 0x2b, 0xbb, 0xf4, 0x8f, 0x51, 0xd4, 0x4a,
	0x54, 0xfa, 0xee, 0x73, 0xb0, 0x42, 0x9f, 0xfd, 0x55, 0x27, 0x5e, 0x20, 0xff, 0x13, 0xc5, 0xea,
	0x92, 0x83, 0xbc, 0xc8, 0xdc, 0xde, 0xe5, 0x3c, 0x8c, 0x21, 0xcf, 0xf9, 0x04, 0x1b, 0x56, 0x67,
	0xad, 0x15, 0x7e, 0xe1, 0xe0, 0x81, 0x7b, 0x2d, 0x2b, 0x01, 0xb7, 0x2a, 0x91, 0xba, 0xb0, 0x55,
	0xd5, 0x3c, 0x84, 0x07, 0x9f, 0x4a, 0xd0, 0x8d, 0x98, 0xa9, 0x34, 0xe2, 0x79, 0xac, 0x71, 0xa1,
	0x94, 0x16, 0xd8, 0x36, 0x8f, 0xb5, 0x1c, 0x3e, 0x8e, 0xb5, 0x1d, 0x21, 0x4a, 0x3b, 0xa4, 0xb9,
	0x51, 0x0a, 0x53, 0x75, 0x8c, 0xac, 0x54, 0x01, 0x73, 0x3d, 0xdb, 0x28, 0xf0, 0xed, 0xae, 0x7a,
	0xb6, 0xd2, 0xe0, 0x9b, 0xfd, 0xf5, 0x1c, 0x42, 0x6c, 0xda, 0xed, 0x39, 0x5c, 0x6c, 0xe6, 0x10,
	0xae, 0xe7, 0x50, 0xdf, 0x9a, 0xc3, 0xc5, 0x6a, 0x0e, 0xdf, 0x90, 0xfd, 0x04, 0x42, 0x8e, 0xfb,
	0xa5, 0x53, 0xac, 0x54, 0x83, 0x91, 0x04, 0xc2, 0xdf, 0x4b, 0x84, 0x9e, 0x92, 0x23, 0x23, 0x23,
	0x9e, 0x81, 0x81, 0xc4, 0x75, 0x6f, 0xae, 0x50, 0x48, 0x50, 0x78, 0x68, 0x64, 0x34, 0x42, 0x86,
	0x55, 0x04, 0x7d, 0x42, 0x88, 0x59, 0x70, 0x21, 0x63, 0x58, 0xf2, 0x33, 0xec, 0x8c, 0xc7, 0xea,
	0x66, 0x11, 0x38, 0xe0, 0x8c, 0x3e, 0x27, 0x2d, 0xc7, 0x1a, 0xae, 0xa7, 0xd3, 0x5c, 0x5a, 0x7e,
	0x56, 0xd5, 0x65, 0xdf, 0x2c, 0x02, 0xf6, 0x11, 0xb1, 0x33, 0xfa, 0x8c, 0x78, 0x4e, 0x04, 0x16,
	0xf0, 0x8b, 0xd2, 0xaf, 0xba, 0xe3, 0x34, 0x60, 0xc1, 0x7d, 0x3f, 0xfa, 0xf4, 0x11, 0x69, 0x98,
	0x05, 0x2e, 0x14, 0xef, 0x63, 0x7d, 0x3c, 0xb6, 0x67, 0x16, 0x6e, 0x91, 0xfa, 0xf4, 0x47, 0x72,
	0x3c, 0x85, 0xd0, 0x6a, 0xb3, 0xe4, 0x99, 0x91, 0x2e, 0xc6, 0xe9, 0x72, 0xbf, 0xdd, 0xd9, 0xe9,
	0x7a, 0x8c, 0x56, 0xdc, 0x08, 0x29, 0xe7, 0xc8, 0xe9, 0x43, 0x52, 0x4f, 0x60, 0xc1, 0xa5, 0x32,
	0x19, 0x76, 0xc9, 0x63, 0x7b, 0x09, 0x2c, 0xde, 0xbc, 0x67, 0x23, 0xb7, 0x31, 0x8e, 0x12, 0x85,
	0x5d, 0xf2, 0x70, 0x19, 0xc6, 0x12, 0xdb, 0xe4, 0xb1, 0x66, 0x02, 0x8b, 0xa0, 0xb0, 0xcb, 0x0b,
	0x87, 0xd1, 0xe7, 0xc4, 0x5b, 0x6f, 0xcc, 0x9f, 0x5a, 0xa5, 0x55, 0xa5, 0x9a, 0x2b, 0xf0, 0x57,
	0xad, 0x52, 0xfa, 0x98, 0x34, 0xcc, 0x94, 0x1b, 0x19, 0xb9, 0x05, 0x3c, 0xc2, 0x05, 0xac, 0x9b,
	0x29, 0xc3, 0x7b, 0xda, 0x23, 0xc7, 0xeb, 0x11, 0xce, 0xfb, 0x13, 0x65, 0xf9, 0x94, 0x87, 0xa9,
	0xc5, 0x5e, 0xd5, 0xd9, 0xe1, 0x8a, 0x3b, 0xef, 0x0f, 0x94, 0x7d, 0x7b, 0x91, 0x5a, 0x17, 0x99,
	0xc1, 0x32, 0xd6, 0x20, 0x78, 0xa8, 0x85, 0x0c, 0x7d, 0x1f, 0x47, 0x6c, 0x56, 0xe0, 0x85, 0xc3,
	0xe8, 0x4f, 0xe4, 0x64, 0x25, 0x92, 0xa9, 0x93, 0x19, 0x9e, 0x87, 0x46, 0x65, 0xd6, 0x7f, 0x88,
	0xea, 0xe3, 0x8a, 0x7d, 0x53, 0x92, 0x63, 0xe4, 0xb6, 0x5d, 0x42, 0xde, 0x70, 0x3d, 0xba, 0xe1,
	0x0a, 0xe4, 0xb6, 0xeb, 0x05, 0x39, 0x8c, 0xa4, 0x8e, 0x75, 0xc8, 0x27, 0xc5, 0x74, 0x2a, 0x0d,
	0xb7, 0x36, 0xf6, 0x1f, 0xe3, 0x62, 0xb5, 0x4b, 0x62, 0x80, 0xf8, 0xa7, 0x4f, 0x1f, 0xe8, 0x39,
	0x39, 0xa9, 0xb4, 0xee, 0x43, 0x52, 0xe9, 0xf1, 0xf7, 0xf2, 0x04, 0x0d, 0x47, 0x25, 0x3b, 0x54,
	0x69, 0xe9, 0x71, 0x7f, 0x99, 0x17, 0x1d, 0x42, 0xb6, 0x7e, 0x1e, 0x75, 0x72, 0x37, 0x60, 0x1f,
	0x47, 0x07, 0xff, 0x73, 0x57, 0xc3, 0xd7, 0xec, 0xb7, 0x83, 0xda, 0xe0, 0xd5, 0x1f, 0xe7, 0x91,
	0xb2, 0x97, 0xc5, 0xe4, 0x34, 0xd4, 0x49, 0x6f, 0x62, 0x74, 0x08, 0xa6, 0x17, 0x5e, 0x2a, 0x93,
	0xe5, 0x16, 0xc2, 0xd9, 0x4b, 0x77, 0x00, 0x88, 0x74, 0xef, 0xd6, 0x99, 0x60, 0xb2, 0x8b, 0x67,
	0x81, 0xf3, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xc4, 0x66, 0x5d, 0x2d, 0x08, 0x00, 0x00,
}

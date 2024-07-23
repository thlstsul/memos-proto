// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: store/workspace_setting.proto

package store

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

type WorkspaceSettingKey int32

const (
	WorkspaceSettingKey_WORKSPACE_SETTING_KEY_UNSPECIFIED WorkspaceSettingKey = 0
	// BASIC is the key for basic settings.
	WorkspaceSettingKey_BASIC WorkspaceSettingKey = 1
	// GENERAL is the key for general settings.
	WorkspaceSettingKey_GENERAL WorkspaceSettingKey = 2
	// STORAGE is the key for storage settings.
	WorkspaceSettingKey_STORAGE WorkspaceSettingKey = 3
	// MEMO_RELATED is the key for memo related settings.
	WorkspaceSettingKey_MEMO_RELATED WorkspaceSettingKey = 4
)

// Enum value maps for WorkspaceSettingKey.
var (
	WorkspaceSettingKey_name = map[int32]string{
		0: "WORKSPACE_SETTING_KEY_UNSPECIFIED",
		1: "BASIC",
		2: "GENERAL",
		3: "STORAGE",
		4: "MEMO_RELATED",
	}
	WorkspaceSettingKey_value = map[string]int32{
		"WORKSPACE_SETTING_KEY_UNSPECIFIED": 0,
		"BASIC":                             1,
		"GENERAL":                           2,
		"STORAGE":                           3,
		"MEMO_RELATED":                      4,
	}
)

func (x WorkspaceSettingKey) Enum() *WorkspaceSettingKey {
	p := new(WorkspaceSettingKey)
	*p = x
	return p
}

func (x WorkspaceSettingKey) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkspaceSettingKey) Descriptor() protoreflect.EnumDescriptor {
	return file_store_workspace_setting_proto_enumTypes[0].Descriptor()
}

func (WorkspaceSettingKey) Type() protoreflect.EnumType {
	return &file_store_workspace_setting_proto_enumTypes[0]
}

func (x WorkspaceSettingKey) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkspaceSettingKey.Descriptor instead.
func (WorkspaceSettingKey) EnumDescriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{0}
}

type WorkspaceStorageSetting_StorageType int32

const (
	WorkspaceStorageSetting_STORAGE_TYPE_UNSPECIFIED WorkspaceStorageSetting_StorageType = 0
	// STORAGE_TYPE_DATABASE is the database storage type.
	WorkspaceStorageSetting_DATABASE WorkspaceStorageSetting_StorageType = 1
	// STORAGE_TYPE_LOCAL is the local storage type.
	WorkspaceStorageSetting_LOCAL WorkspaceStorageSetting_StorageType = 2
	// STORAGE_TYPE_S3 is the S3 storage type.
	WorkspaceStorageSetting_S3 WorkspaceStorageSetting_StorageType = 3
)

// Enum value maps for WorkspaceStorageSetting_StorageType.
var (
	WorkspaceStorageSetting_StorageType_name = map[int32]string{
		0: "STORAGE_TYPE_UNSPECIFIED",
		1: "DATABASE",
		2: "LOCAL",
		3: "S3",
	}
	WorkspaceStorageSetting_StorageType_value = map[string]int32{
		"STORAGE_TYPE_UNSPECIFIED": 0,
		"DATABASE":                 1,
		"LOCAL":                    2,
		"S3":                       3,
	}
)

func (x WorkspaceStorageSetting_StorageType) Enum() *WorkspaceStorageSetting_StorageType {
	p := new(WorkspaceStorageSetting_StorageType)
	*p = x
	return p
}

func (x WorkspaceStorageSetting_StorageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkspaceStorageSetting_StorageType) Descriptor() protoreflect.EnumDescriptor {
	return file_store_workspace_setting_proto_enumTypes[1].Descriptor()
}

func (WorkspaceStorageSetting_StorageType) Type() protoreflect.EnumType {
	return &file_store_workspace_setting_proto_enumTypes[1]
}

func (x WorkspaceStorageSetting_StorageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkspaceStorageSetting_StorageType.Descriptor instead.
func (WorkspaceStorageSetting_StorageType) EnumDescriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{4, 0}
}

type WorkspaceSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key WorkspaceSettingKey `protobuf:"varint,1,opt,name=key,proto3,enum=memos.store.WorkspaceSettingKey" json:"key,omitempty"`
	// Types that are assignable to Value:
	//
	//	*WorkspaceSetting_BasicSetting
	//	*WorkspaceSetting_GeneralSetting
	//	*WorkspaceSetting_StorageSetting
	//	*WorkspaceSetting_MemoRelatedSetting
	Value isWorkspaceSetting_Value `protobuf_oneof:"value"`
}

func (x *WorkspaceSetting) Reset() {
	*x = WorkspaceSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_workspace_setting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceSetting) ProtoMessage() {}

func (x *WorkspaceSetting) ProtoReflect() protoreflect.Message {
	mi := &file_store_workspace_setting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceSetting.ProtoReflect.Descriptor instead.
func (*WorkspaceSetting) Descriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{0}
}

func (x *WorkspaceSetting) GetKey() WorkspaceSettingKey {
	if x != nil {
		return x.Key
	}
	return WorkspaceSettingKey_WORKSPACE_SETTING_KEY_UNSPECIFIED
}

func (m *WorkspaceSetting) GetValue() isWorkspaceSetting_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *WorkspaceSetting) GetBasicSetting() *WorkspaceBasicSetting {
	if x, ok := x.GetValue().(*WorkspaceSetting_BasicSetting); ok {
		return x.BasicSetting
	}
	return nil
}

func (x *WorkspaceSetting) GetGeneralSetting() *WorkspaceGeneralSetting {
	if x, ok := x.GetValue().(*WorkspaceSetting_GeneralSetting); ok {
		return x.GeneralSetting
	}
	return nil
}

func (x *WorkspaceSetting) GetStorageSetting() *WorkspaceStorageSetting {
	if x, ok := x.GetValue().(*WorkspaceSetting_StorageSetting); ok {
		return x.StorageSetting
	}
	return nil
}

func (x *WorkspaceSetting) GetMemoRelatedSetting() *WorkspaceMemoRelatedSetting {
	if x, ok := x.GetValue().(*WorkspaceSetting_MemoRelatedSetting); ok {
		return x.MemoRelatedSetting
	}
	return nil
}

type isWorkspaceSetting_Value interface {
	isWorkspaceSetting_Value()
}

type WorkspaceSetting_BasicSetting struct {
	BasicSetting *WorkspaceBasicSetting `protobuf:"bytes,2,opt,name=basic_setting,json=basicSetting,proto3,oneof"`
}

type WorkspaceSetting_GeneralSetting struct {
	GeneralSetting *WorkspaceGeneralSetting `protobuf:"bytes,3,opt,name=general_setting,json=generalSetting,proto3,oneof"`
}

type WorkspaceSetting_StorageSetting struct {
	StorageSetting *WorkspaceStorageSetting `protobuf:"bytes,4,opt,name=storage_setting,json=storageSetting,proto3,oneof"`
}

type WorkspaceSetting_MemoRelatedSetting struct {
	MemoRelatedSetting *WorkspaceMemoRelatedSetting `protobuf:"bytes,5,opt,name=memo_related_setting,json=memoRelatedSetting,proto3,oneof"`
}

func (*WorkspaceSetting_BasicSetting) isWorkspaceSetting_Value() {}

func (*WorkspaceSetting_GeneralSetting) isWorkspaceSetting_Value() {}

func (*WorkspaceSetting_StorageSetting) isWorkspaceSetting_Value() {}

func (*WorkspaceSetting_MemoRelatedSetting) isWorkspaceSetting_Value() {}

type WorkspaceBasicSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecretKey string `protobuf:"bytes,1,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
}

func (x *WorkspaceBasicSetting) Reset() {
	*x = WorkspaceBasicSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_workspace_setting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceBasicSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceBasicSetting) ProtoMessage() {}

func (x *WorkspaceBasicSetting) ProtoReflect() protoreflect.Message {
	mi := &file_store_workspace_setting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceBasicSetting.ProtoReflect.Descriptor instead.
func (*WorkspaceBasicSetting) Descriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{1}
}

func (x *WorkspaceBasicSetting) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

type WorkspaceGeneralSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// additional_script is the additional script.
	AdditionalScript string `protobuf:"bytes,3,opt,name=additional_script,json=additionalScript,proto3" json:"additional_script,omitempty"`
	// additional_style is the additional style.
	AdditionalStyle string `protobuf:"bytes,4,opt,name=additional_style,json=additionalStyle,proto3" json:"additional_style,omitempty"`
	// custom_profile is the custom profile.
	CustomProfile *WorkspaceCustomProfile `protobuf:"bytes,5,opt,name=custom_profile,json=customProfile,proto3" json:"custom_profile,omitempty"`
}

func (x *WorkspaceGeneralSetting) Reset() {
	*x = WorkspaceGeneralSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_workspace_setting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceGeneralSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceGeneralSetting) ProtoMessage() {}

func (x *WorkspaceGeneralSetting) ProtoReflect() protoreflect.Message {
	mi := &file_store_workspace_setting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceGeneralSetting.ProtoReflect.Descriptor instead.
func (*WorkspaceGeneralSetting) Descriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{2}
}

func (x *WorkspaceGeneralSetting) GetAdditionalScript() string {
	if x != nil {
		return x.AdditionalScript
	}
	return ""
}

func (x *WorkspaceGeneralSetting) GetAdditionalStyle() string {
	if x != nil {
		return x.AdditionalStyle
	}
	return ""
}

func (x *WorkspaceGeneralSetting) GetCustomProfile() *WorkspaceCustomProfile {
	if x != nil {
		return x.CustomProfile
	}
	return nil
}

type WorkspaceCustomProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	LogoUrl     string `protobuf:"bytes,3,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	Locale      string `protobuf:"bytes,4,opt,name=locale,proto3" json:"locale,omitempty"`
	Appearance  string `protobuf:"bytes,5,opt,name=appearance,proto3" json:"appearance,omitempty"`
}

func (x *WorkspaceCustomProfile) Reset() {
	*x = WorkspaceCustomProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_workspace_setting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceCustomProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceCustomProfile) ProtoMessage() {}

func (x *WorkspaceCustomProfile) ProtoReflect() protoreflect.Message {
	mi := &file_store_workspace_setting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceCustomProfile.ProtoReflect.Descriptor instead.
func (*WorkspaceCustomProfile) Descriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{3}
}

func (x *WorkspaceCustomProfile) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *WorkspaceCustomProfile) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *WorkspaceCustomProfile) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *WorkspaceCustomProfile) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *WorkspaceCustomProfile) GetAppearance() string {
	if x != nil {
		return x.Appearance
	}
	return ""
}

type WorkspaceStorageSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// storage_type is the storage type.
	StorageType WorkspaceStorageSetting_StorageType `protobuf:"varint,1,opt,name=storage_type,json=storageType,proto3,enum=memos.store.WorkspaceStorageSetting_StorageType" json:"storage_type,omitempty"`
	// The template of file path.
	// e.g. assets/{timestamp}_{filename}
	FilepathTemplate string `protobuf:"bytes,2,opt,name=filepath_template,json=filepathTemplate,proto3" json:"filepath_template,omitempty"`
	// The max upload size in megabytes.
	UploadSizeLimitMb int64 `protobuf:"varint,3,opt,name=upload_size_limit_mb,json=uploadSizeLimitMb,proto3" json:"upload_size_limit_mb,omitempty"`
	// The S3 config.
	S3Config *StorageS3Config `protobuf:"bytes,4,opt,name=s3_config,json=s3Config,proto3" json:"s3_config,omitempty"`
}

func (x *WorkspaceStorageSetting) Reset() {
	*x = WorkspaceStorageSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_workspace_setting_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceStorageSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceStorageSetting) ProtoMessage() {}

func (x *WorkspaceStorageSetting) ProtoReflect() protoreflect.Message {
	mi := &file_store_workspace_setting_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceStorageSetting.ProtoReflect.Descriptor instead.
func (*WorkspaceStorageSetting) Descriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{4}
}

func (x *WorkspaceStorageSetting) GetStorageType() WorkspaceStorageSetting_StorageType {
	if x != nil {
		return x.StorageType
	}
	return WorkspaceStorageSetting_STORAGE_TYPE_UNSPECIFIED
}

func (x *WorkspaceStorageSetting) GetFilepathTemplate() string {
	if x != nil {
		return x.FilepathTemplate
	}
	return ""
}

func (x *WorkspaceStorageSetting) GetUploadSizeLimitMb() int64 {
	if x != nil {
		return x.UploadSizeLimitMb
	}
	return 0
}

func (x *WorkspaceStorageSetting) GetS3Config() *StorageS3Config {
	if x != nil {
		return x.S3Config
	}
	return nil
}

// Reference: https://developers.cloudflare.com/r2/examples/aws/aws-sdk-go/
type StorageS3Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessKeyId     string `protobuf:"bytes,1,opt,name=access_key_id,json=accessKeyId,proto3" json:"access_key_id,omitempty"`
	AccessKeySecret string `protobuf:"bytes,2,opt,name=access_key_secret,json=accessKeySecret,proto3" json:"access_key_secret,omitempty"`
	Endpoint        string `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Region          string `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	Bucket          string `protobuf:"bytes,5,opt,name=bucket,proto3" json:"bucket,omitempty"`
}

func (x *StorageS3Config) Reset() {
	*x = StorageS3Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_workspace_setting_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageS3Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageS3Config) ProtoMessage() {}

func (x *StorageS3Config) ProtoReflect() protoreflect.Message {
	mi := &file_store_workspace_setting_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageS3Config.ProtoReflect.Descriptor instead.
func (*StorageS3Config) Descriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{5}
}

func (x *StorageS3Config) GetAccessKeyId() string {
	if x != nil {
		return x.AccessKeyId
	}
	return ""
}

func (x *StorageS3Config) GetAccessKeySecret() string {
	if x != nil {
		return x.AccessKeySecret
	}
	return ""
}

func (x *StorageS3Config) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *StorageS3Config) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *StorageS3Config) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

type WorkspaceMemoRelatedSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// disallow_public_visibility disallows set memo as public visibility.
	DisallowPublicVisibility bool `protobuf:"varint,1,opt,name=disallow_public_visibility,json=disallowPublicVisibility,proto3" json:"disallow_public_visibility,omitempty"`
	// display_with_update_time orders and displays memo with update time.
	DisplayWithUpdateTime bool `protobuf:"varint,2,opt,name=display_with_update_time,json=displayWithUpdateTime,proto3" json:"display_with_update_time,omitempty"`
	// content_length_limit is the limit of content length. Unit is byte.
	ContentLengthLimit int32 `protobuf:"varint,3,opt,name=content_length_limit,json=contentLengthLimit,proto3" json:"content_length_limit,omitempty"`
	// enable_auto_compact enables auto compact for large content.
	EnableAutoCompact bool `protobuf:"varint,4,opt,name=enable_auto_compact,json=enableAutoCompact,proto3" json:"enable_auto_compact,omitempty"`
	// enable_double_click_edit enables editing on double click.
	EnableDoubleClickEdit bool `protobuf:"varint,5,opt,name=enable_double_click_edit,json=enableDoubleClickEdit,proto3" json:"enable_double_click_edit,omitempty"`
	// enable_link_preview enables links preview.
	EnableLinkPreview bool `protobuf:"varint,6,opt,name=enable_link_preview,json=enableLinkPreview,proto3" json:"enable_link_preview,omitempty"`
}

func (x *WorkspaceMemoRelatedSetting) Reset() {
	*x = WorkspaceMemoRelatedSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_workspace_setting_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceMemoRelatedSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceMemoRelatedSetting) ProtoMessage() {}

func (x *WorkspaceMemoRelatedSetting) ProtoReflect() protoreflect.Message {
	mi := &file_store_workspace_setting_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceMemoRelatedSetting.ProtoReflect.Descriptor instead.
func (*WorkspaceMemoRelatedSetting) Descriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{6}
}

func (x *WorkspaceMemoRelatedSetting) GetDisallowPublicVisibility() bool {
	if x != nil {
		return x.DisallowPublicVisibility
	}
	return false
}

func (x *WorkspaceMemoRelatedSetting) GetDisplayWithUpdateTime() bool {
	if x != nil {
		return x.DisplayWithUpdateTime
	}
	return false
}

func (x *WorkspaceMemoRelatedSetting) GetContentLengthLimit() int32 {
	if x != nil {
		return x.ContentLengthLimit
	}
	return 0
}

func (x *WorkspaceMemoRelatedSetting) GetEnableAutoCompact() bool {
	if x != nil {
		return x.EnableAutoCompact
	}
	return false
}

func (x *WorkspaceMemoRelatedSetting) GetEnableDoubleClickEdit() bool {
	if x != nil {
		return x.EnableDoubleClickEdit
	}
	return false
}

func (x *WorkspaceMemoRelatedSetting) GetEnableLinkPreview() bool {
	if x != nil {
		return x.EnableLinkPreview
	}
	return false
}

var File_store_workspace_setting_proto protoreflect.FileDescriptor

var file_store_workspace_setting_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x9a, 0x03, 0x0a,
	0x10, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x32, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20,
	0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x49, 0x0a, 0x0d, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d,
	0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x42, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x48, 0x00, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x4f, 0x0a, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x65, 0x6d, 0x6f,
	0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x48,
	0x00, 0x52, 0x0e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x4f, 0x0a, 0x0f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x65, 0x6d,
	0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x48, 0x00, 0x52, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x5c, 0x0a, 0x14, 0x6d, 0x65, 0x6d, 0x6f, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x12, 0x6d, 0x65,
	0x6d, 0x6f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x36, 0x0a, 0x15, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65,
	0x79, 0x22, 0xbd, 0x01, 0x0a, 0x17, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2b, 0x0a,
	0x11, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x4a, 0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x22, 0xa3, 0x01, 0x0a, 0x16, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x65, 0x61,
	0x72, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x70,
	0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x22, 0xd5, 0x02, 0x0a, 0x17, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x53, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x6d, 0x65, 0x6d, 0x6f,
	0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x66, 0x69, 0x6c, 0x65,
	0x70, 0x61, 0x74, 0x68, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x14, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x62, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x11, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x4d, 0x62, 0x12, 0x39, 0x0a, 0x09, 0x73, 0x33, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x6d, 0x6f,
	0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53,
	0x33, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x73, 0x33, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0x4c, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1c, 0x0a, 0x18, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x53, 0x33, 0x10, 0x03, 0x22,
	0xad, 0x01, 0x0a, 0x0f, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x33, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x22,
	0xdf, 0x02, 0x0a, 0x1b, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4d, 0x65, 0x6d,
	0x6f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x3c, 0x0a, 0x1a, 0x64, 0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x18, 0x64, 0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x37, 0x0a,
	0x18, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x15, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x57, 0x69, 0x74, 0x68, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x75, 0x74,
	0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x12, 0x37, 0x0a, 0x18, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x5f,
	0x65, 0x64, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x45, 0x64, 0x69,
	0x74, 0x12, 0x2e, 0x0a, 0x13, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b,
	0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x2a, 0x73, 0x0a, 0x13, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x21, 0x57, 0x4f, 0x52, 0x4b,
	0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4b, 0x45,
	0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x42, 0x41, 0x53, 0x49, 0x43, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x47, 0x45,
	0x4e, 0x45, 0x52, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x4f, 0x52, 0x41,
	0x47, 0x45, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x45, 0x4d, 0x4f, 0x5f, 0x52, 0x45, 0x4c,
	0x41, 0x54, 0x45, 0x44, 0x10, 0x04, 0x42, 0xa0, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x6d,
	0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x15, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x75, 0x73, 0x65, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0xa2, 0x02,
	0x03, 0x4d, 0x53, 0x58, 0xaa, 0x02, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0xca, 0x02, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0xe2, 0x02, 0x17, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x4d, 0x65, 0x6d,
	0x6f, 0x73, 0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_store_workspace_setting_proto_rawDescOnce sync.Once
	file_store_workspace_setting_proto_rawDescData = file_store_workspace_setting_proto_rawDesc
)

func file_store_workspace_setting_proto_rawDescGZIP() []byte {
	file_store_workspace_setting_proto_rawDescOnce.Do(func() {
		file_store_workspace_setting_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_workspace_setting_proto_rawDescData)
	})
	return file_store_workspace_setting_proto_rawDescData
}

var file_store_workspace_setting_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_store_workspace_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_store_workspace_setting_proto_goTypes = []any{
	(WorkspaceSettingKey)(0),                 // 0: memos.store.WorkspaceSettingKey
	(WorkspaceStorageSetting_StorageType)(0), // 1: memos.store.WorkspaceStorageSetting.StorageType
	(*WorkspaceSetting)(nil),                 // 2: memos.store.WorkspaceSetting
	(*WorkspaceBasicSetting)(nil),            // 3: memos.store.WorkspaceBasicSetting
	(*WorkspaceGeneralSetting)(nil),          // 4: memos.store.WorkspaceGeneralSetting
	(*WorkspaceCustomProfile)(nil),           // 5: memos.store.WorkspaceCustomProfile
	(*WorkspaceStorageSetting)(nil),          // 6: memos.store.WorkspaceStorageSetting
	(*StorageS3Config)(nil),                  // 7: memos.store.StorageS3Config
	(*WorkspaceMemoRelatedSetting)(nil),      // 8: memos.store.WorkspaceMemoRelatedSetting
}
var file_store_workspace_setting_proto_depIdxs = []int32{
	0, // 0: memos.store.WorkspaceSetting.key:type_name -> memos.store.WorkspaceSettingKey
	3, // 1: memos.store.WorkspaceSetting.basic_setting:type_name -> memos.store.WorkspaceBasicSetting
	4, // 2: memos.store.WorkspaceSetting.general_setting:type_name -> memos.store.WorkspaceGeneralSetting
	6, // 3: memos.store.WorkspaceSetting.storage_setting:type_name -> memos.store.WorkspaceStorageSetting
	8, // 4: memos.store.WorkspaceSetting.memo_related_setting:type_name -> memos.store.WorkspaceMemoRelatedSetting
	5, // 5: memos.store.WorkspaceGeneralSetting.custom_profile:type_name -> memos.store.WorkspaceCustomProfile
	1, // 6: memos.store.WorkspaceStorageSetting.storage_type:type_name -> memos.store.WorkspaceStorageSetting.StorageType
	7, // 7: memos.store.WorkspaceStorageSetting.s3_config:type_name -> memos.store.StorageS3Config
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_store_workspace_setting_proto_init() }
func file_store_workspace_setting_proto_init() {
	if File_store_workspace_setting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_workspace_setting_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*WorkspaceSetting); i {
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
		file_store_workspace_setting_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*WorkspaceBasicSetting); i {
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
		file_store_workspace_setting_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*WorkspaceGeneralSetting); i {
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
		file_store_workspace_setting_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*WorkspaceCustomProfile); i {
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
		file_store_workspace_setting_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*WorkspaceStorageSetting); i {
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
		file_store_workspace_setting_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*StorageS3Config); i {
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
		file_store_workspace_setting_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*WorkspaceMemoRelatedSetting); i {
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
	file_store_workspace_setting_proto_msgTypes[0].OneofWrappers = []any{
		(*WorkspaceSetting_BasicSetting)(nil),
		(*WorkspaceSetting_GeneralSetting)(nil),
		(*WorkspaceSetting_StorageSetting)(nil),
		(*WorkspaceSetting_MemoRelatedSetting)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_workspace_setting_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_workspace_setting_proto_goTypes,
		DependencyIndexes: file_store_workspace_setting_proto_depIdxs,
		EnumInfos:         file_store_workspace_setting_proto_enumTypes,
		MessageInfos:      file_store_workspace_setting_proto_msgTypes,
	}.Build()
	File_store_workspace_setting_proto = out.File
	file_store_workspace_setting_proto_rawDesc = nil
	file_store_workspace_setting_proto_goTypes = nil
	file_store_workspace_setting_proto_depIdxs = nil
}

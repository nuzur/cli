// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.28.3
// source: user_team.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserTeamRole int32

const (
	UserTeamRole_USER_TEAM_ROLE_INVALID      UserTeamRole = 0
	UserTeamRole_USER_TEAM_ROLE_ADMIN        UserTeamRole = 1
	UserTeamRole_USER_TEAM_ROLE_DEVELOPER    UserTeamRole = 2
	UserTeamRole_USER_TEAM_ROLE_DATA_MANAGER UserTeamRole = 3
	UserTeamRole_USER_TEAM_ROLE_DATA_ANALYST UserTeamRole = 4
	UserTeamRole_USER_TEAM_ROLE_VIEWER       UserTeamRole = 5
)

// Enum value maps for UserTeamRole.
var (
	UserTeamRole_name = map[int32]string{
		0: "USER_TEAM_ROLE_INVALID",
		1: "USER_TEAM_ROLE_ADMIN",
		2: "USER_TEAM_ROLE_DEVELOPER",
		3: "USER_TEAM_ROLE_DATA_MANAGER",
		4: "USER_TEAM_ROLE_DATA_ANALYST",
		5: "USER_TEAM_ROLE_VIEWER",
	}
	UserTeamRole_value = map[string]int32{
		"USER_TEAM_ROLE_INVALID":      0,
		"USER_TEAM_ROLE_ADMIN":        1,
		"USER_TEAM_ROLE_DEVELOPER":    2,
		"USER_TEAM_ROLE_DATA_MANAGER": 3,
		"USER_TEAM_ROLE_DATA_ANALYST": 4,
		"USER_TEAM_ROLE_VIEWER":       5,
	}
)

func (x UserTeamRole) Enum() *UserTeamRole {
	p := new(UserTeamRole)
	*p = x
	return p
}

func (x UserTeamRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserTeamRole) Descriptor() protoreflect.EnumDescriptor {
	return file_user_team_proto_enumTypes[0].Descriptor()
}

func (UserTeamRole) Type() protoreflect.EnumType {
	return &file_user_team_proto_enumTypes[0]
}

func (x UserTeamRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserTeamRole.Descriptor instead.
func (UserTeamRole) EnumDescriptor() ([]byte, []int) {
	return file_user_team_proto_rawDescGZIP(), []int{0}
}

type UserTeamStatus int32

const (
	UserTeamStatus_USER_TEAM_STATUS_INVALID  UserTeamStatus = 0
	UserTeamStatus_USER_TEAM_STATUS_ACTIVE   UserTeamStatus = 1
	UserTeamStatus_USER_TEAM_STATUS_DISABLED UserTeamStatus = 2
	UserTeamStatus_USER_TEAM_STATUS_INVITED  UserTeamStatus = 3
)

// Enum value maps for UserTeamStatus.
var (
	UserTeamStatus_name = map[int32]string{
		0: "USER_TEAM_STATUS_INVALID",
		1: "USER_TEAM_STATUS_ACTIVE",
		2: "USER_TEAM_STATUS_DISABLED",
		3: "USER_TEAM_STATUS_INVITED",
	}
	UserTeamStatus_value = map[string]int32{
		"USER_TEAM_STATUS_INVALID":  0,
		"USER_TEAM_STATUS_ACTIVE":   1,
		"USER_TEAM_STATUS_DISABLED": 2,
		"USER_TEAM_STATUS_INVITED":  3,
	}
)

func (x UserTeamStatus) Enum() *UserTeamStatus {
	p := new(UserTeamStatus)
	*p = x
	return p
}

func (x UserTeamStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserTeamStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_user_team_proto_enumTypes[1].Descriptor()
}

func (UserTeamStatus) Type() protoreflect.EnumType {
	return &file_user_team_proto_enumTypes[1]
}

func (x UserTeamStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserTeamStatus.Descriptor instead.
func (UserTeamStatus) EnumDescriptor() ([]byte, []int) {
	return file_user_team_proto_rawDescGZIP(), []int{1}
}

type UserTeam struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid      string                 `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	UserEmail     string                 `protobuf:"bytes,3,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	TeamUuid      string                 `protobuf:"bytes,4,opt,name=team_uuid,json=teamUuid,proto3" json:"team_uuid,omitempty"`
	Roles         []UserTeamRole         `protobuf:"varint,5,rep,packed,name=roles,proto3,enum=nem.UserTeamRole" json:"roles,omitempty"`
	Status        UserTeamStatus         `protobuf:"varint,6,opt,name=status,proto3,enum=nem.UserTeamStatus" json:"status,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedByUuid string                 `protobuf:"bytes,9,opt,name=created_by_uuid,json=createdByUuid,proto3" json:"created_by_uuid,omitempty"`
	UpdatedByUuid string                 `protobuf:"bytes,10,opt,name=updated_by_uuid,json=updatedByUuid,proto3" json:"updated_by_uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserTeam) Reset() {
	*x = UserTeam{}
	mi := &file_user_team_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserTeam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTeam) ProtoMessage() {}

func (x *UserTeam) ProtoReflect() protoreflect.Message {
	mi := &file_user_team_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTeam.ProtoReflect.Descriptor instead.
func (*UserTeam) Descriptor() ([]byte, []int) {
	return file_user_team_proto_rawDescGZIP(), []int{0}
}

func (x *UserTeam) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *UserTeam) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *UserTeam) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *UserTeam) GetTeamUuid() string {
	if x != nil {
		return x.TeamUuid
	}
	return ""
}

func (x *UserTeam) GetRoles() []UserTeamRole {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *UserTeam) GetStatus() UserTeamStatus {
	if x != nil {
		return x.Status
	}
	return UserTeamStatus_USER_TEAM_STATUS_INVALID
}

func (x *UserTeam) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserTeam) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UserTeam) GetCreatedByUuid() string {
	if x != nil {
		return x.CreatedByUuid
	}
	return ""
}

func (x *UserTeam) GetUpdatedByUuid() string {
	if x != nil {
		return x.UpdatedByUuid
	}
	return ""
}

var File_user_team_proto protoreflect.FileDescriptor

var file_user_team_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x6e, 0x65, 0x6d, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x03, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x65, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x55, 0x75, 0x69,
	0x64, 0x12, 0x27, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6e, 0x65, 0x6d,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x26, 0x0a,
	0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x55, 0x75, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x62, 0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64, 0x2a, 0xbf, 0x01,
	0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x16, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x52, 0x4f, 0x4c, 0x45,
	0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x41, 0x44, 0x4d,
	0x49, 0x4e, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x45, 0x41,
	0x4d, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x45, 0x52,
	0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f,
	0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45,
	0x52, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x45, 0x41, 0x4d,
	0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x41, 0x4e, 0x41, 0x4c, 0x59,
	0x53, 0x54, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x45, 0x41,
	0x4d, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x45, 0x52, 0x10, 0x05, 0x2a,
	0x88, 0x01, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1c, 0x0a, 0x18, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x1b, 0x0a, 0x17, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x1d, 0x0a,
	0x19, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x44, 0x10, 0x03, 0x42, 0x1e, 0x0a, 0x03, 0x6e, 0x65,
	0x6d, 0x42, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x50, 0x01, 0x5a, 0x0b, 0x6e,
	0x65, 0x6d, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_user_team_proto_rawDescOnce sync.Once
	file_user_team_proto_rawDescData = file_user_team_proto_rawDesc
)

func file_user_team_proto_rawDescGZIP() []byte {
	file_user_team_proto_rawDescOnce.Do(func() {
		file_user_team_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_team_proto_rawDescData)
	})
	return file_user_team_proto_rawDescData
}

var file_user_team_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_user_team_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_team_proto_goTypes = []any{
	(UserTeamRole)(0),             // 0: nem.UserTeamRole
	(UserTeamStatus)(0),           // 1: nem.UserTeamStatus
	(*UserTeam)(nil),              // 2: nem.UserTeam
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_user_team_proto_depIdxs = []int32{
	0, // 0: nem.UserTeam.roles:type_name -> nem.UserTeamRole
	1, // 1: nem.UserTeam.status:type_name -> nem.UserTeamStatus
	3, // 2: nem.UserTeam.created_at:type_name -> google.protobuf.Timestamp
	3, // 3: nem.UserTeam.updated_at:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_user_team_proto_init() }
func file_user_team_proto_init() {
	if File_user_team_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_team_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_team_proto_goTypes,
		DependencyIndexes: file_user_team_proto_depIdxs,
		EnumInfos:         file_user_team_proto_enumTypes,
		MessageInfos:      file_user_team_proto_msgTypes,
	}.Build()
	File_user_team_proto = out.File
	file_user_team_proto_rawDesc = nil
	file_user_team_proto_goTypes = nil
	file_user_team_proto_depIdxs = nil
}

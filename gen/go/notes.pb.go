// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: notes.proto

package notes

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Note struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Note) Reset() {
	*x = Note{}
	mi := &file_notes_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{0}
}

func (x *Note) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Note) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Note) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Note) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Note) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateNoteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateNoteRequest) Reset() {
	*x = CreateNoteRequest{}
	mi := &file_notes_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNoteRequest) ProtoMessage() {}

func (x *CreateNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNoteRequest.ProtoReflect.Descriptor instead.
func (*CreateNoteRequest) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNoteRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateNoteRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CreateNoteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Note          *Note                  `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateNoteResponse) Reset() {
	*x = CreateNoteResponse{}
	mi := &file_notes_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNoteResponse) ProtoMessage() {}

func (x *CreateNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNoteResponse.ProtoReflect.Descriptor instead.
func (*CreateNoteResponse) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{2}
}

func (x *CreateNoteResponse) GetNote() *Note {
	if x != nil {
		return x.Note
	}
	return nil
}

type GetNoteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNoteRequest) Reset() {
	*x = GetNoteRequest{}
	mi := &file_notes_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoteRequest) ProtoMessage() {}

func (x *GetNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNoteRequest.ProtoReflect.Descriptor instead.
func (*GetNoteRequest) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{3}
}

func (x *GetNoteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetNoteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Note          *Note                  `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNoteResponse) Reset() {
	*x = GetNoteResponse{}
	mi := &file_notes_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoteResponse) ProtoMessage() {}

func (x *GetNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNoteResponse.ProtoReflect.Descriptor instead.
func (*GetNoteResponse) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{4}
}

func (x *GetNoteResponse) GetNote() *Note {
	if x != nil {
		return x.Note
	}
	return nil
}

type ListNotesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int32                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListNotesRequest) Reset() {
	*x = ListNotesRequest{}
	mi := &file_notes_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListNotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotesRequest) ProtoMessage() {}

func (x *ListNotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotesRequest.ProtoReflect.Descriptor instead.
func (*ListNotesRequest) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{5}
}

func (x *ListNotesRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListNotesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListNotesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Notes         []*Note                `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListNotesResponse) Reset() {
	*x = ListNotesResponse{}
	mi := &file_notes_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListNotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotesResponse) ProtoMessage() {}

func (x *ListNotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotesResponse.ProtoReflect.Descriptor instead.
func (*ListNotesResponse) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{6}
}

func (x *ListNotesResponse) GetNotes() []*Note {
	if x != nil {
		return x.Notes
	}
	return nil
}

func (x *ListNotesResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type UpdateNoteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateNoteRequest) Reset() {
	*x = UpdateNoteRequest{}
	mi := &file_notes_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNoteRequest) ProtoMessage() {}

func (x *UpdateNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNoteRequest.ProtoReflect.Descriptor instead.
func (*UpdateNoteRequest) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateNoteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateNoteRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateNoteRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UpdateNoteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Note          *Note                  `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateNoteResponse) Reset() {
	*x = UpdateNoteResponse{}
	mi := &file_notes_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNoteResponse) ProtoMessage() {}

func (x *UpdateNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNoteResponse.ProtoReflect.Descriptor instead.
func (*UpdateNoteResponse) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateNoteResponse) GetNote() *Note {
	if x != nil {
		return x.Note
	}
	return nil
}

type DeleteNoteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteNoteRequest) Reset() {
	*x = DeleteNoteRequest{}
	mi := &file_notes_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNoteRequest) ProtoMessage() {}

func (x *DeleteNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNoteRequest.ProtoReflect.Descriptor instead.
func (*DeleteNoteRequest) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteNoteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteNoteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteNoteResponse) Reset() {
	*x = DeleteNoteResponse{}
	mi := &file_notes_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNoteResponse) ProtoMessage() {}

func (x *DeleteNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNoteResponse.ProtoReflect.Descriptor instead.
func (*DeleteNoteResponse) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteNoteResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_notes_proto protoreflect.FileDescriptor

const file_notes_proto_rawDesc = "" +
	"\n" +
	"\vnotes.proto\x12\x05notes\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xbc\x01\n" +
	"\x04Note\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x03 \x01(\tR\acontent\x129\n" +
	"\n" +
	"created_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"C\n" +
	"\x11CreateNoteRequest\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x02 \x01(\tR\acontent\"5\n" +
	"\x12CreateNoteResponse\x12\x1f\n" +
	"\x04note\x18\x01 \x01(\v2\v.notes.NoteR\x04note\" \n" +
	"\x0eGetNoteRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"2\n" +
	"\x0fGetNoteResponse\x12\x1f\n" +
	"\x04note\x18\x01 \x01(\v2\v.notes.NoteR\x04note\"C\n" +
	"\x10ListNotesRequest\x12\x12\n" +
	"\x04page\x18\x01 \x01(\x05R\x04page\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\x05R\bpageSize\"W\n" +
	"\x11ListNotesResponse\x12!\n" +
	"\x05notes\x18\x01 \x03(\v2\v.notes.NoteR\x05notes\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount\"S\n" +
	"\x11UpdateNoteRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x03 \x01(\tR\acontent\"5\n" +
	"\x12UpdateNoteResponse\x12\x1f\n" +
	"\x04note\x18\x01 \x01(\v2\v.notes.NoteR\x04note\"#\n" +
	"\x11DeleteNoteRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"$\n" +
	"\x12DeleteNoteResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id2\xc5\x03\n" +
	"\fNotesService\x12W\n" +
	"\n" +
	"CreateNote\x12\x18.notes.CreateNoteRequest\x1a\x19.notes.CreateNoteResponse\"\x14\x82\xd3\xe4\x93\x02\x0e:\x01*\"\t/v1/notes\x12P\n" +
	"\aGetNote\x12\x15.notes.GetNoteRequest\x1a\x16.notes.GetNoteResponse\"\x16\x82\xd3\xe4\x93\x02\x10\x12\x0e/v1/notes/{id}\x12Q\n" +
	"\tListNotes\x12\x17.notes.ListNotesRequest\x1a\x18.notes.ListNotesResponse\"\x11\x82\xd3\xe4\x93\x02\v\x12\t/v1/notes\x12\\\n" +
	"\n" +
	"UpdateNote\x12\x18.notes.UpdateNoteRequest\x1a\x19.notes.UpdateNoteResponse\"\x19\x82\xd3\xe4\x93\x02\x13:\x01*\x1a\x0e/v1/notes/{id}\x12Y\n" +
	"\n" +
	"DeleteNote\x12\x18.notes.DeleteNoteRequest\x1a\x19.notes.DeleteNoteResponse\"\x16\x82\xd3\xe4\x93\x02\x10*\x0e/v1/notes/{id}B9Z7github.com/singh1203/go-notes-service/proto/notes;notesb\x06proto3"

var (
	file_notes_proto_rawDescOnce sync.Once
	file_notes_proto_rawDescData []byte
)

func file_notes_proto_rawDescGZIP() []byte {
	file_notes_proto_rawDescOnce.Do(func() {
		file_notes_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_notes_proto_rawDesc), len(file_notes_proto_rawDesc)))
	})
	return file_notes_proto_rawDescData
}

var file_notes_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_notes_proto_goTypes = []any{
	(*Note)(nil),                  // 0: notes.Note
	(*CreateNoteRequest)(nil),     // 1: notes.CreateNoteRequest
	(*CreateNoteResponse)(nil),    // 2: notes.CreateNoteResponse
	(*GetNoteRequest)(nil),        // 3: notes.GetNoteRequest
	(*GetNoteResponse)(nil),       // 4: notes.GetNoteResponse
	(*ListNotesRequest)(nil),      // 5: notes.ListNotesRequest
	(*ListNotesResponse)(nil),     // 6: notes.ListNotesResponse
	(*UpdateNoteRequest)(nil),     // 7: notes.UpdateNoteRequest
	(*UpdateNoteResponse)(nil),    // 8: notes.UpdateNoteResponse
	(*DeleteNoteRequest)(nil),     // 9: notes.DeleteNoteRequest
	(*DeleteNoteResponse)(nil),    // 10: notes.DeleteNoteResponse
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_notes_proto_depIdxs = []int32{
	11, // 0: notes.Note.created_at:type_name -> google.protobuf.Timestamp
	11, // 1: notes.Note.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 2: notes.CreateNoteResponse.note:type_name -> notes.Note
	0,  // 3: notes.GetNoteResponse.note:type_name -> notes.Note
	0,  // 4: notes.ListNotesResponse.notes:type_name -> notes.Note
	0,  // 5: notes.UpdateNoteResponse.note:type_name -> notes.Note
	1,  // 6: notes.NotesService.CreateNote:input_type -> notes.CreateNoteRequest
	3,  // 7: notes.NotesService.GetNote:input_type -> notes.GetNoteRequest
	5,  // 8: notes.NotesService.ListNotes:input_type -> notes.ListNotesRequest
	7,  // 9: notes.NotesService.UpdateNote:input_type -> notes.UpdateNoteRequest
	9,  // 10: notes.NotesService.DeleteNote:input_type -> notes.DeleteNoteRequest
	2,  // 11: notes.NotesService.CreateNote:output_type -> notes.CreateNoteResponse
	4,  // 12: notes.NotesService.GetNote:output_type -> notes.GetNoteResponse
	6,  // 13: notes.NotesService.ListNotes:output_type -> notes.ListNotesResponse
	8,  // 14: notes.NotesService.UpdateNote:output_type -> notes.UpdateNoteResponse
	10, // 15: notes.NotesService.DeleteNote:output_type -> notes.DeleteNoteResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_notes_proto_init() }
func file_notes_proto_init() {
	if File_notes_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_notes_proto_rawDesc), len(file_notes_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notes_proto_goTypes,
		DependencyIndexes: file_notes_proto_depIdxs,
		MessageInfos:      file_notes_proto_msgTypes,
	}.Build()
	File_notes_proto = out.File
	file_notes_proto_goTypes = nil
	file_notes_proto_depIdxs = nil
}

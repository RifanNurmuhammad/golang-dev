// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: grpc/proto/movie.proto

package proto

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

type MovieQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination string `protobuf:"bytes,1,opt,name=Pagination,proto3" json:"Pagination,omitempty"`
	SearchWord string `protobuf:"bytes,2,opt,name=SearchWord,proto3" json:"SearchWord,omitempty"`
}

func (x *MovieQuery) Reset() {
	*x = MovieQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieQuery) ProtoMessage() {}

func (x *MovieQuery) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieQuery.ProtoReflect.Descriptor instead.
func (*MovieQuery) Descriptor() ([]byte, []int) {
	return file_grpc_proto_movie_proto_rawDescGZIP(), []int{0}
}

func (x *MovieQuery) GetPagination() string {
	if x != nil {
		return x.Pagination
	}
	return ""
}

func (x *MovieQuery) GetSearchWord() string {
	if x != nil {
		return x.SearchWord
	}
	return ""
}

type MovieDetailQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *MovieDetailQuery) Reset() {
	*x = MovieDetailQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieDetailQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieDetailQuery) ProtoMessage() {}

func (x *MovieDetailQuery) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieDetailQuery.ProtoReflect.Descriptor instead.
func (*MovieDetailQuery) Descriptor() ([]byte, []int) {
	return file_grpc_proto_movie_proto_rawDescGZIP(), []int{1}
}

func (x *MovieDetailQuery) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type MovieList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Year   string `protobuf:"bytes,2,opt,name=Year,proto3" json:"Year,omitempty"`
	ImdbID string `protobuf:"bytes,3,opt,name=ImdbID,proto3" json:"ImdbID,omitempty"`
	Type   string `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Poster string `protobuf:"bytes,5,opt,name=Poster,proto3" json:"Poster,omitempty"`
}

func (x *MovieList) Reset() {
	*x = MovieList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_movie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieList) ProtoMessage() {}

func (x *MovieList) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_movie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieList.ProtoReflect.Descriptor instead.
func (*MovieList) Descriptor() ([]byte, []int) {
	return file_grpc_proto_movie_proto_rawDescGZIP(), []int{2}
}

func (x *MovieList) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieList) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *MovieList) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *MovieList) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MovieList) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

type MovieDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Year       string `protobuf:"bytes,2,opt,name=Year,proto3" json:"Year,omitempty"`
	Rated      string `protobuf:"bytes,3,opt,name=Rated,proto3" json:"Rated,omitempty"`
	Released   string `protobuf:"bytes,4,opt,name=Released,proto3" json:"Released,omitempty"`
	Runtime    string `protobuf:"bytes,5,opt,name=Runtime,proto3" json:"Runtime,omitempty"`
	Genre      string `protobuf:"bytes,6,opt,name=Genre,proto3" json:"Genre,omitempty"`
	Director   string `protobuf:"bytes,7,opt,name=Director,proto3" json:"Director,omitempty"`
	Writer     string `protobuf:"bytes,8,opt,name=Writer,proto3" json:"Writer,omitempty"`
	Actors     string `protobuf:"bytes,9,opt,name=Actors,proto3" json:"Actors,omitempty"`
	Plot       string `protobuf:"bytes,10,opt,name=Plot,proto3" json:"Plot,omitempty"`
	Language   string `protobuf:"bytes,11,opt,name=Language,proto3" json:"Language,omitempty"`
	Country    string `protobuf:"bytes,12,opt,name=Country,proto3" json:"Country,omitempty"`
	Awards     string `protobuf:"bytes,13,opt,name=Awards,proto3" json:"Awards,omitempty"`
	Poster     string `protobuf:"bytes,14,opt,name=Poster,proto3" json:"Poster,omitempty"`
	MetaScore  string `protobuf:"bytes,15,opt,name=MetaScore,proto3" json:"MetaScore,omitempty"`
	ImdbRating string `protobuf:"bytes,16,opt,name=ImdbRating,proto3" json:"ImdbRating,omitempty"`
	ImdbVotes  string `protobuf:"bytes,17,opt,name=ImdbVotes,proto3" json:"ImdbVotes,omitempty"`
	ImdbID     string `protobuf:"bytes,18,opt,name=ImdbID,proto3" json:"ImdbID,omitempty"`
	Type       string `protobuf:"bytes,19,opt,name=Type,proto3" json:"Type,omitempty"`
	DVD        string `protobuf:"bytes,20,opt,name=DVD,proto3" json:"DVD,omitempty"`
	BoxOffice  string `protobuf:"bytes,21,opt,name=BoxOffice,proto3" json:"BoxOffice,omitempty"`
	Production string `protobuf:"bytes,22,opt,name=Production,proto3" json:"Production,omitempty"`
	Website    string `protobuf:"bytes,23,opt,name=Website,proto3" json:"Website,omitempty"`
}

func (x *MovieDetail) Reset() {
	*x = MovieDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_movie_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieDetail) ProtoMessage() {}

func (x *MovieDetail) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_movie_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieDetail.ProtoReflect.Descriptor instead.
func (*MovieDetail) Descriptor() ([]byte, []int) {
	return file_grpc_proto_movie_proto_rawDescGZIP(), []int{3}
}

func (x *MovieDetail) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieDetail) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *MovieDetail) GetRated() string {
	if x != nil {
		return x.Rated
	}
	return ""
}

func (x *MovieDetail) GetReleased() string {
	if x != nil {
		return x.Released
	}
	return ""
}

func (x *MovieDetail) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *MovieDetail) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *MovieDetail) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *MovieDetail) GetWriter() string {
	if x != nil {
		return x.Writer
	}
	return ""
}

func (x *MovieDetail) GetActors() string {
	if x != nil {
		return x.Actors
	}
	return ""
}

func (x *MovieDetail) GetPlot() string {
	if x != nil {
		return x.Plot
	}
	return ""
}

func (x *MovieDetail) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *MovieDetail) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *MovieDetail) GetAwards() string {
	if x != nil {
		return x.Awards
	}
	return ""
}

func (x *MovieDetail) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

func (x *MovieDetail) GetMetaScore() string {
	if x != nil {
		return x.MetaScore
	}
	return ""
}

func (x *MovieDetail) GetImdbRating() string {
	if x != nil {
		return x.ImdbRating
	}
	return ""
}

func (x *MovieDetail) GetImdbVotes() string {
	if x != nil {
		return x.ImdbVotes
	}
	return ""
}

func (x *MovieDetail) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *MovieDetail) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MovieDetail) GetDVD() string {
	if x != nil {
		return x.DVD
	}
	return ""
}

func (x *MovieDetail) GetBoxOffice() string {
	if x != nil {
		return x.BoxOffice
	}
	return ""
}

func (x *MovieDetail) GetProduction() string {
	if x != nil {
		return x.Production
	}
	return ""
}

func (x *MovieDetail) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

type ResponseMovieList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie []*MovieList `protobuf:"bytes,1,rep,name=movie,proto3" json:"movie,omitempty"`
}

func (x *ResponseMovieList) Reset() {
	*x = ResponseMovieList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_movie_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMovieList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMovieList) ProtoMessage() {}

func (x *ResponseMovieList) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_movie_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMovieList.ProtoReflect.Descriptor instead.
func (*ResponseMovieList) Descriptor() ([]byte, []int) {
	return file_grpc_proto_movie_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseMovieList) GetMovie() []*MovieList {
	if x != nil {
		return x.Movie
	}
	return nil
}

var File_grpc_proto_movie_proto protoreflect.FileDescriptor

var file_grpc_proto_movie_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x22,
	0x4c, 0x0a, 0x0a, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1e, 0x0a,
	0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x57, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x57, 0x6f, 0x72, 0x64, 0x22, 0x22, 0x0a,
	0x10, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x22, 0x79, 0x0a, 0x09, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x6d, 0x64, 0x62,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x49, 0x6d, 0x64, 0x62, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x22, 0xd1, 0x04, 0x0a,
	0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6c, 0x6f, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x50, 0x6c, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x41, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x41, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x49, 0x6d, 0x64, 0x62, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x49, 0x6d, 0x64, 0x62, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09,
	0x49, 0x6d, 0x64, 0x62, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x49, 0x6d, 0x64, 0x62, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x6d,
	0x64, 0x62, 0x49, 0x44, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x49, 0x6d, 0x64, 0x62,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x44, 0x56, 0x44, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x44, 0x56, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6f, 0x78, 0x4f,
	0x66, 0x66, 0x69, 0x63, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x6f, 0x78,
	0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x22, 0x3b, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x32, 0x86, 0x01,
	0x0a, 0x0c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37,
	0x0a, 0x08, 0x67, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x11, 0x2e, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x18, 0x2e,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0e, 0x67, 0x65, 0x74, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x2e, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x66, 0x61, 0x6e, 0x6e, 0x75, 0x72, 0x6d, 0x75, 0x68,
	0x61, 0x6d, 0x6d, 0x61, 0x64, 0x2f, 0x30, 0x32, 0x2d, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_movie_proto_rawDescOnce sync.Once
	file_grpc_proto_movie_proto_rawDescData = file_grpc_proto_movie_proto_rawDesc
)

func file_grpc_proto_movie_proto_rawDescGZIP() []byte {
	file_grpc_proto_movie_proto_rawDescOnce.Do(func() {
		file_grpc_proto_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_movie_proto_rawDescData)
	})
	return file_grpc_proto_movie_proto_rawDescData
}

var file_grpc_proto_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grpc_proto_movie_proto_goTypes = []interface{}{
	(*MovieQuery)(nil),        // 0: movie.MovieQuery
	(*MovieDetailQuery)(nil),  // 1: movie.MovieDetailQuery
	(*MovieList)(nil),         // 2: movie.MovieList
	(*MovieDetail)(nil),       // 3: movie.MovieDetail
	(*ResponseMovieList)(nil), // 4: movie.ResponseMovieList
}
var file_grpc_proto_movie_proto_depIdxs = []int32{
	2, // 0: movie.ResponseMovieList.movie:type_name -> movie.MovieList
	0, // 1: movie.MovieService.getMovie:input_type -> movie.MovieQuery
	1, // 2: movie.MovieService.getMovieDetail:input_type -> movie.MovieDetailQuery
	4, // 3: movie.MovieService.getMovie:output_type -> movie.ResponseMovieList
	3, // 4: movie.MovieService.getMovieDetail:output_type -> movie.MovieDetail
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_proto_movie_proto_init() }
func file_grpc_proto_movie_proto_init() {
	if File_grpc_proto_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieQuery); i {
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
		file_grpc_proto_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieDetailQuery); i {
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
		file_grpc_proto_movie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieList); i {
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
		file_grpc_proto_movie_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieDetail); i {
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
		file_grpc_proto_movie_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMovieList); i {
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
			RawDescriptor: file_grpc_proto_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_movie_proto_goTypes,
		DependencyIndexes: file_grpc_proto_movie_proto_depIdxs,
		MessageInfos:      file_grpc_proto_movie_proto_msgTypes,
	}.Build()
	File_grpc_proto_movie_proto = out.File
	file_grpc_proto_movie_proto_rawDesc = nil
	file_grpc_proto_movie_proto_goTypes = nil
	file_grpc_proto_movie_proto_depIdxs = nil
}
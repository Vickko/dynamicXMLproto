package protofactory

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

func Test_If_ProtoFileBuilder_IsInit(t *testing.T) {
	literal := ProtoFileBuilder{}
	newVari := NewFileBuilder()
	fmt.Println(literal, newVari)
	// cannot compare, maybe using cmp package
	//fmt.Println(literal == newVari)
}

func Test_Sets(t *testing.T) {
	b := NewFileBuilder()
	fmt.Println(b)
	b.SetSyntax("proto3")
	fmt.Println(b)
}

func Test_Gets(t *testing.T) {
	b := NewFileBuilder()
	fmt.Println(b.GetSyntax())
	b.SetSyntax("proto3")
	fmt.Println(b.GetSyntax())
	b.SetSyntax("proto3666")
	fmt.Println(b.GetSyntax())
}

func Test_build(t *testing.T) {
	f := NewFileBuilder().SetConfig(
		"ai.momenta/hdmap/data/mdf/test.proto",
		"proto3",
		"hdmap.data.mdf",
		"ai.momenta/hdmap/data/mdf",
	).AppendMsg(
		NewMsgBuilder().Set("Man",
			NewField("String", "firstName", 1),
			NewField("String", "lastName", 2),
			NewField("Bool", "isAlive", 3),
			NewField("Int64", "age", 4),
		).ExportMsg(),
		NewMsgBuilder().Set("Man2",
			NewField("String", "firstName", 1),
			NewField("String", "lastName", 2),
			NewField("Bool", "isAlive", 3),
			NewField("Int64", "age", 4),
		).ExportMsg(),
		NewMsgBuilder().Set("TestId",
			NewField("Int64", "id", 1),
		).ExportMsg(),
	)
	f.Build()
	fmt.Println(f)
	fd, _ := protodesc.NewFile(&f.FileDescriptorProto, &protoregistry.Files{})
	md := fd.Messages().ByName(protoreflect.Name("Man"))
	man1 := dynamicpb.NewMessage(md)
	fmt.Println("man1", man1)
	man1.Set(md.Fields().ByJSONName("firstName"), protoreflect.ValueOfString("John"))
	man1.Set(md.Fields().ByJSONName("lastName"), protoreflect.ValueOfString("Smith"))
	man1.Set(md.Fields().ByJSONName("isAlive"), protoreflect.ValueOfBool(true))
	man1.Set(md.Fields().ByJSONName("age"), protoreflect.ValueOfInt64(27))
	fmt.Println("man1", man1)

	fmt.Println(md.Fields().ByJSONName("firstName").Kind())
}

// This example build protobuf msg type and instance on runtime using wrapped protoFactory API
func Test_build_wrapped(t *testing.T) {
	f := NewFileBuilder().SetConfig(
		"ai.momenta/hdmap/data/mdf/test.proto",
		"proto3",
		"hdmap.data.mdf",
		"ai.momenta/hdmap/data/mdf",
	).AppendMsg(
		NewMsgBuilder().Set("Man",
			NewField("String", "firstName", 1),
			NewField("String", "lastName", 2),
			NewField("Bool", "isAlive", 3),
			NewField("Int64", "age", 4),
		).ExportMsg(),
	)
	man1 := f.MsgInstance("Man")
	man1.SetField("firstName", "John")
	man1.SetField("lastName", "Smith")
	man1.SetField("isAlive", true)
	man1.SetField("age", int64(27))
	fmt.Println(man1)
}


// This example build protobuf msg type and instance on runtime using original official protoreflect API
func Test_build_original(t *testing.T) {
	firstName := &descriptorpb.FieldDescriptorProto{
		Name:     proto.String("firstName"),
		Number:   proto.Int32(1),
		Label:    optional,
		Type:     descriptorpb.FieldDescriptorProto_Type(protoreflect.StringKind).Enum(),
		JsonName: proto.String("firstName"),
	}
	lastName := &descriptorpb.FieldDescriptorProto{
		Name:     proto.String("lastName"),
		Number:   proto.Int32(2),
		Label:    optional,
		Type:     descriptorpb.FieldDescriptorProto_Type(protoreflect.StringKind).Enum(),
		JsonName: proto.String("lastName"),
	}
	isAlive := &descriptorpb.FieldDescriptorProto{
		Name:     proto.String("isAlive"),
		Number:   proto.Int32(3),
		Label:    optional,
		Type:     descriptorpb.FieldDescriptorProto_Type(protoreflect.BoolKind).Enum(),
		JsonName: proto.String("isAlive"),
	}
	age := &descriptorpb.FieldDescriptorProto{
		Name:     proto.String("age"),
		Number:   proto.Int32(4),
		Label:    optional,
		Type:     descriptorpb.FieldDescriptorProto_Type(protoreflect.Int64Kind).Enum(),
		JsonName: proto.String("age"),
	}
	m1 := descriptorpb.DescriptorProto{
		Name:    proto.String("Man"),
		Options: new(descriptorpb.MessageOptions),
	}
	m1.Field = append(m1.Field, firstName, lastName, isAlive, age)
	f := descriptorpb.FileDescriptorProto{
		Name:    proto.String("ai.momenta/hdmap/data/mdf/test.proto"),
		Syntax:  proto.String("proto3"),
		Package: proto.String("hdmap.data.mdf"),
		Options: new(descriptorpb.FileOptions),
	}
	f.Options.GoPackage = proto.String("hdmap.data.mdf")
	f.MessageType = append(f.MessageType, &m1)
	fd, _ := protodesc.NewFile(&f, &protoregistry.Files{})
	md := fd.Messages().ByName(protoreflect.Name("Man"))
	man1 := dynamicpb.NewMessage(md)
	man1.Set(md.Fields().ByJSONName("firstName"), protoreflect.ValueOfString("John"))
	man1.Set(md.Fields().ByJSONName("lastName"), protoreflect.ValueOfString("Smith"))
	man1.Set(md.Fields().ByJSONName("isAlive"), protoreflect.ValueOfBool(true))
	man1.Set(md.Fields().ByJSONName("age"), protoreflect.ValueOfInt64(27))
	fmt.Println("man1", man1)
}

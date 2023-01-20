package dynamicxmlproto

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
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
	f := NewFileBuilder().Set(
		"ai.momenta/hdmap/data/mdf/test.proto",
		"proto3",
		"hdmap.data.mdf",
		"ai.momenta/hdmap/data/mdf",
	)

	f.AppendMsg(
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
	fmt.Println(f)
	fd, _ := protodesc.NewFile(&f.FileDescriptorProto, &protoregistry.Files{})
	md := fd.Messages().ByName(protoreflect.Name("Man"))
	man1 := dynamicpb.NewMessage(md)
	fmt.Println("man1",man1)
	man1.Set(md.Fields().ByJSONName("firstName"), protoreflect.ValueOfString("John"))
	man1.Set(md.Fields().ByJSONName("lastName"), protoreflect.ValueOfString("Smith"))
	man1.Set(md.Fields().ByJSONName("isAlive"), protoreflect.ValueOfBool(true))
	man1.Set(md.Fields().ByJSONName("age"), protoreflect.ValueOfInt64(27))
	fmt.Println("man1",man1)

}

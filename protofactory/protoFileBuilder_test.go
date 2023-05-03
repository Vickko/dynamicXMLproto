package protofactory

import (
	"fmt"
	"testing"

	"github.com/Vickko/dynamicXMLproto/workspace"
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
			NewField("", "String", "firstName", 1),
			NewField("", "String", "lastName", 2),
			NewField("", "Bool", "isAlive", 3),
			NewField("", "Int64", "age", 4),
		).ExportMsg(),
		NewMsgBuilder().Set("Man2",
			NewField("", "String", "firstName", 1),
			NewField("", "String", "lastName", 2),
			NewField("", "Bool", "isAlive", 3),
			NewField("", "Int64", "age", 4),
		).ExportMsg(),
		NewMsgBuilder().Set("TestId",
			NewField("", "Int64", "id", 1),
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
			NewField("", "String", "firstName", 1),
			NewField("", "String", "lastName", 2),
			NewField("", "Bool", "isAlive", 3),
			NewField("", "Int64", "age", 4),
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

func Test_nestedType_original(t *testing.T) {
	name := &descriptorpb.FieldDescriptorProto{
		Name:     proto.String("name"),
		Number:   proto.Int32(1),
		Label:    optional,
		Type:     descriptorpb.FieldDescriptorProto_Type(protoreflect.StringKind).Enum(),
		JsonName: proto.String("name"),
	}
	m1 := descriptorpb.DescriptorProto{
		Name:    proto.String("Man"),
		Options: new(descriptorpb.MessageOptions),
	}
	m1.Field = append(m1.Field, name)

	country := &descriptorpb.FieldDescriptorProto{
		Name:     proto.String("country"),
		Number:   proto.Int32(1),
		Label:    optional,
		Type:     descriptorpb.FieldDescriptorProto_Type(protoreflect.StringKind).Enum(),
		JsonName: proto.String("country"),
	}
	city := &descriptorpb.FieldDescriptorProto{
		Name:     proto.String("city"),
		Number:   proto.Int32(2),
		Label:    optional,
		Type:     descriptorpb.FieldDescriptorProto_Type(protoreflect.StringKind).Enum(),
		JsonName: proto.String("city"),
	}
	a1 := descriptorpb.DescriptorProto{
		Name:    proto.String("Address"),
		Options: new(descriptorpb.MessageOptions),
	}
	a1.Field = append(a1.Field, country, city)

	a1f := &descriptorpb.FieldDescriptorProto{
		Name:     proto.String("address"),
		Number:   proto.Int32(2),
		Type:     descriptorpb.FieldDescriptorProto_Type(protoreflect.MessageKind).Enum(),
		TypeName: proto.String("Address"),
	}
	_ = a1f
	m1.Field = append(m1.Field, a1f)
	m1.NestedType = append(m1.NestedType, &a1)

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

	md2 := md.Messages().ByName(protoreflect.Name("Address"))
	fmt.Printf("md2: %v\n", md2)
	//fmt.Printf("fd.Messages().Len(): %v\n", fd.Messages().Len())
	address := dynamicpb.NewMessage(md2)

	address.Set(md2.Fields().ByJSONName("country"), protoreflect.ValueOfString("USA"))
	address.Set(md2.Fields().ByJSONName("city"), protoreflect.ValueOfString("NewYork"))

	man1.Set(md.Fields().ByJSONName("name"), protoreflect.ValueOfString("John"))
	man1.Set(md.Fields().ByJSONName("address"), protoreflect.ValueOfMessage(address))
	fmt.Println(man1)

	fmt.Printf("man1.ProtoReflect(): %v\n", man1.ProtoReflect())

}
func Test_nestedType_original2(t *testing.T) {
	name := &descriptorpb.FieldDescriptorProto{
		Name:     proto.String("name"),
		Number:   proto.Int32(1),
		Label:    optional,
		Type:     descriptorpb.FieldDescriptorProto_Type(protoreflect.StringKind).Enum(),
		JsonName: proto.String("name"),
	}
	m1 := descriptorpb.DescriptorProto{
		Name:    proto.String("Man"),
		Options: new(descriptorpb.MessageOptions),
	}
	m1.Field = append(m1.Field, name)

	country := &descriptorpb.FieldDescriptorProto{
		Name:     proto.String("country"),
		Number:   proto.Int32(1),
		Label:    optional,
		Type:     descriptorpb.FieldDescriptorProto_Type(protoreflect.StringKind).Enum(),
		JsonName: proto.String("country"),
	}
	city := &descriptorpb.FieldDescriptorProto{
		Name:     proto.String("city"),
		Number:   proto.Int32(2),
		Label:    optional,
		Type:     descriptorpb.FieldDescriptorProto_Type(protoreflect.StringKind).Enum(),
		JsonName: proto.String("city"),
	}
	a1 := descriptorpb.DescriptorProto{
		Name:    proto.String("Address"),
		Options: new(descriptorpb.MessageOptions),
	}
	a1.Field = append(a1.Field, country, city)

	a1f := &descriptorpb.FieldDescriptorProto{
		Name:     proto.String("address"),
		Number:   proto.Int32(2),
		Type:     descriptorpb.FieldDescriptorProto_Type(protoreflect.MessageKind).Enum(),
		TypeName: proto.String("Address"),
	}
	_ = a1f
	m1.Field = append(m1.Field, a1f)
	m1.NestedType = append(m1.NestedType, &a1)

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

	md2 := md.Messages().ByName(protoreflect.Name("Address"))
	fmt.Printf("md2: %v\n", md2)
	//fmt.Printf("fd.Messages().Len(): %v\n", fd.Messages().Len())
	address := dynamicpb.NewMessage(md2)

	address.Set(md2.Fields().ByJSONName("country"), protoreflect.ValueOfString("USA"))
	address.Set(md2.Fields().ByJSONName("city"), protoreflect.ValueOfString("NewYork"))

	man1.Set(md.Fields().ByJSONName("name"), protoreflect.ValueOfString("John"))
	man1.Set(md.Fields().ByJSONName("address"), protoreflect.ValueOfMessage(address))
	fmt.Println("man1:\n", man1)

	adr := man1.Get(md.Fields().ByJSONName("address"))
	fmt.Println("protoreflectMsg: ", adr)
	adrm := adr.Message()
	adrm.Set(md2.Fields().ByJSONName("city"), protoreflect.ValueOfString("NewYorkk"))
	adrm.Set(md2.Fields().ByJSONName("country"), protoreflect.ValueOfString("USAA"))

	fmt.Printf("man1.ProtoReflect():\n %v\n", man1.ProtoReflect())

}

func Test_nestedType_wrapped(t *testing.T) {
	f := NewFileBuilder().SetConfig(
		"ai.momenta/hdmap/data/mdf/test.proto",
		"proto3",
		"hdmap.data.mdf",
		"ai.momenta/hdmap/data/mdf",
	).AppendMsg(
		NewMsgBuilder().Set("Man",
			NewField("", "String", "name", 1),
			NewField("", "Address", "address", 2),
		).RegisterNestedMsgs(
			NewMsgBuilder().Set("Address",
				NewField("", "String", "country", 1),
				NewField("", "String", "city", 2),
			).ExportMsg(),
		).ExportMsg(),
	)
	man1 := f.MsgInstance("Man")
	fmt.Println("f.Build().Messages().Len(): ", f.Build().Messages().Len())
	fmt.Println("f.Build().Messages():")
	for i := 0; i < f.Build().Messages().Len(); i++ {
		fmt.Println("", &ProtoMessage{f.Build().Messages().Get(i)})
	}
	man1.SetField("name", "John")
	fmt.Println("man1:\n\t", man1)
	addr1 := man1.SelectChildMsg("address")
	addr1.SetField("city", "NewYork")
	addr1.SetField("country", "USA")
	fmt.Println("man1:\n\t", man1)
}

func Test_compareTrueMsg(t *testing.T) {
	var ma workspace.Man
	fmt.Println("ma.ProtoReflect().Descriptor().Messages(): \n\t", ma.ProtoReflect().Descriptor().Messages())
	fmt.Println(workspace.File_workspace_test_proto.Messages().Len())
	fmt.Println("workspace.File_workspace_test_proto.Messages(): \n\t", workspace.File_workspace_test_proto.Messages().Get(0))
	fmt.Println()
	fmt.Println()
	var maa workspace.Man_Address
	fmt.Println("maa.ProtoReflect().Descriptor(): \n\t", maa.ProtoReflect().Descriptor())
	fmt.Println("maa(protoreflect.Message): \n\t", maa.ProtoReflect())
	fmt.Println("maa(protoreflect.ProtoMessgae): \n\t", maa.ProtoReflect().Interface())
}

func Test_repeatedType_original(t *testing.T) {
	num := &descriptorpb.FieldDescriptorProto{
		Name:     proto.String("num"),
		Number:   proto.Int32(1),
		Label:    descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
		Type:     descriptorpb.FieldDescriptorProto_Type(protoreflect.Int32Kind).Enum(),
		JsonName: proto.String("num"),
	}
	nums := descriptorpb.DescriptorProto{
		Name:    proto.String("nums"),
		Options: new(descriptorpb.MessageOptions),
	}
	nums.Field = append(nums.Field, num)

	f := descriptorpb.FileDescriptorProto{
		Name:    proto.String("ai.momenta/hdmap/data/mdf/test.proto"),
		Syntax:  proto.String("proto3"),
		Package: proto.String("hdmap.data.mdf"),
		Options: new(descriptorpb.FileOptions),
	}
	f.Options.GoPackage = proto.String("hdmap.data.mdf")
	f.MessageType = append(f.MessageType, &nums)
	fd, _ := protodesc.NewFile(&f, &protoregistry.Files{})
	md := fd.Messages().ByName(protoreflect.Name("nums"))
	nums1 := dynamicpb.NewMessage(md)
	//nums1.Set(md.Fields().ByJSONName("num"), protoreflect.ValueOfList())
	fmt.Println(nums1)

	// lv,_ := structpb.NewList([]any{1, 2, 3})
	// structpb.NewListValue(lv).
	nn := nums1.ProtoReflect().Mutable(nums1.Descriptor().Fields().ByJSONName("num")).List()
	nn.Append(protoreflect.ValueOfInt32(1))
	nn.Append(protoreflect.ValueOfInt32(2))
	nn.Append(protoreflect.ValueOfInt32(3))
	nn.Append(protoreflect.ValueOfInt32(4))
	fmt.Println(nums1)
	fmt.Println(nn)
	fmt.Printf("nums1.ProtoReflect().Descriptor(): %v\n", nums1.ProtoReflect().Descriptor().Fields().ByJSONName("num").Kind())

	nnn := nums1.ProtoReflect().Get(nums1.Descriptor().Fields().ByJSONName("num")).List()
	fmt.Printf("nnn.List(): %v\n", nnn)
}

func Test_repeatedType_wrapped(t *testing.T) {
	f := NewFileBuilder().SetConfig(
		"ai.momenta/hdmap/data/mdf/test.proto",
		"proto3",
		"hdmap.data.mdf",
		"ai.momenta/hdmap/data/mdf",
	).AppendMsg(
		NewMsgBuilder().Set("nums",
			NewField("repeated", "Int32", "num", 1),
		).ExportMsg(),
	)
	nums := f.MsgInstance("nums")
	num := nums.SelectChildListMutable("num")
	num.ListAppend(1)
	num.ListAppend(2)
	num.ListAppend(3)
	fmt.Println(nums)
	num.ListSet(0, 3)
	fmt.Println(nums)

}

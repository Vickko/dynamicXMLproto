package protofactory

import (
	"math"
	"strconv"

	log "github.com/Vickko/dynamicXMLproto/logplus"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

type ProtoInstance struct {
	dynamicpb.Message
}

// TODO: ParseValue function still fragile,
// need to take a deeper research into kind-value relation
func ParseValue(k protoreflect.Kind, v any) protoreflect.Value {
	switch k {
	case protoreflect.BoolKind:
		return protoreflect.ValueOfBool(v.(bool))
	case protoreflect.EnumKind:
		return protoreflect.ValueOfEnum(v.(protoreflect.EnumNumber))
	case protoreflect.Int32Kind:
		return protoreflect.ValueOfInt32(int32(v.(int)))
	case protoreflect.Sint32Kind:
		return protoreflect.ValueOfInt32(int32(v.(int)))
	case protoreflect.Uint32Kind:
		return protoreflect.ValueOfUint32(uint32(v.(int)))
	case protoreflect.Int64Kind:
		return protoreflect.ValueOfInt64(v.(int64))
	case protoreflect.Sint64Kind:
		return protoreflect.ValueOfInt64(int64(v.(int)))
	case protoreflect.Uint64Kind:
		return protoreflect.ValueOfUint64(uint64(v.(int)))
	case protoreflect.Sfixed32Kind:
		return protoreflect.ValueOfInt32(int32(v.(int)))
	case protoreflect.Fixed32Kind:
		return protoreflect.ValueOfInt32(int32(v.(int)))
	case protoreflect.FloatKind:
		return protoreflect.ValueOfFloat32(v.(float32))
	case protoreflect.Sfixed64Kind:
		return protoreflect.ValueOfInt64(int64(v.(int)))
	case protoreflect.Fixed64Kind:
		return protoreflect.ValueOfInt64(int64(v.(int)))
	case protoreflect.DoubleKind:
		return protoreflect.ValueOfFloat64(v.(float64))
	case protoreflect.StringKind:
		return protoreflect.ValueOfString(v.(string))
	case protoreflect.BytesKind:
		return protoreflect.ValueOfBytes(v.([]byte))
	case protoreflect.MessageKind:
		return protoreflect.ValueOfMessage(v.(protoreflect.Message))
	// TODO: what is GroupKind?
	// case protoreflect.GroupKind:
	// 	return protoreflect.ValueOf
	default:
		log.Errorln("Try parse value into invalid protoreflect.Kind")

	}
	return protoreflect.Value{}
}

// TODO: 这两个 setField 在ParseValue时的类型断言逻辑有问题，需要解决
// 老 setField 的参数是期望来自于字面量，而字面量的数字会默认为 int
// 新 setField 的参数来自于 parseInt， 会被 parse 为确切的 int32 或 int64
func (p *ProtoInstance) SetField(name string, value any) *ProtoInstance {
	fd := p.Descriptor().Fields().ByJSONName(name)
	// fmt.Println(ParseValue(fd.Kind(), value).String())
	// fmt.Println(p)
	p.Set(fd, ParseValue(fd.Kind(), value))
	return p
}

// func AssertTypeFromStrValue(str string) (res string) {
// 	if n1, err := strconv.ParseInt(str, 10, 64); err == nil {
// 		if n1 < math.MaxInt32 {
// 			res = "int"
// 		} else {
// 			res = "int64"
// 		}
// 	} else if _, err := strconv.ParseFloat(str, 64); err == nil {
// 		res = "float"
// 	} else if _, err := strconv.ParseBool(str); err == nil {
// 		res = "bool"
// 	}
// 	res = "string"

// 	return res
// }

func JudgeValueType(str string) string {
	n, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		if n < math.MaxInt32 {
			return "Int32"
		} else {
			return "Int64"
		}
	}
	_, err = strconv.ParseFloat(str, 64)
	if err == nil {
		return "Double"
	}
	_, err = strconv.ParseBool(str)
	if err == nil {
		return "Bool"
	}
	return "String"
}

func (p *ProtoInstance) SetFieldByType(name string, str string) *ProtoInstance {
	var value any

	if n1, err := strconv.ParseInt(str, 10, 64); err == nil {
		if n1 < math.MaxInt32 {
			value = int(n1)
		} else {
			value = int64(n1)
		}
	} else if n2, err := strconv.ParseFloat(str, 64); err == nil {
		value = n2
	} else if n3, err := strconv.ParseBool(str); err == nil {
		value = n3
	} else {
		value = str
	}

	return p.SetField(name, value)

}

// nested msg is a nil pointer by default when parent msg instance is created.
// so if you want to get or set nested msg, you need to create it first
// by create I mean childMsg := dynamicpb.NewMessage(childMsgDesc)
// TODO: consider create nested msg instance should be exec when parentMsg.MsgInstance(),
// or when SelectChildMsg().
func (p *ProtoInstance) SelectChildMsg(name string) *ProtoInstance {
	parentMsgDesc := p.Descriptor()
	childMsgDesc := parentMsgDesc.Messages().ByName(protoreflect.Name("Address"))
	childMsg := dynamicpb.NewMessage(childMsgDesc)
	// childMsg.Set(childMsgDesc.Fields().ByJSONName("country"), protoreflect.ValueOfString("USA"))
	// childMsg.Set(childMsgDesc.Fields().ByJSONName("city"), protoreflect.ValueOfString("NewYork"))
	p.Set(parentMsgDesc.Fields().ByJSONName("address"), protoreflect.ValueOfMessage(childMsg))
	// fmt.Println("childMsg:\n\t", childMsg)
	// fmt.Println("parentMsg:\n\t", p)
	return &ProtoInstance{*childMsg}
}

// if you select child list for set propose, mutable is required
func (p *ProtoInstance) SelectChildListMutable(name string) ProtoListField {
	return ProtoListField{
		p.ProtoReflect().Mutable(p.Descriptor().Fields().ByJSONName(name)).List(),
		p.Descriptor().Fields().ByJSONName(name).Kind(),
	}
}

func (p *ProtoInstance) SelectChildList(name string) ProtoListField {
	return ProtoListField{
		p.ProtoReflect().Get(p.Descriptor().Fields().ByJSONName(name)).List(),
		p.Descriptor().Fields().ByJSONName(name).Kind(),
	}
}

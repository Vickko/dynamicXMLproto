package protofactory

import (
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
		return protoreflect.ValueOfInt64(int64(v.(int)))
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

func (p *ProtoInstance) SetField(name string, value any) *ProtoInstance {
	fd := p.Descriptor().Fields().ByJSONName(name)
	p.Set(fd, ParseValue(fd.Kind(), value))
	return p
}

package benchmark

import (
	"encoding/xml"
	"math/rand"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"
)

type student struct {
	Name string
	Id   int
}

var num int = 100000
var students = make([]student, num)
var studentsProto = make([]StudentProto, num)

func init() {
	// 随机数种子
	rand.Seed(time.Now().UnixNano())

	var name = rand.Intn(6) + 3 // 生成字符串的长度在 3 到 8 之间

	randStr := func(len int) string {
		str := make([]byte, len)

		for j := 0; j < len; j++ {
			str[j] = byte(rand.Intn('z'-'a'+1) + 'a')
		}
		return string(str)
	}

	for i := 0; i < num; i++ {
		students[i].Name = randStr(name)
		students[i].Id = rand.Int()%89999999 + 10000000 // 1000,0000 ~ 9999,9999

		studentsProto[i].Name = students[i].Name
		studentsProto[i].Id = int32(students[i].Id)
	}

}

func Benchmark_tinyEncXML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := range students {
			xml.Marshal(students[j])
		}
	}
}

func Benchmark_tinyDecXML(b *testing.B) {
	data := make([][]byte, num)
	for i := range students {
		data[i], _ = xml.Marshal(students[i])
	}
	t := student{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := range data {
			xml.Unmarshal(data[j], &t)
		}
	}
}

func Benchmark_tinyEncProto(b *testing.B) {
	// b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := range studentsProto {
			proto.Marshal(&studentsProto[j])
		}
	}
}

func Benchmark_tinyDecProto(b *testing.B) {
	data := make([][]byte, num)
	for i := range students {
		data[i], _ = proto.Marshal(&studentsProto[i])
	}
	t := StudentProto{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := range studentsProto {
			proto.Unmarshal(data[j], &t)
		}
	}
}

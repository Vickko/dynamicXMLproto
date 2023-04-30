// go test -benchmem -run=^$ -bench ^Benchmark_common1Dec.\* github.com/Vickko/dynamicXMLproto/benchmark -count=16

// <book>
//   <title>The Great Gatsby</title>
//   <author>
//     <first-name>F. Scott</first-name>
//     <last-name>Fitzgerald</last-name>
//   </author>
//   <publisher>
//     <name>Scribner</name>
//     <year>1925</year>
//   </publisher>
// </book>

package benchmark

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"
)

type book struct {
	Title  string
	Author struct {
		FirstName string
		LastName  string
	}
	Publisher struct {
		Name string
		Year int
	}
}

var num int = 100000
var books = make([]book, num)
var booksProto = make([]BookProto, num)

func init() {
	// 随机数种子
	rand.Seed(time.Now().UnixNano())

	var title = rand.Intn(46) + 3 // 生成字符串的长度在 1x3 到 6x8 之间
	var name = rand.Intn(6) + 3   // 生成字符串的长度在 3 到 8 之间

	randStr := func(len int) string {
		str := make([]byte, len)

		for j := 0; j < len; j++ {
			str[j] = byte(rand.Intn('z'-'a'+1) + 'a')
		}
		return string(str)
	}

	for i := 0; i < num; i++ {
		//fmt.Println(i)
		books[i].Title = randStr(title)
		books[i].Author.FirstName = randStr(name)
		books[i].Author.LastName = randStr(name)
		books[i].Publisher.Name = randStr(name)
		books[i].Publisher.Year = rand.Int()%2000 + 1000 // 1000~2999

		booksProto[i].Title = books[i].Title
		booksProto[i].Author = &BookProtoAuthor{}
		booksProto[i].Author.FirstName = books[i].Author.FirstName
		booksProto[i].Author.LastName = books[i].Author.LastName
		booksProto[i].Publisher = &BookProtoPublisher{}
		booksProto[i].Publisher.Name = books[i].Publisher.Name
		booksProto[i].Publisher.Year = int32(books[i].Publisher.Year)
	}
}

func Test_common1XML(t *testing.T) {
	fmt.Println(books[1])
	fmt.Println(time.Now())
	x, _ := xml.Marshal(books[1])
	fmt.Println(books[1])
	fmt.Println(string(x))
	b := book{}
	fmt.Println(b)
	xml.Unmarshal(x, &b)
	fmt.Println(b)
}

func Test_common1proto(t *testing.T) {
	fmt.Println(booksProto[0].Title)
	//proto.Marshal(&booksProto[1])
}

func Benchmark_common1EncXML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := range books {
			xml.Marshal(books[j])
		}
	}
}
func Benchmark_common1DecXML(b *testing.B) {
	data := make([][]byte, num)
	for i := range books {
		data[i], _ = xml.Marshal(books[i])
	}
	t := book{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := range data {
			xml.Unmarshal(data[j], &t)
		}
	}
}

func Benchmark_common1EncProto(b *testing.B) {
	// b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := range booksProto {
			proto.Marshal(&booksProto[j])
		}
	}
}

func Benchmark_common1DecProto(b *testing.B) {
	data := make([][]byte, num)
	for i := range books {
		data[i], _ = proto.Marshal(&booksProto[i])
	}
	t := BookProto{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := range booksProto {
			proto.Unmarshal(data[j], &t)
		}
	}
}

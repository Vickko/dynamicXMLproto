package benchmark

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"
)

type date struct {
	Year  int
	Month int
	Day   int
}

type job struct {
	Title     string
	Company   string
	StartDate date
	EndDate   date
}

type resume struct {
	Name    string
	Email   string
	Phone   int64
	Address struct {
		Street string
		City   string
		State  string
		Zip    int
	}
	PersonalInfo struct {
		Education struct {
			Degree   string
			Major    string
			School   string
			Graduate date
		}
		Experience struct {
			Job []job
		}
	}
}

var num int = 100000
var resumes = make([]resume, num)
var resumesProto = make([]ResumeProto, num)

func init() {
	// 随机数种子
	rand.Seed(time.Now().UnixNano())

	var word = rand.Intn(6) + 3 // 生成字符串的长度在 3 到 8 之间

	randStr := func(len int) string {
		str := make([]byte, len)

		for j := 0; j < len; j++ {
			str[j] = byte(rand.Intn('z'-'a'+1) + 'a')
		}
		return string(str)
	}

	for i := 0; i < num; i++ {
		resumes[i].Name = randStr(word)
		resumes[i].Email = randStr(2*word) + "@" + randStr(2*word) + "." + randStr(word)
		resumes[i].Phone = int64(rand.Int()%90000000000 + 10000000000) // 11 digits
		resumes[i].Address.Street = randStr(4 * word)
		resumes[i].Address.City = randStr(word)
		resumes[i].Address.State = randStr(word)
		resumes[i].Address.Zip = rand.Int()%90000 + 10000 // 5 digits
		resumes[i].PersonalInfo.Education.Degree = randStr(3 * word)
		resumes[i].PersonalInfo.Education.Major = randStr(2 * word)
		resumes[i].PersonalInfo.Education.School = randStr(5 * word)
		resumes[i].PersonalInfo.Education.Graduate.Year = rand.Int()%2000 + 1000
		resumes[i].PersonalInfo.Education.Graduate.Month = rand.Int()%12 + 1
		resumes[i].PersonalInfo.Education.Graduate.Day = rand.Int()%31 + 1
		jobCount := rand.Int()%4 + 1
		for j := 0; j < jobCount; j++ {
			resumes[i].PersonalInfo.Experience.Job = append(resumes[i].PersonalInfo.Experience.Job,
				job{Title: randStr(3 * word),
					Company: randStr(word),
					StartDate: date{
						Year:  rand.Int()%2000 + 1000,
						Month: rand.Int()%12 + 1,
						Day:   rand.Int()%31 + 1},
					EndDate: date{
						Year:  rand.Int()%2000 + 1000,
						Month: rand.Int()%12 + 1,
						Day:   rand.Int()%31 + 1}})
		}

		resumesProto[i].Name = resumes[i].Name
		resumesProto[i].Email = resumes[i].Email
		resumesProto[i].Phone = resumes[i].Phone
		resumesProto[i].Address = &ResumeProtoAddress{}
		resumesProto[i].Address.Street = resumes[i].Address.Street
		resumesProto[i].Address.City = resumes[i].Address.City
		resumesProto[i].Address.State = resumes[i].Address.State
		resumesProto[i].Address.Zip = int32(resumes[i].Address.Zip)
		resumesProto[i].PersonalInfo = &ResumeProtoPersonalInfo{}
		resumesProto[i].PersonalInfo.Education = &ResumeProtoPersonalInfoEducation{}
		resumesProto[i].PersonalInfo.Experience = &ResumeProtoPersonalInfoExperience{}
		resumesProto[i].PersonalInfo.Education.Degree = resumes[i].PersonalInfo.Education.Degree
		resumesProto[i].PersonalInfo.Education.Major = resumes[i].PersonalInfo.Education.Major
		resumesProto[i].PersonalInfo.Education.School = resumes[i].PersonalInfo.Education.School
		resumesProto[i].PersonalInfo.Education.Graduate = &ResumeProtoPersonalInfoDate{}
		resumesProto[i].PersonalInfo.Education.Graduate.Year = int32(resumes[i].PersonalInfo.Education.Graduate.Year)
		resumesProto[i].PersonalInfo.Education.Graduate.Month = int32(resumes[i].PersonalInfo.Education.Graduate.Month)
		resumesProto[i].PersonalInfo.Education.Graduate.Day = int32(resumes[i].PersonalInfo.Education.Graduate.Day)
		for j := 0; j < jobCount; j++ {
			resumesProto[i].PersonalInfo.Experience.Job = append(resumesProto[i].PersonalInfo.Experience.Job,
				&ResumeProtoPersonalInfoExperienceJob{
					Title:   resumes[i].PersonalInfo.Experience.Job[j].Title,
					Company: resumes[i].PersonalInfo.Experience.Job[j].Company,
					StartDate: &ResumeProtoPersonalInfoDate{
						Year:  int32(resumes[i].PersonalInfo.Experience.Job[j].StartDate.Year),
						Month: int32(resumes[i].PersonalInfo.Experience.Job[j].StartDate.Month),
						Day:   int32(resumes[i].PersonalInfo.Experience.Job[j].StartDate.Day),
					},
					EndDate: &ResumeProtoPersonalInfoDate{
						Year:  int32(resumes[i].PersonalInfo.Experience.Job[j].StartDate.Year),
						Month: int32(resumes[i].PersonalInfo.Experience.Job[j].StartDate.Month),
						Day:   int32(resumes[i].PersonalInfo.Experience.Job[j].StartDate.Day),
					},
				})
		}

	}

	_ = resumes
	_ = resumesProto
}

func Test_common2All(t *testing.T) {
	fmt.Println("hello from test")
	fmt.Println(resumes[1])
	fmt.Println(resumesProto[1])

}

func Benchmark_common2EncXML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := range resumes {
			xml.Marshal(resumes[j])
		}
	}
}

func Benchmark_common2DecXML(b *testing.B) {
	data := make([][]byte, num)
	for i := range resumes {
		data[i], _ = xml.Marshal(resumes[i])
	}
	t := resume{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := range data {
			xml.Unmarshal(data[j], &t)
		}
	}
}

func Benchmark_common2EncProto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := range resumesProto {
			proto.Marshal(&resumesProto[j])
		}
	}
}

func Benchmark_common2DecProto(b *testing.B) {
	data := make([][]byte, num)
	for i := range resumesProto {
		data[i], _ = proto.Marshal(&resumesProto[i])
	}
	t := ResumeProto{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := range resumesProto {
			proto.Unmarshal(data[j], &t)
		}
	}
}
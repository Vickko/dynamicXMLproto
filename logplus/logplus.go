// Copyright 2023 Vickko. All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package logplus

import "log"

func Fatal(v ...any) {
	log.SetPrefix("[Fatal] ")
	log.Fatal(v...)
	log.SetPrefix("")
}

func Fatalf(format string, v ...any) {
	log.SetPrefix("[Fatal] ")
	log.Fatalf(format, v...)
	log.SetPrefix("")
}

func Fatalln(v ...any) {
	log.SetPrefix("[Fatal] ")
	log.Fatalln(v...)
	log.SetPrefix("")
}

func Error(v ...any) {
	log.SetPrefix("[Error] ")
	log.Panic(v...)
	log.SetPrefix("")
}

func Errorf(format string, v ...any) {
	log.SetPrefix("[Error] ")
	log.Panicf(format, v...)
	log.SetPrefix("")
}

func Errorln(v ...any) {
	log.SetPrefix("[Error] ")
	log.Panicln(v...)
	log.SetPrefix("")
}

func Warning(v ...any) {
	log.SetPrefix("[Warning] ")
	log.Print(v...)
	log.SetPrefix("")
}

func Warningf(format string, v ...any) {
	log.SetPrefix("[Warning] ")
	log.Printf(format, v...)
	log.SetPrefix("")
}

func Warningln(v ...any) {
	log.SetPrefix("[Log] ")
	log.Println(v...)
	log.SetPrefix("")
}

func Log(v ...any) {
	log.SetPrefix("[Log] ")
	log.Print(v...)
	log.SetPrefix("")
}

func Logf(format string, v ...any) {
	log.SetPrefix("[Log] ")
	log.Printf(format, v...)
	log.SetPrefix("")
}

func Logln(v ...any) {
	log.SetPrefix("[Log] ")
	log.Println(v...)
	log.SetPrefix("")
}

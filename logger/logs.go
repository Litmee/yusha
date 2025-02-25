// Copyright 2022 The Yusha Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"log"
	"os"
)

var (
	logChan   = make(chan *yuShaLog, 300)
	infoFile  *os.File
	warnFile  *os.File
	errorFile *os.File
	infoLog   *log.Logger
	warnLog   *log.Logger
	errorLog  *log.Logger
)

/**
后续日志模块的功能在这包下实现
具体实现还要规划一下
*/

// 日志结构体模型
type yuShaLog struct {
	t int
	v string
}

// 日志服务总线
func logServer() {
	for {
		l := <-logChan
		switch l.t {
		case INFO_:
			infoLog.Println(l.v)
		case WARN_:
			warnLog.Println(l.v)
		case ERROR_:
			errorLog.Println(l.v)
		default:
			infoLog.Println(l.v)
		}
	}
}

func init() {
	go logServer()
	_, err := os.Stat("./log")
	if err != nil {
		os.Mkdir("log", 0777)
	}
	infoFile, _ = os.OpenFile("./log/yusha-info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	warnFile, _ = os.OpenFile("./log/yusha-warn.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	errorFile, _ = os.OpenFile("./log/yusha-error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	infoLog = log.New(infoFile, "[INFO] ", 3)
	warnLog = log.New(warnFile, "[WARN] ", 3)
	errorLog = log.New(errorFile, "[ERROR] ", 3)
}

func INFO(val string) {
	logChan <- &yuShaLog{t: INFO_, v: val}
}

func WARN(val string) {
	logChan <- &yuShaLog{t: WARN_, v: val}
}

func ERROR(val string) {
	logChan <- &yuShaLog{t: ERROR_, v: val}
}

func CheckLogChan() {
	for {
		if len(logChan) == 0 {
			break
		}
	}
}

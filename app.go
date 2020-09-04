/**
 * Copyright 2014 @ to2.net.
 * name : app1.go
 * author : jarryliu
 * date : 2015-04-27 20:43:
 * description :
 * history :
 */
package gof

import (
	"github.com/ixre/gof/db"
	"github.com/ixre/gof/log"
	"github.com/ixre/gof/shell"
	"github.com/ixre/gof/storage"
	"time"
)

// 应用当前的上下文
var CurrentApp App

type App interface {
	// Provided db access
	Db() db.Connector
	// Return application configs.
	Config() *Config
	// return registry
	Registry() *Registry
	// Storage
	Storage() storage.Interface
	// Return a logger
	Log() log.ILogger
	// Application is running debug mode
	Debug() bool
}

// 自动安装包
func AutoInstall(d time.Duration) {
	execInstall()
	if d == 0 {
		d = time.Second * 15
	}
	t := time.NewTimer(d)
	for {
		select {
		case <-t.C:
			if err := execInstall(); err == nil {
				t.Reset(d)
			} else {
				break
			}
		}
	}
}

func execInstall() error {
	_, _, err := shell.Run("go install .")
	if err != nil {
		log.Println("[ Gof][ Install]:", err)
	}
	return err
}


// 数组参数
type ArrayFlags []string

// Value ...
func (i *ArrayFlags) String() string {
	return fmt.Sprint(*i)
}

// Set 方法是flag.Value接口, 设置flag Value的方法.
// 通过多个flag指定的值， 所以我们追加到最终的数组上.
func (i *ArrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}


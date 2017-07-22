/**
 * Copyright 2014 @ z3q.net.
 * name : app1.go
 * author : jarryliu
 * date : 2015-04-27 20:43:
 * description :
 * history :
 */
package gof

import (
	"github.com/jsix/gof/db"
	"github.com/jsix/gof/log"
	"github.com/jsix/gof/shell"
	"github.com/jsix/gof/storage"
	"time"
)

// 应用当前的上下文
var CurrentApp App

type App interface {
	// Provided db access
	Db() db.Connector
	// Return application configs.
	Config() *Config
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

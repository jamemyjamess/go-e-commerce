package config

import (
	"fmt"
	"time"
)

type IAppConfig interface {
	URL() string // host:port
	Host() string
	Port() int
	Name() string
	Version() string
	ReadTimeOut() time.Duration
	WriteTimeOut() time.Duration
	BodyLimit() int
	FileLitmit() int
	GCPBucket() string
}

type app struct {
	host         string
	port         int
	name         string
	version      string
	readTimeout  time.Duration
	writeTimeout time.Duration
	bodyLimit    int // bytes
	fileLimit    int // bytes
	gcpBucket    string
}

func (a *app) URL() string {
	return fmt.Sprintf("%s:%d", a.host, a.port)
}

func (a *app) Host() string { return a.host }

func (a *app) Port() int { return a.port }

func (a *app) Name() string { return a.name }

func (a *app) Version() string { return a.version }

func (a *app) ReadTimeOut() time.Duration { return a.readTimeout }

func (a *app) WriteTimeOut() time.Duration { return a.writeTimeout }

func (a *app) BodyLimit() int { return a.bodyLimit }

func (a *app) FileLitmit() int { return a.fileLimit }

func (a *app) GCPBucket() string { return a.gcpBucket }

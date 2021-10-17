package mongondb

import "time"

type Option func(opt *Options)

type Options struct {
	URI       string
	Username  string
	Password  string
	Mechanism string
	Database  string
	TimeOut   time.Duration
}

func WithUrl(uri string) Option {
	return func(opt *Options) {
		opt.URI = uri
	}
}

func WithUsernameAndPassword(name, password string) Option {
	return func(opt *Options) {
		opt.Username = name
		opt.Password = password
	}
}

func WithTimeOut(timeout time.Duration) Option {
	return func(opt *Options) {
		opt.TimeOut = timeout
	}
}

func WithDatabase(database string) Option {
	return func(opt *Options) {
		opt.Database = database
	}
}

func WithMechanism(mechanism string) Option {
	return func(opt *Options) {
		opt.Mechanism = mechanism
	}
}

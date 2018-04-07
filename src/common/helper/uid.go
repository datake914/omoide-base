package helper

import "github.com/rs/xid"

func UniqueId() string {
	return xid.New().String()
}

package models

import (
	"github.com/libvirt/libvirt-go"
)

var libvirtConn *libvirt.Connect;

func InitLibvirt(uri string) (*libvirt.Connect, error) {
	conn, err := libvirt.NewConnect(uri)
	if err != nil {
		return nil, err
	}

	libvirtConn = conn

	return conn, nil
}
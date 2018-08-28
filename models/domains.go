package models

import (
	"github.com/libvirt/libvirt-go"
	"log"
)

type Domain struct {
	Name string
}

func (d Domain) GetAll() ([]Domain, error) {
	libvirtDoms, err := libvirtConn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	if err != nil {
		log.Fatalf("failed to get domain list: %v", err)
		return nil, err
	}

	domains := make([]Domain, len(libvirtDoms))

	for i, domain := range libvirtDoms {
		name, err := domain.GetName()

		if err != nil {
			log.Fatalf("failed to get domain name: %v", err)
			return nil, err
		}

		domains[i] = Domain {
			Name: name,
		}
	}

	return domains, nil
}
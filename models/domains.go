package models

import (
	"github.com/libvirt/libvirt-go"
	"log"
	"math"
)

type Domain struct {
	Autostart bool	`json:"autostart"`
	OSType string	`json:"osType"`
	ID *uint		`json:"id"`
	Name string		`json:"name"`
	UUID string		`json:"uuid"`
	State string	`json:"state"`
}

func getDomainStateString(state libvirt.DomainState) string {
	switch state {
	case libvirt.DOMAIN_SHUTOFF:
		return "shut off";
	case libvirt.DOMAIN_RUNNING:
		return "running";
	case libvirt.DOMAIN_PAUSED:
		return "paused";
	case libvirt.DOMAIN_BLOCKED:
		return "blocked";
	case libvirt.DOMAIN_CRASHED:
		return "crashed";
	case libvirt.DOMAIN_PMSUSPENDED:
		return "suspended";
	case libvirt.DOMAIN_NOSTATE:
		fallthrough
	default:
		return "unknown";
	}
}

func (d Domain) GetAll() ([]Domain, error) {
	libvirtDoms, err := libvirtConn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE)
	if err != nil {
		log.Fatalf("failed to get domain list: %v", err)
		return nil, err
	}

	domains := make([]Domain, len(libvirtDoms))

	for i, domain := range libvirtDoms {
		name, err := domain.GetName()

		if err != nil {
			return nil, err
		}

		autostart, err := domain.GetAutostart()

		if err != nil {
			return nil, err
		}

		osType, err := domain.GetOSType()

		if err != nil {
			return nil, err
		}

		id, err := domain.GetID()

		if err != nil {
			return nil, err
		}

		uuid, err := domain.GetUUIDString()

		if err != nil {
			return nil, err
		}

		state, _, err := domain.GetState()

		domains[i] = Domain {
			Autostart: autostart,
			OSType: osType,
			Name: name,
			UUID: uuid,
			State: getDomainStateString(state),
		}

		if id != math.MaxUint32 {
			domains[i].ID = &id
		}
	}

	return domains, nil
}
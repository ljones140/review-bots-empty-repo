package logagg

import "net"

type Page struct {
	Name        string
	Visits      int
	UniqueViews int
	VistorIPs   []net.IP
}

func (p *Page) AddVisit(visitIP net.IP) {
	p.Visits++
	if p.visitAddressUnique(visitIP) {
		p.UniqueViews++
		p.VistorIPs = append(p.VistorIPs, visitIP)
	}
}

func (p *Page) visitAddressUnique(visitIP net.IP) bool {
	for _, ip := range p.VistorIPs {
		if ip.Equal(visitIP) {
			return false
		}
	}
	return true
}

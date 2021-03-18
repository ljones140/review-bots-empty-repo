package logagg

import (
	"net"
)

type Aggregator struct {
	Pages []Page
}

func (a *Aggregator) AddPageView(name string, ip net.IP) {
	for i := range a.Pages {
		if a.Pages[i].Name == name {
			a.Pages[i].AddVisit(ip)
			return
		}
	}

	newPage := Page{Name: name}
	newPage.AddVisit(ip)
	a.Pages = append(a.Pages, newPage)
}

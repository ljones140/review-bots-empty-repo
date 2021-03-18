package logagg_test

import (
	"net"
	"testing"

	logagg "github.com/ljones140/go_view_log_aggregator"
)

func TestPageCreation(t *testing.T) {
	t.Run("it is created with a page name", func(t *testing.T) {
		page := logagg.Page{Name: "index"}
		if page.Name != "index" {
			t.Errorf("got page name %q wanted %q", page.Name, "index")
		}
	})
}

func TestAddingPageVisits(t *testing.T) {
	t.Run("visit count is incremented when visit is added", func(t *testing.T) {
		page := logagg.Page{Name: "index"}
		address := net.ParseIP("192.0.2.1")
		page.AddVisit(address)

		assertPageVisits(t, page.Visits, 1)
	})

	t.Run("unique view count is incremented when visits are from different ips", func(t *testing.T) {
		page := logagg.Page{Name: "index"}
		page.AddVisit(net.ParseIP("192.0.2.1"))
		page.AddVisit(net.ParseIP("192.1.1.1"))

		assertPageVisits(t, page.Visits, 2)
		assertUniqueViews(t, page.UniqueViews, 2)
	})

	t.Run("unique view count is not incremented when visits from same ip", func(t *testing.T) {
		page := logagg.Page{Name: "index"}
		page.AddVisit(net.ParseIP("192.0.2.1"))
		page.AddVisit(net.ParseIP("192.0.2.1"))

		assertPageVisits(t, page.Visits, 2)
		assertUniqueViews(t, page.UniqueViews, 1)
	})
}

func assertPageVisits(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("number of visits got %d wanted %d", got, want)
	}
}

func assertUniqueViews(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("number of unique views got %d wanted %d", got, want)
	}
}

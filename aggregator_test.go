package logagg_test

import (
	"net"
	"testing"

	logagg "github.com/ljones140/go_view_log_aggregator"
)

func TestAddingPages(t *testing.T) {
	t.Run("it builds a page view for a page", func(t *testing.T) {
		agg := &logagg.Aggregator{}
		agg.AddPageView("page-name", net.ParseIP("192.0.1.1"))

		want := []logagg.Page{
			{"page-name", 1, 1, []net.IP{net.ParseIP("192.0.1.1")}},
		}

		logagg.AssertPageViews(t, agg.Pages, want)
	})

	t.Run("it does not add the same page twice", func(t *testing.T) {
		agg := logagg.Aggregator{}
		agg.AddPageView("page-name", net.ParseIP("192.0.1.1"))
		agg.AddPageView("page-name", net.ParseIP("192.0.1.1"))

		want := []logagg.Page{
			{"page-name", 2, 1, []net.IP{net.ParseIP("192.0.1.1")}},
		}

		logagg.AssertPageViews(t, agg.Pages, want)
	})
}

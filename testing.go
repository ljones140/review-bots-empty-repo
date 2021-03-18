package logagg

import (
	"reflect"
	"testing"
)

func AssertPageViews(t testing.TB, got, want []Page) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

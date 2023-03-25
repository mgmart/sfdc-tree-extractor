package sfdcTreeExtractor

import "testing"

func TestFakerCompany(t *testing.T) {

	got := fakeCompany()
	want := ""
	t.Errorf("got %v want %v", got, want)
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

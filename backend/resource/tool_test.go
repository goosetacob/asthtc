package tool

import "testing"

func TestAesthetic(t *testing.T) {
	response, err := Aesthetic("los angeles")
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if response != "~ l o s a n g e l e s ~" {
		t.Errorf("Aesthetic was incorrect, got: %v, want: %v.", response, "~ l o s a n g e l e s ~")
	}
}

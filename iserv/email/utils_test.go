package email

import (
	"fmt"
	"testing"
)

func TestGetToFieldString(t *testing.T) {
	input := []struct {
		addr     string
		dispName string
		ccs      []string
	}{
		{"user@example.com", "User", []string{}},
		{"alex.lanz@hotmail.com", "", []string{}},
		{"peterppp@gmail.com", "Peter P", []string{"joejjj@gmail.com"}},
		{"peterppp@gmail.com", "Peter P", []string{"joejjj@gmail.com", "john.doe@example.com"}},
	}

	want := []string{
		"To: User <user@example.com>",
		"To: alex.lanz@hotmail.com",
		"To: Peter P <peterppp@gmail.com>\r\nCc: joejjj@gmail.com",
		"To: Peter P <peterppp@gmail.com>\r\nCc: joejjj@gmail.com, john.doe@example.com",
	}

	for i, e := range input {
		t.Run(fmt.Sprintf("getToFieldString=%d", i), func(t *testing.T) {
			got := getToFieldString(e.addr, e.dispName, e.ccs)
			if got != want[i] {
				t.Fatalf("got '%v', want '%v'", got, want[i])
			}
		})
	}
}

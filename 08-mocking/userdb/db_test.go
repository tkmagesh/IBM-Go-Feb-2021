package userdb

import "testing"

func TestUserExist(t *testing.T) {
	found := UserExists("magesh@dummy.com")
	if found != true {
		t.Errorf("expected to return true for existing user")
	}
}

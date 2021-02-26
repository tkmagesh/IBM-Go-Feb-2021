package simpleservice

import "testing"

func TestCheckUserExist(t *testing.T) {
	user := User{
		Name:     "Magesh",
		Email:    "magesh@example.com",
		UserName: "magesh",
	}

	err := RegisterUser(user)
	if err == nil {
		t.Error("Expected Register User to throw and error got nil")
	}
}

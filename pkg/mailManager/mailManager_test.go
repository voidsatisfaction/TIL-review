package mailManager

import (
	"testing"
)

func TestMailManagerSend(t *testing.T) {

	mm := New("test@gmail.com", "password", "smtp.gmail.com", "587")
	err := mm.SendHTML("test@gmail.com", []string{"test@gmail.com"}, "subject", "hello world!")

	if err == nil {
		t.Fatalf("expected err not nil, but got nil")
	}
}

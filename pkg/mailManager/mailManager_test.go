package mailManager

import (
	"testing"
)

func TestMailManagerSend(t *testing.T) {

	mm := New("test@gmail.com", "password", "smtp.gmail.com", "587")
	err := mm.Send("test@gmail.com", []string{"test@gmail.com"}, "hello world!")

	if err == nil {
		t.Fatalf("expected err not nil, but got nil")
	}
}

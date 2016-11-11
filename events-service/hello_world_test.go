package events-service

import "testing"

func TestHelloWorld(t *testing.T) {
	uut := events-service{}

	cases := []struct {
		in, want string
	}{
		{"", "Hello, World"},
	}
	for _, c := range cases {
		got, _ := uut.HelloWorld()
		if got != c.want {
			t.Errorf("HelloWorld() == %q, want %q", got, c.want)
		}
	}
}

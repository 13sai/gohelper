package http

import "testing"

func TestIP2long(t *testing.T) {
	t.Log("127.0.0.1", IP2long("127.0.0.1"))
	t.Log("114.114.114.114", IP2long("114.114.114.114"))
}

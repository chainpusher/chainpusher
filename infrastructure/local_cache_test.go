package infrastructure

import "testing"

func TestCache(t *testing.T) {
	_ = SetKey("test", []byte("test"))
	bytes, err := GetKey("test")
	if err != nil {
		t.Error(err)
	}
	if bytes.String() != "test" {
		t.Error("Expected 'test', got", string(bytes))
	}
}

package pokecache

import "testing"

func TestAddGet(t *testing.T) {
	cases := []struct {
		key string
		val []byte
	}{
		{key: "https://example.com", val: []byte("testdata")},
		{key: "https://example.com/path", val: []byte("moretestdata")},
	}

	cache := NewCache()
	for _, c := range cases {
		cache.Add(c.key, c.val)
	}

	for _, c := range cases {
		val, ok := cache.Get(c.key)
		if !ok {
			t.Errorf("expected to find key %q", c.key)
			continue
		}
		if string(val) != string(c.val) {
			t.Errorf("key %q: got %q, want %q", c.key, val, c.val)
		}
	}
}

func TestGetMissingKey(t *testing.T) {
	cache := NewCache()
	val, ok := cache.Get("missing")
	if ok {
		t.Errorf("expected ok to be false for missing key")
	}
	if val != nil {
		t.Errorf("expected nil value for missing key, got %q", val)
	}
}

func TestReplaceValue(t *testing.T) {
	cache := NewCache()
	cache.Add("key", []byte("first"))
	cache.Add("key", []byte("second"))

	val, ok := cache.Get("key")
	if !ok {
		t.Fatalf("expected to find key")
	}
	if string(val) != "second" {
		t.Errorf("got %q, want %q", val, "second")
	}
}

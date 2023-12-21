package jplog

import "testing"

func TestLog(t *testing.T) {
	log := New().WithGroup("myapp")
	log.Info("a msg", "hello", "world", "foo", "bar")
}

package jplog_test

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/jpillora/jplog"
	"github.com/stretchr/testify/require"
)

func TestLog(t *testing.T) {
	b := strings.Builder{}
	log := jplog.New(&b).WithGroup("myapp")
	log.Info("a msg", "hello", "world", "foo", "bar")
	require.Contains(t, b.String(), "INFO")
	require.Contains(t, b.String(), "hello=world")
	require.Contains(t, b.String(), "foo=bar")
}

func TestLogJSON(t *testing.T) {
	b := bytes.Buffer{}
	os.Setenv("LOG_JSON", "1")
	log := jplog.New(&b).WithGroup("myapp")
	log.Info("a msg", "hello", "world", "foo", "bar")
	type line struct {
		Time  string `json:"time"`
		Level string `json:"level"`
		Msg   string `json:"msg"`
		Myapp struct {
			Hello string `json:"hello"`
			Foo   string `json:"foo"`
		} `json:"myapp"`
	}
	d := json.NewDecoder(&b)
	d.DisallowUnknownFields()
	l := line{}
	if err := d.Decode(&l); err != nil {
		t.Fatal(err)
	}
	require.NotEmpty(t, l.Time)
	require.Equal(t, "INFO", l.Level)
	require.Equal(t, "a msg", l.Msg)
	require.Equal(t, "world", l.Myapp.Hello)
	require.Equal(t, "bar", l.Myapp.Foo)
}

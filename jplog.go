package jplog

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

func New(w io.Writer) *slog.Logger {
	// if you want to customise json handler, use slog.NewJSONHandler
	if s := os.Getenv("LOG_JSON"); s == "1" || s == "true" {
		return slog.New(slog.NewJSONHandler(w, nil))
	}
	return slog.New(Handler(w))
}

// JPHandler is jpillora's slog.Handler
type JPHandler interface {
	slog.Handler
	// enabled verbose logging
	Verbose() JPHandler
}

func Handler(w io.Writer) JPHandler {
	return &h{
		w: w,
	}
}

type h struct {
	w       io.Writer
	group   string
	verbose bool
	attrs   []slog.Attr
}

func (h *h) Enabled(_ context.Context, l slog.Level) bool {
	if h.verbose {
		return true
	}
	return l >= slog.LevelInfo
}

func (h h) Verbose() JPHandler {
	h.verbose = true
	return &h
}

func (h *h) WithAttrs(attrs []slog.Attr) slog.Handler {
	h2 := *h
	h2.attrs = append(h.attrs, attrs...)
	return &h2
}

func (h *h) WithGroup(name string) slog.Handler {
	h2 := *h
	if h2.group == "" {
		h2.group = name
	} else {
		h2.group += "." + name
	}
	return &h2
}

func (h *h) Handle(ctx context.Context, r slog.Record) error {
	sb := strings.Builder{}

	const format = `3:04:05PM 2/1/2006`

	sb.WriteString(cyan(r.Time.Format(format)))

	sb.WriteRune(' ')
	sb.WriteString(level(r.Level.String()))

	if h.group != "" {
		sb.WriteRune(' ')
		sb.WriteString(green(h.group))
	}

	sb.WriteRune(' ')
	sb.WriteString(white(r.Message))

	add := func(attr slog.Attr) bool {
		sb.WriteRune(' ')
		sb.WriteString(grey(attr.Key))
		sb.WriteRune('=')
		sb.WriteString(grey(attr.Value.String()))
		return true
	}
	r.Attrs(add)
	for _, attr := range h.attrs {
		add(attr)
	}
	fmt.Fprintln(h.w, sb.String())
	return nil
}

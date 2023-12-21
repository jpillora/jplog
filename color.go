package jplog

import (
	"strings"

	"github.com/muesli/termenv"
)

var (
	color = termenv.EnvColorProfile().Color
	// Azu http://gogh-co.github.io/Gogh/
	red    = termenv.Style{}.Foreground(color("#AC6D74")).Styled
	green  = termenv.Style{}.Foreground(color("#74AC6D")).Styled
	yellow = termenv.Style{}.Foreground(color("#ACA46D")).Styled
	blue   = termenv.Style{}.Foreground(color("#6D74AC")).Styled
	pink   = termenv.Style{}.Foreground(color("#A46DAC")).Styled
	cyan   = termenv.Style{}.Foreground(color("#6DACA4")).Styled
	white  = termenv.Style{}.Foreground(color("#E6E6E6")).Styled
	grey   = termenv.Style{}.Foreground(color("#B3B1AD")).Styled
)

func level(level string) string {
	p := strings.ToUpper(level)
	if len(p) < 5 {
		p = strings.Repeat(" ", 5-len(level)) + p
	}
	switch strings.ToLower(level) {
	case "debug":
		return cyan(p)
	case "info":
		return blue(p)
	case "warn":
		return yellow(p)
	case "error":
		return red(p)
	case "fatal":
		return pink(p)
	}
	return white(p)
}

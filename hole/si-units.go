package hole

import (
	"fmt"
	"strings"
)

var siPrefixes = [...]struct {
	symbol   string
	exponent int
}{
	{"Q", 30},
	{"R", 27},
	{"Y", 24},
	{"Z", 21},
	{"E", 18},
	{"P", 15},
	{"T", 12},
	{"G", 9},
	{"M", 6},
	{"k", 3},
	{"h", 2},
	{"da", 1},
	{"", 0},
	{"d", -1},
	{"c", -2},
	{"m", -3},
	{"μ", -6},
	{"n", -9},
	{"p", -12},
	{"f", -15},
	{"a", -18},
	{"z", -21},
	{"y", -24},
	{"r", -27},
	{"q", -30},
}

var siUnitsMap = [...]struct{ symbol, expansion string }{
	{"s", "s"},
	{"m", "m"},
	{"g", "kg"},
	{"A", "A"},
	{"K", "K"},
	{"mol", "mol"},
	{"cd", "cd"},
	{"rad", ""},
	{"sr", ""},
	{"Hz", "s^-1"},
	{"N", "kg m s^-2"},
	{"Pa", "kg m^-1 s^-2"},
	{"J", "kg m^2 s^-2"},
	{"W", "kg m^2 s^-3"},
	{"C", "A s"},
	{"V", "kg m^2 s^-3 A^-1"},
	{"F", "kg^-1 m^-2 s^4 A^2"},
	{"Ω", "kg m^2 s^-3 A^-2"},
	{"S", "kg^-1 m^-2 s^3 A^2"},
	{"Wb", "kg m^2 s^-2 A^-1"},
	{"T", "kg s^-2 A^-1"},
	{"H", "kg m^2 s^-2 A^-2"},
	{"°C", "K"},
	{"lm", "cd"},
	{"lx", "cd m^-2"},
	{"Bq", "s^-1"},
	{"Gy", "m^2 s^-2"},
	{"Sv", "m^2 s^-2"},
	{"kat", "mol s^-1"},
}

var _ = answerFunc("si-units", func() []Answer {
	tests := make([]test, 0, len(siPrefixes)*len(siUnitsMap))
	for _, prefix := range siPrefixes {
		for _, unit := range siUnitsMap {
			exp := prefix.exponent
			if unit.symbol == "g" {
				exp -= 3
			}
			out := fmt.Sprint("10^", exp, " ", unit.expansion)
			out = strings.ReplaceAll(strings.ReplaceAll(out, "10^1 ", "10 "), "10^0 ", "1 ")
			tests = append(tests, test{prefix.symbol + unit.symbol, strings.Trim(out, " ")})
		}
	}
	return outputMultirunTests(tests)
})

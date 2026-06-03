package gohtml

import (
	"bytes"
	"strconv"
	"strings"
)

func Format(s string) string {
	return parse(strings.NewReader(s)).html()
}

func FormatBytes(b []byte) []byte {
	return parse(bytes.NewReader(b)).bytes()
}

func FormatWithLineNo(s string) string {
	return AddLineNo(Format(s))
}

func AddLineNo(s string) string {
	lines := strings.Split(s, "\n")
	maxLineNoStrLen := len(strconv.Itoa(len(lines)))
	bf := &bytes.Buffer{}
	for i, line := range lines {
		lineNoStr := strconv.Itoa(i + 1)
		if i > 0 {
			bf.WriteString("\n")
		}
		bf.WriteString(strings.Repeat(" ", maxLineNoStrLen-len(lineNoStr)) + lineNoStr + "  " + line)
	}
	return bf.String()

}

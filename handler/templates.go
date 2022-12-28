package handler

import (
	"html/template"
	"strings"
)

// used in html templates for page list
func IntRange(n int) []int {
	res := make([]int, n)
	if n == 0 {
		res[0] = 0
		return res
	}
	for i := 0; i < n; i++ {
		res[i] = i
	}
	return res
}

// add line breakes
func FormatBody(body string) template.HTML {
	return template.HTML(strings.Replace(template.HTMLEscapeString(body), "\n", "<br>", -1))
}

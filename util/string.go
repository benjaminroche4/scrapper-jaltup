package util

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	EmailRegEx = regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}`)
	URLRegEx   = regexp.MustCompile(
		`(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/|\/|\/\/)?[A-z0-9_-]*?[:]?[A-z0-9_-]*?[@]?[A-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?`, // nolint: lll
	)
)

func Truncate(in string, maxlen int) string {
	if len(in) > maxlen {
		return in[:maxlen]
	}

	return in
}

func CleanCityName(in string) string {
	parts := strings.Split(in, "-")
	c := cases.Title(language.French)
	r := regexp.MustCompile("^[^0-9]+$")

	for _, part := range parts {
		field := strings.TrimSpace(part)

		if r.MatchString(field) {
			return c.String(field)
		}
	}

	return c.String(in)
}

func CleanEmail(in string) string {
	return EmailRegEx.FindString(in)
}

func CleanURL(in string) string {
	return URLRegEx.FindString(in)
}

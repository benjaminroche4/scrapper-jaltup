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
	titleRegEx1 = regexp.MustCompile(`(?i)\s*(\()?\s*[hf]\s*[\/_-]\s*[hf]\s*(\))?`)
	titleRegEx2 = regexp.MustCompile(`(?i)\s*(-)?\s*(en\s+)?(en\s+contrat\s+d'\s*)?(alternance|apprentissage)\s*(-)?`)
	titleRegEx3 = regexp.MustCompile(`(?i)\s*([-(])?\s*(bac\s*\+[0-9]+)(\s+[aà]\s+[0-9]+)?\s*([-)])?`)
	titleRegEx4 = regexp.MustCompile(`(?i)\s*([-(]:)?\s*(mesure\s+poei)\s*([-):])?`)
	titleRegEx5 = regexp.MustCompile(`(?i)(mesure\s+poei|[*]{2,})\s*`)
)

func Truncate(in string, maxlen int) string {
	if len(in) > maxlen {
		return in[:maxlen]
	}

	return in
}

func CleanTitle(in string) string {
	out := titleRegEx1.ReplaceAllString(in, "")
	out = titleRegEx2.ReplaceAllString(out, "")
	out = titleRegEx3.ReplaceAllString(out, "")
	out = titleRegEx4.ReplaceAllString(out, "")
	out = titleRegEx5.ReplaceAllString(out, "")
	out = strings.Trim(out, "/")
	out = strings.Trim(out, "-")
	out = strings.TrimSpace(out)
	out = strings.ToLower(out)

	return out
}

func CleanCityName(in string) string {
	c := cases.Title(language.French)
	r := regexp.MustCompile(`[^0-9 \-_:,.=<>][^0-9]{1,}`)

	trimed := strings.TrimSpace(in)
	str := r.FindString(trimed)
	if str != "" {
		return c.String(strings.TrimSpace(str))
	}

	return c.String(in)
}

func CleanEmail(in string) string {
	return EmailRegEx.FindString(in)
}

func CleanURL(in string) string {
	return URLRegEx.FindString(in)
}

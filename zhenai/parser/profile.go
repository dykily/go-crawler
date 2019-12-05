package parser

import (
	"go-crawler/engine"
	"go-crawler/model"
	"regexp"
)
const userRe = `<div class="m-btn purple"[^>]*>([^<]+)</div>`
func ParserProfile(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(userRe)
	matches := re.FindAllSubmatch(contents, -1)
	profile  := model.Profile{}
	for i, m := range matches {
		switch i {
		case 0:
			profile.Marriage = string(m[1])
		case 1:
			profile.Age = string(m[1])
		case 3:
			profile.Height = string(m[1])
		case 4:
			profile.Weight = string(m[1])
		case 5:
			profile.WorkPlace = string(m[1])
		case 6:
			profile.Income = string(m[1])
		case 7:
			profile.Work = string(m[1])
		case 8:
			profile.Education = string(m[1])
		}
	}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}
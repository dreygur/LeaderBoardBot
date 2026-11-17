// Package mention reads Discord mention markup.
package mention

import "regexp"

// Discord dropped the "!" nickname form; old clients and bots may still send it.
var userMention = regexp.MustCompile(`<@!?(\d+)>`)

// UserID returns the user ID inside the first user mention in s, or "".
func UserID(s string) string {
	match := userMention.FindStringSubmatch(s)
	if len(match) < 2 {
		return ""
	}
	return match[1]
}

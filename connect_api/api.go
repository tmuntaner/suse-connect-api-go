package connect_api

import (
	"io"
	"bytes"
	"strings"
	"regexp"
)

type Service struct {
	Username string
	Password string
}

func readBody(closer io.ReadCloser) []byte {

	buf := new(bytes.Buffer)
	buf.ReadFrom(closer)

	body := buf.String()

	return []byte(body)
}

func getNext(linkHeader string) string {
	links := strings.Split(linkHeader, ",")
	next := ""
	r := regexp.MustCompile("<(.*)>; rel=\"(.*)\"")

	for i, linkString := range links {
		if i > 0 && r.Match([]byte(linkString)) {
			matches := r.FindStringSubmatch(linkString)
			if matches[2] == "next" {
				next = matches[1]
			}
		}
	}

	return next
}

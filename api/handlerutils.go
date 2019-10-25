package api

import (
	"net/url"
	"strconv"
)

type Links map[string]string

func getPaginationLinks(url url.URL, total, page, limit int) map[string]string {
	links := make(Links)

	values := url.Query()

	values.Set("page", strconv.Itoa(page))
	url.RawQuery = values.Encode()
	links["self"] = url.RequestURI()

	if page != 0 {
		values.Set("page", strconv.Itoa(page-1))
		url.RawQuery = values.Encode()

		links["prev"] = url.RequestURI()
	}

	if limit != 0 && (page+1)*limit < total {
		values.Set("page", strconv.Itoa(page+1))
		url.RawQuery = values.Encode()

		links["next"] = url.RequestURI()
	}

	return links
}

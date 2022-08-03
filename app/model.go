package app

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"grest.dev/grest/db"
)

type ListModel struct {
	Count         int64                    `json:"count"`
	Page          int64                    `json:"page_context.page"`
	PerPage       int64                    `json:"page_context.per_page"`
	PageCount     int64                    `json:"page_context.total_pages"`
	LinksFirst    string                   `json:"links.first"`
	LinksPrevious string                   `json:"links.previous"`
	LinksNext     string                   `json:"links.next"`
	LinksLast     string                   `json:"links.last"`
	Results       []map[string]interface{} `json:"results"`
}

func (l *ListModel) SetLink(c *fiber.Ctx) {
	q := ParseQuery(c)
	q.Set(db.QueryLimit, strconv.Itoa(int(l.PerPage)))

	path := strings.Split(c.OriginalURL(), "?")[0] + "?"

	first := q
	first.Del(db.QueryPage)
	first.Add(db.QueryPage, "1")
	l.LinksFirst = c.BaseURL() + path + first.Encode()

	if l.Page > 1 && l.PageCount > 1 {
		previous := q
		previous.Set(db.QueryPage, strconv.Itoa(int(l.Page-1)))
		l.LinksPrevious = c.BaseURL() + path + previous.Encode()
	}

	if l.Page < l.PageCount {
		next := q
		next.Set(db.QueryPage, strconv.Itoa(int(l.Page+1)))
		l.LinksNext = c.BaseURL() + path + next.Encode()
	}

	last := q
	last.Set(db.QueryPage, strconv.Itoa(int(l.PageCount)))
	l.LinksLast = c.BaseURL() + path + last.Encode()
}

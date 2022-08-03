package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"grest.dev/grest"
	"grest.dev/grest/telegram"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	lang := "en"
	ctx, ctxOK := c.Locals("ctx").(*Ctx)
	if ctxOK {
		lang = ctx.Lang
	}
	e, ok := err.(grest.Error)
	if !ok {
		fiberErr, isFiberErr := err.(*fiber.Error)
		if isFiberErr {
			e.Err.Code = fiberErr.Code
		}
		e.Err.Message = err.Error()
	} else {
		c.Locals("grestHRMErrors", &e)
	}
	if e.Err.Code < 400 || e.Err.Code > 599 {
		e.Err.Code = http.StatusInternalServerError
	}
	if e.Err.Code == http.StatusInternalServerError {
		errMsg := e.Err.Message
		errTrace := GetTrace(2)
		if ok {
			errTrace = e.TraceSimple()
		}
		go SendAlert(c, errMsg, errTrace)
		e.Err.Detail = map[string]string{"message": e.Err.Message}
		e.Err.Message = grest.Trans(lang, "500_internal_error")
	}
	return c.Status(e.Err.Code).JSON(e)
}

func NewErrorHandler(c *fiber.Ctx, statusCode int, message string, detail ...interface{}) error {
	err := grest.NewError(statusCode, message, detail...)
	return ErrorHandler(c, err)
}

func SendAlert(c *fiber.Ctx, message string, trace map[string]string) {
	env := APP_ENV
	logInfo := NewLogInfo()
	logInfo.Error = message
	logInfo.Method = c.Method()
	logInfo.Path = c.Path()
	logInfo.Trace = trace
	logInfo.IP = c.IPs()
	logInfo.Referer = c.Get("referer")
	if logInfo.Referer == "" {
		logInfo.Referer = c.BaseURL() + c.OriginalURL()
	}

	ctx, ctxOK := c.Locals("ctx").(*Ctx)
	if ctxOK {
		logInfo.ClientName = ctx.Token.ClientName
		logInfo.Email = ctx.Token.Email
		logInfo.Slug = ctx.Company.Slug
	}
	dataJson, _ := json.MarshalIndent(logInfo, "", "  ")
	if env == "production" || env == "development" {
		go telegram.SendAlert("```\n" + string(dataJson) + "```")
	} else {
		fmt.Println("----------------------------------------------------------------------")
		fmt.Println(string(dataJson))
		fmt.Println("----------------------------------------------------------------------")
	}
}

func ParseQuery(c *fiber.Ctx) url.Values {
	u := c.OriginalURL()
	q := strings.Split(u, "?")
	if len(q) > 1 {
		query, _ := url.ParseQuery(q[1])
		return query
	}
	return url.Values{}
}

func ErrorMessage(detail interface{}, defaultMessage string) string {
	message := defaultMessage
	errMap, ok := detail.(map[string]interface{})
	if ok && errMap["error"] != nil {
		errMapError, ok := errMap["error"].(map[string]interface{})
		if ok && errMapError["message"] != nil {
			msg, ok := errMapError["message"].(string)
			if ok {
				message = msg
			}
		}
	}
	return message
}

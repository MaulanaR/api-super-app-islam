package app

import (
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
	"grest.dev/grest"
	"grest.dev/grest/cache"
	"grest.dev/grest/db"
)

// Map is a shortcut for map[string]interface{}, useful for JSON object
type Map map[string]interface{}

func NewCode(tx *gorm.DB, tableName, fieldName, baseName string) string {
	replacer := strings.NewReplacer(",", "", ".", "", ";", "")
	l := ""
	baseName = replacer.Replace(baseName)
	words := strings.Fields(baseName)
	for _, word := range words {
		l += word[0:1]
	}
	l += "-"

	//get next code
	var nextCode int64
	tx.Table(tableName).Where(fieldName+" LIKE ?", l+"%").Count(&nextCode)
	nextCode += 1
	for i := 0; i < (3 - DigitLen(int(nextCode))); i++ {
		l += "0"
	}
	return strings.ToUpper(l) + strconv.FormatInt(nextCode, 10)
}

func IsCodeExists(ctx *Ctx, tableName, fieldName, code string) error {
	var isExists int64
	tx, _ := ctx.DB()
	tx.Table(tableName).Where(fieldName+" = ?", code).Where("deleted_at is null").Count(&isExists)
	if isExists >= 1 {
		langParam := []map[string]string{{"attribute": "Code", "value": code}}
		return grest.NewError(http.StatusBadRequest, grest.Trans(ctx.Lang, "unique", langParam...))
	}
	return nil
}

func DigitLen(number int) int {
	count := 0
	for number != 0 {
		number /= 10
		count += 1
	}
	return count
}

func StripSlashes(str string) string {
	var dstRune []rune
	strRune := []rune(str)
	strLenth := len(strRune)
	for i := 0; i < strLenth; i++ {
		if strRune[i] == []rune{'\\'}[0] {
			i++
		}
		dstRune = append(dstRune, strRune[i])
	}
	return string(dstRune)
}

type PageContext struct {
	Page       int `json:"page"        example:"1"`
	PerPage    int `json:"per_page"    example:"20"`
	TotalPages int `json:"total_pages" example:"1"`
}

var ExamplePageContext = PageContext{
	Page:       1,
	PerPage:    20,
	TotalPages: 1,
}

type CreatedUpdated struct {
	User CreatedUpdatedUser `json:"user"`
	Time time.Time          `json:"time" format:"date-time"`
}

type CreatedUpdatedUser struct {
	ID    string `json:"id"    format:"uuid"`
	Name  string `json:"name"  example:"Fulan"`
	Email string `json:"email" example:"user@email.com"`
}

var ExampleCreatedUpdated = CreatedUpdated{
	User: CreatedUpdatedUser{
		ID:    "3fa85f64-5717-4562-b3fc-2c963f66afa6",
		Name:  "Fulan",
		Email: "user@email.com",
	},
	Time: time.Now(),
}

type Deleted struct {
	Code    int    `json:"code"    example:"200"`
	Message string `json:"message" example:"Data has been deleted"`
}

type BadRequest struct {
	Error struct {
		Code    int    `json:"code" example:"400"`
		Message string `json:"message" example:"A validation exception occurred."`
	} `json:"error"`
}

type Unauthorized struct {
	Error struct {
		Code    int    `json:"code" example:"401"`
		Message string `json:"message" example:"Invalid authorization credentials."`
	} `json:"error"`
}

type Forbidden struct {
	Error struct {
		Code    int    `json:"code" example:"401"`
		Message string `json:"message" example:"User doesn't have permission to access the resource."`
	} `json:"error"`
}

func SetCompany(ctx *Ctx) error {
	companyID := ctx.Action.CompanyID
	if companyID != "" {
		comp := Company{}
		cacheKey := "companies." + companyID
		cache.Get(cacheKey, &comp)
		if comp.ID == "" {
			centralDB, err := db.DB("central")
			if err != nil {
				return err
			}
			if IsValid(companyID, "uuid") {
				centralDB = centralDB.Where("id = ?", companyID)
			} else {
				centralDB = centralDB.Where("slug = ?", companyID).Or("slug_alias = ?", companyID)
			}
			centralDB.Limit(1).Take(&comp)
			if comp.ID != "" {
				if !IS_ON_PREMISE && comp.MembershipID != 0 {
					membershipDB, err := db.DB("membership")
					if err != nil {
						return err
					}
					membershipDB.Table("member_db").Where("dbid = ?", comp.MembershipID).Select("status").Take(&comp)
				}
				comp.ConnName = companyID
				cache.Set(cacheKey, comp)
			}
		}
		if comp.ID != "" {
			ctx.Company = comp
		}
	}
	return nil
}

func FloatPrecision(num float64, precision int) float64 {
	p := math.Pow10(precision)
	value := float64(int(num*p)) / p
	return value
}

func DiffDate(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

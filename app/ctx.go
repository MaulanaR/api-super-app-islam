package app

import (
	"net/http"
	"time"

	//"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"grest.dev/grest"
	"grest.dev/grest/db"
)

const CtxKey = "ctx"

type Ctx struct {
	centralTx  *gorm.DB
	companyTx  *gorm.DB
	Lang       string
	Action     Action  `json:"-"`
	Token      Token   `json:"-"`
	User       User    `json:"-"`
	Company    Company `json:"-"`
	IsTest     bool    `json:"-"`
	IsPrintSQL bool    `json:"-"`
	//MockDB     sqlmock.Sqlmock `json:"-"`
}

type Action struct {
	Method    string
	EndPoint  string
	CompanyID string
	DataID    string
}

type Token struct {
	ID               string      `json:"-"                                   gorm:"column:uuid;primaryKey"`
	IsLocalTokenOnly db.NullBool `json:"is_local_token_only"                 gorm:"column:is_local_token_only"`
	CreatedAt        time.Time   `json:"created_at"                          gorm:"column:created_at"`
	UpdatedAt        time.Time   `json:"updated_at"                          gorm:"column:updated_at"`

	AccessToken    string `json:"access_token"                        gorm:"column:id"`
	ExpireTimeUnix int64  `json:"expires_in"                          gorm:"column:expire_time"`
	RefreshToken   string `json:"refresh_token"                       gorm:"column:refresh_token"`

	ZahirID              int64       `json:"id"                                  gorm:"column:zahir_id"`
	FirstName            string      `json:"first_name"                          gorm:"column:first_name"`
	LastName             string      `json:"last_name"                           gorm:"column:last_name"`
	Email                string      `json:"email"                               gorm:"column:email"`
	MobileNumber         string      `json:"mobile_number"                       gorm:"column:mobile_number"`
	TelegramUsername     string      `json:"telegram_username"                   gorm:"column:telegram_username"`
	WhatsappNumber       string      `json:"whatsapp_number"                     gorm:"column:whatsapp_number"`
	ClientID             string      `json:"client.id"                           gorm:"column:client_id"`
	ClientSecret         string      `json:"client.secret"                       gorm:"column:client_secret"` // diperlukan untuk melakukan refresh token
	ClientName           string      `json:"client.name"                         gorm:"column:client_name"`
	ClientType           string      `json:"client.type"                         gorm:"column:client_type"`
	IsInternalClient     db.NullBool `json:"client.internal_client"              gorm:"column:internal_client"`
	DeveloperID          int64       `json:"client.developer_account.id"         gorm:"column:developer_id"`
	DeveloperUsername    string      `json:"client.developer_account.username"   gorm:"column:developer_username"`
	DeveloperEmail       string      `json:"client.developer_account.email"      gorm:"column:developer_email"`
	DeveloperFirstName   string      `json:"client.developer_account.first_name" gorm:"column:developer_first_name"`
	DeveloperLastName    string      `json:"client.developer_account.last_name"  gorm:"column:developer_last_name"`
	ClientExpireTimeUnix int64       `json:"client.expires_in"                   gorm:"-"`
}

func (Token) TableName() string {
	return "access_tokens"
}

func (Token) TableVersion() string {
	return "22.03.310609"
}

type User struct {
	db.Model
	ID           string          `json:"id,omitempty"             db:"id"                                                     gorm:"column:id"`
	Name         string          `json:"name,omitempty"           db:"name"                                                   gorm:"column:name"`
	Email        string          `json:"email,omitempty"          db:"nama"                                                   gorm:"column:nama"`
	MobileNumber string          `json:"mobile_number,omitempty"  db:"mobile_phone"                                           gorm:"column:mobile_phone"`
	ZahirID      string          `json:"zahir_id,omitempty"       db:"zahir_id"                                               gorm:"column:zahir_id"`
	ContactID    string          `json:"contact_id,omitempty"     db:"contact_id"                                             gorm:"column:contact_id"`
	RoleID       string          `json:"role_id,omitempty"        db:"role_id"                                                gorm:"column:role_id"`
	IsSuperAdmin db.NullBool     `json:"is_super_admin,omitempty" db:"case when lower(master) = 'password' then 1 else 0 end" gorm:"column:master"`
	AclString    string          `json:"acl_string,omitempty"     db:"acl"                                                    gorm:"column:acl"`
	Acl          map[string]bool `json:"acl,omitempty"            db:"-"                                                      gorm:"-"`
}

func (User) TableName() string {
	return "sistem"
}

type Company struct {
	ConnName         string `gorm:"-"`
	ID               string
	Slug             string
	SlugAlias        string
	Name             string
	Description      string
	Community        string
	ImageUrl         string
	BusinessID       string
	BusinessTypeID   int64
	MembershipID     int64
	MembershipStatus string `gorm:"column:status"`
	DbDriver         string
	DbHost           string
	DbPort           int64
	DbDatabase       string
	DbUsername       string
	DbPassword       string
	IsActive         bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func NewCtx(c *fiber.Ctx) (*Ctx, error) {
	ctx, ok := c.Locals("ctx").(*Ctx)
	if !ok {
		return ctx, grest.NewError(http.StatusInternalServerError, "Context is not found")
	}
	return ctx, nil
}

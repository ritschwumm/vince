package models

import (
	"context"
	"database/sql"
	"math"
	"net/http"
	"net/mail"
	"strings"
	"time"

	"github.com/gernest/vince/config"
	"golang.org/x/crypto/bcrypt"
)

type currentUserKey struct{}

type User struct {
	Model
	Name         string
	Email        string `gorm:"uniqueIndex"`
	PasswordHash string
	Sites        []*Site `gorm:"many2many:site_memberships;"`
	OwnedSites   []*Site `gorm:"-"`

	EmailVerificationCodes  []*EmailVerificationCode `gorm:"constraint:OnDelete:CASCADE;"`
	IntroEmails             []*IntroEmail            `gorm:"constraint:OnDelete:CASCADE;"`
	FeedbackEmails          []*FeedbackEmail         `gorm:"constraint:OnDelete:CASCADE;"`
	CreateSiteEmails        []*CreateSiteEmail       `gorm:"constraint:OnDelete:CASCADE;"`
	CheckStatEmail          []*CheckStatEmail        `gorm:"constraint:OnDelete:CASCADE;"`
	SentRenewalNotification []*SentRenewalNotification
	APIKeys                 []*APIKey
	Subscription            *Subscription
	EnterprisePlan          *EnterprisePlan
	GoogleAuth              *GoogleAuth
	LastSeen                time.Time
	TrialExpiryDate         sql.NullTime
	EmailVerified           bool   `gorm:"not null;default:false"`
	Theme                   string `gorm:"not null;default:system"`
	GracePeriod             *GracePeriod
	Invitations             []*Invitation
}

func SetCurrentUser(ctx context.Context, usr *User) context.Context {
	return context.WithValue(ctx, currentUserKey{}, usr)
}

func GetCurrentUser(ctx context.Context) *User {
	if u := ctx.Value(currentUserKey{}); u != nil {
		return u.(*User)
	}
	return nil
}

// LoadUserModel returns a User model will all relations preloaded. This should be
// the biggest bottleneck performance  wise.
func LoadUserModel(ctx context.Context, uid uint64) (*User, error) {
	var u User
	db := Get(ctx)
	err := db.Preload("Sites").
		Preload("EmailVerificationCodes").
		Preload("IntroEmail").
		Preload("FeedbackEmails").
		Preload("CreateSiteEmails").
		Preload("SentRenewalNotification").
		Preload("APIKeys").
		Preload("Subscription").
		Preload("EnterprisePlan").
		Preload("EnterprisePlan").
		Preload("GoogleAuth").
		Preload("GracePeriod").
		Preload("Invitations").
		First(&u).Error
	if err != nil {
		return nil, err
	}
	err = db.Model(&Site{}).
		Joins("inner join  site_memberships on sites.id = site_memberships.site_id and site_memberships.role = 'owner' and site_memberships.user_id = ? ", u.ID).
		Scan(&u.OwnedSites).Error
	if err != nil {
		return nil, err
	}
	return &u, nil

}

func (u *User) SitesLimit(ctx context.Context) int {
	x := config.Get(ctx)
	if x.IsSelfHost {
		return math.MaxInt
	}
	for _, e := range x.SiteLimitExempt {
		if e == u.Email {
			return math.MaxInt
		}
	}
	if u.EnterprisePlan != nil {
		// check if a user has active plan
		if u.Subscription != nil {
			if u.EnterprisePlan.PlanID == u.Subscription.PlanID &&
				u.Subscription.Status == "active" {
				return int(u.EnterprisePlan.SiteLimit)
			}
		}
	}
	return 0
}

func (u *User) New(r *http.Request) (validation map[string]string, err error) {
	conf := config.Get(r.Context())
	u.Name = r.Form.Get("name")
	u.Email = r.Form.Get("email")
	password := r.Form.Get("password")
	passwordConfirm := r.Form.Get("password_confirmation")
	validation = make(map[string]string)
	validate("name", u.Name, "required", validation, func(s string) bool {
		return s != ""
	})
	validate("email", u.Email, "required", validation, func(s string) bool {
		return s != ""
	})
	validate("email", u.Email, "invalid email", validation, func(s string) bool {
		return emailRRe.MatchString(s)
	})
	validate("password", password, "required", validation, func(s string) bool {
		return s != ""
	})
	validate("password", password, "has to be at least 6 characters", validation, func(s string) bool {
		return len(s) >= 6
	})
	validate("password", password, "cannot be longer than 64 characters", validation, func(s string) bool {
		return len(s) <= 64
	})
	validate("password_confirmation", passwordConfirm, "required", validation, func(s string) bool {
		return s != ""
	})
	validate("password_confirmation", passwordConfirm, "must match password", validation, func(s string) bool {
		return s == password
	})
	if len(validation) != 0 {
		return
	}
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.PasswordHash = string(b)
	if conf.IsSelfHost {
		u.TrialExpiryDate = sql.NullTime{
			Time:  time.Now().AddDate(100, 0, 0),
			Valid: true,
		}
	} else {
		u.TrialExpiryDate = sql.NullTime{
			Time:  time.Now().AddDate(0, 0, 30),
			Valid: true,
		}
	}
	if !conf.EnableEmailVerification {
		u.EmailVerified = true
	}
	return
}

func (u *User) Recipient() string {
	return strings.Split(u.Name, " ")[0]
}

func (u *User) Address() *mail.Address {
	return &mail.Address{
		Name:    u.Name,
		Address: u.Email,
	}
}

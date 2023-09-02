package validator

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/oklog/ulid"
)

// NewValidator
func NewValidator() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("ISO8601", iso8601)
	return v
}

// ISO8601
func iso8601(fl validator.FieldLevel) bool {
	regString := `^(\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2}):(\d{2})(\.\d{0,3})?(Z|[+-](\d{2}):(\d{2}))$`
	reg := regexp.MustCompile(regString)
	return reg.MatchString(fl.Field().String())
}

// NewULID
func NewULID(t time.Time) string {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}

// DateParse
func DateParse(t string) (time.Time, error) {
	layouts := []string{
		"2006-01-02T15:04:05.000-07:00",
		"2006-01-02T15:04:05-07:00",
		"2006-01-02T15:04:05.000Z",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05",
		"2006-01-02",
	}
	for _, l := range layouts {
		if d, err := time.Parse(l, t); err == nil {
			return d, nil
		}
	}
	return time.Time{}, fmt.Errorf("unable to parse date %s, unrecognized format", t)
}

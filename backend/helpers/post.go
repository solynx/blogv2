package helpers

import (
	"strings"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

func GetSlug(value string, uid uuid.UUID) string {
	return slug.Make(value) + "-" + strings.Split(uid.String(), "-")[0]
}

package constants

type UserRole string
type Gender string

const (
	ENV_RUN_DEVELOPMENT = "development"
	ENV_RUN_TESTING     = "testing"
	ENV_RUN_PRODUCTION  = "production"

	Admon  UserRole = "Admin"
	Member UserRole = "Member"

	Male   Gender = "Male"
	Female Gender = "Female"
	NotSay Gender = "Prefer Not To Say"
)

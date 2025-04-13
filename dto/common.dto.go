package dto

import "errors"

const (
	FAILED_TO_BIND_DATA         = "Failed to bind data"
	FAILED_GET_DATA_FROM_BODY   = "Failed to get data from body"
	FAILED_GET_DATA_FROM_QUERY  = "Failed to get data from query"
	FAILED_GET_DATA_FROM_PARAMS = "Failed to get data from params"
	FAILED_GET_DATA_FROM_HEADER = "Failed to get data from header"
	FAILED_GET_DATA_FROM_COOKIE = "Failed to get data from cookie"
	FAILED_HEADER_IS_MISSING    = "Header is missing"
)

var (
	ErrMigrate             = errors.New("failed to migrate database")
	ErrSeed                = errors.New("failed to seed database")
	ErrFresh               = errors.New("failed to fresh database")
	ErrCreateEnum          = errors.New("failed to create enum")
	ErrDropEnum            = errors.New("failed to drop enum")
	ErrDropTable           = errors.New("failed to drop table")
	ErrBindData            = errors.New("failed to bind data")
	ErrOpenFile            = errors.New("failed to open file")
	ErrUnmarshalJSON       = errors.New("failed to unmarshal json")
	ErrUnsupportedFileType = errors.New("unsupported file type")
)

type (
	PaginationRequest struct {
		Page    int    `json:"page"     form:"page"`
		PerPage int    `json:"per_page" form:"per_page"`
		Search  string `json:"search"   form:"search"`
	}

	PaginationResponse struct {
		Page    int   `json:"page"`
		PerPage int   `json:"per_page"`
		MaxPage int   `json:"max_page"`
		Count   int64 `json:"count"`
	}
)

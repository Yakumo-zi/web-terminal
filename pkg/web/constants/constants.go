package constants

type CtxKey string

const (
	CtxRequestIdKey CtxKey = "request_id"
	CtxMethodKey    CtxKey = "method"
	CtxPathKey      CtxKey = "path"
)

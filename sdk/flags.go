package sdk

const (
	// constants
	P_OPENAPI = "openapi"
	P_API     = "api"
	P_PARAM2  = "param2"

	// system parameters
	P_ACCESS_TOKEN = "access_token"
	P_TIMESTAMP    = "_aop_timestamp"
	P_SIGN         = "_aop_signature"

	P_ERROR_CODE    = "error_code"
	P_ERROR_MESSAGE = "error_message"
	P_EXCEPTION     = "exception"
	P_REQUEST_ID    = "request_id"
)

var (
	headers = map[string]string{
		"Cache-Control": "no-cache",
		"Connection":    "Keep-Alive",
		"User-Agent":    "Ocean-SDK-Client",
	}
)

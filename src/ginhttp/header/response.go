package header

const (
	ResponseXTokenHeader     = "X-Token"
	ResponseAllowOrigin      = "Access-Control-Allow-Origin"
	ResponseAllowMethods     = "Access-Control-Allow-Methods"
	ResponseAllowHeaders     = "Access-Control-Allow-Headers"
	ResponseAllowCredentials = "Access-Control-Allow-Credentials"
	ResponseExposeHeaders    = "Access-Control-Expose-Headers"
	ResponseAllowMaxAge      = "Access-Control-Max-Age"
)

var ResponseHeaderList = []string{
	ResponseXTokenHeader,
	ResponseAllowOrigin,
	ResponseAllowMethods,
	ResponseAllowHeaders,
	ResponseAllowCredentials,
	ResponseExposeHeaders,
	ResponseAllowMaxAge,
}

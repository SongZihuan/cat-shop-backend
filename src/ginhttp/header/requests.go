package header

const (
	RequestsXTokenHeader    = "X-Token"
	RequestsAccept          = "Accept"
	RequestsContentType     = "Centent-Type"
	RequestsContentLength   = "Content-Length"
	RequestsOrigin          = "Origin"
	RequestsForwarded       = "Forwarded"
	RequestsXForwardedFor   = "X-Forwarded-For"
	RequestsXForwardedHost  = "X-Forwarded-Host"
	RequestsXForwardedProto = "X-Forwarded-Proto"
	RequestsXMessage        = "X-Message"
	RequestsXVia            = "Via"
)

var RequestsHeaderList = []string{
	RequestsXTokenHeader,
	RequestsAccept,
	RequestsContentType,
	RequestsContentLength,
	RequestsOrigin,
	RequestsForwarded,
	RequestsXForwardedFor,
	RequestsXForwardedHost,
	RequestsXForwardedProto,
	RequestsXMessage,
	RequestsXVia,
}

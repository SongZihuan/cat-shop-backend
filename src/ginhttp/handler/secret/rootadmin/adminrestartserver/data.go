package adminrestartserver

type Query struct {
	Password string `json:"password"`
	Secret   string `json:"secret"`
}

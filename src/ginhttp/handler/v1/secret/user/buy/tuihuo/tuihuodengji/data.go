package tuihuodengji

type Query struct {
	ID        uint   `form:"id"`
	KuaiDi    string `form:"kuaidi"`
	KuaiDiNum string `form:"kuaidinum"`
}

package ginplus

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

const NameEniger = "eniger"
const FindGroupURLByNameMaxDeep = 20

type handlerFuncGroup struct {
	get     gin.HandlerFunc
	post    gin.HandlerFunc
	options gin.HandlerFunc
}

func newHandlerFuncGroup() *handlerFuncGroup {
	return &handlerFuncGroup{}
}

func newHandlerFuncGroupAndApply(method string, handler gin.HandlerFunc) *handlerFuncGroup {
	group := newHandlerFuncGroup()
	group.Apply(method, handler)
	return group
}

func (h *handlerFuncGroup) Apply(method string, handler gin.HandlerFunc) bool {
	switch method {
	case http.MethodGet:
		if h.get != nil {
			return false
		}
		h.get = handler
		return true
	case http.MethodPost:
		if h.post != nil {
			return false
		}
		h.post = handler
		return true
	case http.MethodOptions:
		if h.options != nil {
			return false
		}
		h.options = handler
		return true
	default:
		return false
	}
}

type handlerFuncRecord struct {
	method  string
	handler gin.HandlerFunc
	path    string
}

func newHandlerFuncRecord(method string, handler gin.HandlerFunc, path string) *handlerFuncRecord {
	return &handlerFuncRecord{
		method:  method,
		handler: handler,
		path:    path,
	}
}

func (h *handlerFuncRecord) GetMethod() string {
	return h.method
}

func (h *handlerFuncRecord) GetHandler() gin.HandlerFunc {
	return h.handler
}

func (h *handlerFuncRecord) GetPath() string {
	return h.path
}

type handlerFuncRecordList struct {
	list []*handlerFuncRecord
}

func newHandleFuncRecordList() *handlerFuncRecordList {
	return &handlerFuncRecordList{
		list: make([]*handlerFuncRecord, 0, defaultHandlerFuncRecordListSize),
	}
}

func newHandleFuncRecordListAndApply(method string, handler gin.HandlerFunc, path string) *handlerFuncRecordList {
	record := newHandleFuncRecordList()
	record.Apply(method, handler, path)
	return record
}

func (h *handlerFuncRecordList) Apply(method string, handler gin.HandlerFunc, path string) {
	record := newHandlerFuncRecord(method, handler, path)
	h.list = append(h.list, record)
}

func (h *handlerFuncRecordList) Find(method string) []*handlerFuncRecord {
	if method == "" {
		return h.FindAll()
	}
	res := make([]*handlerFuncRecord, 0, len(h.list))
	for _, i := range h.list {
		if i.GetMethod() == method {
			res = append(res, i)
		}
	}
	return res
}

func (h *handlerFuncRecordList) FindAll() []*handlerFuncRecord {
	res := make([]*handlerFuncRecord, 0, len(h.list))
	for _, i := range h.list {
		res = append(res, i)
	}
	return res
}

type Router struct {
	name         string
	father       *Router
	son          []*Router
	engine       *gin.Engine
	router       gin.IRouter
	relativePath string
	path         string
	routerMap    map[string]*handlerFuncGroup
	handlerMap   map[uintptr]*handlerFuncRecordList
}

func debugPrintWARNINGDefault() {
	major, minor, _, err := utils.GetGoVersion()
	if err != nil {
		panic(err)
	}

	if major <= MinSupportGoMajor || minor <= MinSupportGoMinor {
		debugPrint("[WARNING] Now Gin requires Go %d.%d+.", MinSupportGoMajor, MinSupportGoMinor)
	}
}

func NewEngine() (*Router, error) {
	if !config.IsReady() {
		panic("config is not ready")
	}
	cfg := config.Config()

	gin.SetMode(cfg.Yaml.Global.GetGinMode())
	debugPrintWARNINGDefault()

	engine := gin.New()
	if cfg.Yaml.Global.IsDebug() {
		engine.Use(gin.Logger(), gin.Recovery(), Writer(), Recover())
	} else {
		engine.Use(gin.Logger(), Writer(), Recover())
	}

	if cfg.Yaml.Http.Proxy.Enable() {
		engine.ForwardedByClientIP = true
		err := engine.SetTrustedProxies(cfg.Yaml.Http.Proxy.TrustedIPs)
		if err != nil {
			return nil, err
		}
	}

	relativePath := utils.ProcessPath("/")
	return &Router{
		name:         NameEniger,
		father:       nil,
		son:          make([]*Router, 0, defaultRouterSonListSize),
		engine:       engine,
		router:       engine,
		relativePath: relativePath,
		path:         relativePath,
		routerMap:    make(map[string]*handlerFuncGroup, defaultRouterMapSize),
		handlerMap:   make(map[uintptr]*handlerFuncRecordList, defaultHandlerMapSize),
	}, nil
}

func newRouter(relativePath string, name string, r gin.IRouter, father *Router) *Router {
	if name == NameEniger {
		panic("bad router name")
	}

	relativePath = utils.ProcessPath(relativePath)
	routerPath := utils.ProcessPath(father.path + relativePath)

	son := &Router{
		name:         name,
		father:       father,
		son:          make([]*Router, 0, defaultRouterSonListSize),
		engine:       father.engine,
		router:       r,
		relativePath: relativePath,
		path:         routerPath,
		routerMap:    father.routerMap,
		handlerMap:   father.handlerMap,
	}

	father.son = append(father.son, son)

	return son
}

func (e *Router) Group(relativePath string, name ...string) *Router {
	relativePath = utils.ProcessPath(relativePath)
	next := e.router.Group(relativePath)

	if len(name) == 0 {
		return newRouter(relativePath, "", next, e)
	} else if len(name) == 1 {
		return newRouter(relativePath, name[0], next, e)
	} else {
		panic("too many name")
	}
}

func (e *Router) Use(middleware ...gin.HandlerFunc) {
	e.router.Use(middleware...)
}

func (e *Router) GET(relativePath string, handler gin.HandlerFunc) {
	e.router.GET(relativePath, handler)

	path := e.path + relativePath

	if group, ok := e.routerMap[path]; ok {
		if !group.Apply(http.MethodGet, handler) {
			panic("duplicated router [GET]: " + path)
		}
	} else {
		e.routerMap[path] = newHandlerFuncGroupAndApply(http.MethodGet, handler)
	}

	ptr := reflect.ValueOf(handler).Pointer()
	if record, ok := e.handlerMap[ptr]; ok {
		record.Apply(http.MethodGet, handler, path)
	} else {
		e.handlerMap[ptr] = newHandleFuncRecordListAndApply(http.MethodGet, handler, path)
	}
}

func (e *Router) POST(relativePath string, handler gin.HandlerFunc) {
	e.router.POST(relativePath, handler)

	path := e.path + relativePath

	if group, ok := e.routerMap[path]; ok {
		if !group.Apply(http.MethodPost, handler) {
			panic("duplicated router [POST]: " + path)
		}
	} else {
		e.routerMap[path] = newHandlerFuncGroupAndApply(http.MethodPost, handler)
	}

	ptr := reflect.ValueOf(handler).Pointer()
	if record, ok := e.handlerMap[ptr]; ok {
		record.Apply(http.MethodPost, handler, path)
	} else {
		e.handlerMap[ptr] = newHandleFuncRecordListAndApply(http.MethodPost, handler, path)
	}
}

func (e *Router) OPTIONS(relativePath string, handler gin.HandlerFunc) {
	e.router.OPTIONS(relativePath, handler)

	path := e.path + relativePath

	if group, ok := e.routerMap[path]; ok {
		if !group.Apply(http.MethodOptions, handler) {
			panic("duplicated router [OPTIONS]: " + path)
		}
	} else {
		e.routerMap[path] = newHandlerFuncGroupAndApply(http.MethodOptions, handler)
	}

	ptr := reflect.ValueOf(handler).Pointer()
	if record, ok := e.handlerMap[ptr]; ok {
		record.Apply(http.MethodOptions, handler, path)
	} else {
		e.handlerMap[ptr] = newHandleFuncRecordListAndApply(http.MethodOptions, handler, path)
	}
}

func (r *Router) FindURLByHandler(handler gin.HandlerFunc, method string) (string, bool) {
	ptr := reflect.ValueOf(handler).Pointer()
	recordPtr, ok := r.handlerMap[ptr]
	if !ok || recordPtr == nil {
		return "", false
	}

	record := (*recordPtr).Find(method)
	if len(record) == 0 {
		return "", false
	}

	return record[len(record)-1].GetPath(), true
}

func (r *Router) FindGroupURLByName(name string) (string, bool) {
	return r.findGroupURLByName(name, FindGroupURLByNameMaxDeep)
}

func (r *Router) findGroupURLByName(name string, deep int64) (string, bool) {
	if name == "" || deep == 0 {
		return "", false
	}

	if r.name == name {
		return r.path, true
	}

	for _, i := range r.son {
		if t, ok := i.findGroupURLByName(name, deep-1); ok {
			return t, true
		}
	}

	return "", false
}

func (r *Router) FindRouter(NotFound gin.HandlerFunc, NotMethod gin.HandlerFunc) {
	r.engine.NoRoute(NotFound)
	r.engine.NoMethod(NotMethod)
}

func (r *Router) ServeHTTP(wri http.ResponseWriter, req *http.Request) {
	r.engine.ServeHTTP(wri, req)
}

func (r *Router) NotRouter(NotFound gin.HandlerFunc, NotMethod gin.HandlerFunc) {
	r.engine.NoRoute(NotFound)
	r.engine.NoMethod(NotMethod)
}

func (r *Router) DebugMsg(format string, values ...any) {
	debugPrint(format, values...)
}

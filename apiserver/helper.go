package apiserver

import "github.com/gin-gonic/gin"

type Group struct {
	Url      string
	Handlers []gin.HandlerFunc
	get      map[string]gin.HandlerFunc
	post     map[string]gin.HandlerFunc
	put      map[string]gin.HandlerFunc
	patch    map[string]gin.HandlerFunc
	delete   map[string]gin.HandlerFunc
}

func (g *Group) SetUrl(url string) {
	g.Url = url
}

func (g *Group) Use(fn gin.HandlerFunc) {
	if g.Handlers == nil {
		g.Handlers = make([]gin.HandlerFunc, 0)
	}
	g.Handlers = append(g.Handlers, fn)
}

func (g *Group) Get(url string, fn gin.HandlerFunc) {
	if g.get == nil {
		g.get = make(map[string]gin.HandlerFunc)
	}
	g.get[url] = fn
}

func (g *Group) Post(url string, fn gin.HandlerFunc) {
	if g.post == nil {
		g.post = make(map[string]gin.HandlerFunc)
	}
	g.post[url] = fn
}

func (g *Group) Put(url string, fn gin.HandlerFunc) {
	if g.put == nil {
		g.put = make(map[string]gin.HandlerFunc)
	}
	g.put[url] = fn
}

func (g *Group) Patch(url string, fn gin.HandlerFunc) {
	if g.patch == nil {
		g.patch = make(map[string]gin.HandlerFunc)
	}
	g.patch[url] = fn
}

func (g *Group) Delete(url string, fn gin.HandlerFunc) {
	if g.delete == nil {
		g.delete = make(map[string]gin.HandlerFunc)
	}
	g.delete[url] = fn
}

func (g *Group) Export(r *gin.RouterGroup) {
	grp := r.Group(g.Url)
	for _, v := range g.Handlers {
		grp.Use(v)
	}

	for k, v := range g.get {
		grp.GET(k, v)
	}

	for k, v := range g.post {
		grp.POST(k, v)
	}

	for k, v := range g.put {
		grp.PUT(k, v)
	}

	for k, v := range g.patch {
		grp.PATCH(k, v)
	}

	for k, v := range g.delete {
		grp.DELETE(k, v)
	}
}

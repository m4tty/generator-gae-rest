package <%= entity.name %>DataMgr

import (
	"appengine"
	"appengine/datastore"
	"github.com/mjibson/goon"
)

type appEngine<%= _.capitalize(entity.name) %>DataManager struct {
	currentContext *appengine.Context
}

func NewAppEngine<%= _.capitalize(entity.name) %>DataManager(context *appengine.Context) *appEngine<%= _.capitalize(entity.name) %>DataManager {
	a := new(appEngine<%= _.capitalize(entity.name) %>DataManager)
	a.currentContext = context
	return a
}

//trying out goon...
func (dm appEngine<%= _.capitalize(entity.name) %>DataManager) Get<%= _.capitalize(entity.name) %>ById(id string) (<%= entity.name %> <%= _.capitalize(entity.name) %>, err error) {
	var ctx = *dm.currentContext
	g := goon.FromContext(ctx)
	<%= entity.name %> = <%= _.capitalize(entity.name) %>{Id: id}
	ctx.Infof("<%= entity.name %> get")
	err = g.Get(&<%= entity.name %>)
	ctx.Infof("<%= entity.name %> - " + <%= entity.name %>.Id)
	return
}

func (dm appEngine<%= _.capitalize(entity.name) %>DataManager) Get<%= _.capitalize(pluralize(entity.name) %>() (results []*<%= _.capitalize(entity.name) %>, err error) {
	var ctx = *dm.currentContext
	var <%= pluralize(entity.name) %> []*<%= _.capitalize(entity.name) %>

	g := goon.FromContext(ctx)
	q := datastore.NewQuery(g.Key(&<%= _.capitalize(entity.name) %>{}).Kind()).KeysOnly()
	keys, _ := g.GetAll(q, results)

	<%= pluralize(entity.name) %> = make([]*<%= _.capitalize(entity.name) %>, len(keys))
	for j, key := range keys {
		<%= pluralize(entity.name) %>[j] = &<%= _.capitalize(entity.name) %>{Id: key.StringID()}
	}
	err = g.GetMulti(<%= pluralize(entity.name) %>)
	results = <%= pluralize(entity.name) %>
	return
}

func (dm appEngine<%= _.capitalize(entity.name) %>DataManager) Save<%= _.capitalize(entity.name) %>(<%= entity.name %> *<%= _.capitalize(entity.name) %>) (key string, err error) {
	var ctx = *dm.currentContext
	g := goon.FromContext(ctx)
	g.Put(<%= entity.name %>)
	return
}

func (dm appEngine<%= _.capitalize(entity.name) %>DataManager) Delete<%= _.capitalize(entity.name) %>(id string) (err error) {
	var ctx = *dm.currentContext
	g := goon.FromContext(ctx)
	<%= entity.name %> := <%= _.capitalize(entity.name) %>{Id: id}
	err = g.Delete(g.Key(<%= entity.name %>))
	return
}

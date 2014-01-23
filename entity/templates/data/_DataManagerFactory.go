
<% _.each(entities, function (entity) { %>
package <%= entity.name %>DataMgr

import "appengine"

func Get<%= _.capitalize(entity.name) %>DataManager(context *appengine.Context) (<%= entity.name %>DataManager <%= _.capitalize(entity.name) %>DataManager) {
	var <%= entity.name %>DMgr = NewAppEngine<%= _.capitalize(entity.name) %>DataManager(context)
	<%= entity.name %>DataManager = <%= _.capitalize(entity.name) %>DataManager(<%= entity.name %>DMgr)
	return
}
<% }); %>
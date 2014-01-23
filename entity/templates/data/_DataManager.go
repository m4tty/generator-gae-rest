
<% _.each(entities, function (entity) { %>
package <%= entity.name %>DataMgr

type <%= _.capitalize(entity.name) %>DataManager interface {
	Get<%= _.capitalize(pluralize(entity.name)) %>() (results []*<%= _.capitalize(entity.name) %>, err error)
	Get<%= _.capitalize(entity.name) %>ById(id string) (result <%= _.capitalize(entity.name) %>, err error)
	Save<%= _.capitalize(entity.name) %>(<%= entity.name %> *<%= _.capitalize(entity.name) %>) (key string, err error)
	Delete<%= _.capitalize(entity.name) %>(id string) (err error)
}

<% }); %>
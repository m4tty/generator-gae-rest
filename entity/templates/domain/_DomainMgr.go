<% _.each(entities, function (entity) { %>

package <%= pluralize(entity.name) %>Domain

import (
	"<%= projectPath %>/data/<%= pluralize(entity.name) %>"
	"<%= projectPath %>/web/resources"
	"time"
)

type <%= _.capitalize(pluralize(entity.name)) %>Mgr struct {
	<%= entity.name %>DataMgr.<%= _.capitalize(entity.name) %>DataManager
}

func New<%= _.capitalize(pluralize(entity.name)) %>Mgr(bdm <%= entity.name %>DataMgr.<%= _.capitalize(entity.name) %>DataManager) *<%= _.capitalize(pluralize(entity.name)) %>Mgr {
	return &<%= _.capitalize(pluralize(entity.name)) %>Mgr{bdm}
}

func (dm <%= _.capitalize(pluralize(entity.name)) %>Mgr) Get<%= _.capitalize(entity.name) %>ById(id string) (<%= entity.name %> *resources.<%= _.capitalize(entity.name) %>Resource, err error) {
	d<%= _.capitalize(entity.name) %>, err := dm.<%= _.capitalize(entity.name) %>DataManager.Get<%= _.capitalize(entity.name) %>ById(id)

	if err != nil {
		return nil, err
	}
	var <%= entity.name %>Resource *resources.<%= _.capitalize(entity.name) %>Resource = new(resources.<%= _.capitalize(entity.name) %>Resource)

	mapDataToResource(&d<%= _.capitalize(entity.name) %>, <%= entity.name %>Resource)

	return <%= entity.name %>Resource, nil
}

func (dm <%= _.capitalize(pluralize(entity.name)) %>Mgr) Get<%= _.capitalize(pluralize(entity.name)) %>ByUserId(id string) (<%= pluralize(entity.name) %> []*resources.<%= _.capitalize(entity.name) %>Resource, err error) {
	d<%= _.capitalize(pluralize(entity.name)) %>, err := dm.<%= _.capitalize(entity.name) %>DataManager.Get<%= _.capitalize(pluralize(entity.name)) %>ByUserId(id)
	if err != nil {
		return nil, err
	}

	<%= pluralize(entity.name) %> = make([]*resources.<%= _.capitalize(entity.name) %>Resource, len(d<%= _.capitalize(pluralize(entity.name)) %>))
	for j, <%= entity.name %> := range d<%= _.capitalize(pluralize(entity.name)) %> {
		var <%= entity.name %>Resource *resources.<%= _.capitalize(entity.name) %>Resource = new(resources.<%= _.capitalize(entity.name) %>Resource)
		mapDataToResource(<%= entity.name %>, <%= entity.name %>Resource)
		<%= pluralize(entity.name) %>[j] = <%= entity.name %>Resource
	}
	return <%= pluralize(entity.name) %>, nil
}

func (dm <%= _.capitalize(pluralize(entity.name)) %>Mgr) Get<%= _.capitalize(pluralize(entity.name)) %>() (<%= pluralize(entity.name) %> []*resources.<%= _.capitalize(entity.name) %>Resource, err error) {
	d<%= _.capitalize(pluralize(entity.name)) %>, err := dm.<%= _.capitalize(entity.name) %>DataManager.Get<%= _.capitalize(pluralize(entity.name)) %>()
	if err != nil {
		return nil, err
	}

	<%= pluralize(entity.name) %> = make([]*resources.<%= _.capitalize(entity.name) %>Resource, len(d<%= _.capitalize(pluralize(entity.name)) %>))
	for j, <%= entity.name %> := range d<%= _.capitalize(pluralize(entity.name)) %> {
		var <%= entity.name %>Resource *resources.<%= _.capitalize(entity.name) %>Resource = new(resources.<%= _.capitalize(entity.name) %>Resource)
		mapDataToResource(<%= entity.name %>, <%= entity.name %>Resource)
		<%= pluralize(entity.name) %>[j] = <%= entity.name %>Resource
	}
	return <%= pluralize(entity.name) %>, nil
}

func (dm <%= _.capitalize(pluralize(entity.name)) %>Mgr) Save<%= _.capitalize(entity.name) %>(<%= entity.name %> *resources.<%= _.capitalize(entity.name) %>Resource) (key string, err error) {
	var d<%= _.capitalize(entity.name) %> *<%= entity.name %>DataMgr.<%= _.capitalize(entity.name) %> = new(<%= entity.name %>DataMgr.<%= _.capitalize(entity.name) %>)

	mapResourceToData(<%= entity.name %>, d<%= _.capitalize(entity.name) %>)

	key, saveErr := dm.<%= _.capitalize(entity.name) %>DataManager.Save<%= _.capitalize(entity.name) %>(d<%= _.capitalize(entity.name) %>)
	// if saveErr != nil {
	// 	return key, saveErr
	// }
	return key, saveErr
}

func (dm <%= _.capitalize(pluralize(entity.name)) %>Mgr) Delete<%= _.capitalize(entity.name) %>(id string) (err error) {
	deleteErr := dm.<%= _.capitalize(entity.name) %>DataManager.Delete<%= _.capitalize(entity.name) %>(id)
	return deleteErr
}

// mapper...
func mapResourceToData(<%= entity.name %>Resource *resources.<%= _.capitalize(entity.name) %>Resource, <%= entity.name %>Data *<%= entity.name %>DataMgr.<%= _.capitalize(entity.name) %>) {
	<%= entity.name %>Data.Id = <%= entity.name %>Resource.Id
	<% _.each(attrs, function (attr) { %>
	<%= entity.name %>Data.<%= attr.attrName %> = <%= entity.name %>Resource.<%= attr.attrName %>
	<% }); %>
}

func mapDataToResource(<%= entity.name %>Data *<%= entity.name %>DataMgr.<%= _.capitalize(entity.name) %>, <%= entity.name %>Resource *resources.<%= _.capitalize(entity.name) %>Resource) {
	<%= entity.name %>Resource.Id = <%= entity.name %>Data.Id

	<% _.each(attrs, function (attr) { %>
	<%= entity.name %>Resource.<%= _.capitalize(attr.attrName) %> = <%= entity.name %>Data.<%= _.capitalize(attr.attrName) %>
	<% }); %>
}





<% }); %>
<% _.each(entities, function (entity) { %>
	
package <%= pluralize(entity.name) %>Domain

import (
	"./data/<%= pluralize(entity.name) %>"
	"./web/resources"
	"time"
)

type <%= _.capitalize(pluralize(entity.name)) %>Mgr struct {
	<%= entity.name %>DataMgr.<%= _.capitalize(pluralize(entity.name)) %>DataManager
}

func New<%= _.capitalize(pluralize(entity.name)) %>Mgr(bdm <%= entity.name %>DataMgr.<%= _.capitalize(pluralize(entity.name)) %>DataManager) *<%= _.capitalize(pluralize(entity.name)) %>Mgr {
	return &<%= _.capitalize(pluralize(entity.name)) %>Mgr{bdm}
}

func (dm <%= _.capitalize(pluralize(entity.name)) %>Mgr) Get<%= _.capitalize(pluralize(entity.name)) %>ById(id string) (<%= entity.name %> *resources.<%= _.capitalize(pluralize(entity.name)) %>Resource, err error) {
	d<%= _.capitalize(pluralize(entity.name)) %>, err := dm.<%= _.capitalize(pluralize(entity.name)) %>DataManager.Get<%= _.capitalize(pluralize(entity.name)) %>ById(id)

	if err != nil {
		return nil, err
	}
	var <%= entity.name %>Resource *resources.<%= _.capitalize(pluralize(entity.name)) %>Resource = new(resources.<%= _.capitalize(pluralize(entity.name)) %>Resource)

	mapDataToResource(&d<%= _.capitalize(pluralize(entity.name)) %>, <%= entity.name %>Resource)

	return <%= entity.name %>Resource, nil
}



func (dm <%= _.capitalize(pluralize(entity.name)) %>Mgr) Get<%= _.capitalize(pluralize(entity.name)) %>() (<%= pluralize(entity.name) %> []*resources.<%= _.capitalize(pluralize(entity.name)) %>Resource, err error) {
	d<%= _.capitalize(pluralize(entity.name)) %>, err := dm.<%= _.capitalize(pluralize(entity.name)) %>DataManager.Get<%= _.capitalize(pluralize(entity.name)) %>()
	if err != nil {
		return nil, err
	}

	<%= pluralize(entity.name) %> = make([]*resources.<%= _.capitalize(pluralize(entity.name)) %>Resource, len(d<%= _.capitalize(pluralize(entity.name)) %>))
	for j, <%= entity.name %> := range d<%= _.capitalize(pluralize(entity.name)) %> {
		var <%= entity.name %>Resource *resources.<%= _.capitalize(pluralize(entity.name)) %>Resource = new(resources.<%= _.capitalize(pluralize(entity.name)) %>Resource)
		mapDataToResource(<%= entity.name %>, <%= entity.name %>Resource)
		<%= pluralize(entity.name) %>[j] = <%= entity.name %>Resource
	}
	return <%= pluralize(entity.name) %>, nil
}

func (dm <%= _.capitalize(pluralize(entity.name)) %>Mgr) Save<%= _.capitalize(pluralize(entity.name)) %>(<%= entity.name %> *resources.<%= _.capitalize(pluralize(entity.name)) %>Resource) (key string, err error) {
	var d<%= _.capitalize(pluralize(entity.name)) %> *<%= entity.name %>DataMgr.<%= _.capitalize(pluralize(entity.name)) %> = new(<%= entity.name %>DataMgr.<%= _.capitalize(pluralize(entity.name)) %>)

	mapResourceToData(<%= entity.name %>, d<%= _.capitalize(pluralize(entity.name)) %>)

	key, saveErr := dm.<%= _.capitalize(pluralize(entity.name)) %>DataManager.Save<%= _.capitalize(pluralize(entity.name)) %>(d<%= _.capitalize(pluralize(entity.name)) %>)
	// if saveErr != nil {
	// 	return key, saveErr
	// }
	return key, saveErr
}

func (dm <%= _.capitalize(pluralize(entity.name)) %>Mgr) Delete<%= _.capitalize(pluralize(entity.name)) %>(id string) (err error) {
	deleteErr := dm.<%= _.capitalize(pluralize(entity.name)) %>DataManager.Delete<%= _.capitalize(pluralize(entity.name)) %>(id)
	return deleteErr
}

// mapper...
func mapResourceToData(<%= entity.name %>Resource *resources.<%= _.capitalize(pluralize(entity.name)) %>Resource, <%= entity.name %>Data *<%= entity.name %>DataMgr.<%= _.capitalize(pluralize(entity.name)) %>) {
	<%= entity.name %>Data.Id = <%= entity.name %>Resource.Id
	<%= entity.name %>Data.OwnerId = <%= entity.name %>Resource.OwnerId
	<%= entity.name %>Data.Name = <%= entity.name %>Resource.Name
	<%= entity.name %>Data.Description = <%= entity.name %>Resource.Description
	<%= entity.name %>Data.IsPublic = <%= entity.name %>Resource.IsPublic
	<%= entity.name %>Data.CreatedDate = <%= entity.name %>Resource.CreatedDate
	<%= entity.name %>Data.LastModified = time.Now().UTC()
	<%= entity.name %>Data.Stars = <%= entity.name %>Resource.Stars
	<%= entity.name %>Data.Likes = <%= entity.name %>Resource.Likes
	<%= entity.name %>Data.Dislikes = <%= entity.name %>Resource.Dislikes
	<%= entity.name %>Data.LikedBy = <%= entity.name %>Resource.LikedBy
	<%= entity.name %>Data.DislikedBy = <%= entity.name %>Resource.DislikedBy
	<%= entity.name %>Data.Tags = <%= entity.name %>Resource.Tags
}

func mapDataToResource(<%= entity.name %>Data *<%= entity.name %>DataMgr.<%= _.capitalize(pluralize(entity.name)) %>, <%= entity.name %>Resource *resources.<%= _.capitalize(pluralize(entity.name)) %>Resource) {
	<%= entity.name %>Resource.Id = <%= entity.name %>Data.Id
	<%= entity.name %>Resource.OwnerId = <%= entity.name %>Data.OwnerId
	<%= entity.name %>Resource.Name = <%= entity.name %>Data.Name
	<%= entity.name %>Resource.Description = <%= entity.name %>Data.Description
	<%= entity.name %>Resource.IsPublic = <%= entity.name %>Data.IsPublic
	<%= entity.name %>Resource.CreatedDate = <%= entity.name %>Data.CreatedDate
	<%= entity.name %>Resource.LastModified = <%= entity.name %>Data.LastModified
	<%= entity.name %>Resource.Stars = <%= entity.name %>Data.Stars
	<%= entity.name %>Resource.Likes = <%= entity.name %>Data.Likes
	<%= entity.name %>Resource.Dislikes = <%= entity.name %>Data.Dislikes
	<%= entity.name %>Resource.LikedBy = <%= entity.name %>Data.LikedBy
	<%= entity.name %>Resource.DislikedBy = <%= entity.name %>Data.DislikedBy
	<%= entity.name %>Resource.Tags = <%= entity.name %>Data.Tags
}
<% }); %>
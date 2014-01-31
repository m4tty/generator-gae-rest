
<% _.each(entities, function (entity) { %>
package <%= entity.name %>DataMgr

type <%= _.capitalize(entity.name) %> struct {
    Id string `datastore:"-" goon:"id"` 
    <% _.each(attrs, function (attr) { %>
    <%= _.capitalize(attr.attrName) %> <% if (attr.attrType == 'Enum') { %>string<% } else { %><%= attr.attrImplType %><% }; %>
  <% }); %>
}
<% }); %>
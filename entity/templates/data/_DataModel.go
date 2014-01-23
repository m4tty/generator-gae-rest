package models

type <%= _.capitalize(name) %> struct {
    Id string `datastore:"-" goon:"id"` 
    <% _.each(attrs, function (attr) { %>
    <%= _.capitalize(attr.attrName) %> <% if (attr.attrType == 'Enum') { %>string<% } else { %><%= attr.attrImplType %><% }; %>
  <% }); %>
}

package resources

type <%= _.capitalize(name) %>Resource struct {
    Id int `json:"id"`
    <% _.each(attrs, function (attr) { %>
    <%= _.capitalize(attr.attrName) %> <% if (attr.attrType == 'Enum') { %>string<% } else { %><%= attr.attrImplType %><% }; %> `json:"<%= attr.attrName %>"`
  <% }); %>
}

package main

import (
    "appengine"
    "github.com/gorilla/mux"
    "<%= projectPath %>/web/handlers"
    "github.com/mjibson/appstats"
    "net/http"
)


var router = new(mux.Router)

func init() {
    <% _.each(entities, function (entity) { %>
    router.Handle("/<%= pluralize(entity.name) %>", appstats.NewHandler(handlers.GetAll<%= _.capitalize(pluralize(entity.name)) %>Handler)).Name("<%= pluralize(entity.name) %>-getall").Methods("GET")
    router.Handle("/<%= pluralize(entity.name) %>", appstats.NewHandler(handlers.Add<%= _.capitalize(entity.name) %>Handler)).Name("<%= pluralize(entity.name) %>-create").Methods("POST")
    router.Handle("/<%= pluralize(entity.name) %>/{id}", appstats.NewHandler(handlers.Get<%= _.capitalize(entity.name) %>)).Name("<%= pluralize(entity.name) %>-get").Methods("GET")
    router.Handle("/<%= pluralize(entity.name) %>/{id}", appstats.NewHandler(handlers.Update<%= _.capitalize(entity.name) %>Handler)).Name("<%= pluralize(entity.name) %>-update").Methods("PUT")
    router.Handle("/<%= pluralize(entity.name) %>/{id}", appstats.NewHandler(handlers.Delete<%= _.capitalize(entity.name) %>Handler)).Name("<%= pluralize(entity.name) %>-delete").Methods("DELETE")
    <% }); %>
    http.Handle("/", router)
}


<% _.each(entities, function (entity) { %>

package handlers

import "encoding/json"
import "net/http"
import "fmt"

import "github.com/gorilla/mux"

import "./web/resources"
import "./web/domain/<%= pluralize(entity.name) %>"
import "./data/<%= pluralize(entity.name) %>"
import "appengine"
import "appengine/user"
import "appengine/datastore"

const <%= _.capitalize(entity.name) %>TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"



func Get<%= _.capitalize(pluralize(entity.name)) %>Handler(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	u := user.Current(c)

	if u == nil {
		http.Error(w, http.StatusText(401), 401)
		return
	}

	dataManager := <%= entity.name %>DataMgr.GetDataManager(&c)
	dataMgr := <%= pluralize(entity.name) %>Domain.New<%= _.capitalize(pluralize(entity.name)) %>Mgr(dataManager)
	result, err := dataMgr.Get<%= _.capitalize(pluralize(entity.name)) %>()

	//err = errors.New("asdf")
	if err != nil {
		serveError(c, w, err)
		return
	}

	js, error := json.MarshalIndent(result, "", "  ")
	if error != nil {
		serveError(c, w, error)
		return
	}
	w.Write(js)
	return
}

func Get<%= _.capitalize(entity.name) %>Handler(c appengine.Context, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	<%= entity.name %>Id := vars["id"]

	dataManager := <%= entity.name %>DataMgr.GetDataManager(&c)
	dataMgr := <%= pluralize(entity.name) %>Domain.New<%= _.capitalize(pluralize(entity.name)) %>Mgr(dataManager)
	result, err := dataMgr.Get<%= _.capitalize(entity.name) %>ById(<%= entity.name %>Id)
	if err != nil {
		if err == datastore.ErrNoSuchEntity {
			http.Error(w, "Not Found", 404)
			return
		} else {
			serveError(c, w, err)
			return
		}

	} else {
		if result != nil {
			if checkLastModified(w, r, result.LastModified) {
				return
			}

			js, error := json.MarshalIndent(result, "", "  ")
			if error != nil {
				serveError(c, w, error)
				return
			}

			w.Write(js)
		}
	}

}


func Delete<%= _.capitalize(entity.name) %>Handler(c appengine.Context, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	<%= entity.name %>Id := vars["id"]
	dataManager := <%= entity.name %>DataMgr.GetDataManager(&c)
	dataMgr := <%= pluralize(entity.name) %>Domain.New<%= _.capitalize(pluralize(entity.name)) %>Mgr(dataManager)
	err := dataMgr.Delete<%= _.capitalize(entity.name) %>(<%= entity.name %>Id)
	if err != nil {
		serveError(c, w, err)
	}
}

func Add<%= _.capitalize(entity.name) %>Handler(c appengine.Context, w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var <%= entity.name %> resources.<%= _.capitalize(entity.name) %>Resource
	err := decoder.Decode(&<%= entity.name %>)
	if err != nil {
		serveError(c, w, err)
	}

	dataManager := <%= entity.name %>DataMgr.GetDataManager(&c)
	dataMgr := <%= pluralize(entity.name) %>Domain.New<%= _.capitalize(pluralize(entity.name)) %>Mgr(dataManager)
	_, saveErr := dataMgr.Save<%= _.capitalize(entity.name) %>(&<%= entity.name %>)
	//TODO: return location header w/ the id that was created during save
	if saveErr != nil {
		serveError(c, w, saveErr)
	}
}

func CreateUser<%= _.capitalize(entity.name) %>Handler(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	vars := mux.Vars(r)
	<%= entity.name %>Id := vars["<%= entity.name %>Id"]
	userId := vars["userId"]
	fmt.Fprint(w, "single <%= entity.name %>"+<%= entity.name %>Id)

	decoder := json.NewDecoder(r.Body)
	var <%= entity.name %> resources.<%= _.capitalize(entity.name) %>Resource
	err := decoder.Decode(&<%= entity.name %>)
	if err != nil {
		serveError(c, w, err)
	}
	<%= entity.name %>.OwnerId = userId

	dataManager := <%= entity.name %>DataMgr.GetDataManager(&c)
	dataMgr := <%= pluralize(entity.name) %>Domain.New<%= _.capitalize(pluralize(entity.name)) %>Mgr(dataManager)
	_, saveErr := dataMgr.Save<%= _.capitalize(entity.name) %>(&<%= entity.name %>)

	if saveErr != nil {
		serveError(c, w, saveErr)
	}
}
<% }); %>
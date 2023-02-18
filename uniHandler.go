package university_search

import "net/http"

/*
Entry point for handler of NeighbourUni information
*/
func UniHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		handleGetUni(w, r)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' is supported.", http.StatusNotImplemented)
		return
	}
}

func handleGetUni(w http.ResponseWriter, r *http.Request) {

}

package university_search

import "net/http"

/*
Entry point for handler of NeighbourUni information
*/
func NeighbourHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		handleGetNeighbour(w, r)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' is supported.", http.StatusNotImplemented)
		return
	}
}

func handleGetNeighbour(w http.ResponseWriter, r *http.Request) {

}

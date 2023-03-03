package university_search

import (
	"fmt"
	"net/http"
)

/*
Empty handler as default handler
*/
func EmptyHandler(w http.ResponseWriter, r *http.Request) {

	// Ensure that the server interprets requests as HTML from Client (browser)
	w.Header().Set("content-type", "text/html")

	// Give information for correct usage (paths)
	output := fmt.Sprintf(
		`This service does not provide functionality at this path. 
			    Use <a href="%s">%s</a> or <a href="%s">%s</a>.
			    For diagnostic information about the service, visit: <a href="%s">%s</a>`,
		UNIINFO_PATH, UNIINFO_PATH, NEIGHBOUR_PATH, NEIGHBOUR_PATH, DIAG_PATH, DIAG_PATH)

	// Make the output visible to the client
	_, err := fmt.Fprintf(w, "%v", output)

	// Deal with error, if any
	if err != nil {
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
	}
}

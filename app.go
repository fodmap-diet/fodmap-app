// Sample helloworld is an App Engine app.
package main

// [START import]
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	sdk "github.com/fodmap-diet/go-sdk"
)

func main() {
	http.HandleFunc("/search/", searchHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// searchHandler responds to a item search request
func searchHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search/" {
		http.NotFound(w, r)
		return
	}

	keys, ok := r.URL.Query()["item"]
	if !ok {
		log.Println("Url Param 'item' is missing")
		http.Error(w, "Url Param 'item' is missing", http.StatusBadRequest)
		return
	}

	items := make(map[string]interface{})

	for _, key := range keys {
		key = strings.ToLower(key)
		if len(key) == 0 {
			log.Println("Invalid item")
			http.Error(w, "Invalid item", http.StatusBadRequest)
			return
		}

		item, err := sdk.SearchItem(key)
		if err != nil {
			items[key] = struct {
				Error string `json: "error"`
			}{
				err.Error(),
			}
			continue
		}
		items[key] = item
	}

	js, err := json.MarshalIndent(items, "", "    ")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

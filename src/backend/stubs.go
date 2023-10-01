package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type SearchResult struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

var results = []SearchResult{
	{Title: "Stanford University", URL: "https://www.stanford.edu/", Description: "Description of Stanford University."},
	{Title: "The Massachusetts Institute of Technology", URL: "https://www.mit.edu/", Description: "Description of MIT."},
	{Title: "Harvard University", URL: "https://www.harvard.edu/", Description: "Description of Harvard University."},
	{Title: "Columbia University", URL: "https://www.columbia.edu/", Description: "Description of Columbia University."},
	{Title: "Yale University", URL: "https://www.yale.edu/", Description: "Description of Yale University."},
	{Title: "University of Cambridge", URL: "https://www.cam.ac.uk/", Description: "Description of University of Cambridge."},
	{Title: "University of Oxford", URL: "https://www.ox.ac.uk/", Description: "Description of University of Oxford."},
	{Title: "California Institute of Technology", URL: "https://www.caltech.edu/", Description: "Description of California Institute of Technology."},
	{Title: "Princeton University", URL: "https://www.princeton.edu/", Description: "Description of Princeton University."},
	{Title: "University of Chicago", URL: "https://www.uchicago.edu/", Description: "Description of University of Chicago."},
	{Title: "Imperial College London", URL: "https://www.imperial.ac.uk/", Description: "Description of Imperial College London."},
	{Title: "University of California, Berkeley", URL: "https://www.berkeley.edu/", Description: "Description of University of California, Berkeley."},
	{Title: "University of California, Los Angeles", URL: "https://www.ucla.edu/", Description: "Description of University of California, Los Angeles."},
	{Title: "University of Michigan", URL: "https://www.umich.edu/", Description: "Description of University of Michigan."},
	{Title: "University of Pennsylvania", URL: "https://www.upenn.edu/", Description: "Description of University of Pennsylvania."},
	{Title: "University of Toronto", URL: "https://www.utoronto.ca/", Description: "Description of University of Toronto."},
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	// Set the response headers to allow cross-origin requests.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

	// If the request method is OPTIONS, return an empty response with 200 status code.
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Convert the search results to JSON.
	jsonData, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Set the response content type.
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response back to the client.
	w.Write(jsonData)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/search", handleSearch)

	port := "8081" // Choose any available port you prefer.
	log.Println("Golang stub service is running on port", port)
	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}

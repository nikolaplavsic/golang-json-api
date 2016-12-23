package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Welcome to our App")
}

func AttributeList(w http.ResponseWriter, r *http.Request) {
	attributes := Attributes{
		Attribute{Label: "Title"},
		Attribute{Label: "Description"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(attributes); err != nil {
		panic(err)
	}
}

func AttributeDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	attrId := vars["attributeId"]
	fmt.Fprintf(w, "Detail of attribute with ID: %q", attrId)
}

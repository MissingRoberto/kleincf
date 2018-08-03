package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func (cc *CloudController) AssociateRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	response := map[string]interface{}{
		"metadata": map[string]interface{}{
			"guid": cc.Namespace,
		},
		"entity": map[string]interface{}{
			"space_guid": cc.Namespace,
			"host":       vars["appid"],
		},
	}

	bytes, _ := json.Marshal(response)
	w.Write(bytes)
	w.WriteHeader(http.StatusCreated)
}

func (cc *CloudController) AppRoutes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resources := []map[string]interface{}{
		map[string]interface{}{
			"metadata": map[string]interface{}{
				"guid": cc.Namespace,
			},
			"entity": map[string]interface{}{
				"space_guid":  cc.Namespace,
				"host":        vars["id"],
				"domain_guid": cc.Domain,
			},
		},
	}

	response := map[string]interface{}{
		"total_results": 1,
		"total_pages":   1,
		"resources":     resources,
	}

	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

func (cc *CloudController) Routes(w http.ResponseWriter, r *http.Request) {
	resources := []map[string]interface{}{
		map[string]interface{}{
			"metadata": map[string]interface{}{
				"guid": cc.Namespace,
			},
			"entity": map[string]interface{}{
				"space_guid": cc.Namespace,
				"host":       "host-7",
			},
		},
	}

	response := map[string]interface{}{
		"total_results": 1,
		"total_pages":   1,
		"resources":     resources,
	}

	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

func (cc *CloudController) Route(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var body map[string]interface{}
	err = json.Unmarshal(bytes, &body)
	if err != nil {
		panic(err)
	}

	response := map[string]interface{}{
		"metadata": map[string]interface{}{
			"guid": cc.Namespace,
		},
		"entity": map[string]interface{}{
			"space_guid":  cc.Namespace,
			"host":        body["host"],
			"domain_guid": cc.Domain,
		},
	}

	bytes, _ = json.Marshal(response)
	w.Write(bytes)
	w.WriteHeader(http.StatusCreated)
}

func (cc *CloudController) SharedDomain(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"metadata": map[string]interface{}{
			"guid": cc.Domain,
		},
		"entity": map[string]interface{}{
			"name": cc.Domain,
		},
	}
	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

func (cc *CloudController) SharedDomains(w http.ResponseWriter, r *http.Request) {
	domains := []map[string]interface{}{
		map[string]interface{}{
			"metadata": map[string]interface{}{
				"guid": cc.Domain,
			},
			"entity": map[string]interface{}{
				"name": cc.Domain,
			},
		},
	}

	response := map[string]interface{}{
		"total_results": 1,
		"total_pages":   1,
		"resources":     domains,
	}

	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

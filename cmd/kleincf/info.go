package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (cc *CloudController) V3(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"links": map[string]interface{}{
			"apps": map[string]interface{}{
				"href": fmt.Sprintf("%s/%s", cc.Address, "v3/apps"),
			},
			"builds": map[string]interface{}{
				"href": fmt.Sprintf("%s/%s", cc.Address, "v3/builds"),
			},
		},
	}
	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

func (cc *CloudController) API(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"links": map[string]interface{}{
			"cloud_controller_v3": map[string]interface{}{
				"href": fmt.Sprintf("%s/%s", cc.Address, "v3"),
				"meta": map[string]interface{}{
					"version": "3.41.0",
				},
			},
			"uaa": map[string]interface{}{
				"href": cc.Address,
			},
		},
	}
	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

func (cc *CloudController) InfoV2(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"api_version": "2.116.0",
		// "doppler_logging_endpoint": "ws://localhost:8000",
		"authorization_endpoint": fmt.Sprintf("%s/%s", cc.Address, "uaa"),
	}
	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

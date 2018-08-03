package main

import (
	"encoding/json"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (cc *CloudController) SpacesV2(w http.ResponseWriter, r *http.Request) {
	list, err := cc.Client.CoreV1().Namespaces().List(metav1.ListOptions{
		LabelSelector: "cf=enabled",
	})

	if err != nil {
		w.WriteHeader(500)
	}

	spaces := make([]map[string]interface{}, 0)
	for _, n := range list.Items {
		spaces = append(spaces, map[string]interface{}{
			"metadata": map[string]interface{}{
				"guid": n.Labels["space_id"],
			},
			"entity": map[string]interface{}{
				"name": n.Labels["space_name"],
			},
		},
		)
	}

	response := map[string]interface{}{
		"total_results": 1,
		"total_pages":   1,
		"resources":     spaces,
	}

	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

func (cc *CloudController) SpaceSummaryV2(w http.ResponseWriter, r *http.Request) {
	list, err := cc.ServingClient.ServingV1alpha1().Services(cc.Namespace).List(metav1.ListOptions{
		// LabelSelector: "cf=enabled",
	})

	if err != nil {
		w.WriteHeader(500)
	}

	apps := make([]map[string]interface{}, 0)

	for _, n := range list.Items {
		var runningInstances int
		state := "STOPPED"

		if n.Status.IsReady() {
			runningInstances = 1
			state = "STARTED"
		}

		apps = append(apps, map[string]interface{}{
			"routes": []interface{}{
				map[string]interface{}{
					"host": n.GetName(),
					"domain": map[string]interface{}{
						"name": cc.Domain,
					},
				},
			},
			"running_instances": runningInstances,
			"name":              n.GetName(),
			"instances":         runningInstances,
			"disk_quota":        1024,
			"state":             state,
		},
		)
	}

	response := map[string]interface{}{
		"name": cc.Namespace,
		"apps": apps,
	}
	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

package main

import (
	"encoding/json"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (cc *CloudController) OrganizationsV2(w http.ResponseWriter, r *http.Request) {
	list, err := cc.Client.CoreV1().Namespaces().List(metav1.ListOptions{
		LabelSelector: "cf=enabled",
	})

	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(500)
	}

	spaces := make([]map[string]interface{}, 0)
	for _, n := range list.Items {
		spaces = append(spaces, map[string]interface{}{
			"metadata": map[string]interface{}{
				"guid": n.Labels["org_id"],
			},
			"entity": map[string]interface{}{
				"name": n.Labels["org_name"],
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

package main

import (
	"encoding/json"
	"net/http"

	servingclientset "github.com/knative/serving/pkg/client/clientset/versioned"
	"k8s.io/client-go/kubernetes"
)

type CloudController struct {
	Client             *kubernetes.Clientset
	ServingClient      *servingclientset.Clientset
	ServiceAccountName string
	Address            string
	Domain             string
	Repository         string
	Namespace          string
	TempDir            string
}

func (cc *CloudController) Jobs(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"metadata": map[string]interface{}{
			"guid": "job-id",
		},
		"entity": map[string]interface{}{
			"status": "finished",
			"guid":   "job-id",
		},
	}

	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

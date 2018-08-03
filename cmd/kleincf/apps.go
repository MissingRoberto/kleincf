package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jszroberto/kleincf/knative"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (cc *CloudController) UpdateApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	app := map[string]interface{}{
		"metadata": map[string]interface{}{
			"guid": vars["id"],
		},
		"entity": map[string]interface{}{
			"name":              "name",
			"running_instances": 1,
			"instances":         1,
			"disk_quota":        1024,
			"package_state":     "STAGED",
			"state":             "STARTED",
			"stack_guid":        "cflinuxfs2",
		},
	}

	bytes, _ := json.Marshal(app)
	w.Write(bytes)
}

func (cc *CloudController) App(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	route, err := cc.ServingClient.ServingV1alpha1().Routes(cc.Namespace).Get(vars["id"], metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}

	state := "STOPPED"
	var runningInstances int
	if route.Status.IsReady() {
		state = "RUNNING"
		runningInstances = 1
	}

	app := map[string]interface{}{
		"metadata": map[string]interface{}{
			"guid": vars["id"],
		},
		"entity": map[string]interface{}{
			"name":              vars["id"],
			"running_instances": runningInstances,
			"instances":         runningInstances,
			"disk_quota":        1024,
			"package_state":     "STAGED",
			"state":             state,
			"stack_guid":        "cflinuxfs2",
		},
	}

	bytes, _ := json.Marshal(app)
	w.Write(bytes)
}

func (cc *CloudController) CreateApp(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var body map[string]interface{}
	err = json.Unmarshal(bytes, &body)
	if err != nil {
		panic(err)
	}

	name := body["name"].(string)
	image := fmt.Sprintf("%s/%s", cc.Repository, name)
	servicesClient := cc.ServingClient.ServingV1alpha1().Services(cc.Namespace)
	buildSpec := knative.BuildSpec(knative.BuildOptions{
		ServiceAccountName: cc.ServiceAccountName,
		Repository:         image,
		URL:                fmt.Sprintf("%s/v2/apps/%s/bits", cc.Address, name),
	})

	service := knative.ServiceSpec(knative.ServiceOptions{
		Name:               name,
		Namespace:          cc.Namespace,
		ServiceAccountName: cc.ServiceAccountName,
		BuildSpec:          &buildSpec,
		Image:              image,
		Labels: map[string]string{
			"app_id":   name,
			"app_name": name,
			"cf":       "enabled",
		},
	})

	_, err = servicesClient.Create(&service)
	if err != nil {
		panic(err)
	}

	app := map[string]interface{}{
		"metadata": map[string]interface{}{
			"guid": name,
		},
		"entity": map[string]interface{}{
			"name":              name,
			"running_instances": 0,
			"instances":         1,
			"disk_quota":        1024,
			"state":             "STOPPED",
			"stack_guid":        "cflinuxfs2",
		},
	}

	bytes, _ = json.Marshal(app)
	w.Write(bytes)
}

func (cc *CloudController) Apps(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()["q"]

	list, err := cc.ServingClient.ServingV1alpha1().Services(cc.Namespace).List(metav1.ListOptions{
		FieldSelector: "metadata.name=" + strings.Split(queries[0], ":")[1],
		LabelSelector: "cf=enabled",
	})

	if err != nil {
		response := map[string]interface{}{
			"total_results": 0,
			"total_pages":   0,
			"resources":     nil,
		}
		bytes, _ := json.Marshal(response)
		w.Write(bytes)
		return
		// w.WriteHeader(500)
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
			"metadata": map[string]interface{}{
				"guid": n.Labels["app_id"],
			},
			"entity": map[string]interface{}{
				"name":              n.GetName(),
				"running_instances": runningInstances,
				"instances":         runningInstances,
				"disk_quota":        1024,
				"state":             state,
				"stack_guid":        "cflinuxfs2",
			},
		},
		)
	}

	response := map[string]interface{}{
		"total_results": len(apps),
		"total_pages":   1 % (1 + len(apps)),
		"resources":     apps,
	}
	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

func (cc *CloudController) Instances(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	route, err := cc.ServingClient.ServingV1alpha1().Routes(cc.Namespace).Get(vars["id"], metav1.GetOptions{})

	if err != nil {
		fmt.Println(err)
	}

	state := "STARTING"
	if route.Status.IsReady() {
		state = "RUNNING"
	}

	response := map[string]interface{}{
		"0": map[string]interface{}{
			"state": state,
		},
	}
	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

func (cc *CloudController) Stats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	route, err := cc.ServingClient.ServingV1alpha1().Routes(cc.Namespace).Get(vars["id"], metav1.GetOptions{})

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(404)
		return
	}

	state := "STARTING"
	if route.Status.IsReady() {
		state = "RUNNING"
	}

	response := map[string]interface{}{
		"0": map[string]interface{}{
			"state": state,
			"stats": map[string]interface{}{
				"usage": map[string]interface{}{
					"disk": 66392064,
					"mem":  29880320,
					"cpu":  0.13511219703079957,
					"time": "2014-06-19 22:37:58 +0000",
				},
				"name": vars["id"],
				"uris": []string{
					vars["id"] + cc.Domain,
				},
				"host":       "10.0.0.1",
				"port":       61035,
				"uptime":     65007,
				"mem_quota":  536870912,
				"disk_quota": 1073741824,
				"fds_quota":  16384,
			},
		},
	}
	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

func (cc *CloudController) Stack(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"metadata": map[string]interface{}{
			"guid": "cflinuxfs2",
		},
		"entity": map[string]interface{}{
			"name":        "cflinuxfs2",
			"description": "Ubuntu 14.04.2 trusty",
		},
	}

	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

func (cc *CloudController) Stacks(w http.ResponseWriter, r *http.Request) {
	domains := []map[string]interface{}{
		map[string]interface{}{
			"metadata": map[string]interface{}{
				"guid": "cflinuxfs2",
			},
			"entity": map[string]interface{}{
				"name":        "cflinuxfs2",
				"description": "Ubuntu 14.04.2 trusty",
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

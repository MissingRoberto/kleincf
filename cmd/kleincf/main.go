package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"

	servingclientset "github.com/knative/serving/pkg/client/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var err error
	var clientset *kubernetes.Clientset
	var servingClient *servingclientset.Clientset

	if os.Getenv("KUBECONFIG") != "" {
		kubeconfig := filepath.Join(
			os.Getenv("KUBECONFIG"),
		)
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			log.Fatal(err)
		}

		clientset, err = kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
		servingClient, err = servingclientset.NewForConfig(config)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// In-cluster
		config, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
		// creates the clientset
		clientset, err = kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
		servingClient, err = servingclientset.NewForConfig(config)
		if err != nil {
			log.Fatal(err)
		}
	}

	dir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}

	cc := CloudController{
		Client:             clientset,
		ServingClient:      servingClient,
		ServiceAccountName: os.Getenv("REGISTRY_CREDENTIALS"),
		Address:            os.Getenv("SERVICE_ADDRESS"),
		Domain:             os.Getenv("DOMAIN"),
		Repository:         os.Getenv("REPOSITORY"),
		Namespace:          os.Getenv("NAMESPACE"),
		TempDir:            dir,
	}

	bits := Bits{
		TempDir: dir,
	}

	router := mux.NewRouter()
	router.HandleFunc("/", cc.API).Methods("GET")
	router.HandleFunc("/v2/info", cc.InfoV2).Methods("GET")
	router.HandleFunc("/v2/shared_domains", cc.SharedDomains).Methods("GET")
	router.HandleFunc("/v2/shared_domains/{id}", cc.SharedDomain).Methods("GET")
	router.HandleFunc("/v2/organizations", cc.OrganizationsV2).Methods("GET")
	router.HandleFunc("/v2/organizations/{id}/spaces", cc.SpacesV2).Methods("GET")
	router.HandleFunc("/v2/organizations/{id}/private_domains", cc.SharedDomains).Methods("GET")
	router.HandleFunc("/v2/spaces", cc.SpacesV2).Methods("GET")
	router.HandleFunc("/v2/spaces/{id}/summary", cc.SpaceSummaryV2).Methods("GET")
	router.HandleFunc("/v2/jobs/{id}", cc.Jobs).Methods("GET")
	router.HandleFunc("/v2/apps", cc.Apps).Methods("GET")
	router.HandleFunc("/v2/apps", cc.CreateApp).Methods("POST")
	router.HandleFunc("/v2/apps/{id}", cc.App).Methods("GET")
	router.HandleFunc("/v2/apps/{id}", cc.UpdateApp).Methods("PUT")
	router.HandleFunc("/v2/apps/{id}/bits", bits.DownloadBits).Methods("GET")
	router.HandleFunc("/v2/apps/{id}/bits", bits.UploadBits).Methods("PUT")
	router.HandleFunc("/v2/apps/{id}/instances", cc.Instances).Methods("GET")
	router.HandleFunc("/v2/apps/{id}/routes", cc.AppRoutes).Methods("GET")
	router.HandleFunc("/v2/apps/{id}/stacks", cc.Stacks).Methods("GET")
	router.HandleFunc("/v2/apps/{id}/stats", cc.Stats).Methods("GET")
	router.HandleFunc("/v2/stacks/{id}", cc.Stack).Methods("GET")
	router.HandleFunc("/v2/routes", cc.Routes).Methods("GET")
	router.HandleFunc("/v2/routes", cc.Routes).Methods("PUT")
	router.HandleFunc("/v2/routes/{id}/apps/{appid}", cc.AssociateRoute).Methods("PUT")
	router.HandleFunc("/v3", cc.V3).Methods("GET")
	router.HandleFunc("/login", UAALogin).Methods("GET")
	router.HandleFunc("/uaa/login", UAALogin).Methods("GET")
	router.HandleFunc("/uaa/oauth/token", UAALogin).Methods("POST")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))

	router.HandleFunc("/v3/info", cc.InfoV2).Methods("GET")
	router.HandleFunc("/v3/shared_domains", cc.SharedDomains).Methods("GET")
	router.HandleFunc("/v3/shared_domains/{id}", cc.SharedDomain).Methods("GET")
	router.HandleFunc("/v3/organizations", cc.OrganizationsV2).Methods("GET")
	router.HandleFunc("/v3/organizations/{id}/spaces", cc.SpacesV2).Methods("GET")
	router.HandleFunc("/v3/organizations/{id}/private_domains", cc.SharedDomains).Methods("GET")
	router.HandleFunc("/v3/spaces", cc.SpacesV2).Methods("GET")
	router.HandleFunc("/v3/spaces/{id}/summary", cc.SpaceSummaryV2).Methods("GET")
	router.HandleFunc("/v3/jobs/{id}", cc.Jobs).Methods("GET")
	router.HandleFunc("/v3/apps", cc.Apps).Methods("GET")
	router.HandleFunc("/v3/apps", cc.CreateApp).Methods("POST")
	router.HandleFunc("/v3/apps/{id}", cc.App).Methods("GET")
	router.HandleFunc("/v3/apps/{id}", cc.UpdateApp).Methods("PUT")
	router.HandleFunc("/v3/apps/{id}/bits", bits.DownloadBits).Methods("GET")
	router.HandleFunc("/v3/apps/{id}/bits", bits.UploadBits).Methods("PUT")
	router.HandleFunc("/v3/apps/{id}/instances", cc.Instances).Methods("GET")
	router.HandleFunc("/v3/apps/{id}/routes", cc.AppRoutes).Methods("GET")
	router.HandleFunc("/v3/apps/{id}/stacks", cc.Stacks).Methods("GET")
	router.HandleFunc("/v3/apps/{id}/stats", cc.Stats).Methods("GET")
	router.HandleFunc("/v3/stacks/{id}", cc.Stack).Methods("GET")
	router.HandleFunc("/v3/routes", cc.Routes).Methods("GET")
	router.HandleFunc("/v3/routes/{id}/apps/{appid}", cc.AssociateRoute).Methods("PUT")
}

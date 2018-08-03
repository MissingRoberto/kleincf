# KleinCF

Experiment to rewrite basic CloudFoundry functionalities on top of Knative.


## How to install

1. Install Istio and Knative
1. Install kleincf
    ```
    kubectl apply -f release.yml
    ```
1. Add docker credentials
  ```
  kubectl -n cf-default apply -f docker.yml
  kubectl -n cf-system apply -f docker.yml
  ```
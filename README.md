# hyperfoil-operator
This operator installs and configures Hyperfoil Controller.

## Description
This Operator installs, configures and manages Hyperfoil Controller instances on a Kubernetes/Openshift Based environment

## Getting Started

### Prerequisites
- go version v1.20.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/hyperfoil-operator:tag
```

**NOTE:** This image ought to be published in the personal registry you specified. 
And it is required to have access to pull the image from the working environment. 
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/hyperfoil-operator:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin 
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

**NOTE:** Run `make help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

# GO/V3 Docs
# Hyperfoil Operator

This operator installs and configures Hyperfoil Controller. See example resource in [config/samples/_v1alpha2_hyperfoil.yaml](config/samples/_v1alpha2_hyperfoil.yaml).

## Building the operator

```bash
# Define your container repository.
export CONTAINER_REPO_OVERRIDE="quay.io/hyperfoil" # default: quay.io/hyperfoil
# This creates and pushes the image (${CONTAINER_REPO_OVERRIDE}/hyperfoil-operator:0.15.0)
# Build is executed in a builder container.
make docker-build docker-push
# Creates image (${CONTAINER_REPO_OVERRIDE}/hyperfoil-operator-bundle:0.15.0) with ClusterServiceVersion and other resources
make bundle-build
podman push ${CONTAINER_REPO_OVERRIDE}/hyperfoil-operator-bundle:0.15.0
# Creates CatalogSource and Subscription in the cluster
operator-sdk run bundle ${CONTAINER_REPO_OVERRIDE}/hyperfoil-operator-bundle:0.15.0
```  

## Deploy the operator and run Hyperfoil in minikube
        
First start by [installing minikube](https://minikube.sigs.k8s.io/docs/start/). Start the cluster with `minikube start`. (*Suggestion:* use `minikube dashboard` to monitor the cluster)
                       
Build and install Hyperfoil operator (Go 1.19 is required) with `make build install`.

Deploy the example resource in [config/samples/_v1alpha2_hyperfoil.yaml](config/samples/_v1alpha2_hyperfoil.yaml) in the cluster with `make deploy-samples`

Run the operator with `make run`. Once the hyperfoil controller has started the operator can be stopped, as it only reacts to changes. 

To expose the controller to the host network, need to forward port 8090 with `kubectl port-forward --address 0.0.0.0 service/hyperfoil 8090:8090`

You can connect to the Hyperfoil controller using the hyperfoil-cli or the [web-cli](http://127.0.0.1:8090)
                                           
From here on, follow the [Hyperfoil quickstart guide](https://hyperfoil.io/quickstart) for instructions on how to deploy and run a benchmark.

### Undeploy Hyperfoil

Undeploy samples from cluster `make undeploy-samples` (*Note:* this does not require the operator to be running)

Stop the minikube cluster with `minikube stop`. Optionally delete the cluster with `minikube delete --all`

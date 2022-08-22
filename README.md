# log-broadcast
This operator's purpose to explore multiversion APIs at the simplest level. It will expand to exploring operator potential eventually. Right now, it does not do much, only sends the schema of the API to the mentioned pod in the CRD as <AppName>. Pod needs to be listening for it.
For demonstrative purposes it is not necessary to build. More on this below. 
After building, storage version will be v1beta1.

## Demonstrative setup
 Use
`kubectl run <AppName> --image mahamfirdous/basicschemalogger:v0.0.2 --port=8003` for an image to pair with the CRD. Or alternatively user your own listener. Check the pod's logs to see results. Sample CRDs are in config/samples. Remember to specify a containerPort in the manifest regardless of the directives in the used docker image.

## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster. The cluster used for testing was a single node k3s installation on WSL2.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster
1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Provide or provision an image. Either (*Recommended*) use mine: mahamfirdous/log-broadcast:v3.0.6
	Or just
2a. Build and push your image to the location specified by `IMG`:
	
```sh
make docker-build docker-push IMG=<some-registry>/log-broadcast:tag

```
	
3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/log-broadcast:tag
Just use my registry mentioned above
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller to the cluster:

```sh
make undeploy
```
	
	
### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/) 
which provides a reconcile function responsible for synchronizing resources untile the desired state is reached on the cluster 

### Test It Out
1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

	
## References
https://book.kubebuilder.io
https://developer.ibm.com/tutorials/create-a-conversion-webhook-with-the-operator-sdk/
https://sdk.operatorframework.io/docs/best-practices/best-practices/
https://sdk.operatorframework.io/
http://heidloff.net/article/converting-custom-resource-versions-kubernetes-operators/
https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definition-versioning/
https://github.com/kubernetes-sigs/controller-runtime
	
	
## License

Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

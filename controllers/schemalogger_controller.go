/*
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
*/

package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	broadcastv1beta1 "broadcast.logger/schemalogger/api/v1beta1"
)

// SchemaLoggerReconciler reconciles a SchemaLogger object
type SchemaLoggerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=broadcast.broadcast.logger,resources=schemaloggers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=broadcast.broadcast.logger,resources=schemaloggers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=broadcast.broadcast.logger,resources=schemaloggers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SchemaLogger object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *SchemaLoggerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	log.Info("reconciling schemalogger CR")
	// TODO(user): your logic here
	var schemaMessage broadcastv1beta1.SchemaLogger
	if err := r.Client.Get(ctx, req.NamespacedName, &schemaMessage); err != nil {
		log.Error(err, "unable to fetch CR SchemaLogger")
	}

	var podlist corev1.PodList
	var linkfound bool
	if err := r.List(ctx, &podlist); err != nil {
		log.Error(err, "Unable to list the pods")
	} else {
		for _, piece := range podlist.Items {
			if name := piece.GetName(); name == schemaMessage.Spec.AppName {
				log.Info("Found a pod linked to the CR", "name", name, "CR", req.Name)
				log.Info("sending schema log to pod linked to CR, named", "name", name)
				// TODO(user): add checks for image instead of zero indexing to first container
				port := strconv.Itoa(int(piece.Spec.Containers[0].Ports[0].ContainerPort))
				r.SchemaLogBroadcast("http://"+piece.Status.PodIP+":"+port, &schemaMessage.Spec)
				linkfound = true
			}
		}
	}
	schemaMessage.Status.Discovery = linkfound
	if err := r.Status().Update(ctx, &schemaMessage); err != nil {
		log.Error(err, "Unable to update schemaLogger discovery status", "status", linkfound)
		return ctrl.Result{}, err
	}
	log.Info("SchemaLogger Discovery status updated")

	log.Info("SchemaLogger CR reconciliation done")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SchemaLoggerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&broadcastv1beta1.SchemaLogger{}).
		Watches(
			&source.Kind{Type: &corev1.Pod{}},
			handler.EnqueueRequestsFromMapFunc(r.MapPodstoSchemaLogger),
		).
		Complete(r)

}

func (r *SchemaLoggerReconciler) MapPodstoSchemaLogger(o client.Object) []reconcile.Request {
	ctx := context.Background()
	log := log.FromContext(ctx)

	//List of broadcast resources
	request := []reconcile.Request{}
	var resourcelist broadcastv1beta1.SchemaLoggerList
	if err := r.Client.List(context.TODO(), &resourcelist); err != nil {
		log.Error(err, "Unable to list schemalogger resources")
	} else {
		for _, piece := range resourcelist.Items {
			if piece.Spec.AppName == o.GetName() {
				request = append(request, reconcile.Request{
					NamespacedName: types.NamespacedName{Name: piece.Name, Namespace: piece.Namespace},
				})
				log.Info("event issued for pod linked to schemalogger resource ", "name", o.GetName())
			}
		}
	}
	return request
}

func (r *SchemaLoggerReconciler) SchemaLogBroadcast(URL string, schema *broadcastv1beta1.SchemaLoggerSpec) {
	newSchema, prob := json.Marshal(schema)
	if prob != nil {
		fmt.Printf("Could not marshal request -- %v %s", prob, URL)
	}

	newReader := bytes.NewReader(newSchema)
	contents, prob := http.Post(URL+"/log", "application/json", newReader)
	if prob != nil {
		fmt.Printf("Could not send request -- %v %s", prob, URL)
	} else {

		fmt.Printf("Sending schema to %s\n", URL)
		contentsrep, prob := ioutil.ReadAll(contents.Body)
		if prob != nil {
			fmt.Printf("Could not read response -- %v %s", prob, URL)
		}
		fmt.Printf("%s\n", contentsrep)
	}
}

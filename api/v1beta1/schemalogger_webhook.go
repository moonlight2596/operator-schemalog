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

package v1beta1

import (
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var schemaloggerlog = logf.Log.WithName("schemalogger-resource")

func (r *SchemaLogger) SetupWebhookWithManager(mgr ctrl.Manager) error {

	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
//+kubebuilder:webhook:path=/mutate-broadcast-kubebuilder-io-v1beta1-schemalogger,mutating=true,failurePolicy=fail,groups=broadcast.broadcast.logger,resources=schemaloggers,verbs=create;update,versions=v1,name=mbroadcast.logger,sideEffects=None,admissionReviewVersions=v1beta1

//webhook:verbs=create;update;delete,path=/validate-batch-tutorial-kubebuilder-io-v1-cronjob,mutating=false,failurePolicy=fail,groups=broadcast.broadcast.logger,resources=cronjobs,versions=v1,name=vbroadcast.logger,sideEffects=None,admissionReviewVersions=v1
//kubebuilder:

var _ webhook.Defaulter = &SchemaLogger{}

func (r *SchemaLogger) Default() {
	schemaloggerlog.Info("default", "name", r.Name)
	if r.Annotations != nil {
		r.Annotations["free-rain-check"] = "FREErainValueGhazal"
	}
	schemaloggerlog.Info("webhook", "rain-check", r.Name)
}

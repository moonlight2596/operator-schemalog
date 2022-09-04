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
	"strings"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

/*
type Validator interface {
	runtime.Object
	ValidateCreate() error
	ValidateUpdate(old runtime.Object) error
	ValidateDelete() error
}*/

// log is for logging in this package.
var schemaloggerlog = logf.Log.WithName("schemalogger-resource")

func (r *SchemaLogger) SetupWebhookWithManager(mgr ctrl.Manager) error {

	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
//+kubebuilder:webhook:path=/mutate-broadcast-broadcast-logger-v1beta1-schemalogger,mutating=true,failurePolicy=fail,groups=broadcast.broadcast.logger,resources=schemaloggers,verbs=create;update,versions=v1beta1,name=mbroadcast.broadcast.logger,sideEffects=None,admissionReviewVersions=v1beta1

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-broadcast-broadcast-logger-v1beta1-schemalogger,mutating=false,failurePolicy=fail,groups=broadcast.broadcast.logger,resources=cronjobs,versions=v1,name=vbroadcast.broadcast.logger,sideEffects=None,admissionReviewVersions=v1beta1

var _ webhook.Defaulter = &SchemaLogger{}

func (r *SchemaLogger) Default() {
	schemaloggerlog.Info("default", "name", r.Name)
	if r.Annotations != nil {
		r.Annotations["free-rain-check"] = "FREErainValueGhazal"
	}
	schemaloggerlog.Info("webhook", "rain-check", r.Name)
}

var _ webhook.Validator = &SchemaLogger{}

func (r *SchemaLogger) ValidateCreate() error {
	var Err field.Error
	schemaloggerlog.Info("Validate Creation", "name", r.Name)
	if !(strings.Contains(r.Spec.Image, ":")) {
		Err = *field.Invalid(field.NewPath("Spec").Child("Image"), r.Name, "Need to have tag with image")
		return &Err
	}
	return nil
}

func (r *SchemaLogger) ValidateUpdate(old runtime.Object) error {
	return nil
}

func (r *SchemaLogger) ValidateDelete() error {
	return nil
}

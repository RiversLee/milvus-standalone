/*
Copyright 2021.

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
	"context"
	"fmt"
	"github.com/go-logr/logr"
	"helm.sh/helm/v3/pkg/cli"
	"k8s.io/apimachinery/pkg/api/errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	milvusiov1alpha1 "github.com/milvus-io/mivlus-standalone-operator/api/v1alpha1"
)
const (
	MSFinalizerName = "milvusstandalone.milvus.io/finalizer"
)
// MilvusstandaloneReconciler reconciles a Milvusstandalone object
type MilvusstandaloneReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	logger logr.Logger
	helmSettings *cli.EnvSettings
}

//+kubebuilder:rbac:groups=milvus.io.milvus.io,resources=milvusstandalones,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=milvus.io.milvus.io,resources=milvusstandalones/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=milvus.io.milvus.io,resources=milvusstandalones/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Milvusstandalone object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *MilvusstandaloneReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	milvusstandalone := &milvusiov1alpha1.Milvusstandalone{}
	if err :=  r.Get(ctx, req.NamespacedName, milvusstandalone); err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, fmt.Errorf("error get milvus standalone: #{err}")
	}

	//Finalize
	if milvusstandalone.ObjectMeta.DeletionTimestamp.IsZero() {
		if !controllerutil.ContainsFinalizer(milvusstandalone,MSFinalizerName) {
			controllerutil.AddFinalizer(milvusstandalone, MSFinalizerName)
			if err := r.Update(ctx, milvusstandalone);err != nil {
				return ctrl.Result{},nil
			}
		}
	} else {

	}

	//start reconcile
	r.logger.Info("start reconcile")
	old := milvusstandalone.DeepCopy()

	if err := r.SetDefault(ctx,milvusstandalone);err != nil{
		return ctrl.Result{},err
	}

	if !IsEqual(old,milvusstandalone){
		diff,_ := diffObject(old,milvusstandalone)
		r.logger.Info("SetDefault: "+string(diff),"name",old.Name,"namespace",old.Namespace)
		return ctrl.Result{},r.Update(ctx,milvusstandalone)
	}


	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MilvusstandaloneReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&milvusiov1alpha1.Milvusstandalone{}).
		Complete(r)
}

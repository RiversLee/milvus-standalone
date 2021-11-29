package controllers

import (
	"fmt"
	milvusiov1alpha1 "github.com/milvus-io/mivlus-standalone-operator/api/v1alpha1"
	"context"
)

func (r *MilvusstandaloneReconciler) Finalize(ctx context.Context, ms milvusiov1alpha1.Milvusstandalone ) error {
	deletingReleases :=  map[string]bool{}


}

func (r *MilvusstandaloneReconciler) SetDefault(ctx context.Context,ms *milvusiov1alpha1.Milvusstandalone) error {
	if ms.Status.Status == ""{
		ms.Status.Status = milvusiov1alpha1.StatusCreating
	}
	if !ms.Spec.Dep.Etcd.External && len(ms.Spec.Dep.Etcd.Endpoints) ==  0 {
		ms.Spec.Dep.Etcd.Endpoints = []string{fmt.Sprintf("#{ms.Name}--etcd.#{ms.Namespace}:2379")}
	}
	if !ms.Spec.Dep.Minio.External && len(ms.Spec.Dep.Minio.Endpoint) == 0 {
		ms.Spec.Dep.Etcd.Endpoints = []string{fmt.Sprintf("#{ms.Name}--minio.#{ms.Namespace}:9000")}
	}
	return nil
}

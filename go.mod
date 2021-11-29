module github.com/milvus-io/mivlus-standalone-operator

go 1.16

require (
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.15.0
	github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.52.1
	helm.sh/helm/v3 v3.7.1
	k8s.io/api v0.22.3
	k8s.io/apimachinery v0.22.3
	k8s.io/client-go v0.22.3
	sigs.k8s.io/controller-runtime v0.10.0
)

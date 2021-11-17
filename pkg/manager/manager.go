package manager

import (
	milvusiov1alpha1 "github.com/milvus-io/mivlus-standalone-operator/api/v1alpha1"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"time"
)

var (
	scheme = runtime.NewScheme()
	mgrLog = ctrl.Log.WithName("manager")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(milvusiov1alpha1.AddToScheme(scheme))
	utilruntime.Must(monitoringv1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func NewManager(metricsAddr, probeAddr string, enableLeaderElection bool) (ctrl.Manager, error){
	syncPeriod := 1 * time.Minute
	ctrlOptions := ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     metricsAddr,
		Port:                   9443,
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       "2dc33955.milvus.io",
		SyncPeriod: 			&syncPeriod,
	}
	conf := ctrl.GetConfigOrDie()
	mgr, err := ctrl.NewManager(conf, ctrlOptions)
	if err != nil {
		mgrLog.Error(err, "unable to start manager")
		return nil, err
	}

	//+kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		mgrLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		mgrLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	return mgr, nil
}
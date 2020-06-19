package weekly

import (
	"context"

	"github.com/cloudnative-id/community-operator/dispatcher"

	communityv1alpha1 "github.com/cloudnative-id/community-operator/pkg/apis/community/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_weekly")
var dispatcherController dispatcher.Dispatcher

// Create dispatcherController and adds it to the Manager. The Manager will set the dispacherController when the Manager is Started.
func AddDispatcher(disp dispatcher.Dispatcher) {
	dispatcherController = disp
}

// AddManager creates a new Weekly Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func AddManager(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileWeekly{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("weekly-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Weekly
	err = c.Watch(&source.Kind{Type: &communityv1alpha1.Weekly{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileWeekly implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileWeekly{}

// ReconcileWeekly reconciles a Weekly object
type ReconcileWeekly struct {
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Weekly object and makes changes based on the status send
// and what is in the Weekly.Spec
func (r *ReconcileWeekly) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Weekly")

	// Fetch the weekly
	weekly := &communityv1alpha1.Weekly{}

	// Populate weekly
	err := r.client.Get(context.TODO(), request.NamespacedName, weekly)

	if err != nil {
		if errors.IsNotFound(err) {
			reqLogger.Info("Weekly resource not found. Ignoring since object must be deleted.")
			return reconcile.Result{}, nil
		}
		reqLogger.Error(err, "Failed to get Weekly.")
		return reconcile.Result{}, err
	}

	// Send Weekly to dispatcher based on status
	if weekly.Status.Send {
		reqLogger.Info("Weekly resource status send is True. Ignoring send to dispacher.")
		return reconcile.Result{}, nil
	} else {
		reqLogger.Info("Weekly resource status send is False. Send to dispacher.")
		dispatcherController.SendWeekly(weekly.Spec)

		// Update status.Send
		weekly.Status.Send = true
		err := r.client.Status().Update(context.TODO(), weekly)
		if err != nil {
			reqLogger.Error(err, "Failed to update Weekly status")
			return reconcile.Result{}, err
		}
	}

	return reconcile.Result{}, nil
}

package meetup

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

var log = logf.Log.WithName("controller_meetup")
var dispatcherController dispatcher.Dispatcher

// Create dispatcherController and adds it to the Manager. The Manager will set the dispacherController when the Manager is Started.
func AddDispatcher(disp dispatcher.Dispatcher) {
	dispatcherController = disp
}

// AddManager creates a new Meetup Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func AddManager(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileMeetup{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("meetup-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Meetup
	err = c.Watch(&source.Kind{Type: &communityv1alpha1.Meetup{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileMeetup implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileMeetup{}

// ReconcileMeetup reconciles a Meetup object
type ReconcileMeetup struct {
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Meetup object and makes changes based on the status send
// and what is in the Meetup.Spec
func (r *ReconcileMeetup) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Meetup")

	// Fetch the Meetup
	meetup := &communityv1alpha1.Meetup{}

	// Populate Meetup
	err := r.client.Get(context.TODO(), request.NamespacedName, meetup)

	if err != nil {
		if errors.IsNotFound(err) {
			reqLogger.Info("Meetup resource not found. Ignoring since object must be deleted.")
			return reconcile.Result{}, nil
		}
		reqLogger.Error(err, "Failed to get Meetup.")
		return reconcile.Result{}, err
	}

	// Send Meetup to dispatcher based on status
	if meetup.Status.Send {
		reqLogger.Info("Meetup resource status send is True. Ignoring send to dispacher.")
		return reconcile.Result{}, nil
	} else {
		reqLogger.Info("Meetup resource status send is False. Send to dispacher.")
		dispatcherController.SendMeetup(meetup.Spec)

		// Update status.Send
		meetup.Status.Send = true
		err := r.client.Status().Update(context.TODO(), meetup)
		if err != nil {
			reqLogger.Error(err, "Failed to update Meetup status")
			return reconcile.Result{}, err
		}
	}

	return reconcile.Result{}, nil
}

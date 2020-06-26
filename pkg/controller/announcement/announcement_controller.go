package announcement

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

var log = logf.Log.WithName("controller_announcement")
var dispatcherController dispatcher.Dispatcher

// Create dispatcherController and adds it to the Manager. The Manager will set the dispacherController when the Manager is Started.
func AddDispatcher(disp dispatcher.Dispatcher) {
	dispatcherController = disp
}

// Add creates a new Announcement Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func AddManager(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileAnnouncement{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("announcement-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Announcement
	err = c.Watch(&source.Kind{Type: &communityv1alpha1.Announcement{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileAnnouncement implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileAnnouncement{}

// ReconcileAnnouncement reconciles a Announcement object
type ReconcileAnnouncement struct {
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Anncouncement object and makes changes based on the status send
// and what is in the Anncouncement.Spec
func (r *ReconcileAnnouncement) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Announcement")

	// Fetch the Announcement instance
	announcement := &communityv1alpha1.Announcement{}

	// Populate Meetup
	err := r.client.Get(context.TODO(), request.NamespacedName, announcement)

	if err != nil {
		if errors.IsNotFound(err) {
			reqLogger.Info("Announcement resource not found. Ignoring since object must be deleted.")
			return reconcile.Result{}, nil
		}
		reqLogger.Error(err, "Failed to get Anncouncement.")
		return reconcile.Result{}, err
	}

	// Send Meetup to dispatcher based on status
	if announcement.Status.Send {
		reqLogger.Info("Announcement resource status send is True. Ignoring send to dispacher.")
		return reconcile.Result{}, nil
	} else {
		reqLogger.Info("Announcement resource status send is False. Send to dispacher.")
		dispatcherController.SendAnnouncement(announcement.Spec)

		// Update status.Send
		announcement.Status.Send = true
		err := r.client.Status().Update(context.TODO(), announcement)
		if err != nil {
			reqLogger.Error(err, "Failed to update Announcement status")
			return reconcile.Result{}, err
		}
	}

	return reconcile.Result{}, nil
}

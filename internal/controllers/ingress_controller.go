package controllers

import (
	"context"

	"github.com/drumont/connect-haproxy-operator/internal/haproxy"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

type IngressReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *IngressReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx)

	var ingress networkingv1.Ingress
	if err := r.Get(ctx, req.NamespacedName, &ingress); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info(
		"Ingress creation detected",
		"name", ingress.Name,
		"namespace", ingress.Namespace)

	return ctrl.Result{}, nil
}

func (r *IngressReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&networkingv1.Ingress{}).
		WithOptions(controller.Options{}).
		WithEventFilter(predicate.Funcs{
			CreateFunc: func(e event.TypedCreateEvent[client.Object]) bool {
				log := mgr.GetLogger().WithName("ingress-predicate")
				log.Info("Ingress create event", "name", e.Object.GetName())
				haproxy.ReconcileIngress("kubernetes")
				return true
			},
			UpdateFunc: func(e event.TypedUpdateEvent[client.Object]) bool {
				return false
			},
			DeleteFunc: func(e event.TypedDeleteEvent[client.Object]) bool {
				return false
			},
		}).
		Complete(r)
}

package controllers

import (
	"fmt"

	"github.com/enna-systems/jitsi-kubernetes-operator/api/v1alpha1"

	"github.com/presslabs/controller-util/pkg/syncer"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewJVBPodMonitorSyncer(jitsi *v1alpha1.Jitsi, c client.Client) syncer.Interface {
	mon := &monitoringv1.PodMonitor{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-jvb", jitsi.Name),
			Namespace: jitsi.Namespace,
		},
	}

	return syncer.NewObjectSyncer("PodMonitor", jitsi, mon, c, func() error {
		mon.Labels = jitsi.ComponentLabels("jvb")

		mon.Spec.Selector = metav1.LabelSelector{
			MatchLabels: jitsi.ComponentLabels("jvb"),
		}
		mon.Spec.PodMetricsEndpoints = []monitoringv1.PodMetricsEndpoint{
			{
				Port: "metrics",
			},
		}

		return nil
	})

}

func NewJicofoServiceMonitorSyncer(jitsi *v1alpha1.Jitsi, c client.Client) syncer.Interface {
	mon := &monitoringv1.PodMonitor{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-jicofo", jitsi.Name),
			Namespace: jitsi.Namespace,
		},
	}

	return syncer.NewObjectSyncer("PodMonitor", jitsi, mon, c, func() error {
		mon.Labels = jitsi.ComponentLabels("jicofo")

		mon.Spec.Selector = metav1.LabelSelector{
			MatchLabels: jitsi.ComponentLabels("jicofo"),
		}
		mon.Spec.PodMetricsEndpoints = []monitoringv1.PodMetricsEndpoint{
			{
				Port: "metrics",
			},
		}

		return nil
	})

}

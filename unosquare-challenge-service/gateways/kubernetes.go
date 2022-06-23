package gateways

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"unosquare-challenge/gateways/models"
)

type KubernetesGateway interface {
	GetServices() ([]*models.Service, error)
}

type kubernetesGateway struct {
}

func NewKubernetesGateway() KubernetesGateway {
	return &kubernetesGateway{}
}

func (kg *kubernetesGateway) GetServices() ([]*models.Service, error) {

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	response, err := clientset.CoreV1().Services("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	services := make([]*models.Service, 0)
	for _, svc := range response.Items {
		services = append(services, &models.Service{
			Metadata: models.Metadata{
				Name:              svc.Name,
				Namespase:         svc.Namespace,
				Uid:               string(svc.UID),
				ResourceVersion:   svc.ResourceVersion,
				CreationTimestamp: svc.CreationTimestamp.Time,
			},
			Spec: models.Spec{
				ClusterIP:       svc.Spec.ClusterIP,
				ClusterIPs:      svc.Spec.ClusterIPs,
				Type:            string(svc.Spec.Type),
				SessionAffinity: svc.Spec.SessionAffinityConfig.String(),
				//IpFamilies: svc.Spec.IPFamilies,
				InternalTrafficPolicy: string(*svc.Spec.InternalTrafficPolicy),
			},
		})
	}

	return services, nil
}

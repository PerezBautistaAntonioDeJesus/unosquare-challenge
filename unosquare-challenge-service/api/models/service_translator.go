package models

import (
	gatewayModels "unosquare-challenge/gateways/models"
)

func ServiceGatewayToServiceAPI(services []*gatewayModels.Service) []*Service {
	servicesApi := make([]*Service, 0, len(services))
	for _, svc := range services {
		servicesApi = append(servicesApi, &Service{
			Metadata: Metadata{
				Name:              svc.Metadata.Name,
				Namespace:         svc.Metadata.Namespase,
				Uid:               svc.Metadata.Uid,
				ResourceVersion:   svc.Metadata.ResourceVersion,
				CreationTimestamp: svc.Metadata.CreationTimestamp,
			},
			Spec: Spec{
				ClusterIP:             svc.Spec.ClusterIP,
				ClusterIPs:            svc.Spec.ClusterIPs,
				Type:                  svc.Spec.Type,
				SessionAffinity:       svc.Spec.SessionAffinity,
				InternalTrafficPolicy: svc.Spec.InternalTrafficPolicy,
			},
		})
	}
	return servicesApi
}

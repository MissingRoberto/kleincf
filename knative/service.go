package knative

import (
	buildv1alpha1 "github.com/knative/build/pkg/apis/build/v1alpha1"
	"github.com/knative/serving/pkg/apis/serving/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ServiceOptions struct {
	Name               string
	Namespace          string
	ServiceAccountName string
	BuildSpec          *buildv1alpha1.BuildSpec
	Image              string
	Labels             map[string]string
}

func ServiceSpec(opts ServiceOptions) v1alpha1.Service {
	service := v1alpha1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      opts.Name,
			Namespace: opts.Namespace,
			Labels:    opts.Labels,
		},

		Spec: v1alpha1.ServiceSpec{
			RunLatest: &v1alpha1.RunLatestType{
				Configuration: v1alpha1.ConfigurationSpec{
					Build: opts.BuildSpec,
					RevisionTemplate: v1alpha1.RevisionTemplateSpec{
						Spec: v1alpha1.RevisionSpec{
							ServiceAccountName: opts.ServiceAccountName,
							Container: corev1.Container{
								Image: opts.Image,
							},
						},
					},
				},
			},
		},
	}

	return service
}

package knative

import (
	buildv1alpha1 "github.com/knative/build/pkg/apis/build/v1alpha1"
)

type BuildOptions struct {
	ServiceAccountName string
	Repository         string
	URL                string
}

func BuildSpec(opts BuildOptions) buildv1alpha1.BuildSpec {
	return buildv1alpha1.BuildSpec{
		ServiceAccountName: opts.ServiceAccountName,
		Template: &buildv1alpha1.TemplateInstantiationSpec{
			Name: "buildpack-bits",
			Arguments: []buildv1alpha1.ArgumentSpec{
				buildv1alpha1.ArgumentSpec{
					Name:  "URL",
					Value: opts.URL,
					// Value: "http://localhost:8000/v2/apps/myapp/bits",
					// Value: "https://github.com/cloudfoundry-samples/dotnet-core-hello-world/archive/master.zip",
				},
				buildv1alpha1.ArgumentSpec{
					Name:  "IMAGE",
					Value: opts.Repository,
				},
			},
		},
	}
}

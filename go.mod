module github.com/23technologies/gardener-extension-provider-hcloud

go 1.16

require (
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/ahmetb/gen-crd-api-reference-docs v0.2.0
	github.com/coreos/go-systemd/v22 v22.1.0
	github.com/gardener/etcd-druid v0.5.0
	github.com/gardener/gardener v1.21.0
	github.com/gardener/machine-controller-manager v0.35.0
	github.com/go-logr/logr v0.3.0
	github.com/gobuffalo/packr/v2 v2.8.1
	github.com/golang/mock v1.5.0
	github.com/hetznercloud/hcloud-go v1.25.0
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/onsi/ginkgo v1.14.2
	github.com/onsi/gomega v1.10.5
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	k8s.io/api v0.20.7
	k8s.io/apiextensions-apiserver v0.20.6
	k8s.io/apimachinery v0.20.7
	k8s.io/apiserver v0.20.7
	k8s.io/autoscaler v0.0.0-20190805135949-100e91ba756e
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/code-generator v0.20.7
	k8s.io/component-base v0.20.7
	k8s.io/kubelet v0.20.6
	sigs.k8s.io/controller-runtime v0.8.3
)

replace (
	github.com/emicklei/go-restful => github.com/emicklei/go-restful v2.9.5+incompatible // keep this value in sync with k8s.io/apiserver
	github.com/gardener/gardener => github.com/gardener/gardener v1.23.2
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.4.1
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v1.7.1 // keep this value in sync with sigs.k8s.io/controller-runtime
	google.golang.org/grpc => google.golang.org/grpc v1.27.0 // keep this value in sync with k8s.io/apiserver
	k8s.io/api => k8s.io/api v0.20.7
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.20.7
	k8s.io/apimachinery => k8s.io/apimachinery v0.20.7
	k8s.io/apiserver => k8s.io/apiserver v0.20.7
	k8s.io/client-go => k8s.io/client-go v0.20.7
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.20.7
	k8s.io/code-generator => k8s.io/code-generator v0.20.7
	k8s.io/component-base => k8s.io/component-base v0.20.7
	k8s.io/helm => k8s.io/helm v2.13.1+incompatible
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.20.7
	sigs.k8s.io/controller-runtime => github.com/gardener/controller-runtime v0.8.3-gardener.1
)

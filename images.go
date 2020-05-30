package main

import (
	"strings"
)

// sudo docker images | awk '{print $1":"$2}' | sort -u | uniq
var IMAGES = `
ceph/ceph:v14.2.8
quay.io/cephcsi/cephcsi:v1.2.2
quay.io/k8scsi/csi-attacher:v1.2.0
quay.io/k8scsi/csi-node-driver-registrar:v1.2.0
quay.io/k8scsi/csi-provisioner:v1.4.0
quay.io/k8scsi/csi-snapshotter:v1.2.2
rook/ceph:master TO rook-ceph:master
rook/ceph:v1.2.7 TO rook-ceph:v1.2.7

gcr.io/google_containers/volume-nfs:0.8
ocdr/mysql-operator:0.3.9
percona:ps-5.7.26
docker.io/istio/proxy_init:1.1.6
codeskyblue/gohttpserver
gcr.io/knative-releases/knative.dev/serving/cmd/activator:v0.8.0
gcr.io/knative-releases/knative.dev/serving/cmd/autoscaler-hpa:v0.8.0
gcr.io/knative-releases/knative.dev/serving/cmd/autoscaler:v0.8.0
gcr.io/knative-releases/knative.dev/serving/cmd/controller:v0.8.0
gcr.io/knative-releases/knative.dev/serving/cmd/networking/istio:v0.8.0
gcr.io/knative-releases/knative.dev/serving/cmd/queue:v0.8.0
gcr.io/knative-releases/knative.dev/serving/cmd/webhook:v0.8.0
gcr.io/kubeflow-images-public/profile-controller:v20191205-v0.7.1
gcr.io/kubeflow-images-public/kfam:v20191014-v0.7.0-rc.0-10-gdf3c9366-e3b0c4
gcr.io/kubeflow-images-public/katib/v1alpha3/katib-manager:v0.7.0
gcr.io/kubeflow-images-public/katib/v1alpha3/katib-controller:v0.7.0
gcr.io/kubeflow-images-public/katib/v1alpha3/katib-ui:v0.7.0
gcr.io/ml-pipeline/api-server:0.1.31
gcr.io/kubeflow-images-public/jupyter-web-app:v0.5.0
gcr.io/ml-pipeline/persistenceagent:0.1.31
gcr.io/ml-pipeline/scheduledworkflow:0.1.31
gcr.io/kubeflow-images-public/notebook-controller:v20190614-v0-160-g386f2749-e3b0c4
gcr.io/kubeflow-images-public/profile-controller:v20190619-v0-219-gbd3daa8c-dirty-1ced0e
gcr.io/kubeflow-images-public/pytorch-operator:v0.6.0-18-g5e36a57
gcr.io/kubeflow-images-public/tf_operator:kubeflow-tf-operator-postsubmit-v1-5adee6f-6109-a25c
gcr.io/kubeflow-images-public/profile-controller:v20190619-v0-219-gbd3daa8c-dirty-1ced0e
gcr.io/kubeflow-images-public/kfam:v20190612-v0-170-ga06cdb79-dirty-a33ee4

ocdr/d3-katib-controller:v0.7.0
ocdr/d3-kfplapiserver:2.0.3
ocdr/dkube-d3api:2.0.3
ocdr/dkube-d3auth:2.0.3
ocdr/dkube-d3downloader:2.0.3
ocdr/dkube-d3ext:2.0.3
ocdr/dkube-d3inf:2.0.3
ocdr/dkube-d3installer:2.0.3
ocdr/dkube-d3stashcontroller:2.0.3
ocdr/dkube-d3storagexporter:2.0.3
ocdr/dkube-d3watcher:2.0.3
ocdr/dkube-dfabproxy:2.0.3
ocdr/dkube-docs:2.0.3
ocdr/dkube-inf-watcher:2.0.3
ocdr/dkube-uiserver:2.0.3
ocdr/dkubeadm:2.0.3
ocdr/dkubegc:latest
ocdr/fluentd-kubernetes-daemonset:v1.7-debian-s3-1
ocdr/kfserving-controller:0.2.2
ocdr/ml-pipeline-ui:2.0.3
ocdr/workflow-controller:dkube

alpine:latest
argoproj/argoexec:v2.3.0
argoproj/argoui:v2.3.0
argoproj/workflow-controller:v2.3.0
calico/kube-controllers:v3.7.3
ceph/ceph:v14.2.5
codeskyblue/gohttpserver:latest
coredns/coredns:1.6.0
gcr.io/etcd-development/etcd:latest
gcr.io/google-containers/addon-resizer-amd64:2.1
gcr.io/google-containers/cluster-proportional-autoscaler-amd64:1.6.0
gcr.io/google-containers/k8s-dns-node-cache:1.15.8
gcr.io/google-containers/kube-apiserver:v1.15.3
gcr.io/google-containers/kube-controller-manager:v1.15.3
gcr.io/google-containers/kube-proxy:v1.15.3
gcr.io/google-containers/kube-scheduler:v1.15.3
gcr.io/google-containers/pause:3.1
gcr.io/google_containers/kube-state-metrics:v1.2.0
gcr.io/google_containers/kubernetes-dashboard-amd64:v1.10.1
gcr.io/google_containers/pause-amd64:3.1
gcr.io/google_containers/spartakus-amd64:v1.1.0
gcr.io/google_containers/volume-nfs:0.8
gcr.io/kfserving/kfserving-controller:0.2.2
gcr.io/kfserving/storage-initializer:0.2.2
gcr.io/kubebuilder/kube-rbac-proxy:v0.4.0
gcr.io/kubeflow-images-public/admission-webhook:v20190520-v0-139-gcee39dbc-dirty-0d8f4c
gcr.io/kubeflow-images-public/centraldashboard:<none>
gcr.io/kubeflow-images-public/centraldashboard:v20190823-v0.6.0-rc.0-69-gcb7dab59
gcr.io/kubeflow-images-public/ingress-setup:latest
gcr.io/kubeflow-images-public/jupyter-web-app:9419d4d
gcr.io/kubeflow-images-public/katib/katib-ui:v0.1.2-alpha-157-g3d4cd04
gcr.io/kubeflow-images-public/katib/studyjob-controller:v0.1.2-alpha-157-g3d4cd04
gcr.io/kubeflow-images-public/katib/suggestion-grid:v0.1.2-alpha-157-g3d4cd04
gcr.io/kubeflow-images-public/katib/suggestion-hyperband:v0.1.2-alpha-157-g3d4cd04
gcr.io/kubeflow-images-public/katib/suggestion-random:v0.1.2-alpha-157-g3d4cd04
gcr.io/kubeflow-images-public/katib/v1alpha3/katib-manager:v0.7.0
gcr.io/kubeflow-images-public/katib/v1alpha3/katib-ui:v0.7.0
gcr.io/kubeflow-images-public/katib/vizier-core-rest:v0.1.2-alpha-157-g3d4cd04
gcr.io/kubeflow-images-public/katib/vizier-core:v0.1.2-alpha-157-g3d4cd04
gcr.io/kubeflow-images-public/kfam:<none>
gcr.io/kubeflow-images-public/kubernetes-sigs/application:1.0-beta
gcr.io/kubeflow-images-public/metadata-frontend:v0.1.8
gcr.io/kubeflow-images-public/metadata:v0.1.11
gcr.io/kubeflow-images-public/metadata:v0.1.8
gcr.io/kubeflow-images-public/notebook-controller:<none>
gcr.io/kubeflow-images-public/profile-controller:<none>
gcr.io/kubeflow-images-public/pytorch-operator:v0.7.0
gcr.io/kubeflow-images-public/tf_operator:v0.7.0
gcr.io/kubernetes-helm/tiller:v2.16.1
gcr.io/kubernetes-helm/tiller:v2.16.3
gcr.io/ml-pipeline/envoy:metadata-grpc
gcr.io/ml-pipeline/frontend:0.1.23
gcr.io/ml-pipeline/frontend:0.1.31
gcr.io/ml-pipeline/persistenceagent:0.1.23
gcr.io/ml-pipeline/persistenceagent:0.1.31
gcr.io/ml-pipeline/scheduledworkflow:0.1.23
gcr.io/ml-pipeline/scheduledworkflow:0.1.31
gcr.io/ml-pipeline/viewer-crd-controller:0.1.23
gcr.io/ml-pipeline/viewer-crd-controller:0.1.31
gcr.io/ml-pipeline/visualization-server:0.1.27
gcr.io/tfx-oss-public/ml_metadata_store_server:0.15.1
google/cadvisor:latest
grafana/grafana:5.0.0
grafana/grafana:6.0.2
istio/citadel:1.1.6
istio/galley:1.1.6
istio/kubectl:1.1.6
istio/mixer:1.1.6
istio/pilot:1.1.6
istio/proxy_init:1.1.6
istio/proxyv2:1.1.6
istio/sidecar_injector:1.1.6
jaegertracing/all-in-one:1.9
k8s.gcr.io/etcd-amd64:3.1.12
k8s.gcr.io/k8s-dns-node-cache:1.15.4
kiali/kiali:v0.16
metacontroller/metacontroller:v0.3.0
minio/minio:RELEASE.2018-02-09T22-40-05Z
minio/minio:RELEASE.2018-12-13T02-04-19Z
mysql:5.6
mysql:8
mysql:8.0.3
nginx:1.15
nginx:1.17
nginx:<none>
nvidia/k8s-device-plugin:1.11
percona:<none>
prom/mysqld-exporter:v0.11.0
prom/prometheus:v2.12.0
prom/prometheus:v2.3.1
quay.io/cephcsi/cephcsi:v1.2.2
quay.io/coreos/etcd:v3.3.10
quay.io/coreos/flannel-cni:v0.3.0
quay.io/coreos/flannel:v0.11.0
quay.io/coreos/grafana-watcher:v0.0.8
quay.io/datawire/ambassador:0.53.1
quay.io/k8scsi/csi-attacher:v1.2.0
quay.io/k8scsi/csi-node-driver-registrar:v1.2.0
quay.io/k8scsi/csi-provisioner:v1.4.0
quay.io/k8scsi/csi-snapshotter:v1.2.2
quay.io/presslabs/mysql-operator-orchestrator:0.3.9
quay.io/presslabs/mysql-operator-sidecar:0.3.9
quay.io/presslabs/mysql-operator:0.3.9
quay.io/prometheus/node-exporter:v0.15.2
rook/ceph:master
rook/ceph:v1.2.4
rook/ceph:v1.2.5
seldonio/seldon-core-operator:0.4.1
tensorflow/serving:<none>
tensorflow/tensorflow:1.8.0
`

func getImages() []string {
	images := strings.Split(IMAGES, "\n")
	images = unique(images)

	return images
}

func unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, present := keys[entry]; !present {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

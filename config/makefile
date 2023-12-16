# doc of install script

appNamespace :=default

httpbinDelete:
	kubectl delete -f https://raw.githubusercontent.com/istio/istio/release-1.7/samples/httpbin/httpbin.yaml -n ${appNamespace}

httpbincreate:
	kubectl create -f https://raw.githubusercontent.com/istio/istio/release-1.7/samples/httpbin/httpbin.yaml -n ${appNamespace}

sleeperCreate:
	kubectl create -f https://raw.githubusercontent.com/istio/istio/release-1.7/samples/sleep/sleep.yaml -n ${appNamespace}

ingressControllerCreate:
	kubectl create -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.8.2/deploy/static/provider/cloud/deploy.yaml

httpbinIngressInstall:
	kubectl apply -f ./httpbin-ingress.yaml

fooServiceCreate:
	kubectl create -f ./foo-service-deployment.yaml


imageBuild:
	docker  build -t mk100120/ingress-manager:1.0.0 ../


ingressManagerDryRun:
	kubectl create deployment ingress-manager --image=mk100120/ingress-manager:1.0.0 --dry-run=client -o yaml > ingress-manager-dry-run.yaml
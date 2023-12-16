package main

import (
	"github.com/ekko100120/informer-demo/pkg"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {

	// 1. get client config
	// 2. create clientset
	// 3. create informer factory
	// 4. create controller and add eventHandler
	// 5. start informer factory
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		clusterConfig, err := rest.InClusterConfig()
		if err != nil {
			log.Fatal("cannot get config")
		}
		config = clusterConfig
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("cannot get clientset")
	}
	factory := informers.NewSharedInformerFactory(clientset, 0)
	serviceInformer := factory.Core().V1().Services()
	ingressesInformer := factory.Networking().V1().Ingresses()
	controller := pkg.NewController(clientset, serviceInformer, ingressesInformer)
	stopChan := make(chan struct{})
	factory.Start(stopChan)
	factory.WaitForCacheSync(stopChan)
	controller.Run(stopChan)

}

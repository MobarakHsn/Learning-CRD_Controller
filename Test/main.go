package main

import (
	"bufio"
	"flag"
	"fmt"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
	"time"
)

func int32Ptr(i int32) *int32 {
	return &i
}

func prompt() {
	fmt.Println("-> Press Return key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()
}

func addPod(newObj interface{}) {

	// Here we can either call some method to send a notification or to make it simple simply print some message
	fmt.Println("Pod is added")
}

func deletePod(obj interface{}) {

	fmt.Println("inside delete function")

}
func updatePod(old, new interface{}) {

	fmt.Println("inside update function")

}

func main() {
	var kuberconfig *string
	if home := homedir.HomeDir(); home != "" {
		kuberconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kuberconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kuberconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	ch := make(chan struct{})
	defer close(ch)
	informerfactory := informers.NewSharedInformerFactory(clientset, 30*time.Second)
	c := newController(clientset, informerfactory.Apps().V1().Deployments())
	informerfactory.Start(ch)
	c.run(ch)
}

package helper

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var k8sClient client.Client

const (
	timeout  = time.Minute * 5
	interval = time.Millisecond * 250
)

// checks if a given resource is present in the cluster
// continouslly polls until it returns nil or a timeout occurs
func CheckIfPresent(ns types.NamespacedName, obj client.Object) {
	Eventually(func() error {
		err := k8sClient.Get(context.TODO(), ns, obj)
		if err != nil {
			return err
		}
		return nil
	}, timeout, interval).ShouldNot(HaveOccurred())
}

// checks if a given resource is deleted
// continouslly polls until the object is deleted or a timeout occurs
func CheckIfDeleted(ns types.NamespacedName, obj client.Object) {
	Eventually(func() error {
		err := k8sClient.Get(context.TODO(), ns, obj)
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}, timeout, interval).ShouldNot(HaveOccurred())
}

// copyKubeConfigFile copies default kubeconfig file into current temporary context config file
func copyKubeConfigFile(kubeConfigFile, tempConfigFile string) {
	info, err := os.Stat(kubeConfigFile)
	Expect(err).NotTo(HaveOccurred())
	err = copyFile(kubeConfigFile, tempConfigFile, info)
	Expect(err).NotTo(HaveOccurred())
	os.Setenv("KUBECONFIG", tempConfigFile)
	fmt.Fprintf(GinkgoWriter, "Setting KUBECONFIG=%s\n", tempConfigFile)
}

// LocalKubeconfigSet sets the KUBECONFIG to the temporary config file
func LocalKubeconfigSet(context string) {
	originalKubeCfg := os.Getenv("KUBECONFIG")
	if originalKubeCfg == "" {
		homeDir := GetUserHomeDir()
		originalKubeCfg = filepath.Join(homeDir, ".kube", "config")
	}
	copyKubeConfigFile(originalKubeCfg, filepath.Join(context, "config"))
}

func copyFile(src, dest string, info fs.FileInfo) error {
	bytesRead, err := os.ReadFile(src)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if err := os.WriteFile(dest, bytesRead, info.Mode()); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// GetUserHomeDir gets the user home directory
func GetUserHomeDir() string {
	homeDir, err := os.UserHomeDir()
	Expect(err).NotTo(HaveOccurred())
	return homeDir
}

func ListPods() {

}
func CreateNewContext() string {
	directory, err := os.MkdirTemp("", "")
	Expect(err).NotTo(HaveOccurred())
	fmt.Fprintf(GinkgoWriter, "Created dir: %s\n", directory)
	return directory
}

func InitializeKubeconfg() *kubernetes.Clientset {
	Rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(Rules, nil)
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		panic(err)
	}
	clientset := kubernetes.NewForConfigOrDie(config)
	return clientset
}

// Return string of length passed as parameter
func RandomString(length int) string {
	rand.Seed(time.Now().Unix())
	ran_str := make([]byte, length)

	// Generating Random string
	for i := 0; i < length; i++ {
		ran_str[i] = byte(65 + rand.Intn(25))
	}

	str := string(ran_str)
	return str
}

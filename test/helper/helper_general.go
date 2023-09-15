package helper

import (
	"context"
	"time"

	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
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

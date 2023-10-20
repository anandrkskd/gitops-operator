package helper

import (
	"fmt"

	argoapp "github.com/argoproj-labs/argocd-operator/api/v1beta1"
)

func CreateArgoApplication() {
	argoAppication := &argoapp.ArgoCD{}
	fmt.Println(argoAppication)
}

func main() {
	CreateArgoApplication()
}

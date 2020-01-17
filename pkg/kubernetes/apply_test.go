package kubernetes_test

import (
	"strings"
	"testing"

	"github.com/thought-machine/dracon/pkg/kubernetes"
	"github.com/thought-machine/dracon/plz-out/go/src/github.com/stretchr/testify/assert"
)

func TestGetCmd(t *testing.T) {
	var tests = []struct {
		desc   string
		inOpts kubernetes.KubectlOpts
		outCmd string
	}{
		{"none", kubernetes.KubectlOpts{}, "kubectl apply -f -"},
		{"namespace", kubernetes.KubectlOpts{Namespace: "default"}, `kubectl apply -f - --namespace=default`},
		{"context", kubernetes.KubectlOpts{Context: "minikube"}, `kubectl apply -f - --context=minikube`},
		{"namespace&context", kubernetes.KubectlOpts{Context: "minikube", Namespace: "default"}, `kubectl apply -f - --context=minikube --namespace=default`},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			outCmd := strings.Join(kubernetes.GetCmd(&tt.inOpts), " ")
			assert.Equal(t, tt.outCmd, outCmd)
		})
	}
}

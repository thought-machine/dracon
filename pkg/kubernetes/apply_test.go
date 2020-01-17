package kubernetes

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCmd(t *testing.T) {
	var tests = []struct {
		desc   string
		inOpts KubectlOpts
		outCmd string
	}{
		{"none", KubectlOpts{}, "kubectl apply -f -"},
		{"namespace", KubectlOpts{Namespace: "default"}, `kubectl apply -f - --namespace=default`},
		{"context", KubectlOpts{Context: "minikube"}, `kubectl apply -f - --context=minikube`},
		{"namespace&context", KubectlOpts{Context: "minikube", Namespace: "default"}, `kubectl apply -f - --context=minikube --namespace=default`},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			outCmd := strings.Join(GetCmd(&tt.inOpts), " ")
			assert.Equal(t, tt.outCmd, outCmd)
		})
	}
}

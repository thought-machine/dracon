package kubernetes

import (
	"fmt"
	"io"
	"log"
	"os/exec"

	"github.com/pkg/errors"
)

// KubectlOpts represents options/flags to pass to kubectl
type KubectlOpts struct {
	Context   string
	Namespace string
}

// Apply config using kubectl
func Apply(resources string, opts *KubectlOpts) error {
	shCmd := GetCmd(opts)
	cmd := exec.Command(shCmd[0], shCmd[1:]...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return errors.Wrap(err, "could not create stdin pipe")
	}
	go func() {
		defer stdin.Close()
		_, err := io.WriteString(stdin, resources)
		if err != nil {
			log.Fatal(err)
		}
	}()

	output, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, resources)
	}
	if !cmd.ProcessState.Success() {
		return errors.Wrap(err, string(output))
	}
	log.Println(string(output))
	return nil
}

// GetCmd returns the kubectl command
func GetCmd(opts *KubectlOpts) []string {
	cmd := []string{"kubectl", "apply", "-f", "-"}

	if opts.Context != "" {
		cmd = append(cmd, fmt.Sprintf(`--context=%s`, opts.Context))
	}
	if opts.Namespace != "" {
		cmd = append(cmd, fmt.Sprintf(`--namespace=%s`, opts.Namespace))
	}

	return cmd
}

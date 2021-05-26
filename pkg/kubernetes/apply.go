package kubernetes

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

// KubectlOpts represents options/flags to pass to kubectl
type KubectlOpts struct {
	Context   string
	Namespace string
}

// Apply config using kubectl
func Apply(resources string, opts *KubectlOpts) error {
	shCmd := GetCmd(opts, "apply")
	cmd := exec.Command(shCmd[0], shCmd[1:]...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("could not create stdin pipe: %w", err)
	}
	go func() {
		defer stdin.Close()
		_, err := io.WriteString(stdin, resources)
		if err != nil {
			log.Fatal(err)
		}
	}()

	output, err := cmd.CombinedOutput()
	if err != nil || !cmd.ProcessState.Success() {
		return fmt.Errorf("%s\n%s:%w", resources, output, err)
	}
	log.Printf("%s\n", output)
	return nil
}

// GetCmd returns the kubectl command
func GetCmd(opts *KubectlOpts, arg string) []string {
	cmd := []string{"kubectl", arg, "-f", "-"}

	if opts.Context != "" {
		cmd = append(cmd, fmt.Sprintf(`--context=%s`, opts.Context))
	}
	if opts.Namespace != "" {
		cmd = append(cmd, fmt.Sprintf(`--namespace=%s`, opts.Namespace))
	}

	return cmd
}

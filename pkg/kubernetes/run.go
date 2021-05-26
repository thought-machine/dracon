package kubernetes

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

// Create config using kubectl
func Create(resources string, opts *KubectlOpts) error {
	shCmd := GetCmd(opts, "create")
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

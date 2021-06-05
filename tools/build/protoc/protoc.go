package protoc

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

var (
	includes = []string{
		"-Ibackend/core/proto",
		"-Ithird_party/googleapis",
		"-Ithird_party/envoyproxy",
	}
	outs = []string{
		"--go_out=%s",
		"--go-grpc_out=%s",
		`--validate_out="lang=go:%s"`,
		`--openapiv2_out=%s`,
		`--grpc-gateway_out=%s`,
	}
	opts = []string{
		`--openapiv2_opt`,
		`logtostderr=true`,
		`--grpc-gateway_opt`,
		`logtostderr=true`,
	}
)

func buildArgs(tempDir string, input string) []string {
	args := []string{"protoc"}
	args = append(args, includes...)
	for i := 0; i < len(outs); i++ {
		args = append(args, fmt.Sprintf(outs[i], tempDir))
	}
	args = append(args, opts...)
	args = append(args, input)
	return args
}

func Action(c *cli.Context) error {
	input := "backend/core/proto/*.proto"
	output := "tmp"
	args := buildArgs(output, input)
	command := strings.Join(args, " ")
	fmt.Println(command)
	cmd := exec.Command("/bin/sh", "-c", command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(out))
	}
	err = filepath.Walk(output,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			command := fmt.Sprintf(`cp %s %s`, path, "backend/core/api/")
			cmd := exec.Command("/bin/sh", "-c", command)
			out, err := cmd.CombinedOutput()
			if err != nil {
				return errors.New(string(out))
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	cmd = exec.Command("/bin/sh", "-c", "rm -rf /tmp/*")
	out, err = cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(out))
	}
	return nil
}

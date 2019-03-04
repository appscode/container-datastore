package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/kubedb/cli/pkg/cmds"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"kmodules.xyz/client-go/logs"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	logs.InitLogs()
	defer logs.FlushLogs()

	cmd := cmds.NewKubeDBCommand(os.Stdin, os.Stdout, os.Stderr)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

package plugins

import (
	"os"
	"path/filepath"

	"github.com/dickeyxxx/gonpm/cli"
)

func init() {
	os.Setenv("NODE_PATH", filepath.Join(appDir, "lib", "node_modules"))
	os.Setenv("NPM_CONFIG_GLOBAL", "true")
	os.Setenv("NPM_CONFIG_PREFIX", appDir)
}

var Topic = &cli.Topic{
	Name: "plugins",
	Run:  Run,
	Help: Help,
}

func Run(command string, args ...string) {
	switch command {
	case "install":
		if len(args) != 1 {
			Help(command, args...)
			cli.Exit(2)
		}
		install(args[0])
	case "list":
	case "":
		list()
	default:
		Help(command, args...)
	}
}

func Help(command string, args ...string) {
	cli.Stderrln("plugins help for:", command, args)
}
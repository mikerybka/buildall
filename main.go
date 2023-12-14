package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mikerybka/infra"
	"github.com/mikerybka/util"
)

func main() {
	cmds := []string{}
	f, err := os.Open(filepath.Join(util.HomeDir(), "src/github.com/mikerybka/cmds.json"))
	if err != nil {
		log.Fatalf("no cmd.json file")
	}
	err = json.NewDecoder(f).Decode(&cmds)
	if err != nil {
		log.Fatalf("unable to parse cmds.json file: %s", err)
	}
	for _, pkg := range cmds {
		outDir := filepath.Join(util.HomeDir(), "src/github.com/mikerybka/builds.mikerybka.com/public/main")
		err := infra.Build(pkg, outDir)
		if err != nil {
			log.Fatalf("Error building %s: %s", pkg, err)
		}
	}
	cmd := exec.Command("git", "add", "--all")
	cmd.Dir = filepath.Join(util.HomeDir(), "src/github.com/mikerybka/builds.mikerybka.com")
	err = util.Run(cmd)
	if err != nil {
		panic(err)
	}
	cmd = exec.Command("git", "commit", "-m", "wip")
	cmd.Dir = filepath.Join(util.HomeDir(), "src/github.com/mikerybka/builds.mikerybka.com")
	err = util.Run(cmd)
	if err != nil {
		return
	}
	cmd = exec.Command("git", "push")
	cmd.Dir = filepath.Join(util.HomeDir(), "src/github.com/mikerybka/builds.mikerybka.com")
	err = util.Run(cmd)
	if err != nil {
		panic(err)
	}
}

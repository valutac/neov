package main

import (
	"fmt"
	"log"
	"os/user"

	"github.com/Valutac/neov/authentication"
	"github.com/Valutac/neov/cfg"
	"github.com/Valutac/neov/compute"
)

func main() {
	fmt.Println("neov: neo-cli by Valutac\n")
	log.SetFlags(log.Lshortfile)
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("[FATAL] get current user failed: %s\n", err.Error())
	}
	cfg := cfg.Init(fmt.Sprintf("%s/.neov/config.toml", usr.HomeDir))
	loginSrv := authentication.NewService(cfg)
	if token, err := loginSrv.Login(); err != nil {
		log.Fatalf("[FATAL] authentication failed: %s\n", err.Error())
	} else {
		computeSrv := compute.NewService(token)
		servers := computeSrv.GetServers(cfg.Credential.ProjectID)
		compute.DisplayServerList(servers)
	}

}

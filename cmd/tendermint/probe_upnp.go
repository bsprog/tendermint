package main

import (
	"encoding/json"
	"fmt"

	"github.com/tendermint/tendermint/p2p/upnp"
)

func probe_upnp() {

	capabilities, err := upnp.Probe()
	if err != nil {
		fmt.Println("Probe failed: %v", err)
	} else {
		fmt.Println("Probe success!")
		jsonBytes, err := json.Marshal(capabilities)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(jsonBytes))
	}

}

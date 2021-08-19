package main

import "fmt"

type Mobile struct {
	MobileName string
	price      int
	network    Network
}
type Network struct {
	networkName string
	singnal     bool
}

func shop() {
	brand := Mobile{
		MobileName: "realme",
		price:      20000,
		network: Network{
			networkName: "idea",
			singnal:     true,
		},
	}
	fmt.Printf("\n%v", brand)
	fmt.Printf("\n%v", brand.network.networkName)

}

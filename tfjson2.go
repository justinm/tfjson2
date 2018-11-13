package main

import (
	"flag"
	"fmt"
	"github.com/justinm/tfjson2/tfjson2"
	"log"
	"os"
)

func getCwd() string {
	return os.Getenv("PWD")
}

func main() {
	pathToPlan := flag.String("plan", "", "Path to the plan")

	flag.Parse()

	if len(*pathToPlan) == 0 {
		log.Fatal("--plan must be supplied")
	}

	plan, err := tfjson2.OpenPlan(*pathToPlan)
	if err != nil {
		log.Fatal(err)
	}

	exporter := tfjson2.JsonExporter{Plan: plan}

	json, err := exporter.ToJson()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(*json)
}

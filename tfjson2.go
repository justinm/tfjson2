package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/terraform/terraform"
	"github.com/cloudvar/tfjson_with_open-policy-agent/tfjson2"
	"log"
	"os"
)

func getCwd() string {
	return os.Getenv("PWD")
}

func main() {
	var plan *terraform.Plan
	pathToPlan := flag.String("plan", "", "Path to the plan")
	useStdin := flag.Bool("stdin", false, "Use stdin to read the plan data")

	flag.Parse()

	if !*useStdin {
		if len(*pathToPlan) == 0 {
			log.Fatal("--plan must be supplied")
		}

		p, err := tfjson2.OpenPlan(*pathToPlan)
		if err != nil {
			log.Fatal(err)
		}
		plan = p
	} else {
		p, err := terraform.ReadPlan(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		plan = p
	}

	exporter := tfjson2.JsonExporter{Plan: plan}

	json, err := exporter.ToJson()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(*json)
}

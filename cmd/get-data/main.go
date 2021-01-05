package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"flag"

	"github.com/micheam/wiseman/scrumwise"
)

var dataVersion bool

func init() {
	log.SetOutput(os.Stderr)
}

func main() {
	ctx := context.Background()
	flag.BoolVar(&dataVersion, "data-version", false, "show only data-version")
	flag.Parse()

	if dataVersion {
		var dversion int64
		dversion, err := scrumwise.GetDataVersion(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d", dversion)
		return
	}

	projectID := os.Getenv("SCRUMWISE_PROJECT")
	param := scrumwise.NewGetDataParam(projectID)
	data, err := scrumwise.GetData(ctx, *param)
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(b))
	log.Println("Success")
}

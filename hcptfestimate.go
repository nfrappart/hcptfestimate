package main

import (

    "fmt"
    "encoding/json"
    "io/ioutil"
    "os"
//    "flag"
)

type Resource struct {
  Mode string `json:"mode"`
  Type string `json:"type"`
  Name string `json:"name"`
  Provider string `json:"provider"`
  Instances []interface{} `json:"instances"`
}

type Workspace struct {
  CheckResults interface{} `json:"check_results"`
  Lineage string `json:"lineage"`
  Output interface{} `json:"outputs"`
  Resources []Resource `json:"resources"`
  Serial int `json:"serial"`
  TerraformVersion string `json:"terraform_version"`
  Version int `json:"version"`
}

func countFilteredResources(jsonData []byte) int {
  var workspace Workspace
	err := json.Unmarshal(jsonData, &workspace)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	var count int
	for _, resource := range workspace.Resources {
    if resource.Mode == "managed" && resource.Type != "null_resource" && resource.Type != "terraform_data" {
      for _, _ = range resource.Instances {
        count++
      }
		}
	}
	return count
}

func main() {
  // Define ANSI color codes
  const (
    Red    = "\033[31m"
    Green  = "\033[32m"
    Yellow = "\033[33m"
    Blue   = "\033[34m"
    Reset  = "\033[0m" // Reset to default color
  )

  // Check the number of arguments
  if len(os.Args) != 2 {
    fmt.Println("\nError:\nhcptfestimate require to pass the location of the state file as argument.\nExample: hcptfestimate ./terraform.tfstate\n")
      os.Exit(1)
  }

  jsonState, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }

  defer jsonState.Close()
  
  state, err := ioutil.ReadAll(jsonState)
  if err != nil {
    panic(err)
  }

  var count int = countFilteredResources(state)
  var hourlyrate float64 = 0.00014
  var monthlyprice float64 = float64(count - 500) * float64(730) * hourlyrate
  fmt.Println(Green + "\n### HCP Terraform cost estimation for the provided state file: ###" + Reset +"\n(hcptfestimate is freely provided by Ryzhom SAS.)")
  fmt.Println("\n# Pricing for Standard Tier: #")

  if count < 500 {
    fmt.Println("Resources managed in this workspace:", Yellow , count, Reset)
    fmt.Println("Price per month in HCP:", Yellow +"$0", Reset)
    fmt.Println("(First 500 managed resources are free.)\n")
  }

  if count > 500 {
    fmt.Println("Resources managed in this workspace:", Yellow , count, Reset)
    fmt.Println("Price per month in HCP:", Yellow +"$", monthlyprice, Reset)
    fmt.Println("(First 500 managed resources are free.)\n")
  }
}

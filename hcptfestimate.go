package main

import (

    "fmt"
    "encoding/json"
    "io/ioutil"
    "os"
)

type Resource struct {
  Mode string `json:"mode"`
  Type string `json:"type"`
  Name string `json:"name"`
  Provider string `json:"provider"`
  Instance []interface{} `json:"instances"`
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
		if resource.Mode == "managed" && resource.Type != "null_resource" {
			count++
		}
	}
	return count
}

func main() {
//  jsonState, err := os.Open("./backupstate.json")
  jsonState, err := os.Open(os.Args[1]  )
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
  var monthlyprice float64 = float64(count) * float64(730) * hourlyrate
  fmt.Println("Resources managed in this workspace:", count, "\nPrice per month in HCP: $", monthlyprice)

}



//Counting Filtered JSON Objects
//To count filtered objects in JSON with Go, you can use the encoding/json package to unmarshal the JSON data into a Go struct, and then use a loop to iterate over the objects and count the ones that match your filter criteria.
//
//Here’s an example:
//
//package main
//
//import (
//	"encoding/json"
//	"fmt"
//)
//
//type Bird struct {
//	ID    int    `json:"Id"`
//	Location string `json:"Location"`
//	Content string `json:"Content"`
//}
//
//func countFilteredBirds(jsonData []byte, location string) int {
//	var birds []Bird
//	err := json.Unmarshal(jsonData, &birds)
//	if err != nil {
//		fmt.Println(err)
//		return 0
//	}
//
//	var count int
//	for _, bird := range birds {
//		if bird.Location == location {
//			count++
//		}
//	}
//	return count
//}
//
//func main() {
//	jsonData := []byte(`[
//		{"Id": 13, "Location": "Australia", "Content": "Another string"},
//		{"Id": 145, "Location": "England", "Content": "SomeString"},
//		{"Id": 12, "Location": "England", "Content": "SomeString"},
//		{"Id": 12331, "Location": "Sweden", "Content": "SomeString"},
//		{"Id": 213123, "Location": "England", "Content": "SomeString"}
//	]`)
//
//	location := "England"
//	count := countFilteredBirds(jsonData, location)
//	fmt.Println("Count of birds in", location, ":", count)
//}
//
//In this example, we define a Bird struct to represent the JSON objects, and a countFilteredBirds function that takes a JSON byte slice and a location string as input. The function unmarshals the JSON data into a slice of Bird structs, and then iterates over the slice to count the number of birds that match the specified location.
//
//In the main function, we define the JSON data as a byte slice, and call the countFilteredBirds function with the desired location (“England”). The output will be the count of birds in England.
//
//Note that this example assumes that the JSON data is an array of objects, where each object represents a bird. If your JSON data has a different structure, you may need to modify the code accordingly.
//
//Also, if you’re working with a large JSON dataset, you may want to consider using a more efficient approach, such as using a JSON parsing library like jsoniter or gonum, or using a database to store and query the data.

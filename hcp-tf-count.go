package main

import (

    "fmt"
    "encoding/json"
    "io/ioutil"
    "os"
)

//type project struct {
//  Resources []string `json:"resources"`
//  Data []string `json:"data"`
//  Null []string `json:"null"`
//}

func countResources(jsonData []byte, location string) int {
	var resources []Resource
	err := json.Unmarshal(jsonData, &resource)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	var count int
	for _, resource := range resources {
		if bird.Location == location {
			count++
		}
	}
	return count
}

func main() {
  jsonFile, err := os.Open("statebackup.json")
  if err != nil {
    panic(err)
  }

  defer jsonFile.Close()
  
  state, err := ioutil.ReadAll(jsonFile)
  if err != nil {
    panic(err)
  }
  // commands to call terraform binary and display state file
    //command:= exec.Command("/usr/bin/terraform", "states", "pull")
    //stdout, state := command.Output()

    //if state != nil {
    //    fmt.Println(state.Error())
    //    return
    //}

    //fmt.Println(string(stdout))

  // commands to parse state file (json) and count resources, data sources and null resources. Then calculate price for the workspace for HCP Terraform standard


}



Counting Filtered JSON Objects
To count filtered objects in JSON with Go, you can use the encoding/json package to unmarshal the JSON data into a Go struct, and then use a loop to iterate over the objects and count the ones that match your filter criteria.

Here’s an example:

package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	ID    int    `json:"Id"`
	Location string `json:"Location"`
	Content string `json:"Content"`
}

func countFilteredBirds(jsonData []byte, location string) int {
	var birds []Bird
	err := json.Unmarshal(jsonData, &birds)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	var count int
	for _, bird := range birds {
		if bird.Location == location {
			count++
		}
	}
	return count
}

func main() {
	jsonData := []byte(`[
		{"Id": 13, "Location": "Australia", "Content": "Another string"},
		{"Id": 145, "Location": "England", "Content": "SomeString"},
		{"Id": 12, "Location": "England", "Content": "SomeString"},
		{"Id": 12331, "Location": "Sweden", "Content": "SomeString"},
		{"Id": 213123, "Location": "England", "Content": "SomeString"}
	]`)

	location := "England"
	count := countFilteredBirds(jsonData, location)
	fmt.Println("Count of birds in", location, ":", count)
}

In this example, we define a Bird struct to represent the JSON objects, and a countFilteredBirds function that takes a JSON byte slice and a location string as input. The function unmarshals the JSON data into a slice of Bird structs, and then iterates over the slice to count the number of birds that match the specified location.

In the main function, we define the JSON data as a byte slice, and call the countFilteredBirds function with the desired location (“England”). The output will be the count of birds in England.

Note that this example assumes that the JSON data is an array of objects, where each object represents a bird. If your JSON data has a different structure, you may need to modify the code accordingly.

Also, if you’re working with a large JSON dataset, you may want to consider using a more efficient approach, such as using a JSON parsing library like jsoniter or gonum, or using a database to store and query the data.

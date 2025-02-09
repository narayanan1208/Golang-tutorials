package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"` // dash means it will remove while displaying
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("More about JSON")
	// EncodeJSON()
	DecodeJson()
}

func EncodeJSON() {
	courses := []course{
		{Name: "React", Price: 199, Platform: "learncode.in", Password: "abc123", Tags: []string{"web-dev", "js"}},
		{Name: "MERN", Price: 299, Platform: "learncode.in", Password: "bcd123", Tags: []string{"full-stack", "js"}},
		{Name: "Angular", Price: 399, Platform: "learncode.in", Password: "cde123", Tags: nil},
	}

	// package this data as JSON data
	// finalJSON, err := json.Marshal(courses)
	finalJSON, err := json.MarshalIndent(courses, "", "\t")
	checkNil(err)
	fmt.Printf("%s\n", finalJSON)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		[
			{
							"coursename": "ReactJS Bootcamp",
							"price": 299,
							"website": "learncodeonline.in",
							"tags": [
											"web-dev",
											"js"
							]
			},
			{
							"coursename": "MERN Bootcamp",
							"price": 199,
							"website": "learncodeonline.in",
							"tags": [
											"full-stack",
											"js"
							]
			},
			{
							"coursename": "Angular Bootcamp",
							"price": 299,
							"website": "learncodeonline.in"
			}
		]
	`)

	var lcoCourses []course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid!")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses) // passing reference, 'cause i want to save it here
		fmt.Printf("%#v\n", lcoCourses)
	} else {
		fmt.Println("JSON was invalid")
	}

	fmt.Println("--------------------------")

	// sometimes i don't waana create the struct
	var myOnlineCourses []map[string]interface{} // interface - 'cause we do not know the type of it's value
	json.Unmarshal(jsonDataFromWeb, &myOnlineCourses)

	fmt.Printf("%#v\n", myOnlineCourses)

	fmt.Println("--------------------------")

	for _, course := range myOnlineCourses {
		fmt.Println(course)
		for key, value := range course {
			fmt.Printf("[%v] = %v\n", key, value)
		}
		fmt.Println("")
	}

}

func checkNil(err error) {
	if err != nil {
		panic(err)
	}
}

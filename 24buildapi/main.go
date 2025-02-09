package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/exp/rand"
)

// Model for courses - file
type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` // type is a pointer
}

type Author struct {
	Fullname string `json:"full_name"`
	Website  string `json:"website"`
}

// fake database
var courses []Course

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Building APIs")
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Narayanan S", Website: "go.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Mern stack", CoursePrice: 199, Author: &Author{Fullname: "Abhilash S", Website: "go.dev"}})

	// routing
	r.HandleFunc("/", servHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listening
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - file

func servHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Home page</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// get id from request
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// response := map[string]string{
	// 	"message": "No course for given id:",
	// 	"id":      params["id"],
	// }
	response := "No course for given id: " + params["id"]
	json.NewEncoder(w).Encode(response)
	return
}

// create a course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode(`{"msg": "Please send some data"}`)
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	// what about - {}
	if course.IsEmpty() {
		json.NewEncoder(w).Encode(`{"msg": "No data in JSON"}`)
		return
	}

	// generate a unique id, convert to string
	rand.Seed(uint64(time.Now().UnixNano()))
	course.CourseId = strconv.Itoa(rand.Intn(100))

	// append this course to course slice
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

// update one course
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses, remove it and add with exsisting id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var updatedCourse Course
			_ = json.NewDecoder(r.Body).Decode(&updatedCourse)

			updatedCourse.CourseId = params["id"]
			courses = append(courses, updatedCourse)
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}

	json.NewEncoder(w).Encode(`{"msg": "No course found with the given ID"}`)
}

// delete a course
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses, remove it
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(`{"msg": "Delete successful!"}`)
			return
		}
	}

	json.NewEncoder(w).Encode(`{"msg": "No course found with the given ID"}`)
}

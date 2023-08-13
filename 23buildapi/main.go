package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fakeDB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API in golang")

	fmt.Println("Seeding in memory database...")
	courses = append(courses, Course{CourseId: "1", CourseName: "1st course", CoursePrice: 200, Author: &Author{Fullname: "Hasibul Hasan", Website: "hasib@go.dev"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "2nd course", CoursePrice: 300, Author: &Author{Fullname: "Hasibul Hasan", Website: "hasib@somethingelse.dev"}})

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3400", r))
}

// controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to api by hasib</h1>"))
}


func getOneCourse(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from Request
	params := mux.Vars(r)

	// loop through the courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found")
	return
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Get all courses")
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(courses)
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside")
		return
	}

	// generate unique id, string
	// append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from the request
	params := mux.Vars(r)

	// loop through the value
	// get id and remove
	// add the value with myId
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting one course")

	param := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == param["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(course.CourseId)
			break
		}
	}
}

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

// Model for course - file

type Course struct {
	CourseId    string  `json:"courseid`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type CourseValidation struct {
	CourseId    string `json:"courseid,validate:"required"`
	CourseName  string `json:"coursename",validate:"required"`
	CoursePrice string `json:"price,validate:"required"`
	Author      int    `json:"author",validate:"required"`
}

type Author struct {
	FullName string `json:"full_name"`
	Website  string `json:"website"`
}

//fake DB

var courses []Course

// middleware, helper - file

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""

}

func main() {
	fmt.Println("Server started")
	courses = append(courses, Course{CourseId: "1", CourseName: "Python", CoursePrice: 299, Author: &Author{FullName: "Sanket", Website: "go.dev"}})
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", addCourse).Methods("POST")
	r.HandleFunc("/courses", getAllCourse).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to </h1>"))
}

func addCourse(w http.ResponseWriter, err interface{}, r *http.Request) {

	fmt.Println("Add one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		// json.NewEncoder(w).Encode("Please send some data")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func getAllCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func updateCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for idx, val := range courses {
		if val.CourseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// Todo : send a response when id is not found

}

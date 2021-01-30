package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Course Struct (Model)
type Course struct {
	ID             string      `json:"id"`
	Category       string      `json:"category"`
	Title          string      `json:"title"`
	Instructor     *Instructor `json:"instructor"`
	CourseDuration string      `json:"cousreDuration"`
}

// Init courses var as slice Course struct
var courses []Course

// Get All Courses
func getCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// Get Single Course
func getCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through courses and find with id
	for _, item := range courses {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Course{})
}

// Create a New Course
func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	course.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

// Update Course
func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range courses {
		if item.ID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.ID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode(courses)
}

// Delete Course
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range courses {
		if item.ID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(courses)
}

// Instructor Struct
type Instructor struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	// Init Router
	router := mux.NewRouter()

	// Mock Data - @todo - implement DB
	courses = append(
		courses, Course{
			ID:       "1",
			Category: "Web Development",
			Title:    "Build Web Apps with Go Language",
			Instructor: &Instructor{
				FirstName: "Rob",
				LastName:  "Pike",
			},
			CourseDuration: "4 months",
		},
	)
	courses = append(
		courses, Course{
			ID:       "2",
			Category: "Mobile App Development",
			Title:    "Build Mobile Apps with Flutter Dart",
			Instructor: &Instructor{
				FirstName: "Lars",
				LastName:  "Bak",
			},
			CourseDuration: "5 months",
		},
	)

	// Route Handler / Endpoints
	router.HandleFunc("/api/courses", getCourses).Methods("GET")
	router.HandleFunc("/api/courses/{id}", getCourse).Methods("GET")
	router.HandleFunc("/api/courses", createCourse).Methods("POST")
	router.HandleFunc("/api/courses/{id}", updateCourse).Methods("PUT")
	router.HandleFunc("/api/courses/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":1400", router))
}

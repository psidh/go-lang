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

// Course Model - file (generally)
type Course struct {
	CourseId   string  `json:"courseid"`
	CourseName string  `json:"coursename"`
	Price      int     `json:"price"`
	Author     *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var db []Course

// Middleware / Helper functions
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Dummy database
	db = append(db, Course{
		CourseId:   "2",
		CourseName: "ReactJS",
		Price:      199,
		Author:     &Author{Fullname: "Sidharth", Website: "https://psidh.vercel.app"},
	})

	fmt.Println("Server starting....")
	router := mux.NewRouter()

	router.HandleFunc("/", serveHome)
	router.HandleFunc("/getAll", getAllCourses).Methods("GET")
	router.HandleFunc("/create", createOneCourse).Methods("POST")
	router.HandleFunc("/getId/{id}", getCourseById).Methods("GET")
	router.HandleFunc("/update/{id}", updateCourse).Methods("PUT")
	router.HandleFunc("/delete/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", router))
}

// Controllers
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Serving home...</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all the courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db)
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, course := range db {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	http.Error(w, "No course with given ID", http.StatusNotFound)
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating one course")
	w.Header().Set("Content-Type", "application/json")

	var course Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil || course.IsEmpty() {
		http.Error(w, "Please send valid course data", http.StatusBadRequest)
		return
	}

	course.CourseId = strconv.Itoa(rand.Intn(100))
	db = append(db, course)
	json.NewEncoder(w).Encode(course)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var updatedCourse Course
	if err := json.NewDecoder(r.Body).Decode(&updatedCourse); err != nil || updatedCourse.IsEmpty() {
		http.Error(w, "Invalid course data", http.StatusBadRequest)
		return
	}

	for index, course := range db {
		if course.CourseId == params["id"] {
			updatedCourse.CourseId = params["id"]
			db[index] = updatedCourse
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}

	http.Error(w, "No course with given ID", http.StatusNotFound)
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, course := range db {
		if course.CourseId == params["id"] {
			db = append(db[:index], db[index+1:]...)
			w.Write([]byte("Deleted the course"))
			return
		}
	}

	http.Error(w, "No course with given ID", http.StatusNotFound)
}

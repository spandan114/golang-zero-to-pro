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

//Models
type Course struct {
	CourseId    string  `json:"id"`
	CourseName  string  `json:"name"`
	CoursePrice int     `json:"-"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

//Fake db
var courses []Course

//Middle ware
func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getCourse).Methods("GET")
	r.HandleFunc("/get-one-courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", addCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	//Seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{FullName: "Spandan Joshi", Website: "go.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{FullName: "Spandan Joshi", Website: "go.dev"}})

	//Run server
	fmt.Println("Server start")
	log.Fatal(http.ListenAndServe(":8080", r))
}

//controllers

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello world !")
}

func addCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Body is requeres")
		return
	}

	var myCourse Course
	json.NewDecoder(r.Body).Decode(&myCourse)

	for _, course := range courses {
		if course.CourseName == myCourse.CourseName {
			json.NewEncoder(w).Encode("Course name alredy exist")
			return
		}
	}

	if myCourse.isEmpty() {
		json.NewEncoder(w).Encode("Course is empty")
		return
	}

	rand.Seed(time.Now().UnixNano())
	myCourse.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, myCourse)
	json.NewEncoder(w).Encode(courses)
	return

}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get single course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found")
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get single course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var updatedCourse Course
			json.NewDecoder(r.Body).Decode(&updatedCourse)
			updatedCourse.CourseId = params["id"]
			courses = append(courses, updatedCourse)
			json.NewEncoder(w).Encode("Course " + params["id"] + " Updated successfully")
			return
		}
	}
	json.NewEncoder(w).Encode(params["id"] + " is not valid")

}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get single course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(params["id"] + " is deleted")
			break
		}
	}
	json.NewEncoder(w).Encode(params["id"] + " is not valid")
}

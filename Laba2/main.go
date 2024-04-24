package main

import (
	"log"
	"net/http"
	"Laba2/controllers" 
	"Laba2/service"     
	"Laba2/models"
)

func initializeData(service *service.Service) {
	// Создаем преподавателей
	teacher1 := models.Teacher{Name: "Alex Kov", Email: "alex.doe@gmail.com"}
	teacher2 := models.Teacher{Name: "Ulia Ykubovskay", Email: "ulia.smith@gmail.com"}
	service.CreateTeacher(teacher1)
	service.CreateTeacher(teacher2)

	// Создаем курсы
	course1 := models.Course{Title: "Introduction to Programming", TeacherID: 1, Price: 100}
	course2 := models.Course{Title: "Web Development", TeacherID: 2, Price: 150}
	service.CreateCourse(course1)
	service.CreateCourse(course2)

	// Создаем студентов
	student1 := models.Student{Name: "Misha Fedotov", Email: "misha@gmail.com"}
	student2 := models.Student{Name: "Slava Popov", Email: "slava@gmail.com"}
	service.CreateStudent(student1)
	service.CreateStudent(student2)
}

func main() {
	// Создание экземпляра источника данных и сервиса
	dataSource := service.NewDataSource()
	service := service.NewService(dataSource)
	controller := controllers.NewController(service)

	initializeData(service)
	// Регистрация обработчиков маршрутов
	http.HandleFunc("/teachers", controller.GetAllTeachersHandler)
	http.HandleFunc("/teachers/create", controller.CreateTeacherHandler)
	http.HandleFunc("/teachers/update", controller.UpdateTeacherHandler)
	http.HandleFunc("/teachers/delete", controller.DeleteTeacherHandler)

	http.HandleFunc("/courses", controller.GetAllCoursesHandler)
	http.HandleFunc("/courses/create", controller.CreateCourseHandler)
	http.HandleFunc("/courses/update", controller.UpdateCourseHandler)
	http.HandleFunc("/courses/delete", controller.DeleteCourseHandler)

	http.HandleFunc("/students", controller.GetAllStudentsHandler)
	http.HandleFunc("/students/create", controller.CreateStudentHandler)
	http.HandleFunc("/students/update", controller.UpdateStudentHandler)
	http.HandleFunc("/students/delete", controller.DeleteStudentHandler)

	// Запуск сервера на порту 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

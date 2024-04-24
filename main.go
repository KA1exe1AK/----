package main

import (
	"encoding/json"
	"log"
	"net/http"

	)

// respondWithError отправляет ответ с ошибкой в формате JSON
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// respondWithJSON отправляет ответ в формате JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Внутренняя ошибка сервера"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func initializeData(service *Service) {
	// Создаем преподавателей
	teacher1 := Teacher{Name: "Alex Kov", Email: "alex.doe@gmail.com"}
	teacher2 := Teacher{Name: "Ulia Ykubovskay", Email: "ulia.smith@gmail.com"}
	service.CreateTeacher(teacher1)
	service.CreateTeacher(teacher2)

	// Создаем курсы
	// course1 := Course{Title: "Introduction to Programming", TeacherID: 1, Price: 100}
	// course2 := Course{Title: "Web Development", TeacherID: 2, Price: 150}
	// service.Create("course", course1)
	// service.Create("course", course2)

	// Создаем студентов
	student1 := Student{Name: "Misha Fedotov", Email: "misha@gmail.com"}
	student2 := Student{Name: "Slava Popov", Email: "slava@gmail.com"}
	service.CreateStudent(student1)
	service.CreateStudent(student2)
}

// HealthCheckHandler обработчик для проверки работоспособности сервера
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Сервер работает"})
}

func main() {
	// Создание экземпляра источника данных и сервиса
	dsn := "postgres://user:password@localhost:5432/postgres?sslmode=disable"

	dataSource, err := NewDataSource(dsn)
	if err != nil {
		log.Panic("couldnt create db, ", err)
		return
	}
	service := NewService(dataSource)
	controller := NewController(service)

	initializeData(service)
	// Регистрация обработчиков маршрутов
	http.HandleFunc("/teachers", controller.GetAllTeachersHandler)
	http.HandleFunc("/teachers/create", controller.CreateTeacherHandler)
	http.HandleFunc("/teachers/update", controller.UpdateTeacherHandler)
	http.HandleFunc("/teachers/delete", controller.DeleteTeacherHandler)

	// http.HandleFunc("/courses", controller.GetAllCoursesHandler)
	// http.HandleFunc("/courses/create", controller.CreateCourseHandler)
	// http.HandleFunc("/courses/update", controller.UpdateCourseHandler)
	// http.HandleFunc("/courses/delete", controller.DeleteCourseHandler)

	http.HandleFunc("/students", controller.GetAllStudentsHandler)
	http.HandleFunc("/students/create", controller.CreateStudentHandler)
	http.HandleFunc("/students/update", controller.UpdateStudentHandler)
	http.HandleFunc("/students/delete", controller.DeleteStudentHandler)

	// Регистрация обработчика проверки состояния сервера
	http.HandleFunc("/health", HealthCheckHandler)

	// Запуск сервера на порту 8080
	log.Fatal(http.ListenAndServe(":8081", nil))
}

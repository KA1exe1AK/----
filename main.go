package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Teacher модель преподавателя
type Teacher struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Course модель курса
type Course struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	TeacherID int     `json:"teacher_id"`
	Price     float64 `json:"price"`
}

// Student модель студента
type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// DataSource объект для хранения коллекции экземпляров сущностей
type DataSource struct {
	teachers []Teacher
	courses  []Course
	students []Student
	nextID   map[string]int // Карта для отслеживания следующего доступного ID
}

// Service сервис выполнения CRUD операций
type Service struct {
	dataSource *DataSource
}

// NewDataSource создает новый экземпляр DataSource
func NewDataSource() *DataSource {
	return &DataSource{
		teachers: []Teacher{},
		courses:  []Course{},
		students: []Student{},
		nextID: map[string]int{
			"teacher": 1, // Начальное значение ID для преподавателей
			"course":  1, // Начальное значение ID для курсов
			"student": 1, // Начальное значение ID для студентов
		},
	}
}

// NewService создает новый экземпляр Service
func NewService(dataSource *DataSource) *Service {
	return &Service{
		dataSource: dataSource,
	}
}

// Create добавляет новый объект в соответствующую коллекцию
func (s *Service) Create(dataType string, data interface{}) {
	switch dataType {
	case "teacher":
		teacher := data.(Teacher)
		teacher.ID = s.dataSource.nextID[dataType] // Присваиваем ID
		s.dataSource.teachers = append(s.dataSource.teachers, teacher)
		s.dataSource.nextID[dataType]++ // Увеличиваем следующий ID
	case "course":
		course := data.(Course)
		course.ID = s.dataSource.nextID[dataType] // Присваиваем ID
		s.dataSource.courses = append(s.dataSource.courses, course)
		s.dataSource.nextID[dataType]++ // Увеличиваем следующий ID
	case "student":
		student := data.(Student)
		student.ID = s.dataSource.nextID[dataType] // Присваиваем ID
		s.dataSource.students = append(s.dataSource.students, student)
		s.dataSource.nextID[dataType]++ // Увеличиваем следующий ID
	}
}

// GetAll возвращает все объекты соответствующей коллекции
func (s *Service) GetAll(dataType string) interface{} {
	switch dataType {
	case "teacher":
		return s.dataSource.teachers
	case "course":
		return s.dataSource.courses
	case "student":
		return s.dataSource.students
	}
	return nil
}

// Update обновляет объект с заданным ID
func (s *Service) Update(dataType string, id int, newData interface{}) error {
	switch dataType {
	case "teacher":
		newTeacher := newData.(Teacher)
		for i, teacher := range s.dataSource.teachers {
			if teacher.ID == id {
				// Сохраняем существующий ID у нового преподавателя
				newTeacher.ID = teacher.ID
				s.dataSource.teachers[i] = newTeacher
				return nil
			}
		}
	case "course":
		newCourse := newData.(Course)
		for i, course := range s.dataSource.courses {
			if course.ID == id {
				s.dataSource.courses[i] = newCourse
				return nil
			}
		}
	case "student":
		newStudent := newData.(Student)
		for i, student := range s.dataSource.students {
			if student.ID == id {
				s.dataSource.students[i] = newStudent
				return nil
			}
		}
	}
	return fmt.Errorf("объект с ID %d не найден", id)
}

// Delete удаляет объект с заданным ID
func (s *Service) Delete(dataType string, id int) error {
	switch dataType {
	case "teacher":
		for i, teacher := range s.dataSource.teachers {
			if teacher.ID == id {
				s.dataSource.teachers = append(s.dataSource.teachers[:i], s.dataSource.teachers[i+1:]...)
				return nil
			}
		}
	case "course":
		for i, course := range s.dataSource.courses {
			if course.ID == id {
				s.dataSource.courses = append(s.dataSource.courses[:i], s.dataSource.courses[i+1:]...)
				return nil
			}
		}
	case "student":
		for i, student := range s.dataSource.students {
			if student.ID == id {
				s.dataSource.students = append(s.dataSource.students[:i], s.dataSource.students[i+1:]...)
				return nil
			}
		}
	}
	return fmt.Errorf("объект с ID %d не найден", id)
}

// Controller контроллер для обработки запросов
type Controller struct {
	service *Service
}

// NewController создает новый экземпляр Controller
func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}

// GetAllTeachersHandler обработчик для получения всех преподавателей
func (c *Controller) GetAllTeachersHandler(w http.ResponseWriter, r *http.Request) {
	data := c.service.GetAll("teacher")
	respondWithJSON(w, http.StatusOK, data)
}

// CreateTeacherHandler обработчик для создания нового преподавателя
func (c *Controller) CreateTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var teacher Teacher
	err := json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}
	defer r.Body.Close()

	c.service.Create("teacher", teacher)
	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Преподаватель успешно создан"})
}

// UpdateTeacherHandler обработчик для обновления данных преподавателя
func (c *Controller) UpdateTeacherHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверный ID")
		return
	}

	var teacher Teacher
	err = json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}
	defer r.Body.Close()

	err = c.service.Update("teacher", id, teacher)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Преподаватель успешно обновлен"})
}

// DeleteTeacherHandler обработчик для удаления преподавателя
func (c *Controller) DeleteTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID int `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}
	defer r.Body.Close()

	err = c.service.Delete("teacher", req.ID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Преподаватель успешно удален"})
}

// GetAllCoursesHandler обработчик для получения всех курсов
func (c *Controller) GetAllCoursesHandler(w http.ResponseWriter, r *http.Request) {
	data := c.service.GetAll("course")
	respondWithJSON(w, http.StatusOK, data)
}

// CreateCourseHandler обработчик для создания нового курса
func (c *Controller) CreateCourseHandler(w http.ResponseWriter, r *http.Request) {
	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}
	defer r.Body.Close()

	c.service.Create("course", course)
	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Курс успешно создан"})
}

// UpdateCourseHandler обработчик для обновления данных курса
func (c *Controller) UpdateCourseHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверный ID")
		return
	}

	var course Course
	err = json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}
	defer r.Body.Close()

	err = c.service.Update("course", id, course)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Курс успешно обновлен"})
}

// DeleteCourseHandler обработчик для удаления курса
func (c *Controller) DeleteCourseHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверный ID")
		return
	}

	err = c.service.Delete("course", id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Курс успешно удален"})
}

// GetAllStudentsHandler обработчик для получения всех студентов
func (c *Controller) GetAllStudentsHandler(w http.ResponseWriter, r *http.Request) {
	data := c.service.GetAll("student")
	respondWithJSON(w, http.StatusOK, data)
}

// CreateStudentHandler обработчик для создания нового студента
func (c *Controller) CreateStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}
	defer r.Body.Close()

	c.service.Create("student", student)
	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Студент успешно создан"})
}

// UpdateStudentHandler обработчик для обновления данных студента
func (c *Controller) UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверный ID")
		return
	}

	var student Student
	err = json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}
	defer r.Body.Close()

	err = c.service.Update("student", id, student)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Студент успешно обновлен"})
}

// DeleteStudentHandler обработчик для удаления студента
func (c *Controller) DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверный ID")
		return
	}

	err = c.service.Delete("student", id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Студент успешно удален"})
}

// getRequestID извлекает идентификатор из запроса
func getRequestID(r *http.Request) int {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	return id
}

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
	service.Create("teacher", teacher1)
	service.Create("teacher", teacher2)

	// Создаем курсы
	course1 := Course{Title: "Introduction to Programming", TeacherID: 1, Price: 100}
	course2 := Course{Title: "Web Development", TeacherID: 2, Price: 150}
	service.Create("course", course1)
	service.Create("course", course2)

	// Создаем студентов
	student1 := Student{Name: "Misha Fedotov", Email: "misha@gmail.com"}
	student2 := Student{Name: "Slava Popov", Email: "slava@gmail.com"}
	service.Create("student", student1)
	service.Create("student", student2)
}

// HealthCheckHandler обработчик для проверки работоспособности сервера
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Сервер работает"})
}

func main() {
	// Создание экземпляра источника данных и сервиса
	dataSource := NewDataSource()
	service := NewService(dataSource)
	controller := NewController(service)

	initializeData(service)
	// Регистрация обработчиков маршрутов
	http.HandleFunc("/api/teachers", controller.GetAllTeachersHandler)
	http.HandleFunc("/api/teachers/create", controller.CreateTeacherHandler)
	http.HandleFunc("/api/teachers/update", controller.UpdateTeacherHandler)
	http.HandleFunc("/api/teachers/delete", controller.DeleteTeacherHandler)

	http.HandleFunc("/api/courses", controller.GetAllCoursesHandler)
	http.HandleFunc("/api/courses/create", controller.CreateCourseHandler)
	http.HandleFunc("/api/courses/update", controller.UpdateCourseHandler)
	http.HandleFunc("/api/courses/delete", controller.DeleteCourseHandler)

	http.HandleFunc("/api/students", controller.GetAllStudentsHandler)
	http.HandleFunc("/api/students/create", controller.CreateStudentHandler)
	http.HandleFunc("/api/students/update", controller.UpdateStudentHandler)
	http.HandleFunc("/api/students/delete", controller.DeleteStudentHandler)

	// Запуск сервера на порту 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

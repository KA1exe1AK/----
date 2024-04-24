package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// GetAllTeachersHandler обработчик для получения всех преподавателей
func (c *Controller) GetAllTeachersHandler(w http.ResponseWriter, r *http.Request) {
	data, err := c.service.GetAllTeachers()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}
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

	c.service.CreateTeacher(teacher)
	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Преподаватель успешно создан"})
}

// UpdateTeacherHandler обработчик для обновления данных преподавателя
func (c *Controller) UpdateTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var teacher Teacher
	err := json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}
	defer r.Body.Close()

	err = c.service.UpdateTeacher(teacher)
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

	c.service.DeleteTeacher(req.ID)

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Преподаватель успешно удален"})
}

// // GetAllCoursesHandler обработчик для получения всех курсов
// func (c *Controller) GetAllCoursesHandler(w http.ResponseWriter, r *http.Request) {
// 	data := c.service.GetAll("course")
// 	respondWithJSON(w, http.StatusOK, data)
// }

// UpdateCourseHandler обработчик для обновления данных курса
func (c *Controller) UpdateCourseHandler(w http.ResponseWriter, r *http.Request) {
	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}
	defer r.Body.Close()

	// todo
	// err = c.service.Update("course", id, course)
	// if err != nil {
	// 	respondWithError(w, http.StatusNotFound, err.Error())
	// 	return
	// }

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Курс успешно обновлен"})
}

func (c *Controller) GetAllStudentsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := c.service.GetAllStudents()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}
	fmt.Println("Get all students: found : ", len(data))
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

	c.service.CreateStudent(student)
	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Студент успешно создан"})
}

// UpdateStudentHandler обработчик для обновления данных студента
func (c *Controller) UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}
	defer r.Body.Close()

	c.service.UpdateStudent(student)

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
	c.service.DeleteStudent(id)

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Студент успешно удален"})
}

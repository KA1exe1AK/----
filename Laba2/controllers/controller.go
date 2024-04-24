package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Laba2/service"
	"Laba2/models"
	"Laba2/utils"
)

type Controller struct {
	service *service.Service
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) GetAllTeachersHandler(w http.ResponseWriter, r *http.Request) {
	data := c.service.GetAllTeachers()
	utils.RespondWithJSON(w, http.StatusOK, data)
}

func (c *Controller) CreateTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var teacher models.Teacher
	err := json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}
	defer r.Body.Close()

	c.service.CreateTeacher(teacher)
	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Преподаватель успешно создан"})
}

func (c *Controller) UpdateTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var teacher models.Teacher
	err := json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}
	defer r.Body.Close()

	err = c.service.UpdateTeacher(teacher)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Преподаватель успешно обновлен"})
}

func (c *Controller) DeleteTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID int `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}
	defer r.Body.Close()

	c.service.DeleteTeacher(req.ID)

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Преподаватель успешно удален"})
}

func (c *Controller) GetAllCoursesHandler(w http.ResponseWriter, r *http.Request) {
	data := c.service.GetAllCourses()
	utils.RespondWithJSON(w, http.StatusOK, data)
}

func (c *Controller) UpdateCourseHandler(w http.ResponseWriter, r *http.Request) {
	var course models.Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}
	defer r.Body.Close()

	err = c.service.UpdateCourse(course)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Курс успешно обновлен"})
}

func (c *Controller) CreateCourseHandler(w http.ResponseWriter, r *http.Request) {
	var course models.Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}
	defer r.Body.Close()

	c.service.CreateCourse(course)
	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Курс успешно создан"})
}

func (c *Controller) DeleteCourseHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID int `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}
	defer r.Body.Close()

	c.service.DeleteCourse(req.ID)

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Курс успешно удален"})
}

func (c *Controller) GetAllStudentsHandler(w http.ResponseWriter, r *http.Request) {
	data := c.service.GetAllStudents()
	utils.RespondWithJSON(w, http.StatusOK, data)
}

func (c *Controller) CreateStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}
	defer r.Body.Close()

	c.service.CreateStudent(student)
	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Студент успешно создан"})
}

func (c *Controller) UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}
	defer r.Body.Close()

	c.service.UpdateStudent(student)
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Студент успешно обновлен"})
}

func (c *Controller) DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Неверный ID")
		return
	}
	c.service.DeleteStudent(id)
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Студент успешно удален"})
}

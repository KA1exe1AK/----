package service

import (
	"errors"
	"fmt"
	. "Laba2/models"
  )
  
  var (
	TeacherID int = 0
	StudentID int = 0
	CourseID int = 0

  )
  
  // DataSource объект для хранения коллекции экземпляров сущностей
  type DataSource struct {
	teachers map[int]Teacher
	courses  map[int]Course
	students map[int]Student
  }
  
  // Service сервис выполнения CRUD операций
  type Service struct {
	dataSource *DataSource
  }
  
  // NewDataSource создает новый экземпляр DataSource
  func NewDataSource() *DataSource {
	return &DataSource{
	  teachers: make(map[int]Teacher),
	  courses:  make(map[int]Course),
	  students: make(map[int]Student),
	}
  }
  
  // NewService создает новый экземпляр Service
  func NewService(dataSource *DataSource) *Service {
	return &Service{
	  dataSource: dataSource,
	}
  }
  
  func (s *Service) GetAllTeachers() []Teacher {
	res := []Teacher{}
  
	for _, v := range s.dataSource.teachers {
	  res = append(res, v)
	}
  
	return res
  }
  
  func (s *Service) GetAllStudents() []Student {
	res := []Student{}
  
	for _, v := range s.dataSource.students {
	  res = append(res, v)
	}
  
	return res
  }
  
  func (s *Service) GetAllCourses() []Course {
	res := []Course{}
  
	for _, v := range s.dataSource.courses {
	  res = append(res, v)
	}
  
	return res
  }

  func (s *Service) CreateTeacher(teacher Teacher) {
	teacher.ID = TeacherID
	s.dataSource.teachers[TeacherID] = teacher
	fmt.Println("new teacher created", TeacherID)
	TeacherID++
  }
  
  func (s *Service) CreateStudent(student Student) {
	student.ID = StudentID
	s.dataSource.students[StudentID] = student
	fmt.Println("new student created", StudentID)
	StudentID++
  }

  func (s *Service) CreateCourse(course Course) {
	course.ID = CourseID
	s.dataSource.courses[CourseID] = course
	fmt.Println("new course created", CourseID)
	CourseID++
  }
  
  func (s *Service) UpdateTeacher(teacher Teacher) error {
	if _, ok := s.dataSource.teachers[teacher.ID]; !ok {
	  fmt.Println("user not found ", teacher.ID)
	  return errors.New("user not found")
	}
  
	s.dataSource.teachers[teacher.ID] = teacher
	return nil
  }
  
  func (s *Service) UpdateStudent(student Student) error {
	if _, ok := s.dataSource.students[student.ID]; !ok {
	  fmt.Println("user not found ", student.ID)
	  return errors.New("syudent not found")
	}
  
	s.dataSource.students[student.ID] = student
	return nil
  }

  func (s *Service) UpdateCourse(course Course) error {
	if _, ok := s.dataSource.courses[course.ID]; !ok {
	  fmt.Println("user not found ", course.ID)
	  return errors.New("course not found")
	}
  
	s.dataSource.courses[course.ID] = course
	return nil
  }
  
  func (s *Service) DeleteTeacher(id int) {
  
	delete(s.dataSource.teachers, id)
  }
  
  func (s *Service) DeleteStudent(id int) {
  
	delete(s.dataSource.students, id)
  }

  func (s *Service) DeleteCourse(id int) {
  
	delete(s.dataSource.courses, id)
  }
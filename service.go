package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type DataSource struct {
	db *sql.DB
}

type Service struct {
	dataSource *DataSource
}

// NewDataSource создает новый экземпляр DataSource с подключением к PostgreSQL
func NewDataSource(dbURL string) (*DataSource, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DataSource{db: db}, nil
}

// NewService создает новый экземпляр Service
func NewService(dataSource *DataSource) *Service {
	return &Service{
		dataSource: dataSource,
	}
}

func (s *Service) GetAllTeachers() ([]Teacher, error) {
	rows, err := s.dataSource.db.Query("SELECT id, name FROM teachers")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var teachers []Teacher
	for rows.Next() {
		var teacher Teacher
		if err := rows.Scan(&teacher.ID, &teacher.Name); err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}

	return teachers, nil
}

func (s *Service) GetAllStudents() ([]Student, error) {
	rows, err := s.dataSource.db.Query("SELECT id, name FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.ID, &student.Name); err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}

func (s *Service) CreateTeacher(teacher Teacher) error {
	_, err := s.dataSource.db.Exec("INSERT INTO teachers (name) VALUES ($1)", teacher.Name)
	return err
}

func (s *Service) CreateStudent(student Student) error {
	_, err := s.dataSource.db.Exec("INSERT INTO students (name) VALUES ($1)", student.Name)

	if err != nil {
		fmt.Println(" error creating student", err)
	}
	return err
}

func (s *Service) UpdateTeacher(teacher Teacher) error {
	result, err := s.dataSource.db.Exec("UPDATE teachers SET name = $1 WHERE id = $2", teacher.Name, teacher.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("teacher not found")
	}
	return nil
}

func (s *Service) UpdateStudent(student Student) error {
	result, err := s.dataSource.db.Exec("UPDATE students SET name = $1 WHERE id = $2", student.Name, student.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("student not found")
	}
	return nil
}

func (s *Service) DeleteTeacher(id int) error {
	result, err := s.dataSource.db.Exec("DELETE FROM teachers WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("teacher not found")
	}
	return nil
}

func (s *Service) DeleteStudent(id int) error {
	result, err := s.dataSource.db.Exec("DELETE FROM students WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("student not found")
	}
	return nil
}

package main

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

package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-api/src/domain"
	"go-api/src/infrastructure/repository/models"
	"log"
)

type StudentPostgres struct {
	db *sqlx.DB
}

func NewStudentPostgres(db *sqlx.DB) *StudentPostgres {
	return &StudentPostgres{db: db}
}

// GetStudentById Получить по ID
func (p *StudentPostgres) GetStudentById(studentId int) (*domain.StudentDTO, error) {
	var student models.StudentTable
	var query = fmt.Sprintf("SELECT * FROM %s WHERE id = $1", studentsTable)
	err := p.db.Get(&student, query, studentId)
	if err != nil {
		log.Printf("Level: repos; func GetStudentById(): student with name = %d does not exist",
			studentId,
		)
		return nil, err
	}

	studentDto := models.ConvertStudentTableToStudentDTO(&student)
	return studentDto, err
}

// GetAllStudents Получить все
func (p *StudentPostgres) GetAllStudents() (*[]domain.StudentDTO, error) {
	var students []models.StudentTable
	var query = fmt.Sprintf("SELECT * FROM %s ORDER BY id", studentsTable)
	err := p.db.Select(&students, query)
	log.Printf("func GetAllStudents(): []domain.StudentDTO=%v", students)

	studentDTOs := make([]domain.StudentDTO, len(students))
	for i, student := range students {
		studentDTOs[i] = *models.ConvertStudentTableToStudentDTO(&student)
	}
	return &studentDTOs, err
}

// CreateStudent Создать
func (p *StudentPostgres) CreateStudent(student *domain.StudentDTO) (int, error) {
	var id int
	var createStudentQuery = fmt.Sprintf(
		"INSERT INTO %s (name, age, specialty) VALUES ($1, $2, $3) RETURNING id",
		studentsTable)
	row := p.db.QueryRow(createStudentQuery,
		student.Name,
		student.Age,
		student.Specialty,
	)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	log.Printf("func createStudent(): domain.StudentDTO=%v", student)
	return id, nil
}

// UpdateStudent Обновить данные
func (p *StudentPostgres) UpdateStudent(student *domain.StudentDTO) (int, error) {
	var updateStudentQuery = fmt.Sprintf(
		"UPDATE %s SET name = $1, age = $2, specialty = $3 WHERE id = $4",
		studentsTable,
	)
	_, err := p.db.Exec(updateStudentQuery,
		student.Name,
		student.Age,
		student.Specialty,
		student.Id,
	)
	if err != nil {
		return 0, err
	}

	log.Printf("func UpdateStudent(): domain.StudentDTO=%v", student)
	return student.Id, nil
}

// DeleteStudent Удалить
func (p *StudentPostgres) DeleteStudent(studentId int) error {
	var deleteStudentQuery = fmt.Sprintf("DELETE FROM %s WHERE id = $1", studentsTable)
	_, err := p.db.Exec(deleteStudentQuery,
		studentId,
	)
	if err != nil {
		return err
	}

	log.Printf("func DeleteStudent(): Student id=%d", studentId)
	return nil
}

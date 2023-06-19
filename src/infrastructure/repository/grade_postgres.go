package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-api/src/domain"
	"go-api/src/infrastructure/repository/models"
	"log"
)

type GradePostgres struct {
	db *sqlx.DB
}

func NewGradePostgres(db *sqlx.DB) *GradePostgres {
	return &GradePostgres{db: db}
}

// GetGradeById Получить по ID
func (p *GradePostgres) GetGradeById(gradeId int) (*domain.GradeDTO, error) {
	var grade models.GradeTable
	var query = fmt.Sprintf("SELECT * FROM %s WHERE id = $1", gradesTable)
	err := p.db.Get(&grade, query, gradeId)
	if err != nil {
		log.Printf("Level: repos; func GetGradeById(): grade with name = %d does not exist",
			gradeId,
		)
		return nil, err
	}

	gradeDto := models.ConvertGradeTableToGradeDTO(&grade)
	return gradeDto, err
}

// GetGradesByStudentId Получить все
func (p *GradePostgres) GetGradesByStudentId(studentId int) (*[]domain.GradeDTO, error) {
	var grades []models.GradeTable
	var getStudentGradesQuery = fmt.Sprintf(`
		SELECT * 
		FROM %s 
		WHERE id IN (
			SELECT grade_id 
			FROM %s 
			WHERE student_id = $1 
			ORDER BY grade_id) 
		ORDER BY id`,
		gradesTable,
		studentsGradesTable,
	)
	err := p.db.Select(&grades, getStudentGradesQuery)
	log.Printf("func GetAllGrades(): []domain.GradeDTO=%v", grades)

	gradeDTOs := make([]domain.GradeDTO, len(grades))
	for i, grade := range grades {
		gradeDTOs[i] = *models.ConvertGradeTableToGradeDTO(&grade)
	}
	return &gradeDTOs, err
}

// CreateGrade Создать
func (p *GradePostgres) CreateGrade(grade *domain.GradeDTO) (int, error) {
	var id int
	var createGradeQuery = fmt.Sprintf(
		"INSERT INTO %s (grade, subject_name, student_id) VALUES ($1, $2, $3) RETURNING id",
		gradesTable)
	row := p.db.QueryRow(createGradeQuery,
		grade.Grade,
		grade.SubjectName,
		grade.StudentId,
	)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	log.Printf("func createGrade(): domain.GradeDTO=%v", grade)
	return id, nil
}

// UpdateGrade Обновить данные
func (p *GradePostgres) UpdateGrade(grade *domain.GradeDTO) (int, error) {
	var updateGradeQuery = fmt.Sprintf(
		"UPDATE %s SET grade = $1, subject_name = $2, student_id = $3 WHERE id = $4",
		gradesTable,
	)
	_, err := p.db.Exec(updateGradeQuery,
		grade.Grade,
		grade.SubjectName,
		grade.StudentId,
		grade.Id,
	)
	if err != nil {
		return 0, err
	}

	log.Printf("func UpdateGrade(): domain.GradeDTO=%v", grade)
	return grade.Id, nil
}

// DeleteGrade Удалить
func (p *GradePostgres) DeleteGrade(gradeId int) error {
	var deleteGradeQuery = fmt.Sprintf("DELETE FROM %s WHERE id = $1", gradesTable)
	_, err := p.db.Exec(deleteGradeQuery,
		gradeId,
	)
	if err != nil {
		return err
	}

	log.Printf("func DeleteGrade(): Grade id=%d", gradeId)
	return nil
}

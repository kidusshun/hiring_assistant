package evaluationcritera

import (
	"database/sql"

	"github.com/google/uuid"
)

type store struct {
	db *sql.DB
}


func NewStore(db *sql.DB) *store {
	return &store{
		db: db,
	}
}


func (s *store) CreateEvaluationCriteria(jobPostingID uuid.UUID, criteriaName, description string, weight float32) (*EvaluationCriteria, error) {
	rows := s.db.QueryRow("INSERT INTO evaluation_criteria (job_posting_id, criteria_name, description, weight) VALUES ($1, $2, $3, $4)", jobPostingID, criteriaName, description, weight)

	createdEvaluationCriteria, err := ScanRowToEvaluationCriteria(rows)
	if err != nil {
		return nil, err
	}

	return createdEvaluationCriteria, nil


}

func (s *store) GetEvaluationCriteriaByJobPostingID(jobPostingID uuid.UUID) ([]*EvaluationCriteria, error) {
	rows, err := s.db.Query("SELECT * FROM evaluation_criteria WHERE job_posting_id = $1", jobPostingID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	evaluationCriterias := make([]*EvaluationCriteria, 0)

	for rows.Next() {
		var evaluationCriteria *EvaluationCriteria
		err := rows.Scan(&evaluationCriteria.ID, &evaluationCriteria.JobPostingID, &evaluationCriteria.CriteriaName, &evaluationCriteria.Description, &evaluationCriteria.Weight, &evaluationCriteria.CreatedAt, &evaluationCriteria.UpdatedAt)

		if err != nil {
			return nil, err
		}

		evaluationCriterias = append(evaluationCriterias, evaluationCriteria)
	}

	return evaluationCriterias, nil
}

func ScanRowToEvaluationCriteria(rows *sql.Row) (*EvaluationCriteria, error) {

	evaluationCriteria := new(EvaluationCriteria)
	err := rows.Scan(
		&evaluationCriteria.ID,
		&evaluationCriteria.JobPostingID,
		&evaluationCriteria.CriteriaName,
		&evaluationCriteria.Description,
		&evaluationCriteria.Weight,
		&evaluationCriteria.CreatedAt,
		&evaluationCriteria.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return evaluationCriteria, nil
}
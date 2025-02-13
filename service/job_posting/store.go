package jobposting

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


func (s *store) CreateJobPosting(jobPosting *JobPosting) (*JobPosting, error) {
	_, err := s.db.Exec("INSERT INTO job_postings (id, user_id, title, description, location, department, employment_type) VALUES ($1, $2, $3, $4, $5, $6, $7)", jobPosting.ID, jobPosting.UserID, jobPosting.Title, jobPosting.Description, jobPosting.Location, jobPosting.Department, jobPosting.EmploymentType)

	if err != nil {
		return nil, err
	}

	return jobPosting, nil
}

func (s *store) GetJobPostings(userID uuid.UUID, limit, offset int) ([]*JobPosting, error) {
	rows, err := s.db.Query("SELECT * FROM job_postings WHERE user_id = $1 LIMIT $2 OFFSET $3", userID, limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	jobPostings := make([]*JobPosting, 0)

	for rows.Next() {
		var jobPosting JobPosting
		err := rows.Scan(&jobPosting.ID, &jobPosting.UserID, &jobPosting.Title, &jobPosting.Description, &jobPosting.Location, &jobPosting.Department, &jobPosting.EmploymentType, &jobPosting.CreatedAt, &jobPosting.UpdatedAt)

		if err != nil {
			return nil, err
		}

		jobPostings = append(jobPostings, &jobPosting)
	}

	return jobPostings, nil
}
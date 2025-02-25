package resumes

import "database/sql"


type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}


func (s *Store) AddResumes(resumes []Resume) ([]Resume, error) {
	var storedResumes []Resume;
	for _, resume := range resumes {
		row := s.db.QueryRow("INSERT INTO resumes (job_posting_id, applicant_name, applicant_email, resume_file_path, resume_text, status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, job_posting_id, applicant_name, applicant_email, resume_file_path, resume_text, status, created_at, updated_at", resume.JobPostingID, resume.ApplicantName, resume.ApplicantEmail, resume.ResumePath, resume.ResumeText, resume.Status)
		createdResume, err := ScanRowToResume(row)
		if err != nil {
			return nil, err
		}
		storedResumes = append(storedResumes, *createdResume)
	}

	return storedResumes, nil
}

func ScanRowToResume(rows *sql.Row) (*Resume, error) {

	resume := new(Resume)
	err := rows.Scan(
		&resume.ID,
		&resume.JobPostingID,
		&resume.ApplicantName,
		&resume.ApplicantEmail,
		&resume.ResumePath,
		&resume.ResumeText,
		&resume.Status,
		&resume.CreatedAt,
		&resume.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return resume, nil
}

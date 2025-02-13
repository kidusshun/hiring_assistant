CREATE TABLE resumes (
    id UUID NOT NULL PRIMARY KEY,
    job_posting_id UUID NOT NULL,
    applicant_name VARCHAR(255),
    applicant_email VARCHAR(255),
    resume_file_path TEXT NOT NULL,
    status VARCHAR(50) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (job_posting_id) REFERENCES job_postings(id) ON DELETE CASCADE
);

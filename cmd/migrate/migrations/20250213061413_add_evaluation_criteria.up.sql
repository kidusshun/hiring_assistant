CREATE TABLE evaluation_criteria (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    job_posting_id UUID NOT NULL,
    criteria_name VARCHAR(255) NOT NULL,
    description TEXT,
    weight NUMERIC(5,2) DEFAULT 0.00,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (job_posting_id) REFERENCES job_postings(id) ON DELETE CASCADE
);

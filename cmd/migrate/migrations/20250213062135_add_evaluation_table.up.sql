CREATE TABLE evaluations (
    id UUID NOT NULL PRIMARY KEY,
    resume_id UUID NOT NULL,
    overall_score NUMERIC(5,2) NOT NULL,
    pass BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (resume_id) REFERENCES resumes(id)
);

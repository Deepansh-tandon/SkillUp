package db

import (
	"time"

	"github.com/pgvector/pgvector-go"
)

// Users
type User struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Email     string    `gorm:"uniqueIndex;size:255;not null"`
	Name      string    `gorm:"size:100"`
	Password  string    `gorm:"size:255;not null"` 
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// Goals
type Goal struct {
	ID         string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID     string    `gorm:"index;not null"`
	Title      string    `gorm:"not null"`
	TargetDate *time.Time
	Status     string    `gorm:"type:text;check:status IN ('active','completed')"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

// Topics (hierarchical)
type Topic struct {
	ID               string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID           string    `gorm:"index;not null"`
	GoalID           *string   `gorm:"index"`
	ParentID         *string   `gorm:"index"`
	Title            string    `gorm:"size:200;not null"`
	SourceType       string    `gorm:"type:varchar(20);default:'document';check:source_type IN ('syllabus','document')"`
	CompletionStatus string    `gorm:"type:varchar(20);default:'pending';check:completion_status IN ('pending','completed')"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
}

// Documents
type Document struct {
	ID               string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID           string    `gorm:"index;not null"`
	Filename         string    `gorm:"size:255;not null"`
	FilePath         string    `gorm:"size:500"` // if we later add S3
	ProcessingStatus string    `gorm:"type:varchar(20);default:'uploaded';check:processing_status IN ('uploaded','processed','failed')"`
	UploadDate       time.Time `gorm:"autoCreateTime"`
}

// Raw full-text extracted from document (optional)
type DocumentRaw struct {
	ID         string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	DocumentID string    `gorm:"index;not null"`
	UserID     string    `gorm:"index;not null"`
	Text       string    `gorm:"type:text;not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

// Chunks for embedding + search
type DocumentChunk struct {
	ID        string             `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	DocumentID string            `gorm:"index;not null"`
	UserID    string             `gorm:"index;not null"`
	ChunkText string             `gorm:"type:text;not null"`
	Embedding pgvector.Vector    `gorm:"type:vector;size:1536"` // pgvector-go Vector
	CreatedAt time.Time          `gorm:"autoCreateTime"`
}

// Quizzes
type Quiz struct {
	ID             string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID         string    `gorm:"index;not null"`
	TopicID        *string   `gorm:"index"`
	QuizType       string    `gorm:"type:varchar(20);check:quiz_type IN ('topic','weekly','goal')"`
	Questions      string    `gorm:"type:jsonb"` // stored as JSON
	Score          float64
	TotalQuestions int
	AttemptedAt    *time.Time
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}

// Study activity log
type StudyActivity struct {
	ID             string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID         string    `gorm:"index;not null"`
	TopicID        *string   `gorm:"index"`
	ActivityType   string    `gorm:"type:varchar(20);check:activity_type IN ('reading','quiz','flashcard','chat')"`
	DurationMinutes int
	Data           string    `gorm:"type:jsonb"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}

// Chat messages (user question + answer)
type ChatMessage struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID    string    `gorm:"index;not null"`
	Question  string    `gorm:"type:text;not null"`
	Answer    string    `gorm:"type:text;not null"`
	Sources   string    `gorm:"type:jsonb"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

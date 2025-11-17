package controllers

import (
	"net/http"
	"skillup-backend/db"
	"skillup-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UploadDocument accepts multipart form "file"
func UploadDocument(c *gin.Context) {
	userId := c.GetString("user_id")

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing file"})
		return
	}
	defer file.Close()

	data, err := services.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not read file"})
		return
	}

	doc := db.Document{
		ID:       uuid.NewString(),
		UserID:   userId,
		Filename: header.Filename,
	}
	db.DB.Create(&doc)

	// Extract text synchronously (MVP)
	text := services.PDFToText(data)

	raw := db.DocumentRaw{
		ID:         uuid.NewString(),
		DocumentID: doc.ID,
		UserID:     userId,
		Text:       text,
	}
	db.DB.Create(&raw)

	// Chunk + embed + store
	chunks := services.ChunkText(text)
	for _, ch := range chunks {
		emb := services.GetEmbedding(ch)
		db.DB.Create(&db.DocumentChunk{
			ID:         uuid.NewString(),
			DocumentID: doc.ID,
			UserID:     userId,
			ChunkText:  ch,
			Embedding:  emb,
		})
	}

	// mark processed
	doc.ProcessingStatus = "processed"
	db.DB.Save(&doc)

	c.JSON(http.StatusOK, gin.H{"status": "ok", "document_id": doc.ID})
}

func GetDocuments(c *gin.Context) {
	userId := c.GetString("user_id")
	var docs []db.Document
	db.DB.Where("user_id = ?", userId).Find(&docs)
	c.JSON(http.StatusOK, docs)
}

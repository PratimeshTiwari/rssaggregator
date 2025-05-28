package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/pratimeshtiwari/rssaggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateCourseEnrolled(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		CourseID uuid.UUID `json:"course_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid request body: %v", err))
		return
	}

	coursesEnrolled, err := apiCfg.DB.CreateCoursesEnrolled(r.Context(), database.CreateCoursesEnrolledParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		CourseID:  params.CourseID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error enroll to course failed: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseCourseEnrolledToCourseEnrolled(coursesEnrolled))
}

func (apiCfg *apiConfig) handlerGetCoursesEnrolled(w http.ResponseWriter, r *http.Request, user database.User) {

	coursesEnrolled, err := apiCfg.DB.GetCoursesEnrolled(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error couldn't get courses enrolled: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseCoursesEnrolledToCoursesEnrolled(coursesEnrolled))
}

func (apiCfg *apiConfig) handlerDeleteCourseEnrolled(w http.ResponseWriter, r *http.Request, user database.User) {
	courseIDStr := chi.URLParam(r, "courseID")
	courseID, err := uuid.Parse(courseIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid course ID: %v", err))
		return
	}
	err = apiCfg.DB.DeleteCoursesEnrolled(r.Context(), database.DeleteCoursesEnrolledParams{
		ID:     courseID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error deleting course enrolled: %v", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})
}

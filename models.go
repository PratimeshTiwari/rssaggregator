package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/pratimeshtiwari/rssaggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

type Course struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type CoursesEnrolled struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	CourseID  uuid.UUID `json:"course_id"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

func databaseCourseToCourse(dbCourse database.Course) Course {
	return Course{
		ID:        dbCourse.ID,
		CreatedAt: dbCourse.CreatedAt,
		UpdatedAt: dbCourse.UpdatedAt,
		Name:      dbCourse.Name,
		Url:       dbCourse.Url,
		UserID:    dbCourse.UserID,
	}
}

func databaseCoursesToCourses(dbCourses []database.Course) []Course {

	courses := []Course{}
	for _, dbCourse := range dbCourses {
		courses = append(courses, databaseCourseToCourse(dbCourse))
	}
	return courses
}

func databaseCourseEnrolledToCourseEnrolled(dbCourseEnrolled database.CoursesEnrolled) CoursesEnrolled {
	return CoursesEnrolled{
		ID:        dbCourseEnrolled.ID,
		CreatedAt: dbCourseEnrolled.CreatedAt,
		UpdatedAt: dbCourseEnrolled.UpdatedAt,
		UserID:    dbCourseEnrolled.UserID,
		CourseID:  dbCourseEnrolled.CourseID,
	}
}

func databaseCoursesEnrolledToCoursesEnrolled(dbCoursesEnrolled []database.CoursesEnrolled) []CoursesEnrolled {

	coursesEnrolled := []CoursesEnrolled{}
	for _, dbCourseEnrolled := range dbCoursesEnrolled {
		coursesEnrolled = append(coursesEnrolled, databaseCourseEnrolledToCourseEnrolled(dbCourseEnrolled))
	}
	return coursesEnrolled
}

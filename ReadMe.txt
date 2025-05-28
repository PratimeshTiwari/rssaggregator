Admission Management System

A backend system designed to manage students, courses, and their admissions.

Overview

This project implements a RESTful API server for an Admission Management System. It allows for:

* Student registration and profile management
* Course creation and listing
* Enrollment of students into courses
* Administrative functionalities

The system is built with scalability and maintainability in mind :

Features

* Student Management: Sign-up, login, profile update, and deletion
* Course Management: Create, read, update, and delete courses
* Enrollment: Enroll students into courses and view enrollments
* Authentication: Secure endpoints with token-based authentication
* Error Handling: Comprehensive error responses for client and server errors
* Validation: Input validation to ensure data integrity

Technologies Used

* Language**: Go (Golang)
* Framework**: net/http
* Database**: PostgreSQL
* Authentication**: API Tokens (sha256 randomised)
* Version Control**: Git

Setup Instructions

Prerequisites

* Go (1.20+ recommended)
* PostgreSQL

Steps

1.Clone the Repository

   ```bash
   git clone https://github.com/PratimeshTiwari/rssaggregator.git
   cd rssaggregator
   ```

2. Configure Environment Variables

   Create a `.env` file in the root directory:

   ```env
   DB_SOURCE=postgresql://user:password@localhost:5432/admission_db?sslmode=disable
   ```

3. Initialize the Database**

   Create a PostgreSQL database: `admission_db`
   Run schema and seed SQL scripts if provided

4. Run the Application**

   ```bash
   go run main.go
   ```

   The API server will be accessible at `http://localhost:8080`.

API Documentation

### Authentication

Login

  * `POST /auth/login`

Students

* Register: `POST /students`
* Get: `GET /students/{id}`
* Update: `PUT /students/{id}`
* Delete: `DELETE /students/{id}`

Courses

* Create: `POST /courses`
* List: `GET /courses`
* Get: `GET /courses/{id}`
* Update: `PUT /courses/{id}`
* Delete: `DELETE /courses/{id}`

Enrollments

* Enroll: `POST /enrollments`
* List: `GET /enrollments`



Assumptions

* Each student has a unique email.
* A student can enroll in multiple courses.
* Auth is required for most operations.

Additional Stuff : 
Postman is used for testing and no endpoints that require auth passes without proper api key
Token Based Endpoint Support 
Input Validation is taken care


Deployment

1. Set up PostgreSQL and update `.env`.
2. Run the Go server using:

   ```bash
   go build && ./file #here file = ./resaggregator 
   ```

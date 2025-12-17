package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"session-18/model"
	"session-18/service"
	"session-18/utils"
	"strconv"
	"strings"
)

type AssignmentHandler struct {
	AssignmentService service.AssignmentService
}

func NewAssignmentHandler(assignmenetService service.AssignmentService) AssignmentHandler {
	return AssignmentHandler{
		AssignmentService: assignmenetService,
	}
}

func (AssignmentHandler *AssignmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var assignment model.Assignment
	if err := json.NewDecoder(r.Body).Decode(&assignment); err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "erro data", nil)
		return
	}
	// // validation

	// create assignment service
	err := AssignmentHandler.AssignmentService.Create(&assignment)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success created assignment", nil)
}

func (AssignmentHandler *AssignmentHandler) List(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid page", nil)
		return
	}

	// config limit pagination
	limit := 3

	// Get data assignment form service all assignment
	assignments, pagination, err := AssignmentHandler.AssignmentService.GetAllAssignments(page, limit)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to fetch assignments: "+err.Error(), nil)
		return
	}

	utils.ResponsePagination(w, http.StatusOK, "success get data", assignments, *pagination)

}

func (AssignmentHandler *AssignmentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// assignmentIDstr := chi.URLParam(r, "assignment_id")
	// conversi to int
}

func (AssignmentHandler *AssignmentHandler) Update(w http.ResponseWriter, r *http.Request) {
	// assignmentIDstr := chi.URLParam(r, "assignment_id")

	// var assignment model.Assignment
	// if err := json.NewDecoder(r.Body).Decode(&assignment); err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, "erro data", nil)
	// 	return
	// }
	// assignments, err := AssignmentHandler.AssignmentService.GetAllAssignments()
	// if err != nil {
	// 	return
	// }

}

func (AssignmentHandler *AssignmentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// assignmentIDstr := chi.URLParam(r, "assignment_id")
	// assignments, err := AssignmentHandler.AssignmentService.GetAllAssignments()
	// if err != nil {
	// 	return
	// }

}

func (AssignmentHandler *AssignmentHandler) SubmitAssignment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "error file size", http.StatusBadRequest)
			return
		}
	}

	// get assignment id
	assignmentID, err := strconv.Atoi(r.FormValue("assignment_id"))
	if err != nil {
		http.Error(w, "Invalid assignment ID", http.StatusBadRequest)
		return
	}

	// get student id
	c, _ := r.Cookie("session")
	idStr := strings.TrimPrefix(c.Value, "lumos-")
	studentID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	// get file
	file, fileHeander, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "error file", http.StatusBadRequest)
		return
	}

	status, err := AssignmentHandler.AssignmentService.SubmitAssignment(studentID, assignmentID, file, fileHeander)
	if err != nil {
		http.Error(w, "error submit", http.StatusBadRequest)
		return
	}

	fmt.Println(status)
	http.Redirect(w, r, "/user/success-submit", http.StatusSeeOther)
}

// func (AssignmentHandler *AssignmentHandler) SuccessSubmit(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	if err := AssignmentHandler.Templates.ExecuteTemplate(w, "success_submit", nil); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samandar2605/practice_api/api/models"
	"github.com/samandar2605/practice_api/storage/repo"
)

// @Router /students [post]
// @Summary Create a student
// @Description Create a student
// @Tags students
// @Accept json
// @Produce json
// @Param student body []models.CreateStudent true "student"
// @Success 201 {object} models.ResponseOK
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateStudent(c *gin.Context) {
	var (
		req []models.Student
	)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	var students []*repo.Student

	for _, i := range req {
		students = append(students, &repo.Student{
			Id:          i.Id,
			FirstName:   i.FirstName,
			LastName:    i.LastName,
			UserName:    i.UserName,
			Email:       i.Email,
			PhoneNumber: i.PhoneNumber,
			CreatedAt:   i.CreatedAt,
		})
	}

	err = h.storage.Students().Create(students)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Barakalla!!!\nYugirib kelib Kalla qo'yding ;)",
	})
}

// @Router /students [get]
// @Summary Get all students
// @Description Get all students
// @Tags students
// @Accept json
// @Produce json
// @Param filter query models.GetAllParams false "Filter"
// @Success 200 {object} models.GetAllResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllStudent(c *gin.Context) {
	req, err := studentsParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.storage.Students().GetAll(repo.GetStudentsQuery{
		Page:       req.Page,
		Limit:      req.Limit,
		Search:     req.Search,
		SortByDate: req.SortByDate,
		SortByName: req.SortByName,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, studentsResponse(result))
}

func studentsParams(c *gin.Context) (*models.GetAllParams, error) {
	var (
		limit      int = 10
		page       int = 1
		err        error
		sortByDate string
		SortByName string
	)

	if c.Query("limit") != "" {
		limit, err = strconv.Atoi(c.Query("limit"))
		if err != nil {
			return nil, err
		}
	}

	if c.Query("page") != "" {
		page, err = strconv.Atoi(c.Query("page"))
		if err != nil {
			return nil, err
		}
	}

	if c.Query("sort_by_date") != "" &&
		(c.Query("sort_by_date") == "desc" || c.Query("sort_by_date") == "asc" || c.Query("sort_by_date") == "none") {
		sortByDate = c.Query("sort_by_date")
	}

	if c.Query("sort_by_name") != "" &&
		(c.Query("sort_by_name") == "asc" || c.Query("sort_by_name") == "desc" || c.Query("sort_by_name") == "none"){
		SortByName = c.Query("sort_by_name")
	}

	return &models.GetAllParams{
		Limit:      limit,
		Page:       page,
		Search:     c.Query("search"),
		SortByDate: sortByDate,
		SortByName: SortByName,
	}, nil
}

func studentsResponse(data *repo.GetAllStudentsResult) *models.GetAllResponse {
	response := models.GetAllResponse{
		Students: make([]*models.Student, 0),
		Count:    data.Count,
	}

	for _, post := range data.Students {
		p := parsePostModel(post)
		response.Students = append(response.Students, &p)
	}

	return &response
}

func parsePostModel(student *repo.Student) models.Student {
	return models.Student{
		Id:          student.Id,
		FirstName:   student.FirstName,
		LastName:    student.LastName,
		UserName:    student.UserName,
		Email:       student.Email,
		PhoneNumber: student.PhoneNumber,
		CreatedAt:   student.CreatedAt,
	}
}

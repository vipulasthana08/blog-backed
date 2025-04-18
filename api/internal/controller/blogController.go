package blogcontroller

import (
	structure "blog-backend/api/internal/config"
	blogservice "blog-backend/api/internal/service"
	"blog-backend/api/internal/utils"
	"encoding/json"
	"net/http"
	"regexp"
	"time"
)

func CreateNewBlog(w http.ResponseWriter, r *http.Request) {
	if utils.ThrowMethodNotAllowedError(w, r, http.MethodPost) {
		return
	}

	op := "CreateNewBlog"

	var payload structure.NewBlog
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.CreateResponse(w, http.StatusBadRequest, "Invalid Request Body", op)
		return
	}

	reg := regexp.MustCompile(`\s+`)
	if reg.ReplaceAllString(payload.Title, "") == "" {
		utils.CreateResponse(w, http.StatusBadRequest, "Title can not be empty.", op)
		return
	} else if reg.ReplaceAllString(payload.Description, "") == "" {
		utils.CreateResponse(w, http.StatusBadRequest, "Description can not be empty.", op)
		return
	} else if (payload.Date == time.Time{}) {
		utils.CreateResponse(w, http.StatusBadRequest, "Date can not be empty.", op)
		return
	}

	writeErr := blogservice.WriteNewBlog(payload, op)
	if writeErr != nil {
		utils.CreateResponse(w, http.StatusInternalServerError, writeErr.Message, writeErr.Op)
		return
	}
	utils.CreateResponse(w, http.StatusOK, "")
}

func CheckBlogExist(w http.ResponseWriter, r *http.Request) {
	if utils.ThrowMethodNotAllowedError(w, r, http.MethodGet) {
		return
	}

	op := "CheckBlogExistForDate"

	dateToday, errExist := getDateFromQueryParams(w, r, op)
	if errExist {
		return
	}

	blogservice.CheckForBlogFileExist(w, dateToday, op)
}

func GetBlog(w http.ResponseWriter, r *http.Request) {
	if utils.ThrowMethodNotAllowedError(w, r, http.MethodGet) {
		return
	}

	op := "GetBlog"

	dateToday, errExist := getDateFromQueryParams(w, r, op)
	if errExist {
		return
	}

	blogservice.GetBlogForDate(w, dateToday, op)
}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	if utils.ThrowMethodNotAllowedError(w, r, http.MethodDelete) {
		return
	}

	op := "DeleteBlog"

	dateToday, errExist := getDateFromQueryParams(w, r, op)
	if errExist {
		return
	}

	blogservice.DeleteBlogForDate(w, dateToday, op)
}

func getDateFromQueryParams(resWriter http.ResponseWriter, req *http.Request, op string) (string, bool) {
	queryParams := req.URL.Query()
	date, err := utils.ParseFlexibleDate(queryParams.Get("date"))
	if err != nil {
		utils.CreateResponse(resWriter, http.StatusBadRequest, err.Error(), op)
		return "", true
	}
	return date, false
}

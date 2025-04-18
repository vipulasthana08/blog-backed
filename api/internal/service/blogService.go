package blogservice

import (
	structure "blog-backend/api/internal/config"
	"blog-backend/api/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func WriteNewBlog(payload structure.NewBlog, op string) *structure.Error {
	fileName := fmt.Sprintf("%s.txt", payload.Date.Format("2006-01-02"))

	file, err := os.Create(fileName)
	if err != nil {
		resErr := structure.Error{
			Message: "Error in Creating file.",
			Op:      op,
			Data:    err.Error(),
		}
		return &resErr
	}

	content := fmt.Sprintf("Title: %s\n\nDescription: %s", payload.Title, payload.Description)
	_, err = file.WriteString(content)
	if err != nil {
		resErr := structure.Error{
			Message: "Error in writing in file.",
			Op:      op,
			Data:    err.Error(),
		}
		return &resErr
	}
	defer file.Close()
	return nil
}

func CheckForBlogFileExist(w http.ResponseWriter, date string, op string) {
	fileName := fmt.Sprintf("%s.txt", date)

	fileDetails, err := os.Stat(fileName)
	if fileDetails == nil || os.IsNotExist(err) {
		utils.CreateResponse(w, http.StatusOK, "")
		return
	}
	if err != nil {
		utils.CreateResponse(w, http.StatusInternalServerError, "Error occurred while fetching blog", op)
		return
	}
	utils.CreateResponse(w, http.StatusConflict, "Blog already exists", op)
}

func GetBlogForDate(w http.ResponseWriter, date string, op string) {
	fileName := fmt.Sprintf("%s.txt", date)

	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		utils.CreateResponse(w, http.StatusInternalServerError, "Error reading file", op)
		return
	}

	lines := strings.Split(string(fileContent), "\n")

	var title, description string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Title:") {
			title = strings.TrimPrefix(line, "Title: ")
		}
		if strings.HasPrefix(line, "Description:") {
			description = strings.TrimPrefix(line, "Description: ")
		}
	}
	res := structure.ResBlog{
		Title:       title,
		Description: description,
	}
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func DeleteBlogForDate(w http.ResponseWriter, date string, op string) {
	fileName := fmt.Sprintf("%s.txt", date)

	err := os.Remove(fileName)
	if err != nil {
		utils.CreateResponse(w, http.StatusInternalServerError, "Error deleting blog", op)
		return
	}
	utils.CreateResponse(w, http.StatusOK, "")
}

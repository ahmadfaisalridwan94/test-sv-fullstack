package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"backend/api/models"
	"backend/api/responses"
	"backend/api/utils/formaterror"
	"github.com/gorilla/mux"
)

func (server *Server) CreatePost(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post := models.Post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	post.Prepare()
	errors := post.Validate()
	if len(errors) > 0 {
		responses.JSON(w, http.StatusUnprocessableEntity, errors)
		return
	}
	
	postCreated, err := post.SavePost(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, postCreated.Id))
	responses.JSON(w, http.StatusCreated, postCreated)
}

func (server *Server) GetPosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	limitStr := vars["limit"]
	offsetStr := vars["offset"]
	status := r.URL.Query().Get("status")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	post := models.Post{}

	var statusPtr *string
	if status != "" {
		statusPtr = &status
	}

	posts, err := post.FindAllPosts(server.DB, limit, offset, statusPtr)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

func (server *Server) GetPost(w http.ResponseWriter, r *http.Request) {

	fmt.Printf(r.Method)
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		log.Printf("Error parsing post ID: %v", err)
		response := responses.Response{
			ResponseCode:    "01",
			ResponseStatus:  false,
			ResponseMessage: "Invalid post ID",
		}
		responses.JSON(w, http.StatusBadRequest, response)
		//http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	if pid == 0 {
		log.Printf("Error: post ID is empty")
		response := responses.Response{
			ResponseCode:    "02",
			ResponseStatus:  false,
			ResponseMessage: "Empty post ID",
		}
		responses.JSON(w, http.StatusBadRequest, response)
		//http.Error(w, "Empty post ID", http.StatusBadRequest)
		return
	}
	post := models.Post{}
	postReceived, err := post.FindPostById(server.DB, pid)
	if err != nil {
		log.Printf("Error retrieving post: %v", err)
		response := responses.Response{
			ResponseCode:    "03",
			ResponseStatus:  false,
			ResponseMessage: "Failed to retrieve post",
		}
		responses.JSON(w, http.StatusInternalServerError, response)
		//http.Error(w, "Failed to retrieve post", http.StatusInternalServerError)
		return
	}
	

	responses.JSON(w, http.StatusOK, postReceived)
}

func (server *Server) UpdatePost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the post id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Check if the post exist
	post := models.Post{}
	err = server.DB.Debug().Model(models.Post{}).Where("id = ?", pid).Take(&post).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Post not found"))
		return
	}
	
	// Read the data posted
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	postUpdate := models.Post{}
	err = json.Unmarshal(body, &postUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	
	postUpdate.Prepare()
	errors := postUpdate.Validate()
	if len(errors) > 0 {
		responses.JSON(w, http.StatusUnprocessableEntity, errors)
		return
	}

	postUpdate.Id = post.Id //this is important to tell the model the post id to update, the other update field are set above

	postUpdated, err := postUpdate.UpdateAPost(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, postUpdated)
}

func (server *Server) DeletePost(w http.ResponseWriter, r *http.Request) {

	fmt.Printf(r.Method)
	vars := mux.Vars(r)

	// Is a valid post id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}	

	// Check if the post exist
	post := models.Post{}
	err = server.DB.Debug().Model(models.Post{}).Where("id = ?", pid).Take(&post).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Post doesn't exist"))
		return
	}

	// _, err = post.DeleteAPost(server.DB, pid)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusBadRequest, err)
	// 	return
	// }

	// Update the post status to 'trash'
	post.Status = "trash"
	err = server.DB.Debug().Save(&post).Error
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	response := responses.Response{
		ResponseCode:    "00",
		ResponseStatus:  true,
		ResponseMessage: "Post moved to trash!",
	}
	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCountPosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status := vars["status"]

	post := models.Post{}

	var statusPtr *string
	if status != "" {
		statusPtr = &status
	}

	c, err := post.CountData(server.DB, statusPtr)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response := responses.Response{
		Data: c,
	}
	responses.JSON(w, http.StatusOK, response)
}


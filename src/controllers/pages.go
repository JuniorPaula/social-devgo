package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}
	utils.Render(w, "login.html", nil)
}

func SignupPage(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, "signup.html", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.APIURL)
	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		responses.ResponseJON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		responses.VerifyStatusCodeErrors(w, resp)
		return
	}

	var posts []models.PostDTO
	if err = json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		responses.ResponseJON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.Render(w, "home.html", struct {
		Posts  []models.PostDTO
		UserID uint64
	}{
		Posts:  posts,
		UserID: userID,
	})
}

func EditPostPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		responses.ResponseJON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.APIURL, postID)
	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		responses.ResponseJON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		responses.VerifyStatusCodeErrors(w, resp)
	}

	var post models.PostDTO
	if err = json.NewDecoder(resp.Body).Decode(&post); err != nil {
		responses.ResponseJON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}

	utils.Render(w, "edit-post.html", post)
}

func FindUsersPage(w http.ResponseWriter, r *http.Request) {
	nameOrNickname := strings.ToLower(r.URL.Query().Get("user"))
	url := fmt.Sprintf("%s/users?user=%s", config.APIURL, nameOrNickname)

	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		responses.ResponseJON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		responses.VerifyStatusCodeErrors(w, resp)
		return
	}

	var users []models.UserDTO
	if err = json.NewDecoder(resp.Body).Decode(&users); err != nil {
		responses.ResponseJON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}

	utils.Render(w, "users.html", users)
}

func UserProfilePage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.ResponseJON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userLoggedID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if userID == userLoggedID {
		http.Redirect(w, r, "/profile", http.StatusFound)
	}

	user, err := models.FindUser(userID, r)
	if err != nil {
		responses.ResponseJON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}

	utils.Render(w, "user.html", struct {
		User         models.UserDTO
		UserLoggedID uint64
	}{
		User:         user,
		UserLoggedID: userLoggedID,
	})
}

func UserLoggedProfilePage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, err := models.FindUser(userID, r)
	if err != nil {
		responses.ResponseJON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}

	utils.Render(w, "profile.html", user)
}

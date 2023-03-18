package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requests"
)

type UserDTO struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"createdAt"`
	Followers []UserDTO `json:"followers"`
	Following []UserDTO `json:"following"`
	Posts     []PostDTO `json:"posts"`
}

func FindUser(userID uint64, r *http.Request) (UserDTO, error) {
	userChannel := make(chan UserDTO)
	followersChannel := make(chan []UserDTO)
	followingChannel := make(chan []UserDTO)
	postsChannel := make(chan []PostDTO)

	go GetUserData(userChannel, userID, r)
	go GetFollowers(followersChannel, userID, r)
	go GetFollowing(followingChannel, userID, r)
	go GetPosts(postsChannel, userID, r)

	var (
		user      UserDTO
		followers []UserDTO
		following []UserDTO
		posts     []PostDTO
	)

	for i := 0; i < 4; i++ {
		select {
		case userResp := <-userChannel:
			if userResp.ID == 0 {
				return UserDTO{}, errors.New("user not found")
			}
			user = userResp
		case followersResp := <-followersChannel:
			followers = followersResp
		case followingsResp := <-followingChannel:
			following = followingsResp
		case postsResp := <-postsChannel:
			posts = postsResp
		}
	}

	user.Followers = followers
	user.Following = following
	user.Posts = posts

	return user, nil
}

func GetUserData(channel chan<- UserDTO, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- UserDTO{}
		return
	}
	defer resp.Body.Close()

	var user UserDTO
	if err = json.NewDecoder(resp.Body).Decode(&user); err != nil {
		channel <- UserDTO{}
		return
	}

	channel <- user
}

func GetFollowers(channel chan<- []UserDTO, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.APIURL, userID)
	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer resp.Body.Close()

	var userFollowers []UserDTO
	if err = json.NewDecoder(resp.Body).Decode(&userFollowers); err != nil {
		channel <- nil
		return
	}

	channel <- userFollowers
}

func GetFollowing(channel chan<- []UserDTO, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.APIURL, userID)
	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer resp.Body.Close()

	var userFollowings []UserDTO
	if err = json.NewDecoder(resp.Body).Decode(&userFollowings); err != nil {
		channel <- nil
		return
	}

	channel <- userFollowings
}

func GetPosts(channel chan<- []PostDTO, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/posts", config.APIURL, userID)
	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer resp.Body.Close()

	var userPosts []PostDTO
	if err = json.NewDecoder(resp.Body).Decode(&userPosts); err != nil {
		channel <- nil
		return
	}

	channel <- userPosts
}

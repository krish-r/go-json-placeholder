package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Post struct {
	UserId int    `json:"user_id,omitempty"`
	Id     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
}

type PostPatch struct {
	UserId *int    `json:"user_id,omitempty"`
	Title  *string `json:"title,omitempty"`
	Body   *string `json:"body,omitempty"`
}

type Posts []Post

func (p Posts) search(id int) (int, *Post, error) {
	for i := 0; i < len(p); i++ {
		if p[i].Id == id {
			return i, &p[i], nil
		}
	}
	return -1, nil, errors.New("id not found")
}

func (p Posts) searchPosts(userId int) Posts {
	posts := Posts{}
	for i := 0; i < len(p); i++ {
		if p[i].UserId == userId {
			posts = append(posts, p[i])
		}
	}
	return posts
}

func listAllPosts(w http.ResponseWriter, r *http.Request) {
	err := writeJSON(w, http.StatusOK, posts)
	check(err)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	post := &Post{}
	err := readJSON(r.Body, post)
	if err != nil {
		err = writeJSON(w, http.StatusBadRequest, err)
		check(err)
		return
	}
	defer r.Body.Close()

	post.Id = <-postIdChan

	posts = append(posts, *post)
	err = writeJSON(w, http.StatusCreated, post)
	check(err)
}

func getPostIdFromRequest(w http.ResponseWriter, r *http.Request) (int, error) {
	cId, err := strconv.Atoi(chi.URLParam(r, "postId"))
	if err != nil {
		return -1, err
	}
	return cId, nil
}

func searchPost(w http.ResponseWriter, r *http.Request) (int, *Post, error) {
	id, err := getPostIdFromRequest(w, r)
	if err != nil {
		err := writeJSON(w, http.StatusNotAcceptable, err)
		return -1, nil, err
	}

	idx, post, err := posts.search(id)
	if err != nil {
		err := writeJSON(w, http.StatusNotFound, err)
		return -1, nil, err
	}
	return idx, post, nil
}

func getPost(w http.ResponseWriter, r *http.Request) {
	idx, post, err := searchPost(w, r)
	if err != nil || idx == -1 {
		return
	}

	err = writeJSON(w, http.StatusOK, post)
	check(err)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	post := &Post{}
	err := readJSON(r.Body, post)
	if err != nil {
		err = writeJSON(w, http.StatusBadRequest, err)
		check(err)
	}
	defer r.Body.Close()

	idx, originalPost, err := searchPost(w, r)
	if err != nil || idx == -1 {
		return
	}

	originalPost.UserId = post.UserId
	originalPost.Title = post.Title
	originalPost.Body = post.Body

	err = writeJSON(w, http.StatusOK, originalPost)
	check(err)
}

func editPost(w http.ResponseWriter, r *http.Request) {
	post := &PostPatch{}
	err := readJSON(r.Body, post)
	if err != nil {
		err = writeJSON(w, http.StatusBadRequest, err)
		check(err)
	}
	defer r.Body.Close()

	idx, originalPost, err := searchPost(w, r)
	if err != nil || idx == -1 {
		return
	}

	if post.UserId != nil {
		originalPost.UserId = *post.UserId
	}
	if post.Title != nil {
		originalPost.Title = *post.Title
	}
	if post.Body != nil {
		originalPost.Body = *post.Body
	}

	err = writeJSON(w, http.StatusOK, originalPost)
	check(err)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	idx, post, err := searchPost(w, r)
	if err != nil || idx == -1 {
		return
	}
	deletedPost := *post

	posts = append(posts[0:idx], posts[idx+1:]...)
	err = writeJSON(w, http.StatusOK, deletedPost)
	check(err)
}

func getPostComments(w http.ResponseWriter, r *http.Request) {
	comments := searchPostComments(w, r, -1)

	err := writeJSON(w, http.StatusOK, comments)
	check(err)
}

func searchPostComments(w http.ResponseWriter, r *http.Request, id int) Comments {
	if id == -1 {
		var err error
		id, err = getPostIdFromRequest(w, r)
		if err != nil {
			err := writeJSON(w, http.StatusNotAcceptable, err)
			check(err)
		}
	}

	comments := comments.searchPosts(id)
	return comments
}

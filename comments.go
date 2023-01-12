package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Comment struct {
	PostId int    `json:"post_id,omitempty"`
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Body   string `json:"body,omitempty"`
}

type CommentPatch struct {
	PostId *int    `json:"post_id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Email  *string `json:"email,omitempty"`
	Body   *string `json:"body,omitempty"`
}

type Comments []Comment

func (c Comments) search(id int) (int, *Comment, error) {
	for i := 0; i < len(c); i++ {
		if c[i].Id == id {
			return i, &c[i], nil
		}
	}
	return -1, nil, errors.New("id not found")
}

func (c Comments) searchPosts(postId int) Comments {
	comments := Comments{}
	for i := 0; i < len(c); i++ {
		if c[i].PostId == postId {
			comments = append(comments, c[i])
		}
	}
	return comments
}

func listAllComments(w http.ResponseWriter, r *http.Request) {
	postId := r.URL.Query().Get("postId")
	if postId == "" {
		err := writeJSON(w, http.StatusOK, comments)
		check(err)
		return
	}

	id, err := strconv.Atoi(postId)
	if err != nil {
		err := writeJSON(w, http.StatusNotAcceptable, err)
		check(err)
		return
	}
	filteredComments := searchPostComments(w, r, id)

	err = writeJSON(w, http.StatusOK, filteredComments)
	check(err)
}

func createComment(w http.ResponseWriter, r *http.Request) {
	comment := &Comment{}
	err := readJSON(r.Body, comment)
	if err != nil {
		err = writeJSON(w, http.StatusBadRequest, err)
		check(err)
		return
	}
	defer r.Body.Close()

	comment.Id = <-commentIdChan

	comments = append(comments, *comment)
	err = writeJSON(w, http.StatusCreated, comment)
	check(err)
}

// Approach suggested in chi documentation, using a context to retrieve and store the value
type ctxKey string

func commentCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idx, comment, err := searchComment(w, r)
		if err != nil || idx == -1 {
			return
		}

		ctx := context.WithValue(r.Context(), ctxKey("comment"), comment)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getCommentFromCtx(w http.ResponseWriter, r *http.Request) *Comment {
	ctx := r.Context()
	comment, ok := ctx.Value(ctxKey("comment")).(*Comment)
	if !ok {
		err := writeJSON(w, http.StatusNotFound, "Not Found!")
		check(err)
	}
	return comment
}

//

func getCommentIdFromRequest(w http.ResponseWriter, r *http.Request) (int, error) {
	cId, err := strconv.Atoi(chi.URLParam(r, "commentId"))
	if err != nil {
		return -1, err
	}
	return cId, nil
}

func searchComment(w http.ResponseWriter, r *http.Request) (int, *Comment, error) {
	id, err := getCommentIdFromRequest(w, r)
	if err != nil {
		err := writeJSON(w, http.StatusNotAcceptable, err)
		return -1, nil, err
	}

	idx, comment, err := comments.search(id)
	if err != nil {
		err := writeJSON(w, http.StatusNotFound, err)
		return -1, nil, err
	}
	return idx, comment, nil
}

func getComment(w http.ResponseWriter, r *http.Request) {
	idx, comment, err := searchComment(w, r)
	if err != nil || idx == -1 {
		return
	}

	err = writeJSON(w, http.StatusOK, comment)
	check(err)
}

func updateComment(w http.ResponseWriter, r *http.Request) {
	comment := &Comment{}
	err := readJSON(r.Body, comment)
	if err != nil {
		err = writeJSON(w, http.StatusBadRequest, err)
		check(err)
	}
	defer r.Body.Close()

	idx, originalComment, err := searchComment(w, r)
	if err != nil || idx == -1 {
		return
	}

	originalComment.PostId = comment.PostId
	originalComment.Name = comment.Name
	originalComment.Email = comment.Email
	originalComment.Body = comment.Body

	err = writeJSON(w, http.StatusOK, originalComment)
	check(err)
}

func editComment(w http.ResponseWriter, r *http.Request) {
	comment := &CommentPatch{}
	err := readJSON(r.Body, comment)
	if err != nil {
		err = writeJSON(w, http.StatusBadRequest, err)
		check(err)
	}
	defer r.Body.Close()

	idx, originalComment, err := searchComment(w, r)
	if err != nil || idx == -1 {
		return
	}

	if comment.PostId != nil {
		originalComment.PostId = *comment.PostId
	}
	if comment.Name != nil {
		originalComment.Name = *comment.Name
	}
	if comment.Email != nil {
		originalComment.Email = *comment.Email
	}
	if comment.Body != nil {
		originalComment.Body = *comment.Body
	}

	err = writeJSON(w, http.StatusOK, originalComment)
	check(err)
}

func deleteComment(w http.ResponseWriter, r *http.Request) {
	idx, comment, err := searchComment(w, r)
	if err != nil || idx == -1 {
		return
	}
	deletedComment := *comment

	comments = append(comments[0:idx], comments[idx+1:]...)
	err = writeJSON(w, http.StatusOK, deletedComment)
	check(err)
}

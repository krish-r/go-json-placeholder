package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type User struct {
	Id       int         `json:"id,omitempty"`
	Name     string      `json:"name,omitempty"`
	Username string      `json:"username,omitempty"`
	Email    string      `json:"email,omitempty"`
	Address  UserAddress `json:"address,omitempty"`
	Phone    string      `json:"phone,omitempty"`
	Website  string      `json:"website,omitempty"`
	Company  UserCompany `json:"company,omitempty"`
}

type UserAddress struct {
	Street  string `json:"street,omitempty"`
	Suite   string `json:"suite,omitempty"`
	City    string `json:"city,omitempty"`
	Zipcode string `json:"zipcode,omitempty"`
	Geo     Geo    `json:"geo,omitempty"`
}

type Geo struct {
	Lat string `json:"lat,omitempty"`
	Lng string `json:"lng,omitempty"`
}

type UserCompany struct {
	Name        string `json:"name,omitempty"`
	CatchPhrase string `json:"catch_prase,omitempty"`
	Bs          string `json:"bs,omitempty"`
}

type UserPatch struct {
	Id       *int              `json:"id,omitempty"`
	Name     *string           `json:"name,omitempty"`
	Username *string           `json:"username,omitempty"`
	Email    *string           `json:"email,omitempty"`
	Address  *UserAddressPatch `json:"address,omitempty"`
	Phone    *string           `json:"phone,omitempty"`
	Website  *string           `json:"website,omitempty"`
	Company  *UserCompanyPatch `json:"company,omitempty"`
}

type UserAddressPatch struct {
	Street  *string   `json:"street,omitempty"`
	Suite   *string   `json:"suite,omitempty"`
	City    *string   `json:"city,omitempty"`
	Zipcode *string   `json:"zipcode,omitempty"`
	Geo     *GeoPatch `json:"geo,omitempty"`
}

type GeoPatch struct {
	Lat *string `json:"lat,omitempty"`
	Lng *string `json:"lng,omitempty"`
}

type UserCompanyPatch struct {
	Name        *string `json:"name,omitempty"`
	CatchPhrase *string `json:"catch_prase,omitempty"`
	Bs          *string `json:"bs,omitempty"`
}

type Users []User

var (
	ErrIdNotFound = errors.New("id not found")
)

func (u Users) search(id int) (int, *User, error) {
	for i := 0; i < len(u); i++ {
		if u[i].Id == id {
			return i, &u[i], nil
		}
	}
	return -1, nil, ErrIdNotFound
}

func listAllUsers(w http.ResponseWriter, r *http.Request) {
	err := writeJSON(w, http.StatusOK, users)
	check(err)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	err := readJSON(r.Body, user)
	if err != nil {
		err = writeJSON(w, http.StatusBadRequest, err)
		check(err)
		return
	}
	defer r.Body.Close()

	user.Id = <-userIdChan

	users = append(users, *user)
	err = writeJSON(w, http.StatusCreated, user)
	check(err)
}

func getUserIdFromRequest(w http.ResponseWriter, r *http.Request) (int, error) {
	cId, err := strconv.Atoi(chi.URLParam(r, "userId"))
	if err != nil {
		return -1, err
	}
	return cId, nil
}

func searchUser(w http.ResponseWriter, r *http.Request) (int, *User, error) {
	id, err := getUserIdFromRequest(w, r)
	if err != nil {
		err := writeJSON(w, http.StatusNotAcceptable, err)
		return -1, nil, err
	}

	idx, user, err := users.search(id)
	if err != nil {
		err := writeJSON(w, http.StatusNotFound, err)
		return -1, nil, err
	}
	return idx, user, nil
}

func getUser(w http.ResponseWriter, r *http.Request) {
	idx, user, err := searchUser(w, r)
	if err != nil || idx == -1 {
		return
	}

	err = writeJSON(w, http.StatusOK, user)
	check(err)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	err := readJSON(r.Body, user)
	if err != nil {
		err = writeJSON(w, http.StatusBadRequest, err)
		check(err)
	}
	defer r.Body.Close()

	idx, originalUser, err := searchUser(w, r)
	if err != nil || idx == -1 {
		return
	}

	originalUser.Name = user.Name
	originalUser.Username = user.Username
	originalUser.Email = user.Email
	originalUser.Address = user.Address
	originalUser.Phone = user.Phone
	originalUser.Website = user.Website
	originalUser.Company = user.Company

	err = writeJSON(w, http.StatusOK, originalUser)
	check(err)
}

func editUser(w http.ResponseWriter, r *http.Request) {
	user := &UserPatch{}
	err := readJSON(r.Body, user)
	if err != nil {
		err = writeJSON(w, http.StatusBadRequest, err)
		check(err)
	}
	defer r.Body.Close()

	idx, originalUser, err := searchUser(w, r)
	if err != nil || idx == -1 {
		return
	}

	if user.Name != nil {
		originalUser.Name = *user.Name
	}
	if user.Username != nil {
		originalUser.Username = *user.Username
	}
	if user.Email != nil {
		originalUser.Email = *user.Email
	}
	if user.Address != nil {
		if user.Address.Street != nil {
			originalUser.Address.Street = *user.Address.Street
		}
		if user.Address.Suite != nil {
			originalUser.Address.Suite = *user.Address.Suite
		}
		if user.Address.City != nil {
			originalUser.Address.City = *user.Address.City
		}
		if user.Address.Zipcode != nil {
			originalUser.Address.Zipcode = *user.Address.Zipcode
		}
		if user.Address.Geo != nil {
			if user.Address.Geo.Lat != nil {
				originalUser.Address.Geo.Lat = *user.Address.Geo.Lat
			}
			if user.Address.Geo.Lng != nil {
				originalUser.Address.Geo.Lng = *user.Address.Geo.Lng
			}
		}
	}
	if user.Phone != nil {
		originalUser.Phone = *user.Phone
	}
	if user.Website != nil {
		originalUser.Website = *user.Website
	}
	if user.Company != nil {
		if user.Company.Name != nil {
			originalUser.Company.Name = *user.Company.Name
		}
		if user.Company.CatchPhrase != nil {
			originalUser.Company.CatchPhrase = *user.Company.CatchPhrase
		}
		if user.Company.Bs != nil {
			originalUser.Company.Bs = *user.Company.Bs
		}
	}

	err = writeJSON(w, http.StatusOK, originalUser)
	check(err)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	idx, user, err := searchUser(w, r)
	if err != nil || idx == -1 {
		return
	}
	deletedUser := *user

	users = append(users[0:idx], users[idx+1:]...)
	err = writeJSON(w, http.StatusOK, deletedUser)
	check(err)
}

func getUserPosts(w http.ResponseWriter, r *http.Request) {
	posts := searchUserPosts(w, r)

	err := writeJSON(w, http.StatusOK, posts)
	check(err)
}

func searchUserPosts(w http.ResponseWriter, r *http.Request) Posts {
	id, err := getUserIdFromRequest(w, r)
	if err != nil {
		err := writeJSON(w, http.StatusNotAcceptable, err)
		check(err)
	}

	posts := posts.searchPosts(id)
	return posts
}

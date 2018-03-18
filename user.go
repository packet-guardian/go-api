package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

type userResponse struct {
	Message string
	Data    *User
}

type User struct {
	api              *PGApi
	Username         string            `json:"username"`
	HasPassword      bool              `json:"has_password"`
	DeviceLimit      int               `json:"device_limit"`
	DeviceExpiration *DeviceExpiration `json:"device_expiration"`
	ValidForever     bool              `json:"valid_forever"`
	CanManage        bool              `json:"can_manage"`
	CanAutoreg       bool              `json:"can_autoreg"`
	ValidStart       string            `json:"valid_start"`
	ValidEnd         string            `json:"valid_end"`
	Blacklisted      bool              `json:"blacklisted"`
}

type DeviceExpiration struct {
	Mode  string `json:"mode"`
	Value int64  `json:"value"`
}

func (pg *PGApi) GetUser(username string) (*User, error) {
	resp, err := pg.makeRequest(http.MethodGet, "/api/user/%s", username)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("user not found")
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, errors.New("invalid credientials")
	}

	var userResp userResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&userResp); err != nil {
		return nil, err
	}

	user := userResp.Data
	user.api = pg
	return user, nil
}

func (u *User) Blacklist() error {
	if u.Blacklisted || u.Username == "" {
		return nil
	}

	resp, err := u.api.makeRequest(http.MethodPost, "/api/blacklist/user/%s", u.Username)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusForbidden:
		fallthrough
	case http.StatusUnauthorized:
		return errors.New("invalid credentials or permissions")
	case http.StatusInternalServerError:
		return errors.New(u.api.decodeGenericResponse(resp.Body).Message)
	}
	u.Blacklisted = true
	return nil
}

func (u *User) Unblacklist() error {
	if !u.Blacklisted || u.Username == "" {
		return nil
	}

	resp, err := u.api.makeRequest(http.MethodDelete, "/api/blacklist/user/%s", u.Username)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusForbidden:
		fallthrough
	case http.StatusUnauthorized:
		return errors.New("invalid credentials or permissions")
	case http.StatusInternalServerError:
		return errors.New(u.api.decodeGenericResponse(resp.Body).Message)
	}
	u.Blacklisted = false
	return nil
}

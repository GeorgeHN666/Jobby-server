package main

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (app *application) ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {

	maxBytes := int64(1048576)

	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		return err
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain on single json element")
	}
	return nil
}

func (app *application) WriteJSON(w http.ResponseWriter, r *http.Request, data interface{}, status int, headers ...http.Header) error {

	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for i, v := range headers[0] {
			w.Header()[i] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)

	return nil
}

func (app *application) ErrorBadRequest(w http.ResponseWriter, r *http.Request, err error) error {

	var response struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}
	response.Error = true
	response.Message = err.Error()

	out, err := json.Marshal(response)
	if err != nil {
		return err
	}

	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(out)
	return nil
}

func (app *application) EncryptPassword(pass string) (string, error) {

	cost := 14

	hash, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (app *application) CompareHash(pass string, hash string) (string, error) {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return err.Error(), err

		default:
			return "something happened", err
		}
	}

	return "all good", nil
}

func (app *application) GenerateAuthCode() (string, error) {

	lenght := 6

	chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	bytes := make([]byte, lenght)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	for i, v := range bytes {
		bytes[i] = chars[v%byte(len(chars))]
	}

	return string(bytes), nil
}

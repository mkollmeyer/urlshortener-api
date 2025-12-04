package handler

import (
	"fmt"
	"net/http"
)

type Url struct{}

func (u *Url) Gen(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Generated a url")
}
func (u *Url) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Retrieved a url")
}

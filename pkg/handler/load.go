package handler

import "net/http"

type LoadRequest struct {
	ProjectName string `json:"project_name"`
}

func LoadHandler(w http.ResponseWriter, r *http.Request) {

}

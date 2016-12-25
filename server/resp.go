package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SuperSpaceNinja/sameyourei/model"
)

/*
Resp response */
type Resp struct {
	ResponseCode int    `json:"response_code"`
	URL          string `json:"url"`
	Method       string `json:"method"`
}

/*
BResp response with body */
type BResp struct {
	Resp
	Body struct {
		BodyCount int         `json:"body_count"`
		Todos     model.Todos `json:"todos"`
	} `json:"body"`
}

/*
EResp response with error string */
type EResp struct {
	Resp
	Error string `json:"error"`
}

/*
Body response body */
type Body struct {
	BodyCount int         `json:"body_count"`
	Todos     model.Todos `json:"todos"`
}

func respt(w http.ResponseWriter, r *http.Request, statusCode int, todos model.Todos) {
	resp := Resp{
		statusCode,
		string(r.RequestURI),
		r.Method,
	}
	w.WriteHeader(statusCode)
	bresp := BResp{resp, Body{len(todos), todos}}
	if err := json.NewEncoder(w).Encode(bresp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("error: %s\n", err.Error())
		return
	}
	return
}

func respe(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	resp := Resp{
		statusCode,
		string(r.RequestURI),
		r.Method,
	}
	w.WriteHeader(statusCode)
	eresp := EResp{resp, err.Error()}
	if err = json.NewEncoder(w).Encode(eresp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("error: %s\n", err.Error())
		return
	}
	return
}

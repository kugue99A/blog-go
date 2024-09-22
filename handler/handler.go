package handler

type Status struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

var (
	StatusInternalServerError = Status{Code: 500, Status: "Internal Server Error"}
	StatusNotFound            = Status{Code: 404, Status: "Not Found"}
	StatusBadRequest          = Status{Code: 400, Status: "Bad Request"}
)

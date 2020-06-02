package utils

import "net/http"

//CheckError checks for errors and return appropriate error code
func CheckError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

package api

import (
	"fmt"
	"io"
	"net/http"
)

func ReportInvestment(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Print("ERROR reading body: ", err)
		return
	}

	_, err = w.Write([]byte("Investment Reported w/ body: " + string(bodyBytes) + "\n"))
	if err != nil {
		return
	}
	fmt.Print("Investment Reported w/ body: " + string(bodyBytes) + "\n")
}

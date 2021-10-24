package internal

import (
	"encoding/json"
	"fmt"
	qrcode "github.com/yeqown/go-qrcode"
	"net/http"
)

func GetOne(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	data := v.Get("data")
	qr, err := qrcode.New(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Sprintf("could not generate QRCode: %v", err))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	qr.SaveTo(w)
}

func GetRandomOne(w http.ResponseWriter, r *http.Request) {
	qr, err := qrcode.New(randString(32))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Sprintf("could not generate QRCode: %v", err))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	qr.SaveTo(w)
}

func GetCustomOne(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	if _, ok := data["info"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmt.Sprintf("could not generate QRCode without info!"))
		return
	}

	qr, err := qrcode.New(data["info"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Sprintf("could not generate QRCode: %v", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	qr.SaveTo(w)
}

package controller

import (
	"encoding/json"
	"main/api/response"
	"main/logger"
	"main/pkg/data"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		param1 := r.URL.Query().Get("key")

		if param1 == "" {

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.SetResponse(http.StatusBadRequest, ""))
			logger.Log.Println("key parameter is missing")
			return
		}

		result, _ := data.GetValue(param1)

		if result.Key == "" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response.SetResponse(http.StatusNotFound, ""))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(""))
		return
	}

}

func Set(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		var reqBody data.KeyValue
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			logger.Log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("fail"))
			return
		}

		if reqBody.Key == "" {
			logger.Log.Println("Key missing")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.SetResponse(http.StatusBadRequest, "Key missing!"))
			return
		}
		if reqBody.Value == "" {
			logger.Log.Println("Value missing")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.SetResponse(http.StatusBadRequest, "Value missing!"))
			return
		}

		data.SetValue(reqBody)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(reqBody)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response.SetResponse(http.StatusMethodNotAllowed, ""))
		return
	}
}

func Flush(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodDelete:
		data.GlobalStore = make(map[string]string)
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(response.SetResponse(http.StatusNoContent, ""))
		return
	}
}

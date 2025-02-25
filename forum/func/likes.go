package forum

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LikeRequest struct {
	PostID int    `json:"post_id"`
	Action string `json:"action"`
	Change int    `json:"change"`
}

func LikeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed !", http.StatusMethodNotAllowed)
		return
	}
	
	cookie, err := r.Cookie("session_token")
	if err != nil || cookie.Value == "" {
		sendJSONResponse(w, http.StatusUnauthorized, map[string]interface{}{
			"success": false,
		})
		return
	}
	
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request !", http.StatusBadRequest)
		return
	}

	var req LikeRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Bad Request !", http.StatusBadRequest)
		return
	}


	var column string
	switch req.Action {
	case "like":
		column = "likes"
	case "dislike":
		column = "dislikes"
	default:
		http.Error(w, "Bad Request !!", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE posts SET "+column+" = "+column+" + ? WHERE id = ?", req.Change, req.PostID)
	if err != nil {
		fmt.Println("errrrrr", err)
		http.Error(w, "Internal Server Error !!", http.StatusInternalServerError)
		return
	}

	// var count int
	// err = db.QueryRow("SELECT "+column+" FROM posts WHERE id = ?", req.PostID).Scan(&count)
	// if err != nil {
	// 	http.Error(w, "Internal Server Error !!!", http.StatusInternalServerError)
	// 	return
	// }

	sendJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		// "count":   count,
	})
	
}

func sendJSONResponse(w http.ResponseWriter, statusCode int, response map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
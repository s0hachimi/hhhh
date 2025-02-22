package forum

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LikeRequest struct {
	PostID int    `json:"post_id"`
	Action string `json:"action"`
	Change int    `json:"change"`
}

func LikeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowd", http.StatusMethodNotAllowed)
		return
	}

	var req LikeRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "bad request !", http.StatusBadRequest)
		return
	}

	// fmt.Println("0")
	var column string
	if req.Action == "like" {
		column = "likes"
	} else if req.Action == "dislike" {
		column = "dislikes"
	} else {
		http.Error(w, "bad request !!", http.StatusBadRequest)
		return
	}

	// fmt.Println("1")

	_, err = db.Exec("UPDATE posts SET "+column+" = "+column+" + ? WHERE id = ?", req.Change, req.PostID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error !", http.StatusInternalServerError)
		return
	}

	// fmt.Println("2")
	var count int
	err = db.QueryRow("SELECT "+column+" FROM posts WHERE id = ?", req.PostID).Scan(&count)
	if err != nil {
		http.Error(w, "Internal Server Error !!", http.StatusInternalServerError)
		return
	}

	// // fmt.Println("3")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"count": count , "column": column})
}

package forum

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CommentLikeRequest struct {
	CommentID int    `json:"comment_id"`
	Action    string `json:"action"`
	Change    int    `json:"change"`
}

func CommentLikeHandler(w http.ResponseWriter, r *http.Request) {
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

	ex := CheckCookie(cookie.Value)
	if !ex {
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

	var req CommentLikeRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Bad Request !", http.StatusBadRequest)
		return
	}

	fmt.Println(req)
	
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

	_, err = db.Exec("UPDATE comments SET "+column+" = "+column+" + ? WHERE id = ?", req.Change, req.CommentID)
	if err != nil {
		fmt.Println("errrrrr", err)
		http.Error(w, "Internal Server Error !!", http.StatusInternalServerError)
		return
	}

	if req.Change != -1 {
		err = LikedComment(cookie.Value, req.CommentID, req.Action)
		if err != nil {
			fmt.Println("err", err)
			http.Error(w, "Internal Server Error !!", http.StatusInternalServerError)
			return
		}
	}

	sendJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
	})
}

func LikedComment(cookieValue string, commentID int, likeType string) error {
	var userID int
	err := db.QueryRow("SELECT id FROM users WHERE session_token = ?", cookieValue).Scan(&userID)
	if err != nil {
		fmt.Println("Error fetching user ID:", err)
		return err
	}

	var n int
	if likeType == "like" {
		n = 1
	} else {
		n = -1
	}

	fmt.Println(n, likeType)

	_, err = db.Exec(`
        INSERT INTO comment_likes (user_id, comment_id, like_type) 
        VALUES (?, ?, ?) 
        ON CONFLICT(user_id, comment_id) 
        DO UPDATE SET like_type = excluded.like_type;
    `, userID, commentID, n)
	if err != nil {
		fmt.Println("Error inserting/updating post like:", err)
		return err
	}

	return nil
}

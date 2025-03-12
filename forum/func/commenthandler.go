package forum

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CommentRequest struct {
	PostID   int    `json:"post_id"`
	Username string `json:"nameOfUser"`
	Content  string `json:"comment"`
	Time     string `json:"time"`
}

func CommentHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Bad Request ! !", http.StatusBadRequest)
		return
	}

	var req CommentRequest
	if err := json.Unmarshal(body, &req); err != nil {
		fmt.Println(err)
		http.Error(w, "Bad Request !!", http.StatusBadRequest)
		return
	}
	var userID int
	err = db.QueryRow("SELECT id FROM users WHERE username = ? ", req.Username).Scan(&userID)
	if err != nil {
		http.Error(w, "Bad Request ! !", http.StatusBadRequest)
		return
	}

	res, err := db.Exec("INSERT INTO comments (user_id, post_id, comment_text, time) VALUES (?, ?, ?, ?)", userID, req.PostID, req.Content, req.Time)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad Request ! !", http.StatusBadRequest)
		return
	}

	n, _ := res.LastInsertId()

	sendJSONResponse(w, http.StatusUnauthorized, map[string]interface{}{
		"success": true,
		"id": n,
	})
}

func GetComment(r *http.Request, postID int) ([]comment, error) {

	rows, err := db.Query("SELECT id, post_id, user_id, comment_text, time, likes, dislikes FROM comments WHERE post_id = ? ORDER BY time DESC;", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commentArr []comment

	for rows.Next() {
		var c comment
		err := rows.Scan(&c.Id, &c.PostId, &c.UserId, &c.Text, &c.Time, &c.Likes, &c.Dislikes)
		if err != nil {
			return nil, err
		}

		R := GetUserReactionComments(r, c.Id)

		if R == 1 {
			c.Reaction.Like = true
		} else if R == -1 {
			c.Reaction.Dislike = true
		} else {
			c.Reaction.Like = false
			c.Reaction.Dislike = false
		}

		var username string
		err = db.QueryRow("SELECT username FROM users WHERE id = ? ", c.UserId).Scan(&username)
		if err != nil {
			return nil, err
		}
		c.Username = username

		commentArr = append(commentArr, c)
	}

	return commentArr, nil

}

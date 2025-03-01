package forum

import "net/http"



func CommentHandler(w http.ResponseWriter, r *http.Request)  {
	

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed !", http.StatusMethodNotAllowed)
		return
	}

	cookie, err := r.Cookie("session_token")
	if err != nil || cookie.Value == "" {
		// sendJSONResponse(w, http.StatusUnauthorized, map[string]interface{}{
		// 	"success": false,
		// })
		return
	}





}
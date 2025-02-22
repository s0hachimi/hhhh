document.addEventListener("DOMContentLoaded", function () {
    document.querySelectorAll(".post").forEach(post => {
        let likeBtn = post.querySelector(".like")
        let dislikeBtn = post.querySelector(".dislike")
        let likeSpan = likeBtn.querySelector("span")
        let dislikeSpan = dislikeBtn.querySelector("span")
        let postId = post.getAttribute("data-post-id") 

        likeBtn.addEventListener("click", function () {
            let likes = Number(likeSpan.textContent)
            let dislikes = Number(dislikeSpan.textContent)

            if (likeBtn.classList.contains("active")) {
                likeBtn.classList.remove("active")
                likeSpan.textContent = likes - 1
                updateLikes(postId, "like", -1) 
            } else {
                likeBtn.classList.add("active")
                likeSpan.textContent = likes + 1
                updateLikes(postId, "like", 1) 

               
                if (dislikeBtn.classList.contains("active")) {
                    dislikeBtn.classList.remove("active")
                    dislikeSpan.textContent = dislikes - 1
                    updateLikes(postId, "dislike", -1)
                }
            }
        })

        dislikeBtn.addEventListener("click", function () {
            let likes = Number(likeSpan.textContent)
            let dislikes = Number(dislikeSpan.textContent)

            if (dislikeBtn.classList.contains("active")) {
                dislikeBtn.classList.remove("active")
                dislikeSpan.textContent = dislikes - 1
                updateLikes(postId, "dislike", -1)
            } else {
                dislikeBtn.classList.add("active")
                dislikeSpan.textContent = dislikes + 1
                updateLikes(postId, "dislike", 1)

              
                if (likeBtn.classList.contains("active")) {
                    likeBtn.classList.remove("active")
                    likeSpan.textContent = likes - 1
                    updateLikes(postId, "like", -1)
                }
            }
        })
    })
})


function updateLikes(postId, action, change) {
    fetch("/like", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ post_id: Number(postId), action: action, change: change })
    })
    // .then(response => response.json())
    // .then(data => {
    //     console.log(data)
    // })
    // .catch(error => console.error(error))
}

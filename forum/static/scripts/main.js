document.addEventListener("DOMContentLoaded", function () {

    document.querySelectorAll(".post").forEach(post => {
        let likeBtn = post.querySelector(".like");
        let dislikeBtn = post.querySelector(".dislike");
        let likeSpan = likeBtn.querySelector("span");
        let dislikeSpan = dislikeBtn.querySelector("span");
        let postId = post.getAttribute("data-post-id");

        likeBtn.addEventListener("click", async function () {
            let likes = Number(likeSpan.textContent);
            let dislikes = Number(dislikeSpan.textContent);

            if (likeBtn.classList.contains("active")) {
                let success = await updateLikes(postId, "like", -1);
                if (success) {
                    likeBtn.classList.remove("active");
                    likeSpan.textContent = likes - 1;
                }
            } else {
                let success = await updateLikes(postId, "like", 1);
                if (success) {
                    likeBtn.classList.add("active");
                    likeSpan.textContent = likes + 1;

                    if (dislikeBtn.classList.contains("active")) {
                        await updateLikes(postId, "dislike", -1);
                        dislikeBtn.classList.remove("active");
                        dislikeSpan.textContent = dislikes - 1;
                    }
                }
            }
        });

        dislikeBtn.addEventListener("click", async function () {
            let likes = Number(likeSpan.textContent);
            let dislikes = Number(dislikeSpan.textContent);

            if (dislikeBtn.classList.contains("active")) {
                let success = await updateLikes(postId, "dislike", -1);
                if (success) {
                    dislikeBtn.classList.remove("active");
                    dislikeSpan.textContent = dislikes - 1;
                }
            } else {
                let success = await updateLikes(postId, "dislike", 1);
                if (success) {
                    dislikeBtn.classList.add("active");
                    dislikeSpan.textContent = dislikes + 1;

                    if (likeBtn.classList.contains("active")) {
                        await updateLikes(postId, "like", -1);
                        likeBtn.classList.remove("active");
                        likeSpan.textContent = likes - 1;
                    }
                }
            }
        });
    })

    document.querySelectorAll(".main-comment").forEach(but => {
        let button = but.querySelector("#showComment")
       let n = 0
       button.addEventListener("click", function () {
        if (n === 0) {
            but.querySelector(".comment-text").style.display = "flex"
            button.innerHTML = '<i class="fa-solid fa-comment-slash"></i>'
            n = 1
        } else {
            but.querySelector(".comment-text").style.display = "none"
            button.innerHTML = '<i class="fa-solid fa-comment"></i>'
            n = 0
        }
       })

       let commentText = but.querySelector("#com")
       let text = but.querySelector("input")

       commentText.addEventListener("click", function () {
        
       })
       
    })




});

async function updateLikes(postId, action, change) { 
    try {
        let response = await fetch("/like", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ post_id: Number(postId), action: action, change: change })
        });

        let data = await response.json();

        if (!data.success) {
            window.location.href = "/login-page"
            return false
        }
        
        return true

    } catch (error) {
        console.error("Fetch error:", error);
        return false
    }
}


function openNav() {
    document.getElementById("mySidenav").style.width = "250px";
}
function closeNav() {
    document.getElementById("mySidenav").style.width = "0";
}

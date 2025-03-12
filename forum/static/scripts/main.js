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
                if (likes !== 0) {
                    let success = await updateLikes(postId, "like", -1);
                    if (success) {
                        likeBtn.classList.remove("active");
                        likeSpan.textContent = likes - 1;
                    }
                }

            } else {
                let success = await updateLikes(postId, "like", 1);
                if (success) {
                    likeBtn.classList.add("active");
                    likeSpan.textContent = likes + 1;

                    if (dislikeBtn.classList.contains("active")) {
                        if (dislikes !== 0) {
                            await updateLikes(postId, "dislike", -1);
                            dislikeBtn.classList.remove("active");
                            dislikeSpan.textContent = dislikes - 1;
                        }

                    }
                }
            }
        });

        dislikeBtn.addEventListener("click", async function () {
            let likes = Number(likeSpan.textContent);
            let dislikes = Number(dislikeSpan.textContent);

            if (dislikeBtn.classList.contains("active")) {
                if (dislikes !== 0) {
                    let success = await updateLikes(postId, "dislike", -1);
                    if (success) {
                        dislikeBtn.classList.remove("active");
                        dislikeSpan.textContent = dislikes - 1;
                    }
                }

            } else {
                let success = await updateLikes(postId, "dislike", 1);
                if (success) {
                    dislikeBtn.classList.add("active");
                    dislikeSpan.textContent = dislikes + 1;

                    if (likeBtn.classList.contains("active")) {
                        if (likes !== 0) {
                            await updateLikes(postId, "like", -1);
                            likeBtn.classList.remove("active");
                            likeSpan.textContent = likes - 1;
                        }

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
        let username = document.getElementById("username")
        let u = 0
        if (username === null) {
            u = 1
        } else {
            username = username.textContent
        }
        let postID = but.getAttribute("data-post-id")


        commentText.addEventListener("click", function () {
            if (text.value === "") {
                return
            }
            document.querySelectorAll(".comment-content").forEach(async function (el) {

                let elID = el.getAttribute("data-post-id")

                let nameOfUser = document.createElement("h3")
                let content = document.createElement("pre")
                let time = document.createElement("span")
                let d = document.createElement("div")
                let div = document.createElement("div")
                if (u === 1) {
                    location.href = "/login-page"
                    return
                }
                let date = new Date()
                let arrDate = date.toISOString().split("T")
                let dateNow = arrDate[0] + " " + arrDate[1].slice(0, 8)


                nameOfUser.textContent = username
                content.textContent = text.value
                time.textContent = dateNow


                if (elID === postID) {
                    text.value = ""
                    n = await addComment(postID, nameOfUser.textContent, content.textContent, dateNow)
                    div.setAttribute("id", n)
                    d.innerHTML = `
                <div class="box comment-like" data-comment-id="${n}">
                    <button class="like">
                        <i class="fa-solid fa-thumbs-up"></i>
                        <span>0</span>
                    </button>

                    <button class="dislike">
                        <i class="fa-solid fa-thumbs-down"></i>
                        <span>0</span>
                    </button>
                `
                d.className = "reactions"
                div.className = "div"
                
                
                div.append(nameOfUser, content, time, d)

                    el.querySelector(".hihi").prepend(div)
                    bindCommentLikeDislikeEvents(div.querySelector(".comment-like")); // Bind events for the new comment
                }

            })

        })

    })

    document.querySelectorAll(".comment-content").forEach(el => {
        let c = 0
        el.querySelector(".hide").addEventListener("click", function () {
            if (c === 0) {
                el.querySelector(".hihi").style.display = "block"
                this.textContent = "Hide Comments"
                c = 1
            } else {
                el.querySelector(".hihi").style.display = "none"
                this.textContent = "Show Comments"
                c = 0
            }
        })
    })

    
    Array.from(document.getElementsByClassName("comment-like")).forEach(el => {
        bindCommentLikeDislikeEvents(el);
    })


    function bindCommentLikeDislikeEvents(el) {
        let likeBtn = el.querySelector(".like");
        let dislikeBtn = el.querySelector(".dislike");
        let likeSpan = likeBtn.querySelector("span");
        let dislikeSpan = dislikeBtn.querySelector("span");
        let commentId = el.getAttribute("data-comment-id");

        likeBtn.addEventListener("click", async function () {
            let likes = Number(likeSpan.textContent);
            let dislikes = Number(dislikeSpan.textContent);

            if (likeBtn.classList.contains("active")) {
                if (likes !== 0) {
                    let success = await updateCommentLikes(commentId, "like", -1);
                    if (success) {
                        likeBtn.classList.remove("active");
                        likeSpan.textContent = likes - 1;
                    }
                }

            } else {
                let success = await updateCommentLikes(commentId, "like", 1);
                if (success) {
                    likeBtn.classList.add("active");
                    likeSpan.textContent = likes + 1;

                    if (dislikeBtn.classList.contains("active")) {
                        if (dislikes !== 0) {
                            await updateCommentLikes(commentId, "dislike", -1);
                            dislikeBtn.classList.remove("active");
                            dislikeSpan.textContent = dislikes - 1;
                        }

                    }
                }
            }
        });

        dislikeBtn.addEventListener("click", async function () {
            let likes = Number(likeSpan.textContent);
            let dislikes = Number(dislikeSpan.textContent);

            if (dislikeBtn.classList.contains("active")) {
                if (dislikes !== 0) {
                    let success = await updateCommentLikes(commentId, "dislike", -1);
                    if (success) {
                        dislikeBtn.classList.remove("active");
                        dislikeSpan.textContent = dislikes - 1;
                    }
                }

            } else {
                let success = await updateCommentLikes(commentId, "dislike", 1);
                if (success) {
                    dislikeBtn.classList.add("active");
                    dislikeSpan.textContent = dislikes + 1;

                    if (likeBtn.classList.contains("active")) {
                        if (likes !== 0) {
                            await updateCommentLikes(commentId, "like", -1);
                            likeBtn.classList.remove("active");
                            likeSpan.textContent = likes - 1;
                        }

                    }
                }
            }
        });
    }

});

async function addComment(postID, username, content, time) {
    try {
        let response = await fetch("/comment", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ post_id: Number(postID), nameOfUser: username, comment: content, time: time })
        })

        let data = await response.json();
        if (!data.success) {
            window.location.href = "/login-page"
            return false
        }

        return true, data.id
    } catch (error) {
        console.error(error);
        return false
    }

}

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

async function updateCommentLikes(commentId, action, change) {
    try {
        let response = await fetch("/comment-like", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ comment_id: Number(commentId), action: action, change: change })
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
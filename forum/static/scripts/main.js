let buttonComment = document.querySelector(".button-comment")

buttonComment.addEventListener("click", function(){

    let title = document.querySelector("#title")
    let comment = document.querySelector("#comment")

    let divTitle = document.createElement("div")
    let divComment = document.createElement("div")

    divTitle.textContent = title.value
    divComment.textContent = comment.value

    let main = document.querySelector("main")

    main.append(divTitle, divComment)
})
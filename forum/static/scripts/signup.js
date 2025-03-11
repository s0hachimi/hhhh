document.addEventListener("DOMContentLoaded", function () {


    let signupForm = document.querySelector("#signupForm")
    let errorMsg = document.querySelector("#errorMsg")
    

    signupForm.addEventListener("submit", async function (e) {
        e.preventDefault()

        const formData = new FormData(signupForm);
        const data = Object.fromEntries(formData.entries());

        let res = await send(data)

        if (res.message === "users.email") {
            res.message = 'Email Already Registered'
        } else if (res.message === "users.username"){
            res.message = 'Username Already Taken'
        }

        
        if (!res.success) {
            errorMsg.textContent = res.message;
            errorMsg.style.display = "block";
            e.preventDefault()
            return
        }

        window.location.href = "/"

    })

})


async function send(data) {
    try {
        const response = await fetch("/signup", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(data)
        })

        const result = await response.json();

        return result

    } catch (error) {
        console.error(error);
        return false
    }

}


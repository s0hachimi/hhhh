document.addEventListener("DOMContentLoaded", function () {
    const loginForm = document.getElementById("loginForm");
    const errorMsg = document.getElementById("errorMsg");

    const pass = document.querySelector("#password")
    document.querySelector("#i").addEventListener("click", () => {
        if (pass.className === "hide") {
            pass.setAttribute("type", "text")
            pass.className = "show"
        } else {
            pass.setAttribute("type", "password")
            pass.className = "hide"
        }
    })


    loginForm.addEventListener("submit", async function (event) {
        event.preventDefault();



        const formData = new FormData(loginForm);
        const data = Object.fromEntries(formData.entries());


        try {
            const response = await fetch("/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(data)
            });

            const result = await response.json();
            console.log(result);


            if (result.error) {
                errorMsg.textContent = result.error;
                errorMsg.style.display = "block";
            } else {
                window.location.href = "/";
            }
        } catch (error) {
            console.error("Fetch error:", error);
        }
    });
});
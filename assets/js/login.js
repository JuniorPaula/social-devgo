const loginForm = document.querySelector("#login-form")
loginForm.addEventListener("submit", loginHandle)

function loginHandle(e) {
    e.preventDefault()
    const email = document.querySelector("#email").value
    const password = document.querySelector("#password").value

    const data = { email, password }

    fetch('/login', {
        method: 'POST',
        body: JSON.stringify(data),
        headers: {"Content-type": "application/json"}
    }).then((data) => {
        console.log('data', data)
        if (data.status !== 200) {
            Swal.fire("Ops!!", "Usuário ou senha inválidos", "error")
                .then(function() {
                    email = ""
                    password = ""
                })
            return
        }

        window.location.href = "/home"
    }).catch(e => {
        Swal.fire("Ops!!", "Usuário ou senha inválidos", "error")
            .then(function() {
                email = ""
                password = ""
            })        
    })
}
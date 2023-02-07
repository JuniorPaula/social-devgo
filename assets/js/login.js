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
        console.log(data)
        window.location.href = "/home"
    }).catch(e => {
        console.log(e)
        alert('email ou senha inválidos')
    })
}
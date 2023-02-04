const formSignup = document.querySelector("#signup-form")
formSignup.addEventListener("submit", createUser)

async function createUser(e) {
    e.preventDefault()
    
    const name = document.querySelector("#name").value
    const email = document.querySelector("#email").value
    const nickname = document.querySelector("#nickname").value
    const password = document.querySelector("#password").value
    const confirmPassword = document.querySelector("#confirmPassword").value

    if (password !== confirmPassword) {
        alert('As senhas não conferem!')
        return
    }

    const data = {
        name,
        email,
        nickname,
        password
    }

    fetch('/users', {
        method: 'POST',
        body: JSON.stringify(data),
        headers: {"Content-type": "application/json"}
    })
}
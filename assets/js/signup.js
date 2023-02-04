const formSignup = document.querySelector("#signup-form")

formSignup.addEventListener("submit", createUser)

function createUser(e) {
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

    fetch("/users", {
        method: "POST",
        body: {
            name,
            email,
            nickname,
            password
        }
    }).then(response => {
        console.log(response)
        
    }).catch(e => console.log(e))
}
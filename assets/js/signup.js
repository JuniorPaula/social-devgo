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
        Swal.fire("Ops!!", "As senhas não conferem!", "error")
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
    .then(data => {
        if (data.status !== 201) {
            Swal.fire("Ops!!", "Erro ao criar a conta!", "error")
            alert("Erro ao atualizar publicação :/")
            return
        }

        Swal.fire('Sucesso', 'Conta criada com sucesso!', 'success')
    })
    .catch(e => {
        Swal.fire("Ops!!", "Algo deu errado!", "error")
    })
}
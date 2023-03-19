const formUpdatePassword = document.querySelector("#update-password")
const oldPassword = document.querySelector("#old-password")
const newPassword = document.querySelector("#new-password")
const confirmPassword = document.querySelector("#confirm-password")

formUpdatePassword.addEventListener("submit", updatePassword)

function updatePassword(e) {
    e.preventDefault()

    if (newPassword.value !== confirmPassword.value) {
        Swal.fire("Ops...", "As senhas não conferem :(", "warning")
        return
    }

    const data = {
        old_password: oldPassword.value,
        new_password: newPassword.value
    }

    fetch(`/update-password`, {
        method: 'POST',
        body: JSON.stringify(data),
        headers: {"Content-type": "application/json"}
    })
    .then(data => {
        if (data.status !== 204) {
            Swal.fire("Ops!!", "Error ao atualizar a senha!", "error")
            return
        }

        Swal.fire('Sucesso', 'Senha atualizada com sucesso!', 'success')
            .then(function() {
                window.location = "/profile"
            })
    })
    .catch(e => {
        Swal.fire("Ops!!", "Error ao atualizar senha!", "error")
    })
}
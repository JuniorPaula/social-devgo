const formEditUser = document.querySelector("#edit-user")

formEditUser.addEventListener("submit", editUser)

function editUser(e) {
    e.preventDefault()
    
    const data = {
        name: document.querySelector("#name").value,
        email: document.querySelector("#email").value,
        nickname: document.querySelector("#nickname").value
    }

    fetch(`/edit-user`, {
        method: 'PUT',
        body: JSON.stringify(data),
        headers: {"Content-type": "application/json"}
    })
    .then(data => {
        if (data.status !== 204) {
            Swal.fire("Ops!!", "Error ao editar usuário!", "error")
            return
        }

        Swal.fire('Sucesso', 'Usuário editado com sucesso!', 'success')
            .then(function() {
                window.location = "/profile"
            })
    })
    .catch(e => {
        Swal.fire("Ops!!", "Error ao editar usuário!", "error")
    })
}
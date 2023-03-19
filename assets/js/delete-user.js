const btnDeleteUse = document.querySelector("#delete-user")

btnDeleteUse.addEventListener("click", deleteUser)

function deleteUser(e) {
    Swal.fire({
        title: "Atenção",
        text: "Tem certeza que deseja excluir sua conta? Essa ação é inrreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirm) {
        if (confirm.value) {
            fetch(`/delete-user`, {
                method: 'DELETE',
            })
            .then(data => {
                if (data.status !== 204) {
                    Swal.fire("Ops!!", "Error ao deletar sua conta!", "error")
                    return
                }
        
                Swal.fire('Sucesso', 'Conta deletada com sucesso!', 'success')
                    .then(function() {
                        window.location = "/logout"
                    })
            })
            .catch(e => {
                Swal.fire("Ops!!", "Error ao deletar sua conta!", "error")
            })
        }
    })
}
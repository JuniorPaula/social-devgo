const formPost = document.querySelector("#new-post")
const likePostBtn = document.querySelectorAll(".like-post")
const btnDelete = document.querySelectorAll(".delete-post")

formPost.addEventListener("submit", createPost)

function createPost(e) {
    e.preventDefault()
    const title = document.querySelector("#title").value
    const content = document.querySelector("#content").value

    const data = {
        title,
        content
    }

    fetch('/posts', {
        method: 'POST',
        body: JSON.stringify(data),
        headers: {"Content-type": "application/json"}
    })
    .then(data => {
        if (data.status !== 201) {
            Swal.fire("Ops!!", "Erro ao criar publicação :/", "error")    
            return
        }

        window.location.href = "/home"
    })
    .catch(e => {
        Swal.fire("Ops!!", "Erro ao criar publicação :/", "error")
    })

}

likePostBtn.forEach(btn => {
    btn.addEventListener("click", likePost)
})

btnDelete.forEach(btn => {
    btn.addEventListener("click", deletePost)
})

function likePost(e, btn) {
    e.preventDefault()
    const postId = e.target.parentElement.parentElement.dataset.postId

    fetch(`/posts/${postId}/like`, {
        method: 'POST',
    })
    .then(data => {
        if (data.status !== 204) {
            Swal.fire("Ops!!", "Erro ao curtir publicação :/", "error")    
            return
        }
        const counterLikes = e.srcElement.nextElementSibling
        const qtdLikes = parseInt(counterLikes.innerText)
        counterLikes.innerText = qtdLikes + 1
    })
    .catch(e => {
        Swal.fire("Ops!!", "Erro ao curtir publicação :/", "error")
    })
}

function deletePost(e) {
    e.preventDefault()
    const postId = e.target.parentElement.parentElement.dataset.postId

    Swal.fire({
        title: "Atenção!!",
        text: "Deseja realmente excluir essa publicação ?",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    })
    .then(function(isConfirm) {
        if (!isConfirm.value) return

        fetch(`/posts/${postId}`, {
            method: 'DELETE',
        })
        .then(data => {
            if (data.status !== 204) {
                Swal.fire("Ops!!", "Erro ao deletar publicação :/", "error")
                return
            }
            Swal.fire("Sucesso!", "Publicação deletar com sucesso", "success")
                .then(function() {
                    window.location.href = "/home"
                })
        })
        .catch(e => {
            Swal.fire("Ops!!", "Erro ao deletar publicação :/", "error")
        })
    })
    
}
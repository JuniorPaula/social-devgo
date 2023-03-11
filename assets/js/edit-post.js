const btnEditPost = document.querySelector("#btn-edit-post")

btnEditPost.addEventListener("click", editPost)

function editPost(e) {
    e.preventDefault()
    e.target.disabled = true

    const postId = e.target.dataset.postId
    
    const data = {
        title: document.querySelector("#title").value,
        content: document.querySelector("#content").value
    }

    fetch(`/posts/${postId}`, {
        method: 'PUT',
        body: JSON.stringify(data),
        headers: {"Content-type": "application/json"}
    })
    .then(data => {
        if (data.status !== 204) {
            alert("Erro ao atualizar publicação :/")
            return
        }

        alert("Publicação atualizada com sucesso!")
    })
    .catch(e => {
        alert("Erro ao atualizar publicação :/")
        console.log("[ERROR::] ", e)
    })
    .finally(() => {
        e.target.disabled = false
    })
}
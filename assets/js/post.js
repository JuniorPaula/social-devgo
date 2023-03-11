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
            alert("Erro ao criar publicação :/")    
            return
        }

        window.location.href = "/home"
    })
    .catch(e => {
        alert("Erro ao criar publicação :/")
        console.log("[ERROR::] ", e)
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
            alert("Erro ao curtir publicação :/")    
            return
        }
        const counterLikes = e.srcElement.nextElementSibling
        const qtdLikes = parseInt(counterLikes.innerText)
        counterLikes.innerText = qtdLikes + 1
    })
    .catch(e => {
        alert("Erro ao curtir publicação :/")
        console.log("[ERROR::] ", e)
    })
}

function deletePost(e) {
    e.preventDefault()
    const postId = e.target.parentElement.parentElement.dataset.postId
    
    fetch(`/posts/${postId}`, {
        method: 'DELETE',
    })
    .then(data => {
        if (data.status !== 204) {
            alert("Erro ao deletar publicação :/")
            return
        }
        alert("Publicação deletada com sucesso!")
        window.location.href = "/home"
    })
    .catch(e => {
        alert("Erro ao deletar publicação :/")
        console.log("[ERROR::] ", e)
    })
}
const formPost = document.querySelector("#new-post")
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
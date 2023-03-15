const btnUnfollower = document.querySelector("#unfollower")
const btnFollower = document.querySelector("#follower")

btnFollower.addEventListener("click", followerUser)
btnUnfollower.addEventListener("click", unfollowerUser)

function unfollowerUser(e) {
    console.log(e)
    e.preventDefault()
    e.target.disabled = true

    const userId = e.target.dataset.userId

    fetch(`/users/${userId}/unfollower`, {
        method: 'POST',
    })
    .then(data => {
        if (data.status !== 204) {
            Swal.fire("Ops!!", "Ocorreu um erro inesperado!", "error")
            return
        }
        window.location = `/users/${userId}`
        
    })
    .catch(e => {
        Swal.fire("Ops!!", "Error ao editar a publicação!", "error")
        e.target.disabled = false
    })
}

function followerUser(e) {
    console.log(e)
    e.preventDefault()
    e.target.disabled = true

    const userId = e.target.dataset.userId

    fetch(`/users/${userId}/follower`, {
        method: 'POST',
    })
    .then(data => {
        if (data.status !== 204) {
            Swal.fire("Ops!!", "Ocorreu um erro inesperado!", "error")
            return
        }
        window.location = `/users/${userId}`
        
    })
    .catch(e => {
        Swal.fire("Ops!!", "Error ao editar a publicação!", "error")
        e.target.disabled = false
    })
}
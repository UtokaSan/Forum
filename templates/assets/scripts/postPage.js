var url = new URL(window.location.href);
var param1Value = url.searchParams.get('id');
var parameterUrl = {
    id: parseInt(param1Value)
};
fetch("/api/takepostid", {
    method: "POST",
    headers: {
        "Content-Type": "application/json"
    },
    body: JSON.stringify(parameterUrl),
})
    .then(response => response.json())
    .then(data => {
        console.log(data);
        var title = document.querySelector('.Title');
        title.innerText = data.info[0].title;
        const para = document.createElement('p');
        var description= document.querySelector('.contentPost');
        para.innerText = data.info[0].text;
        description.appendChild(para)
        var like= document.querySelector('.text-like');
        like.innerText = data.info[0].like;
        var dislike= document.querySelector('.text-dislike');
        dislike.innerText = data.info[0].dislike;
        const image = document.createElement('img');
        image.src = data.info[0].image;
        var imageSelect= document.querySelector('.contentPost');
        imageSelect.appendChild(image)
    })
    .catch(error => {
        console.error("Error sending data id", error);
    });
const likeButton = document.querySelector('.like');
const dislikeButton = document.querySelector('.dislike');
likeButton.addEventListener('click', envoyerLike);
dislikeButton.addEventListener('click', envoyerDislike);
function envoyerLike() {
    sendData('like');
}
function envoyerDislike() {
    sendData('dislike');
}
function sendData (formDataSend) {

    const data = {
        reactions: formDataSend,
        post_id: param1Value,
    };
    fetch("/api/likeordislike", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data),
    })
        .then(response => {
            if (!response.ok) {
                if (response.status === 400) {
                    alert("already liked");
                } else {
                    console.log("problem");
                }
            } else {
                location.reload();
                alert("reaction added");
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

const editPostSubmit = document.querySelector('.editForm');

editPostSubmit.addEventListener('submit', function(event) {
    event.preventDefault()
    fetch("/api/editPost", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            id:parseInt(url.searchParams.get('id')),
            title: document.querySelector('#title').value,
            texte: document.querySelector('#message').value
        }),
    })
        .then(response => {
            if (!response.ok) {
                if (response.status === 400) {
                    alert("tu n'as pas le droit de faire ça");
                } else {
                    console.log("problem");
                }
            } else {
                // location.reload();
                alert("test");
                location.reload();
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });

})
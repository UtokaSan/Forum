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
        console.log(data.info[0].title);
        var title = document.querySelector('.Title');
        title.innerText = data.info[0].title;

        var description= document.querySelector('.contentPost');
        description.innerText = data.info[0].text;

        var like= document.querySelector('.text-like');
        like.innerText = data.info[0].like;

        var dislike= document.querySelector('.text-dislike');
        dislike.innerText = data.info[0].dislike;
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
        post_id: parseInt(param1Value),
    };
    fetch("/api/likeordislike", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: data,
    })
        .then(response => response.text())
        .then(data => {
            console.log(data);
        })
        .catch(error => {
            console.error('Error:', error);
        });
}
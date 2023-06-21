
fetch("/api/getComments", {
    method: "POST",
    headers: {
        "Content-Type": "application/json"
    },
    body:JSON.stringify({
        id:parseInt(url.searchParams.get('id'))
    })
})
    .then(response => response.json())
    .then(data => {
        console.log(data)
        data.comment.forEach(post => {
            var newDiv = document.createElement("div");
            newDiv.id = "comment";
            var anchorElement = document.createElement("p");
            anchorElement.textContent = post.IDCreator + " : " + post.text;
            newDiv.appendChild(anchorElement);
            console.log(newDiv);
            console.log(document.getElementById('comments'))
            document.getElementById("comments").appendChild(newDiv);

        });
    })
    .catch(error => {
        console.error("Error update", error);
    });

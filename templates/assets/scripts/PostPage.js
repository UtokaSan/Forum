

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
        var texttitle = data.info[0].title;
        title.innerText = texttitle;

        var description= document.querySelector('.contentPost');
        var textdescription = data.info[0].text;
        description.innerText = textdescription;

    })
    .catch(error => {
        console.error("Error sending data id", error);
    });

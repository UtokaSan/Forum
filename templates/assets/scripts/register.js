const form = document.querySelector('form');

form.addEventListener('submit', function(event) {
    event.preventDefault();
    const url = 'api/register';
    const dataUser = {
        pseudo: document.getElementById("pseudo").value ,
        email: document.getElementById("mailAdress").value,
        password: document.getElementById("password").value,
    };
    console.log("dataUser")
    console.log(dataUser)
    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(dataUser)
    })
        .then(response => response.json())
        .then(response => {
            console.log("response")
            if (!response.ok) {
                    alert(response.message)
            } else {
                console.log("user logged")
            }
        })
        .catch(error => {
            console.error('Erreur lors de la requête:', error);
        });
});
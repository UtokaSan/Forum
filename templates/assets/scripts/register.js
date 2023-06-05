const form = document.querySelector('form');

form.addEventListener('submit', function(event) {
    event.preventDefault();
    const url = 'api/register';
    const formData = new FormData(form);
    const dataUser = {
        pseudo: formData.get('pseudo'),
        email: formData.get('mailAdress'),
        password: formData.get('password'),
    };
    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(dataUser)
    })
        .then(response => {
            console.log(response.body)
            if (!response.ok) {
                    alert(response.body)
            } else {
                console.log("user logged")
            }
        })
        .catch(error => {
            console.error('Erreur lors de la requête:', error);
        });
});
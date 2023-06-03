const form = document.querySelector('form');

form.addEventListener('submit', function(event) {
    event.preventDefault();
    const url = '/api/login';
    const formData = new FormData(form);
    const payload = {
        email: formData.get('email'),
        password: formData.get('password')
    };
    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
    })
        .then(response => {
            if (response.ok) {
                console.log('Formulaire soumis avec succès');
                // const tokenCookie = document.cookie;
                // console.log(tokenCookie)
            } else {
                console.error('Erreur lors de la soumission du formulaire');
            }
        })
        .catch(error => {
            console.error('Erreur lors de la requête:', error);
        });
});

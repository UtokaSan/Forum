const form = document.querySelector('form');

form.addEventListener('submit', function(event) {
    event.preventDefault();
    const url = '/api/login';
    const formData = new FormData(form);
    const payload = {
        email: formData.get('email'),
        password: formData.get('password'),
        saveinfo: formData.get('saveinfo'),
    };
    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        credentials: 'include',
        body: JSON.stringify(payload)
    })
        .then(response => {
            if (!response.ok) {
                if (response.status === 403) {
                    console.log('user ban');
                    alert("Votre compte est bannis")
                }
                if (response.status === 401) {
                    alert("Votre compte n'existe pas")
                }
            } else {
                console.log("user logged")
            }
        })
        .catch(error => {
            console.error('Erreur lors de la requête:', error);
        });
});
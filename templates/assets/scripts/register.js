const form = document.querySelector('form');
console.log("load2")

// console.log(pseudo,email,password,confPassword)

form.addEventListener('submit', function(event) {
    event.preventDefault();

    let pseudo = document.getElementById("pseudo").value
    let email = document.getElementById("mailAdress").value
    let password = document.getElementById("password").value
    let confPassword = document.getElementById("confPassword").value

    if (CheckEmail(email)) return;
    if (CheckPassword(password,confPassword)) return;

    const url = 'api/register';
    const dataUser = {
        pseudo: pseudo ,
        email: email,
        password: password,
    };

    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(dataUser)
    })
        .then(response => response.json())
        .then(response => {
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

function CheckPassword(password,confPassword) {
    if (password !== confPassword) {
        alert("Le mot de passe et la confirmation du mot de passe ne sont pas les même")
        return true
    }
    return false
}

function CheckEmail(email) {
    console.log("mince1")
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    console.log(emailRegex.test(email))
    if (!emailRegex.test(email)) {
        console.log("bug with email")
        alert("L'adresse mail est érronée ")
    }

    return !emailRegex.test(email);
}


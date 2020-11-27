async function Save(user) {
    console.log(user);
    fetch(
        '/users/create',
        {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json;charset=utf-8'
            },
            body: JSON.stringify(user)
        }
    ).then((response) => {
        if (response.ok) {
            let data = response.json(); 
            console.log(data);
        }
    })
}

export default { Save }
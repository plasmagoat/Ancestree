const BASE_URL = 'https://ancestree-api.procopius.dk/api/'

class AuthService {
  login(user) {
    return fetch(`${BASE_URL}person/${user.id}`, {
      method: 'GET',
      //mode: 'cors',
      //   headers: {
      //     'Content-Type': 'application/json',
      //   },
      //body: JSON.stringify(user),
    })
      .then(response => response.json())
      .then(data => {
        localStorage.setItem('user', JSON.stringify(data.person))
        //cb(data.body)
        return data.person
      })
  }

  logout() {
    localStorage.removeItem('user')
  }

  register(user) {
    return fetch(`${BASE_URL}login`, {
      method: 'POST',
      mode: 'cors',
      //   headers: {
      //     'Content-Type': 'application/json',
      //   },
      body: JSON.stringify(user),
    }).then(response => response.json())
  }
}

export default new AuthService()

const BASE_URL = 'http://localhost:9000/api/person'

export default {
  createPerson(person, cb) {
    fetch(`${BASE_URL}`, {
      method: 'POST',
      mode: 'cors',
      //   headers: {
      //     'Content-Type': 'application/json',
      //   },
      body: JSON.stringify(person),
    })
      .then(response => response.json())
      .then(data => cb(data.body))
  },
  updatePerson(id, person, cb) {
    fetch(`${BASE_URL}/${id}`, {
      method: 'PUT',
      mode: 'cors',
      //   headers: {
      //     'Content-Type': 'application/json',
      //   },
      body: JSON.stringify(person),
    })
      .then(response => response.json())
      .then(data => cb(data.body))
  },
  createChild(id, person, cb) {
    fetch(`${BASE_URL}/child/${id}`, {
      method: 'POST',
      mode: 'cors',
      //   headers: {
      //     'Content-Type': 'application/json',
      //   },
      body: JSON.stringify(person),
    })
      .then(response => response.json())
      .then(data => cb(data.body))
  },
  createParent(id, person, cb) {
    fetch(`${BASE_URL}/parent/${id}`, {
      method: 'POST',
      mode: 'cors',
      body: JSON.stringify(person),
    })
      .then(response => response.json())
      .then(data => cb(data.body))
  },
  createLink(childId, parentId, cb) {
    fetch(`${BASE_URL}/link/${childId}/${parentId}`, {
      method: 'POST',
      mode: 'cors',
      //   headers: {
      //     'Content-Type': 'application/json',
      //   },
    })
      .then(response => response.json())
      .then(data => cb(data.body))
  },
}

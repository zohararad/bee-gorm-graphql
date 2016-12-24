import http from './http'

let queries = {}

queries.loadUsers = `query LoadUsers {
  users {
    id
    firstName
    lastName
    email
  }
}`

queries.loadUser = `query User($id: ID!) {
  user(id: $id) {
    id
    firstName
    lastName
    email
  }
}`

queries.createUser = `mutation CreateUser($firstName: String!, $lastName: String!, $email: String!, $password: String!) {
  createUser(firstName: $firstName, lastName: $lastName, email: $email, password: $password) {
    id
    firstName
    lastName
    email
  }
}`

queries.updateUser = `mutation UpdateUser($id: ID!, $firstName: String!, $lastName: String!, $email: String!, $password: String!) {
  updateUser(id: $id, firstName: $firstName, lastName: $lastName, email: $email, password: $password) {
    id
    firstName
    lastName
    email
  }
}`

queries.deleteUser = `mutation DeleteUser($id: ID!) {
  deleteUser(id: $id) {
    id
  }
}`

export default {
  loadUsers (cb) {
    http.client.query(queries.loadUsers).then(result => {
      cb(result.users)
    })
  },
  loadUser (args, cb) {
    http.client.query(queries.loadUser, args).then(result => {
      cb(result.user)
    })
  },
  createUser (args, cb) {
    http.client.query(queries.createUser, args).then(result => {
      cb(result.createUser)
    })
  },
  updateUser (args, cb) {
    http.client.query(queries.updateUser, args).then(result => {
      cb(result.updateUser)
    })
  },
  deleteUser (args, cb) {
    http.client.query(queries.deleteUser, args).then(result => {
      cb(result.deleteUser)
    })
  }
}

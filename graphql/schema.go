package graphql

var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}
	# The query type, represents all of the entry points into our object graph
	type Query {
	  users(): [User]!
	  user(id: ID!): User
	}
	type Mutation {
	  createUser(firstName: String!, lastName: String!, email: String!, password: String!): User
	  updateUser(id: ID!, firstName: String, lastName: String, email: String, password: String): User
	  deleteUser(id: ID!): ID
	}
	type User {
	  id: ID!
	  firstName: String!
	  lastName: String!
	  email: String!
	  password: String!
	}
`


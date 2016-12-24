# Beego, Gorm, GraphQL & Vue.js demo

This app is a demo of using a Beego powered backend, with a Vue.js powered frontend.

## Server-Side

* **Beego** - Golang powered Web framework [Website](https://beego.me/)
* **gorm** - Golang ORM - used instead of Beego's ORM due to better support of Postgres features (and fun) [Website](http://jinzhu.me/gorm)
* **graphql-go** - Golang implementation of GraphQL. [Repo](https://github.com/neelance/graphql-go)

## Client-Side

* **Vue.js** - used to build the UI components [Website](http://vuejs.org/)
* **Vuex** - state management for Vue.js [Website](http://vuex.vuejs.org/)
* **Vue Router** - router for Vue.js, using for routing support [Website](http://router.vuejs.org/)
* **Bulma** - CSS framework used to for UI components [Website](http://bulma.io/)

## Setup

* Install Node.js & Golang (project was developed on Golang v1.7)
* Install glide for Golang dependency management [Website](https://glide.sh/)
* Install yarn for client-side dependency management [Website](https://yarnpkg.com/)
* Run the following

```bash
$ git clone git@github.com:zohararad/bee-gorm-graphql.git
$ cd bee-gorm-graphql
$ yarn install
$ glide up
$ go get github.com/beego/bee
$ npm run build
$ bee run #open http://localhost:8080
```

## Next Steps

* You need to setup your Postgres DB (take a look at conf/app.conf)
* You need to create at least one user - Easiest way is to do so via the GraphiQL UI - http://localhost:/graphql. Take a look at graphql/schema.go for the query structure
* Log into the system http://localhost:8080/session/login
* Enjoy

## Credits

GraphiQL UI is embedded as is. This was blatantly "ripped off" from the **graphql-go** Github repo [StarWars Example](https://github.com/neelance/graphql-go/blob/master/example/starwars/server/server.go)

## Todo

* Tests are missing completely - Will add them as time goes by

Pull Requests are more than welcome!
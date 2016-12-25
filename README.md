# Beego, Gorm, GraphQL & Vue.js demo

This app is a demo of using a Beego powered backend, with a Vue.js powered frontend.

## Server-Side

* [Beego](https://beego.me/) - Golang powered Web framework.
* [Gorm](http://jinzhu.me/gorm) - Golang ORM - used instead of Beego's ORM due to better support of Postgres features (and fun).
* [graphql-go](https://github.com/neelance/graphql-go) - Golang implementation of GraphQL.

## Client-Side

* [Vue.js](http://vuejs.org/) - used to build the UI components.
* [Vuex](http://vuex.vuejs.org/) - state management for Vue.js.
* [Vue Router](http://router.vuejs.org/) - router for Vue.js, using for routing support.
* [Bulma](http://bulma.io/) - CSS framework used to for UI components.

## Setup

* Install Node.js & Golang (project was developed on Golang v1.7)
* Install [glide](https://glide.sh/) for Golang dependency management.
* Install [yarn](https://yarnpkg.com/) for client-side dependency management.
* Run the following:

```bash
$ go get github.com/zohararad/bee-gorm-graphql
$ cd $GOPATH/src/github.com/zohararad/bee-gorm-graphql
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

## Contributions

Pull Requests are more than welcome!

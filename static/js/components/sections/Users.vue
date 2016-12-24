<template>
  <div>
    <nav class="level">
      <div class="level-left">
        <div class="level-item">
          <h3 class="subtitle is-3">Users</h3>
        </div>
      </div>
      <div class="level-right">
        <div class="level-item">
          <a class="button is-small is-info" @click.prevent="createUser">
            <span class="icon is-small"><i class="fa fa-plus"></i></span>
            <span>Add User</span>
          </a>
        </div>
      </div>
    </nav>
    <div class="columns is-marginless">
      <div class="column is-paddingless">
        <div class="card is-fullwidth">
          <table class="table is-striped">
            <thead>
            <tr>
              <th>Name</th>
              <th>Email</th>
              <th>&nbsp;</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="user in users">
              <td>{{ user.firstName }} {{ user.lastName }}</td>
              <td>{{ user.email }}</td>
              <td>
                <a class="button is-small is-dark" @click.prevent="editUser(user)">
                  <span class="icon is-small"><i class="fa fa-edit"></i></span>
                  <span>Edit</span>
                </a>
                <a class="button is-small is-danger" @click.prevent="doDeleteUser(user.id)">
                  <span class="icon is-small"><i class="fa fa-times"></i></span>
                  <span>Delete</span>
                </a>
              </td>
            </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped>
</style>
<script>
  import { mapGetters, mapActions } from 'vuex'
  export default {
    computed: mapGetters(['users']),
    methods: {
      createUser () {
        this.unsetUser()
        this.$router.push({ name: 'createUser' })
      },
      editUser (user) {
        this.setUser(user)
        this.$router.push({name: 'editUser', params: {id: user.id}})
      },
      doDeleteUser (id) {
        this.deleteUser(id)
      },
      ...mapActions(['loadUsers', 'setUser', 'unsetUser', 'deleteUser'])
    },
    mounted () {
      this.unsetUser()
      this.loadUsers()
    }
  }
</script>

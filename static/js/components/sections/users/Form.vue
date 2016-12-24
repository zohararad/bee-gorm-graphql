<template>
  <div>
    <nav class="level">
      <div class="level-left">
        <div class="level-item">
          <h3 class="subtitle is-3">{{ action }}</h3>
        </div>
      </div>
    </nav>
    <div class="columns is-marginless">
      <div class="column is-paddingless">
        <form class="card is-fullwidth" autocomplete="off">
          <div class="card-content">
            <div class="content">
              <div class="control is-horizontal">
                <div class="control-label">
                  <label class="label" for="firstName">Full Name</label>
                </div>
                <div class="control is-grouped">
                  <p class="control is-expanded">
                    <input id="firstName" class="input" type="text" placeholder="First Name" v-model="firstName">
                  </p>
                  <p class="control is-expanded">
                    <input id="lastName" class="input" type="text" placeholder="Last Name" v-model="lastName">
                  </p>
                </div>
              </div>
              <div class="control is-horizontal">
                <div class="control-label">
                  <label class="label" for="email">Email</label>
                </div>
                <div class="control is-grouped">
                  <p class="control is-expanded">
                    <input id="email" class="input" type="email" placeholder="Email" autocomplete="new-email" v-model="email">
                  </p>
                </div>
              </div>
              <div class="control is-horizontal">
                <div class="control-label">
                  <label class="label" for="password">Password</label>
                </div>
                <div class="control is-grouped">
                  <p class="control is-expanded">
                    <input id="password" class="input" type="password" placeholder="Password" autocomplete="new-password" v-model="password">
                  </p>
                </div>
              </div>
            </div>
          </div>
          <div class="card-footer form-actions has-text-right">
            <button class="button is-success" @click.prevent="saveUser">
              <span class="icon is-small">
                <i class="fa fa-save"></i>
                <i class="fa fa-check"></i>
              </span>
              <span>Save</span>
            </button>
            <router-link to="/users" class="button is-light">
              Cancel
            </router-link>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>
<style scoped>
  .form-actions {
    padding: 10px 20px;
    justify-content: flex-end;
  }
  .button {
    transition: background-color 0.5s ease;
  }
  .button + .button {
    margin-left: 5px;
  }
  .is-success .fa-check,
  .is-danger .fa-save{
    display: none;
  }
</style>
<script>
  import { mapGetters, mapActions } from 'vuex'
  export default {
    data () {
      return {
        firstName: '',
        lastName: '',
        email: '',
        password: ''
      }
    },
    computed: {
      action () {
        return this.$route.params.id ? 'Edit User' : 'Create User'
      },
      ...mapGetters(['user', 'saved'])
    },
    methods: {
      saveUser () {
        let newUser = !(this.user && this.user.id)
        let passwordFilled = this.password.trim().length > 0
        let user = {
          firstName: this.firstName,
          lastName: this.lastName,
          email: this.email
        }
        if (passwordFilled) {
          user.password = this.password.trim()
        }
        if (newUser && passwordFilled) {
          this.createUser(user)
        } else if (!newUser) {
          user.id = this.user.id
          this.updateUser(user)
        }
      },
      initUser () {
        this.firstName = this.user.firstName
        this.lastName = this.user.lastName
        this.email = this.user.email
      },
      ...mapActions(['loadUser', 'createUser', 'updateUser', 'setUserSaved'])
    },
    watch: {
      user () {
        this.initUser()
      },
      saved () {
        if (this.saved) {
          let router = this.$router
          let button = this.$el.querySelector('.is-success')
          button.classList.remove('is-success')
          button.classList.add('is-danger')
          setTimeout(() => {
            router.push({name: 'users'})
          }, 1500)
        }
      }
    },
    mounted () {
      this.setUserSaved(false)
      if (this.user === null && this.$route.params.id) {
        this.loadUser(this.$route.params.id)
      } else if (this.user !== null) {
        this.initUser()
      }
    }
  }
</script>

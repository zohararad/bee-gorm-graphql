import * as types from '../mutation-types'
import api from '../../api/users'

// initial state
const state = {
  users: [],
  user: null,
  saved: false
}

// getters
const getters = {
  users: state => state.users,
  user: state => state.user,
  saved: state => state.saved
}

// actions
const actions = {
  loadUsers ({ commit }) {
    api.loadUsers((users) => {
      commit(types.RECEIVE_USERS, users)
    })
  },
  loadUser ({ commit }, id) {
    api.loadUser({ id }, (user) => {
      commit(types.SET_USER, user)
    })
  },
  createUser ({ commit }, userToCreate) {
    api.createUser(userToCreate, (user) => {
      commit(types.CREATE_USER, user)
      commit(types.USER_SAVED, true)
    })
  },
  updateUser ({ commit }, userToUpdate) {
    api.updateUser(userToUpdate, (user) => {
      commit(types.UPDATE_USER, user)
      commit(types.USER_SAVED, true)
    })
  },
  deleteUser ({ commit }, id) {
    api.deleteUser({ id }, (id) => {
      commit(types.DELETE_USER, id)
    })
  },
  setUser ({ commit }, user) {
    commit(types.SET_USER, user)
  },
  unsetUser ({ commit }) {
    commit(types.UNSET_USER)
  },
  setUserSaved ({ commit }, saved) {
    commit(types.USER_SAVED, saved)
  }
}

// mutations
const mutations = {
  [types.RECEIVE_USERS] (state, users) {
    state.users = users
  },
  [types.SET_USER] (state, user) {
    state.user = user
  },
  [types.UNSET_USER] (state) {
    state.user = null
  },
  [types.CREATE_USER] (state, user) {
    state.user = user
  },
  [types.UPDATE_USER] (state, user) {
    state.user = user
  },
  [types.DELETE_USER] (state, id) {
    state.users = state.users.filter((user) => user.id !== id)
  },
  [types.USER_SAVED] (state, saved) {
    state.saved = saved
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}

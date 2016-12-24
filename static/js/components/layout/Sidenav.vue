<template>
  <aside class="menu app-sidebar animated slideInLeft">
    <ul class="menu-list">
      <li>
        <router-link to="/" active-class="is-active" :exact="true">
          <span class="icon is-small"><i class="fa fa-home"></i></span>
          <span>Home</span>
        </router-link>
        <router-link to="/users" active-class="is-active" :exact="true">
          <span class="icon is-small"><i class="fa fa-users"></i></span>
          <span>Users</span>
        </router-link>
      </li>
    </ul>
  </aside>
</template>

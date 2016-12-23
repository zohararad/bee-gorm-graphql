<template>
  <aside class="menu app-sidebar animated slideInLeft">
    <ul class="menu-list">
      <li>
        <router-link :to="{ name: 'users'}">
          <span class="icon is-small"><i class="fa fa-tachometer"></i></span>
          Users
        </router-link>
      </li>
    </ul>
  </aside>
</template>
<style>
  .menu-list a.router-link-active {
    background-color: #00d1b2;
    color: white;
  }
</style>
<script>
</script>

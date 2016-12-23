<template>
  <div id="app">
    <sidenav></sidenav>
    <main class="app-main">
      <div class="container is-fluid is-marginless app-content">
        <transition name="slide-fade">
          <router-view></router-view>
        </transition>
      </div>
    </main>
  </div>
</template>
<script>
import Sidenav from './components/layout/Sidenav'
export default {
  name: 'app',
  components: { Sidenav }
}
</script>
<template>
  <a-layout class="container">
    <Nav></Nav>
    <a-layout>
      <a-layout-header class="headerBtn">
        <Header></Header>
      </a-layout-header>
      <a-layout-content>
        <router-view :key="$route.path"> </router-view>
      </a-layout-content>
      <a-layout-footer>
        <Footer></Footer>
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script>
import Footer from '../components/admin/Footer.vue';
import Header from '../components/admin/Header.vue';
import Nav from '../components/admin/Nav.vue';
export default {
  components: { Nav, Footer, Header },
}
</script>

<style scoped>
.headerBtn{
  display: flex;
  justify-content: right;
  align-items: center;
}
.container {
  height: 100%;
}
</style>

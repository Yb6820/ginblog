<template>
  <v-footer padless color="indigo darken-2">
    <v-row justify="center" no-gutters>
      <v-col class="py-4 text-center white--text" cols="10"
        >{{ new Date().getFullYear() }} - GinBlog
      </v-col>

      <div class="text-center white--text">
        <a class="text-center white--text" href="https://beian.miit.gov.cn/" style="text-decoration: none;display: inline-block;">{{
          icp_record
        }}</a>
        <span>&nbsp;&nbsp;</span>
        <img src="http://img.centyoubet.xyz/beiangov.png" alt="">
        <a class="text-center white--text" href="https://www.beian.gov.cn/portal/registerSystemInfo?recordcode=44011302004001" target="_blank" style="text-decoration: none;"> 
          公安备案号{{ gov_record }} 
        </a>
      </div>
    </v-row>
  </v-footer>
</template>

<script>
export default {
  data() {
    return {
      icp_record: '',
      gov_record: '',
    }
  },
  mounted() {
    this.$root.$on('msg', (msg) => {
      this.icp_record = msg.icp_record
      this.gov_record = msg.gov_record
    })
  },
}
</script>

<style>
img{
  width: 16px;
  height: 16px;
  display: inline-block;
  margin-right: 4px;
  margin-bottom: 0px;
  padding: 0px;
}
</style>

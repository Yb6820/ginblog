<template>
  <v-col>
    <v-card class="ma-3" v-for="item in artList" :key="item.ID" link @click="$router.push(`detail/${item.ID}`)">
        <v-row no-gutters>
            <v-col class="d-flex justify-center align-center mx-3" cols="2">
                <v-img max-height="100" max-width="100" :src="item.img"></v-img>
            </v-col>
            <v-col>
                <v-card-title class="my-2">
                    <v-chip color="pink" label class="mr-3 white--text">{{ item.Category.name }}</v-chip>
                    {{ item.title }}
                </v-card-title>
                <v-card-subtitle>{{ item.desc }}</v-card-subtitle>
                <v-divider></v-divider>
                <v-card-text>
                    <v-icon>{{ 'mdi-calendar-month' }}</v-icon>
                    <span>{{ item.CreatedAt | dateformat('YYYY-MM-DD HH:mm') }}</span>
                </v-card-text>
            </v-col>
        </v-row>
    </v-card>
    <div class="text-center">
        <v-pagination 
        v-model="queryParam.pagenum" 
        :length="Math.ceil(this.total/queryParam.pagesize)"
        total-visible="5"
        @input="getArtlist()"
        ></v-pagination>
    </div>
  </v-col>
</template>

<script>
export default {
    data(){
        return{
            artList: [],
            queryParam: {
                pagesize:5,
                pagenum:1
            },
            total:0
        }
    },
    created(){
        this.getArtlist()
    },
    methods:{
        //获取文章列表
        async getArtlist(){
            const{data:res}=await this.$http.get('articles',{params:{
                pagesize:this.queryParam.pagesize,
                pagenum:this.queryParam.pagenum
            }})
            this.artList=res.data
            this.total=res.total
        }
    }
}
</script>

<style>

</style>
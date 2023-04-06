<template>
  <div>
    <a-card>
        <h3>个人设置</h3>
        <a-form-model>
                <a-form-model-item label="作者名称" v-model="profileInfo.name">
                    <a-input style="width:300px"></a-input>
                </a-form-model-item>
                
                <a-form-model-item label="个人简介" v-model="profileInfo.desc">
                    <a-input type="textarea"></a-input>
                </a-form-model-item>

                <a-form-model-item label="QQ号码" v-model="profileInfo.qq_chat">
                    <a-input style="width:300px"></a-input>
                </a-form-model-item>

                <a-form-model-item label="微信" v-model="profileInfo.wechat">
                    <a-input style="width:300px"></a-input>
                </a-form-model-item>

                <a-form-model-item label="微博" v-model="profileInfo.weibo">
                    <a-input style="width:300px"></a-input>
                </a-form-model-item>

                <a-form-model-item label="B站地址" v-model="profileInfo.bili">
                    <a-input style="width:300px"></a-input>
                </a-form-model-item>

                <a-form-model-item label="email" v-model="profileInfo.email">
                    <a-input style="width:300px"></a-input>
                </a-form-model-item>

                <a-form-model-item label="头像">
                    <a-upload
                    list-type="picture"
                    name="file"
                    :action="upUrl"
                    :headers="headers"
                    @change="avatarChange"
                    >
                    <a-button> <a-icon type="upload" /> 点击上传 </a-button>
                    <br/>
                    <template v-if="profileInfo.avatar">
                      <img :src="profileInfo.avatar" style="width:150px;height:80px">
                    </template>
                    </a-upload>
                </a-form-model-item>

                <a-form-model-item label="头像背景图">
                    <a-upload
                    list-type="picture"
                    name="file"
                    :action="upUrl"
                    :headers="headers"
                    @change="imgChange"
                    >
                    <a-button> <a-icon type="upload" /> 点击上传 </a-button>
                    <br/>
                    <template v-if="profileInfo.img">
                      <img :src="profileInfo.img" style="width:150px;height:80px">
                    </template>
                    </a-upload>
                </a-form-model-item>

                <a-form-model-item>
                    <a-button type="primary" style="margin-right:15px" @click="updateProfile"> 更新</a-button>
                </a-form-model-item>
            </a-form-model>
    </a-card>
  </div>
</template>

<script>
import { Url } from '../../plugins/http'
export default {
data(){
    return{
        profileInfo:{
            id:1,
            name:'',
            desc:'',
            qq_chat:'',
            wechat:'',
            weibo:'',
            bili:'',
            email:'',
            img:'',
            avatar:''
        },
        upUrl: Url + 'upload',
        headers: {}
    }
},
created(){
    this.getProfileInfo()
    this.headers = { Authorization: `Bearer ${window.sessionStorage.getItem('token')}` }
},
methods:{
    // 查询文章信息
    async getProfileInfo () {
      const { data: res } = await this.$http.get(`profile/${this.profileInfo.id}`)
      if (res.status !== 200) return this.$message.error(res.message)
      this.profileInfo = res.data
    },
    // 上传头像
    avatarChange(info) {
            if (info.file.status !== 'uploading') {
                console.log(info.file, info.fileList)
            }
            if (info.file.status === 'done') {
                this.$message.success('图片上传成功')
                this.profileInfo.avatar = info.file.response.url
            } else if (info.file.status === 'error') {
                this.$message.error('图片上传失败')
            }
        },

        //上传头像背景图
        imgChange(info) {
            if (info.file.status !== 'uploading') {
                console.log(info.file, info.fileList)
            }
            if (info.file.status === 'done') {
                this.$message.success('图片上传成功')
                this.profileInfo.img = info.file.response.url
            } else if (info.file.status === 'error') {
                this.$message.error('图片上传失败')
            }
        },
    // 更新
    async updateProfile(){
        const {data:res}=await this.$http.put(`profile/${this.profileInfo.id}`,this.profileInfo)
        if(res.status!==200)return this.$message.error(res.message)
        this.$message.success('个人信息更新成功')
        this.$router.push('/index')
    }
}
}
</script>

<style>

</style>
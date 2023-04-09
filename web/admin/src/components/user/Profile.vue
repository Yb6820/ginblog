<template>
  <div>
    <a-card>
        <a-form-model labelAlign="left" :label-col="{ span: 2 }" :wrapper-col="{ span: 12 }">
                <a-form-model-item label="作者名称">
                    <a-input style="width:300px" v-model="profileInfo.name"></a-input>
                </a-form-model-item>
                
                <a-form-model-item label="个人简介">
                    <a-input type="textarea" v-model="profileInfo.desc"></a-input>
                </a-form-model-item>

                <a-form-model-item label="QQ号码">
                    <a-input style="width:300px" v-model="profileInfo.qq_chat"></a-input>
                </a-form-model-item>

                <a-form-model-item label="微信">
                    <a-input style="width:300px" v-model="profileInfo.wechat"></a-input>
                </a-form-model-item>

                <a-form-model-item label="微博">
                    <a-input style="width:300px" v-model="profileInfo.weibo"></a-input>
                </a-form-model-item>

                <a-form-model-item label="B站地址">
                    <a-input style="width:300px" v-model="profileInfo.bili"></a-input>
                </a-form-model-item>

                <a-form-model-item label="email">
                    <a-input style="width:300px" v-model="profileInfo.email"></a-input>
                </a-form-model-item>

                <a-form-model-item label="头像">
                    <a-upload
                    name="file"
                    list-type="picture-card"
                    class="avatar-uploader"
                    :show-upload-list="false"
                    :action="upUrl"
                    :headers="headers"
                    @change="avatarChange"
                    >
                    <img style="width:120px;height:120px" v-if="profileInfo.avatar" :src="profileInfo.avatar" alt="avatar" />
                    <div v-else>
                        <a-icon :type="loading ? 'loading' : 'plus'" />
                        <div class="ant-upload-text">
                            Upload
                        </div>
                    </div>
                    </a-upload>
                </a-form-model-item>

                <a-form-model-item label="头像背景图">
                    <a-upload
                    name="file"
                    list-type="picture-card"
                    class="avatar-uploader"
                    :show-upload-list="false"
                    :action="upUrl"
                    :headers="headers"
                    @change="imgChange"
                    >
                    <img style="width:150px;height:120px" v-if="profileInfo.img" :src="profileInfo.img" alt="avatar" />
                    <div v-else>
                        <a-icon :type="loading ? 'loading' : 'plus'" />
                        <div class="ant-upload-text">
                            Upload
                        </div>
                    </div>
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
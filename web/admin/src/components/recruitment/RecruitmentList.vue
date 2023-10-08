<template>
  <div>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search
            v-model="queryParam.username"
            placeholder="输入用户名查找"
            enter-button
            allowClear
            @search="getUserList"/>
        </a-col>
        <a-col :span="4">
          <a-button type="primary" @click="addUserVisible=true">新增</a-button>
        </a-col>
      </a-row>

      <a-table
        :columns="columns"
        :pagination='pagination'
        :dataSource="userList"
        rowKey="ID"
        bordered
        @change="handleTableChange"
      >
        <span slot="role" slot-scope="role">{{ role==1?'管理员':'订阅者' }}</span>
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button type="primary" icon="edit" style="margin-right:15px" @click="editUser(data.ID)">编辑</a-button>
            <a-button type="danger" icon="delete" style="margin-right:15px" @click="deleteUser(data.ID)">删除</a-button>
            <a-button type="info" icon="info" @click="reSetPassword(data.ID)">修改密码</a-button>
          </div>
        </template>
      </a-table>
    </a-card>
    <!--   新增用户区域   -->
    <a-modal
      closable
      :visible="addUserVisible"
      width="60%"
      title="新增用户"
      @ok="addUserOk"
      @cancel="addUserCancel"
    >
      <a-form-model :model="newuser" :rules="addUserRules" ref="addUserRef">
        <a-form-model-item label="用户名" prop="username">
          <a-input v-model="newuser.username"></a-input>
        </a-form-model-item>
        <a-form-model-item has-feedback label="密码" prop="password">
          <a-input-password v-model="newuser.password"></a-input-password>
        </a-form-model-item>
        <a-form-model-item has-feedback label="确认密码" prop="checkpass">
          <a-input-password v-model="newuser.checkpass"></a-input-password>
        </a-form-model-item>
        <a-form-model-item label="是否为管理员" prop="role">
          <a-switch checked-children="是" un-checked-children="否" @change="newRoleChange" />
        </a-form-model-item>
      </a-form-model>
    </a-modal>
    <!--   编辑用户区域   -->
    <a-modal
      closable
      :visible="editUserVisible"
      width="60%"
      title="编辑用户"
      @ok="editUserOk"
      @cancel="editUserCancel"
    >
      <a-form-model :model="userinfo" :rules="editUserRules" ref="editUserRef">
        <a-form-model-item label="用户名" prop="username">
          <a-input v-model="userinfo.username"></a-input>
        </a-form-model-item>
        <a-form-model-item label="是否为管理员" prop="role">
          <a-select :defaultValue=userinfo.role.toString() style="120px" @change="editRoleChange">
            <a-select-option key="1" value="1">是</a-select-option>
            <a-select-option key="2" value="2">否</a-select-option>
          </a-select>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '10%',
    key: 'id',
    align: 'center'
  },
  {
    title: '用户名',
    dataIndex: 'username',
    width: '20%',
    key: 'username',
    align: 'center'
  },
  {
    title: '角色',
    dataIndex: 'role',
    width: '20%',
    key: 'role',
    align: 'center',
    scopedSlots: { customRender: 'role' }
  },
  {
    title: '操作',
    width: '30%',
    key: 'action',
    align: 'center',
    scopedSlots: { customRender: 'action' }
  }
]

export default {
  data () {
    return {
      pagination: {
        pageSizeOptions: ['5', '10', '20'],
        defaultCurrent: 1,
        defaultPageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`
      },
      userList: [],
      columns,
      queryParam: {
        username: '',
        pagesize: 5,
        pagenum: 1
      },
      newuser: {
        username: '',
        password: '',
        checkpass: '',
        role: 2
      },
      userinfo: {
        id: 0,
        username: '',
        role: 2
      },
      addUserRules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 4, max: 12, message: '用户名应当在4到12之间', trigger: 'blur' }
        ],
        password: [{
          validator: (rule, value, callback) => {
            if (value === '') {
              callback(new Error('请输入密码'))
            }
            if (value.length < 6 || value.length > 20) {
              callback(new Error('密码应当在6到20为之间'))
            } else {
              callback()
            }
          },
          trigger: 'blur'
        }],
        checkpass: [{
          validator: (rule, value, callback) => {
            if (value === '') {
              callback(new Error('请输入密码'))
            }
            if (value !== this.newuser.password) {
              callback(new Error('密码不一致，请重新输入'))
            } else {
              callback()
            }
          },
          trigger: 'blur'
        }]
      },
      editUserRules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 4, max: 12, message: '用户名应当在4到12之间', trigger: 'blur' }
        ]
      },
      addUserVisible: false,
      editUserVisible: false,
      visible: false
    }
  },
  created () {
    this.getUserList()
  },
  methods: {
    async getUserList () {
      const { data: res } = await this.$http.get('users', {
        params: {
          username: this.queryParam.username,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 200) return this.$message.error(res.message)
      this.userList = res.data
      this.pagination.total = res.total
    },
    // 修改分页时触发
    handleTableChange (pagination, filters, sorter) {
      const pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.queryParam.pagesize = pagination.pageSize
      this.queryParam.pagenum = pagination.current

      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1
        pager.current = 1
      }
      this.pagination = pager
      this.getUserList()
    },
    // 删除用户
    deleteUser (id) {
      this.$confirm({
        title: '确定删除该用户吗？',
        content: '一旦删除，无法恢复',
        onOk: async () => {
          const { data: res } = await this.$http.delete(`user/${id}`)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$message.success('删除成功')
          this.getUserList()
        },
        onCancel: async () => {
          this.$message.info('已取消删除')
        }
      })
    },
    // 新增用户
    addUserOk () {
      this.$refs.addUserRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.post('/user/add', {
          username: this.newuser.username,
          password: this.newuser.password,
          role: this.newuser.role
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.addUserVisible = false
        this.$message.info('添加用户成功')
        this.getUserList()
      })
    },
    addUserCancel () {
      this.$refs.addUserRef.resetFields()
      this.addUserVisible = false
      this.$message.info('新增用户已取消')
    },
    newRoleChange (checked) {
      this.newuser.role = checked ? 1 : 2
    },
    // 编辑用户
    async editUser (id) {
      this.editUserVisible = true
      const { data: res } = await this.$http.get(`/user/${id}`)
      if (res.status !== 200) this.$message.error(res.message)
      this.userinfo.username = res.data.username
      this.userinfo.role = res.data.role
      this.userinfo.id = id
    },
    editUserOk () {
      this.$refs.editUserRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.put(`user/${this.userinfo.id}`, {
          username: this.userinfo.username,
          role: this.userinfo.role
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.editUserVisible = false
        this.$message.info('编辑用户信息成功')
        this.getUserList()
      })
    },
    editRoleChange (value) {
      this.userinfo.role = Number(value)
    },
    editUserCancel () {
      this.editUserVisible = false
      this.$refs.editUserRef.resetFields()
      this.$message.info('编辑已取消')
    }
  }
}

</script>

<style scoped>
.actionSlot{
  display: flex;
  justify-content: center;
}
</style>

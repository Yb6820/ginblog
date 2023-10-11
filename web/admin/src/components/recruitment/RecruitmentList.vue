<template>
  <div>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="5">
          <a-input
            v-model="queryParam.companyName"
            placeholder="输入公司名称"
            @keyup.enter="getRecruitmentList"
          ></a-input>
        </a-col>
        <a-col :span="5">
          <a-input
            v-model="queryParam.deliveryJob"
            placeholder="输入投递职位名称"
            @keyup.enter="getRecruitmentList"
          ></a-input>
        </a-col>
        <a-col :span="5">
          <a-date-picker
            style="width: 220px"
            v-model="queryParam.deliveryTime"
            placeholder="输入投递时间"
            @change="getRecruitmentList"
          ></a-date-picker>
        </a-col>
        <a-col :span="5">
          <a-select
            v-model="queryParam.deliveryStatus"
            placeholder="选择投递状态"
            style="width: 220px"
            enter-button
            allowClear
            @change="getRecruitmentList">
            <a-select-option v-for="item in recruitmentStatusLabels" :key="item.key">
              {{ item.label }}
            </a-select-option>
          </a-select>
        </a-col>
        <a-col :span="4">
          <a-button type="primary" @click="addRecruitmentVisible=true">新增</a-button>
        </a-col>
      </a-row>

      <a-table
        :columns="columns"
        :pagination='pagination'
        :dataSource="recruitmentList"
        rowKey="ID"
        bordered
        @change="handleTableChange"
      >
        <span slot="role" slot-scope="role">{{ role==1?'管理员':'订阅者' }}</span>
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button type="primary" icon="edit" style="margin-right:15px" @click="editRecruitment(data.ID)">编辑</a-button>
            <a-button type="danger" icon="delete" style="margin-right:15px" @click="deleteRecruitment(data.ID)">删除</a-button>
            <a-button type="info" icon="info" @click="reSetPassword(data.ID)">修改密码</a-button>
          </div>
        </template>
      </a-table>
    </a-card>
    <!--   新增应聘记录区域   -->
    <a-modal
      closable
      :visible="addRecruitmentVisible"
      width="60%"
      title="新增应聘记录"
      @ok="addRecruitmentOk"
      @cancel="addRecruitmentCancel"
    >
      <a-form-model :model="newRecruitment" :rules="addRecruitmentRules" ref="addRecruitmentRef">
        <a-form-model-item label="投递公司" prop="companyName">
          <a-input v-model="newRecruitment.companyName"></a-input>
        </a-form-model-item>
        <a-form-model-item label="投递职位" prop="deliveryJob">
          <a-input v-model="newRecruitment.deliveryJob"></a-input>
        </a-form-model-item>
        <a-form-model-item label="第几自愿" prop="deliveryVoluntary">
          <a-select v-model="newRecruitment.voluntaryNum">
            <a-select-option v-for="item in voluntaries" :key="item.key">{{ item.label }}</a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item label="投递状态">
          <a-select v-model="newRecruitment.deliveryStatus">
            <a-select-option v-for="item in recruitmentStatusLabels"   :key="item.key">{{ item.label }}</a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item has-feedback label="投递时间" prop="checkpass">
          <a-input v-model="newRecruitment.deliveryTime"></a-input>
        </a-form-model-item>
        <a-form-model-item label="投递结果">
          <a-input v-model="newRecruitment.deliveryRes"></a-input>
        </a-form-model-item>
        <a-form-model-item label="原因">
          <a-input v-model="newRecruitment.reason"></a-input>
        </a-form-model-item>
        <a-form-model-item label="其他描述">
          <a-textarea v-model="newRecruitment.desc" :auto-size="{ minRows: 3,maxRows:5 }"></a-textarea>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
    <!--   编辑用户区域   -->
    <a-modal
      closable
      :visible="editRecruitmentVisible"
      width="60%"
      title="编辑记录"
      @ok="editRecruitmentOk"
      @cancel="editRecruitmentCancel"
    >
      <a-form-model :model="Recruitmentinfo" :rules="editRecruitmentRules" ref="editRecruitmentRef">
        <a-form-model-item label="用户名" prop="Recruitmentname">
          <a-input v-model="Recruitmentinfo.Recruitmentname"></a-input>
        </a-form-model-item>
        <a-form-model-item label="是否为管理员" prop="role">
          <a-select :defaultValue=Recruitmentinfo.role.toString() style="width: 120px" @change="editRoleChange">
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
    width: '5%',
    key: 'id',
    align: 'center'
  },
  {
    title: '公司名称',
    dataIndex: 'company_name',
    width: '10%',
    key: 'company_name',
    align: 'center'
  },
  {
    title: '投递职位',
    dataIndex: 'job_name',
    width: '10%',
    key: 'job_name',
    align: 'center',
  },
  {
    title: '投递日期',
    dataIndex: 'delivery_time',
    width: '10%',
    key: 'delivery_time',
    align: 'center',
  },
  {
    title: '投递结果',
    dataIndex: 'result',
    width: '10%',
    key: 'result',
    align: 'center',
  },
  {
    title: '投递状态',
    dataIndex: 'status',
    width: '10%',
    key: 'status',
    align: 'center',
  },
  {
    title: '原因',
    dataIndex: 'reason',
    width: '10%',
    key: 'reason',
    align: 'center',
  },
  {
    title: '描述',
    dataIndex: 'desc',
    width: '10%',
    key: 'desc',
    align: 'center'
  },
  {
    title: '操作',
    width: '15%',
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
      recruitmentStatusLabels:[
        { key:0, label:'全部' },
        { key:1, label:'已投递' },
      ],
      voluntaries:[
        { key:1 ,label: '第一自愿' },
        { key:2 ,label: '第二自愿' },
        { key:3 ,label: '第三自愿'}
      ],
      recruitmentList: [],
      columns,
      queryParam: {
        companyName: '',
        deliveryJob:'',
        deliveryTime:'',
        deliveryStatus:0,
        pagesize: 5,
        pagenum: 1
      },
      newRecruitment: {
        companyName: '',
        voluntaryNum:1,
        deliveryTime: '',
        deliveryJob: '',
        deliveryStatus: 1,
        deliveryRes:'',
        reason:'',
        desc:'',
      },
      Recruitmentinfo: {
        id: 0,
        Recruitmentname: '',
        role: 2
      },
      addRecruitmentRules: {
        companyName: [
          { required: true, message: '请输入公司名称', trigger: 'blur' },
          { min: 2, max: 50, message: '公司名称不能为空', trigger: 'blur' }
        ],
        deliveryJob:[
          { required:true, message:'请输入投递的职位' ,trigger:'blur' }
        ],
        deliveryVoluntary:[
          { required:true, message:'请选择第几自愿', trigger:'blur' }
        ],
        deliveryTime:[
          { required:true,message:'请输入投递时间',trigger:'blur' }
        ],
        deliveryStatus:[
          { required:true,message:'请输入投递状态',trigger:'blur' }
        ],
      },
      editRecruitmentRules: {
        Recruitmentname: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 4, max: 12, message: '用户名应当在4到12之间', trigger: 'blur' }
        ]
      },
      addRecruitmentVisible: false,
      editRecruitmentVisible: false,
      visible: false
    }
  },
  created () {
    this.getRecruitmentList()
  },
  methods: {
    async getRecruitmentList () {
      const { data: res } = await this.$http.get('recruitments', {
        params: {
          company_name: this.queryParam.companyName,
          delivery_job:this.queryParam.deliveryJob,
          delivery_time:this.queryParam.deliveryTime,
          delivery_status:this.queryParam.deliveryStatus,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 200) return this.$message.error(res.message)
      this.recruitmentList = res.data
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
      this.getRecruitmentList()
    },
    // 删除用户
    deleteRecruitment (id) {
      this.$confirm({
        title: '确定删除该用户吗？',
        content: '一旦删除，无法恢复',
        onOk: async () => {
          const { data: res } = await this.$http.delete(`Recruitment/${id}`)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$message.success('删除成功')
          this.getRecruitmentList()
        },
        onCancel: async () => {
          this.$message.info('已取消删除')
        }
      })
    },
    // 新增用户
    addRecruitmentOk () {
      this.$refs.addRecruitmentRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.post('/recruitment/add', {
          company_name: this.newRecruitment.companyName,
          voluntary: this.newRecruitment.voluntaryNum,
          job_name: this.newRecruitment.deliveryJob,
          delivery_time:this.newRecruitment.deliveryTime,
          result:this.newRecruitment.deliveryRes,
          status:this.newRecruitment.deliveryStatus,
          reason:this.newRecruitment.reason,
          desc:this.newRecruitment.desc,
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.addRecruitmentVisible = false
        this.$message.info('添加应聘记录成功')
        this.getRecruitmentList()
      })
    },
    addRecruitmentCancel () {
      this.$refs.addRecruitmentRef.resetFields()
      this.addRecruitmentVisible = false
      this.$message.info('新增用户已取消')
    },
    newRoleChange (checked) {
      this.newRecruitment.role = checked ? 1 : 2
    },
    // 编辑用户
    async editRecruitment (id) {
      this.editRecruitmentVisible = true
      const { data: res } = await this.$http.get(`/Recruitment/${id}`)
      if (res.status !== 200) this.$message.error(res.message)
      this.Recruitmentinfo.Recruitmentname = res.data.Recruitmentname
      this.Recruitmentinfo.role = res.data.role
      this.Recruitmentinfo.id = id
    },
    editRecruitmentOk () {
      this.$refs.editRecruitmentRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.put(`Recruitment/${this.Recruitmentinfo.id}`, {
          Recruitmentname: this.Recruitmentinfo.Recruitmentname,
          role: this.Recruitmentinfo.role
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.editRecruitmentVisible = false
        this.$message.info('编辑用户信息成功')
        this.getRecruitmentList()
      })
    },
    editRoleChange (value) {
      this.Recruitmentinfo.role = Number(value)
    },
    editRecruitmentCancel () {
      this.editRecruitmentVisible = false
      this.$refs.editRecruitmentRef.resetFields()
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

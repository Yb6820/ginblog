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
            format="YYYY-MM-DD"
            value-format="YYYY-MM-DD"
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
        rowKey="id"
        bordered
        @change="handleTableChange"
      >
        <template slot="deliveryTime" slot-scope="text">
          <span>{{ text }}</span>
        </template>
        <template slot="voluntary_num" slot-scope="text">{{ voluntaries[text-1].label }}</template>
        <template slot="status" slot-scope="text">{{ recruitmentStatusLabels[text].label }}</template>
        <template slot="action" slot-scope="text,record">
          <div class="actionSlot">
            <a-button type="primary" icon="edit" size="small" style="margin-right:15px" @click="editRecruitment(record)">编辑</a-button>
            <a-button type="primary" size="small" @click="nextRecruitment(record)">
              下一步<a-icon type="right" />
            </a-button>
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
            <a-select-option v-for="item in voluntaries" :key="item.key" :value="item.key">{{ item.label }}</a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item label="投递状态">
          <a-select v-model="newRecruitment.deliveryStatus" prop="deliveryStatus">
            <a-select-option v-for="item in recruitmentStatusLabels" :key="item.key">{{ item.label }}</a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item has-feedback label="投递时间" prop="deliveryTime">
          <a-date-picker v-model="newRecruitment.deliveryTime" format="YYYY-MM-DD" value-format="YYYY-MM-DD"></a-date-picker>
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

    <!--   编辑投递信息区域   -->
    <a-modal
      closable
      :visible="editRecruitmentVisible"
      width="60%"
      title="编辑记录"
      @ok="editRecruitmentOk"
      @cancel="editRecruitmentCancel"
    >
      <a-form-model :model="recruitmentInfo" :rules="editRecruitmentRules" ref="editRecruitmentRef">
        <a-form-model-item label="投递公司" prop="companyName">
          <a-input v-model="recruitmentInfo.companyName"></a-input>
        </a-form-model-item>
        <a-form-model-item label="投递职位" prop="deliveryJob">
          <a-input v-model="recruitmentInfo.deliveryJob"></a-input>
        </a-form-model-item>
        <a-form-model-item label="第几自愿" prop="deliveryVoluntary">
          <a-select v-model="recruitmentInfo.voluntaryNum">
            <a-select-option v-for="item in voluntaries" :key="item.key">{{ item.label }}</a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item label="投递状态">
          <span>{{ recruitmentStatusLabels[recruitmentInfo.deliveryStatus].label }}</span>
        </a-form-model-item>
        <a-form-model-item has-feedback label="投递时间" prop="deliveryTime">
          <a-date-picker v-model="recruitmentInfo.deliveryTime" format="YYYY-MM-DD" value-format="YYYY-MM-DD"></a-date-picker>
        </a-form-model-item>
        <a-form-model-item label="投递结果">
          <a-input v-model="recruitmentInfo.deliveryRes"></a-input>
        </a-form-model-item>
        <a-form-model-item label="原因">
          <a-input v-model="recruitmentInfo.reason"></a-input>
        </a-form-model-item>
        <a-form-model-item label="其他描述">
          <a-textarea v-model="recruitmentInfo.desc" :auto-size="{ minRows: 3,maxRows:5 }"></a-textarea>
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!--   投递信息下一步区域，投递的基本信息不可修改   -->
    <a-modal
      closable
      :visible="nextStepVisible"
      width="60%"
      title="下一步应聘状态"
      @ok="nextRecruitmentOk"
      @cancel="nextRecruitmentCancel"
    >
      <a-form-model :model="nextRecruitmentInfo" :rules="nextRecruitmentRules" ref="nextRecruitmentRef">
        <a-form-model-item label="投递公司" prop="companyName">
          <span>{{ nextRecruitmentInfo.companyName }}</span>
        </a-form-model-item>
        <a-form-model-item label="投递职位" prop="deliveryJob">
          <span>{{ nextRecruitmentInfo.deliveryJob }}</span>
        </a-form-model-item>
        <a-form-model-item label="第几自愿" prop="deliveryVoluntary">
          <span>{{ voluntaries[nextRecruitmentInfo.voluntaryNum-1].label }}</span>
        </a-form-model-item>
        <a-form-model-item label="投递状态">
          <a-select v-model="nextRecruitmentInfo.deliveryStatus" prop="deliveryStatus">
            <a-select-option v-for="item in recruitmentStatusLabels" :key="item.key">{{ item.label }}</a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item has-feedback label="投递时间" prop="deliveryTime">
          <a-date-picker v-model="nextRecruitmentInfo.deliveryTime" format="YYYY-MM-DD" value-format="YYYY-MM-DD"></a-date-picker>
        </a-form-model-item>
        <a-form-model-item label="投递结果">
          <a-input v-model="nextRecruitmentInfo.deliveryRes"></a-input>
        </a-form-model-item>
        <a-form-model-item label="原因">
          <a-input v-model="nextRecruitmentInfo.reason"></a-input>
        </a-form-model-item>
        <a-form-model-item label="其他描述">
          <a-textarea v-model="nextRecruitmentInfo.desc" :auto-size="{ minRows: 3,maxRows:5 }"></a-textarea>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
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
    scopedSlots: { customRender: 'deliveryTime' },
  },
  {
    title: '第几自愿',
    dataIndex: 'voluntary_num',
    width: '10%',
    key: 'voluntary_num',
    align: 'center',
    scopedSlots: { customRender: 'voluntary_num' },
  },
  {
    title: '投递状态',
    dataIndex: 'status',
    width: '10%',
    key: 'status',
    align: 'center',
    scopedSlots: { customRender: 'status' },
  },
  {
    title: '投递结果',
    dataIndex: 'result',
    width: '10%',
    key: 'result',
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
        { key:1, label: '已投递' },
        { key:2, label:'待测评' },
        { key:3, label:'已测评' },
        { key:4, label:'待笔试' },
        { key:5, label:'笔试完成' },
        { key:6, label:'笔试挂' },
        { key:7, label:'待一面' },
        { key:8, label:'一面完成' },
        { key:9, label:'一面挂' },
        { key:10, label:'待二面' },
        { key:11, label:'二面完成' },
        { key:12, label:'二面挂' },
        { key:13, label:'待三面' },
        { key:14, label:'三面完成' },
        { key:15, label:'三面挂' },
        { key:16, label:'待hr' },
        { key:17, label:'hr面完成' },
        { key:18, label:'hr面挂'  }
      ],
      voluntaries:[
        { key:1 ,label: '第一自愿' },
        { key:2 ,label: '第二自愿' },
        { key:3 ,label: '第三自愿' }
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
      recruitmentInfo: {
        id: 0,
        companyName: '',
        voluntaryNum:1,
        deliveryTime: '',
        deliveryJob: '',
        deliveryStatus: 1,
        deliveryRes:'',
        reason:'',
        desc:'',
      },
      nextRecruitmentInfo:{
        id: 0,
        companyName: '',
        voluntaryNum:1,
        deliveryTime: '',
        deliveryJob: '',
        deliveryStatus: 1,
        deliveryRes:'',
        reason:'',
        desc:'',
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
          { required:false, message:'请选择第几自愿', trigger:'change' }
        ],
        deliveryTime:[
          { required:true,message:'请输入投递时间',trigger:'change' }
        ],
        deliveryStatus:[
          { required:false,message:'请输入投递状态',trigger:'change' }
        ],
      },
      editRecruitmentRules: {
        companyName: [
          { required: true, message: '请输入公司名称', trigger: 'blur' },
          { min: 2, max: 50, message: '公司名称不能为空', trigger: 'blur' }
        ],
        deliveryJob:[
          { required:true, message:'请输入投递的职位' ,trigger:'blur' }
        ],
        deliveryVoluntary:[
          { required:false, message:'请选择第几自愿', trigger:'change' }
        ],
        deliveryTime:[
          { required:true,message:'请输入投递时间',trigger:'change' }
        ],
        deliveryStatus:[
          { required:false,message:'请输入投递状态',trigger:'change' }
        ]
      },
      nextRecruitmentRules:{
        deliveryTime:[
          { required:true,message:'请输入投递时间',trigger:'change' }
        ],
        deliveryStatus:[
          { required:false,message:'请输入投递状态',trigger:'change' }
        ]
      },
      addRecruitmentVisible: false,
      editRecruitmentVisible: false,
      nextStepVisible:false,
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
    // 新增应聘信息
    addRecruitmentOk () {
      this.$refs.addRecruitmentRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        console.log(this.newRecruitment)
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
      this.$message.info('新增应聘信息已取消')
    },
    newRoleChange (checked) {
      this.newRecruitment.role = checked ? 1 : 2
    },
    // 编辑用户
    editRecruitment (data) {
      this.editRecruitmentVisible = true
      this.recruitmentInfo.id=data.id
      this.recruitmentInfo.companyName=data.company_name
      this.recruitmentInfo.deliveryJob=data.job_name
      this.recruitmentInfo.deliveryStatus=data.status
      this.recruitmentInfo.deliveryTime=data.delivery_time
      this.recruitmentInfo.voluntaryNum=data.voluntary_num
      this.recruitmentInfo.deliveryRes=data.result
      this.recruitmentInfo.reason=data.reason
      this.recruitmentInfo.desc=data.desc
    },
    editRecruitmentOk () {
      this.$refs.editRecruitmentRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        console.log("编辑用户信息时的参数值",this.recruitmentInfo)
        const { data: res } = await this.$http.post(`recruitment/edit`, {
          recruit_id:this.recruitmentInfo.id,
          company_name:this.recruitmentInfo.companyName,
          job_name:this.recruitmentInfo.deliveryJob,
          voluntary:this.recruitmentInfo.voluntaryNum,
          delivery_time:this.recruitmentInfo.deliveryTime,
          result:this.recruitmentInfo.deliveryRes,
          status:this.recruitmentInfo.deliveryStatus,
          reason:this.recruitmentInfo.reason,
          desc:this.recruitmentInfo.desc,
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.editRecruitmentVisible = false
        this.$message.info('编辑应聘信息成功')
        this.getRecruitmentList()
      })
    },
    editRecruitmentCancel () {
      this.editRecruitmentVisible = false
      this.$refs.editRecruitmentRef.resetFields()
      this.$message.info('编辑已取消')
    },
    nextRecruitment(data){
      this.nextStepVisible=true
      this.nextRecruitmentInfo.id=data.id
      this.nextRecruitmentInfo.companyName=data.company_name
      this.nextRecruitmentInfo.deliveryJob=data.job_name
      this.nextRecruitmentInfo.deliveryStatus=data.status
      this.nextRecruitmentInfo.deliveryTime=data.delivery_time
      this.nextRecruitmentInfo.voluntaryNum=data.voluntary_num
      this.nextRecruitmentInfo.deliveryRes=data.result
      this.nextRecruitmentInfo.reason=data.reason
      this.nextRecruitmentInfo.desc=data.desc
    },
    nextRecruitmentOk(){
      this.$refs.nextRecruitmentRef.validate(async (valid)=>{
        if (!valid)return this.$message.error("参数不符合要求，请重新输入")
        const { data: res }=await this.$http.post(`recruitment/next`,{
          recruit_id:this.nextRecruitmentInfo.id,
          delivery_time:this.nextRecruitmentInfo.deliveryTime,
          status:this.nextRecruitmentInfo.deliveryStatus,
          result:this.nextRecruitmentInfo.deliveryRes,
          reason:this.nextRecruitmentInfo.reason,
          desc:this.nextRecruitmentInfo.desc
        })
        if(res.status!=200)return this.$message.error(res.message)
        this.nextStepVisible=false
        this.$message.info('更改应聘状态成功')
        this.getRecruitmentList()
      })
    },
    nextRecruitmentCancel(){
      this.nextStepVisible=false
      this.$refs.nextRecruitmentRef.resetFields()
      this.$message.info('更改下一步状态取消')
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

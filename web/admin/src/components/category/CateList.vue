<template>
    <div>
        <a-card>
            <a-row :gutter="20">
                <a-col :span="4">
                    <a-button type="primary" @click="addCateVisible = true">新增分类</a-button>
                </a-col>
            </a-row>

            <a-table :columns="columns" :pagination='pagination' :dataSource="cateList" rowKey="id" bordered
                @change="handleTableChange">
                <template slot="action" slot-scope="data">
                    <div class="actionSlot">
                        <a-button type="primary" icon="edit" style="margin-right:15px"
                            @click="editCate(data.id)">编辑</a-button>
                        <a-button type="danger" icon="delete" style="margin-right:15px"
                            @click="deleteCate(data.id)">删除</a-button>
                    </div>
                </template>
            </a-table>
        </a-card>
        <!--   新增分类区域   -->
        <a-modal closable :visible="addCateVisible" width="60%" title="新增分类" @ok="addCateOk" @cancel="addCateCancel">
            <a-form-model :model="newcate" :rules="addCateRules" ref="addCateRef">
                <a-form-model-item label="分类名称" prop="name">
                    <a-input v-model="newcate.name"></a-input>
                </a-form-model-item>
            </a-form-model>
        </a-modal>
        <!--   编辑分类区域   -->
        <a-modal closable :visible="editCateVisible" width="60%" title="编辑分类" @ok="editCateOk" @cancel="editCateCancel">
            <a-form-model :model="cateinfo" :rules="editCateRules" ref="editCateRef">
                <a-form-model-item label="分类名称" prop="name">
                    <a-input v-model="cateinfo.name"></a-input>
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
    width: '10%',
    key: 'id',
    align: 'center'
  },
  {
    title: '分类名',
    dataIndex: 'name',
    width: '20%',
    key: 'name',
    align: 'center'
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
      cateList: [],
      columns,
      queryParam: {
        pagesize: 5,
        pagenum: 1
      },
      newcate: {
        name: ''
      },
      cateinfo: {
        id: 0,
        name: ''
      },
      addCateRules: {
        name: [
          { required: true, message: '请输入分类名', trigger: 'blur' }
        ]
      },
      editCateRules: {
        name: [
          { required: true, message: '请输入分类名', trigger: 'blur' }
        ]
      },
      addCateVisible: false,
      editCateVisible: false,
      visible: false
    }
  },
  created () {
    this.getCateList()
  },
  methods: {
    // 获取分类列表
    async getCateList () {
      const { data: res } = await this.$http.get('categories', {
        params: {
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 200) return this.$message.error(res.message)
      this.cateList = res.data
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
      this.getCateList()
    },
    // 删除分类
    deleteCate (id) {
      this.$confirm({
        title: '确定删除该用分类吗？',
        content: '一旦删除，无法恢复',
        onOk: async () => {
          const { data: res } = await this.$http.delete(`category/${id}`)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$message.success('删除成功')
          this.getCateList()
        },
        onCancel: async () => {
          this.$message.info('已取消删除')
        }
      })
    },
    // 新增分类
    addCateOk () {
      this.$refs.addCateRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.post('/category/add', {
          name: this.newcate.name
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.addCateVisible = false
        this.$message.info('添加分类成功')
        await this.getCateList()
      })
    },
    addCateCancel () {
      this.$refs.addCateRef.resetFields()
      this.addCateVisible = false
      this.$message.info('新增分类已取消')
    },
    // 编辑用户
    async editCate (id) {
      this.editCateVisible = true
      const { data: res } = await this.$http.get(`/category/${id}`)
      if (res.status !== 200) this.$message.error(res.message)
      this.cateinfo.name = res.data.name
      this.cateinfo.id = id
    },
    editCateOk () {
      this.$refs.editCateRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.put(`category/${this.cateinfo.id}`, {
          name: this.cateinfo.name
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.editCateVisible = false
        this.$message.info('更新分类信息成功')
        this.getCateList()
      })
    },
    editCateCancel () {
      this.editCateVisible = false
      this.$refs.editCateRef.resetFields()
      this.$message.info('编辑已取消')
    }
  }
}

</script>

<style scoped>
.actionSlot {
    display: flex;
    justify-content: center;
}
</style>

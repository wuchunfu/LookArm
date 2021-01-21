<template>
  <div>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="4">
          <a-button type="primary" @click="addTagVisible = true">新增状态</a-button>
        </a-col>
      </a-row>

      <a-table
        rowKey="id"
        :columns="columns"
        :pagination="pagination"
        :dataSource="Taglist"
        bordered
        @change="handleTableChange"
      >
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button
              type="primary"
              icon="edit"
              style="margin-right: 15px"
              @click="editTag(data.id)"
            >编辑</a-button>
            <a-button
              type="danger"
              icon="delete"
              style="margin-right: 15px"
              @click="deleteTag(data.id)"
            >删除</a-button>
          </div>
        </template>
      </a-table>
    </a-card>

    <!-- 新增状态区域 -->
    <a-modal
      closable
      title="新增状态"
      :visible="addTagVisible"
      width="60%"
      @ok="addTagOk"
      @cancel="addTagCancel"
      destroyOnClose
    >
      <a-form-model :model="newTag" :rules="addTagRules" ref="addTagRef">
        <a-form-model-item label="状态名称" prop="tag_name">
          <a-input v-model="newTag.tag_name"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!-- 编辑状态区域 -->
    <a-modal
      closable
      destroyOnClose
      title="编辑状态"
      :visible="editTagVisible"
      width="60%"
      @ok="editTagOk"
      @cancel="editTagCancel"
    >
      <a-form-model :model="TagInfo" :rules="TagRules" ref="addTagRef">
        <a-form-model-item label="状态名称" prop="tag_name">
          <a-input v-model="TagInfo.tag_name"></a-input>
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
    align: 'center',
  },
  {
    title: '状态状态',
    dataIndex: 'tag_name',
    width: '20%',
    key: 'tag_name',
    align: 'center',
  },
  {
    title: '操作',
    width: '30%',
    key: 'action',
    align: 'center',
    scopedSlots: { customRender: 'action' },
  },
]

export default {
  data() {
    return {
      pagination: {
        pageSizeOptions: ['5', '10', '20'],
        pageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
      },
      Taglist: [],
      TagInfo: {
        tag_name: '',
        id: 0,
      },
      newTag: {
        tag_name: '',
      },
      columns,
      queryParam: {
        pagesize: 5,
        pagenum: 1,
      },
      editVisible: false,
      TagRules: {
        tag_name: [
          {
            validator: (rule, value, callback) => {
              if (this.TagInfo.tag_name === '') {
                callback(new Error('请输入状态名'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          },
        ],
      },
      addTagRules: {
        name: [
          {
            validator: (rule, value, callback) => {
              if (this.newTag.tag_name === '') {
                callback(new Error('请输入状态名'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          },
        ],
      },
      editTagVisible: false,
      addTagVisible: false,
    }
  },
  created() {
    this.getTagList()
  },
  methods: {
    // 获取状态列表
    async getTagList() {
      const { data: res } = await this.$http.get('tag/list', {
        pagesize: this.queryParam.pagesize,
        pagenum: this.queryParam.pagenum,
      })

      if (res.status !== 200) {
        if (res.status === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear()
          this.$router.push('/login')
        }
        this.$message.error(res.message)
      }
      this.Taglist = res.data
      this.pagination.total = res.total
    },
    // 更改分页
    handleTableChange(pagination, filters, sorter) {
      var pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.queryParam.pagesize = pagination.pageSize
      this.queryParam.pagenum = pagination.current

      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1
        pager.current = 1
      }
      this.pagination = pager
      this.getTagList()
    },
    // 删除状态
    deleteTag(id) {
      this.$confirm({
        title: '提示：请再次确认',
        content: '确定要删除该状态吗？一旦删除，无法恢复',
        onOk: async () => {
          const { data: res } = await this.$http.delete(`tag/delete/${id}`)
          if (res.status != 200) return this.$message.error(res.message)
          this.$message.success('删除成功')
          this.getTagList()
        },
        onCancel: () => {
          this.$message.info('已取消删除')
        },
      })
    },
    // 新增状态
    addTagOk() {
      this.$refs.addTagRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.post('tag/add', {
          tag_name: this.newTag.tag_name,
        })
        if (res.status != 200) return this.$message.error(res.message)
        this.$refs.addTagRef.resetFields()
        this.addTagVisible = false
        this.$message.success('添加状态成功')
        await this.getTagList()
      })
    },
    addTagCancel() {
      this.$refs.addTagRef.resetFields()
      this.addTagVisible = false
      this.$message.info('新增状态已取消')
    },
    // 编辑状态
    async editTag(id) {
      this.editTagVisible = true
      const { data: res } = await this.$http.get(`tag/info/${id}`)
      this.TagInfo = res.data
    },
    editTagOk() {
      this.$refs.addTagRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.put(`tag/edit/${this.TagInfo.id}`, {
          tag_name: this.TagInfo.tag_name,
        })
        if (res.status != 200) return this.$message.error(res.message)
        this.editTagVisible = false
        this.$message.success('更新状态信息成功')
        this.getTagList()
      })
    },
    editTagCancel() {
      this.$refs.addTagRef.resetFields()
      this.editTagVisible = false
      this.$message.info('编辑已取消')
    },
  },
}
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
</style>
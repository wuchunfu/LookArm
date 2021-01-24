<template>
  <div>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search
            v-model="queryParam.app_name"
            placeholder="输入表单名查找"
            enter-button
            allowClear
            @search="getPostInfoCateList"
          />
        </a-col>

        <a-col :span="3" :offset="1">
          <a-select placeholder="请选择分类" style="width: 200px" @change="gotoCatePage">
            <a-select-option
              v-for="item in Catelist"
              :key="item.id"
              :value="item.id"
            >{{ item.name }}</a-select-option>
          </a-select>
        </a-col>
        <a-col :span="3" :offset="1">
          <a-select placeholder="请选择状态" style="width: 200px" @change="gotoTagPage">
            <a-select-option
              v-for="item in Taglist"
              :key="item.id"
              :value="item.id"
            >{{ item.tag_name }}</a-select-option>
          </a-select>
        </a-col>
        <a-col :span="1" :offset="1">
          <a-button type="info" @click="$router.push('/postinfo')">显示全部</a-button>
        </a-col>
      </a-row>

      <a-table
        rowKey="ID"
        :columns="columns"
        :pagination="pagination"
        :dataSource="PostInfoCatelist"
        bordered
        @change="handleTableChange"
      >
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button
              size="small"
              type="primary"
              icon="edit"
              style="margin-right: 15px"
              @click="editPostInfoCate(data.ID)"
            >编辑</a-button>
            <a-button
              size="small"
              type="danger"
              icon="delete"
              style="margin-right: 15px"
              @click="deletePostInfoCate(data.ID)"
            >删除</a-button>
          </div>
        </template>
      </a-table>
    </a-card>

    <!-- 编辑表单区域 -->
    <a-modal
      closable
      :visible="editPostInfoCateVisible"
      width="60%"
      @ok="editPostInfoCateOk"
      @cancel="editPostInfoCateCancel"
      destroyOnClose
    >
      <a-form-model v-model="PostInfoCate">
        <a-row :gutter="20">
          <a-col :span="6">
            <a-form-model-item label="提供人">
              <a-input v-model="PostInfoCate.user_name"></a-input>
            </a-form-model-item>
          </a-col>
          <a-col :span="6" :offset="4">
            <a-form-model-item label="提供人邮箱">
              <a-input v-model="PostInfoCate.email"></a-input>
            </a-form-model-item>
          </a-col>
        </a-row>

        <a-row :gutter="20">
          <a-col :span="6">
            <a-form-model-item label="App分类">
              <a-select v-model="PostInfoCate.category_id" @change="cateChange">
                <a-select-option v-for="item in Catelist" :key="item.id">{{item.name}}</a-select-option>
              </a-select>
            </a-form-model-item>
          </a-col>

          <a-col :span="6" :offset="4">
            <a-form-model-item label="状态">
              <a-select v-model="PostInfoCate.tag_id" @change="tagChange">
                <a-select-option v-for="item in Taglist" :key="item.id">{{item.tag_name}}</a-select-option>
              </a-select>
            </a-form-model-item>
          </a-col>
        </a-row>
        <a-form-model-item label="App名称">
          <a-input v-model="PostInfoCate.app_name"></a-input>
        </a-form-model-item>
        <a-form-model-item label="App版本">
          <a-input v-model="PostInfoCate.app_verison"></a-input>
        </a-form-model-item>
        <a-form-model-item label="App网站">
          <a-input v-model="PostInfoCate.app_website"></a-input>
        </a-form-model-item>
        <a-form-model-item label="开发者">
          <a-input v-model="PostInfoCate.app_developer"></a-input>
        </a-form-model-item>
        <a-form-model-item label="App描述">
          <a-input type="textarea" v-model="PostInfoCate.app_desc"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>
import moment from 'moment'

const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '5%',
    key: 'id',
    align: 'center',
  },
  {
    title: '提交日期',
    dataIndex: 'CreatedAt',
    width: '10%',
    key: 'CreatedAt',
    align: 'center',
    customRender: (val) => {
      return val ? moment(val).format('YYYY年MM月DD日') : '暂无'
    },
  },
  {
    title: '分类',
    dataIndex: 'Category.name',
    width: '10%',
    key: 'Category.name',
    align: 'center',
  },
  {
    title: 'App名称',
    dataIndex: 'app_name',
    width: '10%',
    key: 'app_name',
    align: 'center',
  },
  {
    title: 'App状态',
    dataIndex: 'Tag.tag_name',
    width: '15%',
    key: 'Tag.tag_name',
    align: 'center',
  },
  {
    title: '开发者',
    dataIndex: 'app_developer',
    width: '15%',
    key: 'app_developer',
    align: 'center',
  },
  {
    title: '提交者',
    dataIndex: 'user_name',
    width: '10%',
    key: 'user_name',
    align: 'center',
  },
  {
    title: '提交者邮箱',
    dataIndex: 'email',
    width: '10%',
    key: 'email',
    align: 'center',
  },
  {
    title: '操作',
    width: '15%',
    key: 'action',
    align: 'center',
    scopedSlots: { customRender: 'action' },
  },
]

export default {
  props: ['id'],
  data() {
    return {
      pagination: {
        pageSizeOptions: ['5', '10', '20'],
        pageSize: 10,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
      },
      PostInfoCatelist: [],
      Catelist: [],
      Taglist: [],
      columns,
      addPostInfoCateVisible: false,
      editPostInfoCateVisible: false,
      queryParam: {
        app_name: '',
        pagesize: 10,
        pagenum: 1,
      },
      newPostInfoCate: {
        app_name: '',
        app_version: '',
        app_webpage: '',
        app_desc: '',
        app_developer: '',
        user_name: '',
        email: '',
        category_id: 2,
        tag_id: 1,
      },
      PostInfoCate: {
        app_name: '',
        app_version: '',
        app_webpage: '',
        app_desc: '',
        app_developer: '',
        user_name: '',
        email: '',
        category_id: 2,
        tag_id: 1,
      },
    }
  },
  created() {
    this.getPostInfoCateList()
    this.getCateList()
    this.getTaglist()
  },
  methods: {
    // 获取表单列表
    async getPostInfoCateList() {
      const { data: res } = await this.$http.get(`postinfo/category/${this.id}`, {
        params: {
          app_name: this.queryParam.app_name,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum,
        },
      })
      if (res.status !== 200) {
        if (res.status === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear()
          this.$router.push('/login')
        }
        this.$message.error(res.message)
      }

      this.PostInfoCatelist = res.data
      this.pagination.total = res.total
    },
    // 获取分类
    async getCateList() {
      const { data: res } = await this.$http.get('category/list')
      if (res.status !== 200) return this.$message.error(res.message)
      this.Catelist = res.data
    },
    // 获取状态
    async getTaglist() {
      const { data: res } = await this.$http.get('tag/list')
      if (res.status !== 200) return this.$message.error(res.message)
      this.Taglist = res.data
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
      this.getPostInfoCateList()
    },

    // 新增表单
    //  this.$refs.addTagRef.validate(async (valid) => {
    //     if (!valid) return this.$message.error('参数不符合要求，请重新输入'）
    //   })
    // async addPostInfoCateOk() {
    //   const { data: res } = await this.$http.post('PostInfoCate/add', this.newPostInfoCate)
    //   if (res.status != 200) return this.$message.error(res.message)
    //   this.addPostInfoCateVisible = false
    //   this.$message.success('添加表单成功')
    //   await this.getPostInfoCateList()
    // },
    // addPostInfoCateCancel() {
    //   this.addPostInfoCateVisible = false
    //   this.$message.info('新增表单已取消')
    // },

    cateChange(value) {
      this.newPostInfoCate.category_id = value
    },

    tagChange(value) {
      this.newPostInfoCate.tag_id = value
    },

    // 删除表单
    deletePostInfoCate(id) {
      this.$confirm({
        title: '提示：请再次确认',
        content: '确定要删除该文章吗？一旦删除，无法恢复',
        onOk: async () => {
          const { data: res } = await this.$http.delete(`postinfo/delete/${id}`)
          if (res.status != 200) return this.$message.error(res.message)
          this.$message.success('删除成功')
          this.getPostInfoCateList()
        },
        onCancel: () => {
          this.$message.info('已取消删除')
        },
      })
    },

    // 编辑分类
    async editPostInfoCate(id) {
      this.editPostInfoCateVisible = true
      const { data: res } = await this.$http.get(`postinfo/info/${id}`)
      this.PostInfoCate = res.data
      this.PostInfoCate.id = id
    },
    async editPostInfoCateOk() {
      const { data: res } = await this.$http.put(`postinfo/edit/${this.PostInfoCate.id}`, this.PostInfoCate)
      if (res.status != 200) return this.$message.error(res.message)
      this.editPostInfoCateVisible = false
      this.$message.success('更新分类信息成功')
      this.getPostInfoCateList()
    },
    editPostInfoCateCancel() {
      this.editPostInfoCateVisible = false
      this.$message.info('编辑已取消')
    },
    // 查询分类下的表单
    gotoCatePage(value) {
      this.$router.push(`/postinfo/catelist/${value}`)
    },
    // 查询状态下的表单
    gotoTagPage(value) {
      this.$router.push(`/postinfo/taglist/${value}`)
    },
  },
}
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
.PostInfoCateImg {
  height: 100%;
  width: 100%;
}

.PostInfoCateImg img {
  width: 100px;
  height: 80px;
}
</style>

<template>
  <div class="mt-2" app>
    <div v-if="total == 0 && isLoad" class="d-flex justify-center align-center">
      <div>
        <v-alert class="ma-5" dense outlined type="error"
          >抱歉，你搜索的APP还未搜录，提交个表单告诉我们吧！</v-alert
        >
      </div>
    </div>

    <v-row justify="center" align="center">
      <v-col cols="6">
        <v-card
          outlined
          elevation="2"
          class="my-4"
          v-for="item in appInfoList"
          :key="item.id"
        >
          <v-row class="mt-3" align="center" justify="space-around" no-gutters>
            <v-col cols="9">
              <v-card-title class="title text-h5 font-weight-bold headline">{{
                item.app_name
              }}</v-card-title>
            </v-col>
            <v-col cols="3">
              <v-chip outlined label>{{ item.Category.name }}</v-chip>
            </v-col>
          </v-row>
          <v-card-text>
            <v-row align="center" justify="center">
              <v-col class="font-weight-bold" cols="4"
                >版本：{{ item.app_version }}</v-col
              >
              <v-col class="font-weight-bold" cols="4"
                >当前状态: {{ item.Tag.tag_name }}</v-col
              >
              <v-col class="font-weight-bold" cols="4"
                >开发商: {{ item.app_developer }}</v-col
              >
            </v-row>
          </v-card-text>
          <v-card-subtitle class="font-weight-bold"
            >简短描述：{{ item.app_desc }}</v-card-subtitle
          >

          <v-divider class="mx-3"></v-divider>
          <v-card-actions>
            <v-row align="center" justify="center">
              <v-col cols="6">
                <v-btn :href="item.app_webpage" target="_blank" outlined
                  >下载地址</v-btn
                >
              </v-col>
              <v-col cols="4">
                <v-btn text
                  >本条信息由{{ item.user_name }}({{ item.email }})提供</v-btn
                >
              </v-col>
            </v-row>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>

    <div v-if="total > 0 && !isLoad" class="mt-2 text-center">
      <v-pagination
        color="darken-2"
        dark
        total-visible="7"
        v-model="queryParam.pagenum"
        :length="Math.ceil(total / queryParam.pagesize)"
        @input="getArtList()"
      ></v-pagination>
    </div>
  </div>
</template>
<script>
export default {
  props: ['title'],
  data() {
    return {
      appInfoList: [],
      queryParam: {
        app_name: '',
        pagesize: 10,
        pagenum: 1,
      },
      total: 0,
      isLoad: false,
    }
  },
  created() {
    this.getAppInfoList()
  },
  methods: {
    // 获取app信息列表
    async getAppInfoList() {
      const { data: res } = await this.$http.get('appinfo/list', {
        params: {
          app_name: this.title,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum,
        },
      })
      this.appInfoList = res.data
      this.total = res.total
      this.isLoad = true
    },
  },
}
</script>
<style>
</style>
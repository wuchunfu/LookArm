<template>
  <div class="mt-2" app>
    <v-row justify="center" align="center">
      <v-col cols="6">
        <v-card
          outlined
          elevation="3"
          class="my-4"
          justify-center
          v-for="item in appInfoList"
          :key="item.id"
        >
          <v-row class="mt-3" align="center" justify="center" no-gutters>
            <v-col cols="8">
              <v-card-title class="title text-h5 font-weight-bold headline">
                {{ item.app_name }}
              </v-card-title>
            </v-col>
            <v-col cols="4">
              <v-chip outlined label>{{ item.Category.name }}</v-chip>
            </v-col>
          </v-row>
          <v-card-text>
            <v-row align="center" justify="start">
              <v-col class="font-weight-bold" cols="6"
                >📋版本：{{ item.app_version }}</v-col
              >
              <v-col class="font-weight-bold" cols="6"
                >💡状态: {{ item.Tag.tag_name }}</v-col
              >
              <v-col class="font-weight-bold" cols="6"
                >🖥️开发: {{ item.app_developer }}</v-col
              >
              <v-col class="font-weight-bold" cols="6"
                >🕒更新: {{ item.UpdatedAt | dateFormat }}</v-col
              >
            </v-row>
          </v-card-text>

          <v-card-subtitle class="font-weight-bold"
            >简短描述：{{ item.app_desc }}</v-card-subtitle
          >

          <v-divider class="mx-3"></v-divider>
          <v-card-actions>
            <v-row align="center" justify="space-around">
              <v-col cols="5" offset="1">
                <v-btn :href="item.app_webpage" target="_blank" outlined
                  >下载地址</v-btn
                >
              </v-col>
              <v-col cols="6">
                <v-card-text
                  >本条信息由{{ item.user_name }}({{
                    item.email
                  }})提供</v-card-text
                >
              </v-col>
            </v-row>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>

    <div class="mt-2 text-center">
      <v-pagination
        color="darken-2"
        dark
        total-visible="7"
        v-model="queryParam.pagenum"
        :length="Math.ceil(total / queryParam.pagesize)"
        @input="getAppInfoList()"
      ></v-pagination>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      appInfoList: [],
      queryParam: {
        pagesize: 10,
        pagenum: 1,
      },
      total: 0,
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
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum,
        },
      })
      this.appInfoList = res.data
      this.total = res.total
    },
  },
}
</script>
<style>
</style>
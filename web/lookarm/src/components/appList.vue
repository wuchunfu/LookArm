<template>
  <div>
    <v-sheet class="mt-5 d-flex justify-center">
      <v-row>
        <v-col cols="6" offset="3">
          <v-text-field label="输入App名查找"></v-text-field>
        </v-col>
        <v-col cols="6" offset="3">
          <v-text-field
            outlined
            clearable
            drak
            prepend-inner-icon="mdi-feature-search-outline"
            label="输入App名查找"
          ></v-text-field>
        </v-col>
      </v-row>
    </v-sheet>

    <v-container>
      <v-sheet>
        <v-list class="ma-3" outlined v-for="item in appInfoList" :key="item.id">
          <v-list-group></v-list-group>
          <v-list-item-title>App名：{{item.app_name}}</v-list-item-title>
        </v-list>
      </v-sheet>
    </v-container>
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
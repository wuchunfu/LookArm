<template>
  <div app>
    <TopBar></TopBar>
    <v-main mobile-breakpoint="sm">
      <div class="mt-16 text-center justify-center text-h3">
        <p>快速查看哪些应用支持</p>
        <p>Apple Silicon - m1</p>
      </div>

      <div class="d-flex justify-center align-center" mobile-breakpoint="sm">
        <div v-for="item in tagList" :key="item.id">
          <v-btn large text>{{item.tag_name}}</v-btn>
        </div>
      </div>

      <v-row class="mt-6" justify="center" align="center" align-content="center">
        <v-col cols="6">
          <v-text-field
            color="indigo"
            outlined
            clearable
            prepend-inner-icon="mdi-feature-search-outline"
            label="输入App名查找"
            v-model="searchName"
            @change="searchTitle(searchName)"
          ></v-text-field>
        </v-col>
      </v-row>
      <v-divider class="mx-14"></v-divider>

      <router-view :key="$route.path"></router-view>
    </v-main>

    <Footer></Footer>
  </div>
</template>

<script>
import TopBar from '../components/layuot/topBar'
import Footer from '../components/layuot/footer'
export default {
  components: { TopBar, Footer },
  data() {
    return {
      total: 0,
      tagList: [],
      searchName: '',
    }
  },
  created() {
    this.getTaglist()
  },
  methods: {
    // 获取tag信息列表
    async getTaglist() {
      const { data: res } = await this.$http.get('tag/list')
      this.tagList = res.data
    },
    // 查找App信息
    searchTitle(title) {
      this.$router.push(`/appinfo/${title}`)
    },
  },
}
</script>

<template>
  <v-main app>
    <TopBar></TopBar>
    <v-container>
      <v-row align="center" justify="center">
        <v-col cols="12" sm="4">
          <div class="mt-16 text-center text-h3">
            <p>快速查看哪些应用支持</p>
            <p>Apple Silicon - M1</p>
          </div>
        </v-col>
      </v-row>

      <v-row align="center" justify="center">
        <v-col cols="12">
          <div class="d-flex justify-center align-center">
            <div v-for="item in tagList" :key="item.id">
              <v-btn text>{{ item.tag_name }}</v-btn>
            </div>
          </div>
        </v-col>
      </v-row>

      <v-row class="mt-6" justify="center" align="center">
        <v-col cols="6">
          <v-text-field
            max-width="800"
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
    </v-container>

    <Footer></Footer>
  </v-main>
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

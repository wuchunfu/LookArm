<template>
  <v-app app>
    <TopBar></TopBar>
    <v-sheet min-height="900px">
      <v-main app>
        <v-container>
          <v-row align="center" justify="center">
            <v-col cols="12">
              <div
                :class="
                  $vuetify.breakpoint.sm
                    ? 'mt-16 text-center text-h3'
                    : 'mt-16 text-center text-h4'
                "
              >
                <p>快速查看哪些应用支持</p>
                <p>Apple Silicon - M1</p>
              </div>
            </v-col>

            <v-col cols="6" md="12">
              <div class="d-flex justify-center align-center">
                <div>
                  <v-btn v-for="item in tagList" :key="item.id" text>{{
                    item.tag_name
                  }}</v-btn>
                </div>
              </div>
            </v-col>

            <v-col cols="8" md="6" class="mt-5">
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
          <v-divider class="mx-10"></v-divider>
        </v-container>

        <router-view :key="$route.path"></router-view>
      </v-main>
    </v-sheet>
    <Footer></Footer>
  </v-app>
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

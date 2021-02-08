<template>
  <div>
    <v-app-bar mobileBreakpoint="sm" dark flat app color="grey darken-4">
      <v-app-bar-nav-icon
        class="hidden-md-and-up"
        @click.stop="drawer = !drawer"
      ></v-app-bar-nav-icon>
      <v-toolbar-title dark
        ><router-link to="/">
          <v-avatar class="mx-15 hidden-sm-and-down" size="43">
            <v-img src="../../assets/CPU.png"></v-img>
          </v-avatar>
        </router-link>
      </v-toolbar-title>

      <v-tabs center-active centered class="hidden-sm-and-down">
        <v-tab href="/"> <v-icon small>mdi-home</v-icon>首页 </v-tab>
        <v-tab
          v-for="item in CateList"
          :key="item.id"
          @click="gotoCate(item.id)"
          >{{ item.name }}</v-tab
        >
      </v-tabs>

      <v-spacer></v-spacer>

      <v-dialog max-width="800" v-model="dialog">
        <template v-slot:activator="{ on, attrs }">
          <v-btn
            class="hidden-md-and-up"
            outlined
            color="orange lighten-3"
            dark
            v-bind="attrs"
            v-on="on"
          >
            <v-icon left>mdi-post-outline</v-icon>更新数据
          </v-btn>
          <v-btn
            class="hidden-md-and-down"
            outlined
            color="orange lighten-3"
            dark
            v-bind="attrs"
            v-on="on"
          >
            <v-icon left>mdi-post-outline</v-icon>OH!帮助我们更新数据
          </v-btn>
        </template>
        <template v-slot:default="dialog">
          <v-form ref="postInfoformRef" v-model="valid">
            <v-card flat>
              <v-toolbar flat color="grey darken-3" dark
                >欢迎提交App表单，如通过，该条信息将标识由您提供</v-toolbar
              >
              <v-card-text class="mt-5">
                <v-row>
                  <v-col cols="6">
                    <v-text-field
                      v-model="postInfo.user_name"
                      hint="可以不输入，空为‘匿名’"
                      :rules="user_nameRules"
                      label="请输入您的昵称"
                    ></v-text-field>
                  </v-col>
                  <v-col cols="6">
                    <v-text-field
                      v-model="postInfo.email"
                      hint="可以不输入，空为‘匿名’"
                      :rules="emailRules"
                      label="请输入您的email"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-text-field
                  v-model="postInfo.app_name"
                  hint="不能为空"
                  :rules="app_nameRules"
                  label="请输入APP名"
                ></v-text-field>
                <v-text-field
                  v-model="postInfo.app_version"
                  :rules="app_versionRules"
                  hint="不能为空"
                  label="请输入APP版本"
                ></v-text-field>
                <v-select
                  v-model="postInfo.category_id"
                  label="请选择APP分类"
                  :items="CateList"
                  :rules="category_idRules"
                  item-text="name"
                  item-value="id"
                ></v-select>
                <v-select
                  v-model="postInfo.tag_id"
                  label="请选择APP当前状态"
                  :items="TagList"
                  :rules="tag_idRules"
                  item-text="tag_name"
                  item-value="id"
                ></v-select>
                <v-text-field
                  v-model="postInfo.app_developer"
                  :rules="app_developerRules"
                  hint="不能为空"
                  label="APP开发商/开发者"
                ></v-text-field>
                <v-text-field
                  v-model="postInfo.app_webpage"
                  :rules="app_webpageRules"
                  hint="不能为空，以http:// 或 https:// 开头"
                  label="APP官网或下载地址"
                ></v-text-field>
                <v-text-field
                  v-model="postInfo.app_desc"
                  :rules="app_descRules"
                  hint="不能为空"
                  label="APP用途的简短介绍"
                ></v-text-field>
              </v-card-text>
              <v-card-actions class="justify-end">
                <v-btn text @click="postInfoForm">确定</v-btn>
                <v-btn text @click="dialog.value = false">取消</v-btn>
              </v-card-actions>
            </v-card>
          </v-form>
        </template>
      </v-dialog>

      <v-btn
        text
        dark
        href="https://gitee.com/wejectchan/lookarm/issues"
        target="_blank"
      >
        <v-icon left>mdi-information-outline</v-icon>有建议？提交Issue
      </v-btn>
    </v-app-bar>

    <v-navigation-drawer v-model="drawer" dark app temporary>
      <v-list>
        <v-list-item-title>
          <v-btn href="/" dark text>
            <v-icon small>mdi-home</v-icon>首页
          </v-btn>
        </v-list-item-title>

        <v-list-item
          v-model="group"
          active-class="deep-purple--text text--accent-4"
          v-for="item in CateList"
          :key="item.id"
        >
          <v-list-item-title>
            <v-btn dark text @click="gotoCate(item.id)">{{ item.name }}</v-btn>
          </v-list-item-title>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </div>
</template>
<script>
export default {
  data: () => ({
    drawer: false,
    group: null,
    valid: true,
    CateList: [],
    TagList: JSON.parse(window.sessionStorage.getItem('tag')),
    postInfo: {
      app_name: '',
      app_version: '',
      app_webpage: '',
      app_desc: '',
      app_developer: '',
      user_name: 'LookArm',
      email: 'Unkown@LookArm.cn',
      category_id: 0,
      tag_id: 0
    },
    dialog: false,
    user_nameRules: [v => (v && v.length <= 20) || '昵称不能大于20个字符'],
    emailRules: [v => /.+@.+/.test(v) || 'E-mail需填入有效的形式'],
    app_nameRules: [
      v => !!v || 'APP名称不能为空',
      v => (v && v.length <= 50) || 'APP名不能大于50个字符'
    ],

    app_versionRules: [
      v => !!v || 'APP版本不能为空',
      v => (v && v.length <= 50) || 'APP版本不能大于50个字符'
    ],
    category_idRules: [v => !!v || '请选择APP分类'],
    tag_idRules: [v => !!v || '请选择APP状态'],
    app_webpageRules: [
      v => !!v || '请提供APP官网或下载地址',
      v =>
        /(http|https):\/\/([\w.]+\/?)\S*/.test(v) ||
        '请以以http:// 或 https:// 开头',
      v => (v && v.length <= 200) || '网址不能大于200个字符'
    ],
    app_descRules: [
      v => !!v || '请给APP一个简短的用途描述吧',
      v => (v && v.length <= 300) || '不能大于300个字符'
    ],
    app_developerRules: [
      v => !!v || '请提供APP的开发商或者开发者吧',
      v => (v && v.length <= 20) || '不能大于20个字符'
    ]
  }),
  created() {
    this.getCateList()
    // this.getTagList()
  },

  watch: {
    group() {
      this.drawer = false
    }
  },
  methods: {
    // 获取分类列表
    async getCateList() {
      const { data: res } = await this.$http.get('category/list')
      this.CateList = res.data
    },
    // 获取状态列表
    // async getTagList() {
    //   const { data: res } = await this.$http.get('tag/list')
    //   this.TagList = res.data
    // },
    // 提交表单
    async postInfoForm() {
      if (!this.$refs.postInfoformRef.validate())
        return this.$message.error('你提交的表单不符合要求哦~请检查')
      const { data: res } = await this.$http.post('postinfo/add', this.postInfo)
      if (res.status !== 200) return this.$message.error(res.message)
      this.$message.success('提交表单成功')
      this.dialog = false
    },
    // 前往分类
    gotoCate(id) {
      this.$router.push(`/appinfo/category/${id}`).catch(err => err)
    }
  }
}
</script>
<style lang=""></style>

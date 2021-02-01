<template>
  <v-footer padless max-height="150" color="grey darken-4" class="text-center">
    <v-card width="100%" color="grey darken-4" class="text-center">
      <v-card-text dark width="100%">
        <v-btn text dark href="/">
          <v-icon size="24px">mdi-home</v-icon>
        </v-btn>
        <v-btn text dark href="weject@gmail.com" target="blank">
          <v-icon size="24px">mdi-email</v-icon>
        </v-btn>
        <v-btn text dark href="https://github.com/wejectchen" target="blank">
          <v-icon size="24px">mdi-github</v-icon>
        </v-btn>
        <v-btn
          text
          dark
          href="https://space.bilibili.com/402177130"
          target="blank"
        >
          <v-icon size="24px">mdi-blogger</v-icon>
        </v-btn>
      </v-card-text>
      <v-divider color="grey"></v-divider>
      <v-card-text>
        <p class="text-center white--text">Made With ❤️ By ©LookArm 🎉</p>
        <p class="mb-4 text-center white--text">
          <a
            class="beian"
            style=" text-decoration: none"
            href="https://beian.miit.gov.cn"
            target="_blank"
            >蜀ICP备19035448号-2</a
          >
        </p>
      </v-card-text>
    </v-card>
  </v-footer>
</template>
<script>
export default {}
</script>
<style scoped>
.v-card {
  width: 100%;
  padding: 0;
  margin: 0;
}

a:link {
  color: white;
} /* 未访问的链接 */
a:visited {
  color: white;
} /* 已访问的链接 */
a:hover {
  color: rgb(216, 206, 69);
} /* 鼠标移动到链接上 */
a:active {
  color: yellow;
} /* 选定的链接，即鼠标按下去的时候不松开显示的状态 */
</style>

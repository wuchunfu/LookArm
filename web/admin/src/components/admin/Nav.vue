<template >
  <a-layout-sider breakpoint="lg" v-model="collapsed">
    <div class="log">
      <span>{{ collapsed ? 'Arm' : 'Look Arm' }}</span>
    </div>

    <a-menu theme="dark" mode="inline" @click="goToPage">
      <a-menu-item key="index">
        <a-icon type="dashboard"></a-icon>
        <span>仪表盘</span>
      </a-menu-item>
      <a-menu-item key="appinfo">
        <a-icon type="appstore" />
        <span>App信息</span>
      </a-menu-item>
      <a-menu-item key="postinfo">
        <a-icon type="form" />
        <span>表单管理</span>
      </a-menu-item>
      <a-menu-item key="tag">
        <a-icon type="tag" />
        <span>App状态管理</span>
      </a-menu-item>
      <a-menu-item key="catelist">
        <a-icon type="book" />
        <span>分类列表</span>
      </a-menu-item>
      <a-menu-item key="userlist">
        <a-icon type="user" />
        <span>用户列表</span>
      </a-menu-item>
    </a-menu>
  </a-layout-sider>
</template>

<script>
export default {
  data() {
    return {
      collapsed: false,
    }
  },
  methods: {
    goToPage(item) {
      this.$router.push('/' + item.key).catch((err) => err)
    },
  },
}
</script>

<style scoped>
.log {
  height: 32px;
  margin: 16px;
  background-color: #fff;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 17px;
}
</style>

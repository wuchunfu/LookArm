import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

//获取原型对象上的push函数
const originalPush = VueRouter.prototype.push
//修改原型对象中的push方法
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

const routes = [


  {
    path: '/',
    name: 'appInfo',
    component: () => import('../components/appList.vue')
  },
  {
    path: '/appinfo/:title',
    name: 'search',
    meta: {
      title: '搜索结果'
    },
    component: () => import('../components/search.vue'),
    props: true
  },
  {
    path: '/appinfo/category/:id',
    name: 'category',
    meta: {
      title: 'APP分类'
    },
    component: () => import('../components/appCate.vue'),
    props: true
  }
]

// {
//   path: '/about',
//   name: 'About',
//   // route level code-splitting
//   // this generates a separate chunk (about.[hash].js) for this route
//   // which is lazy-loaded when the route is visited.
//   component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
// }


const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router

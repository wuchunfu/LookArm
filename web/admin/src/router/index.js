import Vue from 'vue'
import VueRouter from 'vue-router'
const Login = () => import(/* webpackChunkName: "Login" */ '../views/Login.vue')
const Admin = () => import(/* webpackChunkName: "Admin" */ '../views/Admin.vue')

// 页面路由组件
const Index = () => import(/* webpackChunkName: "Index" */ '../components/admin/Index.vue')
const PostInfo = () => import(/* webpackChunkName: "AppList" */ '../components/postinfo/postInfo.vue')
const PostInfoCate = () => import(/* webpackChunkName: "AppList" */ '../components/postinfo/postInfoCate.vue')
const PostInfoTag = () => import(/* webpackChunkName: "AppList" */ '../components/postinfo/postInfoTag.vue')
const AppInfo = () => import(/* webpackChunkName: "AppInfoList" */ '../components/appinfo/appInfo.vue')
const CateList = () => import(/* webpackChunkName: "CateList" */ '../components/category/CateList.vue')
const UserList = () => import(/* webpackChunkName: "UserList" */ '../components/user/UserList.vue')
const Tag = () => import('../components/tag/tagList.vue')

// 路由重复点击捕获错误
const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location, onResolve, onReject) {
  if (onResolve || onReject) return originalPush.call(this, location, onResolve, onReject)
  return originalPush.call(this, location).catch(err => err)
}

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'login',
    meta: {
      title: '请登录'
    },
    component: Login
  },
  {
    path: '/',
    name: 'admin',
    meta: {
      title: 'GinBlog 后台管理页面'
    },
    component: Admin,
    children: [
      {
        path: 'index',
        component: Index,
        meta: {
          title: 'GinBlog 后台管理页面'
        }
      },
      {
        path: 'appinfo',
        component: AppInfo,
        meta: {
          title: 'App信息列表'
        }
      },
      {
        path: 'postinfo',
        component: PostInfo,
        meta: {
          title: '表单列表'
        }
      },
      {
        path: 'postinfo/catelist/:id',
        component: PostInfoCate,
        meta: {
          title: '分类表单'
        },
        props: true
      },
      {
        path: 'postinfo/taglist/:id',
        component: PostInfoTag,
        meta: {
          title: '分类表单'
        },
        props: true
      },
      {
        path: 'catelist',
        component: CateList,
        meta: {
          title: '分类列表'
        }
      },
      {
        path: 'tag',
        component: Tag,
        meta: {
          title: 'app状态管理'
        }
      },
      {
        path: 'userlist',
        component: UserList,
        meta: {
          title: '用户列表'
        }
      }
    ]
  }
]

const router = new VueRouter({
  routes
})

router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = to.meta.title
  }
  next()

  const userToken = window.sessionStorage.getItem('token')
  if (to.path === '/login') return next()
  if (!userToken) {
    next('/login')
  } else {
    next()
  }
})

export default router

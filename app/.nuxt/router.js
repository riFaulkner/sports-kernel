import Vue from 'vue'
import Router from 'vue-router'
import { normalizeURL, decode } from 'ufo'
import { interopDefault } from './utils'
import scrollBehavior from './router.scrollBehavior.js'

const _7523c12e = () => interopDefault(import('../pages/draft-tools.vue' /* webpackChunkName: "pages/draft-tools" */))
const _39a10d88 = () => interopDefault(import('../pages/league-home.vue' /* webpackChunkName: "pages/league-home" */))
const _10154228 = () => interopDefault(import('../pages/login.vue' /* webpackChunkName: "pages/login" */))
const _487b5284 = () => interopDefault(import('../pages/strategy.vue' /* webpackChunkName: "pages/strategy" */))
const _466d9f11 = () => interopDefault(import('../pages/index.vue' /* webpackChunkName: "pages/index" */))

const emptyFn = () => {}

Vue.use(Router)

export const routerOptions = {
  mode: 'history',
  base: '/',
  linkActiveClass: 'nuxt-link-active',
  linkExactActiveClass: 'nuxt-link-exact-active',
  scrollBehavior,

  routes: [{
    path: "/draft-tools",
    component: _7523c12e,
    name: "draft-tools"
  }, {
    path: "/league-home",
    component: _39a10d88,
    name: "league-home"
  }, {
    path: "/login",
    component: _10154228,
    name: "login"
  }, {
    path: "/strategy",
    component: _487b5284,
    name: "strategy"
  }, {
    path: "/",
    component: _466d9f11,
    name: "index"
  }],

  fallback: false
}

export function createRouter (ssrContext, config) {
  const base = (config._app && config._app.basePath) || routerOptions.base
  const router = new Router({ ...routerOptions, base  })

  // TODO: remove in Nuxt 3
  const originalPush = router.push
  router.push = function push (location, onComplete = emptyFn, onAbort) {
    return originalPush.call(this, location, onComplete, onAbort)
  }

  const resolve = router.resolve.bind(router)
  router.resolve = (to, current, append) => {
    if (typeof to === 'string') {
      to = normalizeURL(to)
    }
    return resolve(to, current, append)
  }

  return router
}

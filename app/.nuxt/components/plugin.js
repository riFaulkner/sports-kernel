import Vue from 'vue'
import { wrapFunctional } from './utils'

const components = {
  ContractCard: () => import('../../components/ContractCard.vue' /* webpackChunkName: "components/contract-card" */).then(c => wrapFunctional(c.default || c)),
  Logo: () => import('../../components/Logo.vue' /* webpackChunkName: "components/logo" */).then(c => wrapFunctional(c.default || c)),
  PlayerCard: () => import('../../components/PlayerCard.vue' /* webpackChunkName: "components/player-card" */).then(c => wrapFunctional(c.default || c))
}

for (const name in components) {
  Vue.component(name, components[name])
  Vue.component('Lazy' + name, components[name])
}

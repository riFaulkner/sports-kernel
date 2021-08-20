import { wrapFunctional } from './utils'

export { default as ContractCard } from '../../components/ContractCard.vue'
export { default as Logo } from '../../components/Logo.vue'
export { default as PlayerCard } from '../../components/PlayerCard.vue'

export const LazyContractCard = import('../../components/ContractCard.vue' /* webpackChunkName: "components/contract-card" */).then(c => wrapFunctional(c.default || c))
export const LazyLogo = import('../../components/Logo.vue' /* webpackChunkName: "components/logo" */).then(c => wrapFunctional(c.default || c))
export const LazyPlayerCard = import('../../components/PlayerCard.vue' /* webpackChunkName: "components/player-card" */).then(c => wrapFunctional(c.default || c))

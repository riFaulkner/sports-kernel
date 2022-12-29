import { USER_PREFERENCES_QUERY } from '@/graphql/queries/user/userGraphQl'

export const state = () => ({
  userId: null,
  userPreferences: {
    leagues: null,
    preferredLeagueId: null,
    isAdmin: null
  }
})

export const getters = {
  getUserId (state) {
    return state.userId
  },
  getUserLeagues (state) {
    return state.userPreferences.leagues
  },
  getIsUserAdmin (state) {
    return state.userPreferences.isAdmin
  },
  getPreferredLeagueId (state) {
    return state.userPreferences.preferredLeagueId
  }
}

export const mutations = {
  setUserPreferences (state, payload) {
    state.userId = payload.id
    state.userPreferences.leagues = payload.leagues
    state.userPreferences.preferredLeagueId = payload.preferredLeagueId
    state.userPreferences.isAdmin = payload.isAdmin
  }
}
export const actions = {
  async initializeUserPreferences ({ context, state, commit }, payload) {
    const userId = payload.userId
    if (userId) {
      const result = await payload.apolloClient.query(
        {
          query: USER_PREFERENCES_QUERY,
          variables: {
            userId
          }
        }
      ).then((response) => {
        return response.data.userPreferences
      }).catch((error) => {
        console.error("error loading user preferences:", error)
        throw error
      })

      await commit('setUserPreferences', result)
    }
  },
  setUserPreferences(context, payload) {
    context.commit("setUserPreferences", payload)
  }
}

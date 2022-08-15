export const state = () => ({
  isInitialized: false,
  activeLeague: null,
  activeTab: null,
  alert: {
    alert: false,
    alertPayload: {
      color: 'primary',
      timeout: -1,
      message: ''
    }
  }
})

export const mutations = {
  saveActiveLeague (state, payload) {
    state.activeLeague = payload
  },
  updateActiveTab (state, payload) {
    state.activeTab = payload
  },
  UPDATE_ALERT (state, payload) {
    state.alert = payload
  }
}

export const actions = {
  updateActiveLeague (context, payload) {
    context.commit('saveActiveLeague', payload)
    // localStorage.userPreferences = context.state;
  },
  updateActiveTab (context, payload) {
    context.commit('updateActiveTab', payload)
  },
  alertDismiss (context) {
    const dismissAlert = {
      alert: false,
      alertPayload: {
        color: 'primary',
        timeout: -1,
        message: ''
      }
    }
    context.commit('UPDATE_ALERT', dismissAlert)
  },
  alertSuccess (context, payload) {
    const successAlert = {
      alert: true,
      alertPayload: {
        color: 'success',
        timeout: -1,
        message: payload.message
      }
    }
    context.commit('UPDATE_ALERT', successAlert)
  },
  alertInfo (context, payload) {
    const infoAlert = {
      alert: true,
      alertPayload: {
        color: 'info',
        timeout: -1,
        message: payload.message
      }
    }
    context.commit('UPDATE_ALERT', infoAlert)
  },
  alertError (context, payload) {
    const errorAlert = {
      alert: true,
      alertPayload: {
        color: 'error',
        timeout: -1,
        message: payload.message
      }
    }
    context.commit('UPDATE_ALERT', errorAlert)
  }
}

export const getters = {
  getActiveLeague (state) {
    // TODO: This is routed correctly to be able to support localStorage,
    //  but there are some issues w/SSR and localStorage See link: https://forum.vuejs.org/t/how-to-use-localstorage-in-vuex/2029/5
    return state.activeLeague
  },
  getActiveTab (state) {
    return state.activeTab
  }
}

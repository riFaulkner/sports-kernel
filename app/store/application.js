export const state = () => ({
    isInitialized: false,
    activeLeague: null,
    submenu: null,
    activeTab: null
});

export const mutations = {
    saveActiveLeague(state, payload) {
        state.activeLeague = payload;
    },
    updateActiveTab(state, payload) {
        state.activeTab = payload;
    },
    updateSubmenu(state, payload) {
        state.submenu = payload;
    }
}

export const actions = {
    updateActiveLeague(context, payload) {
        context.commit("saveActiveLeague", payload);
        // localStorage.userPreferences = context.state;
    },
    updateActiveTab(context, payload) {
        context.commit("updateActiveTab", payload);
    },
    updateSubmenu(context, payload) {
        context.commit("updateSubmenu", payload);
    }
}

export const getters = {
    getActiveLeague(state) {
        // TODO: This is routed correctly to be able to support localStorage,
        //  but there are some issues w/SSR and localStorage See link: https://forum.vuejs.org/t/how-to-use-localstorage-in-vuex/2029/5
        return state.activeLeague;
    },
    getActiveTab(state) {
        return state.activeTab;
    }
}
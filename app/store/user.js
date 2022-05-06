import {USER_PREFERENCES_QUERY} from "@/graphql/queries/user/userGraphQl";

export const state = () => ({
    userId: null,
    userPreferences: {
        leagues: null,
        preferredLeagueId: "",
        isAdmin: null
    },
});

export const getters = {
    getUserLeagues(state) {
        return state.userPreferences.leagues;
    },
    getIsUserAdmin(state) {
        return state.userPreferences.isAdmin;
    },
}

export const mutations = {
    async setUserPreferences(state, payload) {
        state.userId = payload.id;
        state.userPreferences.leagues = payload.leagues;
        state.userPreferences.preferredLeagueId = payload.preferredLeagueId;
        state.userPreferences.isAdmin = payload.isAdmin;
    }
}
export const actions = {
    async initializeUserPreferences({ state, commit }, payload) {
        const userId = payload.userId;
        if (userId) {
            const result = await payload.apolloClient.query(
                {
                    query: USER_PREFERENCES_QUERY,
                    variables: {
                        userId: userId
                    }
                }
            ).then((response) => {
                return response.data.userPreferences
            });
            await commit('setUserPreferences', result);
        }
    }
}
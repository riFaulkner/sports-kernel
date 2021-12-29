import {USER_PREFERENCES_QUERY} from "@/graphql/queries/user/userGraphQl";

export const state = () => ({
    userId: 'x3IOCU28HGe2GHAF4azJ',
    userPreferences: {
        leagues: null,
        preferredLeagueId: ""
    },
});

export const getters = {
    getUserId(state) {
        const userString = this.$auth.user.userId
        // TODO: remove console log
        console.log("userId: ", userString);
        return userString;
    },
    getUserLeagues(state) {
        return state.userPreferences.leagues;
    },
    async getDefaultLeague(state) {
        console.info("Getting default league");
        const leagues = state.userPreferences.leagues;
        const preferredLeagueId = state.userPreferences.preferredLeagueId;
        if (leagues && preferredLeagueId) {
            console.info("Filtering for default league");
            return leagues.find(league => league.id === preferredLeagueId);
        }
        return null;
    },
}

export const mutations = {
    async setUserPreferences(state, payload) {
        state.userId = payload.id;
        state.userPreferences.leagues = payload.leagues;
        state.userPreferences.preferredLeagueId = payload.preferredLeagueId;
    }
}
export const actions = {
    async initializeUserPreferences({ state, commit }, payload) {
        const userId = state.userId;
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
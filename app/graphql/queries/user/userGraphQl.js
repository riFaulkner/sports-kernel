import gql from "graphql-tag";


export const USER_PREFERENCES_QUERY = gql`
query UserPreferences($userId: ID!) {
    UserPreferences(userId: $userId) {
        id
        ownerName
        preferredLeagueId
        leagues {
            id
            leagueName
        }
    }
}
`
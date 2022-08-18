import gql from 'graphql-tag'

export const USER_PREFERENCES_QUERY = gql`
query UserPreferences($userId: ID!) {
    userPreferences(userId: $userId) {
        id
        ownerName
        preferredLeagueId
        isAdmin
        leagues {
            id
            leagueName
        }
    }
}
`
export const ADD_USER = gql`
mutation addUserToTeam($accessCode: string!){
    addUserToTeam(accessCode: $accessCode){
        id
        ownerName
        preferredLeagueId
        isAdmin
        leagues{
            id
            leagueName
        }
    }
}
`

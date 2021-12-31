import gql from "graphql-tag"

export const LEAGUE_BY_ID_QUERY = gql`
    query League($leagueId: ID!) {
        league(leagueId: $leagueId) {
            id
            leagueName
            divisions {
                divisionName
                leadingWins
            }
        }
    }
`
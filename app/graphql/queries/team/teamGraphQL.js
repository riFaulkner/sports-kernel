import gql from "graphql-tag"

export const TEAM_CONTRACTS = gql`
    query teamContracts($leagueId: ID!, $teamId: ID!) {
        teamContracts(leagueId:$leagueId, teamId: $teamId) {
            id
            playerId
            player {
                id
                playerName
            }
            teamId
            currentYear
            restructureStatus
            totalContractValue
            contractLength
            playerPosition
            contractDetails{
                year
                totalAmount
                paidAmount
                guaranteedAmount
            }
        }
    }
`

export const TEAM_DRAFT_PICKS = gql`
    query getTeamById($leagueId:ID!, $teamId: ID!) {
        teamById(leagueId: $leagueId, teamId:$teamId) {
            teamAssets{
                draftPicks{
                    year
                    picks{
                        round
                        value
                    }
                }
            }
        }
    }
`

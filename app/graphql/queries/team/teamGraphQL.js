import gql from 'graphql-tag'

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
export const TEAMS_WITH_SCORING = gql`
    query GetTeams($leagueId: ID!) {
        teams(leagueId: $leagueId) {
            teamName
            id
            division
            teamScoring {
                year
                summary {
                    wins
                    losses
                    ties
                    currentStreak
                    totalPointsFor
                    totalPointsAgainst
                }
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
                    picks {
                        round
                        value
                    }
                }
            }
        }
    }
`
export const TEAM_DEAD_CAP = gql`
    query getTeamById($leagueId:ID!, $teamId: ID!) {
        teamById(leagueId: $leagueId, teamId:$teamId) {
            id
            teamLiabilities {
                deadCap {
                    year
                    deadCapAccrued {
                        associatedContractId
                        amount
                        contract {
                            player {
                                playerName
                            }
                        }
                    }
                }
            }
        }
    }
`
export const GENERATE_ACCESS_CODE = gql`
    mutation generateAccessCode($leagueId:ID!, $teamId:ID!, $role:Role!){
        generateAccessCode(leagueId: $leagueId, teamId: $teamId, role: $role)
    }
`

export const ONBOARD_USER = gql`
mutation onboardUserToTeamWithAccessCode($accessCode: String!){
    onboardUserToTeamWithAccessCode(accessCode: $accessCode){
        id
        ownerName
        preferredLeagueId
        isAdmin
        leagues{
            id
            leagueName
            roleInLeague
        }
    }
}
`

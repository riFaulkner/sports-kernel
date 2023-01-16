import gql from 'graphql-tag'

export const LEAGUE_BY_ID_QUERY = gql`
    query League($leagueId: ID!) {
        league(leagueId: $leagueId) {
            id
            leagueName
            currentSeason
            divisions {
                divisionName
                leadingWins
            }
        }
    }
`

export const LEAGUE_TEAMS_BY_LEAGUE_ID = gql`
    query Teams($leagueId: ID!) {
        teams(leagueId: $leagueId) {
        id
        teamName
        currentContractsMetadata {
          totalUtilizedCap
          totalAvailableCap
          deadCapUtilizedCap {
              capUtilization
              numContracts
          }
          qbUtilizedCap {
            capUtilization
            numContracts
          }
          rbUtilizedCap {
            capUtilization
            numContracts
          }
          wrUtilizedCap {
            capUtilization
            numContracts
          }
          teUtilizedCap {
            capUtilization
            numContracts
          }
        }
      } 
}
`

export const LEAGUE_TEAMS_BY_LEAGUE_ID_BASE = gql`
    query Teams($leagueId: ID!) {
        teams(leagueId: $leagueId) {
            id
            teamName
        }
    }
`

export const LEAGUE_CONTRACTS = gql`
    query getLeagueContracts($leagueId:ID!){
        leagueContracts(leagueId:$leagueId) {
            id
            playerId
            teamId
            player {
                id
                playerName
                team
                position
            }
            contractLength
            totalContractValue
            restructureStatus
            currentYear
            contractDetails {
                year
                totalAmount
                guaranteedAmount
                paidAmount
            }
        }
    }
`

export const LEAGUE_FILTER_TEAMS_BY_OWNER_ID = gql`
    query GetFullLeagueInfo($leagueId:ID!, $filter:LeagueTeamFiltering) {
        league(leagueId: $leagueId) {
            id
            currentSeason
            teams(search: $filter) {
                id
                teamName
                activeContracts {
                    id
                    playerId
                    player {
                        id
                        playerName
                        team
                        position
                    }
                    currentYear
                    totalContractValue
                    totalRemainingValue
                    teamId
                    contractDetails {
                        guaranteedAmount
                        paidAmount
                        totalAmount
                        year
                    }
                    restructureStatus
                },
                teamAssets {
                    draftPicks {
                        year
                        picks {
                            round
                            value
                        }
                    }
                }
                teamLiabilities {
                    deadCap {
                        deadCapNote
                        deadCapYears {
                            year
                            amount
                        }
                    }
                }
            }
        }
    }
    `
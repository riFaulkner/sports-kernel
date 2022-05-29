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

export const LEAGUE_TEAMS_BY_LEAGUE_ID = gql`
    query Teams($leagueId: ID!) {
        teams(leagueId: $leagueId) {
        id
        ownerID
        teamName
        currentContractsMetadata {
          totalUtilizedCap
          totalAvailableCap
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

export const LEAGUE_CONTRACTS = gql`
    query getLeagueContracts($leagueId:ID!){
        leagueContracts(leagueId:$leagueId) {
            id
            playerId
            player {
                playerName
                teamNFL
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
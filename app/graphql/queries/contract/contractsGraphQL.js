import gql from "graphql-tag"

export const CONTRACT_RESTRUCTURE = gql`
    mutation restructureContract(
        $leagueId: ID!, 
        $restructureDetails: ContractRestructureInput!) {
        contractActionRestructure(
            leagueId: $leagueId
            restructureDetails: $restructureDetails
        ) {
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
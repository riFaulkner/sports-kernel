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
            contractDetails {
                year
                totalAmount
                paidAmount
                guaranteedAmount
            }
            totalContractValue
        }
    }
    `
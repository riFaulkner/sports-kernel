import gql from "graphql-tag"

export const CONTRACT_DROP = gql`
    mutation dropContract($leagueId:ID!, $contractId: ID!, $teamId:ID!) {
        contractActionDrop(leagueId: $leagueId, teamId: $teamId, contractId: $contractId)
    }
`

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

export const CREATE_CONTRACT = gql`
    mutation createContract($contract:ContractInput!, $leagueId:ID!) {
        createContract(leagueId: $leagueId, input: $contract) {
            id
        }
    }
`
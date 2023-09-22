import gql from "graphql-tag";

export const GET_STANDINGS = gql `
    query getStandings($leagueId: ID!, $season: Int!, $week: Int!) {
        standings(leagueId: $leagueId){
            weekStandings(season: $season, week: $week){
                teamName
                divisionId
                divisionName
                teamWins
                teamLosses
                teamTies
                pointsFor
                pointsAgainst
            }
            
        }
    }
`
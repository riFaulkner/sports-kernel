import gql from 'graphql-tag'

export const WEEK_SCORING_MATCH_UPS = gql`
    query weekScoringMatchUps($leagueId: ID!, $season: Int!, $week: Int!) {
        scoring(leagueId: $leagueId){
             weekMatchUps(season: $season, week: $week) {
                awayTeam
                homeTeam
                matchUpNumber
             }
          }
    }
`

export const WEEK_SCORING_FOR_MATCH_UP = gql`
    fragment LineUpPlayerData on PlayerScoring {
        playerName
        points
        gamePlayed
    }
    query scoring($leagueId: ID!, $season: Int!, $week: Int, $matchUpNumber: Int!) {
        scoring(leagueId: $leagueId){
            matchUpScoring(season: $season, week: $week, matchUpNumber: $matchUpNumber) {
                teamName
                isHomeTeam
                totalPoints
                lineUp {
                    qb {
                        ...LineUpPlayerData
                    }
                    rb {
                        ...LineUpPlayerData
                    }
                    wr {
                        ...LineUpPlayerData
                    }
                    te {
                        ...LineUpPlayerData
                    }
                    flex {
                        ...LineUpPlayerData
                        position
                    }
                    superFlex {
                        ...LineUpPlayerData
                        position
                    }
                }
                roster {
                    playerName
                    projectedPoints
                    points
                    nflTeam
                    position
                    isInLineUp
                    injuryStatus
                    gamePlayed
                }
            }
        }
    }
`
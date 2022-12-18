import gql from 'graphql-tag'

export const WEEK_SCORING_MATCH_UPS = gql`
    query weekScoringMatchUps($leagueId: ID!) {
        scoring(leagueId: $leagueId){
             weekMatchUps {
                awayTeam
                homeTeam
                matchUpNumber
             }
          }
    }
`

export const WEEK_SCORING_FOR_MATCH_UP = gql`
    query scoring($leagueId: ID!, $matchUpNumber: Int!) {
        scoring(leagueId: $leagueId){
            matchUpScoring(matchUpNumber: $matchUpNumber) {
                teamName
                isHomeTeam
                totalPoints
                lineUp {
                    qb {
                        playerName
                        points
                        gamePlayed
                    }
                    rb {
                        playerName
                        points
                        gamePlayed
                    }
                    wr {
                        playerName
                        points
                        gamePlayed
                    }
                    te {
                        playerName
                        points
                        gamePlayed
                    }
                    flex {
                        playerName
                        points
                        position
                        gamePlayed
                    }
                    superFlex {
                        playerName
                        points
                        position
                        gamePlayed
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
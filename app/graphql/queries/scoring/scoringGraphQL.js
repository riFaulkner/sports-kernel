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
                    }
                    rb {
                        playerName
                        points
                    }
                    wr {
                        playerName
                        points
                    }
                    te {
                        playerName
                        points
                    }
                    flex {
                        playerName
                        points
                        position
                    }
                    superFlex {
                        playerName
                        points
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
                }
            }
        }
    }
`
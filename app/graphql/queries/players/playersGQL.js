import gql from 'graphql-tag'

export const PLAYER_BY_POSITION = gql`
    query playersByPosition($position: PlayerPosition!) {
        playersByPosition (position:$position) {
            avatar
            birthday
            age
            id
            overallRank
            playerName
            position
            positionRank
            team
        }
    } 
`

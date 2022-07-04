import gql from "graphql-tag";

export const CREATE_PLAYER = gql`
    mutation createPlayer($player: NewPlayerNFL!) {
        createPlayer(input: $player) {
            id
            playerName
        }
    }
`
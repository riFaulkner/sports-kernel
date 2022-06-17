import gql from "graphql-tag"

export const GET_POSTS_BY_LEAGUE = gql`
    query getPosts($leagueId: ID!, $numOfResults: Int) {
            posts(leagueId:$leagueId, numOfResults:$numOfResults){
                id
                author
                title
                postDate
                content
                comments{
                    author
                    content
                }
            }
        } 
    `
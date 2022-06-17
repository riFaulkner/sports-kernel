<template>
  <v-card
    class="mx-auto"
    max-width="344"
  >
    <v-img
      src="https://cdn.vuetifyjs.com/images/cards/sunshine.jpg"
      height="200px"
    ></v-img>

    <v-card-title>
      {{posts[0].title}}
    </v-card-title>

    <v-card-subtitle>
      Author: {{posts[0].author}}
      <p>{{posts[0].postDate}}</p>
    </v-card-subtitle>

    <v-card-actions>
      <v-btn
        color="blue lighten-2"
        text
      >
        Comment
      </v-btn>

      <v-spacer></v-spacer>

      <v-btn
        icon
        @click="show = !show"
      >
        <v-icon>{{ show ? 'mdi-chevron-up' : 'mdi-chevron-down' }}</v-icon>
      </v-btn>
    </v-card-actions>

    <v-expand-transition>
      <div v-show="show">
        <v-divider></v-divider>

        <v-card-text>
          {{posts[0].content}}
        </v-card-text>
      </div>
    </v-expand-transition>
  </v-card>
</template>

<script>
import {GET_POSTS_BY_LEAGUE} from "@/graphql/queries/discussions/discussionsGraphQL"

  export default {
    name: "Discussions.vue",
    props: {
        leagueInfo: {
            type: Object,
            require: true,
            validator: function (value) {
                // Object.values(value).includes("divisions")
                return true
            }
        }
    },
    data: function() {
      return{
          posts: [],
          show: false,
      }
    },
    apollo: {
      posts: {
          query: GET_POSTS_BY_LEAGUE,
          variables() {
              return {leagueId: this.leagueInfo.id, numOfResults: 10}
          }
      }
    }
  }
  
</script>

<style scoped>

</style>
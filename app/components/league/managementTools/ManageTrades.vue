<template>
  <div>
    <v-row>
      <v-col>
        <v-select
            label="Team 1"
            :items="teams"
            item-text="teamName"
            item-value="id"
            :loading="$apollo.queries.teams.loading"
            v-model="teamOne"
        />
      </v-col>
      <v-col>
        <v-select
            label="Team 2"
            :items="teams"
            item-text="teamName"
            item-value="id"
            :loading="$apollo.queries.teams.loading"
            v-model="teamTwo"
        />
      </v-col>
    </v-row>

    <trades :initiating-team-id="teamOne" :secondary-team-id="teamTwo"/>

  </div>
</template>

<script>
import Trades from "@/components/league/trades/Trades";
import {LEAGUE_TEAMS_BY_LEAGUE_ID_BASE} from "@/graphql/queries/league/leagueGraphQL";
export default {
  name: "ManageTrades",
  components: {Trades},
  props: {
    leagueId: {
      type: String,
      required: true
    },
  },
  data: function() {
    return {
      teamOne: null,
      teamTwo: null,
      teams: null
    }
  },
  apollo: {
    teams: {
      query: LEAGUE_TEAMS_BY_LEAGUE_ID_BASE,
      variables() {
        return {
          leagueId: this.leagueId
        }
      }
    }
  }
}
</script>

<style scoped>

</style>
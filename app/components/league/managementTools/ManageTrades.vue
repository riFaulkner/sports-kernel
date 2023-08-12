<template>
  <div>
    <v-row>
      <v-col cols="12" md="6">
        <v-select
            label="Team 1"
            :items="mutuallyExclusiveTeams(teamTwo)"
            item-text="teamName"
            item-value="id"
            :loading="$apollo.queries.teams.loading"
            v-model="teamOne"
        />
      </v-col>
      <v-col cols="12" md="6">
        <v-select
            label="Team 2"
            :items="mutuallyExclusiveTeams(teamOne)"
            item-text="teamName"
            item-value="id"
            :loading="$apollo.queries.teams.loading"
            v-model="teamTwo"
        />
      </v-col>
    </v-row>

    <trades
        v-if="teamOne && teamTwo"
        :league-id="leagueId"
        :initiating-team-id="teamOne"
        :secondary-team-id="teamTwo"
    />

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
  },
  methods: {
    mutuallyExclusiveTeams(teamId) {
        return this.teams ? this.teams.filter(team => {
          return team.id !== teamId
        }) : null
    }
  }
}
</script>

<style scoped>

</style>
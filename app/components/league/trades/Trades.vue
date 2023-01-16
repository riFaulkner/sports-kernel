<template>
  <div>
    <h1>Trades</h1>
    {{initiatingTeamId}}
    {{secondaryTeamId}}
  </div>
</template>

<script>
import {ALL_TEAM_ASSETS_AND_LIABILITIES} from "@/graphql/queries/team/teamGraphQL";

export default {
  name: "Trades",
  props: {
    leagueId: {
      type: String,
      required: true
    },
    initiatingTeamId: {
      type: String,
    },
    secondaryTeamId: {
      type: String
    }
  },
  data: function() {
    return {
    teamAssets: null
  }},
  apollo: {
    teamAssets: {
      query: ALL_TEAM_ASSETS_AND_LIABILITIES,
      variables: function () {
        return {
          leagueId: this.leagueId,
          teamIds: [this.initiatingTeamId, this.secondaryTeamId]
        }
      },
      skip() {
        return this.initiatingTeamId === null || this.secondaryTeamId === null
      },
      update: data => data.team.teamIds
    }
  }
}
</script>

<style scoped>

</style>
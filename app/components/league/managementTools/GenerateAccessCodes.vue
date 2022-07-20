<template>
  <v-col align-self="center" md=4 offset-md=4>
    <v-card class="mb-12" color="grey darken-3">
    <h1 class="text-center">Generate Access Codes</h1>

    <v-card color="grey darken-3" flat="true">
      <v-col>
        <v-select
                v-model=selectedTeam
                :items="teams"
                item-text="teamName"
                item-value="id"
                :loading="$apollo.loading"
                label=" Select a team "
              >
        </v-select>
        <v-select
                v-model=selectedRole
                :items="roleTypes"
                item-text="text"
                item-value="value"
                label=" Select a Role "
              >
        </v-select>
      </v-col>
    </v-card>
    
    <v-card-actions>
      <v-btn
        color="primary"
      >
        Request Code
      </v-btn>
    </v-card-actions>


  </v-card>
  </v-col>
</template>

<script>
import {LEAGUE_TEAMS_BY_LEAGUE_ID} from "~/graphql/queries/league/leagueGraphQL";

export default {
  name: "GenerateAccessCodes",
  props: {
    leagueId: {
      type: String,
      required: true
    }
  },
  data: function () {
    return {
      teams: [],
      roleTypes: [
        {text: "League Manager", value:"LEAGUE_MANAGER"},
        {text: "Team Owner", value:"TEAM_OWNER"}
      ],
      selectedTeam: {}
    }
  },
  apollo: {
    teams: {
      query: LEAGUE_TEAMS_BY_LEAGUE_ID,
      variables() {
        return {leagueId: this.leagueId}
      }
    }
  },
}
</script>

<style scoped>

</style>
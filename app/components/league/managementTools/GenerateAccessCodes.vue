<template>
  <v-card>
    <h1 class="text-center">Generate Team Access Codes</h1>

    <v-select
                v-model:value="contract.teamId"
                :items="teams"
                item-text="teamName"
                item-value="id"
                :loading="$apollo.loading"
                label="Select a team"
                @change="teamSelected"
            >
    </v-select>

  </v-card>

</template>

<script>
import {LEAGUE_TEAMS_BY_LEAGUE_ID} from "~/graphql/queries/league/leagueGraphQL";

export default {
  name: "GenerateAccessCodes",
  components: {ContractManagementCard, ContractSearch},
  props: {
    leagueId: {
      type: String,
      required: true
    }
  },
  data: function () {
    return {
      selectedContract: null,
    }
  },
  methods: {
    contractSelected(contract) {
      this.selectedContract = contract.contract
    },
    contractDeselected() {
      this.selectedContract = null
    },
    contractModified() {
      this.selectedContract = null
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
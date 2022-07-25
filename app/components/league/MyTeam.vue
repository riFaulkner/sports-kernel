<template>
  <v-skeleton-loader
      :loading="$apollo.loading" type="article, card"
  >
    <v-card>
    <v-card-text>
      TeamID: {{ league }}
      <contract-search
        :contracts="teamContracts"
        :league-id="leagueId"
        :loading="this.$apollo.loading"
        />
    </v-card-text>

    </v-card>
  </v-skeleton-loader>

</template>

<script>
import {LEAGUE_FILTER_TEAMS_BY_OWNER_ID} from "@/graphql/queries/league/leagueGraphQL";
import ContractSearch from "@/components/searches/ContractSearch";

export default {
  name: "MyTeam.vue",
  components: {ContractSearch},
  props: {
    leagueId: {
      type: String,
      required: true
    },
    ownerId: {
      type: String,
      required: true
    }
  },
  data: function() {
    return {
      league: {},
    }
  },
  computed: {
    teamContracts() {
      if (this.league) {
        if (this.league.teams && this.league.teams.length > 0) {
          return this.league.teams[0].activeContracts
        }
        return []
      }
      return []
    },
    totalGuaranteed(contract) {
      contract.reduce((incompleteSum, year) => incompleteSum + year.guaranteedAmount, 0)
    }
  },
  apollo: {
    league: {
      query: LEAGUE_FILTER_TEAMS_BY_OWNER_ID,
      variables: function(){
        return {
          leagueId: this.leagueId,
          filter: {
            ownerId: this.ownerId
          }
        }
      }
    }
  }

}
</script>

<style scoped>

</style>
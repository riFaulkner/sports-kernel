<template>
  <v-skeleton-loader
      :loading="$apollo.loading" type="article, card"
  >
    <v-card>
      <v-card-text>
        <contract-search
            :contracts="teamContracts"
            :league-id="leagueId"
            :loading="this.$apollo.loading"
            :selected="selectedContract"
            @contract-selected="contractSelected"
            @contract-deselected="contractDeselected"
        />
      </v-card-text>
    </v-card>
    <v-card>
      <team-draft-picks-breakdown
          :draft-picks="this.league.teams[0]?.teamAssets?.draftPicks"
      />
    </v-card>
    <v-dialog
        v-model="contractIsSelected"
        max-width="500px"
    >
      <contract-management-card
          v-if="contractIsSelected"
          :contract="selectedContract[0]"
          :league-id="leagueId"
          @contract-restructured="contractModified"
          @contract-dropped="contractModified"
          @contract-management-closed="contractDeselected"
      />
    </v-dialog>
  </v-skeleton-loader>

</template>

<script>
import {LEAGUE_FILTER_TEAMS_BY_OWNER_ID} from "@/graphql/queries/league/leagueGraphQL";
import ContractSearch from "@/components/searches/ContractSearch";
import ContractManagementCard from "@/components/league/contracts/ContractManagementCard";
import TeamDraftPicksBreakdown from "@/components/league/team-assets/TeamDraftPicksBreakdown";

export default {
  name: "MyTeam.vue",
  components: {TeamDraftPicksBreakdown, ContractManagementCard, ContractSearch},
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
  data: function () {
    return {
      contractIsSelected: false,
      league: {},
      selectedContract: []
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
    teamDraftPicks() {

    }
  },
  methods: {
    contractSelected(contract) {
      this.selectedContract = [contract.contract]
      this.contractIsSelected = true
    },
    contractDeselected() {
      this.selectedContract = []
      this.contractIsSelected = false
    },
    contractModified() {
      this.selectedContract = []
    }
  },
  apollo: {
    league: {
      query: LEAGUE_FILTER_TEAMS_BY_OWNER_ID,
      variables: function () {
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
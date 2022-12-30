<template>
  <v-card>
    <h1 class="text-center">Manage Contract</h1>

    <v-card-text>
      <contract-search
          :league-id=leagueId
          :contracts=contracts
          :currentSeason=currentSeason
          :loading=this.$apollo.loading
          :selected=contractSelectList
          @contract-selected="contractSelected"
          @contract-deselected="contractDeselected"
      />
    </v-card-text>
    <div v-if="selectedContract !== null">
      <v-dialog
          :value="true"
          max-width="500px"
      >
        <contract-management-card
            :contract=selectedContract
            :current-season=currentSeason
            :league-id=leagueId
            @contract-restructured="contractModified"
            @contract-dropped="contractModified"
            @contract-management-closed="contractDeselected"
        />
      </v-dialog>

    </div>
  </v-card>

</template>

<script>
import ContractSearch from "~/components/searches/ContractSearch";
import ContractManagementCard from "@/components/league/contracts/ContractManagementCard";
import {LEAGUE_CONTRACTS} from "@/graphql/queries/league/leagueGraphQL";

export default {
  name: "ManageContract",
  components: {ContractManagementCard, ContractSearch},
  props: {
    leagueId: {
      type: String,
      required: true
    },
  },
  data: function () {
    return {
      selectedContract: null,
      contractSelectList:[],
      contracts: [],
      currentSeason: this.$store.getters["application/getActiveLeagueCurrentSeason"]
    }
  },
  methods: {
    contractSelected(contract) {
      this.selectedContract = contract.contract
      this.contractSelectList = [contract.contract]
    },
    contractDeselected() {
      this.selectedContract = null
      this.contractSelectList = []
    },
    contractModified() {
      this.selectedContract = null
      this.contractSelectList = []
    }
  },
  apollo: {
    contracts: {
      query: LEAGUE_CONTRACTS,
      variables() {
        return {leagueId: this.leagueId}
      },
      update: data => data.leagueContracts
    }
  }

}
</script>

<style scoped>

</style>
<template>
  <v-card>
    <h1 class="text-center">Manage Contract</h1>

    <v-card-text v-if="selectedContract === null">
      <contract-search
          :league-id="leagueId"
          :contracts="contracts"
          :loading="this.$apollo.loading"
          @contract-selected="contractSelected"
          @contract-deselected="contractDeselected"
      />
    </v-card-text>
    <div v-else>
      <contract-management-card
          :contract=selectedContract
          :league-id="leagueId"
          @contractRestructured="contractModified"
          @contractDropped="contractModified"
      />
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
    }
  },
  data: function () {
    return {
      selectedContract: null,
      contracts: []
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
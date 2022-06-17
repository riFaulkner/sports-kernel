<template>
  <v-card>
    <h1 class="text-center">Manage Contract</h1>

    <v-card-text v-if="selectedContract === null">
      <contract-search
          :league-id="leagueId"
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
  }
}
</script>

<style scoped>

</style>
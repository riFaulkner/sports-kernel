<template>
  <v-card>
    <h1 class="text-center"> Restructure Contract</h1>

    <v-card-text>
      <contract-search :league-id="leagueId"/>
    </v-card-text>

    <br/>
    TODO: Display the contract, if it is eligible for restructure then show that.

    <br/>
    TODO: Submit button

  </v-card>

</template>

<script>
import ContractSearch from "~/components/searches/ContractSearch";
export default {
  name: "RestructureContract",
  components: {ContractSearch},
  props: {
    leagueId: {
      type: String,
      required: true
    }
  }
}
</script>

<style scoped>

</style>
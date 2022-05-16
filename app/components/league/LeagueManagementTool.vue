<template>
  <v-container fluid>
    <v-row>
      <v-col cols="4" offset="4">
        <v-card>
          <v-select
              :items="toolTypes"
              label="Tools"
              hint="Select to get started"
              persistent-hint
              v-model=selectedTool
          />

        </v-card>
      </v-col>
      <v-col cols="4"/>

      <v-col>
        <restructure-contract
            v-if="selectedTool === 'restructure'"
            :league-id="leagueId"
        />
      </v-col>

    </v-row>
  </v-container>

</template>

<script>
import RestructureContract from "~/components/league/managementTools/RestructureContract";

export default {
  name: "LeagueManagementTool",
  components: {RestructureContract},
  props: {
    leagueId: {
      type: String,
      required: true
    }
  },
  data: function () {
    return {
      toolTypes: [
        {text: "Drop", value: "drop"},
        {text: "Restructure", value: "restructure"}
      ],
      selectedTool: ""
    }
  },
  methods: {
    transactionTypeSelected(event) {
      alert(event)
    }
  }
}
</script>

<style scoped>

</style>
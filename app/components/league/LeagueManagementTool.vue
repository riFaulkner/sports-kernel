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
        <manage-contract
            v-if="selectedTool === 'manageContract'"
            :league-id="leagueId"
        />
      </v-col>

    </v-row>
  </v-container>

</template>

<script>
import RestructureContract from "@/components/league/managementTools/ManageContract";
import ManageContract from "@/components/league/managementTools/ManageContract";

export default {
  name: "LeagueManagementTool",
  components: {ManageContract, RestructureContract},
  props: {
    leagueId: {
      type: String,
      required: true
    }
  },
  data: function () {
    return {
      toolTypes: [
        {text: "Create Contract", value: "createContract"},
        {text: "Manage Contract", value: "manageContract"},
        {text: "Add Player", value: "addPlayer"}
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
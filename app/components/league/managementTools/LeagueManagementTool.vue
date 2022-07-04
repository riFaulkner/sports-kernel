<template>
  <v-container fluid>
    <v-row>
      <v-col md="4" offset-md="4">
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

    </v-row>
    <v-row>
      <v-col v-if="selectedTool === 'manageContract'">
        <manage-contract
            :league-id="leagueId"
        />
      </v-col>
      <v-col md="4" offset-md="4" v-if="selectedTool === 'addPlayer'"
      >
        <add-player
        />
      </v-col>
    </v-row>
  </v-container>

</template>

<script>
import RestructureContract from "@/components/league/managementTools/ManageContract";
import ManageContract from "@/components/league/managementTools/ManageContract";
import AddPlayer from "@/components/players/AddPlayer"

export default {
  name: "LeagueManagementTool",
  components: {AddPlayer, ManageContract, RestructureContract},
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
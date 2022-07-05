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
      <v-col v-if="selectedTool === 'createContract'">
        <contract-creation-tool/>
      </v-col>

      <v-col v-if="selectedTool === 'manageContract'">
        <manage-contract
            :league-id="leagueId"
        />
      </v-col>
    </v-row>
  </v-container>

</template>

<script>
import RestructureContract from "@/components/league/managementTools/ManageContract";
import ManageContract from "@/components/league/managementTools/ManageContract";
import ContractCreationTool from "~/components/league/contracts/ContractCreationTool";

export default {
  name: "LeagueManagementTool",
  components: {ContractCreationTool, ManageContract, RestructureContract},
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
      ],
      selectedTool: ""
    }
  },
  methods: {
  }
}
</script>

<style scoped>

</style>
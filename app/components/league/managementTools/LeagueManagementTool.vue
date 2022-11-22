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
        <contract-creation-tool
            :league-id="leagueId"
        />
      </v-col>

      <v-col v-if="selectedTool === 'manageContract'">
        <manage-contract
            :league-id="leagueId"
        />
      </v-col>

      <v-col v-if="selectedTool === 'accessCodes'">
        <generate-access-codes
            :league-id="leagueId"
        />
      </v-col>
      <v-col v-if="selectedTool === 'deadCap'">
        <apply-dead-cap :league-id="leagueId"/>
      </v-col>
    </v-row>
  </v-container>

</template>

<script>
import RestructureContract from "@/components/league/managementTools/ManageContract";
import ManageContract from "@/components/league/managementTools/ManageContract";
import ContractCreationTool from "~/components/league/contracts/ContractCreationTool";
import GenerateAccessCodes from "@/components/league/managementTools/GenerateAccessCodes";
import ApplyDeadCap from "@/components/league/managementTools/ApplyDeadCap";

export default {
  name: "LeagueManagementTool",
  components: {ApplyDeadCap, ContractCreationTool, ManageContract, RestructureContract, GenerateAccessCodes},
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
        {text: "Generate League Access Codes", value: "accessCodes"},
        {text: "Add Dead Cap", value: "deadCap"}
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
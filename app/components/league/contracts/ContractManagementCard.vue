<template>
  <div v-if="contract">
    <template>
      <v-card
          color="grey darken-3"
          elevation="2"
          max-width="450px"
      >
        <v-card-title> {{ contract.player.playerName }}
          <v-spacer/>
          <v-btn
              icon
              dark
              @click="closeDialog"
          >
            <v-icon>mdi-close</v-icon>
          </v-btn>
        </v-card-title>
        <v-card-subtitle>
          Position: {{ contract.player.position }} &nbsp; | &nbsp; Team: {{ contract.player.team }}<br/>
        </v-card-subtitle>
        <v-divider/>
        <v-card-text>
          Current Year: {{ contract.currentYear }} <br/>
          Contract Duration: {{ contractDuration(contract) }} years <br/>
          Total Value: ${{ contract.totalContractValue.toLocaleString() }} <br/>
          Restructure Status:
          <strong class="success--text text--darken-1" v-if="canRestructure">
            {{ contract.restructureStatus }}
          </strong>
          <strong class="error--text" v-else>
            {{ contract.restructureStatus }}
          </strong> <br/>
          <br/>
            <v-progress-linear
                background-color="warning"
                :buffer-value="getContractPercentRemainingGuaranteed(contract) + getContractPercentPaid(contract)"
                :value=getContractPercentPaid(contract)
                color="green"
                stream
            ></v-progress-linear>
          <br/>

          <v-data-table
              :headers=headers
              :items="contract.contractDetails"
              disable-filtering
              disable-sort
              dense
              hide-default-footer
          >
            <template v-slot:item.season="{item}">
              {{calculateSeason(item, contract.currentYear)}}
            </template>
            <template v-slot:item.totalAmount="{item}">
              ${{ item.totalAmount.toLocaleString() }}
            </template>
            <template v-slot:item.guaranteedAmount="{item}">
              ${{ item.guaranteedAmount.toLocaleString() }}
            </template>
            <template v-slot:item.paidAmount="{item}">
              ${{ item.paidAmount.toLocaleString() }}
            </template>
          </v-data-table>
        </v-card-text>
        <v-card-actions>
          <v-btn
              :disabled="!canRestructure"
              @click="restructureContract"
          >
            Restructure
          </v-btn>
          <v-btn
              :disabled="!canDropPlayer"
              @click="dropContract"
          >
            Drop
          </v-btn>

        </v-card-actions>
      </v-card>
    </template>

    <v-dialog
        v-model="confirmationDialog"
        max-width="500px"
    >
      <v-card>
        <v-card-title>
          Are you sure you want to perform this action, it is irreversible.
        </v-card-title>

        <v-card-actions>
          <v-spacer/>
          <v-btn
              @click="cancelAction"
              color="error"
              text
          >
            Cancel
          </v-btn>
          <v-btn
              @click="performConfirmedAction"
              color="primary"
              text
          >
            Confirm
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-dialog
        v-model="contractRestructureDialog"
        max-width="750px"
    >
      <v-card

      >
        <v-card-title>
          Restructure Contract
        </v-card-title>
        <v-card-subtitle>
          *Restructuring will make all money guaranteed.
        </v-card-subtitle>
        <v-card-text>
          <v-form
              v-model="formValidation"
          >
            <v-container>
              <v-row>
                <v-col
                    cols="12"
                >
                  Original Total: ${{ contract.totalContractValue.toLocaleString() }} <br/>
                  Restructure Total: <strong :class="restructureTotalClass">
                  ${{ restructureTotal.toLocaleString() }}</strong>
                </v-col>
              </v-row>
              <v-row>
                <v-col
                    md="3"
                    cols="12"
                    v-for="year in tempContractRestructure"
                    :key="year.year"
                >
                  <v-text-field
                      :label=getRestructureTextFieldLabel(year.year)
                      :value="year.totalAmount"
                      v-model="year.totalAmount"
                      :disabled="year.year < contract.currentYear"
                      type="number"
                      :rules="restructureRules"
                  >
                  </v-text-field>

                </v-col>
              </v-row>
            </v-container>
          </v-form>
        </v-card-text>

        <v-card-actions>

          <v-spacer/>
          <v-btn
              @click="resetRestructureDialog"
              color="primary"
              text
          >Cancel
          </v-btn>
          <v-btn
              @click="submitRestructure"
              color="primary"
              :disabled="!validateRestructure"
              text
          >
            Submit
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>

</template>

<script>
import {CONTRACT_DROP, CONTRACT_RESTRUCTURE} from "@/graphql/queries/contract/contractsGraphQL";
import {LEAGUE_CONTRACTS} from "@/graphql/queries/league/leagueGraphQL";

export default {
  name: "ContractManagementCard",
  props: {
    contract: {
      type: Object,
      required: true
    },
    leagueId: {
      type: String,
      required: true
    },
    currentSeason: {
      type: Number,
      required: true
    }
  },
  data: function () {
    return {
      actionToPerform: null,
      confirmationDialog: false,
      contractRestructureDialog: false,
      formValidation: false,
      headers: [
        {text: "Season", value: "season"},
        {text: "Total", value: "totalAmount"},
        {text: "Guaranteed", value: "guaranteedAmount"},
        {text: "Paid", value: "paidAmount"}
      ],
      restructureRules: [
        value => !!value || 'Required.',
        value => this.minVal <= value || 'Restructured year must be at least 50% of the original value',
        value => this.maxVal >= value || 'Restructured year must be less than 200% of the original value',
      ],
      maxVal: (this.contract.contractDetails[0].totalAmount * 2),
      minVal: (this.contract.contractDetails[0].totalAmount * .50),
      tempContractRestructure: JSON.parse(JSON.stringify(this.contract.contractDetails)),
    }
  },
  computed: {
    canDropPlayer() {
      return true
    },
    canRestructure() {
      return this.contract.restructureStatus === 'ELIGIBLE'
    },
    restructureTotal() {
      let total = 0
      for (let i = 0; i < this.tempContractRestructure.length; i++) {
        total += parseInt(this.tempContractRestructure[i].totalAmount)
      }
      return total
    },
    validateRestructureTotal() {
      const restructureTotalValue = this.restructureTotal;
      const totalValue = this.contract.totalContractValue;

      return totalValue === restructureTotalValue
    },
    validateRestructure() {
      if (this.validateRestructureTotal && this.formValidation) {
        // Contract year values to not match the original contract details (a change was actually made)
        const originalContractDetails = this.contract.contractDetails
        let hasChange = false

        for (let i = 0; i < originalContractDetails.length; i++) {
          if (hasChange) continue
          const originalYearTotal = parseInt(originalContractDetails[i].totalAmount)
          const newYearTotal = parseInt(this.tempContractRestructure[i].totalAmount)

          if (originalYearTotal !== newYearTotal) {
            hasChange = true
          }
        }

        return hasChange
      }
      return false
    },
    restructureTotalClass() {
      if (this.validateRestructureTotal) {
        return "success--text"
      }
      return "error--text"
    }
  },
  methods: {
    calculateSeason(item, contractOffset){
      const offset = contractOffset - item.year
      return this.currentSeason - offset
    },
    closeDialog() {
      this.$emit("contract-management-closed")
    },
    contractDuration(contract) {
      return contract.contractDetails ? contract.contractDetails.length : 0
    },
    cancelAction() {
      this.actionToPerform = null
      this.confirmationDialog = false
      this.$store.dispatch("application/alertInfo", {message: "Action canceled"})
    },
    getContractPercentPaid(contract) {
      return ((contract.totalContractValue - contract.totalRemainingValue) / contract.totalContractValue) * 100
    },
    getContractPercentRemainingGuaranteed(contract) {
      const remainingGuaranteed = contract.contractDetails.filter(contractYear => contractYear.year >= contract.currentYear)
          .reduce((incompleteSum, currentValue) => incompleteSum + currentValue.guaranteedAmount, 0)
      return (remainingGuaranteed / contract.totalContractValue) * 100
    },
    getRestructureTextFieldLabel(year) {
      return `Year: ${year}`
    },
    performConfirmedAction() {
      if (this.actionToPerform === null) {
        this.$store.dispatch('application/alertError', {message: "Failed to perform action"})
      }

      this.actionToPerform()

      this.actionToPerform = null
      this.confirmationDialog = false
    },
    resetRestructureDialog() {
      this.contractRestructureDialog = false
      this.tempContractRestructure = JSON.parse(JSON.stringify(this.contract.contractDetails))
    },
    restructureContract() {
      if (!this.canRestructure) {
        this.$store.dispatch('application/alertError', {message: "This contract cannot be restructured"})
      }
      this.contractRestructureDialog = true
    },
    dropContract() {
      if (!this.canDropPlayer) {
        this.$store.dispatch('application/alertError', {message: "This contract cannot be dropped"})
      }
      const contractId = this.contract.id

      this.actionToPerform = () => {
        this.$apollo.mutate({
          mutation: CONTRACT_DROP,
          variables: {
            leagueId: this.leagueId,
            teamId: this.contract.teamId,
            contractId: contractId
          },
        }).then(() => {
          this.$store.dispatch("application/alertSuccess", {message: "Contract dropped"})
          this.$emit("contract-dropped", {contractId: contractId})
        }).catch((data) => {
          this.$store.dispatch("application/alertError", {message: "Failed to drop contract"})
          console.error("Failed to restructure contract ", data)
        })
      }

      this.confirmationDialog = true
    },
    submitRestructure() {
      // load function call into the action to perform data field
      const contractRestructure = {
        contractId: this.contract.id,
        contractRestructureDetails: this.tempContractRestructure
      }

      contractRestructure.contractRestructureDetails.forEach((year) => {
        year.guaranteedAmount = year.totalAmount
      });

      this.actionToPerform = () => {
        this.$apollo.mutate({
          mutation: CONTRACT_RESTRUCTURE,
          variables: {
            leagueId: this.leagueId,
            restructureDetails: contractRestructure,
          },
        }).then(() => {
          this.$store.dispatch("application/alertSuccess", {message: "Contract restructured"})
          this.$emit("contract-restructured", {contractId: contractRestructure.contractId})
        }).catch((data) => {
          this.$store.dispatch("application/alertError", {message: "Failed to restructure contract"})
          console.error("Failed to restructure contract ", data)
        })
      }

      this.contractRestructureDialog = false
      this.confirmationDialog = true
    }
  }
}
</script>

<style scoped>
::v-deep .v-progress-linear--visible .v-progress-linear__stream {
  -webkit-animation-play-state: paused;
  animation-play-state: paused;
}

</style>
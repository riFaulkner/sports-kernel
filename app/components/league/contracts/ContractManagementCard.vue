<template>
  <div v-if="contract">
    <template>
      <v-card
          color="grey darken-3"
          elevation="2"
          max-width="450px"
      >
        <v-card-title> {{ contract.player.playerName }}</v-card-title>
        <v-card-subtitle>
          Position: {{ contract.player.position }} <br/>
        </v-card-subtitle>
        <v-divider/>
        <v-card-text>
          Contract Length: {{ contract.contractLength }} years <br/>
          Total Value: ${{ contract.totalContractValue.toLocaleString() }} <br/>
          Restructure Status:
          <strong class="success--text text--darken-1" v-if="canRestructure">
            {{ contract.restructureStatus }}
          </strong>
          <strong class="error--text" v-else>
            {{ contract.restructureStatus }}
          </strong>
          <br/>
          Current Year: {{ contract.currentYear }} <br/>
          <br/>
          <v-data-table
              :headers=headers
              :items="contract.contractDetails"
              disable-filtering
              disable-sort
              dense
              hide-default-footer
          >
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
              disabled
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
          <v-container>
            <v-row>
              <v-col
                cols="12"
                >
                Restructure Total: <strong :class="restructureTotalClass"> ${{restructureTotal.toLocaleString()}}</strong>
              </v-col>
            </v-row>
            <v-row>
              <v-col
                  cols="3"
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
export default {
  name: "ContractManagementCard",
  props: {
    contract: {
      type: Object,
      required: true
    }
  },
  data: function () {
    return {
      actionToPerform: null,
      confirmationDialog: false,
      contractRestructureDialog: false,
      headers: [
        {text: "Year", value: "year"},
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
      tempContractRestructure: this.contract.contractDetails,
    }
  },
  computed: {
    canRestructure() {
      return this.contract.restructureStatus === 'ELIGIBLE'
    },
    restructureTotal() {
      let total = 0
      for (let i=0; i< this.tempContractRestructure.length; i++) {
        total += this.tempContractRestructure[i].totalAmount
      }
      return total
    },
    restructureTotalClass(){
      if (this.restructureTotal  === this.contract.totalValue){
        return "text--success"
      }
      return "text--error"
    }
  },
  methods: {
    cancelAction() {
      this.actionToPerform = null
      this.confirmationDialog = false
      this.$store.dispatch("application/alertInfo", {message: "Action canceled"})
    },
    getRestructureTextFieldLabel(year){
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
      this.tempContractRestructure = this.contract.contractDetails
      this.contractRestructureDialog = false
    },
    restructureContract() {
      if (!this.canRestructure) {
        this.$store.dispatch('application/alertError', {message: "This contract cannot be restructured"})
      }
      this.contractRestructureDialog = true
    },
    submitRestructure() {
      // generate the function call to make

      // load function call into the action to perform data field
      this.actionToPerform = () => {
        alert("Testing 123 perfoming an action")
      }

      // open up the confirmation dialog
      this.contractRestructureDialog = false
      this.confirmationDialog = true
    }
  }
}
</script>

<style scoped>

</style>
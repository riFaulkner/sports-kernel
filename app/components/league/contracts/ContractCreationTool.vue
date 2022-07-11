<template>
  <div>
    <v-stepper v-model="currentStep">
      <v-stepper-header>
        <v-stepper-step
            step="1"
            :complete="currentStep > 1"
        >
          Select Team
        </v-stepper-step>
        <v-divider/>
        <v-stepper-step
            step="2"
            :complete="currentStep > 2"
        >
          Select Player
        </v-stepper-step>
        <v-divider/>
        <v-stepper-step
            step="3"
            :complete="currentStep > 3"
        >
          Set Contract Parameters
        </v-stepper-step>
      </v-stepper-header>

      <v-stepper-items>
        <v-stepper-content
            step="1"
        >
          <v-card
              class="mb-12"
              color="grey darken-3"
          >
            <v-select
                v-model:value="contract.teamId"
                :items="teams"
                item-text="teamName"
                item-value="id"
                :loading="$apollo.loading"
                label="Select a team"
                @change="teamSelected"
            >
            </v-select>
          </v-card>

        </v-stepper-content>
        <v-stepper-content
            step="2"
        >
          <v-card>
            <player-search
                @player-selected="playerSelected"
            />
          </v-card>
          <br/>
          <v-btn
              color="primary"
              text
              @click="currentStep = 1"
          >
            Previous
          </v-btn>
        </v-stepper-content>
        <v-stepper-content
            step="3"
        >
          <v-card
              color="grey darken-3"
              elevation="5"
          >
            <v-card-title>Create Contract</v-card-title>
            <v-card-text>
              <v-form
                  v-model="formValidation"
              >
                <v-row>
                  <v-col cols="2">
                    <v-text-field
                        v-model="contract.contractLength"
                        label="Contract Length (Years)"
                        type="number"
                        disabled
                    />
                  </v-col>
                  <v-col cols="10">
                    <v-text-field
                        v-model="contract.currentYear"
                        label="Current Year"
                        type="number"
                        :rules="this.rules.currentContractYearRules"
                    />
                  </v-col>
                </v-row>
                <v-row>
                  <v-col>
                    <v-text-field
                        label="Contract Total Amount (Dollars)"
                        v-model="contractTotal"
                        disabled
                    />
                  </v-col>
                  <v-col>
                    <v-text-field
                        label="Contract Guaranteed Amount (Dollars)"
                        v-model="contractGuaranteed"
                        disabled
                    />
                  </v-col>
                  <v-col>
                    <v-text-field
                        label="Contract Remaining (Dollars)"
                        v-model="contractRemaining"
                        disabled
                    />
                  </v-col>
                </v-row>
                <br/>
                <v-divider/>
                <br/>

                <div v-for="year in contract.contractDetails">
                  <v-row>
                    <v-col cols="1">
                      <v-text-field
                          v-model="year.year"
                          label="Year: "
                          disabled
                      />
                    </v-col>
                    <v-col>
                      <v-text-field
                          v-model="year.totalAmount"
                          label="Total contract value:"
                          @input="yearTotalUpdated(year)"
                          :rules="rules.yearTotalAmountRules"
                          type="number"
                      />
                    </v-col>
                    <v-col>
                      <v-text-field
                          v-model="year.guaranteedAmount"
                          label="Guaranteed value:"
                          type="number"
                      />
                    </v-col>
                    <v-col>
                      <v-text-field
                          v-model="year.paidAmount"
                          label="Amount paid:"
                          type="number"
                          @input="yearPaidUpdated(year)"
                      />
                    </v-col>
                  </v-row>
                </div>
                <v-row>
                  <v-col>
                    <v-icon
                        @click="removeContractYear"
                        :disabled="contract.contractDetails.length === 1"
                    >mdi-minus-circle
                    </v-icon>
                    <v-icon
                        @click="addContractYear"
                        color="primary"
                        :disabled="contract.contractDetails.length === 4"
                    >mdi-plus-circle
                    </v-icon>
                  </v-col>

                  <v-spacer/>
                  <v-col cols="1">
                    <v-spacer/>
                    <v-btn
                        :disabled="!validateContract"
                        @click="createContract"
                        color="primary"
                    >Create</v-btn>
                  </v-col>

                </v-row>
              </v-form>

            </v-card-text>
          </v-card>
          <br/>
          <v-btn
              color="primary"
              text
              @click="currentStep = 2"
          >
            Previous
          </v-btn>
        </v-stepper-content>
      </v-stepper-items>
    </v-stepper>
  </div>
</template>

<script>
import {LEAGUE_TEAMS_BY_LEAGUE_ID} from "~/graphql/queries/league/leagueGraphQL";
import PlayerSearch from "~/components/searches/PlayerSearch";
import {CREATE_CONTRACT} from "~/graphql/queries/contract/contractsGraphQL";

export default {
  name: "ContractCreationTool",
  components: {PlayerSearch},
  props: {
    leagueId: {
      type: String,
      required: true
    }
  },
  data: function () {
    return {
      contract: this.getDefaultContract(),
      currentStep: 1,
      formValidation: false,
      rules: {
        currentContractYearRules: [
          value => !!value || "Current year Required",
          value => 0 < value || "Current year must be greater than 1",
          value => value <= 4 || "Current year must be less than or equal to 4",
          value => value <= this.contract.contractLength || "Current year must be less than contract durration"
        ],
        yearTotalAmountRules: [
            value => !!value || "Year total required",
            value => value > 0 || "Year total must be positive",
        ],
        yearGuaranteedAmountRules: [
          value => !!value || "Guaranteed amount required",
          value => value > 0 || "Guaranteed amount must be positive",
        ],
        yearPaidAmountRules: [
          value => !!value || "Paid amount required",
          value => value >= 0 || "Paid amount must be positive",
        ],
      },
      selectedTeam: {},
      selectedPlayer: {},
      teams: []
    }
  },
  computed: {
    contractTotal() {
      return "$" + this.contract.totalContractValue.toLocaleString()
    },
    contractGuaranteed() {
      return "$" + parseInt(this.contract.contractDetails.reduce((partialSum, year) => partialSum + year.guaranteedAmount, 0)).toLocaleString()
    },
    contractRemaining() {
      return "$" + this.contract.totalRemainingValue.toLocaleString()
    },
    validateContract() {
      if (this.formValidation) {
        return this.contract.contractDetails.filter((year) =>
          year.totalAmount < year.guaranteedAmount || year.totalAmount < year.paidAmount
        ).length === 0
      }
      return false
    }
  },
  methods: {
    addContractYear() {
      const contractLength = this.contract.contractDetails.length
      this.contract.contractDetails.push({
        "guaranteedAmount": this.contract.contractDetails[0].guaranteedAmount,
        "paidAmount": this.contract.contractDetails[0].paidAmount,
        "totalAmount": this.contract.contractDetails[0].totalAmount,
        "year": contractLength + 1
      })
      this.contract.contractLength++
      this.updateContractTotals()
    },
    createContract() {
      if (this.validateContract) {
        if(this.contract.contractLength === this.contract.currentYear) {
          this.contract.restructureStatus = "INELIGIBLE_FINAL_YEAR"
        }
        this.$apollo.mutate({
          mutation: CREATE_CONTRACT,
          variables: {
            contract: this.contract,
            leagueId: this.leagueId
          }
        }).then(() => {
          this.$store.dispatch("application/alertSuccess", {message: "Contract created"})
        }).catch((error)=> {
          console.error(error)
          this.$store.dispatch("application/alertError", {message: "Error occured creating contract"})
        })
      }
    },
    removeContractYear() {
      this.contract.contractDetails.pop()
      this.contract.contractLength--
      this.updateContractTotals()
    },
    playerSelected(event) {
      this.selectedPlayer = event.player
      this.contract.playerId = this.selectedPlayer.id
      this.contract.playerPosition = this.selectedPlayer.position
      this.currentStep = 3
    },
    teamSelected(event) {
      this.selectedTeam = this.teams.filter(team => team.id === event)
      this.currentStep = 2
    },
    getDefaultContract() {
      return {
        "contractDetails": [
          {
            "guaranteedAmount": 0,
            "paidAmount": 0,
            "totalAmount": 0,
            "year": 1
          }
        ],
        "contractLength": 1,
        "contractStatus": "ACTIVE",
        "currentYear": 1,
        "playerId": "",
        "playerPosition": "",
        "restructureStatus": "ELIGIBLE",
        "teamId": "",
        "totalContractValue": 0,
        "totalRemainingValue": 0
      }
    },
    yearTotalUpdated(contractYearDetails) {
      contractYearDetails.guaranteedAmount = parseInt(contractYearDetails.totalAmount) * .5
      this.updateContractTotals()
    },
    yearPaidUpdated() {
      this.updateContractTotals()
    },
    updateContractTotals() {
      const totalContractValue = this.contract.contractDetails.reduce((partialSum, year) => partialSum + parseInt(year.totalAmount), 0)

      this.contract.totalContractValue = totalContractValue
      this.contract.totalRemainingValue = totalContractValue - this.contract.contractDetails.reduce((partialSum, year) => partialSum + parseInt(year.paidAmount), 0)
    }
  },
  apollo: {
    teams: {
      query: LEAGUE_TEAMS_BY_LEAGUE_ID,
      variables() {
        return {leagueId: this.leagueId}
      }
    }
  },
}
</script>

<style scoped>

</style>
<template>
  <v-skeleton-loader
      :loading="$apollo.loading" type="article, card"
  >
    <v-card>
      <v-card-title v-if="this.league.teams?.length > 0">
        <v-spacer/>
        <h2> {{ this.league.teams[0].teamName }} </h2>
        <v-spacer/>
      </v-card-title>
      <v-card-text>
        <contract-search
            :contracts="teamContracts"
            :league-id="leagueId"
            :loading="this.$apollo.loading"
            :selected="selectedContract"
            @contract-selected="contractSelected"
            @contract-deselected="contractDeselected"
        />
      </v-card-text>
      <br/>
      <v-row>
        <v-col cols="12" md="8">
          <team-draft-picks-breakdown
              v-if="this.league.teams?.length > 0"
              :draft-picks="this.league.teams[0].teamAssets?.draftPicks"
          />
        </v-col>
        <v-divider vertical/>
        <v-col cols="12" md="4">
          <team-dead-cap-breakdown
              v-if="this.league.teams?.length > 0"
              :dead-cap="this.league.teams[0].teamLiabilities?.deadCap"
          />
        </v-col>
      </v-row>
    </v-card>

    <v-dialog
        v-model="contractIsSelected"
        max-width="500px"
    >
      <contract-management-card
          v-if="contractIsSelected"
          :contract="selectedContract[0]"
          :league-id="leagueId"
          @contract-restructured="contractRestructured"
          @contract-dropped="contractDropped"
          @contract-management-closed="contractDeselected"
      />
    </v-dialog>
  </v-skeleton-loader>

</template>

<script>
import {LEAGUE_FILTER_TEAMS_BY_OWNER_ID} from "@/graphql/queries/league/leagueGraphQL";
import ContractSearch from "@/components/searches/ContractSearch";
import ContractManagementCard from "@/components/league/contracts/ContractManagementCard";
import TeamDraftPicksBreakdown from "@/components/league/team-assets/TeamDraftPicksBreakdown";
import TeamDeadCapBreakdown from "@/components/league/team-assets/TeamDeadCapBreakdown";
import {TEAM_DEAD_CAP} from "@/graphql/queries/team/teamGraphQL";
import {GET_CONTACT_BY_ID} from "@/graphql/queries/contract/contractsGraphQL";

export default {
  name: "MyTeam.vue",
  components: {TeamDeadCapBreakdown, TeamDraftPicksBreakdown, ContractManagementCard, ContractSearch},
  props: {
    leagueId: {
      type: String,
      required: true
    },
    ownerId: {
      type: String,
      required: true
    }
  },
  data: function () {
    return {
      contractIsSelected: false,
      league: {},
      selectedContract: []
    }
  },
  computed: {
    teamContracts() {
      if (this.league) {
        if (this.league.teams && this.league.teams.length > 0) {
          return this.league.teams[0].activeContracts
        }
        return []
      }
      return []
    },
  },
  methods: {
    contractSelected(contract) {
      this.selectedContract = [contract.contract]
      this.contractIsSelected = true
    },
    contractDeselected() {
      this.selectedContract = []
      this.contractIsSelected = false
    },
    contractDropped(droppedContract) {
      this.contractIsSelected = false
      this.selectedContract = []

      this.league.teams[0].activeContracts = this.league.teams[0]
          .activeContracts
          .slice()
          .filter(contract => contract.id !== droppedContract.contractId)

      this.updateTeamDeadCap()
    },
    contractRestructured(contractInfo) {
      this.contractIsSelected = false
      this.selectedContract = []
      this.updateSingleContractInfo(contractInfo.contractId)
    },
    updateSingleContractInfo(contractId) {
      this.$apollo.query(({
        query: GET_CONTACT_BY_ID,
        variables: {
          leagueId: this.leagueId,
          contractId: contractId
        }
      })).then((response) => {
        const newContract = response.data.contractById

        const allOtherContracts = this.league.teams[0].activeContracts.slice().filter(contract => contract.id !== contractId)

        allOtherContracts.push(newContract)
      }).catch((error) => {
        console.error("ERROR: ", error)
      })
    },
    updateTeamDeadCap() {
      this.$apollo.query(({
        query: TEAM_DEAD_CAP,
        variables: {
          leagueId: this.leagueId,
          teamId: this.league.teams[0].id
        }
      })).then((response) => {
        this.league.teams[0].teamLiabilities.deadCap = response.data.teamById.teamLiabilities.deadCap

      }).catch((error) => {
        this.$store.dispatch("application/alertError", {message: "Error fetching updated dead cap"})
        console.error("ERROR: ", error)
      })
    }
  },
  apollo: {
    league: {
      query: LEAGUE_FILTER_TEAMS_BY_OWNER_ID,
      variables: function () {
        return {
          leagueId: this.leagueId,
          filter: {
            ownerId: this.ownerId
          }
        }
      }
    }
  }
}
</script>

<style scoped>

</style>
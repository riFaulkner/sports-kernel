<template>
  <v-skeleton-loader :loading="$apollo.loading" type="article, card">
    <v-container fluid>
      <v-row >
        <v-col cols="12" sm="3">
          <team-contracts-breakdown :title="assets.quarterBacksTitle" :contracts=quarterBackContracts />
        </v-col>
        <v-col cols="12" sm="3">
          <team-contracts-breakdown :title="assets.runningBacksTitle" :contracts=runningBackContracts />
        </v-col>
        <v-col cols="12" sm="3">
          <team-contracts-breakdown :title="assets.wideReceiversTitle" :contracts=wideReceiverContracts />
        </v-col>
        <v-col cols="12" sm="3">
          <team-contracts-breakdown :title="assets.tightEndsTitle" :contracts=tightEndsContracts />
        </v-col>
      </v-row>
      <v-row>
        <team-draft-picks-breakdown
            :draft-picks="this.teamAssets?.draftPicks"
        />
      </v-row>
      <v-row>
        <v-spacer/>
        <v-btn color="grey darken-3" disabled>Propose Trade</v-btn>
      </v-row>
      <br/>
    </v-container>
  </v-skeleton-loader>
</template>

<script>
import DraftBreakDownList from "@/components/league/team-assets/DraftBreakDownList";
import TeamContractsBreakdown from "@/components/league/team-assets/TeamContractsBreakdown";
import {TEAM_CONTRACTS, TEAM_DRAFT_PICKS} from "@/graphql/queries/team/teamGraphQL";
import TeamDraftPicksBreakdown from "@/components/league/team-assets/TeamDraftPicksBreakdown";
export default {
  name: "TeamAssetsBreakdown",
  components: {TeamDraftPicksBreakdown, TeamContractsBreakdown, DraftBreakDownList},
  props: {
    teamId: {
      type: String,
      required: true
    },
    leagueId: {
      type: String,
      required: true
    }
  },
  data: function () {
    return {
      contracts: [],
      teamAssets: {
        draftPicks: []
      },
      assets: {
        draftPicksTitle: "Draft Picks",
        quarterBacksTitle: "Quarter Backs",
        runningBacksTitle: "Running Backs",
        wideReceiversTitle: "Wide Receivers",
        tightEndsTitle: "Tight Ends",
      }
    }
  },
  computed: {
    quarterBackContracts: function() {
      return this.contracts.filter(contract => contract.playerPosition === "QB")
    },
    runningBackContracts: function() {
      return this.contracts.filter(contract => contract.playerPosition === "RB")
    },
    wideReceiverContracts: function () {
      return this.contracts.filter(contract => contract.playerPosition === "WR")
    },
    tightEndsContracts: function () {
      return this.contracts.filter(contract => contract.playerPosition === "TE")
    }
  },
  apollo: {
    contracts: {
      query: TEAM_CONTRACTS,
      variables() {
        return {
          leagueId: this.leagueId,
          teamId: this.teamId
        }
      },
      update: data => data.teamContracts,
    },
    teamAssets: {
      query: TEAM_DRAFT_PICKS,
      variables() {
        return {
          leagueId: this.leagueId,
          teamId: this.teamId
        }
      },
      update: data => data.teamById.teamAssets,
    }
  }
}
</script>

<style scoped>

</style>
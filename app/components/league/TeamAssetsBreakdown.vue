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
        <v-data-table
            :headers="assets.draftPicksHeaders"
            :items="draftPicks"
            item-key="title"
            hide-default-footer
        >
          <template v-slot:item.draft1="{item}">
            <draft-break-down-list :picks="item.picks[0]" />
          </template>
          <template v-slot:item.draft2="{item}">
            <draft-break-down-list :picks="item.picks[1]" />
          </template>
          <template v-slot:item.draft3="{item}">
            <draft-break-down-list :picks="item.picks[2]" />
          </template>
          <template v-slot:item.draft4="{item}">
            <draft-break-down-list :picks="item.picks[3]" />
          </template>
          <template v-slot:item.draft5="{item}">
            <draft-break-down-list :picks="item.picks[4]" />
          </template>
        </v-data-table>
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
import DraftBreakDownList from "~/components/league/DraftBreakDownList";
import TeamContractsBreakdown from "~/components/league/TeamContractsBreakdown";
import {TEAM_CONTRACTS} from "~/graphql/queries/team/teamGraphQL";
export default {
  name: "TeamAssetsBreakdown",
  components: {TeamContractsBreakdown, DraftBreakDownList},
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
        draftPicks: [
          {
            year: 2022,
            picks: [
              {round:1, pickValue: 1},
              {round:1, pickValue: 15},
              {round:2, pickValue: 33},
              {round:3, pickValue: 65},
            ]
          },
          {
            year: 2023,
            picks: [
              {round:1, pickValue: 1},
              {round:2, pickValue: 33},
              {round:3, pickValue: 65},
            ]
          },
          {
            year: 2024,
            picks: [
              {round:1, pickValue: null},
              {round:2, pickValue: null},
              {round:3, pickValue: null},
            ]
          },
          {
            year: 2025,
            picks: [
              {round:1, pickValue: null},
              {round:2, pickValue: null},
              {round:3, pickValue: null},
              {round:4, pickValue: null},
            ]
          },
          {
            year: 2026,
            picks: [
              {round:1, pickValue: null},
              {round:2, pickValue: null},
              {round:3, pickValue: null},
              {round:4, pickValue: null},
            ]
          }
        ]
      },
      assets: {
        draftPicksTitle: "Draft Picks",
        quarterBacksTitle: "Quarter Backs",
        runningBacksTitle: "Running Backs",
        wideReceiversTitle: "Wide Receivers",
        tightEndsTitle: "Tight Ends",
        draftPicksHeaders: [
          {text: "", value: "round"},
          {text: '2022', value: "draft1"},
          {text: '2023', value: "draft2"},
          {text: '2024', value: "draft3"},
          {text: '2025', value: "draft4"},
          {text: '2026', value: "draft5"},
        ]
      }
    }
  },
  computed: {
    draftPicks: function() {
      let tableObject = [];

      for (let roundNum = 0; roundNum < 4; roundNum++) {
        let roundDataList = [];
        const roundNumber = roundNum+1;

        this.teamAssets.draftPicks.forEach(draftYear => {
            roundDataList.push(draftYear.picks.filter(pick => {
              return pick.round === roundNumber
          }))
        });

        const roundString = "Round " + roundNumber;

        let roundData = {round: roundString, picks: roundDataList}
        tableObject.push(roundData)
      }

      return tableObject
    },
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
    }
  }
}
</script>

<style scoped>

</style>
<template>
  <v-skeleton-loader
      :loading="$apollo.loading" type="article, card"
  >
    <v-card>
    <v-card-text>
      TeamID: {{ league }}
      <v-data-table
          :headers="headers"
          :items="teamContracts"
      >
        <template v-slot:item.>

        </template>
      </v-data-table>
    </v-card-text>

    </v-card>
  </v-skeleton-loader>

</template>

<script>
import {LEAGUE_FILTER_TEAMS_BY_OWNER_ID} from "@/graphql/queries/league/leagueGraphQL";

export default {
  name: "MyTeam.vue",
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
  data: function() {
    return {
      headers: [
        {text: "Player", value: "player.playerName"},
        {text: "NFL team", value: "player.team"},
        {text: "Current Fantasy Points", value: "0"},
        {text: "Points Contributed", value: "0"},
        {text: "Cost per point", value: "0"},
        {text: "Contract Total", value: "totalContractValue"},
        {text: "Contract Paid", value: "totalRemainingValue"},
        {text: "Guaranteed Value", value: "totalGuaranteed(item)"},
        {text: "Year One Total", value: "contractDetails[0].totalAmount"},
        {text: "Year One Paid", value: "contractDetails[0].paidAmount"},
        {text: "Year One Guaranteed ", value: "contractDetails[0].guaranteedAmount"},
        {text: "Year Two Total", value: "contractDetails[1].totalAmount"},
        {text: "Year Two Paid", value: "contractDetails[1].paidAmount"},
        {text: "Year Two Guaranteed ", value: "contractDetails[1].guaranteedAmount"},
        {text: "Year Three Total", value: "contractDetails[2].totalAmount"},
        {text: "Year Three Paid", value: "contractDetails[2].paidAmount"},
        {text: "Year Three Guaranteed ", value: "contractDetails[2].guaranteedAmount"},
        {text: "Edit", value: "0"},
      ],
      league: {},
    }
  },
  computed: {
    teamContracts() {
      if (this.league ) {
        if (this.league.teams && this.league.teams.length > 0) {
          return this.league.teams[0].activeContracts
        }
        return []
      }
      return []
    },
    totalGuaranteed(contract) {
      contract.reduce((incompleteSum, year) => incompleteSum + year.guaranteedAmount, 0)
    }
  },
  apollo: {
    league: {
      query: LEAGUE_FILTER_TEAMS_BY_OWNER_ID,
      variables: function(){
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
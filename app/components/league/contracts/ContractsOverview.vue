<template>
  <v-card>
    <v-card-title class="justify-center">
      {{ leagueInfo.leagueName }}
    </v-card-title>
    <v-data-table
        :headers="headers"
        :items="teams"
        :expanded.sync="expanded"
        item-key="teamName"
        :items-per-page="15"
        hide-default-footer
        show-expand
    >
      <template v-slot:item.currentContractsMetadata.totalUtilizedCap="{item}">
        ${{ item.currentContractsMetadata.totalUtilizedCap.toLocaleString() }}
      </template>
      <template v-slot:item.currentContractsMetadata.totalAvailableCap="{ item }">
        <v-chip
            :color="getColor(item.currentContractsMetadata.totalAvailableCap)"
            dark
        >
          ${{ item.currentContractsMetadata.totalAvailableCap.toLocaleString() }}
        </v-chip>
      </template>
      <template v-slot:item.currentContractsMetadata.qbUtilizedCap="{ item }">
        ${{ item.currentContractsMetadata.qbUtilizedCap.capUtilization.toLocaleString() }}
        ({{ item.currentContractsMetadata.qbUtilizedCap.numContracts }})
      </template>
      <template v-slot:item.currentContractsMetadata.rbUtilizedCap="{ item }">
        ${{ item.currentContractsMetadata.rbUtilizedCap.capUtilization.toLocaleString() }}
        ({{ item.currentContractsMetadata.rbUtilizedCap.numContracts }})
      </template>
      <template v-slot:item.currentContractsMetadata.wrUtilizedCap="{ item }">
        ${{ item.currentContractsMetadata.wrUtilizedCap.capUtilization.toLocaleString() }}
        ({{ item.currentContractsMetadata.wrUtilizedCap.numContracts }})
      </template>
      <template v-slot:item.currentContractsMetadata.teUtilizedCap="{ item }">
        ${{ item.currentContractsMetadata.teUtilizedCap.capUtilization.toLocaleString() }}
        ({{ item.currentContractsMetadata.teUtilizedCap.numContracts }})
      </template>
      <template v-slot:item.currentContractsMetadata.deadCap="{ item }" >
        ${{ item.currentContractsMetadata.deadCap.capUtilization.toLocaleString() }}
        ({{ item.currentContractsMetadata.deadCap.numContracts }})

      </template>
      <template v-slot:expanded-item="{ headers, item }" v-slot:>
        <td :colspan="headers.length">
          <team-assets-breakdown :team-id="item.id" :league-id="leagueInfo.id"/>
        </td>
      </template>
    </v-data-table>
  </v-card>
</template>

<script>
import TeamAssetsBreakdown from '@/components/league/team-assets/TeamAssetsBreakdown'
import {LEAGUE_TEAMS_BY_LEAGUE_ID} from '@/graphql/queries/league/leagueGraphQL'

export default {
  name: 'ContractsOverview.vue',
  components: {TeamAssetsBreakdown},
  props: {
    leagueInfo: {
      type: Object,
      require: true,
      validator: function (value) {
        return true
      }
    }
  },
  data: function () {
    return {
      expanded: [],
      headers: [
        {text: 'Team', value: 'teamName', width: 200},
        {text: 'Total Cap Utilization', value: 'currentContractsMetadata.totalUtilizedCap', width: 105},
        {
          text: 'Available Cap',
          value: 'currentContractsMetadata.totalAvailableCap',
          align: 'start',
          divider: true,
          width: 75
        },
        {
          text: 'QB Value (# Contracts)',
          value: 'currentContractsMetadata.qbUtilizedCap',
          align: 'center',
          sort: (a, b) => (a.capUtilization - b.capUtilization)
        },
        {
          text: 'RB Value (# Contracts)',
          value: 'currentContractsMetadata.rbUtilizedCap',
          align: 'center',
          sort: (a, b) => (a.capUtilization - b.capUtilization)
        },
        {
          text: 'WR Value (# Contracts)',
          value: 'currentContractsMetadata.wrUtilizedCap',
          align: 'center',
          sort: (a, b) => (a.capUtilization - b.capUtilization)
        },
        {
          text: 'TE Value (# Contracts)',
          value: 'currentContractsMetadata.teUtilizedCap',
          align: 'center',
          sort: (a, b) => (a.capUtilization - b.capUtilization)
        },
        {
          text: 'Dead Cap (# Contracts)',
          value: 'currentContractsMetadata.deadCap',
          align: 'center',
        },
        {
          text: '',
          value: 'data-table-expand'
        }
      ],
      teams: []
    }
  },
  methods: {
    getColor(capRemaining) {
      if (capRemaining < 5000000) {
        return 'error'
      }
      if (capRemaining < 10000000) {
        return 'warning'
      }
      return 'success'
    },
  },
  apollo: {
    teams: {
      query: LEAGUE_TEAMS_BY_LEAGUE_ID,
      variables() {
        return {leagueId: this.leagueInfo.id}
      }
    }
  }
}
</script>

<style scoped>

</style>

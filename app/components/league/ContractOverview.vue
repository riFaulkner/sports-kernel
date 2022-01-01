<template>
  <v-card>
    <v-card-title class="justify-center" > {{leagueInfo.leagueName}}</v-card-title>
    <v-data-table
      :headers=headers
      :items=items
      hide-default-footer
    >
      <template v-slot:item.totalUtilizedCap="{item}">
        ${{item.totalUtilizedCap}}
      </template>
      <template v-slot:item.totalAvailableCap="{ item }">
        <v-chip
            :color="getColor(item.totalAvailableCap)"
            dark
        >
          ${{ item.totalAvailableCap }}
        </v-chip>
      </template>
      <template v-slot:item.qbUtilizedCap="{ item }">
        ${{item.qbUtilizedCap.capUtilization}} ({{item.qbUtilizedCap.numContracts}})
      </template>
      <template v-slot:item.rbUtilizedCap="{ item }">
        ${{item.rbUtilizedCap.capUtilization}} ({{item.rbUtilizedCap.numContracts}})
      </template>
      <template v-slot:item.wrUtilizedCap="{ item }">
        ${{item.wrUtilizedCap.capUtilization}} ({{item.wrUtilizedCap.numContracts}})
      </template>
      <template v-slot:item.teUtilizedCap="{ item }">
        ${{item.teUtilizedCap.capUtilization}} ({{item.teUtilizedCap.numContracts}})
      </template>
    </v-data-table>
  </v-card>
</template>

<script>
export default {
  name: "LeagueStandings.vue",
  props: [
    'leagueInfo'
  ],
  data: function() {
    return {
      headers: [
          {text:'Team', value: 'teamName', width: 200},
          {text:'Total Cap Utilization', value: 'totalUtilizedCap', width: 105},
          {text:'Available Cap', value: 'totalAvailableCap', align: 'start',divider: true, width: 75},
          {text:'QB Value (# Contracts)', value: 'qbUtilizedCap', align: 'center', sort: (a, b) => (a.capUtilization - b.capUtilization)},
          {text:'RB Value (# Contracts)', value: 'rbUtilizedCap', align: 'center', sort: (a, b) => (a.capUtilization - b.capUtilization)},
          {text:'WR Value (# Contracts)', value: 'wrUtilizedCap', align: 'center', sort: (a, b) => (a.capUtilization - b.capUtilization)},
          {text:'TE Value (# Contracts)', value: 'teUtilizedCap', align: 'center', sort: (a, b) => (a.capUtilization - b.capUtilization)},
      ],
      items: [
        {
          teamName: "Rick's team",
          totalUtilizedCap: 180,
          totalAvailableCap: 20,
          qbUtilizedCap: { capUtilization: 10, numContracts:3 },
          rbUtilizedCap: { capUtilization: 10, numContracts:3 },
          wrUtilizedCap: { capUtilization: 10, numContracts:3 },
          teUtilizedCap: { capUtilization: 10, numContracts:3 },
        },{
          teamName: "Jeff's team",
          totalUtilizedCap: 199,
          totalAvailableCap: 1,
          qbUtilizedCap: { capUtilization: 20, numContracts:3 },
          rbUtilizedCap: { capUtilization: 20, numContracts:3 },
          wrUtilizedCap: { capUtilization: 20, numContracts:3 },
          teUtilizedCap: { capUtilization: 20, numContracts:3 },
        },
        {
          teamName: "Kyle's team",
          totalUtilizedCap: 190,
          totalAvailableCap: 10,
          qbUtilizedCap: { capUtilization: 30, numContracts:3 },
          rbUtilizedCap: { capUtilization: 30, numContracts:3 },
          wrUtilizedCap: { capUtilization: 30, numContracts:3 },
          teUtilizedCap: { capUtilization: 30, numContracts:3 },
        },
        {
          teamName: "Ramzi's team",
          totalUtilizedCap: 195,
          totalAvailableCap: 5,
          qbUtilizedCap: { capUtilization: 40, numContracts:3 },
          rbUtilizedCap: { capUtilization: 40, numContracts:3 },
          wrUtilizedCap: { capUtilization: 40, numContracts:3 },
          teUtilizedCap: { capUtilization: 40, numContracts:3 },
        }
      ]
    }
  },
  methods: {
    getColor(capRemaining) {
      if (capRemaining < 5) {
        return 'error';
      }
      if (capRemaining < 10) {
        return 'warning'
      }
      return 'success';
    }
  }
}
</script>

<style scoped>

</style>
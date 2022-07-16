<template v-slot:extension>
  <v-tabs-items
    v-if="league"
    v-model="tab"
    touchless
  >
    <v-tab-item>
      <league-standings
        :league-info="league"
      />
    </v-tab-item>
    <v-tab-item>
      <my-team
        :league-id="league.id"
      />
    </v-tab-item>
    <v-tab-item>
      <v-card>
        <v-card-text>
          <contracts-overview
            :league-info="league"
          />
        </v-card-text>
      </v-card>
    </v-tab-item>
    <v-tab-item>
      <v-card>
        <v-card-text>
          <rule-summary />
        </v-card-text>
      </v-card>
    </v-tab-item>
    <v-tab-item>
      <league-management-tool :league-id="league.id" />
    </v-tab-item>
  </v-tabs-items>
  <div v-else class="text-center">
    <h1>
      Please select a league
    </h1>
  </div>
</template>

<script>
import ContractsOverview from '@/components/league/contracts/ContractsOverview'
import LeagueManagementTool from '@/components/league/managementTools/LeagueManagementTool'
import MyTeam from '@/components/league/MyTeam'
import RuleSummary from '@/components/league/RuleSummary'

export default {
  name: 'league-home.vue',
  components: { LeagueManagementTool, ContractsOverview, MyTeam, RuleSummary },
  middleware: 'auth',
  data: function () {
    return {
    }
  },
  computed: {
    league () {
      return this.$store.getters['application/getActiveLeague']
    },
    tab () {
      return this.$store.state.application.activeTab
    },
    submenu () {
      return this.$store.state.application.submenu
    }
  },
  created () {
    this.$store.dispatch('application/updateSubmenu', [
      'Standings', 'My Team', 'Contracts Overview', 'Rules', 'League Management'
    ])
    this.$store.dispatch('application/updateActiveTab', 'Rules')
  },
  destroyed () {
    this.$store.dispatch('application/updateSubmenu', null)
    this.$store.dispatch('application/updateActiveTab', null)
  }
}
</script>

<style scoped>

</style>

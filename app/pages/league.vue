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
      <week-scoring :league-id="league.id"/>
    </v-tab-item>
    <v-tab-item>
      <my-team
        :league-id="league.id"
        :owner-id="ownerId"
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
import WeekScoring from "@/components/league/scoring/WeekScoring";

export default {
  name: 'league-home.vue',
  components: {WeekScoring, LeagueManagementTool, ContractsOverview, MyTeam, RuleSummary },
  middleware: 'auth',
  data: function () {
    return {
    }
  },
  computed: {
    league () {
      return this.$store.getters['application/getActiveLeague']
    },
    ownerId() {
      return this.$store.getters['user/getUserId']
    },
    tab () {
      return this.$store.state.application.activeTab
    },
  },
  created () {
    this.$store.dispatch('application/updateActiveTab', 0)
  },
  destroyed () {
    this.$store.dispatch('application/updateActiveTab', null)
  }
}
</script>

<style scoped>

</style>

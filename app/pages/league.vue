<template v-slot:extension>

  <v-tabs-items
      v-model="tab"
      v-if=league
      touchless
  >
    <v-tab-item>
      <league-standings
          :league-info=league
      />
    </v-tab-item>
    <v-tab-item>
      <v-card>
        <v-card-text>
          Match up tab!
        </v-card-text>
      </v-card>
    </v-tab-item>
    <v-tab-item>
      <v-card>
        <v-card-text>
          <contracts-overview
            :league-info=league
          ></contracts-overview>
        </v-card-text>
      </v-card>
    </v-tab-item>
    <v-tab-item>
      <v-card>
        <v-card-text>
          Trade Center
        </v-card-text>
      </v-card>
    </v-tab-item>

  </v-tabs-items>
  <div v-else class="text-center">
    <h1>
      Please select a league
    </h1>
  </div>
</template>

<script>
import ContractsOverview from "@/components/league/ContractsOverview";

export default {
  name: "league-home.vue",
  components: {ContractsOverview},
  middleware: 'auth',
  data: function () {
    return {
    }
  },
  computed: {
    league() {
      return this.$store.getters["application/getActiveLeague"];
    },
    tab() {
      return this.$store.state.application.activeTab;
    },
    submenu(){
      return this.$store.state.application.submenu;
    }
  },
  created() {
    this.$store.dispatch("application/updateSubmenu", [
      'Standings', 'Match up', 'Contracts Overview', 'Trade Center'
    ]);
  },
  destroyed() {
    this.$store.dispatch("application/updateSubmenu", null);
    this.$store.dispatch("application/updateActiveTab", null);
  }
}
</script>

<style scoped>

</style>

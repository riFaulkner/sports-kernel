<template v-slot:extension>

  <v-tabs-items v-model="tab" v-if=league>
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
    <v-tab-item >
      <v-card>
        <v-card-text>
          Contracts Overview
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
export default {
  name: "league-home.vue",
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
  }
}
</script>

<style scoped>

</style>

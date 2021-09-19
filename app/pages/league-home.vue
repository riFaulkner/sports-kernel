<template v-slot:extension>

  <v-tabs-items v-model="tab" v-if=league>

    <v-tab-item>
      <league-overview
          :league-info=league
      />
    </v-tab-item>
    <v-tab-item>
      <v-card>
        <v-card-text>
          Hello! this is the Teams tab!
        </v-card-text>
      </v-card>
    </v-tab-item>
    <v-tab-item >
      <v-card>
        <v-card-text>
          Hello! this is the match up tab!
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
      text: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.',
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
      'League Home', 'Teams', 'Match up',
    ]);
  },
  destroyed() {
    this.$store.dispatch("application/updateSubmenu", null);
  }
}
</script>

<style scoped>

</style>

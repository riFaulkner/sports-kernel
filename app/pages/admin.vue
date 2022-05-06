<template v-slot:extension>
  <v-tabs-items
      v-model="tab"
      :touchless="true"
  >
    <v-tab-item>
      <h1>Manage leagues</h1>
      {{$auth.user}}
    </v-tab-item>
    <v-tab-item>
      <h1>Contracts</h1>
    </v-tab-item>
    <v-tab-item>
      <h1>Trades</h1>
    </v-tab-item>

  </v-tabs-items>
</template>

<script>
export default {
  name: "admin.vue",
  middleware: 'auth',
  computed: {
    tab() {
      return this.$store.state.application.activeTab;
    },
  },
  created() {
    this.$store.dispatch("application/updateActiveTab", 1);
    this.$store.dispatch("application/updateSubmenu", [
      'Manage Leagues', 'Contracts', 'Trades'
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
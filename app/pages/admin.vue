<template v-slot:extension>
  <v-tabs-items
      v-model="tab"
      :touchless="true"
  >
    <v-tab-item>
      <h1>Manage leagues</h1>
    </v-tab-item>
    <v-tab-item>
      <h1>Add player</h1>
    </v-tab-item>
    <v-tab-item>
      <h1>Contracts</h1>
    </v-tab-item>
    <v-tab-item>
      <h1>Trades</h1>
    </v-tab-item>
    <v-tab-item>
      <v-container fluid>
        <v-row>
          <v-col cols="4"/>
          <v-col cols="4">
            <v-card>
              <v-select
                  :items="dropRestructure.transactionTypes"
                  label="Transaction Type"
                  hint="Select to get started"
                  persistent-hint
                  v-on:change="transactionTypeSelected"
              />
              <v-form>

              </v-form>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-tab-item>
  </v-tabs-items>
</template>

<script>
export default {
  name: "admin.vue",
  middleware: 'auth',
  data: function() {
    return {
      dropRestructure: {
        transactionTypes: [
          {text: "Drop", value: "drop"},
          {text: "Restructure", value: "restructure"}
        ],
        drop: {

        },
        restructure: {

        }
      }

    }
  },
  computed: {
    tab() {
      return this.$store.state.application.activeTab;
    },
  },
  methods: {
    transactionTypeSelected(event) {
      alert(event)
    }
  },
  created() {
    this.$store.dispatch("application/updateActiveTab", 1);
    this.$store.dispatch("application/updateSubmenu", [
      'Manage League', 'Add Player', 'Create Contract', 'Make Trade', 'Restructure or Drop'
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
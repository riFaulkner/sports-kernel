<template v-slot:extension>
  <v-tabs-items
      v-model="tab"
      :touchless="true"
  >
    <v-tab-item>
      <v-col md="4" offset-md="4">
        <add-player/>
      </v-col>
    </v-tab-item>
  </v-tabs-items>
</template>

<script>
import AddPlayer from "@/components/players/AddPlayer.vue"
export default {
  name: "admin.vue",
  components: {AddPlayer},
  middleware: 'auth',
  data: function() {
    return {
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
    this.$store.dispatch("application/updateActiveTab", 0);
  },
  destroyed() {
    this.$store.dispatch("application/updateActiveTab", null);
  }
}
</script>

<style scoped>

</style>
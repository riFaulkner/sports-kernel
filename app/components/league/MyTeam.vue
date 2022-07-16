<template>
  <v-card>
    <v-card-text>
      My Team

    </v-card-text>
  </v-card>
</template>

<script>
export default {
  name: "MyTeam.vue",
  props: {
    leagueId: {
      type: String,
      required: true
    },
    teamId: {
      type: String,
    }
  },
  data: function() {
    return {

    }
  }
}
</script>

<style scoped>

</style>
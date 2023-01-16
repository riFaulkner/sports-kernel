<template>
  <div>
    <h1>Trades</h1>
    {{initiatingTeamId}}
    {{secondaryTeamId}}
  </div>
</template>

<script>
export default {
  name: "Trades",
  props: {
    initiatingTeamId: {
      type: String,
    },
    secondaryTeamId: {
      type: String
    }
  }
}
</script>

<style scoped>

</style>
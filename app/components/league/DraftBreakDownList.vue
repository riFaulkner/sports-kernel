<template>
  <div>
    <div
      v-for="pick in picks"
      :key="pick.pickValue"
    >
      Pick: {{pick.pickValue !=null ? pick.pickValue : "TBD" }}
    </div>
  </div>

</template>

<script>
export default {
  name: "DraftBreakDownList",
  props: {
    picks: {
      type: Array,
      default: function() {
        return []
      }
    }
  }
}
</script>

<style scoped>

</style>
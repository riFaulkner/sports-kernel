<template>
  <div class="text-center">
    <h2> Dead cap </h2>
    <v-data-table
        v-if="deadCap"
        :headers="headers"
        :items=deadCap
        hide-default-footer
    >
      <template v-slot:item.year1DeadCap="{item}">
        {{ getAmount(item.deadCapYears, 0 )}}
      </template>
      <template v-slot:item.year2DeadCap="{item}">
        {{ getAmount(item.deadCapYears, 1 )}}
      </template>

    </v-data-table>
  </div>
</template>

<script>
export default {
  name: "TeamDeadCapBreakdown",
  props: {
    deadCap: {
      type: Array,
      default: function () {
        return []
      }
    }
  },
  computed: {
    headers() {
      const nextDeadCapYear = new Date().getFullYear()
      if (this.deadCap?.length > 0) {
        this.deadCap[0].year
      }

      return [
        {text: "Player", value: "deadCapNote", align: "center"},
        {text: nextDeadCapYear, value: "year1DeadCap", align: "center"},
        {text: nextDeadCapYear + 1, value: "year2DeadCap", align: "center"}
      ]
    },
  },
  methods: {
    getAmount(deadCapMounts, index) {
      if (deadCapMounts?.length <= index) {
        return ""
      }
      const yearValue = deadCapMounts[index]?.amount

      return "$" + yearValue.toLocaleString()
    }
  }
}
</script>

<style scoped>

</style>
<template>
  <v-card>

  </v-card>
</template>

<script>
export default {
  name: "ContractCard.vue",
  data () {
    return {
      maxIndividualCapHitTarget: 50,
      oneYearContractInfo: {
        capHit: 20,
        pctCapHit: 10
      },
      twoYearContractInfo: {
        capHit: 18,
        pctCapHit: 8
      },
      threeYearContractInfo: {
        pctCapHit: 6
      },
      fourYearContractInfo: {
        pctCapHit: 5
      }
    }
  },
  methods: {
    getCapHitColor (capHit) {
      if (capHit > this.maxIndividualCapHitTarget) {
        return "error"
      }
      if (capHit > this.maxIndividualCapHitTarget * 0.7) {
        return "warn"
      }
      return "success"
    }
  }
}
</script>

<style scoped>

</style>

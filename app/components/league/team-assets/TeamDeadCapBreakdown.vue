<template>
  <div>
    <h2> Dead cap </h2>
    <v-data-table
        :headers="headers"
        :items="processedDeadCap"
    >
      <template v-slot:item.year1DeadCap="{item}">
        {{ getAmount(item.amounts, 0 )}}
      </template>
      <template v-slot:item.year2DeadCap="{item}">
        {{ getAmount(item.amounts, 1 )}}
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
  data: function () {
    return {

    }
  },
  computed: {
    processedDeadCap() {
      if (this.deadCap?.length === 0) {
        return []
      }
      let tableView = []

      // For each year of dead cap
      this.deadCap.forEach((year) => {
        // Search through all the dead cap contracts and add them to the table view
        year.deadCapAccrued.forEach((contract) => {
          const filterResult = tableView.filter((item) => {
            return item.id === contract.associatedContractId
          })
          let tableItem = {
            id: contract.associatedContractId,
            name: contract.contract.player.playerName,
            amounts: []
          }
          if(filterResult.length === 1) {
            tableItem = filterResult[0]
          } else {
            tableView.push(tableItem)
          }
          tableItem.amounts.push(contract.amount)
        })
      })

      return tableView
    },
    headers() {
      const nextDeadCapYear = this.deadCap[0].year

      return [
        {text: "Player", value: "name"},
        {text: nextDeadCapYear, value: "year1DeadCap"},
        {text: nextDeadCapYear + 1, value: "year2DeadCap"}
      ]
    },
  },
  methods: {
    getAmount(contractAmounts, index) {
      if (contractAmounts.length <= index) {
        return ""
      }
      const yearValue = contractAmounts[index]

      return "$" + yearValue.toLocaleString()
    }
  }
}
</script>

<style scoped>

</style>
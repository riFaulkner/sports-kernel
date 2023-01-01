<template>
  <div>
    <v-data-table
        :headers="headers"
        :items="contracts"
        item-key="id"
        :search="search"
        :custom-filter="playerNameMatchesSearch"
        :loading="loading"
        :value="selectedList"
        show-select
        :single-select="true"
    >
      <template v-slot:top>
        <v-text-field
            v-model="search"
            label="Player Name"
            class="mx-4"
        ></v-text-field>
      </template>
      <template v-slot:item.data-table-select="{item, isSelected, select}">
        <v-btn v-if="isSelected" @click="contractDeselected(item, select)" class="primary">Selected</v-btn>
        <v-btn v-else @click="contractSelected(item,select)">&nbsp Select &nbsp</v-btn>
      </template>
      <template v-slot:item.totalContractValue="{item}">
        ${{ item.totalContractValue.toLocaleString() }}
      </template>
      <template v-slot:item.totalGuaranteedValue="{item}">
        ${{ totalGuaranteed(item) }}
      </template>
      <template v-slot:item.year1="{item}">
        <div>
          {{ getContractYearDetails(item, 0) }}
        </div>
      </template>
      <template v-slot:item.year2="{item}">
          {{ getContractYearDetails(item, 1) }}
      </template>
      <template v-slot:item.year3="{item}">
          {{ getContractYearDetails(item, 2) }}
      </template>
      <template v-slot:item.year4="{item}">
          {{ getContractYearDetails(item, 3) }}
      </template>
    </v-data-table>
  </div>
</template>

<script>

export default {
  name: "ContractSearch",
  props: {
    contracts: {
      type: Array,
      required: true,
    },
    currentSeason: {
      type: Number,
      required: true
    },
    leagueId: {
      type: String,
      required: true
    },
    loading: {
      type: Boolean,
      required: false,
      default: false
    },
    selected: {
      type: Array,
      default: function () {
        return []
      }
    }
  },
  data: function () {
    return {
      search: "",
      queriedWith: "",
    }
  },
  methods: {
    getContractYearDetails(contract, year) {
      const offset = contract.currentYear + year
      if (contract.contractDetails.length < offset) {
        return ""
      }
      return `$${contract.contractDetails[offset - 1]?.totalAmount.toLocaleString()}`
    },
    getContractYearPaid(contract, year) {
      const currentYear = contract.currentYear

      return currentYear > year
    },
    playerNameMatchesSearch(value, search, item) {
      if (item === undefined || item === null) {
        return false;
      }
      return item.player.playerName.toLowerCase().includes(search.toLowerCase());
    },
    contractSelected(contract, select) {
      select(true)
      this.$emit('contract-selected', {contract: contract})
    },
    contractDeselected(contract, select) {
      select(false)
      this.$emit('contract-deselected', {contract: contract})
    },
    totalGuaranteed(contract) {
      return contract.contractDetails.reduce((incompleteSum, year) => incompleteSum + year.guaranteedAmount, 0)
          .toLocaleString()
    }
  },
  computed: {
    headers() {
      return [
        {text: "Player Name:", value: "player.playerName"},
        {text: "Position", value: "player.position"},
        {text: "NFL Team", value: "player.team"},
        {text: "Total Contract Value", value: "totalContractValue"},
        {text: "Total Guaranteed", value: "totalGuaranteedValue"},
        {text: this.currentSeason, value: "year1"},
        {text: this.currentSeason+1, value: "year2"},
        {text: this.currentSeason+2, value: "year3"},
        {text: this.currentSeason+3, value: "year4"},
        {text: "Restructure Status", value: "restructureStatus", align: 'center'},
      ]
    },
    selectedList: {
      // No set method, instead the event is handled in the contractSelected method above
      get:function () {
        return this.selected
      },
    }
  }
}
</script>

<style scoped>

</style>
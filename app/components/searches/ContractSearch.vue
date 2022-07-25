<template>
  <div>
    <v-data-table
        :headers="headers"
        :items="contracts"
        item-key="id"
        :search="search"
        :custom-filter="playerNameMatchesSearch"
        :loading="loading"
        v-model="selected"
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
        <div v-if=getContractYearPaid(item,1) class="success--text">
          {{ getContractYearDetails(item, 1) }}
        </div>
        <div v-else class="warning--text">
          {{ getContractYearDetails(item, 1) }}
        </div>
      </template>
      <template v-slot:item.year2="{item}">
        <div v-if=getContractYearPaid(item,2) class="success--text">
          {{ getContractYearDetails(item, 2) }}
        </div>
        <div v-else class="warning--text">
          {{ getContractYearDetails(item, 2) }}
        </div>
      </template>
      <template v-slot:item.year3="{item}">
        <div v-if=getContractYearPaid(item,3) class="success--text">
          {{ getContractYearDetails(item, 3) }}
        </div>
        <div v-else class="warning--text">
          {{ getContractYearDetails(item, 3) }}
        </div>
      </template>
      <template v-slot:item.year4="{item}">
        <div v-if=getContractYearPaid(item,4) class="success--text">
          {{ getContractYearDetails(item, 4) }}
        </div>
        <div v-else class="warning--text">
          {{ getContractYearDetails(item, 4) }}
        </div>
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
    leagueId: {
      type: String,
      required: true
    },
    loading: {
      type: Boolean,
      required: false,
      default: false
    }
  },
  data: function () {
    return {
      search: "",
      headers: [
        {text: "Player Name:", value: "player.playerName"},
        {text: "Player Position", value: "player.position"},
        {text: "NFL Team", value: "player.team"},
        {text: "Total Contract Value", value: "totalContractValue"},
        {text: "Total Guaranteed", value: "totalGuaranteedValue"},
        {text: "Year 1", value: "year1"},
        {text: "Year 2", value: "year2"},
        {text: "Year 3", value: "year3"},
        {text: "Year 4", value: "year4"},
        {text: "Contract Restructure Status", value: "restructureStatus", align: 'center'},
      ],
      queriedWith: "",
      selected: [],
    }
  },
  methods: {
    getContractYearDetails(contract, year) {
      if (contract.contractDetails.length < year) {
        return ""
      }
      return `$${contract.contractDetails[year - 1]?.totalAmount.toLocaleString()}`
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

}
</script>

<style scoped>

</style>
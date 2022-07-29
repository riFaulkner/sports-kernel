<template>
  <v-card>
    <v-toolbar
        color="blue lighten-2"
        dark
    >
      <v-spacer></v-spacer>
      <v-toolbar-title class="align-center">{{ title }}</v-toolbar-title>
      <v-spacer></v-spacer>
    </v-toolbar>
    <v-list v-if="contracts.length > 0">
      <v-list-group
          v-for="contract in contracts"
          :key="contract.id"
          v-model="contract.active"
          no-action
      >
        <template v-slot:activator>
          <v-list-item-content>
            <v-list-item-title>{{ contract.player.playerName }}
              <v-spacer/>
              ${{ contract.totalContractValue.toLocaleString() }}
            </v-list-item-title>
          </v-list-item-content>
        </template>
        <v-card>
        </v-card>

        <v-list-item
            v-for="year in contract.contractDetails"
            :key="year.year"
            style="padding-left: 2em"
        >
          <v-list-item-content>
            <v-list-item-title>Year {{ year.year }}: ${{ year.totalAmount.toLocaleString() }}
            </v-list-item-title>
          </v-list-item-content>
          <v-list-item-icon v-if="contract.currentYear === year.year" >
            <v-icon>mdi-star-outline</v-icon>
          </v-list-item-icon>
        </v-list-item>
      </v-list-group>
    </v-list>
    <div v-else class="text-center">
      no contracts for this position
    </div>

  </v-card>


</template>

<script>
export default {
  name: "TeamContractsBreakdown",
  props: {
    title: {
      type: String,
      required: true,
    },
    contracts: {
      type: Array,
      default: () => []
    }
  },
  data: function () {
    return {
      active: []
    }
  }
}
</script>

<style scoped>

</style>
<template>
  <v-card
      color="grey darken-3"
      max-height="750px"
      class="overflow-y-auto"
  >
    <v-card-title> {{team.teamName}}</v-card-title>
    <v-tabs v-model="activeTab">
      <v-tab
          v-for="tab in teamTabs"
          :key="tab.key"
          :href="'#'+tab.key"
          grow
      >
        {{ tab.text }}
      </v-tab>
    </v-tabs>
    <v-tabs-items v-model="activeTab">
      <br/>
      <v-tab-item
          :value="'contracts'"
      >
        <v-checkbox
            v-for="contract in team.activeContracts"
            v-model="value.contracts"
            :key=contract.id
            :label="contract.player.playerName"
            :value="{name: contract.player.playerName, id: contract.id}"
        />
      </v-tab-item>
      <v-tab-item
          :value="'picks'"
      >
        <v-checkbox
            v-for="pick in flattenPicks(team.teamAssets.draftPicks)"
            v-model="value.picks"
            :label="draftPickLabel(pick)"
            :value="`${pick.year}-${pick.round}`"
        />
      </v-tab-item>
    </v-tabs-items>
  </v-card>
</template>

<script>
export default {
  name: "TeamAssetSelector",
  props: {
    team: {
      type: Object,
      required: true
    },
    value: {
      type: Object,
      default: function () {
        return {
          teamId: "",
          contracts: [],
          picks: []
        }
      }
    }
  },
  data: function() {
    return {
      activeTab: "contracts",
      teamTabs: [
        {text: "Contracts", key: "contracts"},
        {text: "Draft Picks", key: "picks"}
      ],
    }
  },
  methods: {
    flattenPicks(draftPicks) {
      const flattenedPicks = []

      draftPicks.forEach(year => {
        year.picks.forEach(pick => {
          const position = pick.value ? pick.value : "NA"
          flattenedPicks.push({year: year.year, originalOwner: pick.originalOwner, round: pick.round, position: position})
        })
      })
      return flattenedPicks
    },
    draftPickLabel(pick) {
      const owner = pick.originalOwner ? `-- Original Owner: ${pick.originalOwner}` : ""
      return `${pick.year} ${pick.round}.${pick.position} ${owner}`
    }
  },
}
</script>

<style scoped>

</style>
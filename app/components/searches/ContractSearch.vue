<template>
  <div>
    <v-data-table
        :headers="headers"
        :items="contracts"
        item-key="playerName"
        :search="search"
        :custom-filter="playerNameMatchesSearch"
    >
      <template v-slot:top>
        <v-text-field
            v-model="search"
            label="Player Name"
            class="mx-4"
        ></v-text-field>
      </template>
    </v-data-table>
  </div>
</template>

<script>
export default {
  name: "ContractSearch",
  props: {
    leagueId: {
      type: String,
      required: true
    }
  },
  data: function () {
    return {
      contracts: [
        {playerName: "player 1", contractValue: 123},
        {playerName: "player 2", contractValue: 123},
      ],
      search: "",
      headers: [
        {text: "Player Name:", value: "playerName"},
        {text: "Contract Value", value: "contractValue"}
      ],
      queriedWith: ""
    }
  },
  methods: {
    playerNameMatchesSearch(value, search, item) {
      if (item === undefined || item === null) {
        return false;
      }
      return item.playerName.includes(search);
    }
  },
  apollo: {
    contracts: {
      query: ,
      variables() {
        return {leagueId: this.leagueId}
      }
    }
  }
}
</script>

<style scoped>

</style>
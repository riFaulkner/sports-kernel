<template>
  <div>
    <v-data-table
        :headers="headers"
        :items="players"
        :search="search"
        :loading="$apollo.loading"
        v-model="selected"
        show-select
        single-select
        @item-selected="playerSelected"
    >
      <template v-slot:top>
        <v-row>
          &nbsp;
          <v-col cols="4">
            <v-select
                v-model="selectedPosition"
                label="Select Player Position"
                :items="playerPositions"
                @change="positionSelected"
            />
          </v-col>
          <v-col>
            <v-text-field
                v-model="search"
                label="Player Name"
            ></v-text-field>
          </v-col>
          &nbsp;
        </v-row>

      </template>

    </v-data-table>
  </div>

</template>

<script>
import {PLAYER_BY_POSITION} from "~/graphql/queries/players/playersGQL";

export default {
  name: "PlayerSearch",
  data: function () {
    return {
      players: [],
      headers: [
        {text: "Player Name", value: "playerName"},
        {text: "Position", value: "position"},
        {text: "Team", value: "team"},
        {text: "Age", value: "age"}
      ],
      search: "",
      selected: [],
      selectedPosition: ""
    }
  },
  computed: {
    playerPositions() {
      return this.$store.getters["globalHelpers/getPlayerPositionTypes"]
    }
  },
  methods: {
    positionSelected() {
      this.$apollo.query({
        query: PLAYER_BY_POSITION,
        variables: {position: this.selectedPosition},
      }).then((data) => {
        this.players = data.data.playersByPosition
      }).catch((error) => {
        this.$store.dispatch("application/alertError", {message: "Error occurred fetching players"})
        console.error(error)
      })
    },
    playerSelected(event) {
      if (event.value) {
        this.$emit("player-selected", {player: event.item})
      }
    }
  }
}
</script>

<style scoped>

</style>
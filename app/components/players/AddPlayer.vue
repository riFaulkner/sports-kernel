<template>
  <v-card
      color="grey darken-3"
      elevation="2"
  >
    <v-card-title>
      Add a Player
    </v-card-title>
    <v-card-subtitle>
      Adding a player will make them available for contract in all leagues.
    </v-card-subtitle>
    <v-card-text>
      <v-form
          v-model="isFormValid"
      >
        <v-text-field
            label="Player Name"
            v-model:value="playerData.playerName"
            :rules="rules.playerName"
        />
        <v-select
            label="Position"
            :items="positionTypes"
            v-model:value="playerData.position"
            :rules="rules.position"
        />
        <v-select
            label="Team"
            v-model:value="playerData.teamNFL"
            :items="teams"
            item-text="longName"
            item-value="abbreviation"
            :rules="rules.teamNFL"
        >
          <template
              v-slot:item="{item}"
          >
            <v-avatar
                size="45"
                tile
            >
              <img
                  :src=item.icon
                  alt="logo"
              />
            </v-avatar> &nbsp; &nbsp; {{ item.longName }}
          </template>
          <template
              v-slot:selection="{item}"
          >
            <v-avatar
                size="45"
                tile
            >
              <img
                  :src=item.icon
                  alt="logo"
              />
            </v-avatar> &nbsp; &nbsp; {{ item.longName }}
          </template>
        </v-select>
        <v-text-field
            label="Avatar URL"
            v-model:value="playerData.avatar"
        />
        <br/>
        <v-row justify="center">
          <v-date-picker justify="center"
                         header-color="primary"
                         elevation="9"
                         v-model="playerData.birthday"
          />
        </v-row>
        <br/>
        <v-text-field
            label="Birthday"
            disabled
            v-model:value="playerData.birthday"
        />

      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-spacer/>
      <v-btn
          @click="resetForm"
          color="primary"
          text
      >
        Reset
      </v-btn>
      <v-btn
          @click="submitForm"
          color="primary"
          :disabled="!isFormValid"
      >
        Submit
      </v-btn>
    </v-card-actions>

  </v-card>

</template>

<script>
import {CREATE_PLAYER} from "~/graphql/queries/nfl/nflTeamsGraphQL";

export default {
  name: "AddPlayer.vue",
  data: function () {
    return {
      isFormValid: "",
      playerData: {
        avatar: "",
        birthday: "",
        playerName: "",
        position: "",
        teamNFL: "",
      },
      rules: {
        playerName: [
          value => !!value || 'Player name required.',
        ],
        position: [
          value => !!value || 'Position required.',
        ],
        teamNFL: [
          value => !!value || 'Team required.',
        ]
      }
    }
  },
  computed: {
    positionTypes() {
      return this.$store.getters["globalHelpers/getPlayerPositionTypes"]
    },
    teams() {
      return this.$store.getters["globalHelpers/getNFLTeams"]
    }
  },
  methods: {
    resetForm() {
      this.playerData = {
        avatar: "",
        birthday: "",
        playerName: "",
        position: "",
        teamNFL: "",
      }
    },
    submitForm() {
      this.$apollo.mutate({
        mutation: CREATE_PLAYER,
        variables:  {
          player: this.playerData
        }
      }).then(() => {
        this.$store.dispatch("application/alertSuccess", {
          message: "Player successfully created."
        })
        this.resetForm()
      }).catch((error) => {
        this.$store.dispatch("application/alertError", {
          message: "Failed to create new player."
        })
        console.error("Failed to create new player: ", error)
      })
    }
  }
}
</script>

<style scoped>

</style>
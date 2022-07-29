<template>
  <v-col align-self="center" md=4 offset-md=4>
    <v-card class="mb-12" color="grey darken-3">
      <h1 class="text-center">Generate Access Codes</h1>
      <v-card color="grey darken-3" elevation=0>
        <v-col>
          <v-select
                v-model=selectedTeam
                :items="teams"
                item-text="teamName"
                item-value="id"
                :loading="$apollo.loading"
                label=" Select a team "
                @change="isTeamSelected=true"
              >
         </v-select>
         <v-select
                v-model=selectedRole
                :items="roleTypes"
                item-text="text"
                item-value="value"
                label=" Select a Role "
                @change="isRoleSelected=true"
              >
          </v-select>
        </v-col>
      </v-card>
    </v-card>
      <v-row justify="center">
      <v-dialog
        v-model="dialog"
        persistent
        max-width="500"
        color="black lighten-2"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-btn
            color="primary"
            dark
            :loading="loading"
            v-bind="attrs"
            v-on="on"
            :disabled="!(isRoleSelected && isTeamSelected)"
            @click="submitAccessCodeRequest"
          >
            Submit
          </v-btn>
        </template>
        <v-skeleton-loader :loading="loading" type="article, actions">
          <v-card>
            <v-card-title class="text-h5">
              Access Code Created!
            </v-card-title>
            <v-card-text>{{accessCode}}</v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn
                color="green darken-1"
                text
                @click="copy"
              >
                Copy Code
              </v-btn>
              <v-btn
                color="red darken-1"
                text
                @click="dialog = false"
              >
                Exit
              </v-btn>
            </v-card-actions>
            <v-overlay
              :absolute="absolute"
              :value="overlay"
            >
              <v-btn
                color="success"
                rounded
                @click="overlayConfirm"
              >
                Copied to Clipboard!
              </v-btn>
            </v-overlay>
          </v-card>
        </v-skeleton-loader>
      </v-dialog>
    </v-row>
  </v-col>
</template>

<script>
import {LEAGUE_TEAMS_BY_LEAGUE_ID} from "~/graphql/queries/league/leagueGraphQL";
import {GENERATE_ACCESS_CODE} from "~/graphql/queries/team/teamGraphQL";

export default {
  name: "GenerateAccessCodes",
  props: {
    leagueId: {
      type: String,
      required: true
    }
  },
  data: function () {
    return {
      teams: [],
      roleTypes: [
        {text: "League Manager", value:"LEAGUE_MANAGER"},
        {text: "Team Owner", value:"TEAM_OWNER"}
      ],
      selectedTeam: {},
      selectedRole: "",
      accessCode: "",
      dialog: false,
      loading: false,
      isTeamSelected: false,
      isRoleSelected: false,
      overlay: false,
      absolute: true
    }
  },
  methods: {
    submitAccessCodeRequest(){
      this.loading = true
      const response = this.$apollo.mutate({
        mutation: GENERATE_ACCESS_CODE,
        variables: {
          leagueId: this.leagueId,
          teamId: this.selectedTeam,
          role: this.selectedRole
        },
      }).then(result => {
          this.accessCode = result.data.generateAccessCode
          this.loading = false
      });
    },
    copy(){
      navigator.clipboard.writeText(this.accessCode)
      //this.$refs.clone.focus()
      //document.execCommand('copy')
      //this.dialog = false
      this.overlay = true
    },
    overlayConfirm(){
      this.overlay = false
      this.dialog = false
    }
  },
  apollo: {
    teams: {
      query: LEAGUE_TEAMS_BY_LEAGUE_ID,
      variables() {
        return {leagueId: this.leagueId}
      }
    }
  },
}
</script>

<style scoped>

</style>
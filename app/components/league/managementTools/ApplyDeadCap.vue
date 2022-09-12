<template>
  <v-container>
    <v-row>
      <v-col>
        <v-select
            v-model=selectedTeam
            :items="teams"
            item-text="teamName"
            item-value="id"
            :loading="$apollo.loading"
            label="Select a team"
            @input="teamSelected"
        />
      </v-col>
    </v-row>
    <template v-if="selectedTeam">
      <v-row>
        <v-col>
          <v-skeleton-loader :loading="this.loadingSelectedTeam" type="table-heading, table-tbody">
            <team-dead-cap-breakdown :dead-cap="this.deadCap"/>
          </v-skeleton-loader>
        </v-col>
      </v-row>
      <v-divider/>
      <v-form
          v-model="isFormValid"
      >
        <v-row>
          <v-col>
            <v-text-field
                v-model:value="yearOneValue"
                type="number"
                :label=yearOneLabel
                :rules="rules.firstYearRules"
            />
          </v-col>
          <v-col>
            <v-text-field
                v-model:value="yearTwoValue"
                type="number"
                :label=yearTwoLabel
                hint="Leave blank if not applicable"
                persistent-hint
            />
          </v-col>
          <v-col cols="12">
            <v-text-field
                v-model:value="deadCapNote"
                label="Player Name/Note"
                hint="This is what will show up when viewing dead cap for the team"
                persistent-hint
            />
          </v-col>
          <v-col offset="11">

            <v-btn
                color="primary"
                text
                elevation="6"
                :disabled="!isFormValid"
                @click="addDeadCapToTeam"
            >
              Save
            </v-btn>
          </v-col>
        </v-row>
      </v-form>

    </template>
  </v-container>
</template>

<script>
import {LEAGUE_TEAMS_BY_LEAGUE_ID} from "@/graphql/queries/league/leagueGraphQL";
import TeamDeadCapBreakdown from "@/components/league/team-assets/TeamDeadCapBreakdown";
import {APPLY_DEAD_CAP, TEAM_DEAD_CAP} from "@/graphql/queries/team/teamGraphQL";

export default {
  name: "ApplyDeadCap",
  components: {TeamDeadCapBreakdown},
  props: {
    leagueId: {
      type: String,
      required: true
    },
  },
  data: function () {
    return {
      currentYear: new Date().getFullYear(),
      deadCapNote: null,
      selectedTeam: null,
      loadingSelectedTeam: false,
      teams: [],
      deadCap: [],
      isFormValid: false,
      rules: {
        firstYearRules: [
            value => !!value || "First year is required",
            value => 0 < value || "First year must be greater than 0",
        ],
        noteRules: [
            value => !!value || "Required, add player name or dead cap note",
            value => 0 < value.length || "Required, add player name or dead cap note"
        ],
      },
      yearOneValue: null,
      yearTwoValue: null,
    }
  },
  computed: {
    yearOneLabel() {
      return `${this.currentYear} Dead Cap amount`
    },
    yearTwoLabel() {
      return `${this.currentYear + 1} Dead Cap amount`
    }
  },
  methods: {
    addDeadCapToTeam(){
      const deadCapYears = []
      deadCapYears.push({year: this.currentYear, amount: this.yearOneValue})
      if (!!this.yearTwoValue && this.yearTwoValue > 0) {
        deadCapYears.push({year: this.currentYear +1 , amount: this.yearTwoValue})
      }
      const deadCapToAdd = {
        deadCapNote: this.deadCapNote,
        deadCapYears: deadCapYears
      }

      console.log("Dead cap to add: ", deadCapToAdd)

      // TODO: add apollo call to make the year
      this.$apollo.mutate({
        mutation: APPLY_DEAD_CAP,
        variables: {
            leagueId: this.leagueId,
            teamId: this.selectedTeam,
            input: deadCapToAdd
          },
      }).then(data => {
        console.log("Data", data)
      }).catch(err=> {
        this.$store.dispatch("application/alertError", {message: "Error adding dead cap"})
        console.log("error", err)
      })
    },
    teamSelected(teamId) {
      this.loadingSelectedTeam = true
      this.$apollo.query({
        query: TEAM_DEAD_CAP,
        variables: {
          leagueId: this.leagueId,
          teamId: teamId
        }
      }).then((data) => {
        this.deadCap = data.data.teamById.teamLiabilities.deadCap
      })
          .catch((error) => {
            console.error("Error getting team dead cap for team: ", error)
            this.$store.dispatch("application/alertError", {message: "Failed to get team's current dead cap"})
          }).finally(() => {
        this.loadingSelectedTeam = false
      })
    }
  },
  apollo: {
    teams: {
      query: LEAGUE_TEAMS_BY_LEAGUE_ID,
      variables() {
        return {leagueId: this.leagueId}
      }
    },
  },
}
</script>

<style scoped>

</style>
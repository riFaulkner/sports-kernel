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
      <v-form>
        <v-row>
          <v-col>
            <v-text-field
                type="number"
                label="Dead Cap amount year 1"
            />
          </v-col>
          <v-col>
            <v-text-field
                type="number"
                label="Dead Cap amount year 2"
            />
          </v-col>
          <v-col cols="12">
            <v-text-field
                label="Player Name/Note"
                hint="This is what will show up when viewing deadcap for the team"
                persistent-hint
            />
          </v-col>
          <v-col offset="11">

            <v-btn
                color="primary"
                text
                elevation="6"
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
import {TEAM_DEAD_CAP} from "@/graphql/queries/team/teamGraphQL";

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
      selectedTeam: null,
      loadingSelectedTeam: false,
      teams: [],
      deadCap: [],
    }
  },
  methods: {
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
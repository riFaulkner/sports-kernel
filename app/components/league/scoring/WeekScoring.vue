<template>
  <v-card>
    <v-container>
      <v-row align="center" justify="center" align-sm="left" justify-sm="left">
        <v-col sm="4" lg="4" cols="4" offset-sm="0">
          <v-select
              v-model="selectedSeason"
              :items="seasons"
              label="Select a season"
          />
        </v-col>
        <v-col sm="8" lg="6" cols="8">
          <v-select
              v-model="selectedWeek"
              :items="weeks"
              label="Select a week" />
        </v-col>
        <v-col sm="7" lg="8" cols="12" offset-sm="1" offset-md="0">
          <v-select
              v-model="selectedMatchUpNumber"
              :items="matchUps"
              :item-text="item => `${item.awayTeam} vs ${item.homeTeam}`"
              item-value="matchUpNumber"
              label="Select a match up"
              :loading="$apollo.queries.matchUps.loading"
              @change="teamSelected"
          />
        </v-col>
        <v-col md="2" offset-sm="0" offset="6" v-if="selectedMatchUpNumber != null">
          <v-btn
              color="primary"
              @click="refreshView"
              :disabled="!canRefresh"
              :loading="isMatchUpRefresh"
          > Refresh Stats
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
    <v-skeleton-loader
        v-if="selectedMatchUpNumber !== null"
        :loading="isInitialMatchUpLoading"
        type="card"
    >
      <v-container v-if="selectedMatchUpNumber !== null">
        <v-row justify="center">
          <v-col
              v-for="team in scoringData"
              cols="12"
              sm="5"
              :key="team.teamName"
          >
            <v-toolbar
                color="blue"
                dark
                class="rounded-t-lg"
            >
              <v-toolbar-title>{{ team.teamName }}</v-toolbar-title>

              <v-spacer></v-spacer>
              <b>Pts: {{ team.totalPoints.toLocaleString() }}</b>
            </v-toolbar>
            <v-list outlined class="rounded-b-lg">
              <v-list-item v-for="qb in team.lineUp.qb" :key="qb.playerName" >
                QB
                <v-spacer/>
                {{ qb.playerName }} -- {{ qb.points }}
              </v-list-item>
              <v-list-item v-for="rb in team.lineUp.rb" :key="rb.playerName">
                RB
                <v-spacer/>
                {{ rb.playerName }} -- {{ rb.points }}
              </v-list-item>
              <v-list-item v-for="wr in team.lineUp.wr" :key="wr.playerName">
                WR
                <v-spacer/>
                {{ wr.playerName }} -- {{ wr.points }}
              </v-list-item>
              <v-list-item v-for="te in team.lineUp.te" :key="te.playerName">
                TE
                <v-spacer/>
                {{ te.playerName }} -- {{ te.points }}
              </v-list-item>
              <v-list-item v-for="flex in team.lineUp.flex" :key="flex.playerName">
                Flex
                <v-spacer/>
                {{ flex.playerName }} -- {{ flex.points }}
              </v-list-item>
              <v-list-item v-for="sFlex in team.lineUp.superFlex" :key="sFlex.playerName">
                SuperFlex
                <v-spacer/>
                {{ sFlex.playerName }} -- {{ sFlex.points }}
              </v-list-item>

            </v-list>
          </v-col>
        </v-row>
        <v-row justify="center">
          <v-col
              v-for="team in scoringData"
              cols="12"
              sm="5"
              :key="team.teamName"
          >
            <v-list
                outlined
                color="grey darken-3"
                class="rounded-lg"
            >
              <v-list-item v-if="$vuetify.breakpoint.mobile">
                Bench: &nbsp; {{ team.teamName }}
              </v-list-item>
              <v-list-item v-for="player in benchPlayers(team)"
                           :disabled="player.gamePlayed === 100"
                           :key="player.playerName"
              >
                <v-list-item-content>
                  <v-list-item-title>
                    {{ player.playerName }} -- Points: {{ player.points }}
                  </v-list-item-title>
                  <v-list-item-subtitle>
                    {{ player.position }} | {{ player.nflTeam }} | Projected: {{ player.projectedPoints }}
                  </v-list-item-subtitle>
                </v-list-item-content>
              </v-list-item>
            </v-list>
          </v-col>
        </v-row>
      </v-container>
    </v-skeleton-loader>
  </v-card>
</template>

<script>
import {WEEK_SCORING_FOR_MATCH_UP, WEEK_SCORING_MATCH_UPS} from "@/graphql/queries/scoring/scoringGraphQL";

export default {
  name: "WeekScoring",
  props: {
    leagueId: {
      type: String,
      required: true
    },
  },
  data: function () {
    return {
      canRefresh: true,
      isInitialFetch: true,
      matchUps: null,
      seasons: [
        {text: '2021', value: 2021},
        {text: '2022', value: 2022},
        {text: '2023', value: 2023},
        {text: '2024', value: 2024},
      ],
      selectedMatchUpNumber: null,
      selectedSeason: new Date().getFullYear(),
      selectedWeek: -1,
      scoringData: null,
      weeks: [
        {text: 'Current Week', value: -1},
        {text: 'Week 1', value: 1},
        {text: 'Week 2', value: 2},
        {text: 'Week 3', value: 3},
        {text: 'Week 4', value: 4},
        {text: 'Week 5', value: 5},
        {text: 'Week 6', value: 6},
        {text: 'Week 7', value: 7},
        {text: 'Week 8', value: 8},
        {text: 'Week 9', value: 9},
        {text: 'Week 10', value: 10},
        {text: 'Week 11', value: 11},
        {text: 'Week 12', value: 12},
        {text: 'Week 13', value: 13},
        {text: 'Week 14', value: 14},
        {text: 'Week 15', value: 15},
        {text: 'Week 16', value: 16},
      ]
    }
  },
  methods: {
    benchPlayers(team) {
      return team.roster.filter(player => !player.isInLineUp)
    },
    refreshView() {
      this.isInitialFetch = false
      this.$apollo.queries.scoringData.refetch()
      this.canRefresh = false

      setTimeout(() => this.canRefresh = true, 60000);
    },
    teamSelected() {
      this.isInitialFetch = true
      this.canRefresh = true
    },

  },
  computed: {
    isInitialMatchUpLoading() {
      if (this.isInitialFetch) {
        return this.$apollo.queries.scoringData.loading
      }
      return false
    },
    isMatchUpRefresh() {
      if (!this.isInitialFetch) {
        return this.$apollo.queries.scoringData.loading
      }
      return false
    }
  },
  apollo: {
    matchUps: {
      query: WEEK_SCORING_MATCH_UPS,
      variables() {
        return {
          leagueId: this.leagueId,
          season: this.selectedSeason,
          week: this.selectedWeek,
        }
      },
      update: data => data.scoring.weekMatchUps,
    },
    scoringData: {
      query: WEEK_SCORING_FOR_MATCH_UP,
      variables() {
        return {
          leagueId: this.leagueId,
          season: this.selectedSeason,
          week: this.selectedWeek,
          matchUpNumber: this.selectedMatchUpNumber
        }
      },
      update: data => {
        return data.scoring.matchUpScoring
      },
      skip() {
        return this.selectedMatchUpNumber === null
      },
    }
  }
}
</script>

<style scoped>
</style>

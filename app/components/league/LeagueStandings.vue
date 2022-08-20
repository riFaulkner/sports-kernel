<template>
  <v-card>
    <v-card-title>
      <v-row>
        <v-col cols="2">
          <v-select
              :items="leagueYears"
              v-model:value="selectedYear"
          />
        </v-col>
        <v-col offset="3">
          {{ leagueInfo.leagueName }} Standings
        </v-col>
      </v-row></v-card-title>
    <v-data-table
        :headers=headers
        :items=processedTeamData
        group-by="division"
        hide-default-footer
        :items-per-page="15"
        dense
    >
      <template v-slot:group.header="{items, isOpen, toggle}">
        <th colspan="100%">
          <v-icon @click="toggle" small
          >{{ isOpen ? 'mdi-minus' : 'mdi-plus' }}
          </v-icon>
          {{ items[0].division }}
        </th>
      </template>

      <template v-slot:item.record="{item}">
        {{ item.seasonSummary.wins }}-{{ item.seasonSummary.losses }}-{{ item.seasonSummary.ties }}
      </template>
      <template v-slot:item.winPct="{item}">
        {{ getWinPercentage(item.seasonSummary) }}%
      </template>
      <template v-slot:item.gamesBack="{item}">
        {{ getGamesBack(item) }}
      </template>
      <template v-slot:item.streak="{item}">
        {{ getStreakDisplay(item.seasonSummary) }}
      </template>
    </v-data-table>
  </v-card>
</template>

<script>
import {TEAMS_WITH_SCORING} from "@/graphql/queries/team/teamGraphQL";

export default {
  name: "LeagueStandings.vue",
  props: {
    leagueInfo: {
      type: Object,
      require: true,
      validator: function (value) {
        // Object.values(value).includes("divisions")
        return true
      }
    }
  },
  data: function () {
    return {
      headers: [
        {text: 'Division', value: 'divisionName', groupable: true},
        {text: 'Team', value: 'teamName'},
        {text: 'Record', value: 'record'},
        {text: 'Win Percentage', value: 'winPct'},
        {text: 'GB', value: 'gamesBack'},
        {text: 'PF', value: 'seasonSummary.totalPointsFor'},
        {text: 'PA', value: 'seasonSummary.totalPointsAgainst'},
        {text: 'Streak', value: 'streak', sortable: false},
      ],
      leagueYears: [2022, 2021],
      selectedYear: new Date().getFullYear(),
      teams: [],
      highestWinsPerDivisionByYear: new Map()
    }
  },
  methods: {
    getWinPercentage(teamOutcomes) {
      const totalGamesPlayed = this.getTotalGamesPlayed(teamOutcomes)
      if (totalGamesPlayed === 0) {
        return 0
      }
      return ((teamOutcomes.wins / totalGamesPlayed) * 100).toPrecision(3);
    },
    getTotalGamesPlayed(teamRecord) {
      return (teamRecord.wins + teamRecord.losses + teamRecord.ties);
    },
    getGamesBack(teamData) {
      let yearWinsPerDivision = this.highestWinsPerDivisionByYear.get(this.selectedYear)
      if (yearWinsPerDivision === undefined) {
        this.highestWinsPerDivisionByYear.set(this.selectedYear, new Map())
        yearWinsPerDivision = this.highestWinsPerDivisionByYear.get(this.selectedYear)
      }
      let highestWins = yearWinsPerDivision.get(teamData.division)
      if (highestWins === undefined) {
        const divisionTeams = this.teams.filter((team) => team.division === teamData.division)
        highestWins = 0
        divisionTeams.forEach((team) => {
          const teamWinsThatYear = team.teamScoring.filter((year) => year.year === this.selectedYear)[0]?.summary?.wins
          if (teamWinsThatYear !== undefined) {
            if (highestWins < teamWinsThatYear) {
              highestWins = teamWinsThatYear
            }
          }
        })
        this.highestWinsPerDivisionByYear.get(this.selectedYear).set(teamData.division, highestWins)
      }

      return highestWins - teamData.seasonSummary.wins
    },
    getStreakDisplay(seasonSummary) {
      const currentStreak = seasonSummary.currentStreak
      if (currentStreak === 0) {
        return currentStreak
      }
      if (currentStreak > 0) {
        return currentStreak + "W";
      }
      return Math.abs(currentStreak) + "L";
    }
  },
  computed: {
    processedTeamData() {
      if (this.teams === null || this.teams.length === 0) {
        return [];
      }

      const tableData = []

      this.teams.forEach((team) => {
        const fullSummary = team.teamScoring.filter(season => season.year === this.selectedYear)

        let summary = {
          currentStreak: 0,
          losses: 0,
          ties: 0,
          wins: 0,
          totalPointsFor: 0,
          totalPointsAgainst: 0
        }
        if (fullSummary.length === 1) {
          summary = fullSummary[0].summary
        }

        const tempTeam = {
          teamName: team.teamName,
          seasonSummary: summary,
          division: team?.division
        }

        tableData.push(tempTeam)
      })

      return tableData
    }
  },
  apollo: {
    teams: {
      query: TEAMS_WITH_SCORING,
      variables: function () {
        return {
          leagueId: this.leagueInfo.id
        }
      }
    }
  }
}
</script>

<style scoped>

</style>
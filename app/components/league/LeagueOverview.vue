<template>
  <v-card>
    <v-card-title class="justify-center" > {{leagueInfo.leagueName}} Standings</v-card-title>
    <v-data-table
      :headers=headers
      :items=items
      hide-default-footer
    >
      <template v-slot:item.record="{item}">
        {{item.outcomes.wins}}-{{item.outcomes.losses}}-{{item.outcomes.ties}}
      </template>

      <template v-slot:item.winPct="{item}">
        {{getWinPercentage(item.outcomes)}}
      </template>
      <template v-slot:item.gamesBack="{item}">
        {{getGamesBack(item)}}
      </template>
    </v-data-table>
  </v-card>
</template>

<script>
export default {
  name: "LeagueOverview.vue",
  props: {
    leagueInfo: {
      type: Object,
      require: true,
      validator: function(value) {
        Object.values(value).includes("divisions")
      }
    }
  },
  data: function() {
    return {
      headers: [
          {text:'Team', value: 'teamName'},
          {text:'Record', value: 'record'},
          {text:'Win Percentage', value: 'winPct'},
          {text:'GB', value: 'gamesBack'},
          {text:'PF', value: 'pointsFor'},
          {text:'PA', value: 'pointsAgainst'},
          {text:'Streak', value: 'currentStreak', sortable:false},
      ],
      items: [
        {
          teamName: "Rick's team",
          outcomes: {
            wins: 10,
            losses: 4,
            ties: 0
          },
          divisionName: "West",
          pointsFor: 1234,
          pointsAgainst: 1200,
          currentStreak: 8,
        },{
          teamName: "Jeff's team",
          outcomes: {
            wins: 9,
            losses: 5,
            ties: 0
          },
          divisionName: "West",
          pointsFor: 1234,
          pointsAgainst: 1200,
          currentStreak: 3,
        },
        {
          teamName: "Kyle's team",
          outcomes: {
            wins: 5,
            losses: 9,
            ties: 0
          },
          divisionName: "East",
          pointsFor: 1234,
          pointsAgainst: 1200,
          currentStreak: -3,
        },
        {
          teamName: "Ramzi's team",
          outcomes: {
            wins: 1,
            losses: 13,
            ties: 0
          },
          divisionName: "West",
          pointsFor: 1234,
          pointsAgainst: 1200,
          currentStreak: -8,
        }
      ],
    }
  },
  methods: {
    getWinPercentage(teamOutcomes) {
      return (teamOutcomes.wins/this.getTotalGamesPlayed(teamOutcomes)).toPrecision(3);
    },
    getTotalGamesPlayed(teamRecord) {
      return (teamRecord.wins + teamRecord.losses + teamRecord.ties);
    },
    getGamesBack(teamData){
      console.info("League Info: ")
      const division = this.leagueInfo.divisions; //.find(it => it.divisionName === teamData.divisionName)
      console.log("division: ",division);
      // return division.leadingWins - teamData.wins;
      return -teamData.wins;
    }
  }
}
</script>

<style scoped>

</style>
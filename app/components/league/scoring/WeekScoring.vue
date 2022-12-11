<template>
   <v-card>
     <v-col md="4" offset-md="4">
       <v-select
           v-model="selectedMatchUpNumber"
           :items="matchUps"
           :item-text="item => `${item.awayTeam} vs ${item.homeTeam}`"
           item-value="matchUpNumber"
           label="Select a match up"
           :loading="$apollo.queries.matchUps.loading"
       />
     </v-col>

     <v-skeleton-loader
         v-if="selectedMatchUpNumber !== null"
         :loading="$apollo.queries.scoringData.loading"
         type="card"
     >
       <v-container v-if="selectedMatchUpNumber !== null">
         <v-row>
           <v-col md="2" cols="0"/>
           <v-col
               v-for="team in scoringData"
               cols="12"
               md="4"
           >
             <v-list outlined>
               <v-list-item>
                 {{team.teamName}} <v-spacer/> <b>Total: {{team.totalPoints.toLocaleString()}}</b>
               </v-list-item>
               <v-list-item v-for="qb in team.lineUp.qb">
                 QB <v-spacer/> {{qb.playerName}} -- {{qb.points}}
               </v-list-item>
               <v-list-item v-for="rb in team.lineUp.rb">
                 RB <v-spacer/> {{rb.playerName}} -- {{rb.points}}
               </v-list-item>
               <v-list-item v-for="wr in team.lineUp.wr">
                 WR <v-spacer/> {{wr.playerName}} -- {{wr.points}}
               </v-list-item>
               <v-list-item v-for="te in team.lineUp.te">
                 TE <v-spacer/> {{te.playerName}} -- {{te.points}}
               </v-list-item>
               <v-list-item v-for="flex in team.lineUp.flex">
                 Flex <v-spacer/> {{flex.playerName}} -- {{flex.points}}
               </v-list-item>
               <v-list-item v-for="sFlex in team.lineUp.superFlex">
                 SuperFlex <v-spacer/> {{sFlex.playerName}} -- {{sFlex.points}}
               </v-list-item>

             </v-list>
           </v-col>
         </v-row>
         <v-row>
           <v-col md="2" cols="0"/>
           <v-col
               v-for="team in scoringData"
               cols="12"
               md="4"
           >
             <v-list
                 outlined
             >
               <v-list-item v-if="$vuetify.breakpoint.mobile">
                  Bench: &nbsp; {{team.teamName}}
               </v-list-item>
               <v-list-item v-for="player in benchPlayers(team)">
                 <v-list-item-content>
                   <v-list-item-title>
                     {{player.playerName}} -- Points Scored: {{player.points}}
                   </v-list-item-title>
                   <v-list-item-subtitle>
                     Projected: {{player.projectedPoints}}
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
  data: function() {
    return {
      matchUps: null,
      selectedMatchUpNumber: null,
      scoringData: null
    }
  },
  methods: {
    benchPlayers(team) {
      return team.roster.filter(player => !player.isInLineUp)
    }
  },
  apollo: {
    matchUps: {
      query: WEEK_SCORING_MATCH_UPS,
      update: data => data.scoring.weekMatchUps,
    },
    scoringData: {
      query: WEEK_SCORING_FOR_MATCH_UP,
      variables () {
        return {
          matchUpNumber: this.selectedMatchUpNumber
        }
      },
      update: data => data.scoring.matchUpScoring,
      skip() {
        return this.selectedMatchUpNumber === null
      }
    }
  }
}
</script>

<style scoped>

</style>
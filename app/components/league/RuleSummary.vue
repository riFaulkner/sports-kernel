<template>
  <div>
    <div class="d-flex">
      <br>
      <p>Below are the proposed rules for a best ball, contract, dynasty fantasy football league. <br>
      The intent is to create a league that is low/no maintenance during the season, but still allows for lots of fun in the off season.</p>
    </div>

    <v-expansion-panels
      v-model="panel"
      :disabled="disabled"
      multiple
    >
      <v-expansion-panel>
        <v-expansion-panel-header>Core League Elements</v-expansion-panel-header>
        <v-expansion-panel-content>
          <h3>Best Ball</h3>
          <p>Your best performing players are automatically "started". You make no start-sit decisions and there is no weekly waiver wire. The roster is a bit bigger as a result to account for injuries and bye weeks. Most standard best ball leagues don't allow trading (they are just draft exercises), but this league will have trading. Matchups are still head-to-head each week.</p>
          <h3>Contract</h3>
          <p>Dynasty version of an auction draft. The start-up draft is an auction of all players (including rookies for the first season only). You can offer a player a 1-year, 2-year, 3-year, or 4-year deal. There is a discount for longer contracts (e.g. a $10 million 1-year deal may be worth less than a $19.5 million 2-year deal). There will be a predetermined "slope" for this that is provided below. </p>
          <h3>Salary Cap</h3>
          <p>There will be a hard salary cap ($200 million/yr) and a minimum contract (1-year, $0.5 million). These loosely match the NFL cap and minimum.</p>
          <h3>Yearly Draft and Auction</h3>
          <p>In the offseason, all players not on rosters plus those that have finished their contracts will go into a free agent pool. Owners will bid in a similar way as the start-up draft. This will occur in the spring.</p>
          <p>There will be a rookie draft every year (except 2021) that follows a 1-12 order for every round based on last season's result (with a lottery for the non-playoff teams similar to NBA style). Rookie contracts are predetermined by pick and are always 4-year deals. More on that to follow also.</p>
          <p>There will be a low-intensity mid-season free agent auction (probably a silent auction) to place any free agents onto rosters for the remainder of the season.</p>
        </v-expansion-panel-content>
      </v-expansion-panel>

      <v-expansion-panel>
        <v-expansion-panel-header>Roster Composition</v-expansion-panel-header>
        <v-expansion-panel-content>
          <p>We want to introduce some new elements with the starting line-up to make it more interesting.
             Remember, you don't select a line-up each week, but the slots are filled with your best players to maximize scoring. 
             It's also heavier on flex spots to make it easier to have a valid line-up each week. The effect of the roster is to increase value for QBs and TEs (rather than doing that through scoring only).</p>
          <ul>
            <li>1 QB</li>
            <li>1 RB</li>
            <li>2 WR</li>
            <li>2 TE</li>
            <li>1 Superflex (QB/RB/WR/TE)</li>
            <li>2 Flex (RB/WR)</li>
          </ul>
          <p></p>
          <p>No kickers or defenses. Obviously. They are stupid in all leagues, but especially stupid in this format.</p>
          <p>Maximum roster size is 35 players based on starting as a 12-team league.
             However, you may not want to roster that many due to the salary cap and minimum salary.
              The maximum roster may adjust down or up based on teams in the league or how we find things work in year 1.</p>
        </v-expansion-panel-content>
      </v-expansion-panel>

      <v-expansion-panel>
        <v-expansion-panel-header>Scoring</v-expansion-panel-header>
        <v-expansion-panel-content>
          <p>Same as roster, we want some new elements to keep things interesting! Goal is to maintain balance but create much higher scores, 
            so I'm sure we'll tweak things a bit once we can dig into projections a bit more (still on to-do list). We'll use decimal scoring and scoring is additive 
            (e.g. 0.5 PPR and 0.2 points per target = 0.7 total PPR).</p>
          <v-simple-table>
                 <tr>
                    <td>Passing</td>
                    <td>Rushing</td>
                    <td>Receiving</td>
                    <td>Miscellaneous</td>
                 </tr>
          </v-simple-table>
        </v-expansion-panel-content>
      </v-expansion-panel>

      <v-expansion-panel>
        <v-expansion-panel-header>Contracts and Free Agent Auction</v-expansion-panel-header>
        <v-expansion-panel-content>
          Some content
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header>Restricted Free Agent Matching Tag</v-expansion-panel-header>
        <v-expansion-panel-content>
          Some content
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header>Salary Cap and Contract Guarantees</v-expansion-panel-header>
        <v-expansion-panel-content>
          Some content
        </v-expansion-panel-content>
      </v-expansion-panel>
       <v-expansion-panel>
        <v-expansion-panel-header>Rookie Contracts</v-expansion-panel-header>
        <v-expansion-panel-content>
          Some content
        </v-expansion-panel-content>
      </v-expansion-panel>
       <v-expansion-panel>
        <v-expansion-panel-header>Mid-Season Free Agent Auction</v-expansion-panel-header>
        <v-expansion-panel-content>
          Some content
        </v-expansion-panel-content>
      </v-expansion-panel>
       <v-expansion-panel>
        <v-expansion-panel-header>Regular Season and Playoff Structure</v-expansion-panel-header>
        <v-expansion-panel-content>
          Some content
        </v-expansion-panel-content>
      </v-expansion-panel>
    <v-expansion-panel>
      <v-expansion-panel-header>Buy-In and Payouts</v-expansion-panel-header>
        <v-expansion-panel-content>
          Some content
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
      <v-expansion-panel-header>Tanking and Collusion</v-expansion-panel-header>
        <v-expansion-panel-content>
          Some content
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
    
  </div>
</template>

<script>
export default {
  name: "RuleSummary.vue",
  data () {
    return {

    }
  }
}
</script>

<style scoped>

</style>
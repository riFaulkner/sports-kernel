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
        <v-expansion-panel-header>Core Elements (What the hell are we doing here?)</v-expansion-panel-header>
        <v-expansion-panel-content>
          <h3>Best Ball</h3>
          <p>"Your best performing players are automatically "started". You make no start-sit decisions and there is no weekly waiver wire. The roster is a bit bigger as a result to account for injuries and bye weeks. Most standard best ball leagues don't allow trading (they are just draft exercises), but this league will have trading. Matchups are still head-to-head each week."</p>
        </v-expansion-panel-content>
      </v-expansion-panel>

      <v-expansion-panel>
        <v-expansion-panel-header>Panel 2</v-expansion-panel-header>
        <v-expansion-panel-content>
          Some content
        </v-expansion-panel-content>
      </v-expansion-panel>

      <v-expansion-panel>
        <v-expansion-panel-header>Panel 3</v-expansion-panel-header>
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
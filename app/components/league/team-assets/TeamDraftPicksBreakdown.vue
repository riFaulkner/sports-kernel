<template>
  <div class="text-center">
    <h2> Draft Picks </h2>
    <v-data-table
        :headers="headers"
        :items="processedDraftPicks"
        item-key="title"
        hide-default-footer
    >
      <template v-slot:item.draft1="{item}">
        <draft-break-down-list :picks="item.picks[0]" />
      </template>
      <template v-slot:item.draft2="{item}">
        <draft-break-down-list :picks="item.picks[1]" />
      </template>
      <template v-slot:item.draft3="{item}">
        <draft-break-down-list :picks="item.picks[2]" />
      </template>
      <template v-slot:item.draft4="{item}">
        <draft-break-down-list :picks="item.picks[3]" />
      </template>
      <template v-slot:item.draft5="{item}">
        <draft-break-down-list :picks="item.picks[4]" />
      </template>
    </v-data-table>
  </div>

</template>

<script>
import DraftBreakDownList from "@/components/league/team-assets/DraftBreakDownList";

export default {
  name: "TeamDraftPicksBreakdown",
  components: {
    DraftBreakDownList
  },
  props: {
    draftPicks: {
      type: Array,
      default: function () {
        return []
      }
    }
  },
  data: function() {
    return {
      headers: [
        {text: "", value: "round"},
        {text: '2022', value: "draft1", align: "center"},
        {text: '2023', value: "draft2", align: "center"},
        {text: '2024', value: "draft3", align: "center"},
        {text: '2025', value: "draft4", align: "center"},
        {text: '2026', value: "draft5", align: "center"},
      ]
    }
  },
  computed: {
    processedDraftPicks: function() {
      if (this.draftPicks === null) {
        return [];
      }

      let tableObject = [];

      for (let roundNum = 0; roundNum < 5; roundNum++) {
        let roundDataList = [];
        const roundNumber = roundNum+1;

        this.draftPicks.forEach(draftYear => {
          roundDataList.push(draftYear.picks.filter(pick => {
            return pick.round === roundNumber
          }))
        });

        const roundString = "Round " + roundNumber;

        let roundData = {round: roundString, picks: roundDataList}
        tableObject.push(roundData)
      }

      return tableObject
    },
  }
}
</script>

<style scoped>

</style>
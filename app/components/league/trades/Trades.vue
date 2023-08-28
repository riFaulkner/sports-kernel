<template>
  <v-container fluid>
    <template v-if="validTrade">
      <v-row>
        <v-col cols="12" md="6" v-for="assets in assetsForTrade">
          <v-data-table
              :items="flattenTradeAssets(assets)"
              :headers="tradeTableHeaders"
              dense
              disabled
              hide-default-footer
              hide-default-header
          />
        </v-col>
      </v-row>
      <v-row justify="end">
        <v-btn
            text
            color="primary"
            :disabled="!validTrade"
            @click="submitTrade"
        >Submit Trade
        </v-btn>
        <br/>
        <br/>
      </v-row>
      <v-row>
        <v-divider/>
      </v-row>
    </template>

    <v-row>
      <v-col cols="12" md="6" v-for="(team, index) in teams">
        <team-asset-selector
            :team="team"
            v-model="assetsForTrade[index]"
        />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import {ALL_TEAM_ASSETS_AND_LIABILITIES} from "@/graphql/queries/team/teamGraphQL";
import TeamAssetSelector from "@/components/league/trades/TeamAssetSelector";

export default {
  name: "Trades",
  components: {TeamAssetSelector},
  props: {
    leagueId: {
      type: String,
      required: true,
    },
    initiatingTeamId: {
      type: String,
      required: true,
    },
    secondaryTeamId: {
      type: String,
      required: true,
    }
  },
  data: function () {
    return {
      teams: null,
      assetsForTrade: [
        {
          team: this.initiatingTeamId,
          contracts: [],
          picks: [],
        },
        {
          team: this.secondaryTeamId,
          contracts: [],
          picks: [],
        }
      ],
      tradeTableHeaders: [
        {text: "Asset", value: "asset"},
        {text: "Type", value: "type"}
      ]
    }
  },
  computed: {
    validTrade() {
      return !!(this.hasTradeAssets(this.assetsForTrade[0]) || this.hasTradeAssets(this.assetsForTrade[1]));
    }
  },
  methods: {
    flattenTradeAssets(assets) {
      const flattenedAssets = []

      assets.contracts ?
        assets.contracts?.forEach(contract => {
          flattenedAssets.push({asset: contract.name, type: "contract"})
        }) : 

      assets.picks?.forEach(pick => {
        flattenedAssets.push({asset: pick, type: "pick"})
      })
      return flattenedAssets
    },
    submitTrade() {
      console.log("Trade submitted!", this.assetsForTrade)
      if (this.validTrade) {

      }
    },
    hasTradeAssets(tradeAssets) {

      return tradeAssets?.contracts?.length !== 0 || tradeAssets?.picks.length !== 0
    }
  },
  apollo: {
    teams: {
      query: ALL_TEAM_ASSETS_AND_LIABILITIES,
      variables: function () {
        return {
          leagueId: this.leagueId,
          teamIds: [this.initiatingTeamId, this.secondaryTeamId]
        }
      },
      skip() {
        return this.initiatingTeamId === null || this.secondaryTeamId === null
      },
      update: data => data.team.teamIds
    }
  },

}
</script>

<style scoped>

</style>

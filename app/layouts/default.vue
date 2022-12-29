<template>
  <v-app dark>
    <v-navigation-drawer
        v-model="drawer"
        :mini-variant=false
        :clipped=true
        disable-resize-watcher
        fixed
        app
    >
      <v-list>
        <v-list-item
            v-for="(item, i) in menuItems"
            v-if="item.adminOnly === false || isUserAdmin"
            :key="i"
            :to="item.to"
            router
            exact
        >
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title v-text="item.title"/>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar
        :clipped-left=true
        fixed
        app
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"/>
      <logo class="SiteHeaderLogo"/>
      <v-toolbar-title v-if="$vuetify.breakpoint.smAndUp" v-text="title"/>

      <v-spacer/>
      <template v-if="$auth.loggedIn">
        <v-menu>
          <template v-slot:activator="{on, attrs}">
            <v-toolbar-items v-if="hasLeagues">
              <v-btn
                  v-bind="attrs"
                  v-on="on"
                  :loading="$apollo.queries.userPreferencesFetch.loading"
              >
                <template v-if="activeLeague">{{ activeLeague.leagueName }}</template>
                <template v-else></template>
                <v-icon>mdi-chevron-down</v-icon>
              </v-btn>
            </v-toolbar-items>
            <v-btn
                v-else
                color="primary"
                medium
            >
              <span>
                Join League
              </span>
            </v-btn>
          </template>
          <v-list v-if=leagues>
            <v-list-item
                v-for="league in leagues"
                :key="league.id"
                @click="changeLeague(league)">
              <v-list-item-title>{{ league.leagueName }}</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
        <v-menu
            left
            bottom
        >
          <template v-slot:activator="{ on, attrs }">
            <v-btn
                icon
                v-bind="attrs"
                v-on="on"
            >
              <v-icon>mdi-dots-vertical</v-icon>
            </v-btn>
          </template>
          <v-list>
            <v-list-item
                v-for="accountAction in userAccountActions"
                :key="accountAction.name"
                @click="accountAction.action()"
            >
              <v-list-item-title>
                <v-icon>{{ accountAction.icon }}</v-icon> &nbsp; {{ accountAction.name }}
              </v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </template>
      <template v-else>
        <v-toolbar-items>
          <v-btn icon @click="login()">
            <v-icon>mdi-account-circle</v-icon>
          </v-btn>
        </v-toolbar-items>
      </template>

      <template v-if="submenu !== null" v-slot:extension>
        <v-tabs
            v-model="activeTab"
            show-arrows
        >
          <v-tab
              v-for="tab in submenu"
              :key="tab"
          >
            {{ tab }}
          </v-tab>
        </v-tabs>
      </template>
    </v-app-bar>
    <v-main
    >
      <v-container
          fluid
      >
        <Nuxt/>
      </v-container>
    </v-main>
    <v-snackbar
        v-model="globalAlert.alert"
        text
        :timeout="globalAlert.alertPayload.timeout"
        disabled="true"
        :color="globalAlert.alertPayload.color"
    > {{ globalAlert.alertPayload.message }}
      <template v-slot:action="{ attrs }">
        <v-btn
            :color="globalAlert.alertPayload.color"
            text
            v-bind="attrs"
            @click="dismissAlert"
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>
    <v-footer
        :absolute=true
        app
    >
      <v-card-text class="text-center">&copy {{ new Date().getFullYear() }}</v-card-text>
    </v-footer>
  </v-app>
</template>

<script>

import {LEAGUE_BY_ID_QUERY} from "@/graphql/queries/league/leagueGraphQL";
import {USER_PREFERENCES_QUERY} from "@/graphql/queries/user/userGraphQl";

export default {
  data() {
    return {
      activeLeagueId: null,
      alert: false,
      alertPayload: {},
      clipped: true,
      drawer: false,
      isInitialLoad: true,
      menuItems: [
        {
          icon: 'mdi-account-group',
          title: "League Home",
          to: '/league',
          name: 'league',
          adminOnly: false,
          submenuItems: [
            'Standings', 'Scoring', 'My Team', 'Contracts Overview', 'Rules'
          ]
        },
        {
          icon: 'mdi-progress-wrench',
          title: "Draft Tools",
          to: '/draft-tools',
          name: 'draft-tools',
          adminOnly: false,
          submenuItems: null
        },
        {
          icon: 'mdi-chart-bubble',
          title: 'Strategy',
          to: '/strategy',
          name: 'strategy',
          adminOnly: false,
          submenuItems: null
        },
        {
          icon: 'mdi-account-wrench',
          title: 'Admin',
          to: '/admin',
          name: 'Admin',
          adminOnly: true,
          submenuItems: ['Add Player']
        },

      ],
      title: 'Sports Kernel',
      userAccountActions: [
        {
          name: "Preferences", icon: "mdi-account", action: () => {
          }
        },
        {
          name: "Logout", icon: "mdi-logout", action: () => {
            this.logout()
          }
        }
      ],
    }
  },
  computed: {
    activeLeague() {
      return this.$store.getters["application/getActiveLeague"];
    },
    globalAlert() {
      return this.$store.state.application.alert;
    },
    hasLeagues() {
      const leagues = this.$store.getters["user/getUserLeagues"]
      return leagues ? leagues.length > 0 : true
    },
    isUserAdmin() {
      const isAdmin = this.$store.getters["user/getIsUserAdmin"];
      return isAdmin !== null ? isAdmin : false;
    },
    leagues() {
      const leagues = this.$store.getters["user/getUserLeagues"];
      return leagues ? leagues : [];
    },
    loggedInUserId() {
      if (this.$auth.loggedIn) {
        return this.$auth.user.sub
      }
      return null
    },
    submenu() {
      const currentRoute = this.$route.name
      const menuItem =  this.menuItems.find((item) => item.name === currentRoute)

      if (menuItem === undefined || menuItem?.submenuItems?.length === 0) {
        return null
      }
      return menuItem.submenuItems
    },
    activeTab: {
      get() {
        return this.$store.state.application.activeTab;
      },
      set(newValue) {
        this.$store.dispatch("application/updateActiveTab", newValue);
      }
    },
    updateSubMenuOptions() {
      const leagueSubMenuItems = [
        'Standings','Scoring', 'My Team', 'Contracts Overview', 'Rules'
      ]

      if (this.leagues.find((league) => league.id === this.activeLeagueId).roleInLeague === 'LEAGUE_MANAGER') {
        leagueSubMenuItems.push('League Management')
      }
      this.menuItems.filter((menuItem) => menuItem.name === "league")[0].submenuItems = leagueSubMenuItems
    }
  },
  methods: {
    dismissAlert() {
      this.$store.dispatch('application/alertDismiss');
    },
    logout: async function () {
      await this.$auth.logout()
    },
    login: function () {
      this.$auth.loginWith('auth0');
    },
    changeLeague(league) {
      this.activeLeagueId = league.id
    },
  },
  apollo: {
    leagueFetch: {
      query: LEAGUE_BY_ID_QUERY,
      variables() {
        return {
          leagueId: this.activeLeagueId
        }
      },
      skip() {
        return this.activeLeagueId === null
      },
      manual: true,
      result({data, loading}) {
        if (!loading) {
          this.$store.dispatch('application/updateActiveLeague', data.league)
        }
      }
    },
    userPreferencesFetch: {
      query: USER_PREFERENCES_QUERY,
      variables() {
        return {
        userId: this.loggedInUserId
      }},
      skip() {
        if (this.$auth) {
          return !(this.$auth.loggedIn && this.$auth.user.sub)
        }
        return true
      },
      manual: true,
      result({data, loading}) {
        if (!loading) {
          this.$store.dispatch("user/setUserPreferences", data.userPreferences).then(() => {
            if (this.isInitialLoad) {
              this.activeLeagueId = this.$store.getters["user/getPreferredLeagueId"]
              this.isInitialLoad = false
            }
          })
        }
      }
    }
  }
}
</script>
<style scoped>
.SiteHeaderLogo {
  margin-bottom: -15px;
}

</style>
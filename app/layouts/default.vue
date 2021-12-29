<template>
  <v-app dark>
    <v-navigation-drawer
        v-model="drawer"
        :mini-variant=false
        :clipped=true
        fixed
        app
    >
      <v-list>
        <v-list-item
            v-for="(item, i) in menuItems"
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
            <v-toolbar-items
                v-if="isInitialized"
            >
              <v-btn
                  v-bind="attrs"
                  v-on="on"
              >
                <template v-if="activeLeague">{{activeLeague.leagueName}}</template>
                <template v-else>Select a League</template>
              <v-icon>mdi-chevron-down</v-icon></v-btn>
            </v-toolbar-items>
            <v-btn
                v-else
                color="primary"
                medium
                :loading="!isInitialized"
            >
              <span v-if="isInitialized">
                Join League
              </span>
            </v-btn>
          </template>
          <v-list v-if=leagues >
            <v-list-item
              v-for="league in leagues"
              :key="league.id"
              @click="changeLeague(league)">
              <v-list-item-title>{{league.leagueName}}</v-list-item-title>
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
              <v-list-item-title><v-icon>{{accountAction.icon}}</v-icon> &nbsp; {{ accountAction.name }}</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </template>
      <template v-else>
        <v-toolbar-items>
          <v-btn icon @click="login()"><v-icon>mdi-account-circle</v-icon></v-btn>
        </v-toolbar-items>
      </template>

      <template v-if=submenu v-slot:extension>
        <v-tabs
            v-model="activeTab"
            align-with-title
        >
          <v-tab
            v-for="tab in submenu"
            :key="tab"
          >
            {{tab}}
          </v-tab>
        </v-tabs>
      </template>
    </v-app-bar>

    <v-main>
      <v-container
        fluid
      >
        <Nuxt/>
      </v-container>
    </v-main>
    <v-footer
        :absolute=true
        app
    >
      <v-card-text class="text-center">&copy {{ new Date().getFullYear() }}</v-card-text>
    </v-footer>
  </v-app>
</template>

<script>

export default {
  data() {
    return {
      clipped: true,
      drawer: false,
      menuItems: [
        {
          icon: 'mdi-account-group',
          title: "League Home",
          to: '/league-home'
        },
        {
          icon: 'mdi-progress-wrench',
          title: "Draft Tools",
          to: '/draft-tools'
        },
        {
          icon: 'mdi-chart-bubble',
          title: 'Strategy',
          to: '/strategy'
        },
      ],
      title: 'Sports Kernel',
      userAccountActions: [
        {name: "Preferences", icon: "mdi-account", action: () => {}},
        {name: "Logout", icon: "mdi-logout", action: () => {this.logout()}}
      ],
      isInitialized: false
    }
  },
  computed: {
    activeLeague() {
      return this.$store.getters["application/getActiveLeague"];
    },
    leagues() {
      const leagues = this.$store.getters["user/getUserLeagues"];
      return leagues ? leagues : [];
    },
    submenu(){
      return this.$store.state.application.submenu;
    },
    activeTab: {
      get() {
        return this.$store.state.application.activeTab;
      },
      set(newValue) {
        this.$store.dispatch("application/updateActiveTab", newValue)
      }
    },
  },
  methods: {
    logout: async function () {
      await this.$auth.logout()
    },
    login: function () {
      this.$auth.loginWith('auth0');
    },
    changeLeague(league) {
      this.$store.dispatch('application/updateActiveLeague', league);
    },
  },
  created() {
      this.$store.dispatch('user/initializeUserPreferences', {apolloClient: this.$apollo})
      this.isInitialized = true;
  }
}
</script>
<style scoped>
.SiteHeaderLogo {
  margin-bottom: -15px;
}

</style>
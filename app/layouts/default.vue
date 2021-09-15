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
            v-for="(item, i) in items"
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
      <!--        {{$auth.user.email}}-->
      <!--        Add Avitar-->
      <template v-if="$auth.loggedIn">
        <v-menu>
          <template v-slot:activator="{on, attrs}">
            <v-toolbar-items>
              <v-btn
                  v-bind="attrs"
                  v-on="on"
              >{{leagues.activeLeague.leagueName}} <v-icon>mdi-chevron-down</v-icon></v-btn>
            </v-toolbar-items>
          </template>
          <v-list>
            <v-list-item
              v-for="league in leagues.allLeagues"
              :key="league.leagueId"
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

    </v-app-bar>
    <v-main>
      <v-container>
        <Nuxt/>
      </v-container>
    </v-main>

    <v-footer
        :absolute=true
        app
    >
      <span>&copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      clipped: true,
      drawer: false,
      items: [
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
      leagues: {
        activeLeague: {
          leagueName: "Dev League",
          leagueId: "asdfkldsfjdf"
        },
        allLeagues: [
          {
            leagueName: "Dev League",
            leagueId: "asdfkldsfjdf"
          },
          {
            leagueName: "Prod League",
            leagueId: "onsdfoiun490cnavkl"
          }
        ]


      },
      userAccountActions: [
        {name: "Preferences", icon: "mdi-account", action: () => {}},
        {name: "Logout", icon: "mdi-logout", action: () => {this.logout()}}
      ]
    }
  },
  methods: {
    logout: async function () {
      await this.$auth.logout()
    },
    login: function () {
      this.$auth.loginWith('auth0');
    },
    changeLeague(league) {
      this.leagues.activeLeague = league
    }
  }
}
</script>
<style scoped>
.SiteHeaderLogo {
  margin-bottom: -15px;
}


</style>

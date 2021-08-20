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
          <v-list-item-content>
            <v-list-item-title v-text="item.title" />
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar
      :clipped-left=true
      fixed
      app
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <logo class="SiteHeaderLogo" />
      <v-toolbar-title v-if="$vuetify.breakpoint.smAndUp" v-text="title" />

        <v-spacer />
      <div v-if="$auth.loggedIn">
        {{$auth.user.email}}
<!--        Add Avitar-->
        <v-btn @click="logout()">Logout</v-btn>
      </div>
      <div v-else>
        <v-btn @click="login()" >Login/Register</v-btn>
      </div>


    </v-app-bar>
    <v-main>
      <v-container>
        <Nuxt />
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
  data () {
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
      right: true,
      rightDrawer: false,
      title: 'Sports Kernel'
    }
  },
  methods: {
    logout: async function () {
      await this.$auth.logout()
    },
    login: function () {
      this.$auth.loginWith('auth0');
    }
  }
}
</script>
<style scoped>
  .SiteHeaderLogo {
    margin-bottom: -10px;
  }


</style>

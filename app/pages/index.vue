<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <v-card class="logo py-4 d-flex justify-center" />
      <v-card>
        <v-card-title class="headline">
          Welcome to Sports-kernel
        </v-card-title>
        <v-card-text>
          Look around and explore the application
        </v-card-text>
        <v-card-actions>
          <v-row justify="end">
            <v-col></v-col>
            <v-col md="auto">
              <v-spacer />
              <v-btn
                color="primary"
                nuxt
                to="/league"
                col=3
              >
                Continue to League
              </v-btn>
            </v-col>
            <v-col md="auto">
              <v-dialog
                v-model="dialog"
                persistent
                width="600">
                <template v-slot:activator="{ on, attrs }">
                  <v-btn
                    color="green lighten-1"
                    dark
                    v-bind="attrs"
                    v-on="on"
                  >
                    Join a League
                  </v-btn>
                </template>

                <v-card>
                  <v-card-title class="text-h5 grey darken-3">
                    Enter Access Code
                  </v-card-title>
                  <v-card-text>
                    Please copy your unique access code from your league welcome email and paste it below.
                  </v-card-text>
                  <v-container>
                     <v-text-field
                     v-model="accessCode"
                     label="Access Code">
                     </v-text-field>
                  </v-container>
                  <v-divider></v-divider>
                  
                  <v-card-actions>
                  <v-spacer></v-spacer>
                    <v-btn
                      color="green lighten-1"
                      text
                      @click="testSubmit"
                    >
                      Join League
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-dialog>
            </v-col>
          </v-row>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import {ONBOARD_USER} from "~/graphql/queries/user/userGraphQl";

export default {
    data () {
       return {
        dialog: false,
        userPreferences: null,
        accessCode: "",
      }
    },
    methods: {
      submitAccessCode(){
        const response = this.$apollo.mutate({
        mutation: ADD_USER,
        variables: {
          accessCode: this.accessCode,
        },
      }).then(result => {
          this.userPreferences = result.data.addUserToTeam
          this.loading = false
      });
      },
      testSubmit(){
        this.dialog = false
        console.log(this.accessCode)
      },
    },
  }


</script>

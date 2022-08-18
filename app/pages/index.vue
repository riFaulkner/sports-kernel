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
                width="500">
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
                  <v-card-title class="text-h5 grey darken-1">
                    Enter Access Code Here
                  </v-card-title>
                    <v-text-field :rules="rules"></v-text-field>
                  <v-divider></v-divider>
                  <v-card-actions>
                  <v-spacer></v-spacer>
                    <v-btn
                      color="green lighten-1"
                      text
                      @click="dialog = false"
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
  export default {
      data () {
        return {
          dialog: false,
        }
      },
    }
</script>

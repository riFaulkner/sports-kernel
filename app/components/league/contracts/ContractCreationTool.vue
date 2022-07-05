<template>
  <div>
    <v-stepper v-model="currentStep">
      <v-stepper-header>
        <v-stepper-step
            step="1"
            :complete="currentStep > 1"
        >
          Select Team
        </v-stepper-step>
        <v-divider/>
        <v-stepper-step
            step="2"
            :complete="currentStep > 2"
        >
          Select Player
        </v-stepper-step>
        <v-divider/>
        <v-stepper-step
            step="3"
            :complete="currentStep > 3"
        >
          Set Contract Parameters
        </v-stepper-step>
      </v-stepper-header>

      <v-stepper-items>
        <v-stepper-content step="1">
          <v-card
              class="mb-12"
              color="grey lighten-1"
              height="200px"
          >
          </v-card>
          <v-spacer/>
          <v-btn
              color="primary"
              @click="currentStep = 2"
          >
            Continue
          </v-btn>
        </v-stepper-content>
        <v-stepper-content
            step="2"
        >
          <v-card
              class="mb-12"
              color="grey lighten-1"
              height="200px"
          >
          </v-card>
          <v-btn
              color="primary"
              text
              @click="currentStep = 1"
          >
            Previous
          </v-btn>
          <v-btn
              color="primary"
              @click="currentStep = 3"
          >
            Continue
          </v-btn>
        </v-stepper-content>
        <v-stepper-content
            step="3"
        >
          <v-card
              class="mb-12"
              color="grey lighten-1"
              height="200px"
          >
          </v-card>
          <v-btn
              color="primary"
              text
              @click="currentStep = 2"
          >
            Previous
          </v-btn>
          <v-btn
              color="primary"
              @click="currentStep = 4"
          >
            Continue
          </v-btn>
        </v-stepper-content>
      </v-stepper-items>
    </v-stepper>
  </div>
</template>

<script>
export default {
  name: "ContractCreationTool",
  data: function () {
    return {
      currentStep: 1
    }
  }
}
</script>

<style scoped>

</style>
<template>

</template>

<script>
export default {
  name: "login.vue",
  mounted() {
    if (!this.$auth.loggedIn) {
      this.$auth.loginWith('auth0');
    }
  }
}
</script>

<style scoped>

</style>
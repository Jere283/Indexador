<template>
  <Bar></Bar>
  <router-view />
</template>

<style>
@import url("https://fonts.googleapis.com/css2?family=Lato&display=swap");
body {
  font-family: "Lato", sans-serif;
}
</style>

<script>
import Bar from "./components/Nav.vue";

export default {
  components: {
    Bar,
  },
};
</script>

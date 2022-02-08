<script setup>
import Accueil from "./components/Accueil.vue";
import Menu from "./components/Menu.vue";
import Team from "./components/Team.vue";
import Quiz from "./components/Quiz.vue";
import { ref } from "vue";
</script>

<template>
  <Menu />
  <RouterView />
</template>

<style>
* {
  /*applique à toute la page*/
  margin: 0; /*supprime le contour blanc*/
  font-family: "Teko"; /*ou "Prompt", j'ai importé 2 polices qui me plaisaient via google front*/
}
</style>

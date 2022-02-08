<script setup>
import Team from "./components/Team.vue";
import Menu from "./components/Menu.vue";
</script>

<template>
  <link
    href="https://fonts.googleapis.com/css2?family=Prompt:wght@200&family=Supermercado+One&family=Teko:wght@300&display=swap"
    rel="stylesheet"
  />
  <Menu />
  <Team />
</template>

<style>
* {
  /*applique à toute la page*/

  margin: 0; /*supprime le contour blanc*/
  font-family: "Teko"; /*ou "Prompt", j'ai importé 2 polices qui me plaisaient via google front*/
}
</style>

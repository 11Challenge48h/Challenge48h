<<<<<<< HEAD
<script>

axios
  .get('https://swapi.dev/api/.json')
  .then(response => (this.info = response.data.bpi))
</script>

<template>
	<div id="app">
  {{ info }}
</div>
</template>

<style>	
=======
<script setup>
import Accueil from './components/Accueil.vue' 
import Menu from './components/Menu.vue' 
</script>

<template>
  <Menu />
  <Accueil />
</template>

<style>
* {
  /*applique à toute la page*/
  margin: 0; /*supprime le contour blanc*/
  font-family: "Teko"; /*ou "Prompt", j'ai importé 2 polices qui me plaisaient via google front*/
}

>>>>>>> origin/Nathan
</style>
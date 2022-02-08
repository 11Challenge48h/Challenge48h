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
</style>
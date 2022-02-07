<script setup>
import Accueil from './components/Accueil.vue';
import Menu from './components/Menu.vue';
import {ref} from 'vue';

async function getMessage() {
	let data = [];
    const options = {
        method: "GET",
        headers: {}
    };
    const response = await fetch('http://swapi.dev/api/people', options)
    if (response.status == 200) {
        response.json().then(data => {
			let data
            console.log(data);
            message.value = data;
        })
    }
}

const message = ref([]);

</script>

<template>

<!--Menu />
<Accueil /-->
<div>
	<span v-bind:key="data.results" v-for="unmsg in message">{{unmsg}}</span>
</div>

</template>

<style>
	#app {
		font-family: Avenir, Helvetica, Arial, sans-serif;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
		text-align: center;
		color: #2c3e50;
		margin-top: 60px;
	}
</style>

<script setup>
import {ref} from 'vue';

const listPeople = ref([]);
const listPlanet = ref([]);
const listStarship = ref([]);
const listSpecies = ref([]);
const listVehicles = ref([]);
const listFilms = ref([]);

const isPeopleShowed = ref(false);
const isPlanetShowed = ref(false);
const isVehicleShowed = ref(false);
const isStarshipShowed = ref(false);
const isSpecieShowed = ref(false);
const isFilmShowed = ref(false);

const getTable = (data) => {
  let { films, created, edited, pilots, url, people, homeworld, planets, characters, starships, vehicles, species, residents, ...allowed } = data;
  return allowed;
};

async function getPeople(url) {
    const received = ref([]);
    const options = {
        method: "GET",
        headers: {}
    };
    const response = await fetch(url, options)
    if (response.status == 200) {
        const data = await response.json();
        received.value = data.results;
        received.value.forEach((element) =>
        listPeople.value.push(element));
        if(data.next != null){
            await getPeople(data.next)
        }
    }
}

async function getPlanets(url) {
    const received = ref([]);
    const options = {
        method: "GET",
        headers: {}
    };
    const response = await fetch(url, options)
    if (response.status == 200) {
        const data = await response.json();
        received.value = data.results;
        received.value.forEach((element) =>
        listPlanet.value.push(element));
        if(data.next != null){
            await getPlanets(data.next);
        }
    }
}

async function getStarships(url) {
    const received = ref([]);
    const options = {
        method: "GET",
        headers: {}
    };
    const response = await fetch(url, options)
    if (response.status == 200) {
        const data = await response.json();
        received.value = data.results;
        received.value.forEach((element) =>
        listStarship.value.push(element));
        if(data.next != null){
            await getStarships(data.next);
        }
    }
}

async function getSpecies(url) {
    const received = ref([]);
    const options = {
        method: "GET",
        headers: {}
    };
    const response = await fetch(url, options)
    if (response.status == 200) {
        const data = await response.json();
        received.value = data.results;
        received.value.forEach((element) =>
        listSpecies.value.push(element));
        if(data.next != null){
            await getSpecies(data.next);
        }
    }
}

async function getVehicles(url) {
    const received = ref([]);
    const options = {
        method: "GET",
        headers: {}
    };
    const response = await fetch(url, options)
    if (response.status == 200) {
        const data = await response.json();
        received.value = data.results;
        received.value.forEach((element) =>
        listVehicles.value.push(element));
        if(data.next != null){
            await getVehicles(data.nextz);
        }
    }
}

async function getFilms(url) {
    const received = ref([]);
    const options = {
        method: "GET",
        headers: {}
    };
    const response = await fetch(url, options)
    if (response.status == 200) {
        const data = await response.json();
        received.value = data.results;
        received.value.forEach((element) =>
        listFilms.value.push(element));
        if(data.next != null){
            await getFilms(data.next);
        }
    }
}

function isPpShowed(){
    isPeopleShowed.value = !isPeopleShowed.value;
}
function isPlShowed(){
    isPlanetShowed.value = !isPlanetShowed.value;
}
function isVShowed(){
    isVehicleShowed.value = !isVehicleShowed.value;
}
function isSpShowed(){
    isSpecieShowed.value = !isSpecieShowed.value;
}
function isSsShowed(){
    isStarshipShowed.value = !isStarshipShowed.value;
}
function isFShowed(){
    isFilmShowed.value = !isFilmShowed.value;
}

getPeople('https://swapi.dev/api/people');
getPlanets('https://swapi.dev/api/planets');
getStarships('https://swapi.dev/api/starships');
getVehicles('https://swapi.dev/api/vehicles');
getSpecies('https://swapi.dev/api/species');
getFilms('https://swapi.dev/api/films');

</script>

<template>

<div id="people">
    <div class="buttonPeople">
        <button @click="isPpShowed">Characters</button>
    </div> 
    <div id="peopleData" v-if="isPeopleShowed">
        <div v-for="people, index in listPeople">
            <div v-for="category, key in getTable(people)">
                {{ key }} : {{ category }}
            </div><br/>
        </div>
    </div>
</div>

<div id="planet">
    <div class="buttonPlanets">
        <button @click="isPlShowed">Planets</button>
    </div> 
    <div id="planetData" v-if="isPlanetShowed">
        <div v-for="planet, index in listPlanet">
            <div v-for="category, key in getTable(planet)">
                {{ key }} : {{ category }}
            </div><br/>
        </div>
    </div>
</div>

<div id="starship">
    <div class="buttonStarships">
        <button @click="isSsShowed">Starships</button>
    </div> 
    <div id="starshipData" v-if="isStarshipShowed">
        <div v-for="starship, index in listStarship">
            <div v-for="category, key in getTable(starship)">
                {{ key }} : {{ category }}
            </div><br/>
        </div>
    </div>
</div>

<div id="species">
    <div class="buttonSpecies">
        <button @click="isSpShowed">Species</button>   
    </div> 
    <div id="speciesData" v-if="isSpecieShowed">
        <div v-for="specie, index in listSpecies">
            <div v-for="category, key in getTable(specie)">
                {{ key }} : {{ category }}
            </div><br/>
        </div>
    </div>
</div>

<div id="vehicles">
    <div class="buttonVehicles">
        <button @click="isVShowed">Vehicles</button> 
    </div> 
    <div id="vehiclesData" v-if="isVehicleShowed">
        <div v-for="vehicle, index in listVehicles">
            <div v-for="category, key in getTable(vehicle)">
                {{ key }} : {{ category }}
            </div><br/>
        </div>
    </div>
</div>

<div id="films">    
    <div class="buttonFilms">
        <button @click="isFShowed">Films</button>  
    </div> 
    <div id="filmsData" v-if="isFilmShowed">
        <div v-for="film, index in listFilms">
            <div v-for="category, key in getTable(film)">
                {{ key }} : {{ category }}
            </div><br/>
        </div>
    </div>
</div>

</template>

<style scoped>
    #peopleData{
        
    }
    #planetData{
        
    }
    #starshipData{
        
    }
    #speciesData{
        
    }
    #vehiclesData{
        
    }
    #filmsData{
        
    }
    
</style>
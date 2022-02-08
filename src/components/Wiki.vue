<script setup>
import {ref} from 'vue';

const received = ref([]);
const listPeople = ref([]);
const listData = ref([]);
const listPlanet = ref([]);
const listStarship = ref([]);
const listSpecies = ref([]);
const listVehicles = ref([]);
const listFilms = ref([]);
const regex = '/https:\/\/swapi.dev\//g'

async function getData(url){
    const options = {
        method: "GET",
        headers: {}
    };
    const response = await fetch(url, options)
    if (response.status == 200) {
        const data = await response.json();
        listData.value = data.results;
        return listData.value;
    }
}

async function getPeople(url) {
    const options = {
        method: "GET",
        headers: {}
    };
    const response = await fetch(url, options)
    if (response.status == 200) {
        const data = await response.json();
        received.value = data.results;
        //console.log(data.results);
        listPlanet.value = listPlanet.value.concat(received.value);
        getUrl(listPlanet);
        //console.log(list);
        if(data.next != null){
            await getPlanet(data.next);
        }
    }
}

async function getterPeople(){
    const options = {
        method: "GET",
        headers: {}
    };
    const response = await fetch(url, options)
    if (response.status == 200) {
        const data = await response.json();

    }
}

async function getPlanet(url) {
    const options = {
        method: "GET",
        headers: {}
    };
    const response = await fetch(url, options)
    if (response.status == 200) {
        const data = await response.json();
        received.value = data.results;
        //console.log(data.results);
        listPlanet.value = listPlanet.value.concat(received.value);
        getUrl(listPlanet);
        //console.log(list);
        if(data.next != null){
            await getPlanet(data.next);
        }
    }
}

async function getStarship(url) {
    const options = {
        method: "GET",
        headers: {}
    };
    const response = await fetch(url, options)
    if (response.status == 200) {
        const data = await response.json();
        received.value = data.results;
        //console.log(data.results);
        listStarship.value = listStarship.value.concat(received.value);
        getUrl(listPeople.value);
        //console.log(list);
        if(data.next != null){
            await getPeople(data.next);
        }
    }
}

async function getSpecies(url) {
    const options = {
        method: "GET",
        headers: {}
    };
    const response = await fetch(url, options)
    if (response.status == 200) {
        const data = await response.json();
        received.value = data.results;
        //console.log(data.results);
        listSpecies.value = listSpecies.value.concat(received.value);
        getUrl(listSpecies.value);
        //console.log(list);
        if(data.next != null){
            await getSpecies(data.next);
        }
    }
}

async function getVehicles(url) {
    const options = {
        method: "GET",
        headers: {}
    };
    const response = await fetch(url, options)
    if (response.status == 200) {
        const data = await response.json();
        received.value = data.results;
        //console.log(data.results);
        listVehicles.value = listVehicles.value.concat(received.value);
        getUrl(listVehicles.value);
        //console.log(list);
        if(data.next != null){
            await getVehicles(data.next);
        }
    }
}

async function getFilms(url) {
    const options = {
        method: "GET",
        headers: {}
    };
    const response = await fetch(url, options)
    if (response.status == 200) {
        const data = await response.json();
        received.value = data.results;
        //console.log(data.results);
        listFilms.value = listFilms.value.concat(received.value);
        getUrl(listFilms.value);
        //console.log(list);
        if(data.next != null){
            await getFilms(data.next);
        }
    }
}

function getUrl(ar){
    for(const [index, element] of Object.entries(ar)){
        for(const [key, value] of Object.entries(element)){
            //console.log(value)
        }
    }
}

function addStyle(){

}

function hidePeople(){
    const targetDiv = document.getElementById("peopleData")
    if (targetDiv.style.display !== "none") {
        targetDiv.style.display = "none";
    } else {
        targetDiv.style.display = "block";
    }
}
function showPeople(){
    const targetDiv = document.getElementById("peopleData")
    if (targetDiv.style.display == "none") {
        targetDiv.style.display = "contents";
    } else {
        targetDiv.style.display = "block";
    }
}

function hidePlanet(){
    const targetDiv2 = document.getElementById("planetData")
    if (targetDiv2.style.display !== "none") {
        targetDiv2.style.display = "none";
    } else {
        targetDiv2.style.display = "block";
    }
}
function showPlanet(){
    const targetDiv2 = document.getElementById("planetData")
    if (targetDiv2.style.display == "none") {
        targetDiv2.style.display = "contents";
    } else {
        targetDiv2.style.display = "block";
    }
}

function hideStarship(){
    const targetDiv3 = document.getElementById("starshipData")
    if (targetDiv3.style.display !== "none") {
        targetDiv3.style.display = "none";
    } else {
        targetDiv3.style.display = "block";
    }
}
function showStarship(){
    const targetDiv3 = document.getElementById("starshipData")
    if (targetDiv3.style.display == "none") {
        targetDiv3.style.display = "contents";
    } else {
        targetDiv3.style.display = "block";
    }
}

function hideSpecies(){
    const targetDiv4 = document.getElementById("speciesData")
    if (targetDiv4.style.display !== "none") {
        targetDiv4.style.display = "none";
    } else {
        targetDiv4.style.display = "block";
    }
}
function showSpecies(){
    const targetDiv4 = document.getElementById("speciesData")
    if (targetDiv4.style.display == "none") {
        targetDiv4.style.display = "contents";
    } else {
        targetDiv4.style.display = "block";
    }
}

function hideVehicles(){
    const targetDiv5 = document.getElementById("vehiclesData")
    if (targetDiv5.style.display !== "none") {
        targetDiv5.style.display = "none";
    } else {
        targetDiv5.style.display = "block";
    }
}
function showVehicles(){
    const targetDiv5 = document.getElementById("vehiclesData")
    if (targetDiv5.style.display == "none") {
        targetDiv5.style.display = "contents";
    } else {
        targetDiv5.style.display = "block";
    }
}

function hideFilms(){
    const targetDiv6 = document.getElementById("filmsData")
    if (targetDiv6.style.display !== "none") {
        targetDiv6.style.display = "none";
    } else {
        targetDiv6.style.display = "block";
    }
}
function showFilms(){
    const targetDiv6 = document.getElementById("filmsData")
    if (targetDiv6.style.display == "none") {
        targetDiv6.style.display = "contents";
    } else {
        targetDiv6.style.display = "block";
    }
}
</script>

<template>

<div id="people">
    <div class="buttonPeople">
        <button @click="getPeople('http://swapi.dev/api/people')">Les personnages de Star Wars</button><br />
        <button @click="hidePeople">Afficher moins</button>
        <button @click="showPeople">Afficher plus</button>   
    </div> 
    <div id="peopleData">
        <table>
            <thead>
                <tr>
                    <td v-for="key, index in listPeople[0]">{{ index }}</td>
                </tr>
            </thead>
            <tbody>
                <tr v-for="unpp in listPeople">
                    <td v-for="item, index in unpp">
                        {{ item }}
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</div>

<div id="planet">
    <div class="buttonPlanets">
        <button @click="getPlanet('http://swapi.dev/api/planets')">Les planètes</button><br />
        <button @click="hidePlanet">Afficher moins</button>
        <button @click="showPlanet">Afficher plus</button>   
    </div> 
    <div id="planetData">
        <table>
            <thead>
                <tr>
                    <td v-for="key, index in listPlanet[0]">{{ index }}</td>
                </tr>
            </thead>
            <tbody>
                <tr v-for="unpp in listPlanet">
                    <td v-for="item, index in unpp">
                        {{ item }}
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</div>

<div id="starship">
    <div class="buttonStarships">
        <button @click="getStarship('http://swapi.dev/api/starships')">Les fameux vaisseaux</button><br />
        <button @click="hideStarship">Afficher moins</button>
        <button @click="showStarship">Afficher plus</button>   
    </div> 
    <div id="starshipData">
        <table>
            <thead>
                <tr>
                    <td v-for="key, index in listStarship[0]">{{ index }}</td>
                </tr>
            </thead>
            <tbody>
                <tr v-for="unpp in listStarship">
                    <td v-for="item, index in unpp">
                        {{ item }}
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</div>

<div id="species">
    <div class="buttonSpecies">
        <button @click="getSpecies('http://swapi.dev/api/species')">Les différentes espèces de la saga</button><br />
        <button @click="hideSpecies">Afficher moins</button>
        <button @click="showSpecies">Afficher plus</button>   
    </div> 
    <div id="speciesData">
        <table>
            <thead>
                <tr>
                    <td v-for="key, index in listSpecies[0]">{{ index }}</td>
                </tr>
            </thead>
            <tbody>
                <tr v-for="unpp in listSpecies">
                    <td v-for="item, index in unpp">
                        {{ item }}
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</div>

<div id="vehicles">
    <div class="buttonVehicles">
        <button @click="getVehicles('http://swapi.dev/api/vehicles')">Les véhicules</button><br />
        <button @click="hideVehicles">Afficher moins</button>
        <button @click="showVehicles">Afficher plus</button>   
    </div> 
    <div id="vehiclesData">
        <table>
            <thead>
                <tr>
                    <td v-for="key, index in listVehicles[0]">{{ index }}</td>
                </tr>
            </thead>
            <tbody>
                <tr v-for="unpp in listVehicles">
                    <td v-for="item, index in unpp">
                        {{ item }}
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</div>

<div id="films">
    <div class="buttonFilms">
        <button @click="getFilms('http://swapi.dev/api/films')">Les films</button><br />
        <button @click="hideFilms">Afficher moins</button>
        <button @click="showFilms">Afficher plus</button>   
    </div> 
    <div id="filmsData">
        <table>
            <thead>
                <tr>
                    <td v-for="key, index in listFilms[0]">{{ index }}</td>
                </tr>
            </thead>
            <tbody>
                <tr v-for="unpp in listFilms">
                    <td v-for="item, index in unpp">
                        {{ item }}
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</div>

</template>
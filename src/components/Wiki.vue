<script setup>
import { ref } from "vue";
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
  let {
    films,
    created,
    edited,
    pilots,
    url,
    people,
    homeworld,
    planets,
    characters,
    starships,
    vehicles,
    species,
    residents,
    ...allowed
  } = data;
  return allowed;
};
async function getPeople(url) {
  const received = ref([]);
  const options = {
    method: "GET",
    headers: {},
  };
  const response = await fetch(url, options);
  if (response.status == 200) {
    const data = await response.json();
    received.value = data.results;
    received.value.forEach((element) => listPeople.value.push(element));
    if (data.next != null) {
      await getPeople(data.next);
    }
  }
}
async function getPlanets(url) {
  const received = ref([]);
  const options = {
    method: "GET",
    headers: {},
  };
  const response = await fetch(url, options);
  if (response.status == 200) {
    const data = await response.json();
    received.value = data.results;
    received.value.forEach((element) => listPlanet.value.push(element));
    if (data.next != null) {
      await getPlanets(data.next);
    }
  }
}
async function getStarships(url) {
  const received = ref([]);
  const options = {
    method: "GET",
    headers: {},
  };
  const response = await fetch(url, options);
  if (response.status == 200) {
    const data = await response.json();
    received.value = data.results;
    received.value.forEach((element) => listStarship.value.push(element));
    if (data.next != null) {
      await getStarships(data.next);
    }
  }
}
async function getSpecies(url) {
  const received = ref([]);
  const options = {
    method: "GET",
    headers: {},
  };
  const response = await fetch(url, options);
  if (response.status == 200) {
    const data = await response.json();
    received.value = data.results;
    received.value.forEach((element) => listSpecies.value.push(element));
    if (data.next != null) {
      await getSpecies(data.next);
    }
  }
}
async function getVehicles(url) {
  const received = ref([]);
  const options = {
    method: "GET",
    headers: {},
  };
  const response = await fetch(url, options);
  if (response.status == 200) {
    const data = await response.json();
    received.value = data.results;
    received.value.forEach((element) => listVehicles.value.push(element));
    if (data.next != null) {
      await getVehicles(data.nextz);
    }
  }
}
async function getFilms(url) {
  const received = ref([]);
  const options = {
    method: "GET",
    headers: {},
  };
  const response = await fetch(url, options);
  if (response.status == 200) {
    const data = await response.json();
    received.value = data.results;
    received.value.forEach((element) => listFilms.value.push(element));
    if (data.next != null) {
      await getFilms(data.next);
    }
  }
}
function isPpShowed() {
  isPeopleShowed.value = !isPeopleShowed.value;
}
function isPlShowed() {
  isPlanetShowed.value = !isPlanetShowed.value;
}
function isVShowed() {
  isVehicleShowed.value = !isVehicleShowed.value;
}
function isSpShowed() {
  isSpecieShowed.value = !isSpecieShowed.value;
}
function isSsShowed() {
  isStarshipShowed.value = !isStarshipShowed.value;
}
function isFShowed() {
  isFilmShowed.value = !isFilmShowed.value;
}
getPeople("https://swapi.dev/api/people");
getPlanets("https://swapi.dev/api/planets");
getStarships("https://swapi.dev/api/starships");
getVehicles("https://swapi.dev/api/vehicles");
getSpecies("https://swapi.dev/api/species");
getFilms("https://swapi.dev/api/films");
</script>

<template>
  <div class="fond">
    <link
      href="https://fonts.googleapis.com/css2?family=Prompt:wght@200&family=Supermercado+One&family=Teko:wght@300&display=swap"
      rel="stylesheet"
    />
    <div class="row" id="people">
      <div class="buttonPeople">
        <button class="button" @click="isPpShowed">Characters</button>
      </div>
      <div class="Data" v-if="isPeopleShowed">
        <div v-for="(people, index) in listPeople">
          <div v-for="(category, key) in getTable(people)">
            {{ key }} : {{ category }}
          </div>
          <br />
        </div>
      </div>
    </div>

    <div class="row" id="planet">
      <div class="buttonPlanets">
        <button class="button" @click="isPlShowed">Planets</button>
      </div>
      <div class="Data" v-if="isPlanetShowed">
        <div v-for="(planet, index) in listPlanet">
          <div v-for="(category, key) in getTable(planet)">
            {{ key }} : {{ category }}
          </div>
          <br />
        </div>
      </div>
    </div>

    <div class="row" id="starship">
      <div class="buttonStarships">
        <button class="button" @click="isSsShowed">Starships</button>
      </div>
      <div class="Data" v-if="isStarshipShowed">
        <div v-for="(starship, index) in listStarship">
          <div v-for="(category, key) in getTable(starship)">
            {{ key }} : {{ category }}
          </div>
          <br />
        </div>
      </div>
    </div>

    <div class="row" id="species">
      <div class="buttonSpecies">
        <button class="button" @click="isSpShowed">Species</button>
      </div>
      <div class="Data" v-if="isSpecieShowed">
        <div v-for="(specie, index) in listSpecies">
          <div v-for="(category, key) in getTable(specie)">
            {{ key }} : {{ category }}
          </div>
          <br />
        </div>
      </div>
    </div>

    <div class="row" id="vehicles">
      <div class="buttonVehicles">
        <button class="button" @click="isVShowed">Vehicles</button>
      </div>
      <div class="Data" v-if="isVehicleShowed">
        <div v-for="(vehicle, index) in listVehicles">
          <div v-for="(category, key) in getTable(vehicle)">
            {{ key }} : {{ category }}
          </div>
          <br />
        </div>
      </div>
    </div>

    <div class="row" id="films">
      <div class="buttonFilms">
        <button class="button" @click="isFShowed">Films</button>
      </div>
      <div class="Data" v-if="isFilmShowed">
        <div v-for="(film, index) in listFilms">
          <div v-for="(category, key) in getTable(film)">
            {{ key }} : {{ category }}
          </div>
          <br />
        </div>
      </div>
    </div>
  </div>
  <div class="fond2"></div>
</template>

<style scoped>
.fond {
  background: linear-gradient(
    to bottom,
    rgba(255, 255, 255, 0.5),
    rgba(241, 225, 0, 0.671) 50%
  );
  margin-top: 50px;
  display: flex;
  justify-content: space-between;
}
.fond2 {
  background: rgba(241, 225, 0, 0.671);
  height: 70vh;
}
.button {
  margin-left: 5%;
  width: 160px;
  height: 50px;
  border: 3px solid white;
  background-color: transparent;
  font-size: 30px;
  cursor: pointer;
  transition: 0.8s ease;
  color: white;
  border-radius: 10%;
  margin-top: 10px;
}
.button:hover {
  color: white;
  background-color: rgba(241, 225, 0, 0.671);
  opacity: 0.8;
}
.row {
  flex-basis: 25%;
  background: black;
  border-radius: 10px;
  padding: 3px 20px;
  box-sizing: border-box;
  margin: 50px;
  box-shadow: 15px 15px 15px;
}
.Data {
  color: white;
}
#peopleData {
}
#planetData {
}
#starshipData {
}
#speciesData {
}
#vehiclesData {
}
#filmsData {
}
</style>

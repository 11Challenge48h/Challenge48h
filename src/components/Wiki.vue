<script setup>
import { ref } from "vue";
const received = ref([]);
const listPeople = ref([]);
const listData = ref([]);
const listPlanet = ref([]);
const listStarship = ref([]);
const listSpecies = ref([]);
const listVehicles = ref([]);
const listFilms = ref([]);
const regex = "/https:\/\/swapi.dev\//g";

const isLoading = ref(false);

async function getData(url) {
  isLoading.value = true;
  const options = {
    method: "GET",
    headers: {},
  };
  const response = await fetch(url, options);
  if (response.status == 200) {
    const data = await response.json();
    listData.value = data.results;
    isLoading.value = false;
    return listData.value;
  }
}
async function getPeople(url) {
  isLoading.value = true;
  const options = {
    method: "GET",
    headers: {},
  };
  const response = await fetch(url, options);
  if (response.status == 200) {
    const data = await response.json();
    received.value = data.results;
    //console.log(data.results);
    listPlanet.value = listPlanet.value.concat(received.value);
    getUrl(listPlanet);
    //console.log(list);
    if (data.next != null) {
      await getPlanet(data.next);
    }
    isLoading.value = false;
  }
}
async function getterPeople() {
  isLoading.value = true;
  const options = {
    method: "GET",
    headers: {},
  };
  const response = await fetch(url, options);
  if (response.status == 200) {
    const data = await response.json();
  }
}
async function getPlanet(url) {
  isLoading.value = true;
  const options = {
    method: "GET",
    headers: {},
  };
  const response = await fetch(url, options);
  if (response.status == 200) {
    const data = await response.json();
    received.value = data.results;
    //console.log(data.results);
    listPlanet.value = listPlanet.value.concat(received.value);
    getUrl(listPlanet);
    //console.log(list);
    if (data.next != null) {
      await getPlanet(data.next);
    }
    isLoading.value = false;
  }
}
async function getStarship(url) {
  isLoading.value = true;
  const options = {
    method: "GET",
    headers: {},
  };
  const response = await fetch(url, options);
  if (response.status == 200) {
    const data = await response.json();
    received.value = data.results;
    //console.log(data.results);
    listStarship.value = listStarship.value.concat(received.value);
    getUrl(listPeople.value);
    //console.log(list);
    if (data.next != null) {
      await getPeople(data.next);
    }
    isLoading.value = false;
  }
}
async function getSpecies(url) {
  isLoading.value = true;
  const options = {
    method: "GET",
    headers: {},
  };
  const response = await fetch(url, options);
  if (response.status == 200) {
    const data = await response.json();
    received.value = data.results;
    //console.log(data.results);
    listSpecies.value = listSpecies.value.concat(received.value);
    getUrl(listSpecies.value);
    //console.log(list);
    if (data.next != null) {
      await getSpecies(data.next);
    }
    isLoading.value = false;
  }
}
async function getVehicles(url) {
  isLoading.value = true;

  const options = {
    method: "GET",
    headers: {},
  };
  const response = await fetch(url, options);
  if (response.status == 200) {
    const data = await response.json();
    received.value = data.results;
    //console.log(data.results);
    listVehicles.value = listVehicles.value.concat(received.value);
    getUrl(listVehicles.value);
    //console.log(list);
    if (data.next != null) {
      await getVehicles(data.next);
    }
    isLoading.value = false;
  }
}
async function getFilms(url) {
  isLoading.value = true;
  const options = {
    method: "GET",
    headers: {},
  };
  const response = await fetch(url, options);
  if (response.status == 200) {
    const data = await response.json();
    received.value = data.results;
    //console.log(data.results);
    listFilms.value = listFilms.value.concat(received.value);
    getUrl(listFilms.value);
    //console.log(list);
    if (data.next != null) {
      await getFilms(data.next);
    }
    isLoading.value = false;
  }
}
function getUrl(ar) {
  for (const [index, element] of Object.entries(ar)) {
    for (const [key, value] of Object.entries(element)) {
      //console.log(value)
    }
  }
}
function addStyle() {}

function hidePeople() {
  const targetDiv = document.getElementById("peopleData");
  if (targetDiv.style.display !== "none") {
    targetDiv.style.display = "none";
  } else {
    targetDiv.style.display = "block";
  }
}
function showPeople() {
  const targetDiv = document.getElementById("peopleData");
  if (targetDiv.style.display == "none") {
    targetDiv.style.display = "contents";
  } else {
    targetDiv.style.display = "block";
  }
}
function hidePlanet() {
  const targetDiv2 = document.getElementById("planetData");
  if (targetDiv2.style.display !== "none") {
    targetDiv2.style.display = "none";
  } else {
    targetDiv2.style.display = "block";
  }
}
function showPlanet() {
  const targetDiv2 = document.getElementById("planetData");
  if (targetDiv2.style.display == "none") {
    targetDiv2.style.display = "contents";
  } else {
    targetDiv2.style.display = "block";
  }
}
function hideStarship() {
  const targetDiv3 = document.getElementById("starshipData");
  if (targetDiv3.style.display !== "none") {
    targetDiv3.style.display = "none";
  } else {
    targetDiv3.style.display = "block";
  }
}
function showStarship() {
  const targetDiv3 = document.getElementById("starshipData");
  if (targetDiv3.style.display == "none") {
    targetDiv3.style.display = "contents";
  } else {
    targetDiv3.style.display = "block";
  }
}
function hideSpecies() {
  const targetDiv4 = document.getElementById("speciesData");
  if (targetDiv4.style.display !== "none") {
    targetDiv4.style.display = "none";
  } else {
    targetDiv4.style.display = "block";
  }
}
function showSpecies() {
  const targetDiv4 = document.getElementById("speciesData");
  if (targetDiv4.style.display == "none") {
    targetDiv4.style.display = "contents";
  } else {
    targetDiv4.style.display = "block";
  }
}
function hideVehicles() {
  const targetDiv5 = document.getElementById("vehiclesData");
  if (targetDiv5.style.display !== "none") {
    targetDiv5.style.display = "none";
  } else {
    targetDiv5.style.display = "block";
  }
}
function showVehicles() {
  const targetDiv5 = document.getElementById("vehiclesData");
  if (targetDiv5.style.display == "none") {
    targetDiv5.style.display = "contents";
  } else {
    targetDiv5.style.display = "block";
  }
}
function hideFilms() {
  const targetDiv6 = document.getElementById("filmsData");
  if (targetDiv6.style.display !== "none") {
    targetDiv6.style.display = "none";
  } else {
    targetDiv6.style.display = "block";
  }
}
function showFilms() {
  const targetDiv6 = document.getElementById("filmsData");
  if (targetDiv6.style.display == "none") {
    targetDiv6.style.display = "contents";
  } else {
    targetDiv6.style.display = "block";
  }
}
</script>

<template>
  <link
    href="https://fonts.googleapis.com/css2?family=Prompt:wght@200&family=Supermercado+One&family=Teko:wght@300&display=swap"
    rel="stylesheet"
  />
  <div class="fond">
    <template v-if="!isLoading">
      <div id="people">
        <div class="buttonPeople">
          <button
            class="button"
            @click="getPeople('http://swapi.dev/api/people')"
          >
            characters</button
          ><br />
          <button class="button" @click="hidePeople">Afficher moins</button>
          <button class="button" @click="showPeople">Afficher plus</button>
        </div>
        <div id="peopleData">
          <table>
            <thead>
              <tr>
                <td v-for="(key, index) in listPeople[0]">{{ index }}</td>
              </tr>
            </thead>
            <tbody>
              <tr v-for="unpp in listPeople">
                <td v-for="(item, index) in unpp">
                  {{ item }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <div id="planet">
        <div class="buttonPlanets">
          <button
            class="button"
            @click="getPlanet('http://swapi.dev/api/planets')"
          >
            planets</button
          ><br />
          <button class="button" @click="hidePlanet">Afficher moins</button>
          <button class="button" @click="showPlanet">Afficher plus</button>
        </div>
        <div id="planetData">
          <table>
            <thead>
              <tr>
                <td v-for="(key, index) in listPlanet[0]">{{ index }}</td>
              </tr>
            </thead>
            <tbody>
              <tr v-for="unpp in listPlanet">
                <td v-for="(item, index) in unpp">
                  {{ item }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <div id="starship">
        <div class="buttonStarships">
          <button
            class="button"
            @click="getStarship('http://swapi.dev/api/starships')"
          >
            spaceships</button
          ><br />
          <button class="button" @click="hideStarship">Afficher moins</button>
          <button class="button" @click="showStarship">Afficher plus</button>
        </div>
        <div id="starshipData">
          <table>
            <thead>
              <tr>
                <td v-for="(key, index) in listStarship[0]">{{ index }}</td>
              </tr>
            </thead>
            <tbody>
              <tr v-for="unpp in listStarship">
                <td v-for="(item, index) in unpp">
                  {{ item }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <div id="species">
        <div class="buttonSpecies">
          <button
            class="button"
            @click="getSpecies('http://swapi.dev/api/species')"
          >
            species</button
          ><br />
          <button class="button" @click="hideSpecies">Afficher moins</button>
          <button class="button" @click="showSpecies">Afficher plus</button>
        </div>
        <div id="speciesData">
          <table>
            <thead>
              <tr>
                <td v-for="(key, index) in listSpecies[0]">{{ index }}</td>
              </tr>
            </thead>
            <tbody>
              <tr v-for="unpp in listSpecies">
                <td v-for="(item, index) in unpp">
                  {{ item }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <div id="vehicles">
        <div class="buttonVehicles">
          <button
            class="button"
            @click="getVehicles('http://swapi.dev/api/vehicles')"
          >
            vehicules</button
          ><br />
          <button class="button" @click="hideVehicles">Afficher moins</button>
          <button class="button" @click="showVehicles">Afficher plus</button>
        </div>
        <div id="vehiclesData">
          <table>
            <thead>
              <tr>
                <td v-for="(key, index) in listVehicles[0]">{{ index }}</td>
              </tr>
            </thead>
            <tbody>
              <tr v-for="unpp in listVehicles">
                <td v-for="(item, index) in unpp">
                  {{ item }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <div id="films">
        <div class="buttonFilms">
          <button
            class="button"
            @click="getFilms('http://swapi.dev/api/films')"
          >
            movies</button
          ><br />
          <button class="button" @click="hideFilms">Afficher moins</button>
          <button class="button" @click="showFilms">Afficher plus</button>
        </div>
        <div id="filmsData">
          <table>
            <thead>
              <tr>
                <td v-for="(key, index) in listFilms[0]">{{ index }}</td>
              </tr>
            </thead>
            <tbody>
              <tr v-for="unpp in listFilms">
                <td v-for="(item, index) in unpp">
                  {{ item }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </template>
    <img v-else src="../assets/R.gif" />
  </div>
</template>

<style>
.fond {
  background: linear-gradient(
    /*fond de la première page dégradé*/ to bottom,
    /* on peut aussi mettre to top, to right, to left*/ rgba(255, 255, 255, 0.5),
    rgba(241, 225, 0, 0.671) 50%
  );
  height: 100vh;
}
.button {
  margin-left: 5%;
  width: 160px;
  height: 50px;
  border: 3px solid rgb(0, 0, 0);
  background-color: transparent;
  font-size: 30px;
  cursor: pointer;
  transition: 0.8s ease;
  color: rgb(0, 0, 0);
  border-radius: 10%;
  margin-top: 10px;
}
</style>

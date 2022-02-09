<script setup>
import Proposal from "./Proposal.vue";
import { ref } from "vue";
const questions = ref([]);
/*questions.value.push([
  "L'ordre des épisodes 1,2, 3 est le suivant :",
  [
    [
      '"Un Nouvel espoir","L\'Empire contre-attaque","Le Retour du Jedi"',
      false,
    ],
    [
      '"La Menace fantôme","L\'Empire contre-attaque","La Revanche des Sith"',
      false,
    ],
    [
      '"La Menace fantôme","L\'Attaque des clones","La Revanche des Sith"',
      false,
    ],
  ],
]);
questions.value.push([
  "Dans la saga, Luke Skywalker et la princesse Leia sont :",
  [
    ["En couple", false],
    ["Frère et soeur", true],
    ["Père et fille", false],
  ],
]);
questions.value.push([
  "Quel est l'ordre 66 ?",
  [
    ["Coruscant", false],
    ["Dagobah", false],
    ["Yavin IV", false],
  ],
]);
questions.value.push([
  "Quelle est la planète d'origine de Chewbacca et de son espèce ?",
  [
    ["Tatooine", false],
    ["Kashyyk", false],
    ["Naboo", false],
  ],
]);
questions.value.push([
  "Les sabres lasers fonctionnent grâce à : ",
  [
    ["Un diamant", false],
    ["La Force", false],
    ["Un cristal", false],
  ],
]);
questions.value.push([
  "Quel est le célèbre vaisseau chasseur de l'Alliance rebelle ?",
  [
    ["Le F-Wing", false],
    ["Le X-Wing", true],
    ["Le A-Wing", false],
  ],
]);
questions.value.push([
  "De quel groupe dépend l'armée des droïdes ?",
  [
    ["L'Empire", false],
    ["La Fédération du Commerce", false],
    ["La République", false],
  ],
]);
questions.value.push([
  "Quel est le vaisseau mythique piloté par Han Solo ?",
  [
    ["L'Aigle millénium", false],
    ["Le Faucon millénium", true],
    ["Le Corbeau millénium", false],
  ],
]);
questions.value.push([
  "Le personnage de Jar Jar Binks appartient à l'espèce des : ",
  [
    ["Ewoks", false],
    ["Wookies", false],
    ["Gungans", false],
  ],
]);*/
const points = ref(0);
function shuffle(array) {
  let currentIndex = array.length,
    randomIndex;

  // While there remain elements to shuffle...
  while (currentIndex !== 0) {
    // Pick a remaining element...
    randomIndex = Math.floor(Math.random() * currentIndex);
    currentIndex--;

    // And swap it with the current element.
    [array[currentIndex], array[randomIndex]] = [
      array[randomIndex],
      array[currentIndex],
    ];
  }

  return array;
}
function createQuestion(category, obj) {
  let name = obj.name;
  let attr = "";
  let correct;

  const r = rand(1, 2);

  console.log("obj = ", obj);

  if (category === "people") {
    if (r === 1) {
      attr = "height";
      correct = obj.height;
    } else {
      attr = "birth year";
      correct = obj.birth_year;
    }
  } else if (category === "planets") {
    if (r === 1) {
      attr = "rotation period";
      correct = obj.rotation_period;
    } else {
      attr = "population";
      correct = obj.population;
    }
  } else if (category === "starships") {
    if (r === 1) {
      attr = "cost in credits";
      correct = obj.cost_in_credits;
    } else {
      attr = "crew";
      correct = obj.crew;
    }
  }

  console.log("Question : [", attr, " - ", correct, "]");

  //On génère les fausses valeurs

  return [
    "What is the " + attr + " of " + name + " ?",
    shuffle([
      [correct, true],
      [rand(0, 175), false],
      [rand(190, 20000), false],
    ]),
  ];
}
function rand(min, max) {
  min = Math.ceil(min);
  max = Math.floor(max);
  return Math.floor(Math.random() * (max - min + 1)) + min;
}
async function loadData() {
  const apiBase = "https://swapi.dev/api";

  for (let i = 0; i < 6; i++) {
    let cat = "";
    let ra = rand(1, 4);

    switch (rand(1, 3)) {
      case 1:
        cat = "people";
        break;
      case 2:
        cat = "starships";
        ra += 8;
        break;
      case 3:
        cat = "planets";
        break;
      default:
        cat = "people";
        break;
    }

    const r = await fetch(`${apiBase}/${cat}/${ra}`);

    if (r.status === 200) {
      const result = await r.json();

      questions.value.push(createQuestion(cat, result));
    }
  }
}
loadData();
function addPoint() {
  points.value += 1;
}
</script>

<template>
  <link
    href="https://fonts.googleapis.com/css2?family=Prompt:wght@200&family=Supermercado+One&family=Teko:wght@300&display=swap"
    rel="stylesheet"
  />
  <div class="fond">
    <h1>Points : {{ points }}</h1>
    <div class="questions">
      <Proposal
        v-for="question in questions"
        :question="question"
        @incPt="addPoint"
      />
    </div>
  </div>
</template>

<style>
div.navbar {
  position: unset !important;
}
.fond {
  background: linear-gradient(
    /*fond de la première page dégradé*/ to bottom,
    /* on peut aussi mettre to top, to right, to left*/ rgba(255, 255, 255, 0.5),
    rgba(241, 225, 0, 0.671) 50%
  );
}
</style>

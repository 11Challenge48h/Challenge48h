<script setup>
import Proposal from "./Proposal.vue";
import { ref } from "vue";

const questions = ref([]);

questions.value.push([
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
]);

const points = ref(0);

console.log("questions : ", questions.value);

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

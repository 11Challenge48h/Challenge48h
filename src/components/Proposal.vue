<script setup>
import { ref } from "vue";

defineProps({
  question: Array,
});

const alreadyAnswered = ref(false);

const emit = defineEmits(["incPt"]);
function getAnswerText(answer) {
  return answer[0];
}
function isAnswerCorrect(answer) {
  return answer[1] === true;
}
function getQuestionText(q) {
  return q[0];
}
function getAnswers(q) {
  return q[1];
}
function verifySelectedAnswer(ans) {
  if (ans[1] === true && alreadyAnswered.value === false) {
    alreadyAnswered.value = true;
    emit("incPt");
  }
}
function getClass(ans) {
  if (ans[1] === true) {
    return "s";
  }
  return "e";
}
</script>

<template>
  <h2>{{ getQuestionText(question) }}</h2>
  <p
    v-for="(a, i) in getAnswers(question)"
    @click="verifySelectedAnswer(a)"
    :class="getClass(a)"
  >
    {{ i }}. {{ getAnswerText(a) }}
  </p>
</template>

<style scoped>
p {
  background-color: black;
  border-radius: 10px;
  width: fit-content;

  color: white;

  margin-bottom: 15px;
  padding-left: 10px;
  padding-right: 10px;
}

p:active {
  background-color: #2c3e50;
}

.s:active {
  background-color: green;
}

.e:active {
  background-color: red;
}
</style>

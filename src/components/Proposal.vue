<script setup>
	defineProps({
		question: Array
	})
	
	const emit = defineEmits(["incPt"])
	
	function getAnswerText(answer) {
		return answer[0]
	}
	
	function isAnswerCorrect(answer) {
		return answer[1] === true
	}
	
	function getQuestionText(q) {
		return q[0]
	}
	
	function getAnswers(q) {
		return q[1]
	}
	
	function verifySelectedAnswer(ans) {
		if(ans[1] === true) {
			emit('incPt')
		}
	}
	
	function getClass(ans) {
		if(ans[1] === true) {
			return "s"
		}
		
		return "e"
	}
</script>

<template>
	<h2>{{ getQuestionText(question) }}</h2>
	<p v-for="(a, i) in getAnswers(question)" @click="verifySelectedAnswer(a)" :class="getClass(a)">{{ i }}. {{ getAnswerText(a) }}</p>
</template>

<style scoped>
	p {
		background-color: dodgerblue;
		border-radius: 10px;
		width: 400px;
		
		/*margin-left: auto;*/
		/*margin-right: auto;*/
		
		color: white;
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
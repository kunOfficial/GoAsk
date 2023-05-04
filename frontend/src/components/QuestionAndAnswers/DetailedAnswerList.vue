<template>
    <div class="DetailedAnswerList">
        <DetailedAnswer v-for="answer in answers" :answer="answer" />
    </div>
</template>

<script setup lang="ts">
import AnswerType from '@/types/Answer/AnswerType';
import DetailedAnswer from '@/components/QuestionAndAnswers/DetailedAnswer.vue'
const props = defineProps<{
    answers: Array<AnswerType>
}>()
</script>

<style scoped>

</style>
<template>
    <div class="HotQuestionsItem">
        <div class="title">
            <div class="index">{{ index }}</div>
            <h2>{{ question.title }}</h2>
        </div>
        <p>
            {{ question.description }}
        </p>
        <div class="view">
            <font-awesome-icon icon="fa-solid fa-fire" color="rgb(255, 1, 1)" />
            热度
            {{ question.view }}
        </div>
    </div>
</template>

<script setup lang="ts">
import QuestionType from '@/types/Question/QuestionType';

const props = defineProps<{
    index: number,
    question: QuestionType
}>()

</script>

<style scoped>
.HotQuestionsItem {
    margin: 20px;
    background-color: antiquewhite;
    padding: 20px;
    border-radius: 10px;
    cursor: pointer;
}

.title {
    display: flex;
    align-items: center;
}

.title>.index {
    font-size: 30px;
    color: rgb(5, 118, 214);
    margin-right: 20px;
}

.view {
    color: red;
}
</style>
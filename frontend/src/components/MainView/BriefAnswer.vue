<template>
    <div class="BriefAnswer">
        <div class="body">
            <h2 class="questionTitle">{{ answer.question_title }}</h2>
            <p><span class="questioner">{{ answer.answerer_nick_name }}:</span>
                {{ answer.content }}
            </p>
        </div>
        <LikeButton :answerID="answer.id" />
    </div>
</template>

<script setup lang="ts">
import AnswerType from '@/types/Answer/AnswerType';
import LikeButton from "@/components/Common/LikeButton.vue";


const props = defineProps<{
    answer: AnswerType
    //   bar?: number
}>()
</script>


<style scoped>
.BriefAnswer {
    background: rgb(220, 230, 243);
    /* border: 10px; */
    width: 100%;
    border-radius: 5px;
    position: relative;
    cursor: pointer;
}

.BriefAnswer:hover {
    background: rgb(206, 223, 244);
}

h2 {
    margin-left: 10px;
}

.body .questioner {
    margin-left: 10px;
    font-size: 25px;
    color: grey;
    font-weight: bolder;
}

.body p {
    padding-right: 5px;
    padding-left: 5px;
    padding-bottom: 50px;
}
</style>

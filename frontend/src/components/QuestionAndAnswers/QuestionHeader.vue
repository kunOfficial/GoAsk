<template>
    <div class="QuestionHeader">
        <UserBriefInfoBar :userID="question.questioner_id"></UserBriefInfoBar>
        <h2>{{ question.title }}</h2>
        <p>{{ question.description }}</p>
        <p class="createdAt">{{ question.updated_at }}</p>
        <button class="answer-button" @click="$emit('write-answer')">
            <font-awesome-icon icon="fa-solid fa-pen" /> 写回答
        </button>
        <hr>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from "vue"
import QuestionType from "@/types/Question/QuestionType"
import UserBriefInfoBar from "@/components/Common/UserBriefInfoBar.vue";
import client from '@/utils/client'
import error from "@/utils/info"

onMounted(async () => {
    try {
        const response = await client.getQuestion(props.question_id)
        const body = response.data
        if (response.status === 200 && body.status === 200) {
            question.value = body.data
            // console.log(body.data)
        } else {
            error.ErrorAlert(body)
        }
    } catch (err) {
        console.log(err)
    }
})

const props = defineProps<{
    question_id: number
}>()

const question = ref<QuestionType>({
    id: 0,
    questioner_id: 1,
    title: "",
    description: "",
    view: 0,
    updated_at: "",
    questioner_nick_name: "",
    answers: []
})

const emit = defineEmits<{
    (e: 'write-answer'): void
}>()


</script>

<style scoped>
.QuestionHeader {
    position: relative;
    margin-bottom: 100px;
}

p {
    padding-bottom: 20px;
}

.createdAt {
    color: grey;
}

.answer-button {
    position: absolute;
    right: 10px;
    bottom: 15px;
    width: 70px;
    height: 40px;
    background-color: rgb(45, 118, 214);
    border-style: none;
    color: aliceblue;
    border-radius: 15px;
    cursor: pointer;
    margin-left: 5px;
}
</style>
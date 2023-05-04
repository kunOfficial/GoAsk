<template>
    <div class="MyQuestions">
        <MyQuestionsItem v-for="question in questions" :question="question" @update-question="(q) => updateQuestion(q)"
            @delete-question="(qid) => deleteQuestion(qid)" />
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue"
import QuestionType from "@/types/Question/QuestionType"
import MyQuestionsItem from "@/components/Home/MyQuestion.vue"
import client from "@/utils/client"
import error from "@/utils/info"

onMounted(async () => {
    try {
        const response = await client.getUserQuestions(10, 1, parseInt(props.user_id, 10))
        const body = response.data
        if (response.status == 200 && body.status == 200) {
            questions.value = body.data
        } else {
            error.ErrorAlert(body)
        }
    } catch (err) {
        console.log(err)
    }
})


const props = defineProps<{
    user_id: string,
}>()


const questions = ref<Array<QuestionType>>([]
)

function updateQuestion(target_question: QuestionType) {
    // 由于最新修改的问题得放在最前面 (unshift)，所以先删再加
    deleteQuestion(target_question.id)
    questions.value.unshift(target_question)
}

function deleteQuestion(qid: number) {
    console.log("delete" + qid)
    questions.value = questions.value.filter((question) => { return question.id != qid })
}

</script>

<style scoped>

</style>
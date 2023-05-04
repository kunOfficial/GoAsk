<template>
    <div class="HotQuestions">
        <HotQuestionsItem v-for="(question, index) in questions" :index="index + 1" :question="question"
            @click="$router.push({ name: 'questions', params: { question_id: question.id } })">
        </HotQuestionsItem>
    </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'
import { ref, onMounted, watch } from "vue"
import QuestionType from "@/types/Question/QuestionType"
import HotQuestionsItem from "@/components/MainView/BriefQuestion.vue"
import client from "@/utils/client"
import error from "@/utils/info"

onMounted(async () => {
    updateQuestions()
})


const route = useRoute()

watch(() => route.query.filter,
    () => {
        // console.log(route.query.filter)
        updateQuestions()
    }
)


async function updateQuestions() {
    const filterString = route.query.filter as string
    if (filterString != undefined && filterString != "") {
        try {
            const response = await client.searchQuestions(filterString, 10, 1)
            const body = response.data
            if (response.status == 200 && body.status === 200) {
                questions.value = body.data
            } else {
                error.ErrorAlert(body)
            }
        } catch (err) {
            console.log(err)
        }
    } else {
        try {
            const response = await client.getQuestions(10, 1)
            const body = response.data
            if (response.status == 200 && body.status === 200) {
                questions.value = body.data
            } else {
                error.ErrorAlert(body)
            }
        } catch (err) {
            console.log(err)
        }
    }
}
const questions = ref<Array<QuestionType>>([])

</script>

<style scoped>

</style>
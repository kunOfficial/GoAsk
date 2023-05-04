<template>
    <div class="QuestonsAndAnswers">
        <QuestionHeader :question_id="question.id" @write-answer="answerFormVisible = true" />
        <AnswerEditDialog :dialogFormVisible="answerFormVisible" :question_id="question.id"
            @close="() => { answerFormVisible = false; }"
            @cancel="() => { answerFormVisible = false; }"
            @post-success="(newAnswer) => { answerFormVisible = false; question.answers.push(newAnswer) }" />
        <DetailedAnswerList :answers="question.answers"></DetailedAnswerList>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue"
import QuestionHeader from "@/components/QuestionAndAnswers/QuestionHeader.vue"
import DetailedAnswerList from "@/components/QuestionAndAnswers/DetailedAnswerList.vue"
import QuestionType from "@/types/Question/QuestionType";
import AnswerEditDialog from "@/components/Dialog/AnswerEditDialog.vue";
import client from "@/utils/client"
import error from "@/utils/info"

onMounted(async () => {
    try {
        const response = await client.getQuestionAndAnswers(parseInt(props.question_id, 10))
        const body = response.data
        if (body.status === 200) {
            // vue 里面似乎有自动将驼峰命名方式和下划线命名方式自动转换的功能
            question.value = body.data
            // console.log(body.data)
            // console.log(question.value.questioner_nick_name)
        } else {
            error.ErrorAlert(body)
        }
    } catch (err) {
        console.log(err)
    }
})

const props = defineProps<{
    // 该变量从 path 里接收参数
    question_id: string,
}>()

// 用 question 来接收后端返回的问题和回答 
const question = ref<QuestionType>({
    id: parseInt(props.question_id, 10),
    questioner_nick_name: "匿名提问者",
    questioner_id: 0,
    title: "",
    description: "",
    view: 0,
    updated_at: "",
    answers: [],
})

const answerFormVisible = ref(false)

</script>

<style scoped>

</style>

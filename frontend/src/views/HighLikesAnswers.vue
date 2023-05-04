<template>
    <div class="HighLikesAnswers">
        <BriefAnswer v-for="answer in answers" :answer="answer"
            @click="$router.push({ name: 'questions', params: { question_id: answer.question_id } })" />
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue"
import AnswerType from '@/types/Answer/AnswerType';
import BriefAnswer from '@/components/MainView/BriefAnswer.vue';
import client from "@/utils/client";
import error from "@/utils/info"

// 进入推荐回答页面时
onMounted(async () => {
    try {
        // 获取一页 answer 信息
        const response = await client.getAnswers(10, 1)
        const body = response.data
        if (body.status === 200) {
            answers.value = body.data
        } else {
            error.ErrorAlert(body)
        }
    } catch (err) {
        console.log(err)
    }
})

// answers 用于接收后端返回list
const answers = ref<Array<AnswerType>>([])


</script>

<style scoped>
.HighLikesAnswers {
    margin: 10px;
    width: 95%;
}
</style>

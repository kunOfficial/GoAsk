<template>
    <div class="MyAnswers">
        <MyAnswerItem v-for="answer in answers" :answer="answer" @delete-answer="(aid) => deleteAnswer(aid)"
            @update-answer="(newAnswer) => updateAnswer(newAnswer)" />
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue"
import AnswerType from '@/types/Answer/AnswerType';
import MyAnswerItem from "@/components/Home/MyAnswer.vue";
import client from "@/utils/client"
import error from "@/utils/info"

onMounted(async () => {
    try {
        const response = await client.getUserAnswers(10, 1, parseInt(props.user_id, 10))
        const body = response.data
        if (response.status == 200 && body.status == 200) {
            answers.value = body.data
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

const answers = ref<Array<AnswerType>>([])

function updateAnswer(target_answer: AnswerType) {
    // 由于最新修改的问题得放在最前面 (unshift)，所以先删再加
    deleteAnswer(target_answer.id)
    answers.value.unshift(target_answer)
}

function deleteAnswer(aid: number) {
    console.log("delete" + aid)
    answers.value = answers.value.filter((answer) => { return answer.id != aid })
}

</script>

<style scoped>

</style>
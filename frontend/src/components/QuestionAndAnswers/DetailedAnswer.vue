<template>
    <div class="DetailedAnswer">
        <UserBriefInfoBar :userID="answer.answerer_id"></UserBriefInfoBar>
        <p class="content">{{ answer.content }}</p>
        <p class="updatedAt">{{ answer.updated_at }}</p>
        <LikeButton :answerID="answer.id" />
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import AnswerType from '@/types/Answer/AnswerType';
import LikeButton from '@/components/Common/LikeButton.vue';
import UserBriefInfoBar from '@/components/Common/UserBriefInfoBar.vue';
import UserType from '@/types/User/UserType'
import client from "@/utils/client"
import error from "@/utils/info"

onMounted(async () => {
    try {
        const response = await client.getUserInfo(props.answer.answerer_id)
        const body = response.data
        if (response.status === 200 && body.status === 200) {
            answerer.value = body.data
            // console.log(body.data)
        } else {
            error.ErrorAlert(body)
        }
    } catch (err) {
        console.log(err)
    }
})

const props = defineProps<{
    answer: AnswerType
}>()

const answerer = ref<UserType>({
    user_id: props.answer.answerer_id,
    user_name: "",
    nick_name: props.answer.answerer_nick_name,
    about_me: "",
    profession: "",
})
</script>

<style scoped>
.DetailedAnswer {
    position: relative;
    border: 2px solid rgba(45, 118, 214, 0.9);
    padding: 20px;
    margin: 20px;
    border-radius: 10px;
}

p {
    padding-bottom: 20px;
}

.updatedAt {
    color: grey;
}
</style>
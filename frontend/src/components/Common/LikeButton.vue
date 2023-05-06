<template>
    <!-- style 动态绑定  -->
    <button class="LikeButton" :class="{ active: isLiked }" @click.stop="likeAction()">
        <font-awesome-icon icon="fa-solid fa-thumbs-up" />
        {{ likes }}
    </button>
</template>

<script setup lang="ts">
import client from '@/utils/client';
import info from '@/utils/info';
import store from '@/utils/store'
import { ElMessageBox } from 'element-plus'

import { onMounted, ref } from "vue"

const props = defineProps<{
    answerID: number
}>()

const isLiked = ref(false)
const likes = ref(0)

onMounted(async () => {
    try {
        const response = await client.getAnswerLikes(props.answerID)
        const body = response.data
        if (response.status === 200 && body.status === 200) {
            isLiked.value = body.data.is_liked
            likes.value = body.data.likes
            // console.log(body.data)
        } else {
            info.ErrorAlert(body)
        }
    } catch (err) {
        console.log(err)
    }
})

async function likeAction() {
    if (!store.isLogin()) { // 如果用户没有登录
        ElMessageBox.alert("点赞需要登录，请点击右上角登录")
    } else {
        if (isLiked.value) { // 如果是已经点赞的状态，则取消点赞
            try {
                const response = await client.cancelLikeAnswer(props.answerID)
                const body = response.data
                if (response.status === 200 && response.data.status === 200) {
                    isLiked.value = !isLiked.value
                    likes.value--
                } else {
                    info.ErrorAlert(body)
                }
            } catch (err) {
                console.log(err)
            }
        } else { // 如果是未点赞的状态，则点赞
            try {
                const response = await client.likeAnswer(props.answerID)
                const body = response.data
                if (response.status === 200 && response.data.status === 200) {
                    isLiked.value = !isLiked.value
                    likes.value++
                } else {
                    info.ErrorAlert(body)
                }
            }
            catch (err) {
                console.log(err)
            }
        }
    }
}

</script>

<style scoped>
.LikeButton {
    width: 70px;
    height: 30px;
    position: absolute;
    right: 10px;
    bottom: 10px;
    background-color: rgb(45, 118, 214);
    border-style: none;
    color: aliceblue;
    border-radius: 3px;
    cursor: pointer;
}

.LikeButton.active {
    background-color: rgb(62, 55, 141);
}

.LikeButton:hover {
    background-color: rgb(18, 61, 230);
}
</style>

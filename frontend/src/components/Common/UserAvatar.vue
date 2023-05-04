<template>
    <div class="UserAvatar">
        <el-avatar v-if="login" :src="imgURL.toString()" size="default" shape="circle" @error="() => defaultAvatar()">
            <!-- <img :src="imgURL" crossorigin='anonymous' /> -->
        </el-avatar>
        <el-avatar v-else size="default" shape="circle">登录</el-avatar>
        <!-- <el-avatar v-else size="80" :src="getAssetsImages(userID)" shape="circle"></el-avatar> -->
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import config from '@/config'
// import { useRouter, useRoute } from 'vue-router'
const props = defineProps<{
    userID: number,
    login: boolean
}>()

//TODO: 使用URL的方法在获取图片时避免了跨域问题，但src接收类型为string(不是URL), 所以老是在控制台有warning
const imgURL = ref<URL>(new URL(`/avatar/user_${props.userID}.jpeg`, config.HttpURL))
// const imgURL = ref<string>(`${config.HttpURL}/avatar/user_${props.userID}.jpeg`)

function defaultAvatar() {
    imgURL.value = new URL(`/avatar/user_0.jpeg`, config.HttpURL)
}

</script>


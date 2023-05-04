<template>
    <div class="UserBriefInfoBar">
        <UserAvatar :userID="userID" :login="true"></UserAvatar>
        <p>{{ user.nick_name }}</p>
        <div class="additionalInfo">
            <ul>
                <li>{{ user.about_me }}</li>
                <li>{{ user.profession }}</li>
            </ul>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue"
import UserType from "@/types/User/UserType"
import UserAvatar from '@/components/Common/UserAvatar.vue';
import client from '@/utils/client'
import info from '@/utils/info'

onMounted(async () => {
    try {
        const response = await client.getUserInfo(props.userID)
        const body = response.data
        if (response.status === 200 && body.status === 200) {
            user.value = body.data
            // console.log(body.data)
        } else {
            info.ErrorAlert(body)
        }
    } catch (err) {
        console.log(err)
    }
})
const props = defineProps<{
    userID: number
}>()

const user = ref<UserType>({
    user_id: 0,
    user_name: "",
    nick_name: "匿名用户",
    about_me: "",
    profession: "",
})

</script>

<style scoped>
.UserBriefInfoBar {
    max-width: 1000px;
    display: flex;
    align-items: center;
    padding-bottom: 10px;
    /* border:solid 1px; */
}

/* .circleImg {
    border-radius: 30px;
    width: 50px;
    height: 50px;
    border: solid rgb(29, 6, 161) 2px;
    cursor: pointer;
} */

p {
    font-size: 20px;
    margin-left: 20px;
}

.additionalInfo {

    right: 10px;
}

.additionalInfo>ul {
    margin-left: 20px;
    list-style: none;
}

.additionalInfo>ul li {
    color: grey;
}
</style>
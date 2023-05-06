<template>
    <div class="HeaderBar">
        <div class="logo">
            <font-awesome-icon icon="fa-solid fa-clipboard-question" size="2xl" color="rgb(45, 118, 214)" />
            <h1>Go Ask</h1>
        </div>
        <SearchBar />
        <button class="ask-button" @click="$emit('ask-question')">提问</button>
        <UserAvatar class="avatar" :login="user.login" :userID="user.user_id" @click="$emit('sign-in-or-sign-out')" />
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import SearchBar from '@/components/HeaderBar/SearchBar.vue';
import UserAvatar from '@/components/Common/UserAvatar.vue';
import BriefUserType from '@/types/User/BriefUserType'


const props = defineProps<{
    user: BriefUserType
}>()

</script>

<style scoped>
.HeaderBar {
    position: fixed;
    top: 0;
    width: 100%;
    height: 60px;
    background-color: whitesmoke;
    /* z-index 越高就在越上层 */
    z-index: 1000;
    /* 加阴影 */
    box-shadow: 2px 2px 2px 1px rgba(0, 0, 0, 0.2);
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.logo {
    margin-left: 80px;
    display: flex;
    align-items: center;
}

.avatar {
    margin-right: 20px;
    cursor: pointer;
}

h1 {
    font-family: Arial, Helvetica, sans-serif;
    color: rgb(31, 118, 224);
    font-weight: bolder;
    margin: 10px;
}

.ask-button {
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
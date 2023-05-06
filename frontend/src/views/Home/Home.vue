<template>
    <div class="Home">
        <div class="UserInfoSection">
            <UserBriefInfoBar :userID="myProfile.user_id" />
            <EditButton class="EditProfileButton" @click="dialogFormVisible = true">编辑资料</EditButton>
            <ProfileEditDialog :userProfileService="myProfile" :dialogFormVisible="dialogFormVisible"
                @update-success="(u) => { myProfile = u; dialogFormVisible = false }"
                @cancel="(dialogFormVisible = false)" @close="(dialogFormVisible = false)" />
        </div>
        <div class="selector">
            <span>
                <router-link :to="{ name: 'MyQuestions', params: { user_id: props.user_id } }">我的问题</router-link>
            </span>
            <span>
                <router-link :to="{ name: 'MyAnswers' }">我的回答</router-link>
            </span>
        </div>
        <router-view></router-view>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import client from '@/utils/client'
import UserBriefInfoBar from '@/components/Common/UserBriefInfoBar.vue';
import EditButton from '@/components/Common/EditButton.vue'
import ProfileEditDialog from '@/components/Dialog/ProfileEditDialog.vue'
import UserType from '@/types/User/UserType';
import ResponseBodyType from '@/types/ResponseBodyType'
import error from '@/utils/info';
import { useRouter } from 'vue-router'


const rouer = useRouter()

onMounted(async () => {
    try {
        const response = await client.getUserInfo(parseInt(props.user_id, 10))
        const body = response.data as ResponseBodyType<UserType>
        if (body.status == 200) {
            myProfile.value = body.data
        } else {
            error.ErrorAlert(body)
            // router.back()
        }
    } catch (err) {
        console.log(err)
    }
})

const props = defineProps<{
    user_id: string,
}>()
const dialogFormVisible = ref(false)
const myProfile = ref<UserType>({
    // userID 直接从 path的param中的id获取，如果等后端返回userID, 会有问题
    user_id: parseInt(props.user_id, 10),
    user_name: "",
    nick_name: "",
    about_me: "",
    profession: "",
})

</script>

<style scoped>
.UserInfoSection {
    padding: 30px;
    position: relative;
}

.EditProfileButton {
    position: absolute;
    right: 10px;
    bottom: 15px;
}


.selector a {
    /* display: flex; */
    padding: 15px 2px;
    color: gray;
    margin-right: 20px;
    transition: 0.7s;
    display: inline-block;
    cursor: pointer;
    text-decoration: none;
    transition: 0.5s;
}

.selector a:hover {
    transform: scale(1.2);
    color: black;
}
</style>
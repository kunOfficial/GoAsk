<template>
    <el-dialog v-model="dialogFormVisible" title="编辑资料">
        <el-form :model="userProfileService" size="large">
            <!-- <input type="file" /> -->
            <el-form-item label="头像" label-width="80px">
                <input type="file" @change="(e) => { uploadAvatar(e) }" />
            </el-form-item>
            <el-form-item label="用户名" label-width="80px">
                <el-input v-model="userProfileService.user_name" :disabled="true" />
            </el-form-item>
            <el-form-item label="昵称" label-width="80px">
                <el-input v-model="userProfileService.nick_name" autocomplete="off" maxlength="10" show-word-limit />
            </el-form-item>
            <el-form-item label="职业" label-width="80px">
                <el-input v-model="userProfileService.profession" autocomplete="off" maxlength="10" show-word-limit />
            </el-form-item>
            <el-form-item label="个人简介" label-width="80px">
                <el-input v-model="userProfileService.about_me" type="textarea" autocomplete="off" maxlength="30"
                    show-word-limit />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button size="large" @click="$emit('cancel')">取消</el-button>
                <el-button size="large" type="primary" @click="SubmitProfileUpdate()">
                    提交
                </el-button>
            </span>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref } from "vue"
import UserType from '@/types/User/UserType'
import client from '@/utils/client'
import error from '@/utils/info'

const props = defineProps<
    {
        dialogFormVisible: boolean,
        userProfileService: UserType,
    }>()

const emit = defineEmits<{
    (e: 'update-success', u: UserType): void
    (e: 'cancel'): void
}>()


const newAvaterFile = ref<File>()

// TODO:修改密码的功能
async function SubmitProfileUpdate() {
    try {
        const updateResponse = await client.updateUserInfo(props.userProfileService)
        var uploadResponse = null
        // console.log(newAvaterFile.value)
        if (newAvaterFile.value != null) {
            const fd = new FormData()
            fd.append("avatar", newAvaterFile.value, newAvaterFile.value.name)
            uploadResponse = await client.uploadAvatar(fd)
            const body = uploadResponse.data
            // console.log(body)
            if (uploadResponse.status == 200 && body.status === 200) {
                error.SuccessMessage("用户头像更新成功！")
            } else {
                error.ErrorAlert(body)
            }
        }
        const body = updateResponse.data
        if (updateResponse.status == 200 && body.status === 200) {
            emit('update-success', props.userProfileService)
            error.SuccessMessage("更新成功！")
        } else {
            error.ErrorAlert(body)
        }
    } catch (err) {
        console.log(err)
    }
}

function uploadAvatar(e: Event) {
    // console.log(e)
    const target = (<HTMLInputElement>e.target)
    const fileList = target.files
    if (fileList != null) {
        newAvaterFile.value = fileList[0]
    }
}


</script>
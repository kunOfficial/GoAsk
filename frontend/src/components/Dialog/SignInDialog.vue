<template>
    <el-dialog v-model="signInFormVisible" title="登录" width="30%">
        <h1>欢迎来到Go Ask!</h1>
        <el-form :model="userService" size="large">
            <el-form-item label="用户名" label-width="80px">
                <el-input v-model="userService.user_name" autocomplete="off" />
            </el-form-item>
            <el-form-item label="密码" label-width="80px">
                <el-input type="password" show-password v-model="userService.password" autocomplete="off" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button id="SignUp" size="large" @click="$emit('redirect-to-sign-up');">注册</el-button>
                <el-button size="large" @click="$emit('close')">取消</el-button>
                <el-button size="large" type="primary" @click="Login()">
                    登录
                </el-button>
            </span>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref, inject } from 'vue'
import UserLoginService from '@/types/User/UserSignService'
import BriefUserType from '@/types/User/BriefUserType'
import ResponseBodyType from '@/types/ResponseBodyType'
import { ElMessageBox, ElMessage } from 'element-plus'
import store from '@/utils/store'
import client from '@/utils/client'

const prop = defineProps<{
    signInFormVisible: boolean,
}>()
const emit = defineEmits<{
    (e: 'sign-in-success', user: BriefUserType): void
    (e: 'redirect-to-sign-up'): void
}>()


const userService = ref<UserLoginService>({
    user_name: "",
    password: "",
    nick_name: "",
})


async function Login() {
    try {
        let response = await client.login(userService.value)
        const body = response.data
        if (body.status != 200) {
            ElMessageBox.alert(body.msg)
        } else {
            // TODO: 其实可以为不同的data创建interface
            const data = body.data
            const user: BriefUserType = {
                user_id: data.user_id,
                nick_name: data.nick_name,
                login: true,
            }
            console.log(user, data.token)
            // 设置jwt token
            store.setToken(data.token)
            store.updateUser(user)
            emit("sign-in-success", user)
            ElMessage({
                type: 'success',
                message: `登录成功, 欢迎${body.data.nick_name}`,
            })
        }
    } catch (err) {
        console.log(err)
    }
}
</script>

<style scoped>
#SignUp {
    position: absolute;
    left: 40px;
}

h1 {
    padding: 30px;
    color: rgb(64, 158, 255);
}

p {
    margin-bottom: 50px;
}
</style>
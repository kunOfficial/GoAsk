<template>
    <el-dialog v-model="signUpFormVisible" title="登录" width="30%">
        <h1>欢迎来到Go Ask!</h1>
        <el-form :model="userService" size="large">
            <el-form-item label="用户名" label-width="80px">
                <el-input v-model="userService.user_name" placeholder="6-15位阿拉伯数字、字母、下划线的组合" autocomplete="off"
                    minlength=6 maxlength=15 />
            </el-form-item>
            <el-form-item label="密码" label-width="80px">
                <el-input type="password" show-password v-model="userService.password"
                    placeholder="6-15位阿拉伯数字、字母、下划线的组合" autocomplete="off" />
            </el-form-item>
            <el-form-item label="确认密码" label-width="80px">
                <el-input type="password" show-password v-model="recheckedPassword" placeholder="请再次输入密码"
                    autocomplete="off" />
            </el-form-item>
            <el-form-item label="昵称" label-width="80px">
                <el-input v-model="userService.nick_name" placeholder="昵称,最长10位" autocomplete="off" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button size="large" @click="$emit('close')">取消</el-button>
                <el-button size="large" type="primary" @click="SignUp()">
                    注册
                </el-button>
            </span>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref, inject } from 'vue'
import UserLoginService from '@/types/User/UserSignService'
import client from '@/utils/client'
import { ElMessage, ElMessageBox } from 'element-plus'

const prop = defineProps<{
    signUpFormVisible: boolean,
}>()

const emit = defineEmits(['sign-up-success', 'close'])
// function checkeLength(str:string, min:number, max:number) :boolean {
//     return str.length>=min && str.length<max;
// }

const userService = ref<UserLoginService>({
    user_name: "",
    password: "",
    nick_name: "",
})

const recheckedPassword = ref<string>("")


function reset() {
    userService.value.user_name = ""
    userService.value.password = ""
    recheckedPassword.value = ""
}

async function SignUp() {
    let reg: RegExp = new RegExp("[\\w]{5,15}");
    // console.log(reg)
    // console.log(reg.test(userService.value.userName))
    if (!reg.test(userService.value.user_name)) {
        ElMessageBox.alert("用户名应为6-15位阿拉伯数字、字母、下划线的组合")
    } else if (!reg.test(userService.value.password)) {
        ElMessageBox.alert("密码应为6-15位阿拉伯数字、字母、下划线的组合")
    } else if (recheckedPassword.value != userService.value.password) {
        ElMessageBox.alert("两次密码输入不一致！")
        userService.value.password = ""
        recheckedPassword.value = ""
    } else if (userService.value.nick_name.length > 10) {
        ElMessageBox.alert("昵称应在10个字符以内!")
    }
    else {
        if (userService.value.nick_name.length == 0) {
            userService.value.nick_name = "匿名用户"
        }
        try {
            let response = await client.register(userService.value)
            const body = response.data 
            if (body.status != 200) {
                ElMessageBox.alert(body.msg)
            } else {
                emit("sign-up-success")
                ElMessage({
                    type: 'success',
                    message: `注册成功!`,
                })
            }
        } catch (err) {
            console.log(err)
        }
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
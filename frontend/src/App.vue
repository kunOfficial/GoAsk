<template>
    <HeaderBar :user="user" @sign-in-or-sign-out="SignInOrSignOut()" @ask-question="askQuestion()" />
    <SignInDialog :signInFormVisible="signInFormVisible" @close="signInFormVisible = false"
        @sign-in-success="(u) => SignInSuccess(u)"
        @redirect-to-sign-up="() => { signInFormVisible = false; signUpFormVisible = true }">
    </SignInDialog>
    <SignUpDialog :signUpFormVisible="signUpFormVisible" @close="signUpFormVisible = false"
        @signUp-success="() => { signUpFormVisible = false; signInFormVisible = true }" />
    <QuestionEditDialog :dialogFormVisible="questionFormVisible" @close="(questionFormVisible = false)"
        @cancel="(questionFormVisible = false)" @success="(q) => redirectToQuestion(q.id)" />
    <main>
        <div class="navbar">
            <ul>
                <li>
                    <font-awesome-icon icon="fa-solid fa-couch" color="rgb(45, 118, 214)" />
                    <router-link to="/"> 逛一逛</router-link>
                </li>
                <li>
                    <font-awesome-icon icon="fa-solid fa-fire" color="rgb(45, 118, 214)" />
                    <router-link to="/hot-questions"> 热榜</router-link>
                </li>
                <li>
                    <font-awesome-icon icon="fa-solid fa-house" color="rgb(45, 118, 214)" />
                    <router-link :to="{ name: 'Home', params: { user_id: user.user_id } }"> 我的</router-link>
                </li>
            </ul>
        </div>
        <div class="view-container">
            <router-view />
        </div>
    </main>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import HeaderBar from '@/components/HeaderBar/HeaderBar.vue';
import BriefUserType from '@/types/User/BriefUserType'
import SignInDialog from '@/components/Dialog/SignInDialog.vue';
import SignUpDialog from '@/components/Dialog/SignUpDialog.vue';
import QuestionEditDialog from './components/Dialog/QuestionEditDialog.vue';
import { ElMessageBox, ElMessage } from 'element-plus'
import store from '@/utils/store'
import router from './utils/routes';

const signInFormVisible = ref(false)
const signUpFormVisible = ref(false)
const questionFormVisible = ref(false)

// 从localstore中初始化用户信息
const user = ref<BriefUserType>(
    store.initUser()
)
console.log(user.value)


function SignInSuccess(u: BriefUserType) {
    signInFormVisible.value = false;
    user.value = u
}

function SignInOrSignOut() {
    if (user.value.login === true) { //已经登录，则询问是否退出登录
        ElMessageBox.confirm(
            '是否退出登录？',
            'Warning',
            {
                confirmButtonText: '退出',
                cancelButtonText: '取消',
                type: 'warning',
            }
        ).then(() => { // 退出登录
            user.value.login = false
            // 删除登录信息
            store.removeUser()
            store.removeToken()
            ElMessage({
                type: 'success',
                message: '退出登录成功',
            })
        })
        // .catch(() => {
        //     ElMessage({
        //         type: 'info',
        //         message: 'Delete canceled',
        //     })
        // })
    } else { // 未登录，则登录
        signInFormVisible.value = true
    }
}

function redirectToQuestion(question_id: number) {
    questionFormVisible.value = false
    router.push({ name: 'questions', params: { question_id: question_id } })
}

function askQuestion() {
    if (!store.isLogin()) { // 如果用户没有登录
        ElMessageBox.alert("用户未登录，请点击右上角登录")
        return
    }
    questionFormVisible.value = true
}
</script >


<style scoped>
* {
    box-sizing: border-box;
    font-family: Helvetica, "PingFang SC", "Microsoft Yahei", sans-serif;
    margin: 0;
    padding: 0;
}

main {
    width: 100%;
    min-height: 100vh;
    background: rgba(193, 202, 234, 0.7);
    justify-content: center;
}


.navbar {
    width: 80%;
    margin-left: auto;
    margin-right: auto;
    margin-top: 60px;
    height: 60px;
    padding: 10px;
    display: flex;
    /*主轴对齐*/
    align-items: center;
    /*交叉轴对齐*/
    /* justify-content: space-between; */
    background-color: rgba(234, 246, 252, 0.6);
    border-bottom: 1px solid blue;
}

.navbar {
    background-color: (83, 156, 225, 0.6);
    box-shadow: 0 25px 20px -20px rgba(0, 0, 0, 0.9);
}

.navbar ul li {
    list-style: none;
    display: inline-block;
    font-size: 20px;
    margin: 0px 100px;
    cursor: pointer;
    position: relative;
    transition: 0.2s;
}

.navbar ul li>a {
    color: rgb(45, 118, 214);
    text-decoration: none;
}

/*::after 在元素后面加东西*/
.navbar ul li::after {
    content: "";
    /* background: rgba(32, 34, 52, 1); */
    background: black;
    width: 110%;
    height: 110%;
    border-radius: 30px;
    /*absolute 布局方式：是相对祖先的box来定位，父元素必须设置为relative*/
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    z-index: -1;
    opacity: 0;
    transition: 0.2s;
}

.navbar ul li:hover {
    color: whitesmoke;
}

.navbar ul li:hover::after {
    top: 50%;
    opacity: 1;
}

.view-container {
    background: rgba(244, 246, 250, 0.9);
    padding: 50px;
    width: 80%;
    margin: auto;
}
</style>

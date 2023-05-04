// views
import HotQuestions from '@/views/HotQuestions.vue';
import Home from '@/views/Home/Home.vue';
import HighLikesAnswers from '@/views/HighLikesAnswers.vue';
import QuestionAndAnswers from '@/views/QuestionAndAnswers.vue';
import MyQuestions from '@/views/Home/MyQuestions.vue';
import MyAnswers from '@/views/Home/MyAnswers.vue';
import NotFound from '@/views/NotFound.vue';
import { createRouter, createWebHistory, RouteLocationNormalized } from 'vue-router'
import { ElMessageBox } from 'element-plus'
import client from '@/utils/client'
import store from '@/utils/store'
import ResponseBodyType from "@/types/ResponseBodyType"
import UserType from "@/types/User/UserType"

async function checkLogin(to: RouteLocationNormalized) {
    const to_id = parseInt(to.params["user_id"] as string)
    try {
        const response = await client.getUserInfo(to_id)
        const body = response.data as ResponseBodyType<UserType>
        if (body.status == 1001) {
            ElMessageBox.alert("该用户不存在")
            return false
        } else if (!store.isLogin()) { // 如果用户没有登录
            ElMessageBox.alert("用户未登录，请点击右上角登录")
            store.clear()
            return false
        } else if (store.getUserID() !== to_id) {
            ElMessageBox.alert("用户信息不匹配,您无权访问")
            return false
        }
        return true;
    } catch (err) {
        console.log(err)
    }
}


const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: "/:pathMath(.*)*", name: "NotFound", component: NotFound },
        { path: "/", component: HighLikesAnswers },
        { path: "/hot-questions", name: "HotQuestions", component: HotQuestions },
        { path: "/search", name: "SearchQuestions", props: true, component: HotQuestions },
        { path: "/questions/:question_id(\\d+)", props: true, name: "questions", component: QuestionAndAnswers },
        {
            path: "/home/:user_id(\\d+)", name: "Home", component: Home,
            props: true,
            beforeEnter: [checkLogin],
            children: [
                {
                    path: "questions",
                    name: "MyQuestions",
                    component: MyQuestions,
                    props: true,
                },
                {
                    path: "answers",
                    name: "MyAnswers",
                    component: MyAnswers,
                    props: true,
                }
            ]
        },
    ]
})

export default router;
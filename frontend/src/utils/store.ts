import UserBriefInfoBar from '@/components/Common/UserBriefInfoBar.vue'
import BriefUserType from '@/types/User/BriefUserType'

// vue-cookies组件在 vue3 composition api 中已经不好使了(因为setup中无法通过this来方位app)
// 强制转换，默认上层组件一定provide了$cookies
// const $cookies = inject<VueCookies>('$cookies')
// console.log($cookies == undefined)
// $cookies.set("token", data.token) // cookies保存token

// 读取local storage 中的登录信息
function initUser(): BriefUserType {
    const user = localStorage.getItem("user")
    if (user == null) {
        const new_user: BriefUserType = {
            user_id: 0x0,
            nick_name: "游客",
            login: false,
        }
        localStorage.setItem("user", JSON.stringify(new_user))
        return new_user
    } else {
        return JSON.parse(user) as BriefUserType
    }
}

function removeUser() {
    localStorage.removeItem("user")
}

function updateUser(user: BriefUserType) {
    localStorage.setItem("user", JSON.stringify(user))
}

function isLogin(): boolean {
    const str = localStorage.getItem("user")
    if (str == null) {
        return false;
    } else {
        const user: BriefUserType = JSON.parse(str)
        return user.login
    }
}

function getUserID(): number {
    const str = localStorage.getItem("user")
    if (str == null) { // 默认返回0
        return 0
    } else {
        const user: BriefUserType = JSON.parse(str)
        return user.user_id
    }
}

// token
function setToken(token: string) {
    localStorage.setItem("token", token)
}
function removeToken() {
    localStorage.removeItem("token")
}

function getToken(): string | null {
    return localStorage.getItem("token")
}

// clear 清空local storage
function clear() {
    removeUser()
    removeToken()
}

export default { initUser, updateUser, isLogin, removeUser, removeToken, getUserID, setToken, getToken, clear }
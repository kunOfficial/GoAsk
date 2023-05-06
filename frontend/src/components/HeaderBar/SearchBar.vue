<template>
    <div class="SearchBar">
        <input type="text" @keyup.enter="search()" placeholder="在 Go Ask 上搜索问题" v-model="filterString" />
        <button @click="search()">
            <font-awesome-icon icon="fa-solid fa-magnifying-glass" size="lg" color="white" />
        </button>
    </div>
</template>


<script setup lang="ts">
import { ref } from "vue"
import { useRouter } from "vue-router"

const filterString = ref<string>("")

const router = useRouter()

function search() {
    router.push({ name: 'SearchQuestions', query: { filter: filterString.value } })
}

</script>

<style scoped>
.SearchBar {
    height: 40px;
    width: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    position: relative;
}

.SearchBar>input {
    /* 输入框与输入的文字隔开一点距离 */
    padding: 0 18px;
    border: none;
    outline: none;
    font-size: 20px;
    box-shadow: 2px 2px 6px rgba(0, 0, 0, 0.1);
    height: 40px;
    /*宽度等于父级元素*/
    width: 100%;
    border-radius: 20px;
    /*文字的颜色*/
    color: #626262;
    /* position: absolute;
  right: 0px; */

    cursor: text;
}

.SearchBar>input::placeholder {
    opacity: 0.5;
}

.SearchBar>button {
    position: absolute;
    right: 0px;
    height: 40px;
    width: 20%;
    max-width: 100px;
    border-radius: 20px;
    outline: none;
    border: none;
    background: linear-gradient(#27aae2, #3c5ef5);
    cursor: pointer;
}
</style>
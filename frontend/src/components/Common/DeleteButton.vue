<template>
    <button class="DeleteButton">
        <font-awesome-icon icon="fa-solid fa-trash-can" />
        <slot></slot>
    </button>
</template>

<script setup lang="ts">

</script>

<style scoped>
.DeleteButton {
    font-size: 15px;
    width: 100px;
    height: 40px;
    background-color: rgb(232, 22, 22);
    border-style: none;
    color: aliceblue;
    border-radius: 15px;
    cursor: pointer;

}
</style>
<template>
    <button class="EditButton">
        <font-awesome-icon icon="fa-solid fa-pen" />
        <slot></slot>
    </button>
</template>

<script setup lang="ts">

</script>

<style scoped>
.EditButton {
    font-size: 15px;
    width: 100px;
    height: 40px;
    background-color: rgb(45, 118, 214);
    border-style: none;
    color: aliceblue;
    border-radius: 15px;
    cursor: pointer;

}
</style>
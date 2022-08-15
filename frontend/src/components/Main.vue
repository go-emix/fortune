<script setup>
import {RouterView} from 'vue-router'
import Top from './Top.vue'
import Left from './Left.vue'</script>

<template class="app">
    <Top></Top>
    <Left></Left>
    <router-view></router-view>
</template>

<style>
.app {
    width: 100%;
    height: 100%;
}
</style>

<script setup>
import {useI18n} from 'vue-i18n'
import Left from './Left.vue'
import Top from './Top.vue'

const {t} = useI18n()

</script>

<template>
    <Top></Top>
    <Left></Left>
    <div>{{ t("dashboard") }}</div>
</template>

<style scoped>

</style>
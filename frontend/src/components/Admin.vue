<script setup>
import {useI18n} from 'vue-i18n'
import Left from './Left.vue'

const {t} = useI18n()

</script>

<template>
    <Left></Left>
    <div>{{ t("admin") }}</div>
</template>

<style scoped>

</style>
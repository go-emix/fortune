<script setup lang="ts">
import {useI18n} from 'vue-i18n'

const {t} = useI18n()

</script>

<template>
    <el-empty :description=" t('not_found')" class="empty"/>
</template>

<style scoped>
.empty {
    width: 100%;
    height: 100%;
}
</style>
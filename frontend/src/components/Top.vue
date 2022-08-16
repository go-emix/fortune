<script setup>
import {useI18n} from 'vue-i18n'
import {langs} from '../pkg/i18n'
import {exit, getState, saveState} from "../pkg/session"
import {ref} from "vue";

const {t, locale} = useI18n()

function syncI18n(va) {
    saveState({i18n: va})
}

const admin = ref(getState().admin)

</script>

<template>
    <div class="temp">
        <span class="span">{{ admin.nickname ? admin.nickname : admin.username }}</span>
        <el-select v-model="locale" @change="syncI18n" class="select">
            <el-option
                v-for="item in langs"
                :key="item.i18n"
                :label="item.lang"
                :value="item.i18n">
            </el-option>
        </el-select>
        <el-link class="exit" @click="exit" :underline="false">{{ t("exit") }}
        </el-link>
    </div>
</template>

<style scoped>

.temp {
    margin-top: 15px;
}

.select {
    margin-left: 30px;
    width: 70px;
}

.exit {
    color: #b9abab;
    margin-left: 30px;
}

.span {
    color: #66b1ff;
}

</style>
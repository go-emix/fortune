import {createApp} from 'vue'
import ElementPlus from 'element-plus';
import 'element-plus/theme-chalk/index.css';
import App from './App.vue'
import i18n from './pkg/i18n'

const app = createApp(App)
app.use(ElementPlus)
app.use(i18n)
app.mount('#app')

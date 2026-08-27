import { createApp } from 'vue'
import { createPinia } from 'pinia'
import './assets/index.css'

import App from './App.vue'
import router from './router'
import { useClient } from './composables/useClient'

const app = createApp(App)

app.use(createPinia())
app.use(router)

useClient();


app.mount('#app')

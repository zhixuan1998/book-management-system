import './assets/styles/main.scss'

import { createApp } from 'vue'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

import App from './App.vue'
import router from './router'
import components from './components'
import repositories from '../repositories'
import { messages } from '../models/businessMessages'

const app = createApp(App)

app.use(router)
app.component('FontAwesomeIcon', FontAwesomeIcon);

app.provide('messages', messages)
app.provide('repositories', repositories)

Object.keys(components).forEach((name) => app.component(name, components[name]))

app.mount('#app')

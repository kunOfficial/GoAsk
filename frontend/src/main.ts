import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from '@/utils/routes'

// font-awesome
import { library } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'; // <-- TS2307 squiggles
import { faThumbsUp, faMagnifyingGlass, faClipboardQuestion, faPen, faHouse, faFire, faCouch, faTrashCan } from '@fortawesome/free-solid-svg-icons';

// element-plus
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';








// axios.defaults.baseURL = config.HttpURL;
// axios.defaults.headers.common['Authorization'] = config.JWTToken;
// axios.defaults.headers.post['Content-Type'] = 'application/form';
// axios.defaults.withCredentials = true;




library.add(faThumbsUp, faMagnifyingGlass, faClipboardQuestion, faPen, faHouse, faFire, faCouch, faTrashCan);


// createApp(App).use(router).mount('#app')


const app = createApp(App)
app.component('font-awesome-icon', FontAwesomeIcon)
app.use(ElementPlus, { size: 'small', zIndex: 3000 })
app.use(router)
app.mount('#app')


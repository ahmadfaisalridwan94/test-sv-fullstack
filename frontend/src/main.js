import { createApp } from 'vue'
import './style.css'
import router from "./router";
import App from './App.vue'
import './services/api/axios'
import VueSweetalert2 from 'vue-sweetalert2';
import 'sweetalert2/dist/sweetalert2.min.css';

const app = createApp(App);
app.use(VueSweetalert2);
app.use(router).mount("#app");

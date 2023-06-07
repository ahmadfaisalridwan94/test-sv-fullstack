import { createRouter, createWebHistory } from "vue-router";
import DefaultLayout from "../components/DefaultLayout.vue";
import Home from "../views/Home.vue";
import CreatePost from "../views/CreatePost.vue";
import EditPost from "../views/EditPost.vue";
import Preview from "../views/Preview.vue";

const routes = [
    {
        path: '/',        
        component: DefaultLayout,
        children: [
            {
                path: '/',
                name: 'home',
                component: Home,
            },           
            {
                path: '/post/edit/:id',
                name: 'edit_post',
                component: EditPost,
            },
            {
                path: '/post/create',
                name: 'create_post',
                component: CreatePost,
            },
            {
                path: '/preview',
                name: 'preview',
                component: Preview,
            },
            
        ]
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
  });
  
export default router;
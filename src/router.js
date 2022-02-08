import { createRouter, createWebHistory } from "vue-router" ;
import Accueil from './components/Accueil.vue';
import Team from './components/Team.vue';
//import Quiz from './components/Quiz.vue';
//import Wiki from './components/Wiki.vue';

const routes = [
    {
        path : '/',
        name : 'Accueil',
        component : Accueil,
    },
    {
        path : '/Team',
        name : 'Team',
        component : Team,
    },
/*    {
        path : '/Quiz',
        name : 'Quiz',
        component : Quiz,
    },
    {
        path : '/Wiki',
        name : 'Wiki',
        component : Wiki,
    }*/
];

const router = createRouter({
    history : createWebHistory(),
    routes,
});

export default router ;
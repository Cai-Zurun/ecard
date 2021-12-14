import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);

export default new Router({
    // mode: 'history',
    routes: [
        {
            path: '/login',
            name: 'Login',
            component: () => import('../components/page/Login.vue')
        },
        {
            path: '/',
            name: 'Home',
            component: () => import('../components/common/Home.vue'),
            children: [
                {
                    path: '/dashboard',
                    name: 'Dashboard',
                    component: () => import('../components/page/Dashboard.vue'),
                },
                {
                    path: '/student-info',
                    name: 'StudentInfo',
                    component: () => import('../components/page/StudentInfo.vue'),
                },
                {
                    path: '/alarm-info',
                    name: 'AlarmInfo',
                    component: () => import('../components/page/AlarmInfo.vue'),
                },
                {
                    path: '/location',
                    name: 'Location',
                    component: () => import('../components/page/Location.vue'),
                },
                {
                    path: '/track',
                    name: 'Track',
                    component: () => import('../components/page/Track.vue'),
                },
                {
                    path: '/virtual-device',
                    name: 'VirtualDevice',
                    component: () => import('../components/page/VirtualDevice.vue'),
                },
                {
                    path: '/late-situation',
                    name: 'LateSituation',
                    component: () => import('../components/page/LateSituation.vue'),
                },
            ]
        },
    ]
})
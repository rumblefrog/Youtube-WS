import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Videos from './views/Videos.vue';
import Selection from './views/Selection.vue';

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/Videos',
      name: 'Videos',
      component: Videos,
    },
    {
      path: '/Selection',
      name: 'Selection',
      component: Selection,
    },
    // {
    //   path: '/about',
    //   name: 'about',
    //   // route level code-splitting
    //   // this generates a separate chunk (about.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
    // }
  ]
})

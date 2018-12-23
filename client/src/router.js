import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import NominationsList from '@/views/NominationsList.vue'
import NominationBestHunter from '@/views/NominationBestHunter.vue'
import NominationCreativeСlass from '@/views/NominationCreativeСlass.vue'
import NominationMenthor from '@/views/NominationMenthor.vue'
import NominationCertificate from '@/views/NominationCertificate.vue'
import NominationThanks from '@/views/NominationThanks.vue'
import NominationTransformation from '@/views/NominationTransformation.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/nominations',
      name: 'nominations',
      component: NominationsList
    },
    {
      path: '/nomination-best-hunter',
      name: 'nomination-best-hunter',
      component: NominationBestHunter
    },
    {
      path: '/nomination-creative-class',
      name: 'nomination-creative-class',
      component: NominationCreativeСlass
    },
    {
      path: '/nomination-menthor',
      name: 'nomination-menthor',
      component: NominationMenthor
    },
    {
      path: '/nomination-certificate',
      name: 'nomination-certificate',
      component: NominationCertificate
    },
    {
      path: '/nomination-thanks',
      name: 'nomination-thanks',
      component: NominationThanks
    },
    {
      path: '/nomination-transformation',
      name: 'nomination-transformation',
      component: NominationTransformation
    }
  ]
})

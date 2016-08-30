import Vue from 'vue'
import Router from 'vue-router'
import Resource from 'vue-resource'

import App from './App'
import Stpinyin from './components/Stpinyin.vue'
import Unidecode from './components/Unidecode.vue'
import Slugify from './components/Slugify.vue'
import Pinyin from './components/Pinyin.vue'

Vue.use(Router)
Vue.use(Resource)
var router = new Router({
  linkActiveClass: 'active'
})

router.map({
  '/stpinyin': {
    name: 'stpinyin',
    component: Stpinyin
  },
  '/unidecode': {
    name: 'unidecode',
    component: Unidecode
  },
  '/slugify': {
    name: 'slugify',
    component: Slugify
  },
  '/pinyin': {
    name: 'pinyin',
    component: Pinyin
  }
})
router.alias({
  '/': '/pinyin'
})
router.redirect({
  '*': '/pinyin'
})

router.start(App, 'app')
